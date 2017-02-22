package main

import "fmt"

func main() {
	greeting := []string{
		"Good Morning!",
		"Bonjoiur!",
		"Buenos Dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selmat Pagi!",
		"Gutten Morgen!",
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2]) // start at index 1 upto but not including 2
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2]) // start at index 0 upto but not including 2
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:]) // start at index 5 and all the way through the end
	fmt.Print("[:] ")
	fmt.Println(greeting[:]) // start at index 0 and all the way through the end
}
