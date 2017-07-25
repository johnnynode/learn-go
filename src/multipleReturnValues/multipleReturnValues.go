package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a,b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

/*


$ methods run multiple-return-values.methods
3
7
7

*/