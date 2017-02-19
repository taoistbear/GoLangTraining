/*
4. Create a program that prints to the terminal asking for a user to enter a
small number and a larger number. Print the remaider of the larger number
divided by the smaller number.
*/

package main

import "fmt"

func main() {
	var small int
	var large int
	fmt.Println("Please enter a number...")
	fmt.Scan(&small)
	fmt.Println("Please enter a number larger than", small)
	fmt.Scan(&large)
	for small >= large {
		fmt.Println("You did not enter a number larger than", small)
		fmt.Println("Please enter a number LARGER than", small)
		fmt.Scan(&large)
		if small < large {
			break
		}
	}
	remain := large % small
	fmt.Println(remain, "is th remainder of", large, "divided by", small)
}
