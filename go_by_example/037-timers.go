package go_by_example

import (
	"fmt"
	"time"
)

// We often want to execute Go code at some point in the future, or repeatedly at regular intervals.
// Go's built-in `timer`s and `ticker`s make both of these tasks easy.

func Timers() {
	// Timers represent a single event in the future.
	// You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.
	// This timer will wait 2s.
	timer1 := time.NewTimer(2 * time.Second)

	// The `<-timer1.C` blocks on the timer's channel `C` until it sends a value indicating that the timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// You might wonder why we didn't just use `time.Sleep`.
	// One reason is that timers may be useful if you want to cancel the timer before it fires.
	// Here's an example of that.
	timer2 := time.NewTimer(time.Second)
	go func() {
		// This will block until the timer fires. Since the timer is stopped, this will keep blocking.
		<-timer2.C		
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Give the timer2 enough time to fire if it were not stopped.
	time.Sleep(2 * time.Second)
}