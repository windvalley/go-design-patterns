package objectpool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjectPool(t *testing.T) {
	p := NewPool(3)

	testSeconds := 1
	endTest := make(chan bool)
	go func() {
		for {
			select {
			case <-time.After(time.Duration(testSeconds) * time.Second):
				endTest <- true
				return
			}
		}
	}()

	// case 1
Label1:
	for {
		select {
		case obj := <-*p:
			obj.Do()
			*p <- obj // NOTE
		case <-endTest:
			fmt.Println("Test time has run out")
			break Label1
		default:
			fmt.Println("No more objects left, retry later or fail")
			break Label1
		}
	}

	// case 2
	for {
		select {
		case obj := <-*p:
			obj.Do()
		case <-endTest:
			fmt.Println("Test time has run out")
			return
		default:
			fmt.Println("No more objects left, retry later or fail")
			return
		}
	}
}
