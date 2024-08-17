package go_by_example

import (
	"fmt"
)

func Variables() {
	// Variables are explicitly declared and used by the compiler to e.g. check type-correctness
	// Use `var` to declare a variable
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued
	// For example, the zero value for an int is 0
	// The zero value for a string is an empty string
	// The zero value for a boolean is false
	var e int
	fmt.Println(e)
	var f string
	fmt.Println(f)
	var g bool
	fmt.Println(g)

	// The := syntax is shorthand for declaring and initializing a variable
	// This is equivalent to `var h string = "short"`
	h := "short"
	fmt.Println(h)
}