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
	delete(myGreeting, 7)
	fmt.Println(myGreeting)
}
