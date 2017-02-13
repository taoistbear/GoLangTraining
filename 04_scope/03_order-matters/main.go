package main

import "fmt"

func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 42
}

var y = 42

// Order matters when declaring variables. When declaring variables at the
// package level, order does not matter. y is accessible for the whole package
// similar to JavaSciprt where global variables are compiled before running
// a program or at a minimum there is the ability for the program to search for
// the variable.
