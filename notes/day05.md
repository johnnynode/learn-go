### __Go 中的运算符__

- Go 中的运算符都是从左到右的

- 优先级 (从高到低)
 * 一元运算符：
   ```^  !```
 * 二元运算符：
   ```
   \*   /  %  <<  >>  &  &^
   \+  -  |  ^
   == != < <= >= >
   <-
   &&
   ||

   ```

### __& 与 && 的区别__
- & 会将两边的表达式同时执行

- && 只要左边的表达式不成立，右边的表达式便不会被执行

### __示例代码：__
```
import (
	"fmt"
)

/*
 6: 0110
11: 1011
-------------
&   0010 = 2
|   1111 = 15
^   1101 = 13
&^  0100 = 4
*/

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	PB
	TB
)

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	fmt.Println(!ture) // false
	fmt.Println(1 << 10) // 1024
	fmt.Println(1 << 10 << 10) // 1048576
	fmt.Println(1 << 10 << 10 >> 10) // 1024
	fmt.Println()
	fmt.Println(6 & 11) // 2
	fmt.Println(6 | 11) // 15
	fmt.Println(6 ^ 11) // 13
	fmt.Println(6 &^ 11) // 4
}

func test2() {
	a := -1
	if a > 0 && (10/a) > 1 {
		fmt.Println("fine!") // 不会被执行
	}
}

func test3() {
    fmt.Println(B)  // 1
    fmt.Println(KB) // 1024
    fmt.Println(MB) // 1.048576e+06
    fmt.Println(GB) // 1.073741824e+09
    fmt.Println(PB) // 1.099511627776e+12
    fmt.Println(TB) // 1.125899906842624e+15
}

```