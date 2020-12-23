package factorymethod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFactory(t *testing.T) {
	factory := NewFactory()

	for _, v := range []struct{ input, output string }{
		{"product1", "this is product1"},
		{"product2", "this is product2"},
		{"product3", ""},
	} {
		product := factory.CreateProduct(v.input)
		if product == nil {
			t.Logf("no such product: %s", v.input)
			continue
		}

		assert.Equal(
			t,
			product.GetDescription(),
			v.output,
			"they should be equal",
		)
	}
}
