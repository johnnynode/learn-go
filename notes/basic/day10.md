### __map__

- 类似其他语言中的哈希表或字典，以key-value的形式存储数据
- key必须是支持==或!=比较运算的类型，不可以是函数、map或slice
- map查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍
- map使用make() 创建，支持`:=`这种简写方式。
- make([keyType]valueType,cap), cap表示容量，可省略。
- 超出容量时会自动扩容，但尽量提供一个合理的初始值。
- 使用len() 获取元素个数
- 键值对不存在时自动添加，使用delete() 删除某键值对。
- 使用for range对map和slice进行迭代操作
- 示例代码：
```
package main

import (
	"fmt"
	"sort"
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
}

// 普通创建方式
func test1() {
	var m map[int]string // 中括号里面的是key的类型，外面string是value的类型
	m = map[int]string{}
	fmt.Println(m) // map[]
}

// 使用make关键字创建
func test2() {
	var m map[int]string
	m = make(map[int]string)
	fmt.Println(m) // map[]
}

// 直接声明和赋值创建
func test3() {
	var m map[int]string = make(map[int]string)
	fmt.Println(m) // map[]
}

// 使用更简单的方式
func test4() {
	m := make(map[int]string)
	fmt.Println(m) // map[]
}

func test5() {
	m := make(map[int]string)
	m[1] = "OK"
	a := m[1]
	b := m[2]

	fmt.Println(m) // map[1:OK]
	fmt.Println(a) // OK
	fmt.Println(b) //  // 空的字符串什么也没有
	fmt.Println("------")
	// 中途删除
	delete(m, 1) // 第一个参数是map对象，第二个参数是key
	fmt.Println(a) //  OK // a是字符串的变量，还在
	fmt.Println(m[1]) //  // 删除之后，map中的key为1的值没有了。
}

// 复杂类型map
func test6() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	m[1] = make(map[int]string) // 此处是为了使用make初始化里面的map
	m[1][1] = "OK"
	a := m[1][1]
	fmt.Println(a) // OK
	fmt.Println("....")
}

// 使用多返回值来判断map是否被初始化
func test7() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	a, ok := m[2][1] // 此时并未被初始化
	fmt.Println(a, ok) //  false // 输出一个空值和false
	if !ok {
		m[2] = make(map[int]string)
	}
	m[2][1] = "OK"
	a, ok = m[2][1]
	fmt.Println(a, ok) // OK true // 此处被正确初始化了
}

// 迭代操作
func test8() {
	s1 := []int{1,2,3}
	for i,v:=range s1{
		fmt.Println(i, v)
	}
	fmt.Println("----")
	m1 := make(map[int]string)
	m1[1] = "Hello"
	m1[2] = "World"
	for k,v:=range m1{
		fmt.Println(k, v)
	}
	/*
	// 输出结果为：
	0 1
	1 2
	2 3
	----
	1 Hello
	2 World

	*/


}

// 以map为值的slice类型
func test9() {
	sm := make([]map[int]string, 2)
	for i,v := range sm{
		v = make(map[int]string, 1)
		v[i] = "OK"
		fmt.Println(v)
	}
	fmt.Println("----")
	fmt.Println(sm)
	/*
	map[0:OK]
	map[1:OK]
	----
	[map[] map[]]

	*/

	// 通过上面的运行结果可知， v是map中一个索引的拷贝，改变v，并不改变map，slice同样是这样。
}

func test10() {
	sm := make([]map[int]string, 2)
	for k := range sm{
		sm[k] = make(map[int]string, 1)
		sm[k][k] = "OK"
		fmt.Println(sm[k])
	}
	fmt.Println("----")
	fmt.Println(sm)
	/*
	map[0:OK]
	map[1:OK]
	----
	[map[0:OK] map[1:OK]]

	*/
}

// 对map进行间接排序，map是无序的，通过key来排序
func test11() {
	m := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"} // 创建一个map
	s := make([]int,len(m)) // 创建一个slice
	i := 0
	for k,_ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s) // 对s进行排序
	fmt.Println(s)
	fmt.Println("----")
	// 根据拍好序的k，取出对应map中的值，从而达到对map的排序
	for _, v := range s {
		fmt.Println(m[v])
	}
	/*
	[1 2 3 4 5]
	----
	a
	b
	c
	d
	e

	*/
}
```