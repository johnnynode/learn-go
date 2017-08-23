package main

import (
	s "strings"
	"fmt"
)

var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count:  ", s.Count("test", "t"))
	p("HasPrefix:  ", s.HasPrefix("test", "te"))
	p("HasSuffix:  ", s.HasSuffix("test", "st"))
	p("Index:  ", s.Index("test", "e"))
	p("Join:  ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:  ", s.Repeat("a", 5))
	p("Replace:  ", s.Replace("foo", "o", "0", -1))
	p("Replace:  ", s.Replace("foo", "o", "0", 1))
	p("Split:  ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:  ", s.ToLower("TEST"))
	p("ToUpper:  ", s.ToLower("test"))
	p()

	p("Len", len("hello"))
	p("Char", "hello"[1])
}

/*

$ go run stringFunctions.go
Contains:  true
Count:   2
HasPrefix:   true
HasSuffix:   true
Index:   1
Join:   a-b
Repeat:   aaaaa
Replace:   f00
Replace:   f0o
Split:   [a b c d e]
ToLower:   test
ToUpper:   test

Len 5
Char 101

*/