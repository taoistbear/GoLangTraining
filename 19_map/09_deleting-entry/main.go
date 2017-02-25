package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"zero":  "Good Morning!",
		"one":   "Bonjour!",
		"two":   "Buenos Dias!",
		"three": "Bonjiorno!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "two")
	fmt.Println(myGreeting)
}
