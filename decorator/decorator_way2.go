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
func (co *ConcreteObject) GetScore() int {
	return co.score
}

// ConcreteObjectWithDecorate decorator
type ConcreteObjectWithDecorate struct {
	Object
	scorePlus int
}

// NewConcreteObjectWithDecorate ...
func NewConcreteObjectWithDecorate(co Object, scorePlus int) Object {
	return &ConcreteObjectWithDecorate{co, scorePlus}
}

// GetScore implement Object
func (cod *ConcreteObjectWithDecorate) GetScore() int {
	return cod.Object.GetScore() + cod.scorePlus
}
