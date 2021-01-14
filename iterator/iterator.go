package iterator

// Iterator ...
type Iterator interface {
	CurrentValue() interface{}
	HasNext() bool
	Next()
}

// Container ...
type Container interface {
	GetIterator() Iterator
}

type bookContainer struct {
	books []string
}

// NewBookContainer ...
func NewBookContainer(books []string) Container {
	return &bookContainer{books}
}

func (c *bookContainer) GetIterator() Iterator {
	return &bookIterator{
		cursor:    0,
		container: c,
	}
}

type bookIterator struct {
	cursor    int
	container *bookContainer
}

func (i *bookIterator) CurrentValue() interface{} {
	return i.container.books[i.cursor]
}

func (i *bookIterator) HasNext() bool {
	if i.cursor < len(i.container.books) {
		return true
	}

	return false
}

func (i *bookIterator) Next() {
	if i.HasNext() {
		i.cursor++
	}
}
