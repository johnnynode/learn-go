### __method__

- Go 中没有class, 但依旧有method
- 通过显式说明receier来实现与某个类型的组合
- 只能为同一个包中的类型定义方法
- Receiver可以是类型的值或指针
- 不存在方法重载
- 可以使用值或指针来调用方法，编译器会自动完成转换
- 从某种意义上来说，方法是函数的语法糖，因为reveiver其实就是方法所接收的第一个参数(Method Value vs. Method Expression)
- 如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
- 类型别名不会拥有底层类型所附带的方法
- 方法可以调用结构中的非公开字段。
- 示例代码：
```
package main

import (
	"fmt"
)

type A struct {
	Name string
	Age int
	grade string // 私有字段 以package为级别，都是公有的，不同包都是私有的
}

type TZ int // 初始化一个int型的类型别名 TZ

func main() {
	test1()
	test2()
    test3()
    test4()
    test5()
    test6()
}

// 定义 A 的方法 Print 结构为： func + 接收者 + 方法名 ； go语言中没有重载的概念
func (a A) Print() {
	fmt.Println("A")
}

// 设置名称 内部值拷贝
func (a A) SetName() {
	a.Name = "A"
}

// 设置年龄 内部指针拷贝
func (a *A) SetAge() {
	a.Age = 10
}

// 设置年级 访问私有字段
func (a *A) SetGrade() {
	a.grade = "g3"
}

// 为 TZ 类型的a 绑定一个Print方法
func (a *TZ) Print() {
	fmt.Println("TZ")
}

// 普通测试
func test1() {
	a := A{}
	a.Print() // A
}

// 测试值拷贝
func test2() {
	a := A{}
	a.SetName()
	fmt.Println(a.Name) // "" // 输出零值， 空字符串 , 说明：此处值拷贝，未修改成功。
}

// 测试指针拷贝
func test3() {
	a := A{}
	a.SetAge()
	fmt.Println(a.Age) // 10 // 此处修改成功
}

// 测试类型别名绑定的方法 method value
func test4() {
	var a TZ
	a.Print() // TZ
}

// 测试 method expression 的写法
func test5() {
	var a TZ
	(*TZ).Print(&a) // TZ // 这种方法
}

// 测试私有字段的访问权限
func test6() {
	a := A{}
	(*A).SetGrade(&a) // 通过 method expression 的写法 同 a.SetGrade()
	fmt.Println(a.grade) // "g3"
}

```
