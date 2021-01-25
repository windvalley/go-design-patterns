package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	fmt.Println("Before decorate:")
	c1 := NewComponent1()
	c1.Operate()

	fmt.Println("After decorate:")
	d1 := NewDecorator1()
	d1.DecorateOn(c1)
	d1.Operate()
}
