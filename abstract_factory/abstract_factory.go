package abstractfactory

// Factory factory
type Factory interface {
	CreateProduct1() Product
	CreateProduct2() Product
}

// Product product
type Product interface {
	GetName() string
}

type concreteFactory struct {
}

// NewConcreteFactory create factory
func NewConcreteFactory() Factory {
	return &concreteFactory{}
}

func (f *concreteFactory) CreateProduct1() Product {
	return &product1{"product1"}
}

func (f *concreteFactory) CreateProduct2() Product {
	return &product2{"product2"}
}

type product1 struct {
	name string
}

func (p *product1) GetName() string {
	return p.name
}

type product2 struct {
	name string
}

func (p *product2) GetName() string {
	return p.name
}
