package lsm

import (
	"cmp"
	"math"
	"slices"
	"sync/atomic"
)

// Table is one table's store: its memtables and sstables, in its own
// directory. Writes go through Engine.Apply; reads go through the Table.
type Table struct {
	e   *Engine
	id  TableID
	dir string

	// state is what readers use; it is published while holding Engine.mu.
	state   atomic.Pointer[readState]
	dropped atomic.Bool

	// Guarded by Engine.mu.
	mem        *memTable
	imm        []*immMemTable // oldest first
	files      []*fileMeta    // newest first (by maxSeq)
	flushedSeq uint64
}

// immMemTable is a memtable waiting to be flushed.
type immMemTable struct {
	mem *memTable
	// boundary is the last sequence number assigned when the memtable was
	// rotated: every entry of the table at or below it is in this memtable or
	// an older one.
	boundary uint64
}

// readState is an immutable view of a table's memtables and sstables.
type readState struct {
	mems []*memTable // newest first
	// files are ordered by maximum timestamp, highest first, so a read can
	// stop once a row deletion shadows everything that remains.
	files []*fileMeta
}

// ID returns the table's ID.
func (t *Table) ID() TableID { return t.id }

// publishLocked publishes the table's current memtables and files to
// readers. Engine.mu must be held.
func (t *Table) publishLocked() {
	s := &readState{
		mems:  make([]*memTable, 0, len(t.imm)+1),
		files: slices.Clone(t.files),
	}
	slices.SortStableFunc(s.files, func(a, b *fileMeta) int { return cmp.Compare(b.maxTS, a.maxTS) })
	s.mems = append(s.mems, t.mem)
	for i := len(t.imm) - 1; i >= 0; i-- {
		s.mems = append(s.mems, t.imm[i].mem)
	}
	t.state.Store(s)
}

// memtableBytesLocked returns the memory used by the table's memtables.
func (t *Table) memtableBytesLocked() int64 {
	n := t.mem.approximateSize()
	for _, im := range t.imm {
		n += im.mem.approximateSize()
	}
	return n
}

// Get returns the row stored under key: the merge of all its fragments in
// memtables and sstables (Cassandra's single-partition read). It returns
// ErrNotFound if the key has no fragments. The row may hold only
// tombstones; interpreting it is up to the caller (see the row package).
//
// Get must not run concurrently with DropTable for the same table; the
// schema layer serializes them.
func (t *Table) Get(key []byte) ([]byte, error) {
	if t.dropped.Load() {
		return nil, ErrUnknownTable
	}
	merger := t.e.opts.Merger
	// Load the state before the sequence number: every sstable in the state
	// holds only entries at or below the sequence number loaded after it.
	s := t.state.Load()
	seq := t.e.visibleSeq.Load()
	var acc []byte
	found := false
	for _, m := range s.mems {
		var ok bool
		var err error
		if acc, ok, err = m.getMerged(merger, key, seq, acc); err != nil {
			return nil, err
		}
		found = found || ok
	}
	deletedAt := int64(math.MinInt64)
	if found {
		var err error
		if deletedAt, err = merger.DeletedAt(acc); err != nil {
			return nil, err
		}
	}
	for _, f := range s.files {
		// Files are ordered by maximum timestamp: once a row deletion
		// shadows this file's newest data, it shadows every remaining file.
		if f.maxTS <= deletedAt {
			break
		}
		if !f.mayContain(key) {
			continue
		}
		v, ok, err := f.reader.get(key)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}
		if acc, err = mergeInto(merger, acc, v); err != nil {
			return nil, err
		}
		found = true
		if deletedAt, err = merger.DeletedAt(acc); err != nil {
			return nil, err
		}
	}
	if !found {
		return nil, ErrNotFound
	}
	return acc, nil
}

// Flush flushes the table's mutable memtable and waits until all of the
// table's memtables are written to sstables.
func (t *Table) Flush() error {
	e := t.e
	e.commitMu.Lock()
	e.mu.Lock()
	defer e.mu.Unlock()
	if err := e.checkUsableLocked(); err != nil {
		e.commitMu.Unlock()
		return err
	}
	if t.dropped.Load() {
		e.commitMu.Unlock()
		return ErrUnknownTable
	}
	e.rotateLocked(t)
	e.commitMu.Unlock()
	for len(t.imm) > 0 && !t.dropped.Load() {
		if err := e.checkUsableLocked(); err != nil {
			return err
		}
		e.mu.cond.Wait()
	}
	if t.dropped.Load() {
		return ErrUnknownTable
	}
	return nil
}
