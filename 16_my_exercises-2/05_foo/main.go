/*
5. Write a function, foo, which can be called in all of these ways:
  func main() {
    foo(1, 2)
    foo(1, 2, 3)
    aSlice := []int{1, 2, 3, 4}
    foo(aSlice...)
    foo()
  }
*/

package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(num ...int) {
	for _, n := range num {
		fmt.Print(n)
	}
	fmt.Println("...")
	if num == nil {
		fmt.Println("No int was passed as an argument.")
	}
}
