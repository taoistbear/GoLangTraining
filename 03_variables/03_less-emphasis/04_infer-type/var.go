package main

import "fmt"

// declaring and initializing many variables at once withouit assigning type... Go will infer the type of variable as best as it can but should try to assign type upon declaration.
func main() {

	var message = "Hello World!"
	var a, b, c = 1, 2, 3

	fmt.Println(message, a, b, c)
}
