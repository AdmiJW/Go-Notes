package main

import (
	"embed"
)

// `go:embed` is a compiler directive that allows programs to include arbitrary files
// and folders in the Go binary at build time.
// Import the `embed` package. If you don't use any exported identifiers from this package,
// you can do a blank import with `_ "embed"`.

// `embed` directives accepts paths relative to the directory containing the Go source file.
// This directive embeds the contents of the file into the `string` variable immediately following it.

//go:embed tmp/single_file.txt
var fileString string

// Or embed the contents of the file into a `[]byte`.
//go:embed tmp/single_file.txt
var fileByte []byte

// We can also embed multiple files or even folders with wildcards.
// This uses a variable of the embed.FS type, which implements a simple virtual file system.
//go:embed tmp/single_file.txt
//go:embed tmp/*.hash
var folder embed.FS

func EmbedDirective() {
	// Print out the contents of `single_file.txt`
	print(fileString)
	print(string(fileByte))

	// Retrieve some files from the embedded folder.
	content1, _ := folder.ReadFile("tmp/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("tmp/file2.hash")
	print(string(content2))
}	