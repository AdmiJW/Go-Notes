package go_by_example

import (
	"fmt"
)

// Go supports recursive functions. Here's a classic factorial example.

// This `factorial` function calls itself until it reaches the base case of `fact(0)`.
func fact(n int) int {
	if n == 0 { return 1 }
	return n * fact(n - 1)
}

func Recursion() {
	fmt.Println(fact(7))

	// Closures can also be recursive, but this requires the closure to be assigned to a variable first.
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 { return n }
		return fib(n - 1) + fib(n - 2)
	}
	fmt.Println(fib(7))
}