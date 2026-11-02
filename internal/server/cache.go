package server

import (
	"container/list"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/cassandra-room/cassandra-room/internal/cql"
	"github.com/datastax/go-cassandra-native-protocol/datatype"
	"github.com/datastax/go-cassandra-native-protocol/message"
)

const numLRUShards = 64

type lruEntry[V any] struct {
	key   string
	value V
}

type lruShard[V any] struct {
	mu       sync.Mutex
	capacity int
	ll       *list.List
	items    map[string]*list.Element
}

// lruCache is a sharded, thread-safe bounded LRU cache used for prepared statements
// and parsed ASTs so memory usage stays strictly bounded and hot-path lookups
// do not contend on a single global mutex across connections.
type lruCache[V any] struct {
	shards [numLRUShards]lruShard[V]
}

func newLRUCache[V any](capacity int) *lruCache[V] {
	if capacity <= 0 {
		capacity = 1024
	}
	perShard := (capacity + numLRUShards - 1) / numLRUShards
	if perShard < 1 {
		perShard = 1
	}
	c := &lruCache[V]{}
	for i := range c.shards {
		c.shards[i] = lruShard[V]{
			capacity: perShard,
			ll:       list.New(),
			items:    make(map[string]*list.Element, perShard),
		}
	}
	return c
}

// parseCached parses a CQL query or retrieves its pre-parsed AST from the bounded LRU cache.
func (s *Server) parseCached(query string) (cql.Statement, error) {
	if stmt, ok := s.stmtCache.Get(query); ok {
		return stmt, nil
	}
	stmt, err := cql.ParseStatement(query)
	if err != nil {
		return nil, err
	}
	s.stmtCache.Put(query, stmt)
	return stmt, nil
}

func (s *Server) lookupPrepared(id []byte) (*preparedStmt, bool) {
	return s.preparedCache.Get(hex.EncodeToString(id))
}

func (s *Server) registerPrepared(session *Session, ksOverride, rawCQL string, stmt cql.Statement) *preparedStmt {
	ks, tbl, bindNames := session.extractStatementMeta(ksOverride, stmt)

	// Hash both the resolved keyspace and the raw CQL string (matching Cassandra's
	// QueryProcessor.computeId behavior so unqualified queries in different keyspaces
	// never collide).
	sum := sha256.Sum256([]byte(ks + "\x00" + rawCQL))
	queryID := sum[:16]
	metaSum := sha256.Sum256([]byte("meta:" + ks + "\x00" + rawCQL))
	resultMetaID := metaSum[:16]

	varCols := make([]*message.ColumnMetadata, len(bindNames))
	for i, name := range bindNames {
		if name == "" {
			name = fmt.Sprintf("bind_%d", i)
		}
		varCols[i] = &message.ColumnMetadata{
			Keyspace: ks,
			Table:    tbl,
			Name:     name,
			Index:    int32(i),
			Type:     datatype.Varchar,
		}
	}

	var resultCols []*message.ColumnMetadata
	if sel, ok := stmt.(*cql.SelectStatement); ok {
		for i, name := range sel.Selectors {
			resultCols = append(resultCols, &message.ColumnMetadata{
				Keyspace: ks,
				Table:    tbl,
				Name:     name,
				Index:    int32(i),
				Type:     datatype.Varchar,
			})
		}
	}

	prep := &preparedStmt{
		ID:               queryID,
		ResultMetadataID: resultMetaID,
		Keyspace:         ks,
		RawCQL:           rawCQL,
		Statement:        stmt,
		BindNames:        bindNames,
		VariablesMetadata: &message.VariablesMetadata{
			Columns: varCols,
		},
		ResultMetadata: &message.RowsMetadata{
			ColumnCount: int32(len(resultCols)),
			Columns:     resultCols,
		},
	}

	s.preparedCache.Put(hex.EncodeToString(queryID), prep)
	return prep
}

func (c *lruCache[V]) Get(key string) (V, bool) {
	s := c.shardFor(key)
	s.mu.Lock()
	defer s.mu.Unlock()
	if elem, ok := s.items[key]; ok {
		s.ll.MoveToFront(elem)
		return elem.Value.(*lruEntry[V]).value, true
	}
	var zero V
	return zero, false
}

func (c *lruCache[V]) Put(key string, value V) {
	s := c.shardFor(key)
	s.mu.Lock()
	defer s.mu.Unlock()

	if elem, ok := s.items[key]; ok {
		s.ll.MoveToFront(elem)
		elem.Value.(*lruEntry[V]).value = value
		return
	}

	if s.ll.Len() >= s.capacity {
		oldest := s.ll.Back()
		if oldest != nil {
			s.ll.Remove(oldest)
			delete(s.items, oldest.Value.(*lruEntry[V]).key)
		}
	}

	elem := s.ll.PushFront(&lruEntry[V]{key: key, value: value})
	s.items[key] = elem
}

func (c *lruCache[V]) shardFor(key string) *lruShard[V] {
	return &c.shards[fnv64a(key)&(numLRUShards-1)]
}
