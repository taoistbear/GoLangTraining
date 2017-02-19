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
	fmt.Print("Please enter a number: ")
	fmt.Scan(&small)
	fmt.Print("Please enter a number larger than ")
	fmt.Print(small)
	fmt.Print(": ")
	fmt.Scan(&large)
	for small >= large {
		fmt.Print("You did not enter a number larger than ")
		fmt.Print(small)
		fmt.Print(". Please try again: ")
		fmt.Scan(&large)
		if small < large {
			break
		}
	}
	remain := large % small
	fmt.Print(remain)
	fmt.Print(" is the rmainder of ")
	fmt.Print(large)
	fmt.Print(" divided by ")
	fmt.Print(small)
	fmt.Print(".\n")
}
