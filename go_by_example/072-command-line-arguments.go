package go_by_example

import (
	"fmt"
	"os"
)

// Command-line arguments are a common way to parameterize execution of programs.
// For example, `go run hello.go` uses `run` and `hello.go` arguments to the `go` program.

func CommandLineArguments() {
	// `os.Args` provides access to raw command-line arguments.
	// Note that the first value in this slice is the path to the program,
	// and `os.Args[1:]` holds the arguments to the program.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// You can get individual args with normal indexing.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	// Run this program with a few arguments, such as:
	//   go run . a b c d
}