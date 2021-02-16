package inversionofcontrol

import (
	"fmt"
	"testing"
)

func TestInversionOfControl(t *testing.T) {
	intSet := NewIntSet()

	intSet.Add(1)
	intSet.Add(2)
	fmt.Printf("add 1 and 2: %+v\n", intSet)

	intSet.Delete(1)
	fmt.Printf("delete 1: %+v\n", intSet)

	intSet.Undo()
	fmt.Printf("undo: %+v\n", intSet)

	intSet.Add(1)
	fmt.Printf("add 1: %+v\n", intSet)

	intSet.Delete(3)
	fmt.Printf("delete 3: %+v\n", intSet)

	fmt.Println("@@@@loop undo:")
	for {
		if err := intSet.Undo(); err != nil {
			fmt.Printf("error: %s\n", err)
			break
		}

		fmt.Printf("undo: %+v\n", intSet)
	}
}
