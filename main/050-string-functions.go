package main

import (
	"fmt"
	s "strings"
)

// The standard library's `strings` package provides many useful string-related functions.

func StringFunctions() {
	// We alias `fmt.Println` to a shorter name, `p`, for easier typing.
	p := fmt.Println 

	// Here's a sample of the functions available in `strings`.
	// Since these are functions from the package, not methods on the string object itself, 
	// we need to pass the string in question as the first argument to the function.
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}