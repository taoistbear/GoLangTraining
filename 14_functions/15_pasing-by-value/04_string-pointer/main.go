package main

import "fmt"

func main() {
	name := "Josh"
	fmt.Println(&name) // 0x.....

	changeMe(&name)

	fmt.Println(&name) // 0x.....
	fmt.Println(name)  // Rocky
}

func changeMe(z *string) {
	fmt.Println(z)  // 0x....
	fmt.Println(*z) // Josh
	*z = "Rocy"
	fmt.Println(z)  // 0x.....
	fmt.Println(*z) // Rocky
}
