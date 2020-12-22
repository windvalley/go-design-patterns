package simplefactory

// Product product
type Product interface {
	GetDescription() string
}

// NewProduct create a product
func NewProduct(name string) Product {
	var p Product
	switch name {
	case "product1":
		p = new(product1)
	case "product2":
		p = &product2{"product2"}
	}

	return p
}

type product1 struct {
}

func (*product1) GetDescription() string {
	return "this is product1"
}

type product2 struct {
	Name string
}

func (p *product2) GetDescription() string {
	return "this is " + p.Name
}
