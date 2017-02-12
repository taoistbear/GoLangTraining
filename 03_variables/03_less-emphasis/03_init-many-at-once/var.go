package main

import "fmt"

// declaring, assignig and initializing many variables at once... not reccomended as it makes code less readable and succinct.
func main() {

	var message string = "Hello World!"
	var a, b, c int = 1, 2, 3

	fmt.Println(message, a, b, c)
}
