package main

import "fmt"

// using basic variables with shorthand method and printing their value
func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	// %v will format for standard outpout for variable by type
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
