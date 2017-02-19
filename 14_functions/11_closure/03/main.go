package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by mulitple functions
wihtout closure, for two or more funcs to hav e access to the same vaialbe,
that variable would need to be package scope

anonymous function
a funciton without a name

func expression
assign a func to a variable
*/
