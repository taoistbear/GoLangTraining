package main

import "fmt"

func main() {
	// loops to print the first 200 numbers in decimal, binary and hexademinal format
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
