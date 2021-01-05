package flyweight

import "fmt"

// Shape ...
type Shape interface {
	Draw()
}

// ShapeFactory ...
type ShapeFactory struct {
	pool map[string]Shape
}

// NewShapeFactory ...
func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{pool: make(map[string]Shape)}
}

// GetShape ...
func (sf *ShapeFactory) GetShape(newShape func(string) Shape, color string) Shape {
	shape, ok := sf.pool[color]
	if !ok {
		shape = newShape(color)
		sf.pool[color] = shape
	}
	return shape
}

// Circle ...
type Circle struct {
	color  string
	radius int
}

// NewCircle ...
func NewCircle(color string) Shape {
	return &Circle{color: color}
}

// SetRadius ...
func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}

// Draw ...
func (c *Circle) Draw() {
	fmt.Printf("draw a circle: [color: %s, radius: %d]\n", c.color, c.radius)
}
