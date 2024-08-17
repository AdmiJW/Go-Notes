package main

import (
	"fmt"
	"time"
)

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine.

func Channels() {
	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the `channel <-` syntax.
	// Here we send `"ping"` to the `messages` channel we made above, from a new goroutine.
	go func() {
		fmt.Println("Goroutine started")

		time.Sleep(3 * time.Second)
		fmt.Println("Sending 1st message to channel")
		messages <- "ping"
		fmt.Println("1st message sent to channel")

		time.Sleep(3 * time.Second)
		fmt.Println("Sending 2nd message to channel")
		messages <- "pong"
		fmt.Println("2nd message sent to channel")
	}()

	// The `<-channel` syntax receives a value from the channel.
	// Here we'll receive the `"ping"` message we sent above and print it out.
	fmt.Println("Started receiving messages from channel")

	fmt.Printf("Received 1st message: %s\n", <-messages)
	fmt.Printf("Received 2nd message: %s\n", <-messages)

	// When we run the program, the `"ping"` message is successfully passed from one goroutine to another via our channel.
	// By default, sends and receives are blocking until both the sender and receiver are ready.
	// This property allowed us to wait at the end of our program for the `"ping"` message without having to use any other synchronization.
}