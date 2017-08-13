###　常量的定义

- 常量的值在编译时就已经确定
- 常量的定义格式与变量基本相同
- 等号右侧必须是常量或者常量的表达式
- 常量的表达式中的函数必须是内置函数
- 示例：
 ```
 const a int = 1
 const b = 'A'

 const (
    c = a
    d = a + 1
    e = a * 2
 )

 const f,g,h = 1, "2" ,'t'

 ```

### 常量的初始化规则与枚举

- 在定义常量组时，如果不提供初始值，则表示将使用上行的表达式
- 使用相同的表达式不代表具有相同的值
 ```
 import (
 	"fmt"
 )

 var str = "test" // 全局变量在在编译时不存在, 而const定义的常量存在

 const (
 	a = 2
 	b
 	c = "12345"
 	d = len(c)
 	// dd = len(str) // 报错 `const initializer len(str) is not a constant`

 	e, f = 1, "2"
 	g, h // 备注此处如果想省略同上行，必须个数保持一致，不能只有一项

 )

 func main() {
 	fmt.Println(a) // 2
 	fmt.Println(b) // 2 // 未定义则同上行的值
 	fmt.Println(c) // 12345
 	fmt.Println(d) // 5
 	// fmt.Println(dd) // 无法输出，此处报错
 	fmt.Println()
 	fmt.Println(g) // 1
 	fmt.Println(h) // 2
 }

 ```
- iota是常量的计数器，从0开始，组中每定义1个常量自动递增1
- 通过初始化规则与iota可以达到枚举的效果
- 每遇到一个const关键字， iota就会重置为0
 ```
 import (
 	"fmt"
 )

 const (
 	a = 'A'
 	b
 	c = iota
 	d
 )

 const (
 	e = 'e'
 	f = iota
 )

 func main() {
 	fmt.Println(a) // 65
    fmt.Println(b) // 65
 	fmt.Println(c) // 2
 	fmt.Println(d) // 3
 	fmt.Println(e) // 69
 	fmt.Println(f) // 1  这里f在一个新的const中从0开始数值为1
 }

 ```
- 上述只为举例，一般常量定义都需要使用大写，如果不想被外部访问到，需要使用`_`或`c` ，如下所示：
 ```
 const (
 	MAX_SIZE = 10 // 可以被外部调用
 	_TEST = "ST" // 不可被外部调用
 	cTEST2 = "ST2" // 不可被外部调用
 )

 ```