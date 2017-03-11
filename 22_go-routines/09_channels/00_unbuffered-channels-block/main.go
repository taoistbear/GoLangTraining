package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		This is an unbuffered channel. A buffered channel is for more advance use and
		would look like 'c := make(chan int, 347)' where there is a numerical value
		declared to describe the size of the buffer.
	*/
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	// time.Sleep() presented here merely waits and allows time for the goroutines to finish
	// before exiting the main func
	time.Sleep(time.Second)
}

/*
	The '<-' syntax is describing where the value is being pushed. As with 'c <- i'
	the arrow is pointing towards the declared channel 'c'. Thus pushing the value towards
	it. as with '<-c' is is pulling the value out of the channel 'c'.

	With an unbuffered channel the program will stop until the value is used, or 'pulled'
	off of the channel. Such as:

	Push to channel:  c <- i
	Pull off channel:   <-c

	This is synchronized much like a 2 valve piston! Intake and Exhaust!
*/
