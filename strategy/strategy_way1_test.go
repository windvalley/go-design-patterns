package strategy

import "fmt"

func ExampleStrategy() {
	addStrategy := new(Addition)
	object := NewContext(addStrategy)
	result := object.Operate(1, 3)

	fmt.Println(result)

	multiStrategy := new(Multiply)
	object = NewContext(multiStrategy)
	result = object.Operate(1, 3)

	fmt.Println(result)

	// Output:
	// 4
	// 3
}
