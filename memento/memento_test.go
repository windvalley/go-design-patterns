package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	originator := new(Originator)
	careTaker := new(CareTaker)

	originator.SetState("State #1")
	originator.SetState("State #2")
	careTaker.Add(originator.SaveStateToMemento())

	originator.SetState("State #3")
	careTaker.Add(originator.SaveStateToMemento())

	originator.SetState("State #4")
	fmt.Printf("Originator current state: %s\n", originator.GetState())

	originator.GetStateFromMemento(careTaker.Get(0))
	fmt.Printf("Originator recoverd to first saved state: %s\n", originator.GetState())

	originator.GetStateFromMemento(careTaker.Get(1))
	fmt.Printf("Originator recoverd to second saved state: %s\n", originator.GetState())
}
