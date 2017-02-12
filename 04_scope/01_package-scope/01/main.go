package main

import "fmt"

// declaring the variable outside of the func or block it is inside the scope of the package
var x int = 42

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
