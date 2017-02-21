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

/*
Passing by value is referncing arguments in functions. When passing variables
into functions they do not neccessarily have the same scope as the variable
being passed. With that being said they only receive the value of said variable.

The only difference would be an function expression within the same func
declaration that would have access to said variable in scope or if the variable
was declared at the package level.

This technique should dramatically help keep scope tight.
*/
