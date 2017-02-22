package main

import (
	"fmt"
)

func main() {
	var student []string
	var students [][]string
	// student[0] = "Josh"  // => get same out of range error
	student = append(student, "Josh")
	fmt.Println(student)
	fmt.Println(students)

}
