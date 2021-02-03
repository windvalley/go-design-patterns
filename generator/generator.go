package generator

// FibnacciGenerator ...
func FibnacciGenerator(n int) chan int {
	ret := make(chan int)

	go func(ret chan int) {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			// Blocks on the operation
			ret <- b
			a, b = b, a+b
		}

		close(ret)
	}(ret)

	return ret
}
