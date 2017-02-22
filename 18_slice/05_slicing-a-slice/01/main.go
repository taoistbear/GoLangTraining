package main

import "fmt"

func main() {
	var results []int
	fmt.Println(results)

	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])  // slicing a slice
	fmt.Println(mySlice[2])    // accessing by index - standard
	fmt.Println("myString"[2]) // accessing by index - non-standard
}
