package command

import "fmt"

// Command ...
type Command interface {
	Execute()
}

// Invoker ...
type Invoker struct {
	commands []Command
}

// StoreCommand ...
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

// ClearCommand ...
func (i *Invoker) ClearCommand() {
	i.commands = []Command{}
}

// ExecuteAllCommands ...
func (i *Invoker) ExecuteAllCommands() {
	for _, command := range i.commands {
		command.Execute()
	}

	i.ClearCommand()
}

// Receiver ...
type Receiver struct {
	name string
}

// NewReceiver ...
func NewReceiver(name string) *Receiver {
	return &Receiver{name}
}

// ExecuteCommand1 ...
func (r *Receiver) ExecuteCommand1() {
	fmt.Printf("Recever %s execute command1\n", r.name)
}

// ExecuteCommand2 ...
func (r *Receiver) ExecuteCommand2() {
	fmt.Printf("Receiver %s execute command2\n", r.name)
}

// Command1 ...
type Command1 struct {
	Receiver *Receiver
}

// Execute ...
func (c *Command1) Execute() {
	c.Receiver.ExecuteCommand1()
}

// Command2 ...
type Command2 struct {
	Receiver *Receiver
}

// Execute ...
func (c *Command2) Execute() {
	c.Receiver.ExecuteCommand2()
}
