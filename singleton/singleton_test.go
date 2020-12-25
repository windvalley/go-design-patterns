package singleton

import (
	"sync"
	"testing"
)

func TestNewSingleton(t *testing.T) {
	testTimes := 100
	wg := sync.WaitGroup{}
	wg.Add(testTimes)

	for i := 0; i < testTimes; i++ {
		go func() {
			defer wg.Done()
			instance := NewSingleton()
			instance.Increment()
		}()
	}

	wg.Wait()

	instance := NewSingleton()
	count := instance.GetCount()

	if count != testTimes {
		t.Errorf("got: %d, expected: %d", count, testTimes)
	}
}
