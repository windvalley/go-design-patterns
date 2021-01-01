package decorator

import (
	"fmt"
	"testing"
)

func TestNewConreteObjectWithDecorate(t *testing.T) {
	object := &ConcreteObject{50}
	fmt.Println("pre decorate score: ", object.GetScore())

	newObject := NewConcreteObjectWithDecorate(object, 10)
	fmt.Println("after decorate score: ", newObject.GetScore())
}
