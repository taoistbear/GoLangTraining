package main

import "fmt"

func main() {
	name := "Josh"
	fmt.Println(name) // Josh

	changeMe(name)

	fmt.Println(name) // Josh
}

func changeMe(z string) {
	fmt.Println(z) // Josh
	z = "Rocy"
	fmt.Println(z) // Rocky
}
