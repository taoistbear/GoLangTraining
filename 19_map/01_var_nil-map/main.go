package main

import "fmt"

func main() {
	var myGreeting map[string]string
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
}

// add these lines:
/*
  myGreeting["Josh"] = "Good Morning."
  myGreeting["Jenny"] = "Bonjour."
*/
// and you will get this:
// panic: assignement to entry in nil map
