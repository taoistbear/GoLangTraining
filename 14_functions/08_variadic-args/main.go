package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

/*
To understand the diference between variadic functions with parameters and
agruments look at the place of the elipses. '...' is placed before the parameter
of TYPE when declaring a function. '...' is placed after the the variable when
passing an agrument (if the variable is a singular item referencing a data
structure such as the slice above).
*/
