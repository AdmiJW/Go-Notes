package main

import (
	"errors"
	"fmt"
)

// It's possible to use custom types as `errors` by implementing the `Error()` method on them.
// A custom error type usually has the suffix `Error`.
type argError struct {
	arg int
	message string
}

// Adding the `Error()` method to `argError` makes it implement the `error` interface.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func CustomErrors() {
	f := func(arg int) (int, error) {
		if arg == 42 {
			return -1, &argError{arg, "can't work with it"}
		}
		return arg + 3, nil
	}

	_, err := f(42)
	var ae *argError

	// `errors.As` checks that a given error (or any error in its chain) matches a specific error type.
	// If matches, it converts to a value of that type and assigns it to the variable, then returns `true`.
	// If there's no match, it returns `false`.
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}