package simplefactory

import "testing"

func TestFactory_CreateProduct(t *testing.T) {
	cases := []struct {
		input  ProductType
		output string
	}{
		{Product1, "this is product1"},
		{Product2, "this is product2"},
		{Product3, "this is product3"},
	}

	factory := NewProductFactory()
	for _, v := range cases {
		product, err := factory.CreateProduct(v.input)
		if err != nil {
			t.Log(err)
			continue
		}

		if product.GetDescription() != v.output {
			t.Errorf("got: %s, expected: %s", product.GetDescription(), v.output)
		}
	}
}
