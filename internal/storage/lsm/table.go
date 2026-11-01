package lsm

import (
	"slices"
	"sync/atomic"

	"github.com/cockroachdb/pebble"
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
	mems  []*memTable // newest first
	files []*fileMeta // newest first
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

// Get returns a copy of the newest value of key, or ErrNotFound.
//
// Get must not run concurrently with DropTable for the same table; the
// schema layer serializes them.
func (t *Table) Get(key []byte) ([]byte, error) {
	if t.dropped.Load() {
		return nil, ErrUnknownTable
	}
	// Load the state before the sequence number: every sstable in the state
	// holds only entries at or below the sequence number loaded after it.
	s := t.state.Load()
	seq := t.e.visibleSeq.Load()
	for _, m := range s.mems {
		if n := m.get(key, seq); n != nil {
			if n.kind() == pebble.InternalKeyKindDelete {
				return nil, ErrNotFound
			}
			return append([]byte(nil), n.value...), nil
		}
	}
	for _, f := range s.files {
		if !f.mayContain(key) {
			continue
		}
		kind, value, found, err := f.reader.get(key)
		if err != nil {
			return nil, err
		}
		if found {
			if kind == pebble.InternalKeyKindDelete {
				return nil, ErrNotFound
			}
			return value, nil
		}
	}
	return nil, ErrNotFound
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
