package cql_test

import (
	"reflect"
	"testing"

	"github.com/cassandra-room/cassandra-room/internal/cql"
)

func TestParseCQL50Statements(t *testing.T) {
	validStatements := []struct {
		name  string
		query string
	}{
		{
			name:  "USE keyspace",
			query: `USE my_keyspace;`,
		},
		{
			name:  "CREATE KEYSPACE SimpleStrategy",
			query: `CREATE KEYSPACE IF NOT EXISTS cycling WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3} AND durable_writes = true;`,
		},
		{
			name: "CREATE TABLE with composite partition key, clustering order, static column, and collections",
			query: `CREATE TABLE IF NOT EXISTS cycling.cyclist_category (
				category text,
				points int,
				id uuid,
				lastname text STATIC,
				tags set<text>,
				scores list<int>,
				metadata map<text, frozen<tuple<int, text>>>,
				PRIMARY KEY ((category, points), id)
			) WITH CLUSTERING ORDER BY (id DESC) AND comment = 'test table';`,
		},
		{
			name: "CREATE TABLE with Cassandra 5.0 VECTOR type and Dynamic Data Masking",
			query: `CREATE TABLE IF NOT EXISTS ai.embeddings (
				id uuid PRIMARY KEY,
				secret text MASKED WITH mask_inner(1, 2),
				ssn text MASKED WITH DEFAULT,
				embedding vector<float, 1536>
			);`,
		},
		{
			name:  "CREATE CUSTOM INDEX (SAI)",
			query: `CREATE CUSTOM INDEX IF NOT EXISTS embedding_idx ON ai.embeddings (embedding) USING 'StorageAttachedIndex';`,
		},
		{
			name:  "SELECT with ANN vector search (Cassandra 5.0)",
			query: `SELECT id, secret FROM ai.embeddings ORDER BY embedding ANN OF [0.1, 0.2, 0.3] LIMIT 10;`,
		},
		{
			name:  "SELECT with collection slice range",
			query: `SELECT tags[1..3], tags[..2], tags[1..] FROM cycling.cyclist_category WHERE category = 'sprint' AND points = 100;`,
		},
		{
			name:  "SELECT with arithmetic, functions, writetime, ttl, cast, group by, per partition limit",
			query: `SELECT id, points + 10 AS bonus, WRITETIME(lastname), MAXWRITETIME(tags), TTL(lastname), CAST(points AS double) FROM cycling.cyclist_category WHERE category = ? AND points IN (10, 20, 30) GROUP BY category, points ORDER BY id ASC PER PARTITION LIMIT 5 LIMIT 100 ALLOW FILTERING;`,
		},
		{
			name:  "SELECT DISTINCT",
			query: `SELECT DISTINCT category, points FROM cycling.cyclist_category;`,
		},
		{
			name:  "SELECT JSON",
			query: `SELECT JSON id, lastname FROM cycling.cyclist_category WHERE category = :cat AND points = :pts;`,
		},
		{
			name:  "INSERT normal with IF NOT EXISTS, TTL, and TIMESTAMP",
			query: `INSERT INTO cycling.cyclist_category (category, points, id, lastname, tags) VALUES ('gc', 200, 123e4567-e89b-12d3-a456-426614174000, 'Doe', {'a', 'b'}) IF NOT EXISTS USING TTL 86400 AND TIMESTAMP 1700000000;`,
		},
		{
			name:  "INSERT JSON with DEFAULT UNSET",
			query: `INSERT INTO cycling.cyclist_category JSON '{"category": "gc", "points": 100}' DEFAULT UNSET;`,
		},
		{
			name:  "INSERT pg-style dollar-quoted string and duration literal",
			query: `INSERT INTO cycling.events (id, script, dur) VALUES (1, $$select 'hello'$$, 12h30m);`,
		},
		{
			name:  "UPDATE with counter/collection operations and LWT IF condition",
			query: `UPDATE cycling.cyclist_category USING TTL 300 SET lastname = 'Smith', tags = tags + {'c'}, metadata['k'] = (1, 'v'), points += 5 WHERE category = 'gc' AND points = 200 AND id = 123e4567-e89b-12d3-a456-426614174000 IF lastname = 'Doe' AND tags CONTAINS 'a';`,
		},
		{
			name:  "DELETE columns, map element, and UDT field with IF EXISTS",
			query: `DELETE lastname, metadata['k'], addr.zip FROM cycling.cyclist_category USING TIMESTAMP 1000 WHERE category = 'gc' AND points = 200 IF EXISTS;`,
		},
		{
			name: "BEGIN UNLOGGED BATCH",
			query: `BEGIN UNLOGGED BATCH USING TIMESTAMP 12345
				INSERT INTO cycling.cyclist_category (category, points, id) VALUES ('a', 1, 123e4567-e89b-12d3-a456-426614174000);
				UPDATE cycling.cyclist_category SET lastname = 'Lee' WHERE category = 'a' AND points = 1 AND id = 123e4567-e89b-12d3-a456-426614174000;
				DELETE FROM cycling.cyclist_category WHERE category = 'b' AND points = 2;
			APPLY BATCH;`,
		},
		{
			name:  "ALTER TABLE column mask",
			query: `ALTER TABLE IF EXISTS cycling.cyclist_category ALTER IF EXISTS lastname MASKED WITH DEFAULT;`,
		},
		{
			name:  "CREATE TYPE",
			query: `CREATE TYPE IF NOT EXISTS cycling.address (street text, zip int);`,
		},
		{
			name:  "DESCRIBE statements",
			query: `DESCRIBE FULL SCHEMA WITH INTERNALS;`,
		},
		{
			name:  "TRUNCATE table",
			query: `TRUNCATE TABLE cycling.cyclist_category;`,
		},
		{
			name:  "DROP TABLE",
			query: `DROP TABLE IF EXISTS cycling.cyclist_category;`,
		},
	}

	for _, tc := range validStatements {
		t.Run(tc.name, func(t *testing.T) {
			tree, err := cql.ParseQuery(tc.query)
			if err != nil {
				t.Fatalf("ParseQuery(%q) failed: %v", tc.query, err)
			}
			if tree == nil || tree.CqlStatement() == nil {
				t.Fatalf("ParseQuery(%q) returned nil statement", tc.query)
			}
		})
	}
}

func TestBindMarkersAndSelectorsAST(t *testing.T) {
	bindTests := []struct {
		name      string
		query     string
		wantNames []string
	}{
		{
			name:      "WHERE IN positional markers",
			query:     `SELECT * FROM ks.t WHERE id IN (?, ?, ?);`,
			wantNames: []string{"", "", ""},
		},
		{
			name:      "WHERE tuple equality markers",
			query:     `SELECT * FROM ks.t WHERE (a, b) = (?, ?);`,
			wantNames: []string{"", ""},
		},
		{
			name:      "WHERE and LIMIT markers",
			query:     `SELECT * FROM ks.t WHERE id = ? LIMIT ?;`,
			wantNames: []string{"", ""},
		},
		{
			name:      "UPDATE counter addition marker",
			query:     `UPDATE ks.t SET n = n + ? WHERE id = ?;`,
			wantNames: []string{"", ""},
		},
		{
			name:      "UPDATE shorthand += marker",
			query:     `UPDATE ks.t SET n += ? WHERE id = ?;`,
			wantNames: []string{"", ""},
		},
		{
			name:      "UPDATE map element key and value markers",
			query:     `UPDATE ks.t SET m[?] = ? WHERE id = ?;`,
			wantNames: []string{"", "", ""},
		},
		{
			name:      "UPDATE LWT IF condition marker",
			query:     `UPDATE ks.t SET a = ? WHERE id = ? IF b = ?;`,
			wantNames: []string{"", "", ""},
		},
		{
			name:      "INSERT VALUES and USING TTL markers",
			query:     `INSERT INTO ks.t (id, v) VALUES (?, ?) USING TTL ?;`,
			wantNames: []string{"", "", ""},
		},
		{
			name:      "Named bind markers in source order",
			query:     `SELECT * FROM ks.t WHERE id = :user_id AND status = :"ExactCase" LIMIT :lim;`,
			wantNames: []string{"user_id", "ExactCase", "lim"},
		},
	}

	for _, tc := range bindTests {
		t.Run(tc.name, func(t *testing.T) {
			stmt, err := cql.ParseStatement(tc.query)
			if err != nil {
				t.Fatalf("ParseStatement(%q) failed: %v", tc.query, err)
			}
			var gotNames []string
			switch s := stmt.(type) {
			case *cql.SelectStatement:
				gotNames = s.BindNames
			case *cql.InsertStatement:
				gotNames = s.BindNames
			case *cql.UpdateStatement:
				gotNames = s.BindNames
			case *cql.DeleteStatement:
				gotNames = s.BindNames
			default:
				t.Fatalf("unexpected statement type %T", stmt)
			}
			if !reflect.DeepEqual(gotNames, tc.wantNames) {
				t.Fatalf("BindNames = %#v, want %#v", gotNames, tc.wantNames)
			}
		})
	}

	// Verify SELECT column names (COUNT(*), aliases, unquoted identifiers, quoted identifiers)
	selStmt, err := cql.ParseStatement(`SELECT COUNT(*), id AS foo, Name, "Quoted", "Escaped""Quote", points + 10 AS "BonusPts" FROM "MyKS".riders;`)
	if err != nil {
		t.Fatalf("ParseStatement failed: %v", err)
	}
	sel, ok := selStmt.(*cql.SelectStatement)
	if !ok {
		t.Fatalf("expected *cql.SelectStatement, got %T", selStmt)
	}
	if sel.Keyspace != "MyKS" || sel.Table != "riders" {
		t.Fatalf("got keyspace=%q table=%q, want MyKS.riders", sel.Keyspace, sel.Table)
	}
	wantCols := []string{"count", "foo", "name", "Quoted", `Escaped"Quote`, "BonusPts"}
	if !reflect.DeepEqual(sel.Selectors, wantCols) {
		t.Fatalf("Selectors = %#v, want %#v", sel.Selectors, wantCols)
	}
}

func TestParseInvalidCQL(t *testing.T) {
	invalidQueries := []string{
		`SELECT FROM WHERE;`,
		`CREAT TABLE foo (id int PRIMARY KEY);`,
		`INSERT INTO foo (a, b) VALUES (1, );`,
		`SELECT * FROM ks.tbl WHERE a == 1;`,
	}

	for _, q := range invalidQueries {
		_, err := cql.ParseQuery(q)
		if err == nil {
			t.Errorf("expected syntax error for invalid query %q, got nil", q)
		}
	}
}
