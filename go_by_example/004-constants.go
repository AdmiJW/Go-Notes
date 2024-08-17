package go_by_example

import (
	"fmt"
	"math"
)


func Constants() {
	// `const` declares a constant value
	const s string = "constant"
	
	fmt.Println(s)

	// Anywhere a `var` statement can appear, a `const` statement can appear
	const n = 500000000
	fmt.Println(n)

	// Constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it's given one, such as by an explicit conversion
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call
	// For example, here `math.Sin` expects a `float64`, and `n` has been given a type by the context
	fmt.Println(math.Sin(n))
}