/*
6. Find a problem at projecteuler.net then create the solution. Add a comment
beneath your solution that includes a description of the problem. Upload your
solution to github. Tweet me (Todd) your solution.
*/

package main

import "fmt"

func main() {
	fmt.Println(fibSumEven())
}

func fibSumEven() int {
	var sum int
	var prev = 1
	var curr = 1
	for i := 1; i <= 32; i++ {
		var next int = (prev + curr)
		if next%2 == 0 {
			sum += next
		}
		prev = curr
		curr = next
	}
	return sum
}

/*
Even Fibonacci numbers
Problem 2

Find the sum of all even Fibonnaci numbers with a value of less than 4 million
*/

// It was a few trial by error to find the limit of i to the limit value
