package strategy

import "fmt"

func ExampleStrategy() {
	addStrategy := new(Addition)
	object := NewObject(addStrategy)
	result := object.Operate(1, 3)

	fmt.Println(result)

	multiStrategy := new(Multiplication)
	object = NewObject(multiStrategy)
	result = object.Operate(1, 3)

	fmt.Println(result)

	// Output:
	// 4
	// 3
}
