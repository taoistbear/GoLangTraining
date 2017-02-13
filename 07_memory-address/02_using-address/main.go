package main

import "fmt"

const MetersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * MetersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
