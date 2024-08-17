package go_by_example

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures.
// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// Here's the concrete implementation of the interface for `rect` and `circle` types.
type rect2 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect2) area() float64 {
	return r.width * r.height
}

func (r *rect2) perim() float64 {
	return (2 * r.width) + (2 * r.height)
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interfaces() {
	r := rect2{width: 3, height: 4}
	c := circle{radius: 5}

	measure(&r)
	measure(&c)
}
