package decorator

// Object ...
type Object interface {
	GetScore() int
}

// ConcreteObject concrete object
type ConcreteObject struct {
	score int
}

// GetScore implement Object
func (cb *ConcreteObject) GetScore() int {
	return cb.score
}

// ConcreteObjectWithDecorate decorator
type ConcreteObjectWithDecorate struct {
	Object
	scorePlus int
}

// NewConcreteObjectWithDecorate ...
func NewConcreteObjectWithDecorate(cb Object, scorePlus int) Object {
	return &ConcreteObjectWithDecorate{cb, scorePlus}
}

// GetScore implement Object
func (cbd *ConcreteObjectWithDecorate) GetScore() int {
	return cbd.Object.GetScore() + cbd.scorePlus
}
