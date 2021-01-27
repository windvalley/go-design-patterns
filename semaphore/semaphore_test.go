package semaphore

import (
	"fmt"
	"testing"
	"time"
)

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
