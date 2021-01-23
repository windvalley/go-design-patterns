package strategy

// Strategy ...
type Strategy interface {
	Apply(int, int) int
}

// Context ...
type Context struct {
	Strategy Strategy
}

// NewContext ...
func NewContext(strategy Strategy) *Context {
	return &Context{strategy}
}

// Operate ...
func (c *Context) Operate(value1, value2 int) int {
	return c.Strategy.Apply(value1, value2)
}

// Addition ...
type Addition struct{}

// Apply ...
func (*Addition) Apply(value1, value2 int) int {
	return value1 + value2
}

// Multiply ...
type Multiply struct{}

// Apply ...
func (*Multiply) Apply(value1, value2 int) int {
	return value1 * value2
}
