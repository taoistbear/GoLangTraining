package main

import "fmt"

// declaring multiple variables at once. Not reccomended as code is not as succinct and readable.
func main() {
	var message string
	var a, b, c int
	a = 1

	message = "Hello World"

	fmt.Println(message, a, b, c)
}
