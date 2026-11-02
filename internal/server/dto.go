package server

import (
	"context"

	"github.com/cassandra-room/cassandra-room/internal/cql"
	"github.com/datastax/go-cassandra-native-protocol/message"
	"github.com/datastax/go-cassandra-native-protocol/primitive"
)

// Request bundles the execution context passed to a QueryHandler for a parsed CQL statement.
type Request struct {
	// Ctx is canceled when the client connection closes or the server shuts down.
	Ctx      context.Context
	Session  *Session
	Keyspace string
	// Statement is the parsed CQL AST. Statements are shared across concurrent
	// requests via the statement and prepared caches and MUST be treated as
	// read-only (immutable) by QueryHandler implementations.
	Statement cql.Statement
	Values    []*primitive.Value
}

// preparedStmt stores a parsed CQL statement, its resolved keyspace at PREPARE time,
// and its variables/result metadata.
type preparedStmt struct {
	ID                []byte
	ResultMetadataID  []byte
	Keyspace          string
	RawCQL            string
	Statement         cql.Statement
	BindNames         []string
	VariablesMetadata *message.VariablesMetadata
	ResultMetadata    *message.RowsMetadata
}
