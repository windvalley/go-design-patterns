package strategy

// Strategy2 ...
type Strategy2 interface {
	Apply(int, int) int
}

// Context2 ...
type Context2 struct {
	Strategy2
}

// NewContext2 ...
func NewContext2(strategy Strategy2) *Context2 {
	return &Context2{strategy}
}

// Addition2 ...
type Addition2 struct{}

// Apply ...
func (*Addition2) Apply(value1, value2 int) int {
	return value1 + value2
}

// Multiply2 ...
type Multiply2 struct{}

// Apply ...
func (*Multiply2) Apply(value1, value2 int) int {
	return value1 * value2
}
