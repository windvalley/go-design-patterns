package singleton

import "sync"

var barInstance *Bar

func init() {
	barInstance = &Bar{}
}

// Bar singleton object
type Bar struct {
	count int
	mutex sync.Mutex
}

// NewBar return the init instance
func NewBar() *Bar {
	return barInstance
}

// Increment increment count
func (s *Bar) Increment() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.count++
}

// GetCount get count
func (s *Bar) GetCount() int {
	return s.count
}
