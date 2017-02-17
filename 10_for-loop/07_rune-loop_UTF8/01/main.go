package main

import "fmt"

func main() {
	for i := 250; i <= 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
}

/*
NOTE:
Some operating systems (Windows) might not print characters wehre i < 256

If you have this issue, you can use this code:

fmt.Println(i, " - ", string(i), " - ", []int32(sring(i)))

UTF-8 works with 1-4 bytes.

A byte is 8 bits.

[]byte deals with bytes, taht is, only 1 byte (8 bits) at a time.

[]int32 allows us to store the value of 4 bytes, that is 5 bytes * 8 bits per byte = 32 bits.
*/
