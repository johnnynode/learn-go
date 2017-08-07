### __slice切片__

- 其本身并不是数组，它指向底层的数组
- 作为变长数组的替代方案，可以关联底层数组的局部或全部
- 类型为引用类型
- 可以直接创建或从底层数组获取生成
- 使用len() 获取元素个数， cap() 获取容量
- 一般使用make() 创建
- 如果多个slice指向相同底层数组，其中一个的值改变会影响全部
- make([]T,len,cap)
- 其中cap可以省略，则和len的值相同
- len表示存数的元素个数，cap表示容量
- Reslice 时索引以被slice的切片为准
- 索引不可以超过被slice的切片的容量cap()值
- 索引月切不会导致底层数组的重新分配而是引发错误
- 可以在slice尾部追加元素
- 可以将一个slice追加在另一个slice尾部
- 如果最终长度未超过追加到slice的容量则返回原始slice
- 如果超过追加到的slice的容量则将重新分配数组并拷贝原始数据
- 示例代码

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
}

func test1() {
	var s1 []int
	fmt.Println(s1) // []

}

func test2() {
	a := [5]int{1,2,3,4,5}
	fmt.Println(a) // [1 2 3 4 5]
	s1 := a[2]
	fmt.Println(s1) // 3 此处是取值
	s2 := a[4:5] // 左闭右开 此处是取数组
	s3 := a[4:len(a)]
	s4 := a[4:]
	fmt.Println(s2) // [5]
	fmt.Println(s3) // [5]
	fmt.Println(s4) // [5]
	// 备注 s2, s3, s4 都是从索引为4的位置，取值到最后
}

func test3() {
	a := [5]int{1,2,3,4,5}
	fmt.Println(a) // [1 2 3 4 5]
	s1 := a[:3] // 取前3项，不会包含索引为3的，同样是左闭右开
	fmt.Println(s1) // [1 2 3]
}

func test4() {
	s1 := make([]int, 3, 5) // 初始化3个int放入数组，第三个参数是初始容量
	fmt.Println(s1) // [0 0 0]
	fmt.Println(len(s1),cap(s1)) // 3 5
}

func test5() {
	a := []byte{'a','b','c','d','e','f','g','h','i','j'}
	b := string(a[2:5])
	c := string(a[3:5])
	fmt.Println(b) // cde
	fmt.Println(c) // de
}

func test6() {
	s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1) // 0xc042047f50
	s1 = append(s1,1,2,3)
	fmt.Printf("%v %p\n", s1, s1) // [0 0 0 1 2 3] 0xc042047f50 // 此处表明，在未超过最大容量候，操作的都是同一个地址
	s1 = append(s1,1,2,3)
	fmt.Println("cap:", cap(s1)) // cap:12 // 此处可以看到原容量是6，现在就会扩大一倍，每次都是一倍一倍的增长，地址也是重新分配
	fmt.Printf("%v %p\n", s1, s1) // [0 0 0 1 2 3 1 2 3] 0xc0420420c0 // 此处地址变化了
}

func test7() {
	a := []int{1,2,3,4,5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2) // [3 4 5] [2 3] // 可以看出 s1 和 s2 底层共享元素3的，指向的是底层数组的同一个元素
	s1[0] = 9 // 改变这个元素 在s1中是索引为0的数是3，现在变成9
	fmt.Println(s1, s2) // [9 4 5] [2 9]
}

func test8() {
	a := []int{1,2,3,4,5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2) // [3 4 5] [2 3] // 可以看出 s1 和 s2 底层共享元素3的，指向的是底层数组的同一个元素
	fmt.Println(len(s2),cap(s2)) // 2 4
	s2 = append(s2, 1,1,1) // append 超越s2的容量, 新的内存地址被分配，将不指向同一个底层数组了， 此时就不会被改变
	s1[0] = 9 // 改变这个元素 在s1中是索引为0的数是3，现在变成9
	fmt.Println(s1, s2) // [9 4 5] [2 3 1 1 1]
}

func test9() {
	s1 := []int{1,2,3,4,5,6}
	s2 := []int{7,8,9}
	copy(s1,s2) // 将s2拷贝到s1中
	fmt.Println(s1) // [7 8 9 4 5 6] // 可以看出保持s1的长度
}

func test10() {
	s1 := []int{1,2,3,4,5,6}
	s2 := []int{7,8,9}
	copy(s2,s1) // 将s1拷贝到s2中
	fmt.Println(s2) // [1 2 3] // 可以看出保持上s2的长度
}

func test11() {
	s1 := []int{1,2,3,4,5,6}
	s2 := []int {7,8,9,10,1,1,1,1,1}
	copy(s2[2:4], s1[1:3])
	fmt.Println(s2) // [7 8 2 3 1 1 1 1 1]
}

```