package factorymethod

// Factory factory
type Factory interface {
	CreateProduct(name string) Product
}

// Product product
type Product interface {
	GetDescription() string
}

// NewFactory create factory
func NewFactory() Factory {
	return &factoryInstance{}
}

type factoryInstance struct {
}

func (*factoryInstance) CreateProduct(name string) Product {
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
