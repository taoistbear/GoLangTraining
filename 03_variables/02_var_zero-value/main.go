package main

import "fmt"

// demonstrating variable declaration w/type assignment git and printing their zero value if not initialized
func main() {

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println()
}
