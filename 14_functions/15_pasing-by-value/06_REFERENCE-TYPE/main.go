package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Josh]
}

func changeMe(z []string) {
	z[0] = "Josh"
	fmt.Println(z) // [Josh]
}
