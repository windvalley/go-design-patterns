package singleton

import (
	"sync"
	"testing"
)

func TestNewBar(t *testing.T) {
	testTimes := 100
	wg := sync.WaitGroup{}
	wg.Add(testTimes)

	for i := 0; i < testTimes; i++ {
		go func() {
			defer wg.Done()
			instance := NewBar()
			instance.Increment()
		}()
	}

	wg.Wait()

	instance := NewBar()
	count := instance.GetCount()

	if count != testTimes {
		t.Errorf("got: %d, expected: %d", count, testTimes)
	}
}
