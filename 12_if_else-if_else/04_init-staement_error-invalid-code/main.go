package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	fmt.Println(food)
}

// food is in the scope of the if statement and causes and error.
