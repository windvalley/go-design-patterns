package parallelism

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
	tasks []string,
	taskHandler func(task string) result,
	taskTimeout time.Duration,
	stop <-chan struct{},
) <-chan result {
	resChan := make(chan result)

	go func() {
		var wg sync.WaitGroup
		for _, v := range tasks {
			wg.Add(1)
			go func(task string) {
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

				wg.Done()
			}(v)
		}

		go func() {
			wg.Wait()
			close(resChan)
		}()
	}()

	return resChan
}

func taskHandler(task string) result {
	time.Sleep(3 * time.Second)
	res := task + " done"

	return result{task, res, nil}
}
