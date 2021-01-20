package factorymethod

// Product ...
type Product interface {
	GetDescription() string
}

// Factory ...
type Factory interface {
	CreateProduct() Product
}

type product1Factory struct{}

// NewProduct1Factory ...
func NewProduct1Factory() Factory {
	return &product1Factory{}
}

func (*product1Factory) CreateProduct() Product {
	return new(product1)
}

type product2Factory struct{}

// NewProduct2Factory ...
func NewProduct2Factory() Factory {
	return &product2Factory{}
}

func (*product2Factory) CreateProduct() Product {
	return &product2{"product2"}
}

type product1 struct{}

func (*product1) GetDescription() string {
	return "this is product1"
}

type product2 struct {
	Name string
}

func (p *product2) GetDescription() string {
	return "this is " + p.Name
}
