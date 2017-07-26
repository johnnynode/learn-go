### Go 内置关键字 (25个均为小写)

> break default func interface select
> case defer go map struct
> chan else goto package switch
> const fallthrough if range type
> continue for import return

### Go 的注释方法

- //   : 单行注释
- /**/ : 多行注释

### Go 程序的一般结构
- 通过package来组织结构
- 只有package名称为main的包可以包含main函数
- 一个可执行程序有且只有一个main函数
- 通过import关键字来导入其他非main包
- 通过const关键字来进行常量的定义
- 通过在函数体外部使用var关键字来进行全局变量的声明与赋值
- 通过type关键字来进行结构(struct)或(interface)的声明
- 通过func关键字来进行函数的声明

### 导入package的方式和注意事项
- 逐个导入：
 ```
 import "fmt"
 import "os"
 import "time"
 ```

- 统一导入：
 ```
 import (
    "fmt"
    "os"
    "time"
 )
 ```

- 导入包之后，就可以使用格式<PageageName>.<FuncName> 来对包中的函数进行调用

- 注意： 如果导入了包未被调用，则其中的函数或类型将会报出编译错误

### package 别名的使用

- 当导入第三方包时，包名很可能非常接近或相同，需要使用别名来进行区别和调用如：
 ```
  import std "fmt"
  func main() {
    std.Println("HelloWorld!");
  }

 ```

- 省略调用如：
 ```
 import . "fmt"
 func main() {
    Println("HelloWorld!");
 }
 ```
 省略调用的注意事项：不可与别名同时使用且不建议在实际项目中采用

### 可见性规则

- 在Go语言中通过大小写来决定该变量、常量、类型、接口、结构或函数是否可被外部所调用
- 根据约定，函数名首字母小写即为private, 大写即为public