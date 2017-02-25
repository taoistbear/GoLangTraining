package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good Morning!",
		1: "Bonjour!",
		2: "Buenos Dias!",
		3: "Bonjiorno!",
	}

	fmt.Println(myGreeting)

	if val, exists := myGreeting[7]; exists {
		delete(myGreeting, 7)
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	} else {
		fmt.Println("That value doesn't exists.")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}

	fmt.Println(myGreeting)
}
