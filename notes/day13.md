### __tt__
- Go 中的struct与C中的struct非常相似，并且Go没有class
- 使用type <Name> struct{} 定义结构，名称遵循可见性规则
- 支持指向自身的指针类型成员
- 支持匿名结构，可用作成员或定义成员变量
- 匿名结构也可以用于map的值
- 可以使用字面值对结构进行初始化
- 允许直接通过指针来读写结构成员
- 相同类型的成员可进行直接拷贝赋值
- 支持 == 与 != 比较运算符，但不支持 > 或 <
- 支持匿名字段，本质上是定义了以某个类型名为名称的字段
- 嵌入结构作为匿名字段看起来像继承，但不是继承
- 可以使用匿名字段指针
- 示例代码：
```
package main

import (
	"fmt"
)

type person struct {
	Name string
	Age int
}

func main() {
	test1()
	test2()
	test3()
	test4T()
	test5T()
	test6T()
	test7()
	test8()
	test9()
	test10()
	test11()
	test12()
}

func test1() {
	a := person{}
	fmt.Println(a) // { 0} // 里面一个是空字符串，一个是0 ，它们是string和int的零值
}

func test2() {
	a := person{}
	a.Name = "joe"
	a.Age = 19
	fmt.Println(a) // {joe 19}
}

func test3() {
	a := person {
		Name: "joe",
		Age: 19,
	}
	fmt.Println(a) // {joe 19}
}

func test4(per person) {
	per.Age = 13
	fmt.Println("test4",per)
}

// 测试，对象作为值拷贝，不修改原对象
func test4T() {
	a := person{
		Name:"joe",
		Age:19,
	}
	fmt.Println(a) // {joe 19}
	test4(a) // test4 {joe 13}
	fmt.Println(a) // {joe 19}
}

func test5(per *person) {
	per.Age = 13
	fmt.Println("test4",per)
}

// 测试对象的引用拷贝，修改原对象
func test5T() {
	a := person{
		Name:"joe",
		Age:19,
	}
	fmt.Println(a) // {joe 19}
	test5(&a) // test4 &{joe 13} // 说明此处内部改变了
	fmt.Println(a) // {joe 13} // 说明此处外部也被改变了
}


func test6(per *person) {
	per.Age = 13
	fmt.Println("test4",per)
}

// 测试直接通过地址赋值, 推荐使用取地址符号
func test6T() {
	a := &person{
		Name:"joe",
		Age: 19,
	}
	fmt.Println(a) // {joe 19}
	test6(a) // test4 &{joe 13} // 说明此处内部改变了
	fmt.Println(a) // {joe 13} // 说明此处外部也被改变了
}

// 匿名结构
func test7() {
	a := struct {
		Name string
		Age int
	}{
		Name:"joe",
		Age: 19,
	}

	fmt.Println(a) // {joe 19}
}

// 匿名结构取地址
func test8() {
	a := &struct {
		Name string
		Age int
	}{
		Name:"joe",
		Age: 19,
	}

	fmt.Println(a) // &{joe 19}
}

// 匿名结构的嵌套
func test9() {
	type person struct {
		Name string
		Age int
		Contact struct {
			Phone, City string // 此处的写法是为了简便
		}
	}

	a := person{
		Name:"joe",
		Age:19,
	}
	// 对于匿名struct的初始化必须使用下面的方式
	a.Contact.Phone = "123"
	a.Contact.City = "Bj"
	fmt.Println(a) // {joe 19 {123 Bj}}
}

// 匿名字段
func test10() {
	type person struct {
		string
		int
	}

	a := person{"joe",19} // 一定要注意初始化的顺序，类型要对应，否则报错
	fmt.Println(a) // {joe 19}
}

// 相同类型，可以赋值
func test11() {
	a := person{Name:"joe", Age: 19}
	var b person // 初始化person类型的b
	b = a // 把 a 的值赋给b，开辟了一块新的内存空间
	a.Age = 20 // 修改a的值
	fmt.Println(a) // {joe 20}
	fmt.Println(b) // {joe 19} // 可见 b 未受影响
}

// 测试继承
func test12() {
	type teacher struct {
		person // 继承person
		Class , Grade string
	}

	a := teacher{person:person{Name:"a", Age:29} , Class:"c1", Grade:"g1"}

	fmt.Println(a) // {{a 29} c1 g1}

	a.Class = "c3"
	a.Grade = "g3"
	a.person.Name = "aa" // 可以使用这种方式调用
	a.Age = 30 // 也可以使用这种方式调用, 简单推荐
	fmt.Println(a) // {{aa 30} c3 g3}

}
```

