package main

import "fmt"

func main() {
	greet("Jane", "Doe")
}

func greet(fname, lname string) {
	fmt.Println("Hello", fname, lname)
}

// with both parameters of type string you can declare type once as shown above
// that is pretty cool... :)
