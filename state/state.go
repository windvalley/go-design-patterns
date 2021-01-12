package state

import "fmt"

// State ...
type State interface {
	Set(context *Context)
}

// Context ...
type Context struct {
	state State
}

// SetState ...
func (c *Context) SetState(state State) {
	c.state = state
}

// GetState ...
func (c *Context) GetState() State {
	return c.state
}

type startState struct{}

// NewStartState ...
func NewStartState() State {
	return &startState{}
}

func (s *startState) Set(context *Context) {
	context.SetState(s)
	fmt.Println("Player is in start state")
}

func (s *startState) String() string {
	return "Start State"
}

type stopState struct{}

// NewStopState ...
func NewStopState() State {
	return &stopState{}
}

func (s *stopState) Set(context *Context) {
	context.SetState(s)
	fmt.Println("Player is in stop state")
}

func (s *stopState) String() string {
	return "Stop State"
}
