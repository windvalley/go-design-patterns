package parallelism

import (
	"fmt"
	"testing"
)

func TestParallelism(t *testing.T) {
	tasks := []string{"task1", "task2", "task3"}

	// ok case
	results, _ := getTaskResults(tasks, 5)
	fmt.Printf("%+v\n", results)

	// timeout/error case
	_, err := getTaskResults(tasks, 2)
	fmt.Println(err)
}
