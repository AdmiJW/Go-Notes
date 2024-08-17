package main

import (
	"fmt"
	"time"
)

// We can use channels to synchronize execution across goroutines.
// Here's an example of using a blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish, you may prefer to use a `WaitGroup`.


func ChannelSynchronization() {
	// This is the function we'll run in a goroutine.
	// The `doneChannel` will be used to notify another goroutine that this function's work is done.
	worker := func(doneChannel chan bool) {
		fmt.Print("working...")
		time.Sleep(time.Second)
		fmt.Println("done")

		// Send a value to notify that we're done.
		doneChannel <- true
	}

	// Start a worker goroutine, giving it the channel to notify on.
	doneChannel := make(chan bool, 1)
	go worker(doneChannel)

	// Block until we receive a notification from the worker on the channel.
	<- doneChannel

	fmt.Println("the worker has finished")

	// If you removed the `<- doneChannel` line from this program, the program would exit before the worker even started.
}