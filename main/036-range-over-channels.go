package main

import (
	"fmt"
	"time"
)

// In the previous example we saw how `for` and `range` provide iteration over basic data structures.
// We can also use this syntax to iterate over values received from a channel.

func RangeOverChannels() {
	// We'll iterate over 2 values in the `queue` channel.
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This `range` iterates over each element as it's received from `queue`.
	// Because we `close`d the channel above, the iteration terminates after receiving the 2 elements.
	for elem := range queue {
		fmt.Printf("Received element: %s\n", elem)
	}

	// This example also shows that it's possible to close a non-empty channel but still have the remaining values be received.

	// We can also close a non-empty channel and still receive the remaining values.
	interval := make(chan int)

	go func() {
		for i := range 5 {
			time.Sleep(1 * time.Second)
			interval <- i
		}
		close(interval)
	}()

	for i := range interval {
		fmt.Printf("Received interval: %d\n", i)
	}
	fmt.Println("Interval closed")
}