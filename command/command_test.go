package command

import "testing"

func TestCommand(t *testing.T) {
	receiver1 := NewReceiver("receiver1")

	command1 := &Command1{receiver1}
	command2 := &Command2{receiver1}

	invoker := new(Invoker)

	invoker.StoreCommand(command1)
	invoker.StoreCommand(command2)

	invoker.ExecuteAllCommands()
	invoker.ExecuteAllCommands()
}
