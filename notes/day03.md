### __批量定义常量__
```
const (
    PI = 3.14
    c1 = 1
    c2 = 2
    c3 = 3
)
```
其中大写的外部可调用，小写的不可调用

### __批量定义全局变量__
```
var (
    name = 'n1'
    age = 10
)
```

### __批量声明一般类型__
```
type(
    newType int
    type1 float32
    type2 string
    type3 byte
)
```

### __Go 基本类型__

- 布尔值： bool
 * 长度 1字节
 * 取值范围 true, false
 * 注意事项：不可用数字代表true或false
<br>
- 整形： int/uint
 * 根据运行平台可能为32位或64位
<br>
- 8位整形： int8/uint8
 * 长度： 1字节
 * 取值范围：-128 ~127 / 0 ~ 255
<br>
-  字节型： byte (uint8别名)
<br>
- 16位整形： int16/uint16
 * 长度：2字节
 * 取值范围：-32768 ~ 32767 或 0 ~ 65535
<br>
- 32位整型： int32(rune)/uint32
  * 长度： 4字节
  * 取值范围： -2^32/2 ~ 2^32/2-1 或 0 ~ 2 ^ 32 - 1
  * 说明：rune 是有符号的32位整型的别名, 处理unicode字符相关
<br>
- 64位整形： int64/uint64
 * 长度： 8字节
 * 取值范围： -2^64/2 ~ 2^64/2 - 1 或 0 ~ 2 ^ 64 - 1
<br>
- 浮点型： float32 / float64
  * 长度： 4/8字节
  * 小数位： 精确到7/15小数位
<br>
- 复数：complex64 / complex128
 * 长度： 8/16字节
<br>
- 足够保存指针的32位或64位的整数型：  uintptr
<br>
- 其他值类型： array 、 struct 、 string
<br>
- 引用类型： slice 、 map 、 chan
<br>
- 接口类型： interface
<br>
- 函数类型： func
<br>
### __类型零值__

 - 零值并不等于空值，而是当变量被声明为某种类型后的默认值

 - 通常情况下，值类型默认值为0， bool为false, string为空字符串，关键代码如下：
  ```
   var a int
   var b int32
   var c float32
   var d bool
   var e string
   var f []int
   var g [1]int
   var h [1]bool
   var j [1]byte
   fmt.Println(a) // 0
   fmt.Println(b) // 0
   fmt.Println(c) // 0
   fmt.Println(d) // false
   fmt.Println(e) // '' 什么都没输出
   fmt.Println(f) // []
   fmt.Println(g) // [0]
   fmt.Println(h) // [false]
   fmt.Println(j) // [0]

  ```

### __内置的类的静态常量__

```
 fmt.Println(math.MinInt8) // -128
 fmt.Println(math.MaxInt32) // 2147483647
```

### __类型的别名__

```
type(
    我是自定义string string // 此处给string 自定义别名
)

func main() {
    var b 我是自定义string
    b = "中文测试"
    fmt.Println(b);
}

```

### __单个变量的声明与赋值__

- 变量的声明格式： var 变量名称 变量类型 如：`var a int`
- 变量的赋值格式： 变量名称 = 表达式 如： `a = 1`
- 声明的同时赋值： var 变量名称 变量类型 = 表达式 如：`var a int = 1` 或 `var a = 1` (变量类型可以省略，由系统推断)
- 变量声明与赋值的最简写法：`d := 456`

### __多个变量的声明与赋值__

- 全局变量的声明可使用var()的方式进行简写，如：
 ```
 var (
	    a = 'hello'
	    b,c = 1,5
	    d := 3
 )
 ```

- 全局变量的声明不可以省略var, 但可用并行方式
 ```
 var a,b,c,d int = 1,2,3,4
 fmt.Println(a) // 1
 fmt.Println(b) // 2
 fmt.Println(c) // 3
 fmt.Println(d) // 4
 ```
- 所有变量都可以使用类型推断
- 局部变量不可以使用var()的方式简写，只能使用并行方式
 ```
 e,f,g,h := 1,2,3,4
 fmt.Println(e) // 1
 fmt.Println(f) // 2
 fmt.Println(g) // 3
 fmt.Println(h) // 4
 ```

- `:=` 的方式多用于函数有多个返回值时进行赋值或并行赋值忽略某项时使用,如：
 ```
  a, _, c,d = 1,2,3,4 // 使用 _ 来忽略其中的一项
  fmt.Println(a) // 1
  fmt.Println(d) // 3
  fmt.Println(d) // 4

 ```

### __变量的类型转换__

-  go 中不存在隐式转换，所有类型转换必须显式声明
- 转换只能发生在两种相互兼容的类型之间, 布尔类型不可以和数字进行转换
 ```
   import (
    "fmt"
    "strconv"
   )

   func main() {
    var a int = 65
    b := string(a)
    c := strconv.Itoa(a)
    d := strconv.Atoi(c)
    fmt.Println(b) // A  备注此处系统将65通过ASCii 码的形式转换成了 A
    fmt.Println(c) // 65 备注此处的65是字符形65
    fmt.Println(d) // 65 备注此处将字符型65转换成了数字65
   }

 ```
- 类型的转换格式：`<ValueA> [:]= <TypeOfValueA> (<ValueB>)` 而可选项 `:` 在变量未在上文定义的时候使用。
 ```
 var a float32 = 1.1
 b := int(a)
 fmt.Println(a) // 1.1
 fmt.Println(b) // 1

 ```