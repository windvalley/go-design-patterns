package boundedparallelism

import (
	"errors"
	"sync"
	"time"
)

type result struct {
	id  string
	res string
	err error
}

func batchDoTask(
	concurrency int,
	tasks []string,
	taskHandler func(task string) result,
	taskTimeout time.Duration,
	stop <-chan struct{},
) <-chan result {
	taskChan := make(chan string)

	go func() {
		defer close(taskChan)
		for _, task := range tasks {
			taskChan <- task
		}
	}()

	resChan := make(chan result)
	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			for task := range taskChan {
				done := make(chan struct{})
				go func() {
					defer close(done)
					result := taskHandler(task)
					resChan <- result
				}()

				select {
				case <-done:
				case <-time.After(taskTimeout * time.Second):
					resChan <- result{task, "", errors.New("task timeout")}
				case <-stop:
				}
			}

			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	return resChan
}

func taskHandler(task string) result {
	time.Sleep(3 * time.Second)
	res := task + " done"

	return result{task, res, nil}
}

func getTaskResults(tasks []string, taskTimeout time.Duration) ([]result, error) {
	stop := make(chan struct{})
	defer close(stop)

	resChan := batchDoTask(100, tasks, taskHandler, taskTimeout, stop)

	results := make([]result, 0)
	for r := range resChan {
		if r.err != nil {
			return nil, r.err
		}

		results = append(results, r)
	}

	return results, nil
}
