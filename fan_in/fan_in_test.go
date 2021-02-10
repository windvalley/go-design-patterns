package fanin

import (
	"fmt"
	"testing"
)

func TestMergeChannel(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	out := MergeChannel(c1, c2, c3)

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

	close(c1)
	close(c2)
	close(c3)
}
