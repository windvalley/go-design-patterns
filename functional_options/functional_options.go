package functionaloptions

import (
	"errors"
	"time"
)

// Server ...
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

// NewServer ...
func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {
	server := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1024,
	}

	for _, option := range options {
		option(&server)
	}

	if server.MaxConns > 10240 {
		return nil, errors.New("max connections cannot exceed 10240")
	}

	// ...

	return &server, nil
}

// WithProtocol ...
func WithProtocol(p string) func(*Server) {
	return func(s *Server) {
		s.Protocol = p
	}
}

// WithTimeout ...
func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

// WithMaxConns ...
func WithMaxConns(maxconns int) func(*Server) {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}
