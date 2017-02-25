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

	// delete(myGreeting, 2)

	if val, exists := myGreeting[2]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	} else {
		fmt.Println("That value doesn't exists.")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}

	fmt.Println(myGreeting)
}
