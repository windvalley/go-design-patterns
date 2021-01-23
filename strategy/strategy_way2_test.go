package strategy

import (
	"fmt"
	"testing"
)

func TestStrategyWay2(t *testing.T) {
	addStrategy := new(Addition2)
	context := NewContext2(addStrategy)
	result := context.Apply(1, 3)

	fmt.Println(result)

	multiStrategy := new(Multiply2)
	// context = NewContext2(multiStrategy)
	context.Strategy2 = multiStrategy
	result = context.Apply(1, 3)

	fmt.Println(result)
}
