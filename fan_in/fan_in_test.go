package fanin

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	defer close(c1)
	defer close(c2)
	defer close(c3)

	out := Merge(c1, c2, c3)

	go func() {
		for item := range out {
			fmt.Println(item)
		}
	}()

	for i := 0; i < 5; i++ {
		c1 <- 1
		c2 <- 2
		c3 <- 3
	}
}
