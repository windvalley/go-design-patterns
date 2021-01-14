package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	books := []string{"foo", "bar", "foobar"}

	bookContainer := NewBookContainer(books)

	for iterator := bookContainer.GetIterator(); iterator.HasNext(); iterator.Next() {
		book := iterator.CurrentValue()
		fmt.Println(book)
	}
}
