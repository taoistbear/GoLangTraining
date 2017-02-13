package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	// no access to the scope storing x
	// this does not compile
	fmt.Println(x)
}

// x would need to be declared outside of main at the package level
// in order to get this to compile.
