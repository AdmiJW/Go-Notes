package main

import (
	"fmt"
	"time"
)

// Goroutines are lightweight threads of execution.

func Goroutines() {
	f := func(from string) {
		for i := range 3 {
			fmt.Println(from, ":", i)
		}
	}

	// Suppose we have a function call `f(s)`. Here's how we'd call that in the usual way, running it synchronously.
	f("direct")

	// To invoke this function in a goroutine, use `go f(s)`. This new goroutine will execute concurrently with the calling one.
	go f("goroutine")

	// You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in separate goroutines now.
	// Wait for them to finish (for a more robust approach, use a `WaitGroup`).
	time.Sleep(time.Second)
	fmt.Println("done")

	// When we run this program, we see the output of the blocking call first, then the interleaved output of the two goroutines.
	// This interleaving reflects the goroutines being run concurrently by the Go runtime.
}