package go_by_example

import (
	"fmt"
)

func For() {
	// `for` is Go's only looping construct

	// The most basic type, with a single condition (like a `while` loop)
	i := 1
	for i <= 3 {
		fmt.Println("basic", i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop
	for i := 0; i < 3; i++ {
		fmt.Println("classic", i)
	}

	// Another way of accomplishing the basic `do this N times` loop
	// Use `range` over an integer
	for i := range 3 {
		fmt.Println("range", i)
	}

	// `for` without a condition will loop repeatedly until you `break` out of the loop or `return` from the enclosing function
	for {
		fmt.Println("break")
		break
	}
	ForReturn()

	// You can also `continue` to the next iteration of the loop
	for i := range 6 {
		if i%2 == 0 {
			continue
		}
		fmt.Println("continue", i)
	}
}

func ForReturn() {
	for {
		fmt.Println("return")
		return
	}
}