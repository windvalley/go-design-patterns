package simplefactory

// Product product
type Product interface {
	GetDescription() string
}

// NewProduct create a product
func NewProduct(name string) Product {
	switch name {
	case "product1":
		return new(product1)
	case "product2":
		return &product2{"product2"}
	default:
		return nil
	}
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
