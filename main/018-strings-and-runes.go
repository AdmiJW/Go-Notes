package main

import (
	"fmt"
	"unicode/utf8"
)

func StringsAndRunes() {
	// A Go string is a read-only slice of bytes.
	// The language and the standard library treat strings specially - as containers of text encoded in UTF-8.
	// In other languages, strings are made of characters.
	// In Go, the concept of a character is called a `rune` - an integer value identifying a Unicode code point.

	// `s` is a string assigned a literal value representing the word "hello" in Thai. 
	// Go string literals are UTF-8 encoded text.
	const s = "สวัสดี"

	// Since strings are equivalent to `[]byte`, `len(s)` returns the number of bytes in the string.
	fmt.Println("length of s:", len(s))

	// Indexing into a string produces the raw byte values at each index. 
	// This loop generates the hex values of all the bytes that constitute the code points in `s`.
	for i := range len(s) {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// To count how many runes are in a string, we can use the `utf8` package.
	// Note that the run-time of `RuneCountInString` is O(n), because it has to decode each UTF-8 rune sequentially.
	// Some Thai characters are represented by UTF-8 code points that can span multiple bytes.
	fmt.Println("rune count:", utf8.RuneCountInString(s))

	// A `range` loop handles strings specially and decodes each rune along with its offset in the string.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, idx)
	}

	// The same iteration can be done using the `utf8.DecodeRuneInString` function.
	for i,w := 0,0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width

		// This demonstrates passing a `rune` value to a function.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found t")
	} else if r == 'ส' {
        fmt.Println("found so sua")
    }
}