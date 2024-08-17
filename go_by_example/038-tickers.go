package go_by_example

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future.
// Tickers are for when you want to do something repeatedly at regular intervals.

func Tickers() {
	// Tickers use a similar mechanism to timers: a channel that sends values.
	// Here we'll use the `select` builtin on the channel to await the values as they arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Tickers can be stopped like timers. Once a ticker is stopped,
	// it won't receive any more values on its channel. We'll stop ours after 1600ms, after 3 ticks.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}