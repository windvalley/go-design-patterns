package fanin

import "sync"

// Merge several channels in one channel
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
