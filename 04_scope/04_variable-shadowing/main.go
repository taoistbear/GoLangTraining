package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // max is now the result, not the function
}

// don't do this; bad coding practice to shadow variables in this way. code is
// less readable and variables, functions can easily be overwritten.

// notice there are no issues when this is compliled.
