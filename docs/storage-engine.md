# Storage engine design (`internal/storage/lsm`)

This document describes cassandra-room's own LSM storage engine. It replaces
the plan in MEMO.md to use Pebble as the storage engine.

The engine follows **Cassandra's storage architecture**: one commit log per
node, and a memtable, sstable set and compaction strategy per table. It
follows **Cassandra's semantics**: cells are reconciled by write timestamp,
not by arrival order. It stores data in **Pebble's sstable format**, reusing
Pebble's reader and writer. Compaction is Cassandra 5's Unified Compaction
Strategy (UCS) per table.

Status: M1.5 is done: Cassandra's architecture with key-value semantics
(newest write wins). Sections marked *planned* describe later milestones.

## 1. Goals

- **Cassandra's storage architecture**, so per-table operations (flush,
  snapshot, drop, truncate, repair, streaming) touch only that table and
  tables do not interfere with each other's compaction.
- **Cassandra's semantics**: last-write-wins by cell timestamp, timestamped
  tombstones, TTLs, `gc_grace_seconds`.
- **UCS compaction per table**, as in Cassandra 5, plus a time-window mode
  for insert-only TTL tables (*planned*).
- **Partitions never split across sstables in a level or shard.** Shard
  boundaries are token boundaries, so this follows from the layout.
- **Performance at least equal to the Pebble fork** on the workloads in
  section 10. Anything reused from Pebble that measurably slows the engine
  is replaced.

Non-goals: byte compatibility with Cassandra's sstable or commit log files.
Migrating data from real Cassandra goes through CQL, or through a read-only
Cassandra sstable importer if that becomes necessary.

## 2. Architecture

| Part | Per node | Per table | Cassandra equivalent |
|---|---|---|---|
| Commit log (WAL) | ✓ segmented; a segment is deleted once every table that wrote to it has flushed past it | | `CommitLog` |
| Memtable memory budget | ✓ when mutable memtables exceed a fraction of the budget, the largest is flushed; writes stall when the budget is exhausted | | `memtable_heap_space`, `MemtableCleaner` |
| WAL space limit | ✓ when live segments exceed it, the tables pinning the oldest segment are flushed | | `commitlog_total_space`, `flushDataFrom` |
| Manifest | ✓ one log of edits for all tables | | per-table directories and transaction logs |
| Memtable, flush | | ✓ | `ColumnFamilyStore` memtables |
| sstables | | ✓ in `tables/<table ID>/` | `<keyspace>/<table>-<id>/` |
| Compaction strategy and settings | | ✓ | per-table compaction options |

Keys inside a table are `token | partition key | clustering key`, with no
table ID: the table is the directory.

Per-table operations:

- **Create table:** a manifest edit and a new directory.
- **Drop table:** a manifest edit, then the directory is deleted; WAL entries
  for the table are ignored on replay.
- **Flush table:** rotates only that table's memtable.

## 3. What is reused from Pebble, and what is not

Pebble is pinned at upstream commit a995e89. Only its public packages are
imported.

| Area | Choice | Why |
|---|---|---|
| sstable format, writer, reader, block iterators | **Reuse** `pebble/sstable` (columnar format, `TableFormatPebblev6`) | The most optimised part of Pebble. |
| Bloom filters | **Reuse** `pebble/sstable/tablefilters/bloom` | Same reason. |
| Block cache | **Reuse** Pebble's cache (*planned*, M2) | Sharded, CLOCK-Pro, tuned for these blocks. |
| Files | **Reuse** `pebble/vfs`, `pebble/objstorage/objstorageprovider` | `vfs.NewMem` and `vfs.NewCrashableMem` make tests fast and crash tests possible. |
| WAL and manifest record framing | **Reuse** `pebble/record` | Checksummed, block-aligned records with torn-write detection. Buffered background writing with `record.LogWriter` is *planned* (M2). |
| Memtable | **Own**: skiplist, single writer per table, lock-free readers; keys and values in arena chunks | Pebble's `arenaskl` is internal. If ours stays slower (section 10), port it. |
| Commit log management, manifest, read path, compaction, merge function | **Own** | Must follow Cassandra's architecture and semantics. |

What Pebble's format holds for Cassandra:

- one key per row; the partition header (partition deletion, static row) is
  a key that sorts first in the partition;
- cells are encoded in the value against the table's schema, as Cassandra's
  `SerializationHeader` does (*planned*, data-model layer);
- the Bloom filter is built on the partition prefix through a `Comparer`
  whose `Split` returns its length (*planned*);
- `Statistics.db`-style metadata (min and max timestamps, max expiry time,
  tombstone histogram) is written by block property collectors (*planned*).

## 4. Cassandra semantics (*planned*, M2)

Cassandra resolves conflicts by **write timestamp**; Pebble resolves them by
sequence number. These differ whenever writes arrive out of order, which
repair, hinted handoff and client-supplied timestamps cause routinely. So:

- **Every row write is a merge operand** (Pebble's `MERGE` kind). A merge
  function reconciles cells by timestamp: the highest timestamp wins,
  tombstones are cells with timestamps, TTLs are checked against the current
  time.
- **Reads merge all versions of a row** across memtables and sstables. They
  stop early as Cassandra does: once the result holds a cell newer than the
  maximum timestamp of the remaining sstables, those are skipped.
- **Compaction merges operands with the same function.**
- **Pebble's `DELETE`, `SINGLEDEL` and `RANGEDEL` are never used for user
  data**: they hide older entries by sequence number, which would let an
  older tombstone delete newer data.

M1 and M1.5 still use `SET` and `DELETE` with newest-wins reads; they are the
key-value base the merge semantics are built on.

## 5. Write path

1. `Engine.Apply(batch)` appends the batch, whose entries may belong to
   several tables, to the WAL as one record and assigns it a block of
   sequence numbers. With `Sync`, the WAL is fsynced before returning.
2. Each entry is inserted into its table's mutable memtable. A batch becomes
   visible to readers atomically (readers ignore sequence numbers above the
   last fully applied batch).
3. Before the write, the engine checks:
   - **memory:** if mutable memtables use more than `MemTableCleanupThreshold`
     of `MemTableBudget`, the largest one is rotated and queued for flush; if
     all memtables, including ones waiting to flush, exceed the budget, the
     write waits;
   - **WAL:** the current segment is replaced when it exceeds
     `WALSegmentSize`; if live segments exceed `WALTotalSpace`, the tables
     whose memtables hold entries from the oldest segment are flushed.
4. Flushes run in the background, oldest first. A flush writes one sstable
   into the table's directory, records it and the table's flushed sequence
   number in the manifest, and deletes WAL segments no table still needs.

## 6. Manifest and recovery

The manifest is an append-only log of edits written with `pebble/record`. A
file named `CURRENT` names the live manifest and is replaced atomically. An
edit records:

- tables created and dropped;
- per table: sstables added (number, size, smallest and largest internal
  keys, sequence number range) and deleted, and the **flushed sequence
  number**: every WAL entry of the table at or below it is in sstables;
- the next file number, the last sequence number, and the lowest WAL segment
  still needed.

On open, the engine replays the manifest, then replays WAL segments at or
above the lowest needed one. An entry is applied only if its table exists and
its sequence number is above the table's flushed sequence number. The
recovered memtables are flushed, a new manifest is written, and old segments
are deleted.

## 7. Read path

- **Point lookup (M1):** a table's memtables from newest to oldest, then its
  sstables from newest to oldest, skipping sstables whose key range or Bloom
  filter excludes the key. With merge semantics (M2), versions are merged
  instead of the first one winning.
- **Iterators (*planned*, M3):** a merging iterator over a table's memtables
  and the sstables of the shards that overlap the bounds.
- **Snapshots (*planned*):** a sequence number plus a reference to a table's
  file set.

## 8. Compaction: UCS per table (*planned*, M3–M4)

The rules follow Cassandra 5's Unified Compaction Strategy, checked against
the cassandra-5.0 branch (commit 3e6d5ca):
`src/java/org/apache/cassandra/db/compaction/UnifiedCompactionStrategy.md`
and `unified/Controller.java`.

- **Scaling parameter `W`**, per table. `W < 0` is leveled (fanout
  `F = 2 - W`, threshold `T = 2`), `W > 0` is tiered (`F = T = 2 + W`),
  `W = 0` gives `F = T = 2`.
- **Defaults** (`CassandraRelevantProperties` and `cassandra_latest.yaml`):
  `scaling_parameters` T4 (tiered, `W = 2`), `target_sstable_size` 1 GiB,
  `base_shard_count` 4, `min_sstable_size` 100 MiB, `sstable_growth` 0.333,
  `max_sstables_to_compact` 64. Cassandra 5.0's shipped `cassandra.yaml`
  still defaults tables to STCS; UCS is the default only in
  `cassandra_latest.yaml`. This engine defaults every table to UCS T4,
  which behaves like STCS.
- **Density and levels:** an sstable's density is its size divided by the
  fraction of the token range it covers; level `n` starts at density
  `m·F^n`, where `m` is the flush size.
- **Trigger:** within a level, sstables are grouped into overlapping sets; a
  set with `T` or more sstables is compacted.
- **Sharding:** output is cut at token boundaries, so a partition is never
  split, and compactions in different shards run in parallel.
- **Fully expired sstables** are dropped without rewriting, checked every 10
  minutes (`CompactionController.getFullyExpiredSSTables`, used by every
  Cassandra strategy).
- **Time-window mode** (*planned*): sstables grouped by maximum write
  timestamp, for insert-only TTL tables, so whole files expire together.

## 9. Purging (*planned*, M5)

- **Purge rule (Cassandra's):** a tombstone older than `gc_grace_seconds` is
  dropped if no sstable outside the compaction overlaps its key with older
  data (checked with Bloom filters and per-file timestamp ranges) and no
  memtable holds the key.
- **Due sstables:** the Pebble fork's per-file summary and `IsFileDue` carry
  over, so files whose tombstones become due (by time or by repair
  watermark) are compacted even when UCS would not pick them.

## 10. Testing and benchmarks

- Unit tests per component, with `vfs.NewMem`.
- Crash tests with `vfs.NewCrashableMem`: crash at random points, dropping
  unsynced data, and check that every synced write survives and nothing else
  appears.
- A randomized model test (*planned*, M3): random operations against the
  engine and an in-memory model, with random flushes, compactions, reopens
  and crashes.
- Benchmarks against Pebble on the same workload, re-run after each
  milestone.

**First measurement (M1, in-memory file system, 100-byte values, Apple
M-series).** Pebble at a995e89 with default options, same workload:

| Benchmark | lsm (M1) | Pebble | Cause (from CPU and allocation profiles) | Fix (M2) |
|---|---|---|---|---|
| `Set`, no sync | 610–700 ns/op, 6 allocs, 1.3 KB (M1.5: 680–720 ns, 7 allocs: one more for the per-write table list) | 526 ns/op, 2 allocs, 133 B | One file write per WAL record (61% of allocated bytes), a new batch per write (19%), heap memtable nodes (15%); about 30% of CPU in garbage collection | Buffered WAL writing (`record.LogWriter`) with group commit; pooled batches; arena-allocated nodes |
| `Get` from sstables | 4.4–4.7 µs/op (M1.5: 4.7–5.1) | 2.4 µs/op | No block cache: every lookup re-reads, checksums and decompresses the index and data blocks | Pebble's block cache; reused point iterators |

## 11. Milestones

| Milestone | Contents |
|---|---|
| **M1** (done) | WAL, memtable, flush to sstables, manifest, recovery, point lookups, for one key space. |
| **M1.5** (done) | Cassandra's architecture: one WAL with segments and a memory budget, per-table memtables, sstables and directories, create and drop table. |
| M2 | Merge semantics (timestamp reconciliation); the speed fixes in section 10. |
| M3 | Merging iterator, UCS for one shard per table, model test. |
| M4 | UCS sharding, parallel compactions, time-window mode. |
| M5 | Purge rule, collector summaries, due-file picking. |
| M6 | Snapshots, streaming and repair support, benchmarks against the Pebble fork. |
