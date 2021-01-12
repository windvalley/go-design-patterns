package state

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	ctx := new(Context)
	start := NewStartState()
	stop := NewStopState()

	start.Set(ctx)
	state := ctx.GetState()
	fmt.Println(state)

	stop.Set(ctx)
	state = ctx.GetState()
	fmt.Println(state)
}
