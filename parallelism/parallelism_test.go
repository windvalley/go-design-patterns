package parallelism

import (
	"fmt"
	"testing"
	"time"
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

func getTaskResults(tasks []string, taskTimeout time.Duration) ([]result, error) {
	stop := make(chan struct{})
	defer close(stop)

	resCh := batchDoTask(tasks, taskHandler, taskTimeout, stop)

	results := make([]result, 0)
	for r := range resCh {
		if r.err != nil {
			return nil, r.err
		}

		results = append(results, r)
	}

	return results, nil
}
