package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good Morning!",
		"Jenny": "Bonjour!",
	}

	fmt.Println(myGreeting["Jenny"])
}
