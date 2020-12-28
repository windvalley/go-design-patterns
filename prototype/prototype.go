package prototype

// Prototype prototype
type Prototype interface {
	Build()
	GetName() string
	GetContent() string
	Clone() Prototype
}

// Product product
type Product struct {
	name    string
	content string
}

// NewProduct get an instance of Product
func NewProduct(name string) Prototype {
	return &Product{
		name: name,
	}
}

// Build assuming this is a time-consuming task
func (p *Product) Build() {
	p.content = p.name + " some content"
}

// GetName get product name
func (p *Product) GetName() string {
	return p.name
}

// GetContent get product content
func (p *Product) GetContent() string {
	return p.content
}

// Clone get a cloned product
func (p *Product) Clone() Prototype {
	copyP := *p
	return &copyP
}
