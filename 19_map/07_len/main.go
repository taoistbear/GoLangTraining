package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good Morning!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy!"

	fmt.Println(len(myGreeting))
}
