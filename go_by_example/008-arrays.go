package go_by_example

import (
	"fmt"
)

func Arrays() {
	// In Go, an array is a numbered sequence of elements of a specific length
	// In typical usage, slices are more common than arrays.
	// That said, arrays are useful in some specific cases.

	// To create an array, use the following syntax
	// By default, an array is zero-valued, which for `int`s means `0`s
	var a [5]int
	fmt.Println("emp:", a)

	// Set a value at an index using the `array[index] = value` syntax
	a[4] = 100
	fmt.Println("set:", a)

	// Get a value at an index using the `array[index]` syntax
	fmt.Println("get:", a[4])

	// The built-in `len` function returns the length of an array
	fmt.Println("len:", len(a))

	// Use this syntax to declare and initialize an array in one line
	b := [5]int {1, 2, 3, 4, 5}
	fmt.Println("declare:", b)

	// Using such syntax, you can have the compiler count the length of the array for you using `...`
	c := [...]int {1, 2, 3, 4, 5}
	fmt.Println("auto count:", c)

	// If you specify the index with `:`, the elements in between will be zero-valued
	d := [...]int {100, 3:400, 500}
	fmt.Println("specify index:", d)

	// Arrays are one-dimensional, but you can compose types to build multi-dimensional data structures
	var twoDimensional [2][3]int
	for i := range len(twoDimensional) {
		for j := range len(twoDimensional[i]) {
			twoDimensional[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoDimensional)

	// Multi-dimensional arrays can also be created and initialized at once
	twoDimensional2 := [2][3]int {
		{0, 1, 2},
		{3, 4, 5},
	}
	fmt.Println("2d2:", twoDimensional2)
}