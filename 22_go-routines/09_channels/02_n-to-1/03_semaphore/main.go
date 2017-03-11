package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	/*
		'done' is what is creating the semaphore. Basically it is a flag or toggle that will
		signal action or inaction based on what is happening. This channel of 'done' will be
		waiting to receive a bool in this instance. It is open and sleeping until it recieves
		something.
	*/
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
		}
		// here we send the bool to channel 'done'
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
		}
		// here we send the bool to channel 'done'
		done <- true
	}()

	go func() {
		/*
			This goroutine when launched will wait for channel 'done' to receive it's
			value and then immediately disposes of that value. That is what this does:

			'<-done'

			This dumps the value.

			This goroutine waits do do this twice. There are two goroutines up there waiting
			to finish their loops to push a bool to this channel. As soon as this happens and
			both dumps have occured it closes the main channel 'c' which will end the range
			satement below.

			I need to research this a little bit more. I am asserting that the dump of
			'<-done' subsequently closes that channel or at least stops the sleep/deadlock
		*/
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
