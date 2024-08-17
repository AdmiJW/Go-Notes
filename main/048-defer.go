package main

import (
	"fmt"
	"os"
)

// Defer is used to ensure that a function call is performed later in a program's execution,
// usually for purposes of cleanup.
// `defer` is often used where e.g. `ensure` and `finally` would be used in other languages.

func Defer() {
	createFile := func(p string) *os.File {
		fmt.Println("creating file")
		f, err := os.Create(p)
		if err != nil {
			panic(err)
		}
		return f
	}
	writeFile := func(f *os.File) {
		fmt.Println("writing to file")
		fmt.Fprintln(f, "data")
	}
	// It's important to check for errors when closing a file, even in a deferred function.
	closeFile := func(f *os.File) {
		fmt.Println("closing file")
		err := f.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}

	// Suppose we wanted to create a file, write to it, and then close when we're done.
	// Here's how we could do that with `defer`.
	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)

}