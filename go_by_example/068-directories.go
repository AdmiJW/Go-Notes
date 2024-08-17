package go_by_example

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Go has several useful functions for working with directories in the file system.

func Directories() {
	check := func(e error) {
		if e != nil { panic(e) }
	}

	// Create a new subdirectory in the current working directory.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// When creating temporary directories, it's a good practice to defer their removal.
	// `os.RemoveAll` will delete a whole directory tree, similar to `rm -rf`.
	defer os.RemoveAll("subdir")

	// Helper function to create a new empty file
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	createEmptyFile("subdir/file1")

	// We can create a hierarchy of directories, including parents with `MkdirAll`.
	// This is similar to the command-line `mkdir -p`.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` lists directory contents, returning a slice of `fs.FileInfo` objects.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` changes the current working directory, similar to `cd`.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Now we'll see the contents of subdir/parent/child when listing the current directory.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `cd` back to the original directory.
	err = os.Chdir("../../..")
	check(err)

	// We can also visit a directory recursively, including all its subdirectories.
	// `WalkDir` accepts a callback function to handle every file or directory visited.
	visit := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(" ", path, d.IsDir())
		return nil
	}

	fmt.Println("Visiting subdir")
	filepath.WalkDir("subdir", visit)
}	