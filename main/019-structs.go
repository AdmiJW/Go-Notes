package main

import "fmt"

// Go's structs are typed collections of fields.
// They're useful for grouping data together to form records.

// For example, this `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

// This `newPerson` function constructs a new `person` struct with the given name.
// You can safely return a pointer to local variable as a local variable will survive the scope of the function.
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func Structs() {
	// This syntax creates a new struct.
	p := person{"Bob", 20}
	fmt.Println(p)

	// You can name the fields when initializing a struct.
	p = person{name: "Alice", age: 30}
	fmt.Println(p)

	// Omitted fields will be zero-valued.
	p = person{name: "Fred"}
	fmt.Println(p)

	// An `&` prefix yields a pointer to the struct.
	pointer := &person{name: "Ann", age: 40}
	fmt.Println(pointer)

	// It's idiomatic to encapsulate new struct creation in constructor functions.
	fmt.Println(newPerson("Jon"))

	// To access the field of a struct, use the dot syntax.
	p = person{name: "Sean", age: 50}
	fmt.Println(p.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	pointer = &p
	fmt.Println(pointer.age)

	// Structs are mutable.
	pointer.age = 51
	fmt.Println(pointer.age)

	// If a struct type is only used for a single value, we don't have to give it a name.
	// The value can have an anonymous struct type.
	// Such technique is commonly used for table-driven tests.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
