package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\t", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\t", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\t", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\t", TB)
}
