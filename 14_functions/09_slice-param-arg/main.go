package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data)
	fmt.Println(n)
}

func average(sf []float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

/*
To compare this version to the previous example:

The func 'average' is looking for a data structure, in this case a slice of
float64. Thus the parameter is set. The argument passed is the same type of
data structure, a slice of float64 ([]float64).

I will need to pay attention to these details as they vary quite differently
from what I was used to doing with JavaScript.
*/
