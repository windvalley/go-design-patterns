package boundedparallelism

import (
	"fmt"
	"testing"
)

func TestBoundedParallelism(t *testing.T) {
	tasks := []string{"task1", "task2", "task3"}

	// ok case
	results, _ := getTaskResults(tasks, 5)
	fmt.Printf("%+v\n", results)

	// timeout/error case
	_, err := getTaskResults(tasks, 1)
	fmt.Println(err)
}
