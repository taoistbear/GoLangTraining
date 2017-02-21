/*
2. Modify the previous program to use a func expression.
*/

package main

import "fmt"

func main() {
	var half = func(x int) (int, bool) {
		var half int = x / 2
		if x%2 == 0 {
			return half, true
		} else {
			return half, false
		}
	}
	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(half(16))
	fmt.Println(half(21))
	fmt.Println(half(42))

}
