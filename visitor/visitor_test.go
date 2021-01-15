package visitor

import "testing"

func TestVisitor(t *testing.T) {
	keyboard := new(Keyboard)
	mouse := new(Mouse)
	monitor := new(Monitor)
	computer := NewComputer([]ComputerPart{keyboard, mouse, monitor})

	displayVisitor := new(ComputerPartDisplayVisitor)
	computer.Accept(displayVisitor)

	checkVisitor := new(ComputerPartCheckVisitor)
	computer.Accept(checkVisitor)
}
