package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s) // constant
	const n = 50000000
	const d = 3e20 / n
	fmt.Println(d) // 6e+12
	fmt.Println(int64(d)) // 6000000000000
	fmt.Println(math.Sin(n)) // 0.8256467432733234
}
