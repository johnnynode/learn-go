package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

/*

$ go run environment-variables.go
FOO: 1
BAR:

PATH
SHELL
...

$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...

*/