package main

import "fmt"

// using variables with shorthand method and printing their type
func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	// %T will format for type
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
