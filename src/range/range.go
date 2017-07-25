package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum) // sum: 9

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key", k)
	}

	for i, c := range "methods" {
		fmt.Println(i, c);
	}
}

/*

$ methods run range.methods

sum: 9
index: 1
a -> apple
b -> banana
key a
key b
0 103
1 111

*/
