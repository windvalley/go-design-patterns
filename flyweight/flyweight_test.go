package flyweight

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewShapFactory(t *testing.T) {
	colors := []string{"blue", "red", "yellow", "green", "black"}
	shapeFactory := NewShapeFactory()

	fmt.Println("reuse 5 instances to generate 50 objects: ")
	for i := 0; i < 50; i++ {
		color := colors[rand.Intn(5)]
		circle := shapeFactory.GetShape(NewCircle, color).(*Circle)
		circle.SetRadius(rand.Intn(100) + 1)
		circle.Draw()
	}

	fmt.Println("shapeFactory pool as follows:")
	for k, v := range shapeFactory.pool {
		fmt.Printf("color: %s, %T: %+v\n", k, v, v)
	}
}
