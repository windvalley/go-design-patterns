package singleton

import "sync"

var (
	once              sync.Once
	singletonInstance *Singleton
)

// Singleton singleton
type Singleton struct {
	count int
	mutex sync.Mutex
}

// Increment increment count
func (s *Singleton) Increment() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.count++
}

// GetCount get count
func (s *Singleton) GetCount() int {
	return s.count
}

// NewSingleton create a singleton instance
func NewSingleton() *Singleton {
	once.Do(func() {
		singletonInstance = &Singleton{}
	})

	return singletonInstance
}
