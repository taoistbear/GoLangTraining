/*
3. Write a function with one variadic parameter that finds the greatest number
in a list of numbers.
*/

package main

import "fmt"

func main() {
	var greatest = func(list ...int) int {
		var great int
		for _, n := range list {
			if n > great {
				great = n
			}
		}
		return great
	}
	fmt.Println(greatest(1, 2, 3, 4, 5))
	fmt.Println(greatest(3, 72, 9, 6, 47))
	fmt.Println(greatest(1000, 4, 82, 53, 98))
	fmt.Println(greatest(7, 16, 32, 75, 5))
}
