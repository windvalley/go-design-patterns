package singleton

import (
	"sync"
	"testing"
)

func TestNewFoo(t *testing.T) {
	testTimes := 100
	wg := sync.WaitGroup{}
	wg.Add(testTimes)

	for i := 0; i < testTimes; i++ {
		go func() {
			defer wg.Done()
			instance := NewFoo()
			instance.Increment()
		}()
	}

	wg.Wait()

	instance := NewFoo()
	count := instance.GetCount()

	if count != testTimes {
		t.Errorf("got: %d, expected: %d", count, testTimes)
	}
}
