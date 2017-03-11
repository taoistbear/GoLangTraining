package main

import (
	"fmt"
	"sync"
)

/*
	In this example we are exploring the ability for two goroutines to access the same
	variable. The problem with this is the creation of a race condition by haveing
	two of the goroutines access the WaitGroup seperately.
*/

func main() {
	c := make(chan int)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- 1
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- 1
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
