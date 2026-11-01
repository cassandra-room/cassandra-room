package cql

import (
	"fmt"
	"strings"
	"sync"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cassandra-room/cassandra-room/internal/cql/parser"
)

// SyntaxError represents a CQL parsing or lexing error with line and column position.
type SyntaxError struct {
	Line    int
	Column  int
	Message string
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("line %d:%d %s", e.Line, e.Column, e.Message)
}

type errorListener struct {
	*antlr.DefaultErrorListener
	errors []*SyntaxError
}

func (l *errorListener) reset() {
	l.errors = l.errors[:0]
}

func (l *errorListener) SyntaxError(
	_ antlr.Recognizer,
	_ any,
	line, column int,
	msg string,
	_ antlr.RecognitionException,
) {
	l.errors = append(l.errors, &SyntaxError{
		Line:    line,
		Column:  column,
		Message: msg,
	})
}

func (l *errorListener) Err() error {
	if len(l.errors) == 0 {
		return nil
	}
	if len(l.errors) == 1 {
		return l.errors[0]
	}
	var b strings.Builder
	for i, err := range l.errors {
		if i > 0 {
			b.WriteString("; ")
		}
		b.WriteString(err.Error())
	}
	return fmt.Errorf("multiple syntax errors: %s", b.String())
}

// pooledParser holds a reusable ANTLR lexer, token stream, parser, and error listener
// to avoid heap allocations per query (similar to CockroachDB's parser pooling).
type pooledParser struct {
	lexer   *parser.CqlLexer
	stream  *antlr.CommonTokenStream
	parser  *parser.CqlParser
	errList *errorListener
}

var parserPool = sync.Pool{
	New: func() any {
		errList := &errorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
		lexer := parser.NewCqlLexer(nil)
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(errList)

		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewCqlParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(errList)

		return &pooledParser{
			lexer:   lexer,
			stream:  stream,
			parser:  p,
			errList: errList,
		}
	},
}

func acquireParser(input string) *pooledParser {
	pp := parserPool.Get().(*pooledParser)
	pp.errList.reset()
	is := antlr.NewInputStream(input)
	pp.lexer.SetInputStream(is)
	pp.stream.SetTokenSource(pp.lexer)
	pp.parser.SetTokenStream(pp.stream)
	return pp
}

func releaseParser(pp *pooledParser) {
	pp.errList.reset()
	parserPool.Put(pp)
}

// parseWithSLLFallback runs the ANTLR parser in fast SLL prediction mode first,
// falling back to full LL prediction mode only if SLL reports a syntax error.
func parseWithSLLFallback[T any](input string, ruleFn func(*parser.CqlParser) T) (T, error) {
	pp := acquireParser(input)
	defer releaseParser(pp)

	// Fast path: SLL mode
	pp.parser.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	tree := ruleFn(pp.parser)
	if len(pp.errList.errors) == 0 {
		return tree, nil
	}

	// Fallback path: reset stream and retry with full LL mode for ambiguous rules
	// (or to produce exact LL syntax error diagnostics).
	pp.errList.reset()
	is := antlr.NewInputStream(input)
	pp.lexer.SetInputStream(is)
	pp.stream.SetTokenSource(pp.lexer)
	pp.parser.SetTokenStream(pp.stream)
	pp.parser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLL)

	tree = ruleFn(pp.parser)
	if err := pp.errList.Err(); err != nil {
		var zero T
		return zero, err
	}
	return tree, nil
}

// ParseQuery parses a single CQL statement (with optional trailing semicolons)
// using the Cassandra 5.0 (CQL 3.4.7) ANTLR4 grammar.
func ParseQuery(input string) (parser.IQueryContext, error) {
	return parseWithSLLFallback(input, func(p *parser.CqlParser) parser.IQueryContext {
		return p.Query()
	})
}

// ParseBatchScript parses one or more semicolon-separated CQL statements.
func ParseBatchScript(input string) (parser.IRootContext, error) {
	return parseWithSLLFallback(input, func(p *parser.CqlParser) parser.IRootContext {
		return p.Root()
	})
}
