package go_by_example

import (
	"fmt"
)

func Values() {
	// Strings can be concatenated with the + operator
	fmt.Println("Hello, " + "World!")

	// Integers and floats
	fmt.Println(1 + 1)
	fmt.Println(7.0 / 3.0)

	// Booleans
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}