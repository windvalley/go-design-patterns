package fanout

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	c1 := make(chan int)
	defer close(c1)

	cs := Split(c1, 3)

	go func() {
		for id, c := range cs {
			go func(id int, c chan int) {
				for v := range c {
					fmt.Printf("channel %d do the work %d\n", id, v)
				}
			}(id, c)
		}
	}()

	for i := 0; i < 10; i++ {
		c1 <- i
	}
}
