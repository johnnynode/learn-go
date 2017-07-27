package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var fpln = fmt.Println

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int `json:"page"`
	Fruits []string  `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fpln(string(bolB)) // true

	intB, _ := json.Marshal(1)
	fpln(string(intB)) // 1

	fltB, _ := json.Marshal("2.34")
	fpln(string(fltB)) // 2.34

	strB, _ := json.Marshal("gopher")
	fpln(string(strB)) // gopher

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fpln(string(slcB)) // ["apple","peach","pear"]

	mapD := map[string]int{"apple":5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fpln(string(mapB)) // {"apple":5, "lettuce":7}

	res1D := &Response1 {
		Page : 1,
		Fruits :[]string{"apple","peach","pear"}}
	res1B, _ := json.Marshal(res1D)
	fpln(string(res1B)) // {"Page":1,"Fruits":["apple","peach","pear"]}

	res2D := &Response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fpln(string(res2B)) // {"Page":1,"fruits":["apple","peach","pear"]}

	byt := []byte(`{"num":6.13,"strs":["a", "b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fpln(dat) // map[num:6.13 strs:[a b]]

	num := dat["num"].(float64)
	fpln(num) // 6.13

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fpln(str1) // a

	str := `{"page":1, "fruits":["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fpln(res) // {1 [apple peach]}
	fpln(res.Fruits[0]) // apple

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple":5, "lettuce":7}
	enc.Encode(d) // {"apple":5,"lettuce":7}
}