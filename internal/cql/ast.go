package cql

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cassandra-room/cassandra-room/internal/cql/parser"
)

// Statement is a parsed, high-level CQL AST statement.
type Statement interface {
	statementNode()
}

// UseStatement represents `USE <keyspace>`.
type UseStatement struct {
	Keyspace string
}

func (*UseStatement) statementNode() {}

// CreateKeyspaceStatement represents `CREATE KEYSPACE [IF NOT EXISTS] <ks> WITH ...`.
type CreateKeyspaceStatement struct {
	Keyspace    string
	IfNotExists bool
}

func (*CreateKeyspaceStatement) statementNode() {}

// DropKeyspaceStatement represents `DROP KEYSPACE [IF EXISTS] <ks>`.
type DropKeyspaceStatement struct {
	Keyspace string
	IfExists bool
}

func (*DropKeyspaceStatement) statementNode() {}

// AlterKeyspaceStatement represents `ALTER KEYSPACE [IF EXISTS] <ks> WITH ...`.
type AlterKeyspaceStatement struct {
	Keyspace string
	IfExists bool
}

func (*AlterKeyspaceStatement) statementNode() {}

// ColumnDef represents a column definition in CREATE TABLE.
type ColumnDef struct {
	Name       string
	CQLType    string
	IsStatic   bool
	IsInlinePK bool
}

// CreateTableStatement represents `CREATE TABLE [IF NOT EXISTS] [ks.]<table> (...)`.
type CreateTableStatement struct {
	Keyspace       string
	Table          string
	IfNotExists    bool
	Columns        []ColumnDef
	PartitionKeys  []string
	ClusteringKeys []string
}

func (*CreateTableStatement) statementNode() {}

// DropTableStatement represents `DROP TABLE [IF EXISTS] [ks.]<table>`.
type DropTableStatement struct {
	Keyspace string
	Table    string
	IfExists bool
}

func (*DropTableStatement) statementNode() {}

// AlterTableStatement represents `ALTER TABLE [IF EXISTS] [ks.]<table> ...`.
type AlterTableStatement struct {
	Keyspace string
	Table    string
	IfExists bool
}

func (*AlterTableStatement) statementNode() {}

// TruncateStatement represents `TRUNCATE [TABLE] [ks.]<table>`.
type TruncateStatement struct {
	Keyspace string
	Table    string
}

func (*TruncateStatement) statementNode() {}

// Relation represents a single predicate in a WHERE clause (`col op term`).
type Relation struct {
	Column   string
	Operator string
	Value    string
	IsMarker bool
}

// SelectStatement represents `SELECT ... FROM [ks.]<table> [WHERE ...] [LIMIT ...]`.
type SelectStatement struct {
	Keyspace  string
	Table     string
	SelectAll bool
	Selectors []string
	Where     []Relation
	BindNames []string
}

func (*SelectStatement) statementNode() {}

// InsertStatement represents `INSERT INTO [ks.]<table> (...) VALUES (...)`.
type InsertStatement struct {
	Keyspace    string
	Table       string
	Columns     []string
	Values      []string
	IfNotExists bool
	IsJSON      bool
	BindNames   []string
}

func (*InsertStatement) statementNode() {}

// Assignment represents a single `col = expr` in an UPDATE statement.
type Assignment struct {
	Column string
	Value  string
}

// UpdateStatement represents `UPDATE [ks.]<table> SET ... WHERE ...`.
type UpdateStatement struct {
	Keyspace    string
	Table       string
	Assignments []Assignment
	Where       []Relation
	IfExists    bool
	BindNames   []string
}

func (*UpdateStatement) statementNode() {}

// DeleteStatement represents `DELETE [cols...] FROM [ks.]<table> WHERE ...`.
type DeleteStatement struct {
	Keyspace  string
	Table     string
	Columns   []string
	Where     []Relation
	IfExists  bool
	BindNames []string
}

func (*DeleteStatement) statementNode() {}

// BatchStatement represents `BEGIN [UNLOGGED|COUNTER] BATCH ... APPLY BATCH`.
type BatchStatement struct {
	Type       string // "LOGGED", "UNLOGGED", "COUNTER"
	Statements []Statement
	BindNames  []string
}

func (*BatchStatement) statementNode() {}

// RawStatement wraps any other valid CQL statement (e.g., ROLES, UDFs, DESCRIBE)
// together with its underlying ANTLR parse tree context.
type RawStatement struct {
	Context parser.ICqlStatementContext
}

func (*RawStatement) statementNode() {}

// ParseStatement parses a CQL query string using the ANTLR4 Cassandra 5.0 grammar
// and converts it into a structured Statement AST.
func ParseStatement(input string) (Statement, error) {
	tree, err := ParseQuery(input)
	if err != nil {
		return nil, err
	}
	return BuildStatement(tree.CqlStatement())
}

// BuildStatement converts an ANTLR ICqlStatementContext into a high-level Statement.
func BuildStatement(ctx parser.ICqlStatementContext) (Statement, error) {
	if ctx == nil {
		return nil, fmt.Errorf("empty CQL statement")
	}

	switch {
	case ctx.UseStatement() != nil:
		u := ctx.UseStatement()
		return &UseStatement{
			Keyspace: unquoteIdent(u.KeyspaceName().GetText()),
		}, nil

	case ctx.CreateKeyspaceStatement() != nil:
		ck := ctx.CreateKeyspaceStatement()
		return &CreateKeyspaceStatement{
			Keyspace:    unquoteIdent(ck.KeyspaceName().GetText()),
			IfNotExists: ck.K_IF() != nil,
		}, nil

	case ctx.DropKeyspaceStatement() != nil:
		dk := ctx.DropKeyspaceStatement()
		return &DropKeyspaceStatement{
			Keyspace: unquoteIdent(dk.KeyspaceName().GetText()),
			IfExists: dk.K_IF() != nil,
		}, nil

	case ctx.AlterKeyspaceStatement() != nil:
		ak := ctx.AlterKeyspaceStatement()
		return &AlterKeyspaceStatement{
			Keyspace: unquoteIdent(ak.KeyspaceName().GetText()),
			IfExists: ak.K_IF() != nil,
		}, nil

	case ctx.CreateTableStatement() != nil:
		return buildCreateTable(ctx.CreateTableStatement()), nil

	case ctx.DropTableStatement() != nil:
		dt := ctx.DropTableStatement()
		ks, tbl := splitColumnFamilyName(dt.ColumnFamilyName())
		return &DropTableStatement{
			Keyspace: ks,
			Table:    tbl,
			IfExists: dt.K_IF() != nil,
		}, nil

	case ctx.AlterTableStatement() != nil:
		at := ctx.AlterTableStatement()
		ks, tbl := splitColumnFamilyName(at.ColumnFamilyName())
		return &AlterTableStatement{
			Keyspace: ks,
			Table:    tbl,
			IfExists: len(at.AllK_IF()) > 0,
		}, nil

	case ctx.TruncateStatement() != nil:
		tr := ctx.TruncateStatement()
		ks, tbl := splitColumnFamilyName(tr.ColumnFamilyName())
		return &TruncateStatement{
			Keyspace: ks,
			Table:    tbl,
		}, nil

	case ctx.SelectStatement() != nil:
		return buildSelect(ctx.SelectStatement()), nil

	case ctx.InsertStatement() != nil:
		return buildInsert(ctx.InsertStatement()), nil

	case ctx.UpdateStatement() != nil:
		return buildUpdate(ctx.UpdateStatement()), nil

	case ctx.DeleteStatement() != nil:
		return buildDelete(ctx.DeleteStatement()), nil

	case ctx.BatchStatement() != nil:
		return buildBatch(ctx.BatchStatement())

	default:
		return &RawStatement{Context: ctx}, nil
	}
}

func buildCreateTable(ct parser.ICreateTableStatementContext) *CreateTableStatement {
	ks, tbl := splitColumnFamilyName(ct.ColumnFamilyName())
	stmt := &CreateTableStatement{
		Keyspace:    ks,
		Table:       tbl,
		IfNotExists: ct.K_IF() != nil,
	}

	if def := ct.TableDefinition(); def != nil {
		for _, colCtx := range def.AllTableColumns() {
			if colCtx.TablePartitionKey() != nil {
				for _, pk := range colCtx.TablePartitionKey().AllIdent() {
					stmt.PartitionKeys = append(stmt.PartitionKeys, unquoteIdent(pk.GetText()))
				}
				for _, ck := range colCtx.AllIdent() {
					stmt.ClusteringKeys = append(stmt.ClusteringKeys, unquoteIdent(ck.GetText()))
				}
				continue
			}

			idents := colCtx.AllIdent()
			if len(idents) > 0 && colCtx.ComparatorType() != nil {
				colName := unquoteIdent(idents[0].GetText())
				isPK := colCtx.K_PRIMARY() != nil
				stmt.Columns = append(stmt.Columns, ColumnDef{
					Name:       colName,
					CQLType:    colCtx.ComparatorType().GetText(),
					IsStatic:   colCtx.K_STATIC() != nil,
					IsInlinePK: isPK,
				})
				if isPK {
					stmt.PartitionKeys = append(stmt.PartitionKeys, colName)
				}
			}
		}
	}
	return stmt
}

func buildSelect(sel parser.ISelectStatementContext) *SelectStatement {
	ks, tbl := splitColumnFamilyName(sel.ColumnFamilyName())
	stmt := &SelectStatement{
		Keyspace:  ks,
		Table:     tbl,
		BindNames: collectBindVariables(sel),
	}

	if sc := sel.SelectClause(); sc != nil && sc.Selectors() != nil {
		sels := sc.Selectors().AllSelector()
		if len(sels) == 0 {
			stmt.SelectAll = true
		} else {
			for _, s := range sels {
				stmt.Selectors = append(stmt.Selectors, selectorColumnName(s))
			}
		}
	}

	if wc := sel.WhereClause(); wc != nil {
		stmt.Where = extractRelations(wc)
	}
	return stmt
}

// selectorColumnName resolves the result column name for a SELECT selector:
// - If an explicit alias (`AS ident`) is present, uses the unquoted/lowercased alias.
// - If the selector is `COUNT(*)`, returns `"count"` (matching Cassandra's column naming).
// - If the selector is a simple identifier (`Name` or `"Quoted"`), unquotes/lowercases it.
// - Otherwise (e.g. `writetime(x)`), uses the normalized expression text.
func selectorColumnName(s parser.ISelectorContext) string {
	if s.Noncol_ident() != nil {
		return unquoteIdent(s.Noncol_ident().GetText())
	}
	raw := s.UnaliasedSelector().GetText()
	if strings.EqualFold(raw, "count(*)") || strings.EqualFold(raw, "count(1)") {
		return "count"
	}
	if len(raw) >= 2 && raw[0] == '"' && raw[len(raw)-1] == '"' {
		return unquoteIdent(raw)
	}
	return strings.ToLower(raw)
}

func buildInsert(ins parser.IInsertStatementContext) *InsertStatement {
	ks, tbl := splitColumnFamilyName(ins.ColumnFamilyName())
	stmt := &InsertStatement{
		Keyspace:  ks,
		Table:     tbl,
		BindNames: collectBindVariables(ins),
	}

	if norm := ins.NormalInsertStatement(); norm != nil {
		stmt.IfNotExists = norm.K_IF() != nil
		for _, c := range norm.AllCident() {
			stmt.Columns = append(stmt.Columns, unquoteIdent(c.GetText()))
		}
		for _, t := range norm.AllTerm() {
			stmt.Values = append(stmt.Values, t.GetText())
		}
	} else if js := ins.JsonInsertStatement(); js != nil {
		stmt.IsJSON = true
		stmt.IfNotExists = js.K_IF() != nil
		if js.JsonValue() != nil {
			stmt.Values = append(stmt.Values, js.JsonValue().GetText())
		}
	}
	return stmt
}

func buildUpdate(upd parser.IUpdateStatementContext) *UpdateStatement {
	ks, tbl := splitColumnFamilyName(upd.ColumnFamilyName())
	stmt := &UpdateStatement{
		Keyspace:  ks,
		Table:     tbl,
		IfExists:  upd.K_EXISTS() != nil,
		BindNames: collectBindVariables(upd),
	}

	for _, op := range upd.AllColumnOperation() {
		col := unquoteIdent(op.Cident().GetText())
		val := ""
		if diff := op.ColumnOperationDifferentiator(); diff != nil {
			val = strings.TrimPrefix(diff.GetText(), "=")
		}
		stmt.Assignments = append(stmt.Assignments, Assignment{
			Column: col,
			Value:  val,
		})
	}

	if wc := upd.WhereClause(); wc != nil {
		stmt.Where = extractRelations(wc)
	}
	return stmt
}

func buildDelete(del parser.IDeleteStatementContext) *DeleteStatement {
	ks, tbl := splitColumnFamilyName(del.ColumnFamilyName())
	stmt := &DeleteStatement{
		Keyspace:  ks,
		Table:     tbl,
		IfExists:  del.K_EXISTS() != nil,
		BindNames: collectBindVariables(del),
	}

	if ds := del.DeleteSelection(); ds != nil {
		for _, op := range ds.AllDeleteOp() {
			stmt.Columns = append(stmt.Columns, op.GetText())
		}
	}

	if wc := del.WhereClause(); wc != nil {
		stmt.Where = extractRelations(wc)
	}
	return stmt
}

func buildBatch(b parser.IBatchStatementContext) (*BatchStatement, error) {
	batchType := "LOGGED"
	if b.K_UNLOGGED() != nil {
		batchType = "UNLOGGED"
	} else if b.K_COUNTER() != nil {
		batchType = "COUNTER"
	}

	stmt := &BatchStatement{
		Type:      batchType,
		BindNames: collectBindVariables(b),
	}
	for _, obj := range b.AllBatchStatementObjective() {
		switch {
		case obj.InsertStatement() != nil:
			stmt.Statements = append(stmt.Statements, buildInsert(obj.InsertStatement()))
		case obj.UpdateStatement() != nil:
			stmt.Statements = append(stmt.Statements, buildUpdate(obj.UpdateStatement()))
		case obj.DeleteStatement() != nil:
			stmt.Statements = append(stmt.Statements, buildDelete(obj.DeleteStatement()))
		}
	}
	return stmt, nil
}

// collectBindVariables walks the parse tree in source order and collects every
// positional (`?`) and named (`:name`) bind variable across the entire statement
// (including IN lists, tuples, collection keys, arithmetic expressions, LIMIT,
// USING TTL/TIMESTAMP, and IF conditions).
// Positional markers are represented by "" (empty string) and named markers by
// their unquoted/lowercased identifier name.
func collectBindVariables(tree antlr.Tree) []string {
	var names []string
	var walk func(n antlr.Tree)
	walk = func(n antlr.Tree) {
		if n == nil {
			return
		}
		if term, ok := n.(antlr.TerminalNode); ok {
			sym := term.GetSymbol()
			if sym != nil && sym.GetTokenType() == parser.CqlLexerQMARK {
				names = append(names, "")
			}
			return
		}

		// In Cql.g4, every named bind variable has the shape `':' noncol_ident`.
		childCount := n.GetChildCount()
		for i := 0; i < childCount; i++ {
			child := n.GetChild(i)
			if term, ok := child.(antlr.TerminalNode); ok && term.GetText() == ":" && i+1 < childCount {
				if identCtx, ok := n.GetChild(i + 1).(parser.INoncol_identContext); ok {
					names = append(names, unquoteIdent(identCtx.GetText()))
					i++ // skip the noncol_ident child
					continue
				}
			}
			walk(child)
		}
	}
	walk(tree)
	return names
}

func extractRelations(wc parser.IWhereClauseContext) []Relation {
	var rels []Relation
	for _, roe := range wc.AllRelationOrExpression() {
		rel := roe.Relation()
		if rel == nil {
			continue
		}
		var col, op, val string
		if rel.Cident() != nil {
			col = unquoteIdent(rel.Cident().GetText())
		}
		if rel.RelationType() != nil {
			op = rel.RelationType().GetText()
		} else if rel.K_IN() != nil {
			op = "IN"
		} else if rel.K_LIKE() != nil {
			op = "LIKE"
		}
		terms := rel.AllTerm()
		if len(terms) > 0 {
			val = terms[len(terms)-1].GetText()
		} else if rel.InMarker() != nil {
			val = rel.InMarker().GetText()
		} else if rel.SingleColumnInValues() != nil {
			val = rel.SingleColumnInValues().GetText()
		}
		markers := collectBindVariables(rel)
		rels = append(rels, Relation{
			Column:   col,
			Operator: op,
			Value:    val,
			IsMarker: len(markers) > 0,
		})
	}
	return rels
}

func splitColumnFamilyName(cf parser.IColumnFamilyNameContext) (keyspace, table string) {
	if cf == nil {
		return "", ""
	}
	if cf.KsName() != nil {
		keyspace = unquoteIdent(cf.KsName().GetText())
	}
	if cf.CfName() != nil {
		table = unquoteIdent(cf.CfName().GetText())
	}
	return keyspace, table
}

func unquoteIdent(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return strings.ReplaceAll(s[1:len(s)-1], `""`, `"`)
	}
	return strings.ToLower(s)
}
