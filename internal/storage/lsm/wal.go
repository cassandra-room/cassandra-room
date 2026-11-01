package lsm

import (
	"fmt"
	"io"

	"github.com/cockroachdb/pebble/record"
	"github.com/cockroachdb/pebble/vfs"
)

// walWriter appends batches to one WAL segment, framed by pebble/record.
type walWriter struct {
	num  uint64
	f    vfs.File
	w    *record.Writer
	size int64 // bytes written so far, approximately
}

func createWAL(fs vfs.FS, dir string, num uint64) (*walWriter, error) {
	f, err := fs.Create(fs.PathJoin(dir, walFileName(num)), vfs.WriteCategoryUnspecified)
	if err != nil {
		return nil, err
	}
	return &walWriter{num: num, f: f, w: record.NewWriter(f)}, nil
}

// add appends one batch. The record reaches the OS before add returns; with
// sync it is also fsynced.
func (w *walWriter) add(data []byte, sync bool) error {
	off, err := w.w.WriteRecord(data)
	if err != nil {
		return err
	}
	w.size = off + int64(len(data))
	if err := w.w.Flush(); err != nil {
		return err
	}
	if sync {
		return w.f.Sync()
	}
	return nil
}

// close syncs and closes the WAL.
func (w *walWriter) close() error {
	err := w.w.Close()
	if serr := w.f.Sync(); err == nil {
		err = serr
	}
	if cerr := w.f.Close(); err == nil {
		err = cerr
	}
	return err
}

// replayWAL calls fn for every complete record in the WAL. A torn or zeroed
// tail (from a crash during a write that was never acknowledged as synced)
// ends the replay without error.
func replayWAL(fs vfs.FS, path string, fn func(data []byte) error) error {
	f, err := fs.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r := record.NewReader(f, 0)
	for {
		rr, err := r.Next()
		if err == io.EOF || record.IsInvalidRecord(err) {
			return nil
		}
		if err != nil {
			return fmt.Errorf("lsm: reading WAL %s: %w", path, err)
		}
		data, err := io.ReadAll(rr)
		if record.IsInvalidRecord(err) {
			return nil
		}
		if err != nil {
			return fmt.Errorf("lsm: reading WAL %s: %w", path, err)
		}
		if err := fn(data); err != nil {
			return err
		}
	}
}
