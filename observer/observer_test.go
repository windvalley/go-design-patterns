package observer

import (
	"testing"
	"time"
)

func TestPublisherSubscriber(t *testing.T) {
	publisher := NewPublisherA()

	subscriber1 := NewSubscriberA(1)
	subscriber2 := NewSubscriberA(2)
	subscriber3 := NewSubscriberA(3)

	publisher.Register(subscriber1)
	publisher.Register(subscriber2)
	publisher.Register(subscriber3)

	// A simple loop publishing the current Unix timestamp to subscribers.
	stop := time.NewTimer(3 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			publisher.Notify(Event{Data: t.UnixNano()})
			publisher.Deregister(subscriber3)
		}
	}
}
