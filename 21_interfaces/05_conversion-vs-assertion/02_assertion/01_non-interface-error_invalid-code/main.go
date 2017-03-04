package main

import "fmt"

func main() {
	name := "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}

/*
invalid type assertion: name.(string) (non-interface type string on left)
*/
