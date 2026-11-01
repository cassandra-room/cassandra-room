package lsm

import (
	"bytes"
	"context"

	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/objstorage/objstorageprovider"
	"github.com/cockroachdb/pebble/sstable"
	"github.com/cockroachdb/pebble/sstable/tablefilters/bloom"
	"github.com/cockroachdb/pebble/vfs"
)

func makeWriterOptions(o *Options) sstable.WriterOptions {
	return sstable.WriterOptions{
		TableFormat:  tableFormat,
		BlockSize:    o.BlockSize,
		FilterPolicy: bloom.FilterPolicy(o.BloomBitsPerKey),
		// Value blocks only pay off for MVCC-style older versions; without
		// them the reader needs no value-block provider.
		DisableValueBlocks: true,
	}
}

func makeReaderOptions() sstable.ReaderOptions {
	var o sstable.ReaderOptions
	o.FilterDecoders = append(o.FilterDecoders, bloom.Decoder)
	return o
}

// writeTable writes the memtable to sstable num. It returns nil if the
// memtable is empty.
func writeTable(fs vfs.FS, dir string, num uint64, wo sstable.WriterOptions, mem *memTable) (*fileMeta, error) {
	n := mem.first()
	if n == nil {
		return nil, nil
	}
	f, err := fs.Create(fs.PathJoin(dir, tableFileName(num)), vfs.WriteCategoryUnspecified)
	if err != nil {
		return nil, err
	}
	w := sstable.NewRawWriter(objstorageprovider.NewFileWritable(f), wo)
	for ; n != nil; n = n.next() {
		k := pebble.InternalKey{UserKey: n.key, Trailer: pebble.InternalKeyTrailer(n.trailer)}
		if err := w.Add(k, n.value, false, sstable.KVMeta{}); err != nil {
			w.Close()
			return nil, err
		}
	}
	// Close finishes the table: it syncs and closes the file.
	if err := w.Close(); err != nil {
		return nil, err
	}
	md, err := w.Metadata()
	if err != nil {
		return nil, err
	}
	return &fileMeta{
		num:      num,
		size:     md.Size,
		smallest: md.SmallestPoint.Clone(),
		largest:  md.LargestPoint.Clone(),
		minSeq:   uint64(md.SeqNums.Low),
		maxSeq:   uint64(md.SeqNums.High),
	}, nil
}

// tableReader is an open sstable.
type tableReader struct {
	r *sstable.Reader
}

func openTable(fs vfs.FS, dir string, num uint64, ro sstable.ReaderOptions) (*tableReader, error) {
	path := fs.PathJoin(dir, tableFileName(num))
	f, err := fs.Open(path, vfs.RandomReadsOption)
	if err != nil {
		return nil, err
	}
	readable, err := objstorageprovider.NewFileReadable(f, fs, objstorageprovider.NewReadaheadConfig(), path)
	if err != nil {
		f.Close()
		return nil, err
	}
	r, err := sstable.NewReader(context.Background(), readable, ro)
	if err != nil {
		readable.Close()
		return nil, err
	}
	return &tableReader{r: r}, nil
}

func (t *tableReader) close() error { return t.r.Close() }

// get returns the newest entry for key in the table. The Bloom filter is
// checked first. The returned value is a copy.
func (t *tableReader) get(key []byte) (kind pebble.InternalKeyKind, value []byte, found bool, err error) {
	it, err := t.r.NewPointIter(context.Background(), sstable.IterOptions{
		FilterBlockSizeLimit: sstable.AlwaysUseFilterBlock,
		Env:                  sstable.NoReadEnv,
		ReaderProvider:       sstable.MakeTrivialReaderProvider(t.r),
	})
	if err != nil {
		return 0, nil, false, err
	}
	kv := it.SeekPrefixGE(key, key, 0)
	if kv != nil && bytes.Equal(kv.K.UserKey, key) {
		kind = kv.K.Kind()
		if kind == pebble.InternalKeyKindSet {
			v, _, verr := kv.Value(nil)
			if verr != nil {
				err = verr
			} else {
				value = append([]byte(nil), v...)
			}
		}
		found = err == nil
	}
	if err == nil {
		err = it.Error()
	}
	// Close returns the iterator to a pool; it must not be used afterwards.
	if cerr := it.Close(); err == nil {
		err = cerr
	}
	return kind, value, found, err
}

// mayContain reports whether key is inside the table's key range.
func (f *fileMeta) mayContain(key []byte) bool {
	return bytes.Compare(key, f.smallest.UserKey) >= 0 && bytes.Compare(key, f.largest.UserKey) <= 0
}
