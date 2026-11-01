// Package lsm is cassandra-room's LSM storage engine. See
// docs/storage-engine.md for the design.
package lsm

import (
	"errors"
	"fmt"

	"github.com/cockroachdb/pebble/sstable"
	"github.com/cockroachdb/pebble/vfs"
)

var (
	// ErrNotFound is returned by Get when the key does not exist.
	ErrNotFound = errors.New("lsm: not found")
	// ErrClosed is returned by operations on a closed engine.
	ErrClosed = errors.New("lsm: closed")
	// ErrUnknownTable is returned for a table that does not exist or was
	// dropped.
	ErrUnknownTable = errors.New("lsm: unknown table")
	// ErrTableExists is returned by CreateTable for an existing table.
	ErrTableExists = errors.New("lsm: table exists")
)

// TableID identifies a table. The schema layer maps Cassandra's table UUIDs
// to TableIDs.
type TableID uint64

func (id TableID) dirName() string { return fmt.Sprintf("%016x", uint64(id)) }

// Options configures an Engine. The zero value is usable; see ensureDefaults.
type Options struct {
	// FS is the file system the engine lives on. Defaults to vfs.Default.
	FS vfs.FS
	// MemTableBudget is the memory all memtables may use together, including
	// those waiting to be flushed. Writes stall while it is exceeded.
	// Defaults to 256 MiB. (Cassandra: memtable_heap_space.)
	MemTableBudget int64
	// MemTableCleanupThreshold is the fraction of MemTableBudget that mutable
	// memtables may use before the largest is flushed. Defaults to 0.33.
	// (Cassandra: memtable_cleanup_threshold, 1/(memtable_flush_writers+1).)
	MemTableCleanupThreshold float64
	// WALSegmentSize is the size at which the commit log moves to a new
	// segment. Defaults to 32 MiB. (Cassandra: commitlog_segment_size.)
	WALSegmentSize int64
	// WALTotalSpace is the size of live commit log segments above which the
	// tables pinning the oldest segment are flushed. Defaults to 1 GiB.
	// (Cassandra: commitlog_total_space.)
	WALTotalSpace int64
	// BloomBitsPerKey configures the Bloom filter written into every sstable.
	// Defaults to 10 (about 1% false positives, Cassandra's default).
	BloomBitsPerKey uint32
	// BlockSize is the target uncompressed size of sstable data blocks. Zero
	// uses Pebble's default.
	BlockSize int
}

func (o *Options) ensureDefaults() {
	if o.FS == nil {
		o.FS = vfs.Default
	}
	if o.MemTableBudget <= 0 {
		o.MemTableBudget = 256 << 20
	}
	if o.MemTableCleanupThreshold <= 0 || o.MemTableCleanupThreshold > 1 {
		o.MemTableCleanupThreshold = 0.33
	}
	if o.WALSegmentSize <= 0 {
		o.WALSegmentSize = 32 << 20
	}
	if o.WALTotalSpace <= 0 {
		o.WALTotalSpace = 1 << 30
	}
	if o.BloomBitsPerKey == 0 {
		o.BloomBitsPerKey = 10
	}
}

// tableFormat is the sstable format the engine writes: Pebble's columnar
// format with checksummed footers.
const tableFormat = sstable.TableFormatPebblev6

// WriteOptions configures a single write.
type WriteOptions struct {
	// Sync makes the write durable (the WAL is fsynced) before Apply returns.
	Sync bool
}

var (
	// Sync is a WriteOptions that fsyncs the WAL.
	Sync = &WriteOptions{Sync: true}
	// NoSync is a WriteOptions that hands the write to the OS without fsync.
	NoSync = &WriteOptions{Sync: false}
)
