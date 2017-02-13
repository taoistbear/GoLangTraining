package main

import "fmt"

const (
	A = iota // 0
	B        // 1
	C        // 2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

// odd behavior where iota is only declare once yet assumed for not declared variables.
