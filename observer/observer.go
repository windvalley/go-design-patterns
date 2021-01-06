package observer

import (
	"fmt"
)

type (
	// Event message
	Event struct {
		Data int64
	}

	// Subscriber is observer;
	// In the "real world", error handling would likely be implemented.
	Subscriber interface {
		// receive message
		OnNotify(Event)
	}

	// Publisher being observed
	Publisher interface {
		Register(Subscriber)
		Deregister(Subscriber)
		// send message
		Notify(Event)
	}
)

type subscriberA struct {
	id int
}

// NewSubscriberA ...
func NewSubscriberA(id int) Subscriber {
	return &subscriberA{id}
}

func (s *subscriberA) OnNotify(e Event) {
	fmt.Printf("subscriberA(observer %d) received: %d\n", s.id, e.Data)
}

type publisherA struct {
	subscribers map[Subscriber]struct{}
}

// NewPublisherA ...
func NewPublisherA() Publisher {
	return &publisherA{
		// or subscribers: make(map[Subscriber]struct{}),
		subscribers: map[Subscriber]struct{}{},
	}
}

func (p *publisherA) Register(s Subscriber) {
	p.subscribers[s] = struct{}{}
}

func (p *publisherA) Deregister(s Subscriber) {
	delete(p.subscribers, s)
}

func (p *publisherA) Notify(e Event) {
	for s := range p.subscribers {
		s.OnNotify(e)
	}
}
