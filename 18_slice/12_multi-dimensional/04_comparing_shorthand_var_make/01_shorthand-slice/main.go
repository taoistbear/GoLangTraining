package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{}
	// student[0] = "Josh"  // => get same out of range error
	student = append(student, "Josh")
	fmt.Println(student)
	fmt.Println(students)

}

/*
When using shorthand or var, i.e. NOT using make(), there is no reference, length
and/or capacity. That is why append() is needed to add to the []slice.
*/
