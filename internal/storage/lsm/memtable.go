package lsm

import (
	"bytes"
	"sync/atomic"

	"github.com/cockroachdb/pebble"
)

const (
	maxHeight = 20
	// pBranch is the inverse probability of a node reaching the next level.
	pBranch = 4
	// arenaChunkSize is the size of the chunks keys and values are copied
	// into, which keeps small allocations off the garbage collector's path.
	arenaChunkSize = 64 << 10
	// nodeOverhead approximates the memory a node uses besides its key,
	// value and tower, for size accounting.
	nodeOverhead = 64
)

// memNode is a skiplist node. key, value and trailer never change after the
// node is published; tower entries are only ever set by the single writer.
type memNode struct {
	key     []byte
	value   []byte
	trailer uint64 // seqnum << 8 | kind, as in Pebble's internal keys
	tower   []atomic.Pointer[memNode]
}

func (n *memNode) next() *memNode { return n.tower[0].Load() }

func (n *memNode) kind() pebble.InternalKeyKind { return pebble.InternalKeyKind(n.trailer & 0xff) }

// memTable is a skiplist ordered by user key ascending, then trailer
// descending (newest version first), matching Pebble's internal key order.
// There is a single writer (Engine.Apply holds commitMu) and any number of
// lock-free readers.
type memTable struct {
	head   *memNode
	height atomic.Int32
	size   atomic.Int64
	// used is the engine-wide memtable memory counter this memtable adds to.
	used *atomic.Int64
	// firstSeg is the WAL segment of the memtable's first entry, or 0 while
	// it is empty. Segments from firstSeg on are needed until the memtable is
	// flushed. Guarded by Engine.mu.
	firstSeg uint64

	// Writer-only state.
	rnd   uint64
	chunk []byte
}

var memTableSeed atomic.Uint64

func newMemTable(used *atomic.Int64) *memTable {
	m := &memTable{
		head: &memNode{tower: make([]atomic.Pointer[memNode], maxHeight)},
		used: used,
		rnd:  memTableSeed.Add(1)*0x9E3779B97F4A7C15 | 1,
	}
	m.height.Store(1)
	return m
}

func (m *memTable) empty() bool { return m.head.next() == nil }

func (m *memTable) approximateSize() int64 { return m.size.Load() }

// compare orders node n against (key, trailer).
func compareNode(n *memNode, key []byte, trailer uint64) int {
	if c := bytes.Compare(n.key, key); c != 0 {
		return c
	}
	switch {
	case n.trailer > trailer:
		return -1
	case n.trailer < trailer:
		return 1
	}
	return 0
}

func (m *memTable) randomHeight() int {
	h := 1
	for h < maxHeight {
		// xorshift64*
		m.rnd ^= m.rnd >> 12
		m.rnd ^= m.rnd << 25
		m.rnd ^= m.rnd >> 27
		if (m.rnd*2685821657736338717)%pBranch != 0 {
			break
		}
		h++
	}
	return h
}

// alloc copies b into the arena.
func (m *memTable) alloc(b []byte) []byte {
	if len(b) == 0 {
		return nil
	}
	if len(b) > arenaChunkSize/4 {
		return append([]byte(nil), b...)
	}
	if cap(m.chunk)-len(m.chunk) < len(b) {
		m.chunk = make([]byte, 0, arenaChunkSize)
	}
	start := len(m.chunk)
	m.chunk = append(m.chunk, b...)
	return m.chunk[start:len(m.chunk):len(m.chunk)]
}

// add inserts an entry. Only the single writer may call it. Entries with the
// same key and trailer are not expected (sequence numbers are unique).
func (m *memTable) add(key []byte, trailer uint64, value []byte) {
	var prev [maxHeight]*memNode
	listHeight := int(m.height.Load())
	x := m.head
	for level := listHeight - 1; level >= 0; level-- {
		for {
			nx := x.tower[level].Load()
			if nx == nil || compareNode(nx, key, trailer) >= 0 {
				break
			}
			x = nx
		}
		prev[level] = x
	}

	h := m.randomHeight()
	if h > listHeight {
		for level := listHeight; level < h; level++ {
			prev[level] = m.head
		}
		m.height.Store(int32(h))
	}
	n := &memNode{
		key:     m.alloc(key),
		value:   m.alloc(value),
		trailer: trailer,
		tower:   make([]atomic.Pointer[memNode], h),
	}
	// Link bottom-up. Each new pointer is set before the node is published at
	// that level, so readers always see a well-formed list.
	for level := 0; level < h; level++ {
		n.tower[level].Store(prev[level].tower[level].Load())
		prev[level].tower[level].Store(n)
	}
	sz := int64(len(key) + len(value) + nodeOverhead + 8*h)
	m.size.Add(sz)
	m.used.Add(sz)
}

// seekGE returns the first node at or after (key, trailer), or nil.
func (m *memTable) seekGE(key []byte, trailer uint64) *memNode {
	x := m.head
	var nx *memNode
	for level := int(m.height.Load()) - 1; level >= 0; level-- {
		for {
			nx = x.tower[level].Load()
			if nx == nil || compareNode(nx, key, trailer) >= 0 {
				break
			}
			x = nx
		}
	}
	// Return the node that was compared, not a fresh load of x.tower[0]: a
	// concurrent insert may have linked a smaller node after x since then.
	return nx
}

// get returns the newest entry for key with a sequence number at or below
// seq, or nil.
func (m *memTable) get(key []byte, seq uint64) *memNode {
	n := m.seekGE(key, makeTrailer(seq, pebble.InternalKeyKindMax))
	if n == nil || !bytes.Equal(n.key, key) {
		return nil
	}
	return n
}

// first returns the first node, or nil if the memtable is empty.
func (m *memTable) first() *memNode { return m.head.next() }

func makeTrailer(seq uint64, kind pebble.InternalKeyKind) uint64 {
	return seq<<8 | uint64(kind)
}
