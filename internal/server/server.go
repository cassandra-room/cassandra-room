package server

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"runtime"
	"sync"
	"time"

	"github.com/cassandra-room/cassandra-room/internal/cql"
	"github.com/datastax/go-cassandra-native-protocol/frame"
	"github.com/datastax/go-cassandra-native-protocol/message"
	"github.com/datastax/go-cassandra-native-protocol/primitive"
	"github.com/datastax/go-cassandra-native-protocol/segment"
)

// QueryHandler is an optional callback that allows plugging custom execution logic
// (e.g., storage engine / coordinator) for parsed CQL statements.
type QueryHandler func(req *Request) (message.Message, error)

// Server is a concurrent CQL native protocol server that accepts multiple TCP connections,
// multiplexes up to 32k stream IDs per connection, enforces global admission control
// (both request count and in-flight byte budgets), caches parsed ASTs and prepared
// statements in sharded bounded LRUs, and decodes wire frames using
// datastax/go-cassandra-native-protocol.
type Server struct {
	codec             frame.RawCodec
	handler           QueryHandler
	maxStreamsPerConn int
	admission         *admissionController

	preparedCache *lruCache[*preparedStmt]
	stmtCache     *lruCache[cql.Statement]

	mu           sync.RWMutex
	maxFrameSize int
	writeTimeout time.Duration
	listeners    map[net.Listener]struct{}
	conns        map[net.Conn]struct{}
	closed       bool
	shuttingDown bool
	serverCtx    context.Context
	serverCancel context.CancelFunc
	connWG       sync.WaitGroup
}

// New creates a new CQL protocol server. If handler is nil, the server uses its
// built-in default execution router.
func New(handler QueryHandler) *Server {
	maxGlobal := runtime.NumCPU() * 64
	if maxGlobal < 256 {
		maxGlobal = 256
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		codec:             frame.NewRawCodec(),
		handler:           handler,
		maxStreamsPerConn: DefaultMaxConcurrentStreamsPerConn,
		admission:         newAdmissionController(maxGlobal, DefaultMaxInFlightBytes),
		preparedCache:     newLRUCache[*preparedStmt](DefaultPreparedCacheCapacity),
		stmtCache:         newLRUCache[cql.Statement](DefaultStatementCacheCapacity),
		maxFrameSize:      DefaultMaxFrameSize,
		writeTimeout:      DefaultWriteTimeout,
		listeners:         make(map[net.Listener]struct{}),
		conns:             make(map[net.Conn]struct{}),
		serverCtx:         ctx,
		serverCancel:      cancel,
	}
}

// ListenAndServe listens on the TCP network address addr (e.g. ":9042") and
// serves incoming CQL client connections concurrently.
func (s *Server) ListenAndServe(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.Serve(ln)
}

// Serve accepts incoming connections on the Listener ln, spawning a new goroutine
// for each connection. Transient accept errors (such as EMFILE) are retried with
// exponential backoff from 5ms up to 1s.
func (s *Server) Serve(ln net.Listener) error {
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		_ = ln.Close()
		return net.ErrClosed
	}
	s.listeners[ln] = struct{}{}
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.listeners, ln)
		s.mu.Unlock()
		_ = ln.Close()
	}()

	var tempDelay time.Duration
	for {
		conn, err := ln.Accept()
		if err != nil {
			s.mu.RLock()
			closed := s.closed
			s.mu.RUnlock()
			if closed || errors.Is(err, net.ErrClosed) {
				return nil
			}
			if isTemporaryAcceptError(err) {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if maxDelay := 1 * time.Second; tempDelay > maxDelay {
					tempDelay = maxDelay
				}
				timer := time.NewTimer(tempDelay)
				select {
				case <-timer.C:
				case <-s.serverCtx.Done():
					timer.Stop()
					return nil
				}
				continue
			}
			return err
		}
		tempDelay = 0

		go func(c net.Conn) {
			_ = s.serveConn(c)
		}(conn)
	}
}

// Shutdown gracefully stops the server: it closes listeners and sets an immediate
// read deadline on active connections so reader loops stop reading new frames while
// keeping sockets open for in-flight responses to be written. If ctx expires before
// in-flight requests finish, remaining request contexts are canceled and connections
// are closed.
func (s *Server) Shutdown(ctx context.Context) error {
	now := time.Now()
	s.mu.Lock()
	s.closed = true
	s.shuttingDown = true
	for ln := range s.listeners {
		_ = ln.Close()
	}
	for c := range s.conns {
		_ = c.SetReadDeadline(now)
	}
	s.mu.Unlock()

	done := make(chan struct{})
	go func() {
		s.connWG.Wait()
		close(done)
	}()

	select {
	case <-done:
		s.serverCancel()
		return nil
	case <-ctx.Done():
		s.serverCancel()
		s.mu.Lock()
		for c := range s.conns {
			_ = c.Close()
		}
		s.mu.Unlock()
		<-done
		return ctx.Err()
	}
}
