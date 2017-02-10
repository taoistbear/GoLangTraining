package main

import "fmt"

func main() {
	// loops to print the first 2000 UTF-8 codes in decimal, binary and hexademical format
	for i := 0; i < 2000; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
