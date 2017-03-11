package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
		}
		done <- true
	}()

	/*
		By using this here the semaphore will not allow the program to flow to the
		range loop. It will never complete and there is not concurrent running.
		This creates the deadlock! The goroutine prevents this and sets it aside and
		makes it like an event listening in JavaScript!!!
	*/
	// we block here until done <- true
	<-done
	<-done
	close(c)

	// to unblock above
	// we need to take values off of chan c here
	/// but we never get here, because we're blocked above
	for n := range c {
		fmt.Println(n)
	}
}

/*
fatal error: all goroutines are asleep - deadlock!
*/
