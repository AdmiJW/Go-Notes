package main

import (
	"fmt"
	"slices"
)

func Slices() {
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements)
	// An uninitialized slice equals to `nil` and has a length and capacity of 0
	var s []string
	fmt.Println("uninitialized slice:", s, s == nil, len(s) == 0)

	// To create an empty slice with non-zero length, use the built-in `make` function
	// By default, a new slice's capacity is equal to its length
	// However, you can specify the capacity by providing a third argument to `make`
	s = make([]string, 3, 5)
	fmt.Println("empty slice:", s, "length:", len(s), "capacity:", cap(s))

	// You can set and get values in the same way as arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("length:", len(s))

	// Slices support several more operations than arrays
	// The built-in `append` function adds elements to a slice.
	// Since `append` can return a new slice, you must always store the result of `append`
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	// Slices can also be copied using the `copy` function
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// Slices can be sliced, with the syntax `slice[low:high]`
	// This will return a slice containing the elements `slice[low]`, `slice[low+1]`, ..., `slice[high-1]`
	l := s[2:5]
	fmt.Println("slice:", l)

	// This slices up to (but excluding) `s[5]`
	l = s[:5]
	fmt.Println("up to:", l)

	// And this slices up from (and including) `s[2]`
	l = s[2:]
	fmt.Println("from:", l)

	// A slice can be declared and initialized in one line
	t := []string {"g", "h", "i"}
	fmt.Println("declare:", t)

	// The `slices` package provides a number of useful functions for working with slices
	t2 := []string {"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays
	twoDimensional := make([][]int, 3)
	for i := range len(twoDimensional) {
		innerLength := i + 1
		twoDimensional[i] = make([]int, innerLength)

		for j := range len(twoDimensional[i]) {
			twoDimensional[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoDimensional)
}