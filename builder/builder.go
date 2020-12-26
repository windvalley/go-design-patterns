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

// ConcreteBuilder a concrete builer
type ConcreteBuilder struct {
	product *Product
}

// NewProduct init a product
func (b *ConcreteBuilder) NewProduct() {
	b.product = new(Product)
}

// BuildPartA build part
func (b *ConcreteBuilder) BuildPartA() {
	b.product.content += "parta"
}

// BuildPartB build part
func (b *ConcreteBuilder) BuildPartB() {
	b.product.content += "partb"
}

// GetProduct get product
func (b *ConcreteBuilder) GetProduct() interface{} {
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
