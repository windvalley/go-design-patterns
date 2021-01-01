package decorator

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// FuncFoo ...
type FuncFoo func(int)

// ProfileTimingDecorate decorator
func ProfileTimingDecorate(fn FuncFoo) FuncFoo {
	return func(n int) {
		defer fmt.Printf(
			"%s lasted %s\n",
			runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(),
			time.Since(time.Now()),
		)

		fn(n)
	}
}

func funcFooDemo(n int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}

	fmt.Println(sum)
}
