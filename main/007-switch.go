package main

import (
	"fmt"
	"time"
)

func Switch() {
	// `switch` statements express conditionals across many branches
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use commas to separate multiple expressions in the same `case` statement
	// Use the `default` case to match when no other `case` is satisfied
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// `switch` without an expression is an alternate way to express `if/else` logic
	// `case` expressions can also be non-constants
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type `switch` compares types instead of values.
	// This can be used to discover the type of an interface value

	findType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It's an boolean")
		case int:
			fmt.Println("It's a integer")
		default:
			fmt.Printf("Unknown type %T\n", t)
		}
	}

	findType(true)
	findType(1)
	findType("hello")
}