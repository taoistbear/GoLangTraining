/*
1. Write a funciton which take an integer. the function will have two returns.
The first return should be the argument divided by 2. The second return should
be a bool that let's us know whether or not the argument was even. For example:
 a. half(1) returns(0, false)
 b. half(2) returns(1, true)
*/

package main

import "fmt"

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(half(16))
	fmt.Println(half(21))
	fmt.Println(half(42))

}

func half(x int) (int, bool) {
	var half int = x / 2
	if x%2 == 0 {
		return half, true
	} else {
		return half, false
	}
}
