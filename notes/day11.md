### __function__
- Go 函数不支持嵌套、重载和默认参数
- 支持：无需声明原型、不定长度变参、多返回值、命名返回值参数、匿名函数、闭包
- 定义函数使用关键字`func`,且左大括号不能另起一行
- 函数也可以做为一种类型使用
- 示例代码：
```
package main

import (
	"fmt"
)

func main() {

}

// 参数列表和返回值，注意类型
func test1(a int b string) (int string) {
}

// 无返回值
func test2(a int b string) {
}

// 只有一个返回值
func test3(a int b string) bool {
}

// 参数类型一致，可以只写一个类型
func test4(a,b,c int) int {
}

// 返回值类型一致，可以只写一个类型, 命名返回值
func test5(a,b,c int) (a,b,c int) {
	a, b, c = 1, 2, 3 // 直接赋值= 不需要 :=
	return // 此时可以不写明具体的返回值，已经在上面命名了。
}

// 返回值类型一致，不命名返回值
func test6(a,b,c int) (int, int, int) {
	a, b, c := 1, 2, 3
	return a, b, c // 此处必须跟上返回了什么
}

// 相同类型的参数，但不定长，可以使用不定长变参, 不定长变参如 ...int 必须作为参数中的最后一项
func test7(a ...int) {
	fmt.Println(a) // 此时a代表所有的参数，类型为slice 如：[1 2 3]
}

// 值的拷贝，不影响原值
func test8(s ...int) {
	s[0] = 3 // 此处改变了参数列表的值
	s[1] = 4
	fmt.Println(s)
}

func test8T() {
	a, b := 1, 2
	test8(a, b) // [3, 4]
	fmt.Println(a, b) // 1 2 // 此处说明值并没有被改变
}

// 内存地址的拷贝，改变原值
func test9(s []int) {
	s[0] = 5
	s[1] = 6
	s[2] = 7
	s[3] = 8
	fmt.Println(s)
}

func test9T() {
	s1 := []int{1,2,3,4}
	test9(s1) // [5 6 7 8]
	fmt.Println(s1) // [5 6 7 8] // 可以看到原先传递进来的值被修改了
}

// 想把int型和string型等简单类型进行地址的拷贝，使用指针的方式， 参数传递地址
func test10(a *int) {
	*a = 2 // 改变指向
	fmt.Println(*a)
}

func test10T() {
	a := 1 // 原值
	test10(&a) // 2 // 内部被改变了指向
	fmt.Println(a) // 2 // 可以看到原值被改变
}

// 函数可以被赋值使用
func test11() {
	fmt.Println("Func 11")
}

func test11T() {
	a := test11 //
	a() // Func 11
}

// 闭包的支持
func test12(x int) func(int) int {
	return func (y int) int{
		return x + y
	}
}

func test12T() {
	f := test12(10)
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 12
}

```