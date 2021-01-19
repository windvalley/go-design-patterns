package simplefactory

import "fmt"

// ProductType ...
type ProductType int

const (
	// Product1 ...
	Product1 ProductType = iota
	// Product2 ...
	Product2
	// Product3 ...
	Product3
)

// Product ...
type Product interface {
	GetDescription() string
}

// ProductFactory ...
type ProductFactory struct{}

// NewProductFactory ...
func NewProductFactory() *ProductFactory {
	return &ProductFactory{}
}

// CreateProduct ...
func (*ProductFactory) CreateProduct(productType ProductType) (Product, error) {
	var (
		p   Product
		err error
	)

	switch productType {
	case Product1:
		p = new(product1)
	case Product2:
		p = &product2{"product2"}
	default:
		err = fmt.Errorf("no such product: %v", productType)
	}

	return p, err
}

type product1 struct{}

func (*product1) GetDescription() string {
	return "this is product1"
}

type product2 struct {
	name string
}

func (p *product2) GetDescription() string {
	return "this is " + p.name
}
