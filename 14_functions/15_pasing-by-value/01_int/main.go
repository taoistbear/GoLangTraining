package main

import "fmt"

func main() {
	age := 44
	changeMe(age)
	fmt.Println(age) // 44
}

func changeMe(z int) {
	fmt.Println(z)
	z = 24
	fmt.Println(z)
}

// when changeMe is called on line 8
// the vallue 44 is being passed as an agrument
