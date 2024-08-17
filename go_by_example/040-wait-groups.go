package go_by_example

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multiple goroutines to finish, we can use a wait group.

func worker3(id int) {
	fmt.Printf("worker %d starting\n", id)
	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func WaitGroups() {
	// This WaitGroup is used to wait for all the goroutines launched here to finish.
	// Note: If a WaitGroup is passed to a function, it must be passed as a pointer.
	var waitGroup sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each.
	for i := 1; i <= 5; i++ {
		waitGroup.Add(1)

		// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
		// This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
		go func() {
			// A `defer` statement defers the execution of a function until the surrounding function returns.
			// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
			defer waitGroup.Done()
			worker3(i)
		}()
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they're done.
	waitGroup.Wait()
	fmt.Println("all workers done")

	// Note that this approach has no straightforward way to propogate errors from workers.
	// For more advanced use cases, consider using the `errgroup` package.
}