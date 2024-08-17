package main

import "fmt"

// Go supports embedding of structs and interfaces to express a more seamless composition of types.

type base struct {
	num int
}

func (b *base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// Here, a `container` embeds a `base`. An embedding looks like a field without a name.
type container struct {
	base
	str string
}

func StructEmbedding() {
	// When creating structs with literals, we have to initialize the embedding explicity.
	co := container {
		base: base { num: 1 },
		str: "str",
	}

	// We can access the base's fields directly on `co`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Alternatively, we can spell out the full path using the embedded type's name.
	fmt.Printf("co={num: %v, str: %v}\n", co.base.num, co.str)

	// Since `container` embeds `base`, the methods of `base` also become methods of `container`.
	// Here, we invoke a method that was embedded from `base` directly on `co`.
	fmt.Println("describe:", co.describe())

	// Embedding structs with methods may be used to bestow interface implementations onto other structs.
	// Here we see that a `container` now implements the `describer` interface because it embeds `base`.
	type describer interface {
		describe() string
	}
	var d describer = &co
	fmt.Println("describer:", d.describe())
}