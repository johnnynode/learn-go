### __数组__

- 定义数组的格式： `var <varName> [n]<type>` ,前提：n >= 0
- 数组程度也是类型的一部分，因此具有不同长度的数组为不同类型
- 注意区分指向数组的指针和指针数组
- 数组在Go中为值类型
- 数组之间可以使用 == 或 != 进行比较，但不能比较大小
- 可以使用new来创建数组，此方法返回一个指向数组的指针
- Go 支持多维数组
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
	test6()
	test7()
	test8()
	test9()
	test10()
	test11()
	test12()
}

func test1() {
	var a [2]int // 长度为2的int数组
	var b [1]int

	// 不能直接复制 b = a 这是不合法的，他们长度不同，不能直接赋值
	fmt.Println(a) // [0 0] // 字面值不够，则会用零值来补充
	fmt.Println(b) // [0]
}

// 相同长度类型的可赋值
func test2() {
	// var a [2]int{1,2} // 不能这样用，错误的定义和赋值
	a := [2]int{1,1}
	var b [2]int
	b = a
	fmt.Println(a) // [1,1]
	fmt.Println(b) // [1,1]
}

// 类型不同，无法赋值
func test3() {
	a := [2]string{"1","2"}
	b := [2]int{}

	// b = a // 错误的赋值, 不同类型无法转换，报错

	fmt.Println(a) // [1 2]
	fmt.Println(b) // [0 0]
}

// 使用索引
func test4() {
	a := [5]int{1:1,2} // 指定索引为1的值为1,其后一项为2，其余零值补充
	b := [5]int{4:1} // 指定索引为4的，也就是第五号元素为1，其余零值补充
	fmt.Println(a) // [0 1 2 0 0]
	fmt.Println(b) // [0 0 0 0 1]
}

// 使用...和索引来指定数组
func test5() {
	a := [...]int{1,2,3,4,5} // 使用...来定义一个已知长度的数组
	b := [...]int{0:1, 1:11, 2:22}

	fmt.Println(a) // [1 2 3 4 5]
	fmt.Println(b) // [1 11 22]
}

// 指向数组的指针
func test6() {
	a := [...]int{1,2,3}
	var p *[3]int = &a // p 是指向数组的指针
	fmt.Println(p) // &[1 2 3]
}

// 指针数组
func test7() {
	x, y := 1,2
	a := [...]*int{&x, &y}
	fmt.Println(a) // [0xc0420441d0 0xc0420441d8] // 说明 a 是一个指针数组
}

// 数组之间可以使用 == 或者 != 的比较， 但不可以比较大小 ; 比较前提是：长度，类型相同
func test8() {
	a := [2]int{1,2}
	b := [2]int{1,3}
	c := [2]int{1,2}

	fmt.Println(a == b) // false
	fmt.Println(a == c) // true
}

// 使用new关键字会返回一个指向数组的指针
func test9() {
	p := new([3]int)
	fmt.Print(p) // &[0 0 0]
}

// 定义一个第一维长度为2的，第二维长度为3的 二维数组
func test10() {
	a := [2][3]int{
		{1,1,1},
		{2,2,2}}
	fmt.Println(a) // [[1 1 1] [2 2 2]]
}

// 第一维度数组可以使用...来计算, 注意定义的数组长度，不能溢出
func test11() {
	a := [...][5]int{
		{1,2,3},
		{2,2,2,2}}
	fmt.Println(a) // [[1 2 3 0 0] [2 2 2 2 0]]
}

// 冒泡排序
func test12() {
	a := [...]int{5,2,3,1,4}
	fmt.Println(a)
	num := len(a)
	for i:=0; i < num; i++ {
		for j := i+1; j<num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
	/*
	[5 2 3 1 4]
	[5 4 3 2 1]
	*/
}
```