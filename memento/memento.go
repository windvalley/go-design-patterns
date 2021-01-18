package memento

// Memento ...
type Memento struct {
	state string
}

// SetState ...
func (m *Memento) SetState(state string) {
	m.state = state
}

// GetState ...
func (m *Memento) GetState() string {
	return m.state
}

// Originator ...
type Originator struct {
	state string
}

// SetState ...
func (o *Originator) SetState(state string) {
	o.state = state
}

// GetState ...
func (o *Originator) GetState() string {
	return o.state
}

// SaveStateToMemento ...
func (o *Originator) SaveStateToMemento() *Memento {
	memento := new(Memento)
	memento.SetState(o.state)
	return memento
}

// GetStateFromMemento ...
func (o *Originator) GetStateFromMemento(memento *Memento) {
	o.state = memento.GetState()
}

// CareTaker ...
type CareTaker struct {
	mementoList []*Memento
}

// Add ...
func (c *CareTaker) Add(state *Memento) {
	c.mementoList = append(c.mementoList, state)
}

// Get ...
func (c *CareTaker) Get(index int) *Memento {
	return c.mementoList[index]
}
