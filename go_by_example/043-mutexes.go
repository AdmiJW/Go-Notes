package go_by_example

import (
	"fmt"
	"sync"
)

// In the previous example, we saw how to manage simple counter state using atomic operations.
// For more complex state we can use a mutex (mutual exclusion) to safely access data across multiple goroutines.

// Container holds a map of counters.
// Since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronize access.
// Note that mutexes must not be copied, so if this `struct` is passed around, it should be done by pointer.
type Container struct {
	mu sync.Mutex
	counters map[string]int
}

// Lock the mutex before accessing `counters`. If it is already locked, the goroutine will block until it is unlocked.
// Unlock it at the end of the function, after updating the map, possibly in a `defer` statement.
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func Mutexes() {
	// Note that the zero value of a mutex is usable as-is, so no initialization is necessary.
	c := Container {
		counters: map[string]int { "a": 0, "b": 0 },
	}
	var waitGroup sync.WaitGroup

	// This goroutine increments a named counter in a loop
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		waitGroup.Done()
	}

	// Run several goroutines concurrently.
	// Note that they all access the same `Container` instance,
	// and two of them access the same counter.
	waitGroup.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// Wait for all goroutines to finish.
	waitGroup.Wait()
	fmt.Println(c.counters)
}