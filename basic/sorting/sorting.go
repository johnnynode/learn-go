package main

import "fmt"
import "sort"

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("String:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}

/*

$ go run sorting.go
String: [a b c]
Ints:  [2 4 7]
Sorted: true

*/