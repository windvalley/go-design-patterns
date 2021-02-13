package timingfunctions

import (
	"log"
	"time"
)

// Duration measuring function execution time
func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)

	log.Printf("%s lasted %s", name, elapsed)
}

func functionDemo() {
	defer Duration(time.Now(), "functionDemo")

	sum := 0
	for i := 0; i <= 100000; i++ {
		sum += i
	}
}
