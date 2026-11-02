package server

import (
	"errors"
	"fmt"
	"net"
	"os"
	"syscall"

	"github.com/datastax/go-cassandra-native-protocol/message"
	"github.com/datastax/go-cassandra-native-protocol/primitive"
)

func isTemporaryAcceptError(err error) bool {
	if errors.Is(err, syscall.EMFILE) ||
		errors.Is(err, syscall.ENFILE) ||
		errors.Is(err, syscall.ECONNABORTED) ||
		errors.Is(err, syscall.EINTR) {
		return true
	}
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		return true
	}
	return false
}

func isTimeoutError(err error) bool {
	if errors.Is(err, os.ErrDeadlineExceeded) {
		return true
	}
	var ne net.Error
	return errors.As(err, &ne) && ne.Timeout()
}

// resolveValues validates and extracts bound values from QueryOptions, supporting
// both PositionalValues and NamedValues (ordered according to bindNames).
func resolveValues(opts *message.QueryOptions, bindNames []string) ([]*primitive.Value, error) {
	want := len(bindNames)
	if opts == nil {
		if want > 0 {
			return nil, fmt.Errorf("expected %d bound values, got 0", want)
		}
		return nil, nil
	}

	hasPositional := len(opts.PositionalValues) > 0
	hasNamed := len(opts.NamedValues) > 0

	if want == 0 {
		if hasPositional || hasNamed {
			got := len(opts.PositionalValues) + len(opts.NamedValues)
			return nil, fmt.Errorf("expected 0 bound values, got %d", got)
		}
		return nil, nil
	}

	if hasNamed {
		out := make([]*primitive.Value, want)
		for i, name := range bindNames {
			if name == "" {
				return nil, fmt.Errorf("named values cannot be used with positional '?' bind marker at index %d", i)
			}
			val, ok := opts.NamedValues[name]
			if !ok {
				return nil, fmt.Errorf("missing value for bind variable :%s", name)
			}
			out[i] = val
		}
		return out, nil
	}

	if len(opts.PositionalValues) != want {
		return nil, fmt.Errorf("expected %d bound values, got %d", want, len(opts.PositionalValues))
	}
	return opts.PositionalValues, nil
}

// fnv64a computes the 64-bit FNV-1a hash of s without allocating.
func fnv64a(s string) uint64 {
	const (
		offset64 = 14695981039346656037
		prime64  = 1099511628211
	)
	var h uint64 = offset64
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= prime64
	}
	return h
}
