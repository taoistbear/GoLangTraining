package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hello World!")
	}

	greeting()
	fmt.Printf("%T\n", greeting)
}
