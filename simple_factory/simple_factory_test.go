package simplefactory

import "testing"

func TestNewProduct(t *testing.T) {
	tests := []struct{ input, output string }{
		{"product1", "this is product1"},
		{"product2", "this is product2"},
		{"product3", ""},
	}

	for _, v := range tests {
		product := NewProduct(v.input)
		if product == nil {
			t.Logf("no such product: %s", v.input)
			continue
		}

		if product.GetDescription() != v.output {
			t.Errorf("got: %s, expected: %s", product.GetDescription(), v.output)
		}
	}
}
