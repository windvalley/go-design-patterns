package generator

import (
	"fmt"
	"testing"
)

func TestFibnacciGenerator(t *testing.T) {
	for i := range FibnacciGenerator(20) {
		fmt.Println(i)
	}
}
