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

/*
Using the pointer you are able to change a variable that is in a DIFFERENT SCOPE.
Important concetp to learn.

When declaring TYPE in PARAMETER using '*' indicates that it is pointer of reference.
---- func changeMe(z *int) {}

When declaring a VARIABLE using '*' dereferences that memory address allowing
access to the the actual value (if not you would be changing the memory address)
---- *z = 24

When passing AGRUMENT using '&' indicates this is the memory address and creates
the pointer to be used by other funcs outside of the nomral scope
---- changeMe(&age)
*/
