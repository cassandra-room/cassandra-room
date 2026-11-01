package lsm

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"testing"

	"github.com/cockroachdb/pebble/vfs"
	"github.com/stretchr/testify/require"
)

func mustGet(t *testing.T, tbl *Table, key, want string) {
	t.Helper()
	v, err := tbl.Get([]byte(key))
	require.NoError(t, err, "key %q", key)
	require.Equal(t, want, string(v), "key %q", key)
}

func mustNotFound(t *testing.T, tbl *Table, key string) {
	t.Helper()
	_, err := tbl.Get([]byte(key))
	require.ErrorIs(t, err, ErrNotFound, "key %q", key)
}

func set(t *testing.T, e *Engine, id TableID, key, value string, wo *WriteOptions) {
	t.Helper()
	b := NewBatch()
	b.Set(id, []byte(key), []byte(value))
	require.NoError(t, e.Apply(b, wo))
}

func del(t *testing.T, e *Engine, id TableID, key string) {
	t.Helper()
	b := NewBatch()
	b.Delete(id, []byte(key))
	require.NoError(t, e.Apply(b, NoSync))
}

func tableFiles(e *Engine, id TableID) int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return len(e.mu.tables[id].files)
}

func walFiles(t *testing.T, fs vfs.FS) int {
	names, err := fs.List("db")
	require.NoError(t, err)
	n := 0
	for _, name := range names {
		if kind, _, ok := parseFileName(name); ok && kind == "wal" {
			n++
		}
	}
	return n
}

func TestBasic(t *testing.T) {
	e, err := Open("db", &Options{FS: vfs.NewMem()})
	require.NoError(t, err)
	defer e.Close()
	a, err := e.CreateTable(1)
	require.NoError(t, err)
	b, err := e.CreateTable(2)
	require.NoError(t, err)
	_, err = e.CreateTable(1)
	require.ErrorIs(t, err, ErrTableExists)

	mustNotFound(t, a, "k")
	set(t, e, 1, "k", "a1", NoSync)
	set(t, e, 2, "k", "b1", NoSync)
	mustGet(t, a, "k", "a1")
	mustGet(t, b, "k", "b1") // tables are separate key spaces
	del(t, e, 1, "k")
	mustNotFound(t, a, "k")
	mustGet(t, b, "k", "b1")

	// A batch across tables is atomic and rejected as a whole if any table
	// is unknown.
	batch := NewBatch()
	batch.Set(1, []byte("x"), []byte("1"))
	batch.Set(2, []byte("y"), []byte("2"))
	require.NoError(t, e.Apply(batch, Sync))
	mustGet(t, a, "x", "1")
	mustGet(t, b, "y", "2")
	bad := NewBatch()
	bad.Set(1, []byte("z"), []byte("1"))
	bad.Set(99, []byte("z"), []byte("1"))
	require.ErrorIs(t, e.Apply(bad, NoSync), ErrUnknownTable)
	mustNotFound(t, a, "z")
}

// TestFlushReopen checks reads across memtables and several sstables, with
// overwrites and deletes spanning sstables, in two tables, before and after
// reopening.
func TestFlushReopen(t *testing.T) {
	fs := vfs.NewMem()
	opts := &Options{FS: fs, MemTableBudget: 96 << 10}
	e, err := Open("db", opts)
	require.NoError(t, err)
	for _, id := range []TableID{1, 2} {
		_, err := e.CreateTable(id)
		require.NoError(t, err)
	}

	model := map[TableID]map[string]string{1: {}, 2: {}}
	rng := rand.New(rand.NewPCG(3, 4))
	for i := 0; i < 20000; i++ {
		id := TableID(1 + rng.IntN(2))
		key := fmt.Sprintf("key%05d", rng.IntN(3000))
		if rng.IntN(5) == 0 {
			del(t, e, id, key)
			delete(model[id], key)
		} else {
			val := fmt.Sprintf("v%d", i)
			set(t, e, id, key, val, NoSync)
			model[id][key] = val
		}
	}
	check := func(e *Engine) {
		for id, m := range model {
			tbl, err := e.Table(id)
			require.NoError(t, err)
			for i := 0; i < 3000; i++ {
				key := fmt.Sprintf("key%05d", i)
				if v, ok := m[key]; ok {
					mustGet(t, tbl, key, v)
				} else {
					mustNotFound(t, tbl, key)
				}
			}
		}
	}
	check(e)
	require.Greater(t, tableFiles(e, 1), 1, "expected flushes driven by the memory budget")
	require.NoError(t, e.Flush())
	check(e)
	require.NoError(t, e.Close())

	e, err = Open("db", opts)
	require.NoError(t, err)
	check(e)
	require.NoError(t, e.Close())
}

// TestRecoverFromWAL closes without flushing; reopening must replay the WAL
// into the right tables.
func TestRecoverFromWAL(t *testing.T) {
	fs := vfs.NewMem()
	e, err := Open("db", &Options{FS: fs})
	require.NoError(t, err)
	_, err = e.CreateTable(1)
	require.NoError(t, err)
	_, err = e.CreateTable(2)
	require.NoError(t, err)
	for i := 0; i < 100; i++ {
		set(t, e, TableID(1+i%2), fmt.Sprintf("k%d", i), fmt.Sprintf("v%d", i), NoSync)
	}
	del(t, e, 2, "k5")
	require.NoError(t, e.Close())

	e, err = Open("db", &Options{FS: fs})
	require.NoError(t, err)
	for i := 0; i < 100; i++ {
		tbl, _ := e.Table(TableID(1 + i%2))
		if i == 5 {
			mustNotFound(t, tbl, "k5")
			continue
		}
		mustGet(t, tbl, fmt.Sprintf("k%d", i), fmt.Sprintf("v%d", i))
	}
	// New writes get higher sequence numbers than recovered ones.
	set(t, e, 1, "k0", "new", NoSync)
	a, _ := e.Table(1)
	mustGet(t, a, "k0", "new")
	require.NoError(t, e.Close())

	e, err = Open("db", &Options{FS: fs})
	require.NoError(t, err)
	a, _ = e.Table(1)
	mustGet(t, a, "k0", "new")
	require.NoError(t, e.Close())
}

// TestTableFlushIsPerTable checks that flushing one table leaves the others'
// memtables alone, as Cassandra's per-table flush does.
func TestTableFlushIsPerTable(t *testing.T) {
	e, err := Open("db", &Options{FS: vfs.NewMem()})
	require.NoError(t, err)
	defer e.Close()
	a, _ := e.CreateTable(1)
	_, _ = e.CreateTable(2)
	set(t, e, 1, "a", "1", NoSync)
	set(t, e, 2, "b", "2", NoSync)
	require.NoError(t, a.Flush())
	require.Equal(t, 1, tableFiles(e, 1))
	require.Equal(t, 0, tableFiles(e, 2))
	mustGet(t, a, "a", "1")
}

// TestMemoryBudgetFlushesLargest checks that exceeding the cleanup threshold
// flushes the largest memtable, not every table.
func TestMemoryBudgetFlushesLargest(t *testing.T) {
	e, err := Open("db", &Options{FS: vfs.NewMem(), MemTableBudget: 300 << 10, MemTableCleanupThreshold: 0.5})
	require.NoError(t, err)
	defer e.Close()
	_, _ = e.CreateTable(1)
	quiet, _ := e.CreateTable(2)
	set(t, e, 2, "quiet", "1", NoSync)
	val := string(make([]byte, 100))
	for i := 0; i < 5000; i++ {
		set(t, e, 1, fmt.Sprintf("k%05d", i), val, NoSync)
	}
	busy, _ := e.Table(1)
	require.NoError(t, busy.Flush())
	require.Greater(t, tableFiles(e, 1), 1)
	require.Equal(t, 0, tableFiles(e, 2), "the quiet table must not be flushed")
	mustGet(t, quiet, "quiet", "1")
}

// TestQuietTableDoesNotPinWAL checks Cassandra's commitlog_total_space rule:
// a table whose memtable holds entries from the oldest WAL segment is flushed
// so the segments can be deleted.
func TestQuietTableDoesNotPinWAL(t *testing.T) {
	fs := vfs.NewMem()
	e, err := Open("db", &Options{
		FS:             fs,
		MemTableBudget: 1 << 20,
		WALSegmentSize: 16 << 10,
		WALTotalSpace:  64 << 10,
	})
	require.NoError(t, err)
	defer e.Close()
	_, _ = e.CreateTable(1)
	quiet, _ := e.CreateTable(2)
	set(t, e, 2, "quiet", "1", NoSync)
	val := string(make([]byte, 100))
	for i := 0; i < 5000; i++ {
		set(t, e, 1, fmt.Sprintf("k%05d", i), val, NoSync)
	}
	// Flush only the busy table; the quiet table may only have been flushed by
	// the WAL space rule.
	busy, _ := e.Table(1)
	require.NoError(t, busy.Flush())
	require.Equal(t, 1, tableFiles(e, 2), "the quiet table must have been flushed to free the WAL")
	require.LessOrEqual(t, walFiles(t, fs), 2)
	mustGet(t, quiet, "quiet", "1")
}

// TestDropTable checks that a dropped table's data and directory are gone, and
// that re-creating a table with the same ID does not replay the old table's
// WAL entries.
func TestDropTable(t *testing.T) {
	fs := vfs.NewMem()
	e, err := Open("db", &Options{FS: fs})
	require.NoError(t, err)
	a, _ := e.CreateTable(1)
	_, _ = e.CreateTable(2)
	set(t, e, 1, "flushed", "1", NoSync)
	require.NoError(t, a.Flush())
	set(t, e, 1, "unflushed", "1", NoSync) // only in the WAL
	set(t, e, 2, "other", "1", NoSync)

	require.NoError(t, e.DropTable(1))
	_, err = a.Get([]byte("flushed"))
	require.ErrorIs(t, err, ErrUnknownTable)
	_, err = e.Table(1)
	require.ErrorIs(t, err, ErrUnknownTable)
	_, err = fs.Stat(fs.PathJoin("db", tablesDirName, TableID(1).dirName()))
	require.Error(t, err, "the table's directory must be deleted")

	a2, err := e.CreateTable(1)
	require.NoError(t, err)
	mustNotFound(t, a2, "unflushed")
	require.NoError(t, e.Close())

	e, err = Open("db", &Options{FS: fs})
	require.NoError(t, err)
	defer e.Close()
	a2, _ = e.Table(1)
	mustNotFound(t, a2, "flushed")
	mustNotFound(t, a2, "unflushed")
	b, _ := e.Table(2)
	mustGet(t, b, "other", "1")
}

// TestCrashRecovery crashes the file system at random points, dropping
// unsynced data, and checks that every synced write survives in the right
// table and nothing else appears.
func TestCrashRecovery(t *testing.T) {
	for seed := uint64(0); seed < 20; seed++ {
		t.Run(fmt.Sprint(seed), func(t *testing.T) {
			rng := rand.New(rand.NewPCG(seed, 99))
			fs := vfs.NewCrashableMem()
			opts := &Options{FS: fs, MemTableBudget: 64 << 10, WALSegmentSize: 8 << 10, WALTotalSpace: 32 << 10}
			e, err := Open("db", opts)
			require.NoError(t, err)
			for _, id := range []TableID{1, 2, 3} {
				_, err := e.CreateTable(id)
				require.NoError(t, err)
			}

			type tk struct {
				id  TableID
				key string
			}
			synced := map[tk]string{}  // must survive
			written := map[tk]string{} // may survive
			n := 500 + rng.IntN(3000)
			for i := 0; i < n; i++ {
				k := tk{TableID(1 + rng.IntN(3)), fmt.Sprintf("k%05d", i)} // each key written once
				val := fmt.Sprintf("v%d", i)
				sync := rng.IntN(10) == 0
				wo := NoSync
				if sync {
					wo = Sync
				}
				set(t, e, k.id, k.key, val, wo)
				written[k] = val
				if sync {
					// A synced write makes every earlier write durable too.
					for k, v := range written {
						synced[k] = v
					}
				}
			}
			crashed := fs.CrashClone(vfs.CrashCloneCfg{UnsyncedDataPercent: rng.IntN(100), RNG: rng})
			require.NoError(t, e.Close())

			e2, err := Open("db", &Options{FS: crashed, MemTableBudget: 64 << 10})
			require.NoError(t, err)
			for k, v := range written {
				tbl, err := e2.Table(k.id)
				require.NoError(t, err)
				got, err := tbl.Get([]byte(k.key))
				if _, mustExist := synced[k]; mustExist {
					require.NoError(t, err, "synced key %v", k)
				}
				if err == nil {
					require.Equal(t, v, string(got), "key %v", k)
				} else {
					require.ErrorIs(t, err, ErrNotFound)
				}
			}
			require.NoError(t, e2.Close())
		})
	}
}

// TestConcurrentReadWrite runs readers while a writer writes to two tables
// and memtables are flushed, under the race detector. Each key's value only
// grows, so a reader must never see a value go backwards.
func TestConcurrentReadWrite(t *testing.T) {
	e, err := Open("db", &Options{FS: vfs.NewMem(), MemTableBudget: 48 << 10})
	require.NoError(t, err)
	defer e.Close()
	tables := []*Table{}
	for _, id := range []TableID{1, 2} {
		tbl, err := e.CreateTable(id)
		require.NoError(t, err)
		tables = append(tables, tbl)
	}

	const keys = 200
	var wg sync.WaitGroup
	done := make(chan struct{})
	for r := 0; r < 4; r++ {
		wg.Add(1)
		go func(r int) {
			defer wg.Done()
			tbl := tables[r%2]
			seen := make([]int, keys)
			rng := rand.New(rand.NewPCG(uint64(r), 1))
			for {
				select {
				case <-done:
					return
				default:
				}
				k := rng.IntN(keys)
				v, err := tbl.Get([]byte(fmt.Sprintf("k%03d", k)))
				if err == ErrNotFound {
					if seen[k] != 0 {
						t.Errorf("key %d disappeared after value %d", k, seen[k])
						return
					}
					continue
				}
				if err != nil {
					t.Error(err)
					return
				}
				var n int
				fmt.Sscanf(string(v), "%d", &n)
				if n < seen[k] {
					t.Errorf("key %d went backwards: %d after %d", k, n, seen[k])
					return
				}
				seen[k] = n
			}
		}(r)
	}
	for i := 1; i <= 20000; i++ {
		b := NewBatch()
		b.Set(TableID(1+i%2), []byte(fmt.Sprintf("k%03d", i%keys)), []byte(fmt.Sprintf("%d", i)))
		require.NoError(t, e.Apply(b, NoSync))
	}
	require.NoError(t, e.Flush())
	close(done)
	wg.Wait()
}

func TestClosedAndLocked(t *testing.T) {
	fs := vfs.NewMem()
	e, err := Open("db", &Options{FS: fs})
	require.NoError(t, err)
	_, _ = e.CreateTable(1)
	_, err = Open("db", &Options{FS: fs})
	require.Error(t, err, "a second Open must fail while the engine is locked")
	require.NoError(t, e.Close())
	b := NewBatch()
	b.Set(1, []byte("a"), nil)
	require.ErrorIs(t, e.Apply(b, NoSync), ErrClosed)
	require.ErrorIs(t, e.Close(), ErrClosed)
}

func BenchmarkSet(b *testing.B) {
	e, err := Open("db", &Options{FS: vfs.NewMem()})
	require.NoError(b, err)
	defer e.Close()
	_, _ = e.CreateTable(1)
	val := make([]byte, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		batch := NewBatch()
		batch.Set(1, []byte(fmt.Sprintf("key%012d", i)), val)
		if err := e.Apply(batch, NoSync); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetFromTables(b *testing.B) {
	e, err := Open("db", &Options{FS: vfs.NewMem(), MemTableBudget: 3 << 20})
	require.NoError(b, err)
	defer e.Close()
	tbl, _ := e.CreateTable(1)
	const n = 100000
	val := make([]byte, 100)
	for i := 0; i < n; i++ {
		batch := NewBatch()
		batch.Set(1, []byte(fmt.Sprintf("key%012d", i)), val)
		require.NoError(b, e.Apply(batch, NoSync))
	}
	require.NoError(b, e.Flush())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := tbl.Get([]byte(fmt.Sprintf("key%012d", (i*7919)%n))); err != nil {
			b.Fatal(err)
		}
	}
}
