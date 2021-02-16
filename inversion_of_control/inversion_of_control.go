package inversionofcontrol

import "errors"

// Undo is a control object
type Undo []func()

// Add control functions to slice
func (u *Undo) Add(function func()) {
	*u = append(*u, function)
}

// Undo the last action
func (u *Undo) Undo() error {
	functions := *u
	if len(functions) == 0 {
		return errors.New("No functions to undo")
	}

	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // For garbage collection
	}

	*u = functions[:index]
	return nil
}

// IntSet is a business object
type IntSet struct {
	data map[int]bool
	undo Undo
}

// NewIntSet ...
func NewIntSet() *IntSet {
	return &IntSet{data: make(map[int]bool)}
}

// Undo ...
func (s *IntSet) Undo() error {
	return s.undo.Undo()
}

// Contains ...
func (s *IntSet) Contains(x int) bool {
	return s.data[x]
}

// Add ...
func (s *IntSet) Add(x int) {
	if s.Contains(x) {
		s.undo.Add(nil)
		return
	}

	s.data[x] = true
	s.undo.Add(func() { s.Delete(x) })
}

// Delete ...
func (s *IntSet) Delete(x int) {
	if !s.Contains(x) {
		s.undo.Add(nil)
		return
	}

	delete(s.data, x)
	s.undo.Add(func() { s.Add(x) })
}
