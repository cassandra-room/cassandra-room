package lsm

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"sync"
	"sync/atomic"

	"github.com/cockroachdb/pebble/sstable"
	"github.com/cockroachdb/pebble/vfs"
)

// Engine is a node's storage: one commit log (WAL) and memtable memory
// budget shared by all tables, and a store per table (Cassandra's
// architecture). It is safe for concurrent use.
//
// Lock order: commitMu, then manifestMu, then mu.
type Engine struct {
	opts       Options
	fs         vfs.FS
	dir        string
	tablesDir  string
	lock       io.Closer
	writerOpts sstable.WriterOptions
	readerOpts sstable.ReaderOptions

	// commitMu serializes writes: WAL appends and memtable inserts happen in
	// sequence-number order, and every memtable has a single writer.
	commitMu sync.Mutex
	lastSeq  uint64 // guarded by commitMu
	// visibleSeq is the highest sequence number whose batch is fully applied.
	// Readers ignore memtable entries above it, so batches appear atomically.
	visibleSeq atomic.Uint64
	// memUsed is the memory used by all memtables, including those waiting
	// to be flushed.
	memUsed atomic.Int64

	// manifestMu serializes manifest writes, which happen without mu held.
	manifestMu sync.Mutex

	mu struct {
		sync.Mutex
		cond   sync.Cond
		tables map[TableID]*Table
		wal    *walWriter
		// walSegs are the live WAL segments, oldest first; the last one is
		// being written.
		walSegs     []walSegment
		manifest    *manifestWriter
		nextFileNum uint64
		// flushQueue holds rotated memtables in rotation order.
		flushQueue []flushJob
		// immBytes is the memory used by memtables in flushQueue.
		immBytes int64
		flushing bool
		bgErr    error
		closed   bool
	}
	flushWG sync.WaitGroup
}

type walSegment struct {
	num  uint64
	size int64 // final size; the current segment's size is in mu.wal
}

type flushJob struct {
	t   *Table
	imm *immMemTable
}

// Open opens or creates the engine in dir.
func Open(dir string, opts *Options) (*Engine, error) {
	var o Options
	if opts != nil {
		o = *opts
	}
	o.ensureDefaults()
	e := &Engine{
		opts:       o,
		fs:         o.FS,
		dir:        dir,
		tablesDir:  o.FS.PathJoin(dir, tablesDirName),
		writerOpts: makeWriterOptions(&o),
		readerOpts: makeReaderOptions(),
	}
	e.mu.cond.L = &e.mu.Mutex
	e.mu.tables = map[TableID]*Table{}
	if err := e.fs.MkdirAll(e.tablesDir, 0755); err != nil {
		return nil, err
	}
	// Make the directories durable: each one's entry lives in its parent.
	for _, d := range []string{e.fs.PathDir(dir), dir} {
		if err := syncDir(e.fs, d); err != nil {
			return nil, err
		}
	}
	lock, err := e.fs.Lock(e.fs.PathJoin(dir, "LOCK"))
	if err != nil {
		return nil, fmt.Errorf("lsm: locking %s: %w", dir, err)
	}
	e.lock = lock
	if err := e.recover(); err != nil {
		e.closeFiles()
		lock.Close()
		return nil, err
	}
	return e, nil
}

func (e *Engine) newTable(id TableID) *Table {
	t := &Table{e: e, id: id, dir: e.fs.PathJoin(e.tablesDir, id.dirName())}
	t.mem = newMemTable(&e.memUsed)
	return t
}

// recover loads the manifest, opens every table's sstables, replays the WAL
// into memtables, flushes them, and starts a new manifest and WAL segment.
func (e *Engine) recover() error {
	state, err := e.loadState()
	if err != nil {
		return err
	}
	nextFileNum := max(state.nextFileNum, 1)
	lastSeq := state.lastSeq
	bump := func(num uint64) { nextFileNum = max(nextFileNum, num+1) }

	for id, tm := range state.tables {
		t := e.newTable(id)
		t.flushedSeq = tm.flushedSeq
		for _, f := range tm.files {
			r, err := openTable(e.fs, t.dir, f.num, e.readerOpts)
			if err != nil {
				return fmt.Errorf("lsm: opening table %d file %d: %w", id, f.num, err)
			}
			f.reader = r
			t.files = append(t.files, f)
			lastSeq = max(lastSeq, f.maxSeq)
			bump(f.num)
		}
		sortFilesNewestFirst(t.files)
		e.mu.tables[id] = t
	}

	// File numbers are global; make sure new ones exceed every file on disk,
	// including files no manifest records.
	rootNames, err := e.fs.List(e.dir)
	if err != nil {
		return err
	}
	var wals []uint64
	for _, name := range rootNames {
		if kind, num, ok := parseFileName(name); ok {
			bump(num)
			if kind == "wal" && num >= state.minWAL {
				wals = append(wals, num)
			}
		}
	}
	slices.Sort(wals)
	tableDirs, err := e.fs.List(e.tablesDir)
	if err != nil {
		return err
	}
	for _, d := range tableDirs {
		names, _ := e.fs.List(e.fs.PathJoin(e.tablesDir, d))
		for _, name := range names {
			if _, num, ok := parseFileName(name); ok {
				bump(num)
			}
		}
	}

	// Replay the WAL. An entry is applied only if its table exists and the
	// table has not flushed past it.
	recovered := map[TableID]*memTable{}
	var scratch atomic.Int64
	for _, num := range wals {
		err := replayWAL(e.fs, e.fs.PathJoin(e.dir, walFileName(num)), func(data []byte) error {
			br, err := newBatchReader(data)
			if err != nil {
				return err
			}
			for seq := br.seq; ; seq++ {
				ent, ok, err := br.next()
				if err != nil {
					return err
				}
				if !ok {
					break
				}
				lastSeq = max(lastSeq, seq)
				t := e.mu.tables[ent.table]
				if t == nil || seq <= t.flushedSeq {
					continue
				}
				m := recovered[ent.table]
				if m == nil {
					m = newMemTable(&scratch)
					recovered[ent.table] = m
				}
				m.add(ent.key, makeTrailer(seq, ent.kind), ent.value)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Flush what was recovered; afterwards every table has flushed up to
	// lastSeq, so the old WAL segments are no longer needed.
	for id, m := range recovered {
		t := e.mu.tables[id]
		meta, err := writeTable(e.fs, t.dir, nextFileNum, e.writerOpts, m)
		if err != nil {
			return err
		}
		nextFileNum++
		if meta == nil {
			continue
		}
		if meta.reader, err = openTable(e.fs, t.dir, meta.num, e.readerOpts); err != nil {
			return err
		}
		if err := syncDir(e.fs, t.dir); err != nil {
			return err
		}
		t.files = append([]*fileMeta{meta}, t.files...)
	}

	walNum := nextFileNum
	nextFileNum++
	wal, err := createWAL(e.fs, e.dir, walNum)
	if err != nil {
		return err
	}
	e.mu.wal = wal
	e.mu.walSegs = []walSegment{{num: walNum}}
	manifestNum := nextFileNum
	nextFileNum++
	snapshot := &versionEdit{nextFileNum: nextFileNum, lastSeq: lastSeq, minWAL: walNum}
	for id, t := range e.mu.tables {
		t.flushedSeq = lastSeq
		snapshot.tables = append(snapshot.tables, &tableEdit{
			id: id, created: true, added: slices.Clone(t.files), flushedSeq: lastSeq,
		})
	}
	m, err := createManifest(e.fs, e.dir, manifestNum, snapshot)
	if err != nil {
		return err
	}
	e.mu.manifest = m
	e.mu.nextFileNum = nextFileNum
	e.lastSeq = lastSeq
	e.visibleSeq.Store(lastSeq)
	for _, t := range e.mu.tables {
		t.publishLocked()
	}

	// Remove what the new manifest no longer references: old WAL segments and
	// manifests, directories of dropped tables, and sstables from flushes that
	// crashed before they were recorded.
	for _, name := range rootNames {
		kind, num, ok := parseFileName(name)
		if (ok && ((kind == "wal" && num < walNum) || (kind == "manifest" && num != manifestNum))) ||
			name == currentFileName+".tmp" {
			_ = e.fs.Remove(e.fs.PathJoin(e.dir, name))
		}
	}
	for _, d := range tableDirs {
		var t *Table
		for _, tt := range e.mu.tables {
			if tt.id.dirName() == d {
				t = tt
			}
		}
		path := e.fs.PathJoin(e.tablesDir, d)
		if t == nil {
			_ = e.fs.RemoveAll(path)
			continue
		}
		live := map[uint64]bool{}
		for _, f := range t.files {
			live[f.num] = true
		}
		names, _ := e.fs.List(path)
		for _, name := range names {
			if kind, num, ok := parseFileName(name); ok && kind == "sst" && !live[num] {
				_ = e.fs.Remove(e.fs.PathJoin(path, name))
			}
		}
	}
	return nil
}

func (e *Engine) loadState() (*manifestState, error) {
	state, exists, err := loadManifest(e.fs, e.dir)
	if err != nil {
		return nil, err
	}
	if !exists {
		state = newManifestState()
	}
	return state, nil
}

func sortFilesNewestFirst(files []*fileMeta) {
	slices.SortFunc(files, func(a, b *fileMeta) int {
		switch {
		case a.maxSeq > b.maxSeq:
			return -1
		case a.maxSeq < b.maxSeq:
			return 1
		}
		return 0
	})
}

func (e *Engine) checkUsableLocked() error {
	if e.mu.closed {
		return ErrClosed
	}
	return e.mu.bgErr
}

// CreateTable creates a table and returns its store.
func (e *Engine) CreateTable(id TableID) (*Table, error) {
	// commitMu fixes lastSeq: WAL entries at or below it belong to earlier
	// tables with the same ID, if any, and must never be replayed into this
	// one.
	e.commitMu.Lock()
	defer e.commitMu.Unlock()
	e.manifestMu.Lock()
	defer e.manifestMu.Unlock()
	e.mu.Lock()
	if err := e.checkUsableLocked(); err != nil {
		e.mu.Unlock()
		return nil, err
	}
	if _, ok := e.mu.tables[id]; ok {
		e.mu.Unlock()
		return nil, ErrTableExists
	}
	manifest := e.mu.manifest
	e.mu.Unlock()

	t := e.newTable(id)
	if err := e.fs.MkdirAll(t.dir, 0755); err != nil {
		return nil, err
	}
	if err := syncDir(e.fs, e.tablesDir); err != nil {
		return nil, err
	}
	t.flushedSeq = e.lastSeq
	edit := &versionEdit{tables: []*tableEdit{{id: id, created: true, flushedSeq: e.lastSeq}}}
	if err := manifest.write(edit); err != nil {
		e.setBackgroundError(fmt.Errorf("lsm: manifest write failed: %w", err))
		return nil, err
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	t.publishLocked()
	e.mu.tables[id] = t
	return t, nil
}

// Table returns the store of an existing table.
func (e *Engine) Table(id TableID) (*Table, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if t, ok := e.mu.tables[id]; ok {
		return t, nil
	}
	return nil, ErrUnknownTable
}

// DropTable drops a table: it is removed from the manifest, its pending
// memtables are discarded, and its directory is deleted. WAL entries for the
// table are ignored from then on.
//
// DropTable must not run concurrently with reads of the table; the schema
// layer serializes them.
func (e *Engine) DropTable(id TableID) error {
	e.commitMu.Lock()
	defer e.commitMu.Unlock()
	e.manifestMu.Lock()
	defer e.manifestMu.Unlock()
	e.mu.Lock()
	if err := e.checkUsableLocked(); err != nil {
		e.mu.Unlock()
		return err
	}
	t, ok := e.mu.tables[id]
	if !ok {
		e.mu.Unlock()
		return ErrUnknownTable
	}
	manifest := e.mu.manifest
	e.mu.Unlock()

	if err := manifest.write(&versionEdit{tables: []*tableEdit{{id: id, dropped: true}}}); err != nil {
		e.setBackgroundError(fmt.Errorf("lsm: manifest write failed: %w", err))
		return err
	}

	e.mu.Lock()
	t.dropped.Store(true)
	delete(e.mu.tables, id)
	// Discard the table's memtables. No flush is running: flushes hold
	// manifestMu, which DropTable holds.
	e.memUsed.Add(-t.mem.approximateSize())
	queue := e.mu.flushQueue[:0]
	for _, job := range e.mu.flushQueue {
		if job.t == t {
			e.mu.immBytes -= job.imm.mem.approximateSize()
			e.memUsed.Add(-job.imm.mem.approximateSize())
			continue
		}
		queue = append(queue, job)
	}
	e.mu.flushQueue = queue
	files := t.files
	t.files, t.imm = nil, nil
	e.mu.cond.Broadcast()
	e.mu.Unlock()

	var errs []error
	for _, f := range files {
		errs = append(errs, f.reader.close())
	}
	errs = append(errs, e.fs.RemoveAll(t.dir), syncDir(e.fs, e.tablesDir))
	return errors.Join(errs...)
}

func (e *Engine) setBackgroundError(err error) {
	e.mu.Lock()
	if e.mu.bgErr == nil {
		e.mu.bgErr = err
	}
	e.mu.cond.Broadcast()
	e.mu.Unlock()
}

// Apply writes the batch atomically. Every table it writes to must exist.
func (e *Engine) Apply(b *Batch, wo *WriteOptions) error {
	if b.count == 0 {
		return nil
	}
	sync := wo != nil && wo.Sync
	e.commitMu.Lock()
	defer e.commitMu.Unlock()

	e.mu.Lock()
	if err := e.makeRoomLocked(); err != nil {
		e.mu.Unlock()
		return err
	}
	type target struct {
		id  TableID
		mem *memTable
	}
	targets := make([]target, 0, len(b.tables))
	for _, id := range b.tables {
		t, ok := e.mu.tables[id]
		if !ok {
			e.mu.Unlock()
			return fmt.Errorf("%w: %d", ErrUnknownTable, id)
		}
		if t.mem.firstSeg == 0 {
			t.mem.firstSeg = e.mu.wal.num
		}
		targets = append(targets, target{id, t.mem})
	}
	wal := e.mu.wal
	e.mu.Unlock()

	seq := e.lastSeq + 1
	b.setSeqNum(seq)
	if err := wal.add(b.data, sync); err != nil {
		// The WAL is in an unknown state; refuse further writes.
		e.setBackgroundError(fmt.Errorf("lsm: WAL write failed: %w", err))
		return err
	}
	br, _ := newBatchReader(b.data)
	for s := seq; ; s++ {
		ent, ok, _ := br.next()
		if !ok {
			break
		}
		for _, tg := range targets {
			if tg.id == ent.table {
				tg.mem.add(ent.key, makeTrailer(s, ent.kind), ent.value)
				break
			}
		}
	}
	e.lastSeq = seq + uint64(b.count) - 1
	e.visibleSeq.Store(e.lastSeq)
	return nil
}

// makeRoomLocked enforces the WAL and memory limits before a write, rotating
// memtables and WAL segments as needed and stalling while memtables exceed
// the budget. commitMu and mu are held.
func (e *Engine) makeRoomLocked() error {
	for {
		if err := e.checkUsableLocked(); err != nil {
			return err
		}
		if e.mu.wal.size >= e.opts.WALSegmentSize {
			if err := e.rotateWALLocked(); err != nil {
				return err
			}
		}
		if e.walBytesLocked() > e.opts.WALTotalSpace {
			// Flush the tables whose mutable memtables pin the oldest segment.
			// Rotated memtables are already queued for flush.
			oldest := e.mu.walSegs[0].num
			for _, t := range e.mu.tables {
				if t.mem.firstSeg != 0 && t.mem.firstSeg <= oldest {
					e.rotateLocked(t)
				}
			}
		}
		mutable := e.memUsed.Load() - e.mu.immBytes
		if float64(mutable) > float64(e.opts.MemTableBudget)*e.opts.MemTableCleanupThreshold {
			if t := e.largestMutableLocked(); t != nil {
				e.rotateLocked(t)
				continue
			}
		}
		if e.memUsed.Load() <= e.opts.MemTableBudget || len(e.mu.flushQueue) == 0 {
			return nil
		}
		e.mu.cond.Wait()
	}
}

func (e *Engine) largestMutableLocked() *Table {
	var largest *Table
	for _, t := range e.mu.tables {
		if !t.mem.empty() && (largest == nil || t.mem.approximateSize() > largest.mem.approximateSize()) {
			largest = t
		}
	}
	return largest
}

func (e *Engine) walBytesLocked() int64 {
	var n int64
	for _, s := range e.mu.walSegs[:len(e.mu.walSegs)-1] {
		n += s.size
	}
	return n + e.mu.wal.size
}

// rotateWALLocked starts a new WAL segment. commitMu and mu are held.
func (e *Engine) rotateWALLocked() error {
	num := e.mu.nextFileNum
	e.mu.nextFileNum++
	wal, err := createWAL(e.fs, e.dir, num)
	if err == nil {
		err = syncDir(e.fs, e.dir)
	}
	if err == nil {
		err = e.mu.wal.close()
	}
	if err != nil {
		e.mu.bgErr = fmt.Errorf("lsm: rotating WAL: %w", err)
		return e.mu.bgErr
	}
	e.mu.walSegs[len(e.mu.walSegs)-1].size = e.mu.wal.size
	e.mu.walSegs = append(e.mu.walSegs, walSegment{num: num})
	e.mu.wal = wal
	return nil
}

// rotateLocked makes the table's mutable memtable immutable and queues it for
// flush. It does nothing if the memtable is empty. commitMu and mu are held,
// so no write is in progress and lastSeq bounds the memtable's entries.
func (e *Engine) rotateLocked(t *Table) {
	if t.mem.empty() {
		return
	}
	im := &immMemTable{mem: t.mem, boundary: e.lastSeq}
	t.imm = append(t.imm, im)
	t.mem = newMemTable(&e.memUsed)
	e.mu.immBytes += im.mem.approximateSize()
	e.mu.flushQueue = append(e.mu.flushQueue, flushJob{t: t, imm: im})
	t.publishLocked()
	e.maybeScheduleFlushLocked()
}

func (e *Engine) maybeScheduleFlushLocked() {
	if e.mu.flushing || len(e.mu.flushQueue) == 0 || e.mu.bgErr != nil || e.mu.closed {
		return
	}
	e.mu.flushing = true
	e.flushWG.Add(1)
	go e.flushLoop()
}

// minNeededWALLocked returns the oldest WAL segment some memtable still
// needs, ignoring the memtable being flushed by except.
func (e *Engine) minNeededWALLocked(except *immMemTable) uint64 {
	minSeg := e.mu.wal.num
	for _, t := range e.mu.tables {
		if t.mem.firstSeg != 0 {
			minSeg = min(minSeg, t.mem.firstSeg)
		}
		for _, im := range t.imm {
			if im != except {
				minSeg = min(minSeg, im.mem.firstSeg)
			}
		}
	}
	return minSeg
}

// flushLoop flushes queued memtables in rotation order until none remain.
func (e *Engine) flushLoop() {
	defer e.flushWG.Done()
	for {
		e.mu.Lock()
		if len(e.mu.flushQueue) == 0 || e.mu.bgErr != nil || e.mu.closed {
			e.mu.flushing = false
			e.mu.cond.Broadcast()
			e.mu.Unlock()
			return
		}
		e.mu.Unlock()
		if err := e.flushOne(); err != nil {
			e.setBackgroundError(fmt.Errorf("lsm: flush failed: %w", err))
		}
	}
}

// flushOne flushes the memtable at the head of the queue. The sstable is
// written and the manifest edit is synced before readers switch from the
// memtable to the sstable, so a read never misses data.
func (e *Engine) flushOne() error {
	e.manifestMu.Lock()
	defer e.manifestMu.Unlock()

	e.mu.Lock()
	if len(e.mu.flushQueue) == 0 {
		// DropTable removed the queued memtables.
		e.mu.Unlock()
		return nil
	}
	job := e.mu.flushQueue[0]
	t, im := job.t, job.imm
	num := e.mu.nextFileNum
	e.mu.nextFileNum++
	e.mu.Unlock()

	meta, err := writeTable(e.fs, t.dir, num, e.writerOpts, im.mem)
	if err == nil && meta != nil {
		meta.reader, err = openTable(e.fs, t.dir, num, e.readerOpts)
	}
	if err == nil {
		err = syncDir(e.fs, t.dir)
	}
	if err != nil {
		if t.dropped.Load() {
			// The table's directory may already be gone.
			err = nil
			e.mu.Lock()
			e.finishFlushLocked(job)
			e.mu.Unlock()
		}
		return err
	}

	e.mu.Lock()
	dropped := t.dropped.Load()
	edit := &versionEdit{nextFileNum: e.mu.nextFileNum, minWAL: e.minNeededWALLocked(im)}
	if !dropped {
		te := &tableEdit{id: t.id, flushedSeq: im.boundary}
		if meta != nil {
			te.added = []*fileMeta{meta}
		}
		edit.tables = []*tableEdit{te}
	}
	manifest := e.mu.manifest
	e.mu.Unlock()

	if err := manifest.write(edit); err != nil {
		if meta != nil {
			meta.reader.close()
		}
		return err
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	if t.dropped.Load() {
		if meta != nil {
			meta.reader.close()
		}
	} else {
		if meta != nil {
			t.files = append([]*fileMeta{meta}, t.files...)
		}
		t.flushedSeq = im.boundary
		t.imm = t.imm[1:]
		t.publishLocked()
	}
	e.finishFlushLocked(job)
	e.deleteObsoleteWALLocked(edit.minWAL)
	return nil
}

// finishFlushLocked removes a finished job from the queue and releases its
// memory.
func (e *Engine) finishFlushLocked(job flushJob) {
	e.mu.flushQueue = e.mu.flushQueue[1:]
	sz := job.imm.mem.approximateSize()
	e.mu.immBytes -= sz
	e.memUsed.Add(-sz)
	e.mu.cond.Broadcast()
}

// deleteObsoleteWALLocked deletes WAL segments below minWAL. The manifest
// already records that every table has flushed past them.
func (e *Engine) deleteObsoleteWALLocked(minWAL uint64) {
	for len(e.mu.walSegs) > 1 && e.mu.walSegs[0].num < minWAL {
		_ = e.fs.Remove(e.fs.PathJoin(e.dir, walFileName(e.mu.walSegs[0].num)))
		e.mu.walSegs = e.mu.walSegs[1:]
	}
}

// Flush flushes every table and waits until all memtables are written.
func (e *Engine) Flush() error {
	e.commitMu.Lock()
	e.mu.Lock()
	defer e.mu.Unlock()
	if err := e.checkUsableLocked(); err != nil {
		e.commitMu.Unlock()
		return err
	}
	for _, t := range e.mu.tables {
		e.rotateLocked(t)
	}
	e.commitMu.Unlock()
	for len(e.mu.flushQueue) > 0 {
		if err := e.checkUsableLocked(); err != nil {
			return err
		}
		e.mu.cond.Wait()
	}
	return e.checkUsableLocked()
}

// Close waits for an in-progress flush, then closes the engine. Data in
// memtables is recovered from the WAL on the next Open.
func (e *Engine) Close() error {
	e.commitMu.Lock()
	defer e.commitMu.Unlock()
	e.manifestMu.Lock()
	defer e.manifestMu.Unlock()
	e.mu.Lock()
	if e.mu.closed {
		e.mu.Unlock()
		return ErrClosed
	}
	e.mu.closed = true
	e.mu.cond.Broadcast()
	e.mu.Unlock()
	// A flush in progress needs manifestMu to finish; release it meanwhile.
	e.manifestMu.Unlock()
	e.flushWG.Wait()
	e.manifestMu.Lock()

	e.mu.Lock()
	defer e.mu.Unlock()
	err := e.closeFiles()
	if lerr := e.lock.Close(); err == nil {
		err = lerr
	}
	return err
}

func (e *Engine) closeFiles() error {
	var errs []error
	if e.mu.wal != nil {
		errs = append(errs, e.mu.wal.close())
		e.mu.wal = nil
	}
	if e.mu.manifest != nil {
		errs = append(errs, e.mu.manifest.close())
		e.mu.manifest = nil
	}
	for _, t := range e.mu.tables {
		for _, f := range t.files {
			if f.reader != nil {
				errs = append(errs, f.reader.close())
				f.reader = nil
			}
		}
	}
	return errors.Join(errs...)
}
