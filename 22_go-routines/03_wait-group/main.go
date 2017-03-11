package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2) // This adds to goroutines to the WaitGroup
	go foo()
	go bar()
	wg.Wait() // When the wait group is 0 this signals for main goroutine to exit
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done() // When this executes it removes 1 from the WaitGroup
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done() // When this executes it removes 1 from the WaitGroup
}
