package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var fpln = fmt.Println

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fpln(match) // true

	r, _ := regexp.Compile("p([a-z]+)ch")

	fpln(r.MatchString("peach")) // true

	fpln(r.FindString("peach punch")) // peach

	fpln(r.FindStringIndex("peach punch")) // [0 5]

	fpln(r.FindStringSubmatch("peach punch")) // [peach ea]

	fpln(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	fpln(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

	fpln(r.FindAllStringSubmatchIndex("peach punch pinch", -1)) // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	fpln(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	fpln(r.Match([]byte("peach"))) // true

	r = regexp.MustCompile("p([a-z]+)ch")
	fpln(r) // p([a-z]+)ch

	fpln(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fpln(string(out)) // a PEACH
}
