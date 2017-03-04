package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

// another shape
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

// which implements the shape interface
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
}

/*
The polymorphism of using interface unifies similar (but different) methods on similar
struct(s) etc. It brings things together through an ... interface where methods can
interact differently depending on what is calling the methods. It is a bucket of sorts
that can do the hard work for you via the compiler to really narrow down methodology
to truly be unique and interactive based on how it is being used.
*/
