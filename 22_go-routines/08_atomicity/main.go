package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		// atomic.AddInt64 locks the variable down same as mutex.Lock()/Unlock()
		// sequence may be a little wonky with atomic methods so if exact sequence
		// is needed use mutex
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Coutner:", atomic.LoadInt64(&counter)) // access wihtout race
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go
