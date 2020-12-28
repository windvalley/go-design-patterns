package singleton

import "sync"

var (
	once        sync.Once
	fooInstance *Foo
)

// Foo singleton object
type Foo struct {
	count int
	mutex sync.Mutex
}

// NewFoo create a singleton instance
func NewFoo() *Foo {
	once.Do(func() {
		fooInstance = &Foo{}
	})

	return fooInstance
}

// Increment increment count
func (s *Foo) Increment() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.count++
}

// GetCount get count
func (s *Foo) GetCount() int {
	return s.count
}
