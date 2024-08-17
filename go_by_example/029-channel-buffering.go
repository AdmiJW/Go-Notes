package go_by_example

import (
	"fmt"
)

// By default, channels are unbuffered.
// This means that they will only accept sends (`chan <-`) if there is a corresponding receive (`<- chan`) ready to receive the sent value.
// Buffered channels accept a limited number of values without a corresponding receiver for those values.
// By specifying the buffer length when creating the channel, we can enable buffering.

func ChannelBuffering() {
	// Here, we make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	go func() {
		fmt.Println("Goroutine started")

		fmt.Println("Sending 1st message to channel")
		messages <- "ping"
		fmt.Println("1st message sent to channel")

		fmt.Println("Sending 2nd message to channel")
		messages <- "pong"
		fmt.Println("2nd message sent to channel")
	}()

	// Later, we receive these two messages as usual.
	fmt.Println("Started receiving messages from channel")

	fmt.Printf("Received 1st message: %s\n", <-messages)
	fmt.Printf("Received 2nd message: %s\n", <-messages)
}