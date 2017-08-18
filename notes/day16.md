### __reflection__

- 反射可大大提高程序的灵活性，使得interface{}有更大的发挥余地
- 反射使用TypeOf 和 ValueOf 函数从接口中获取目标对象信息
- 反射会将匿名字段作为独立字段(匿名字段本质)
- 想要利用反射修改对象状态，前提是interface.data 是settable, 即 pointer-interface
- 通过反射可以"动态调用方法"
- 示例代码：
```
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

type Manager struct {
	User
	title string
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("Bad")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("Good")
	}

}

func (u User) Hello() {
	fmt.Println("Hello world")
}

func (u User) SayHello(name string) {
	fmt.Println("Hello", name, ", my name is ", u.Name)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o) // 获得接收的接口类型
	fmt.Println("t:", t) // t: main.User
	fmt.Println("Type:", t.Name()) // Type: User

	// 根据传入参数判断非指针类型不能继续操作
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("不能传递&u，这样的指针类型")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i:= 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface() // 通过这种方法取出字段所对应的值
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
	/**
	// for循环输出
	Fields:
	    Id: int = 1
	  Name: string = OK
	   Age: int = 12

	 */

	  fmt.Println("====================")

	  for i:=0; i<t.NumMethod(); i++ {
	  	m := t.Method(i)
	  	fmt.Printf("%6s: %v\n", m.Name, m.Type)
	  }

	  /**
	    // 打印出的结果：
	    Hello: func(main.User)
	  */
}



func main() {
	test1()
    test2()
    test3()
    test4()
    test5()
    test6()
}

func test1() {
	u := User{1, "OK", 12}
	Info(u)
	// Info(&u) // 如果传入这样的值，就会终止执行
}

// 测试匿名字段
func test2() {
	 m := Manager {User: User{1, "OK", 12}, title:"123"}
	 t := reflect.TypeOf(m)

	 fmt.Printf("%#v\n", t.Field(0))
	 fmt.Println("=================")
	 fmt.Printf("%#v\n", t.Field(1))
	 fmt.Println("-----------------")
	 fmt.Printf("%#v\n", t.FieldByIndex([]int{0,0})) // 通过传递进去的slice，来取匿名字段
	 fmt.Println()
	 fmt.Printf("%#v\n", t.FieldByIndex([]int{0,1})) // 通过传递进去的slice，来取匿名字段
   fmt.Println()

	 /*打印出的结果：*/
	 /*
	 reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x4a6380), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
	 =================
   reflect.StructField{Name:"title", PkgPath:"main", Type:(*reflect.rtype)(0x4987e0), Tag:"", Offset:0x20, Index:[]int{1}, Anonymous:false}
   -----------------
   reflect.StructField{Name:"Id", PkgPath:"", Type:(*reflect.rtype)(0x4987e0), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:false}

   reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x498ea0), Tag:"", Offset:0x8, Index:[]int{1}, Anonymous:false}
	  */
}

// 使用reflect 修改基本类型的值
func test3() {
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x) // 999
}

// 通过反射对接口对象中的值进行修改
func test4() {
	u := User{1, "OK", 12}
	Set(&u)
	fmt.Println(u) // {1 Good 12}
}

// 直接调用修改方法
func test5() {
	u := User{1, "OK", 12}
	u.SayHello("John") // Hello John , my name is  OK
}

// 通过reflect来调用方法
func test6() {
	u := User{1, "OK", 12}
	v := reflect.ValueOf(u)
  mv := v.MethodByName("SayHello")

  args := []reflect.Value{reflect.ValueOf("john")}
  mv.Call(args) // Hello john , my name is  OK
}

```