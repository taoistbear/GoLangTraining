package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good Morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Buenos Dias"
	greeting[3] = "suprabadham"

	fmt.Println(greeting[2])
}

/*
➜  07_append-invalid git:(master) ✗ go run main.go
panic: runtime error: index out of range

goroutine 1 [running]:
panic(0x79c40, 0xc42000a0e0)
	/usr/local/go/src/runtime/panic.go:500 +0x1a1
main.main()
	/Users/NEXSW/Desktop/GoCode/src/github.com/taoistbear/GoLangTraining/18_slice/07_append-invalid/main.go:13 +0x7e
exit status 2
*/
