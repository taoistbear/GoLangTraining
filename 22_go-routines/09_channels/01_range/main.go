package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// using close() it will close the channel and prevent any further items
		// to be added to it; if not the chanel and routine will be asleep and create
		// a deadlock
		close(c)
	}()

	/*
		By using range here it creates automation with the channel to pull and push
		values off of the channel as soon as they are added to essentially 'block' then
		unblock the flow of data seamlessly until the channel is closed.
	*/
	for n := range c {
		fmt.Println(n)
	}
}

/*
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()

by using 'close(c)' this runs. close removes the deadlock.
*/
