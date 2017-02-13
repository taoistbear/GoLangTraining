package main

import "fmt"

const (
	_ = iota      // 0 >> this throws away the zero
	b = iota * 10 // 1 * 10
	c = iota * 10 // 2 * 10
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
}

// iota is applicate go MATH
