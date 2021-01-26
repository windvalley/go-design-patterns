package prototype

// Product ...
type Product interface {
	Build()
	GetName() string
	GetContent() string
	Clone() Product
}

type productA struct {
	name    string
	content string
}

// NewProductA ...
func NewProductA(name string) Product {
	return &productA{
		name: name,
	}
}

// Build assuming this is a time-consuming task
func (p *productA) Build() {
	p.content = p.name + " some content"
}

// GetName ...
func (p *productA) GetName() string {
	return p.name
}

// GetContent ...
func (p *productA) GetContent() string {
	return p.content
}

// Clone get a cloned product
func (p *productA) Clone() Product {
	copyP := *p
	return &copyP
}
