package semaphore

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func doTask(task string) {
	time.Sleep(1 * time.Second)
	fmt.Printf("do %s, %s done\n", task, task)
}

func TestConcurrencyCtrl(t *testing.T) {
	tasks := []string{"task1", "task2", "task3", "task4", "task5", "task6"}

	tickets, timeout := 3, 10*time.Second
	sema := NewConcreteSemaphore(tickets, timeout)

	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(task string) {
			defer wg.Done()
			if err := sema.Acquire(); err != nil {
				fmt.Printf("task '%s' acquire semaphore failed: %s\n", task, err)
				return
			}

			doTask(task)

			if err := sema.Release(); err != nil {
				fmt.Printf("task '%s' release semaphore failed: %s\n", task, err)
			}
		}(task)
	}

	wg.Wait()
}

func TestSemaphore(t *testing.T) {
	// case 1
	tickets, timeout := 1, 2*time.Second
	sema := NewConcreteSemaphore(tickets, timeout)

	if err := sema.Acquire(); err != nil {
		fmt.Printf("Acquire semaphore failed: %s\n", err)
		return
	}

	fmt.Println("This block do important work")

	if err := sema.Release(); err != nil {
		fmt.Printf("Release semaphore failed: %s\n", err)
		return
	}

	// case 2
	tickets, timeout = 0, 2*time.Second
	sema = NewConcreteSemaphore(tickets, timeout)

	if err := sema.Acquire(); err != nil {
		if err != ErrNoTickets {
			panic(err)
		}

		fmt.Printf("No tickets left, can't work, error: %s\n", err)
	}

	// case 3
	tickets, timeout = 1, 2*time.Second
	sema = NewConcreteSemaphore(tickets, timeout)

	if err := sema.Release(); err != nil {
		fmt.Printf("Release semaphore failed: %s\n", err)
		return
	}
}
