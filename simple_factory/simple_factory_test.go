package simplefactory

import "testing"

func TestNewProduct(t *testing.T) {
	tests := []struct{ input, output string }{
		{"product1", "this is product1"},
		{"product2", "this is product2"},
	}

	for _, v := range tests {
		if product := NewProduct(v.input); product.GetDescription() != v.output {
			t.Errorf("got: %s, expected: %s", product.GetDescription(), v.output)
		}
	}
}
