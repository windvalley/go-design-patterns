package buildoptions

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

// ServerBuilder ...
type ServerBuilder struct {
	Server
	err error
}

// Create ...
func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
	sb.Server.Addr = addr
	sb.Server.Port = port
	sb.Protocol = "tcp"
	sb.Timeout = 30 * time.Second
	sb.MaxConns = 1024
	return sb
}

// WithProtocol ...
func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	sb.Server.Protocol = protocol
	return sb
}

// WithMaxConn ...
func (sb *ServerBuilder) WithMaxConn(maxconn int) *ServerBuilder {
	if maxconn > 10240 {
		sb.err = errors.New("max connections cannot exceed 10240")
		return sb
	}
	sb.Server.MaxConns = maxconn
	return sb
}

// WithTimeout ...
func (sb *ServerBuilder) WithTimeout(timeout time.Duration) *ServerBuilder {
	sb.Server.Timeout = timeout
	return sb
}

// Build ...
func (sb *ServerBuilder) Build() (*Server, error) {
	return &sb.Server, sb.err
}
