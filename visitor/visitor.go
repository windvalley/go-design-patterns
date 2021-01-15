package visitor

import "fmt"

// ComputerPart ...
type ComputerPart interface {
	Accept(ComputerPartVisitor)
}

// ComputerPartVisitor ...
type ComputerPartVisitor interface {
	Visit(ComputerPart)
}

// Keyboard ...
type Keyboard struct{}

// Accept ...
func (k *Keyboard) Accept(visitor ComputerPartVisitor) {
	visitor.Visit(k)
}

// Mouse ...
type Mouse struct{}

// Accept ...
func (m *Mouse) Accept(visitor ComputerPartVisitor) {
	visitor.Visit(m)
}

// Monitor ...
type Monitor struct{}

// Accept ...
func (m *Monitor) Accept(visitor ComputerPartVisitor) {
	visitor.Visit(m)
}

// Computer ...
type Computer struct {
	parts []ComputerPart
}

// NewComputer ...
func NewComputer(parts []ComputerPart) *Computer {
	return &Computer{parts}
}

// Accept ...
func (c *Computer) Accept(visitor ComputerPartVisitor) {
	for i := 0; i < len(c.parts); i++ {
		c.parts[i].Accept(visitor)
	}

	visitor.Visit(c)
}

// ComputerPartDisplayVisitor ...
type ComputerPartDisplayVisitor struct{}

// Visit ...
func (c *ComputerPartDisplayVisitor) Visit(computerPart ComputerPart) {
	switch v := computerPart.(type) {
	case *Keyboard:
		fmt.Printf("displaying keyboard: %#v\n", *v)
	case *Mouse:
		fmt.Printf("displaying mouse: %#v\n", *v)
	case *Monitor:
		fmt.Printf("displaying monitor: %#v\n", *v)
	case *Computer:
		fmt.Printf("displaying computer: %#v\n", *v)
	}
}

// ComputerPartCheckVisitor ...
type ComputerPartCheckVisitor struct{}

// Visit ...
func (c *ComputerPartCheckVisitor) Visit(computerPart ComputerPart) {
	switch v := computerPart.(type) {
	case *Keyboard:
		fmt.Printf("checking keyboard ok: %#v\n", *v)
	case *Mouse:
		fmt.Printf("checking mouse ok: %#v\n", *v)
	case *Monitor:
		fmt.Printf("checking monitor error: %#v\n", *v)
	case *Computer:
		fmt.Printf("checking computer not ok: %#v\n", *v)
	}
}
