package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "Josh" // => No Error
	// student = append(student, "Josh") // => No Error
	fmt.Println(student)
	fmt.Println(students)

}

/*
When using make() everything is set for the []slice. append() is only needed
after breaching the length specified in the make() statement.
*/
