package objectpool

import (
	"fmt"
	"time"
)

// Pool ...
type Pool chan *Object

// NewPool ...
func NewPool(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- NewObject(i)
	}

	return &p
}

// Object ...
type Object struct {
	id int
}

// NewObject ...
func NewObject(id int) *Object {
	return &Object{id}
}

// Do ...
func (o *Object) Do() {
	fmt.Printf("Object %d done at %s\n", o.id, time.Now())
}
