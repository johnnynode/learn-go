### __defer与panic/recover__

- 执行方式类似其他语言中的析构函数，在函数体执行结束后按照调用顺序的相反顺序逐个执行
- 即使函数发生严重错误也会执行,类似finally
- 支持匿名函数的调用
- 常用于文件清理、文件关闭、解锁以及记录时间等操作
- 通过与匿名函数配合可在return之后修改函数计算结果
- 如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时即已获得了拷贝，否则则是引用某个变量的地址
- Go 中没有异常机制，但有panic/recover模式来处理错误
- Panic 可以在任何地方引发，但recover只有在defer调用函数中有效
- 示例代码：
```
package main

import (
	"fmt"
)

func main() {

}

// defer 定义的先被调用，逆序向上调用
func test1() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	/*
	// 输出结果
	a
	c
	b

	*/
}

// 继续测试，逆向调用
func test2() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	/*
	// 输出结果
	2
	1
	0

	*/
}

// 使用匿名函数测试
func test3() {
	for i := 0; i < 3; i++ {
		defer func () {
			fmt.Println(i) // 此时defer改变了i作为地址的引用，一直引用这个局部变量i
		}()
	}
	/*
	// 输出结果
	3
	3
	3

	*/
}

// 测试panic
func test4() {
	fmt.Println("Func in 4")
}

func test5() {
	// defer 函数必须先调用，然后再执行panic
	// 通过defer把程序从panic状态recover恢复回来
	defer func() {
		if err:=recover(); err != nil {
			fmt.Println("Recover in 5")
		}
	}
	panic("Panic in 5") // panic之后的所有语句都不会执行
}

func test6() {
	fmt.Println("Func in 6")
}

func test456T() {
	test4() // Func in 4
	test5() // Recover in 5
	test6() // Func in 6
}

```