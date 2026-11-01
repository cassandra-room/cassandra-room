package lsm

// Benchmarks comparing the engine with Pebble (pinned in go.mod) on the same
// workload: the same keys and the same encoded row values (one 100-byte cell).
// Run with: go test -run x -bench . ./internal/storage/lsm/

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/cassandra-room/cassandra-room/internal/storage/row"
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/sstable/tablefilters/bloom"
	"github.com/cockroachdb/pebble/vfs"
)

func benchKey(i int) []byte { return []byte(fmt.Sprintf("key%012d", i)) }

func benchValue() []byte {
	return row.Append(nil, &row.Row{
		Deletion: row.Deletion{MarkedForDeleteAt: row.NoTimestamp},
		Cells:    []row.Cell{{Column: 1, Timestamp: 1, LocalDeletionTime: row.NoDeletionTime, Value: make([]byte, 100)}},
	})
}

type benchStore interface {
	put(key, value []byte, sync bool) error
	get(key []byte) error
	flush() error
	// compact rewrites everything into the fewest files; the engine has no
	// compaction yet and already holds one file per flush.
	compact() error
	close() error
}

type lsmStore struct {
	e   *Engine
	tbl *Table
}

func (s lsmStore) put(key, value []byte, sync bool) error {
	wo := NoSync
	if sync {
		wo = Sync
	}
	return s.e.Put(1, key, value, wo)
}
func (s lsmStore) get(key []byte) error { _, err := s.tbl.Get(key); return err }
func (s lsmStore) flush() error         { return s.e.Flush() }
func (s lsmStore) compact() error       { return nil }
func (s lsmStore) close() error         { return s.e.Close() }

type pebbleStore struct{ db *pebble.DB }

func (s pebbleStore) put(key, value []byte, sync bool) error {
	wo := pebble.NoSync
	if sync {
		wo = pebble.Sync
	}
	return s.db.Set(key, value, wo)
}
func (s pebbleStore) get(key []byte) error {
	_, c, err := s.db.Get(key)
	if err == nil {
		c.Close()
	}
	return err
}
func (s pebbleStore) flush() error { return s.db.Flush() }
func (s pebbleStore) compact() error {
	return s.db.Compact(context.Background(), []byte("a"), []byte("z"), true)
}
func (s pebbleStore) close() error { return s.db.Close() }

func openBench(b *testing.B, engine string) benchStore {
	switch engine {
	case "lsm":
		e, err := Open("db", &Options{FS: vfs.NewMem()})
		if err != nil {
			b.Fatal(err)
		}
		tbl, err := e.CreateTable(1)
		if err != nil {
			b.Fatal(err)
		}
		return lsmStore{e, tbl}
	default:
		// Compactions are off: the engine has none yet. Pebble gets the same
		// block cache size and Bloom filters as the engine's defaults.
		c := pebble.NewCache(64 << 20)
		defer c.Unref()
		opts := &pebble.Options{FS: vfs.NewMem(), Cache: c, DisableAutomaticCompactions: true}
		for i := range opts.Levels {
			opts.Levels[i].TableFilterPolicy = func() pebble.TableFilterPolicy { return bloom.FilterPolicy(10) }
		}
		db, err := pebble.Open("db", opts)
		if err != nil {
			b.Fatal(err)
		}
		return pebbleStore{db}
	}
}

var engines = []string{"lsm", "pebble"}

func BenchmarkPut(b *testing.B) {
	val := benchValue()
	for _, eng := range engines {
		b.Run(eng, func(b *testing.B) {
			s := openBench(b, eng)
			defer s.close()
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := s.put(benchKey(i), val, false); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkPutParallel(b *testing.B) {
	val := benchValue()
	for _, eng := range engines {
		for _, sync := range []bool{false, true} {
			b.Run(fmt.Sprintf("%s/sync=%t", eng, sync), func(b *testing.B) {
				s := openBench(b, eng)
				defer s.close()
				var n atomic.Int64
				b.ReportAllocs()
				b.ResetTimer()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						if err := s.put(benchKey(int(n.Add(1))), val, sync); err != nil {
							b.Error(err)
							return
						}
					}
				})
			})
		}
	}
}

func BenchmarkGetFromSSTables(b *testing.B) {
	const n = 100000
	val := benchValue()
	for _, eng := range engines {
		b.Run(eng, func(b *testing.B) {
			s := openBench(b, eng)
			defer s.close()
			for i := 0; i < n; i++ {
				if err := s.put(benchKey(i), val, false); err != nil {
					b.Fatal(err)
				}
			}
			if err := s.flush(); err != nil {
				b.Fatal(err)
			}
			// Both engines now read from a single file per key.
			if err := s.compact(); err != nil {
				b.Fatal(err)
			}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := s.get(benchKey((i * 7919) % n)); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
