package go_by_example

import "fmt"

// Go has added support for generics, also known as type parameters.

// As an example, `mapKeys` takes a map of any type, and return a slice of its keys.
// This function has two type parameters - `K` and `V`.
// `K` has the `comparable` constraint, which means that the type of `K` must support the `==` and `!=` operators
// which is required for the map key type in Go.
// `V` has the `any` constraint, meaning that it's not restricted in any way. (`any` is an alias for `interface{}`).
func mapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))	// 0 length, but capacity of len(m)
	for k := range m {
		r = append(r, k)
	}
	return r
}

// As an example of a generic type, `List` is a singly-linked list with values of any type.
type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val T
}

// We can define methods on generic types just like on regular types, 
// but we have to keep the type parameters in place.
func (list *List[T]) Push(val T) {
	if list.tail == nil {
		list.head = &element[T]{val: val}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: val}
		list.tail = list.tail.next
	}
}
func (list *List[T]) GetAll() []T {
	var elements []T
	for e := list.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return elements
}

func Generics() {
	m := map[int]string { 1: "2", 2: "4", 4: "8" }

	// When invoking generic functions, we can often rely on type inference to determine the type parameters.
	// Note how we don't have to specify the type parameters `K` and `V` explicitly.
	fmt.Println("keys:", mapKeys(m))

	// Though, we can specify the type parameters explicitly if we want to.
	fmt.Println("keys:", mapKeys[int, string](m))

	list := List[int]{}
	list.Push(10)
	list.Push(13)
	list.Push(23)
	fmt.Println("list:", list.GetAll())

	// Note: the order of iteration over map keys is not guaranteed to be the same between calls.
}