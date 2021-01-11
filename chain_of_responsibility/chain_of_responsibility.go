package chainofresponsibility

import "fmt"

// Handler ...
type Handler interface {
	HaveRight(requestID int) bool
	Handle(requestID int)
}

// concreteHandler ...
type concreteHandler struct {
	id           int
	handlerCount int
	next         Handler
}

// NewConcreteHandler ...
func NewConcreteHandler(
	handlerID, handlerCount int,
	nextHandler Handler) Handler {
	return &concreteHandler{handlerID, handlerCount, nextHandler}
}

func (h *concreteHandler) HaveRight(requestID int) bool {
	if requestID%h.handlerCount+1 == h.id {
		return true
	}
	return false
}

func (h *concreteHandler) Handle(requestID int) {
	if h.HaveRight(requestID) {
		fmt.Printf("Handler %d handled request id %d\n", h.id, requestID)
		return
	}

	if h.next != nil {
		h.next.Handle(requestID)
		return
	}

	fmt.Printf("No handler could handle request id %d\n", requestID)
}
