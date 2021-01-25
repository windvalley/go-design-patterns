package decorator

import "fmt"

// Component ...
type Component interface {
	Operate()
}

// Decorator ...
type Decorator interface {
	Component
	DecorateOn(c Component)
	Do()
}

type component1 struct {
	name string
}

// NewComponent1 ...
func NewComponent1() Component {
	return &component1{"component1"}
}

func (c *component1) Operate() {
	fmt.Printf("%s operate\n", c.name)
}

type decorator1 struct {
	c Component
}

// NewDecorator1 ...
func NewDecorator1() Decorator {
	return &decorator1{}
}

func (d *decorator1) DecorateOn(c Component) {
	d.c = c
}

func (d *decorator1) Do() {
	fmt.Printf("decorate on %v\n", d.c)
}

func (d *decorator1) Operate() {
	d.Do()
	d.c.Operate()
	d.Do()
}
