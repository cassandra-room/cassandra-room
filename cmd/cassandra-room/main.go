package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cassandra-room/cassandra-room/internal/server"
)

func main() {
	var (
		addr                  = flag.String("addr", ":9042", "CQL native transport listen address")
		maxConcurrentRequests = flag.Int("max-concurrent-requests", 0, "Global limit on concurrent active CQL requests (0 = default max(NumCPU*64, 256))")
		maxInFlightBytes      = flag.Int64("max-inflight-bytes", server.DefaultMaxInFlightBytes, "Global byte budget for in-flight request payloads")
		maxFrameSize          = flag.Int("max-frame-size", server.DefaultMaxFrameSize, "Maximum allowed CQL frame body size in bytes")
		writeTimeout          = flag.Duration("write-timeout", server.DefaultWriteTimeout, "Per-frame socket write deadline")
		shutdownTimeout       = flag.Duration("shutdown-timeout", 15*time.Second, "Graceful shutdown deadline for draining in-flight requests")
	)
	flag.Parse()

	srv := server.New(nil)
	if *maxConcurrentRequests > 0 {
		srv.SetMaxConcurrentRequests(*maxConcurrentRequests)
	}
	if *maxInFlightBytes > 0 {
		srv.SetMaxInFlightBytes(*maxInFlightBytes)
	}
	if *maxFrameSize > 0 {
		srv.SetMaxFrameSize(*maxFrameSize)
	}
	srv.SetWriteTimeout(*writeTimeout)

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", *addr, err)
	}
	log.Printf("cassandra-room listening for CQL v3/v4/v5 connections on %s", ln.Addr())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	serveErrCh := make(chan error, 1)
	go func() {
		serveErrCh <- srv.Serve(ln)
	}()

	select {
	case err := <-serveErrCh:
		if err != nil && !errors.Is(err, net.ErrClosed) {
			log.Fatalf("server stopped with error: %v", err)
		}
	case sig := <-sigCh:
		log.Printf("received %s, initiating graceful shutdown (timeout=%s)...", sig, *shutdownTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), *shutdownTimeout)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("graceful shutdown timed out: %v", err)
		} else {
			log.Printf("graceful shutdown complete")
		}
	}
}
//reviewed


