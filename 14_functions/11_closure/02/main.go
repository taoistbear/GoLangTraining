package main

import "fmt"

var x int

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by mulitple functions
wihtout closure, for two or more funcs to hav e access to the same vaialbe,
that variable would need to be package scope
*/
