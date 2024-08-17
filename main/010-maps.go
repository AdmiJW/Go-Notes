package main

import (
	"fmt"
	"maps"
)

func Maps() {
	// Maps are Go's built-in associative data type (sometimes called hashes or dicts in other languages)

	// To create an empty map, use the built-in `make` - `make(map[key-type]val-type)`
	m := make(map[string]int)

	// Set key/value pairs using the `map[key] = val` syntax
	m["k1"] = 7
	m["k2"] = 13
	// Printing a map with `fmt.Println` will show all of its key/value pairs
	fmt.Println("map:", m)

	// Get a value for a key with `map[key]`
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key doesn't exist, you'll get the zero-value for the map's value type
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// The built-in `len` function returns the number of key/value pairs when called on a map
	fmt.Println("len:", len(m))

	// The built-in `delete` function removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("delete:", m)

	// To remove all key/value pairs from a map, use the `clear` built-in function
	clear(m)
	fmt.Println("clear:", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map
	// This can be used to disambiguate between missing keys and keys with zero values
	_, isPresent := m["k2"]
	fmt.Println("isPresent:", isPresent)

	// You can also declare and initialize a new map in one line
	n := map[string]int { "foo": 1, "bar": 2 }
	fmt.Println("declare:", n)

	// The `maps` package provides a number of useful functions for working with maps
	n2 := map[string]int { "foo": 1, "bar": 2 }
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}