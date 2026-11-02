package server

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/cassandra-room/cassandra-room/internal/cql"
	"github.com/datastax/go-cassandra-native-protocol/compression"
	"github.com/datastax/go-cassandra-native-protocol/compression/lz4"
	"github.com/datastax/go-cassandra-native-protocol/datatype"
	"github.com/datastax/go-cassandra-native-protocol/message"
	"github.com/datastax/go-cassandra-native-protocol/primitive"
	"github.com/datastax/go-cassandra-native-protocol/segment"
)

var (
	plainSegmentCodec = segment.NewCodec()
	lz4SegmentCodec   = segment.NewCodecWithCompression(lz4.Compressor{})
)

// Session tracks per-connection state (such as the active keyspace set by USE,
// the frame/segment compression algorithm negotiated during STARTUP, and whether
// the connection has transitioned to Protocol v5 segment framing).
type Session struct {
	mu                  sync.RWMutex
	initialized         bool
	version             primitive.ProtocolVersion
	pendingModernLayout bool
	modernLayout        bool
	keyspace            string
	throwOnOverload     bool
	compressionAlgo     string
	compressor          compression.Compressor
	handler             QueryHandler

	writeMu  sync.Mutex
	writeErr error
}

func (s *Session) executeStatement(
	ctx context.Context,
	ksOverride string,
	stmt cql.Statement,
	values []*primitive.Value,
) message.Message {
	if err := s.validateStatementKeyspace(ksOverride, stmt); err != nil {
		return &message.Invalid{ErrorMessage: err.Error()}
	}
	resolvedKS, _, _ := s.extractStatementMeta(ksOverride, stmt)

	if s.handler != nil {
		resp, err := s.handler(&Request{
			Ctx:       ctx,
			Session:   s,
			Keyspace:  resolvedKS,
			Statement: stmt,
			Values:    values,
		})
		if err != nil {
			var synErr *cql.SyntaxError
			if errors.As(err, &synErr) {
				return &message.SyntaxError{ErrorMessage: synErr.Error()}
			}
			return &message.Invalid{ErrorMessage: err.Error()}
		}
		if resp != nil {
			return resp
		}
	}

	switch st := stmt.(type) {
	case *cql.UseStatement:
		s.setKeyspace(st.Keyspace)
		return &message.SetKeyspaceResult{Keyspace: st.Keyspace}

	case *cql.CreateKeyspaceStatement:
		return &message.SchemaChangeResult{
			ChangeType: primitive.SchemaChangeTypeCreated,
			Target:     primitive.SchemaChangeTargetKeyspace,
			Keyspace:   st.Keyspace,
		}

	case *cql.AlterKeyspaceStatement:
		return &message.SchemaChangeResult{
			ChangeType: primitive.SchemaChangeTypeUpdated,
			Target:     primitive.SchemaChangeTargetKeyspace,
			Keyspace:   st.Keyspace,
		}

	case *cql.DropKeyspaceStatement:
		return &message.SchemaChangeResult{
			ChangeType: primitive.SchemaChangeTypeDropped,
			Target:     primitive.SchemaChangeTargetKeyspace,
			Keyspace:   st.Keyspace,
		}

	case *cql.CreateTableStatement:
		return &message.SchemaChangeResult{
			ChangeType: primitive.SchemaChangeTypeCreated,
			Target:     primitive.SchemaChangeTargetTable,
			Keyspace:   resolvedKS,
			Object:     st.Table,
		}

	case *cql.AlterTableStatement:
		return &message.SchemaChangeResult{
			ChangeType: primitive.SchemaChangeTypeUpdated,
			Target:     primitive.SchemaChangeTargetTable,
			Keyspace:   resolvedKS,
			Object:     st.Table,
		}

	case *cql.DropTableStatement:
		return &message.SchemaChangeResult{
			ChangeType: primitive.SchemaChangeTypeDropped,
			Target:     primitive.SchemaChangeTargetTable,
			Keyspace:   resolvedKS,
			Object:     st.Table,
		}

	case *cql.SelectStatement:
		var cols []*message.ColumnMetadata
		for i, sel := range st.Selectors {
			cols = append(cols, &message.ColumnMetadata{
				Keyspace: resolvedKS,
				Table:    st.Table,
				Name:     sel,
				Index:    int32(i),
				Type:     datatype.Varchar,
			})
		}
		return &message.RowsResult{
			Metadata: &message.RowsMetadata{
				ColumnCount: int32(len(cols)),
				Columns:     cols,
			},
			Data: message.RowSet{},
		}

	case *cql.InsertStatement, *cql.UpdateStatement, *cql.DeleteStatement, *cql.TruncateStatement, *cql.BatchStatement:
		return &message.VoidResult{}

	default:
		return &message.Invalid{
			ErrorMessage: fmt.Sprintf("unsupported CQL statement: %T", stmt),
		}
	}
}

// validateStatementKeyspace ensures that table-scoped statements have a valid
// keyspace (either explicitly in the CQL statement, via per-request/prepared override,
// or from the session's active USE keyspace) before PREPARE or execution.
func (s *Session) validateStatementKeyspace(ksOverride string, stmt cql.Statement) error {
	const errNoKeyspace = "no keyspace has been specified; USE a keyspace, or explicitly specify keyspace.tablename"

	switch st := stmt.(type) {
	case *cql.UseStatement, *cql.CreateKeyspaceStatement, *cql.AlterKeyspaceStatement, *cql.DropKeyspaceStatement:
		return nil
	case *cql.CreateTableStatement, *cql.AlterTableStatement, *cql.DropTableStatement, *cql.TruncateStatement,
		*cql.SelectStatement, *cql.InsertStatement, *cql.UpdateStatement, *cql.DeleteStatement:
		ks, _, _ := s.extractStatementMeta(ksOverride, stmt)
		if ks == "" {
			return errors.New(errNoKeyspace)
		}
		return nil
	case *cql.BatchStatement:
		for _, child := range st.Statements {
			childKS, _, _ := s.extractStatementMeta(ksOverride, child)
			if childKS == "" {
				return errors.New(errNoKeyspace)
			}
		}
		return nil
	default:
		return fmt.Errorf("unsupported CQL statement: %T", stmt)
	}
}

func (s *Session) extractStatementMeta(ksOverride string, stmt cql.Statement) (keyspace, table string, bindNames []string) {
	switch st := stmt.(type) {
	case *cql.UseStatement:
		return st.Keyspace, "", nil
	case *cql.CreateKeyspaceStatement:
		return st.Keyspace, "", nil
	case *cql.AlterKeyspaceStatement:
		return st.Keyspace, "", nil
	case *cql.DropKeyspaceStatement:
		return st.Keyspace, "", nil
	case *cql.CreateTableStatement:
		return s.resolveKeyspace(ksOverride, st.Keyspace), st.Table, nil
	case *cql.AlterTableStatement:
		return s.resolveKeyspace(ksOverride, st.Keyspace), st.Table, nil
	case *cql.DropTableStatement:
		return s.resolveKeyspace(ksOverride, st.Keyspace), st.Table, nil
	case *cql.TruncateStatement:
		return s.resolveKeyspace(ksOverride, st.Keyspace), st.Table, nil
	case *cql.SelectStatement:
		return s.resolveKeyspace(ksOverride, st.Keyspace), st.Table, st.BindNames
	case *cql.InsertStatement:
		return s.resolveKeyspace(ksOverride, st.Keyspace), st.Table, st.BindNames
	case *cql.UpdateStatement:
		return s.resolveKeyspace(ksOverride, st.Keyspace), st.Table, st.BindNames
	case *cql.DeleteStatement:
		return s.resolveKeyspace(ksOverride, st.Keyspace), st.Table, st.BindNames
	case *cql.BatchStatement:
		return s.resolveKeyspace(ksOverride, ""), "", st.BindNames
	default:
		return s.resolveKeyspace(ksOverride, ""), "", nil
	}
}

// resolveKeyspace returns the explicit keyspace in the CQL statement if present,
// otherwise the per-request/prepared keyspace override, otherwise the session's active keyspace.
func (s *Session) resolveKeyspace(ksOverride, explicit string) string {
	if explicit != "" {
		return explicit
	}
	if ksOverride != "" {
		return ksOverride
	}
	if s == nil {
		return ""
	}
	return s.Keyspace()
}

// isInitialized returns true if STARTUP has already completed on this session.
func (s *Session) isInitialized() bool {
	if s == nil {
		return false
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.initialized
}

// setInitializedWithVersion marks the session as having completed STARTUP negotiation
// with the given protocol version.
func (s *Session) setInitializedWithVersion(v primitive.ProtocolVersion) {
	if s == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.initialized = true
	s.version = v
}

// protocolVersion returns the protocol version negotiated during STARTUP (or 0 if not yet initialized).
func (s *Session) protocolVersion() primitive.ProtocolVersion {
	if s == nil {
		return 0
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.version
}

// setPendingModernLayout arms the transition to Protocol v5+ segment framing once
// the STARTUP READY response frame has been written.
func (s *Session) setPendingModernLayout(v bool) {
	if s == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.pendingModernLayout = v
}

// consumePendingModernLayout transitions the session to Protocol v5+ segment framing
// if a v5 STARTUP just completed.
func (s *Session) consumePendingModernLayout() bool {
	if s == nil {
		return false
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.pendingModernLayout {
		s.pendingModernLayout = false
		s.modernLayout = true
		return true
	}
	return false
}

// isModernLayout returns true if this connection has transitioned to Protocol v5+ segment framing.
func (s *Session) isModernLayout() bool {
	if s == nil {
		return false
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.modernLayout
}

// segmentCodec returns the v5 segment codec configured for this session.
func (s *Session) segmentCodec() segment.Codec {
	if s != nil {
		s.mu.RLock()
		algo := s.compressionAlgo
		s.mu.RUnlock()
		if algo == "lz4" {
			return lz4SegmentCodec
		}
	}
	return plainSegmentCodec
}

// Keyspace returns the currently active keyspace for this connection.
func (s *Session) Keyspace() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.keyspace
}

// setKeyspace updates the active keyspace for this connection.
func (s *Session) setKeyspace(ks string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.keyspace = ks
}

// getThrowOnOverload returns true if the client requested THROW_ON_OVERLOAD=1 in STARTUP.
func (s *Session) getThrowOnOverload() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.throwOnOverload
}

// setThrowOnOverload sets whether the connection should immediately receive
// ERROR OVERLOADED instead of applying backpressure when admission control is full.
func (s *Session) setThrowOnOverload(v bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.throwOnOverload = v
}

// compression returns the negotiated compression algorithm ("lz4", "snappy", or "")
// and its Compressor implementation for this session.
func (s *Session) compression() (string, compression.Compressor) {
	if s == nil {
		return "", nil
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.compressionAlgo, s.compressor
}

// setCompression configures the frame/segment compression algorithm for this session.
func (s *Session) setCompression(algo string, comp compression.Compressor) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.compressionAlgo = algo
	s.compressor = comp
}
