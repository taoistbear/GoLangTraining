package main

import "fmt"

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by mulitple functions
wihtout closure, for two or more funcs to hav e access to the same vaialbe,
that variable would need to be package scope
*/
