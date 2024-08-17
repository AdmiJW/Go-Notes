package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// The primary mechanism for managing state in Go is communication over channels.
// We saw this for example with worker pools. There are a few other options for managing state though.
// Here, we'll look at using the `sync/atomic` package for atomic counters accessed by multiple goroutines.

func AtomicCounters() {
	// We'll use an atomic integer type to represent our (always-positive) counter.
	var ops atomic.Uint64
	// As example, we declare another regular integer counter.
	var regularOps uint64

	// A WaitGroup will help us wait for all goroutines to finish their work.
	var workGroup sync.WaitGroup

	// We'll start 50 goroutines that each increment the counter exactly 1000 times.
	for i := 0; i < 50; i++ {
		workGroup.Add(1)

		go func() {
			defer workGroup.Done()

			for c := 0; c < 1000; c++ {
				ops.Add(1)
				regularOps++
			}
		}()
	}

	// Wait until all goroutines are done.
	workGroup.Wait()

	// Although here no goroutines are writing to `ops`, using `Load` it's safe to read a value even while other goroutines are (atomically) updating it.
	fmt.Println("ops:", ops.Load())
	fmt.Println("regularOps:", regularOps)

	// We expect to get exactly 50,000 operations. 
	// Had we used a non-atomic integer and incremented it with `ops++`,
	// we'd likely get a different number less than expected 50,000, changing between runs,
	// because the goroutines would interfere with each other.
	// Moreover, we'd get data race failures when running with the `-race` flag.
}