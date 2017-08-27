package main

import (
	"basic/time"
	"fmt"
	"math/rand"
)

var fpln = fmt.Println
var fp = fmt.Print

func main() {

	fp(rand.Intn(100), ",") // 81
	fp(rand.Intn(100))      // 87
	fpln()

	fpln(rand.Float64()) // 0.6645600532184904

	fp((rand.Float64()*5)+5, ",") // 7.1885709359349015
	fp((rand.Float64()*5 + 5))    // 7.123187485356329
	fpln()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fp(r1.Intn(100), ",") // 79
	fp(r1.Intn(100))      // 21
	fpln()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fp(r2.Intn(100), ",") // 5
	fp(r2.Intn(100))      // 87
	fpln()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fp(r3.Intn(100), ",") // 5
	fpln(r3.Intn(100))    // 87
}
