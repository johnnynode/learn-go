package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}


}

/*

$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/../panic.go:6 +0x64
exit status 2

*/