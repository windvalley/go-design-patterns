package semaphore

import (
	"errors"
	"time"
)

var (
	// ErrNoTickets ...
	ErrNoTickets = errors.New("semaphore acquired timeout")
	// ErrIllegalRelease ...
	ErrIllegalRelease = errors.New("can't release the semaphore without acquiring it first")
)

// Semaphore ...
type Semaphore interface {
	Acquire() error
	Release() error
}

type concreteSemaphore struct {
	sema    chan struct{}
	timeout time.Duration
}

func (s *concreteSemaphore) Acquire() error {
	select {
	case s.sema <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *concreteSemaphore) Release() error {
	select {
	case _ = <-s.sema:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

// NewConcreteSemaphore ...
func NewConcreteSemaphore(tickets int, timeout time.Duration) Semaphore {
	return &concreteSemaphore{
		sema:    make(chan struct{}, tickets),
		timeout: timeout,
	}
}
