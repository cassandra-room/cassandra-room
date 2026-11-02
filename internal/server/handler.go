package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/cassandra-room/cassandra-room/internal/cql"
	"github.com/datastax/go-cassandra-native-protocol/compression/lz4"
	"github.com/datastax/go-cassandra-native-protocol/compression/snappy"
	"github.com/datastax/go-cassandra-native-protocol/message"
	"github.com/datastax/go-cassandra-native-protocol/primitive"
)

func (s *Server) handleMessage(
	ctx context.Context,
	session *Session,
	version primitive.ProtocolVersion,
	msg message.Message,
) message.Message {
	switch msg.(type) {
	case *message.Options, *message.Startup:
	default:
		if session != nil && !session.isInitialized() {
			return &message.ProtocolError{
				ErrorMessage: "Connection is not initialized; STARTUP must be sent before any other request",
			}
		}
	}

	switch m := msg.(type) {
	case *message.Options:
		compOpts := []string{"lz4", "snappy"}
		if version >= primitive.ProtocolVersion5 {
			compOpts = []string{"lz4"}
		}
		return &message.Supported{
			Options: map[string][]string{
				message.StartupOptionCqlVersion:  {"3.4.7"},
				message.StartupOptionCompression: compOpts,
				message.SupportedProtocolVersions: {
					"3/v3",
					"4/v4",
					"5/v5",
				},
			},
		}

	case *message.Startup:
		if session != nil && session.isInitialized() {
			return &message.ProtocolError{
				ErrorMessage: "Connection is already initialized; STARTUP may only be sent once per connection",
			}
		}
		if m.Options != nil {
			if rawComp, ok := m.Options[message.StartupOptionCompression]; ok {
				compAlgo := strings.ToLower(strings.TrimSpace(rawComp))
				switch compAlgo {
				case "":
					if session != nil {
						session.setCompression("", nil)
					}
				case "lz4":
					if session != nil {
						session.setCompression("lz4", lz4.Compressor{})
					}
				case "snappy":
					if version >= primitive.ProtocolVersion5 {
						return &message.ProtocolError{
							ErrorMessage: "snappy compression is not supported in protocol v5; supported algorithm is lz4",
						}
					}
					if session != nil {
						session.setCompression("snappy", snappy.Compressor{})
					}
				default:
					return &message.ProtocolError{
						ErrorMessage: fmt.Sprintf("Unsupported compression algorithm %q; supported algorithms are lz4, snappy", rawComp),
					}
				}
			}
			if session != nil && m.Options[message.StartupOptionThrowOnOverload] == "1" {
				session.setThrowOnOverload(true)
			}
		}
		if session != nil {
			session.setInitializedWithVersion(version)
			if version.SupportsModernFramingLayout() {
				session.setPendingModernLayout(true)
			}
		}
		return &message.Ready{}

	case *message.Register:
		return &message.Ready{}

	case *message.Query:
		var ksOverride string
		if m.Options != nil {
			ksOverride = m.Options.Keyspace
		}
		stmt, err := s.parseCached(m.Query)
		if err != nil {
			return &message.SyntaxError{ErrorMessage: err.Error()}
		}
		_, _, bindNames := session.extractStatementMeta(ksOverride, stmt)
		values, err := resolveValues(m.Options, bindNames)
		if err != nil {
			return &message.Invalid{ErrorMessage: err.Error()}
		}
		return session.executeStatement(ctx, ksOverride, stmt, values)

	case *message.Prepare:
		stmt, err := s.parseCached(m.Query)
		if err != nil {
			return &message.SyntaxError{ErrorMessage: err.Error()}
		}
		if err := session.validateStatementKeyspace(m.Keyspace, stmt); err != nil {
			return &message.Invalid{ErrorMessage: err.Error()}
		}
		prep := s.registerPrepared(session, m.Keyspace, m.Query, stmt)
		return &message.PreparedResult{
			PreparedQueryId:   prep.ID,
			ResultMetadataId:  prep.ResultMetadataID,
			VariablesMetadata: prep.VariablesMetadata,
			ResultMetadata:    prep.ResultMetadata,
		}

	case *message.Execute:
		prep, ok := s.lookupPrepared(m.QueryId)
		if !ok {
			return &message.Unprepared{
				Id:           m.QueryId,
				ErrorMessage: fmt.Sprintf("unprepared query id %x", m.QueryId),
			}
		}
		values, err := resolveValues(m.Options, prep.BindNames)
		if err != nil {
			return &message.Invalid{ErrorMessage: err.Error()}
		}
		// A prepared statement is bound to the keyspace in which it was prepared.
		return session.executeStatement(ctx, prep.Keyspace, prep.Statement, values)

	case *message.Batch:
		for _, child := range m.Children {
			var childStmt cql.Statement
			var childKS string
			var bindNames []string

			if len(child.Id) > 0 {
				prep, ok := s.lookupPrepared(child.Id)
				if !ok {
					return &message.Unprepared{
						Id:           child.Id,
						ErrorMessage: fmt.Sprintf("unprepared batch child id %x", child.Id),
					}
				}
				childStmt = prep.Statement
				// Prepared child uses the keyspace captured at PREPARE time.
				childKS = prep.Keyspace
				bindNames = prep.BindNames
			} else {
				parsed, err := s.parseCached(child.Query)
				if err != nil {
					return &message.SyntaxError{ErrorMessage: err.Error()}
				}
				childStmt = parsed
				// Raw query child uses the Batch request's Keyspace override (if any).
				childKS = m.Keyspace
				_, _, bindNames = session.extractStatementMeta(childKS, childStmt)
			}

			if len(child.Values) != len(bindNames) {
				return &message.Invalid{
					ErrorMessage: fmt.Sprintf("expected %d bound values in batch child, got %d", len(bindNames), len(child.Values)),
				}
			}

			switch childStmt.(type) {
			case *cql.InsertStatement, *cql.UpdateStatement, *cql.DeleteStatement:
			default:
				return &message.Invalid{
					ErrorMessage: fmt.Sprintf("only INSERT, UPDATE, and DELETE are allowed in BATCH, got %T", childStmt),
				}
			}

			resp := session.executeStatement(ctx, childKS, childStmt, child.Values)
			if errMsg, isErr := resp.(message.Error); isErr {
				return errMsg
			}
		}
		return &message.VoidResult{}

	default:
		return &message.ProtocolError{
			ErrorMessage: fmt.Sprintf("unsupported request opcode: %v", msg.GetOpCode()),
		}
	}
}
