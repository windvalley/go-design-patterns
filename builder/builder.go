package builder

// Builder builder
type Builder interface {
	NewProduct()
	BuildPartA()
	BuildPartB()
	GetProduct() interface{}
}

// Director director
type Director struct {
	builder Builder
}

// SetBuilder set a builder
func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

// Build build
func (d *Director) Build() *Product {
	d.builder.NewProduct()
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	return d.builder.GetProduct().(*Product)
}

type concreteBuilder struct {
	product *Product
}

// NewConcreteBuilder ...
func NewConcreteBuilder() Builder {
	return &concreteBuilder{}
}

func (b *concreteBuilder) NewProduct() {
	b.product = new(Product)
}

func (b *concreteBuilder) BuildPartA() {
	b.product.content += "parta"
}

func (b *concreteBuilder) BuildPartB() {
	b.product.content += "partb"
}

func (b *concreteBuilder) GetProduct() interface{} {
	return b.product
}

// Product product
type Product struct {
	content string
}

// Show product
func (p *Product) Show() string {
	return p.content
}
