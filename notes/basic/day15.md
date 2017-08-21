### __interface__

- 接口是一个或多个方法签名的集合
- Go 语言中所有类型都实现了空接口 (空接口是没有任何东西的接口)
- 只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示声明实现了哪个接口，这称为Structural Typing
- 接口只有方法声明，没有实现，没有数据字段
- 接口可以匿名嵌入其他接口，或嵌入到结构中
- 将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，即无法修改复制品的状态，也无法获取指针
- 只有当接口存储的类型和对象都为nil时，接口才等于nil
- 接口调用不会做receiver的自动转换
- 接口同样支持匿名字段方法
- 接口也可实现类似OOP中的多态
- 空接口可以作为任何类型数据的容器
- 通过类型断言的ok pattern 可以判断接口中的数据类型
- 使用 type switch 则可针对空接口进行比较全面的类型判断
- 可以将拥有超集的接口转换为子集的接口
- 示例代码：
```
package main

import (
	"fmt"
)

// 此处演示空接口
type empty interface {

}

type USB interface {
	Name() string // name 方法 返回名称 string是返回类型
	Connect()
	Showing // 此处是接口的嵌套继承
}

type Showing interface {
	Show()
}

// 定义一个结构体为了实现上述的USB接口
type PhoneConnecter struct {
	name string
}

// 模拟实现的方法 Name
func (pc PhoneConnecter) Name() string {
	return pc.name
}

//模拟实现方法 Connect
func (pc PhoneConnecter) Connect () {
	fmt.Println("Connect:", pc.name)
}

//模拟实现方法 Connect
func (pc PhoneConnecter) Show () {
	fmt.Println("Showing:", pc.name)
}

// 测试方法
func Disconnect(usb USB) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected.", pc.name)
		return
	}
	fmt.Println("Unknown device.")
}

// 测试方法2 传入空接口
func Disconnect2(usb interface{}) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected.", pc.name)
		return
	}
	fmt.Println("Unknown device.")
}

// 测试方法3 使用switch来判断不同传入值的类型
func Disconnect3(usb interface{}) {
	switch v:= usb.(type) {
		case PhoneConnecter:
			fmt.Println("Disconnected.", v.name)
		default:
			fmt.Println("Unknown device.")
	}
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
}

// 显式实现
func test1() {
	var a USB // 这个意思是 a 实现了USB的接口
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect() // Connect: PhoneConnecter
	Disconnect(a) // Disconnected. PhoneConnecter // 说明自动实现了
}

// 未显示实现
func test2() {
	a := PhoneConnecter{"PhoneConnecter"} // 此处未直接实现 准备测试 Structural Typing机制
	a.Connect() // Connect: PhoneConnecter
	Disconnect(a) // Disconnected. PhoneConnecter // 说明自动实现了
}

// 未显示实现 使用Disconnect2方法，测试传入空接口
func test3() {
	a := PhoneConnecter{"PhoneConnecter"} // 此处未直接实现 准备测试 Structural Typing机制
	a.Connect() // Connect: PhoneConnecter
	Disconnect2(a) // Disconnected. PhoneConnecter // 说明自动实现了
}

// 未显示实现 使用Disconnect2方法，测试使用switch
func test4() {
	a := PhoneConnecter{"PhoneConnecter"} // 此处未直接实现 准备测试 Structural Typing机制
	a.Connect() // Connect: PhoneConnecter
	Disconnect3(a) // Disconnected. PhoneConnecter // 说明自动实现了
}

// 接口之间的转换，只能从超集的接口转换到子集的接口
func test5() {
	pc := PhoneConnecter{"PhoneConnecter"} // 此处由于Structural Typing机制 默认实现了USB接口
	var a Showing // a 实现 Showing 接口
	a = Showing(pc) // 将USB类型的接口强制转换成Showing类型的接口
	a.Show() // Showing:PhoneConnecter
	// a.Name() // 尝试调用USB接口中的方法，此处会报错, a 已经被降级了
}

// 测试：直接修改属性是可以的
func test6() {
	pc := PhoneConnecter{"PhoneConnecter"}
	pc.name = "pc"
	pc.Connect() // Connect: pc // 说明：此处修改了
}

// 测试 ：将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，即无法修改复制品的状态，也无法获取指针
func test7() {
	pc := PhoneConnecter{"PhoneConnecter"}
	var a Showing
	a = Showing(pc)
	a.Show() // Showing: PhoneConnecter
	pc.name = "pc" // 尝试修改复制品的状态
	a.Show() // Showing: PhoneConnecter // 说明尝试失败
}

// 测试：只有当接口存储的类型和对象都为nil时，接口才等于nil
func test8() {
	var a interface{}
	fmt.Println(a == nil) // true

	var p * int = nil
	a = p // a 指向 p
	fmt.Println(a == nil) // false
}

```