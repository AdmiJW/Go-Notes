package go_by_example

import (
	"fmt"
	"net/http"
	"time"
)

// In the previous example, we looked at setting up a simple HTTP server.
// HTTP servers are useful for demonstrating the usage of `context.Context` for controlling cancellation.
// A `Context` carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines.

func hello2(w http.ResponseWriter, req *http.Request) {
	// A `context.Context` is created for each request by the `net/http` machinery,
	// and is available with the `Context()` method.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Wait for a few seconds before sending a reply to the client.
	// This could simulate some work the server is doing.
	// While working, keep an eye on the context's `Done()` channel for a signal that we should cancel 
	// the work and return as soon as possible.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// The context's `Err()` method returns an error that explains why the `Done()` channel was closed.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func Context() {
	// As before, we register our handler on the "/hello" route, and start serving.
	http.HandleFunc("/hello", hello2)
	http.ListenAndServe(":8090", nil)

	// To test the server, we can issue a request using `curl` from another terminal, or use the browser.
	// Cancel the request (Ctrl+C) to see the server cancel the request.
}