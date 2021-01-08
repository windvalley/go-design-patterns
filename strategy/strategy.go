package strategy

// Strategy ...
type Strategy interface {
	Apply(int, int) int
}

// Object ...
type Object struct {
	Strategy Strategy
}

// NewObject ...
func NewObject(strategy Strategy) *Object {
	return &Object{strategy}
}

// Operate ...
func (o *Object) Operate(value1, value2 int) int {
	return o.Strategy.Apply(value1, value2)
}

// Addition ...
type Addition struct{}

// Apply ...
func (Addition) Apply(value1, value2 int) int {
	return value1 + value2
}

// Multiplication ...
type Multiplication struct{}

// Apply ...
func (Multiplication) Apply(value1, value2 int) int {
	return value1 * value2
}
