package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good Morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Buenos Dias"
	greeting = append(greeting, "Suprabadham!")
	greeting = append(greeting, "Zǎo'ān!")
	greeting = append(greeting, "Ohayou Gozaimasu!")
	greeting = append(greeting, "Gidday!")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
