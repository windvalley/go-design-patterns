package builder

import "testing"

func TestBuiler(t *testing.T) {
	expected := "partapartb"

	director := new(Director)
	builder := new(ConcreteBuilder)

	director.SetBuilder(builder)
	product := director.Build()

	result := product.Show()
	if result != expected {
		t.Errorf("got: %s, expected: %s", result, expected)
	}
}
