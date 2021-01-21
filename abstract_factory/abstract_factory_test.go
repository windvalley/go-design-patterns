package abstractfactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFactory(t *testing.T) {
	factory := NewConcreteFactory()
	assert.Equal(t, factory.CreateProduct1().GetName(), "product1", "The should be equal")
	assert.Equal(t, factory.CreateProduct2().GetName(), "product2", "The should be equal")
}
