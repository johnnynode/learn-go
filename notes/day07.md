### __判断语句if__

- 条件表达式没有括号
- 支持一个初始化表达式 (可以并行方式)
- 做大括号必须和条件语句或else在同一行
- 支持单行模式
- 初始化语句中的变量为block级别，同时隐藏外部同名变量
- 示例代码：
 ```
 package main

 import (
 	"fmt"
 )

 func main() {
 	test1()
 	test2()
 	test3()
 	test4()
 	test5()
 }

 func test1() {
 	a := 1
 	if a == 1 {
 		fmt.Println("a == 1") // a == 1
 	}
 }

 func test2() {
 	if b := 2; b == 2 {
 		fmt.Println("b == 2") // b == 2
 	}
 	// fmt.Println(b) // 不能直接输出，会报错 undefined:b
 }

 func test3() {
 	if c, d := 3,4; c < d {
 		fmt.Println("c < d") // c < d
 	}
 }

 func test4() {
 	e := 10
 	if e:=5; e < 10 {
 		fmt.Println(e) // 5
 	}
 	fmt.Println(e) // 10
 }

 func test5() {
 	if e:=5; e > 10 {
 		fmt.Println("e > 10")
 	} else {
 		fmt.Println("e <= 10") // e <= 10
 	}
 }
 ```
