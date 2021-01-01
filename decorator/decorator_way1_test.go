package decorator

import (
	"testing"
)

func TestProfileTimingDecorate(t *testing.T) {
	number := 100000

	t.Log("origin: ")
	funcFooDemo(number)

	t.Log("after decorator: ")
	fn := ProfileTimingDecorate(funcFooDemo)
	fn(number)
}
