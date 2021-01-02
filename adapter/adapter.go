package adapter

import "fmt"

// Socket ...
type Socket interface {
	GetHoleNumber() int
	Charging(p Plug)
}

// Plug ...
type Plug interface {
	GetPinNumber() int
}

// Adapter ...
type Adapter interface {
	Plug
	Socket
}

type concreteSocket struct {
	holeNumber int
}

// NewConcreteSocket ...
func NewConcreteSocket(holeNumber int) Socket {
	return &concreteSocket{holeNumber}
}

// GetHoleNumber ...
func (s *concreteSocket) GetHoleNumber() int {
	return s.holeNumber
}

// Charging ...
func (s *concreteSocket) Charging(p Plug) {
	plugPinNumber := p.GetPinNumber()
	if plugPinNumber != s.holeNumber {
		fmt.Printf(
			"socket hole number: %d, plug pin number: %d, the plug is not supported\n",
			s.holeNumber,
			plugPinNumber,
		)
		return
	}

	fmt.Printf(
		"socket hole number: %d, pin number: %d, charging...\n",
		s.holeNumber,
		plugPinNumber,
	)
}

type concretePlug struct {
	pinNumber int
}

// NewConcretePlug ...
func NewConcretePlug(pinNumber int) Plug {
	return &concretePlug{pinNumber}
}

func (p *concretePlug) GetPinNumber() int {
	return p.pinNumber
}

type concreteAdapter struct {
	Socket
	pinNumber  int
	holeNumber int
}

// NewConcreteAdapter ...
func NewConcreteAdapter(s Socket, holeNumber int) Adapter {
	return &concreteAdapter{Socket: s, holeNumber: holeNumber}
}

func (a *concreteAdapter) GetPinNumber() int {
	return a.pinNumber
}

func (a *concreteAdapter) Charging(p Plug) {
	plugPinNumber := p.GetPinNumber()

	if plugPinNumber != a.holeNumber {
		fmt.Printf(
			"adapter hole number: %d, plug pin number: %d, the adapter is not fit\n",
			a.holeNumber,
			plugPinNumber,
		)
		return
	}

	fmt.Printf("adapter hole number: %d, plug pin number: %d, ", a.holeNumber, plugPinNumber)
	a.pinNumber = a.Socket.GetHoleNumber()
	a.Socket.Charging(a)
}
