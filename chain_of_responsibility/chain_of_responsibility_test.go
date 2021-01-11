package chainofresponsibility

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	handler1 := NewConcreteHandler(1, 3, nil)
	handler2 := NewConcreteHandler(2, 3, handler1)
	handler3 := NewConcreteHandler(3, 3, handler2)

	handler3.Handle(3023)
	handler3.Handle(3322)
	handler3.Handle(888888)

	// error case
	handler1 = NewConcreteHandler(1, 2, nil)
	handler2 = NewConcreteHandler(3, 2, handler1)

	handler2.Handle(1001)
}
