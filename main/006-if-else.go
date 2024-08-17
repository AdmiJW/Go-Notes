package main

import (
	"fmt"
)

func IfElse() {
	// Branching with `if` and `else` in Go is straight-forward
	if isEven(7) {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// You can have an `if` statement without an `else`
	if 8 % 4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Logical operators `&&`, `||` and `!` are useful for expressing multiple conditions
	if isEven(8) || isEven(7) {
		fmt.Println("Either 8 or 7 is even")
	}

	// A statement can precede conditionals 
	// Variables declared in this statement are available in all branches (current, and subsequent `else` branches)
	if n1, n2 := 2, 2; n1 < n2 {
		fmt.Println("1 is less than 2")
	} else if n1 > n2 {
		fmt.Println("1 is not less than 2")
	} else {
		fmt.Println("1 is equal to 2")
	}

	// Note that you don't need parentheses around conditions in Go
	// However, the braces are always required

	// There is no ternary if in Go, so you'll need to use a full `if` statement even for basic conditions
}

func isEven(n int) bool {
	return n % 2 == 0
}