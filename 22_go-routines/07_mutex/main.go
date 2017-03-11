package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

// mutex is a mutual exclusion object, it is created so that multiple program
// threads can take turns sharing the same reasource
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		// Lock() clamps down the variable and prevents other threads from accessing it
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		// Unlock() opens up the variable in order allow other threads to access it
		mutex.Unlock()
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go
