package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c1 := customer{"Josh", 44}
	fmt.Println(&c1.name) // 0x....

	changeMe(&c1)

	fmt.Println(c1)       // {Rocky 44}
	fmt.Println(&c1.name) // 0x....
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Josh 44}
	fmt.Println(&z.name) // 0x...
	z.name = "Rocky"
	fmt.Println(z)       // &{Rocky 44}
	fmt.Println(&z.name) // 0x....
}
