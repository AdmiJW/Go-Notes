package go_by_example

import (
	"fmt"
)

// Variadic functions can be called with any number of trailing arguments.
// For example, `fmt.Println` is a common variadic function.

// Here's a function that will take an arbitrary number of `int`s as arguments.
// Within the function, `nums` will be available as a `[]int`. We can call `len(nums)`, iterate over it with `range`, etc.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func VariadicFunctions() {
	// Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice, apply them to a variadic function using `func(slice...)` like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}