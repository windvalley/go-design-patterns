package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	facade := NewFacade()

	fmt.Println("@@@ API1: ")
	facade.API1()
	fmt.Println("@@@ API2: ")
	facade.API2()
	fmt.Println("@@@ API3: ")
	facade.API3()
}
