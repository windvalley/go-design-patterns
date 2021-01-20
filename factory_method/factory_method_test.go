package factorymethod

import (
	"fmt"
	"testing"
)

func TestNewFactory(t *testing.T) {
	product1Factory := NewProduct1Factory()
	product1 := product1Factory.CreateProduct()
	fmt.Println(product1.GetDescription())

	product2Factory := NewProduct2Factory()
	product2 := product2Factory.CreateProduct()
	fmt.Println(product2.GetDescription())
}
