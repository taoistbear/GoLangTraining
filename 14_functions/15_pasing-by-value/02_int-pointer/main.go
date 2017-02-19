package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age) // 0x.......

	changeMe(&age)

	fmt.Println(&age) // 0x.......
	fmt.Println(age)  // 24
}

func changeMe(z *int) {
	fmt.Println(z)  // 0x......
	fmt.Println(*z) // 44
	*z = 24
	fmt.Println(z)  // 0x......
	fmt.Println(*z) // 24
}
