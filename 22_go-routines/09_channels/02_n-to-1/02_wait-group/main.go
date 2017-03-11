package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	// by moving the wait group counter out of the goroutines it eliminates the rc
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		// notice the close() is in its own goroutine down here waiting for the
		// WaitGroup to finish before closing the channel
		// without the close this creates the deadlock and sleeping channel
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
