package interpreter

import (
	"fmt"
	"testing"
)

func getMaleExpression() Expression {
	robert := NewTerminalExpression("Robert")
	john := NewTerminalExpression("John")

	return NewOrExpression(robert, john)
}

func getMarriedWomanExpression() Expression {
	julie := NewTerminalExpression("Julie")
	married := NewTerminalExpression("Married")

	return NewAndExpression(julie, married)
}

func TestInterpreter(t *testing.T) {
	isMale := getMaleExpression()
	isMarriedWoman := getMarriedWomanExpression()

	fmt.Printf("John is male? %t\n", isMale.Interpret("John"))
	fmt.Printf("Julie is married woman? %t\n", isMarriedWoman.Interpret("Married Julie"))
}
