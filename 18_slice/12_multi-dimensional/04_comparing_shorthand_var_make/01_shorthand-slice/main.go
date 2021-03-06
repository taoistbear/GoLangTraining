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
When using shorthand there is an undelying array, yet there is no length set,
thus using append() is neccessary to add items to the []slice
*/
