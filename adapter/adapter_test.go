package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	fmt.Println("@@@ case 1:")
	socket := NewConcreteSocket(3)
	plug := NewConcretePlug(3)
	socket.Charging(plug)

	fmt.Println("@@@ case 2:")
	plug = NewConcretePlug(2)
	socket.Charging(plug)

	fmt.Println("@@@ case 3:")
	adapter := NewConcreteAdapter(socket, 3)
	adapter.Charging(plug)

	fmt.Println("@@@ case 4:")
	adapter = NewConcreteAdapter(socket, 2)
	adapter.Charging(plug)
}
