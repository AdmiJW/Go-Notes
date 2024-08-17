package go_by_example

import (
	"fmt"
	"time"
)

// Go's `select` lets you wait on multiple channel operations.
// A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

func Select() {
	// For example, we'll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time, 
	// to simulate e.g. blocking RPC operations executing in concurrent goroutines.	
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We'll use `select` to await both of these values simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("received %s\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("received %s\n", msg2)
		}
	}
}