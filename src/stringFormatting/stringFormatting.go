package main

import (
	"fmt"
	"os"
)

var fp = fmt.Printf // 备注此处不能使用 fmt.Println 会有问题

var fpln = fmt.Println // 支持换行

var fsp = fmt.Sprintf

var ffp = fmt.Fprintf

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	fp("%v\n", p) // {1 2}

	fp("%+v\n", p) // {x:1 y:2}

	fp("%#v\n", p) // main.point{x:1, y:2}

	fp("%T\n", p) // main.point

	fp("%t\n", true) // true

	fp("%d\n", 123) // 123

	fp("%b\n", 14) // 1110

	fp("%c\n", 33) // !

	fp("%x\n", 456) // 1c8

	fp("%f\n", 78.9) // 78.900000

	fp("%e\n", 123400000.0) // 1.234000e+08

	fp("%E\n", 123400000.0) // 1.234000E+08

	fp("%s\n", "\"string\"") // "string"

	fp("%q\n", "\"string\"") // "\"string\""

	fp("%x\n", "hex this") // 6865782074686973

	fp("%p\n", &p) // 0x42135100

	fp("|%6d|%6d|\n", 12, 345) // |    12|   345|

	fp("|%6.2f|%6.2f|\n", 1.2, 3.45) // |  1.20|  3.45|

	fp("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // |1.20  |3.45  |

	fp("|%6s|%6s|\n", "foo", "b") // |   foo|     b|

	fp("|%-6s|%-6s|\n", "foo", "b") // |foo   |b     |

	s := fsp("a %s", "string")

	fpln(s) // a string

	ffp(os.Stderr, "an %s\n", "error") // an error
}
