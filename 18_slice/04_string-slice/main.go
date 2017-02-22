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
		"Gutten Morgen!", // This is not like JS, end objects with ','
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}
}
