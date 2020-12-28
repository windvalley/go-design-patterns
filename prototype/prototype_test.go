package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	productA := NewProduct("foo")
	productA.Build()

	productB := productA.Clone()

	assert.NotSame(t, productB, productA, "they cannot be the same")
	assert.Equal(t, productA, productB, "they should be equal")
	assert.Equal(t, productA.GetName(), productB.GetName(), "they should be equal")
	assert.Equal(t, productA.GetContent(), productB.GetContent(), "they should be equal")
}
