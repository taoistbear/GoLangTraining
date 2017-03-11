package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() chan int {
	in := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()
	return in
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
