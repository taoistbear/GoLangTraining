package main

import "fmt"

func main() {
	var myGreeting map[string]string
	// var myGreeting = map[string]string{}
	// using a composit literal works to initialize the map beyond nil
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

/*
	There is no append method to get out of this nil hole.
*/
