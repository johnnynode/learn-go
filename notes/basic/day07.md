### __判断语句if__

- 条件表达式没有括号
- 支持一个初始化表达式 (可以并行方式)
- 做大括号必须和条件语句或else在同一行
- 支持单行模式
- 初始化语句中的变量为block级别，同时隐藏外部同名变量
- 示例代码：
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
 }

 func test1() {
    a := 1
    if a == 1 {
        fmt.Println("a == 1") // a == 1
    }
 }

 func test2() {
    if b := 2; b == 2 {
        fmt.Println("b == 2") // b == 2
    }
    // fmt.Println(b) // 不能直接输出，会报错 undefined:b
 }

 func test3() {
    if c, d := 3,4; c < d {
        fmt.Println("c < d") // c < d
    }
 }

 func test4() {
    e := 10
    if e:=5; e < 10 {
        fmt.Println(e) // 5
    }
    fmt.Println(e) // 10
 }

 func test5() {
    if e:=5; e > 10 {
        fmt.Println("e > 10")
    } else {
        fmt.Println("e <= 10") // e <= 10
    }
 }
```
---
### __循环语句for__

- Go 只有for一个循环语句关键字，但支持3种形式
- 初始化和步进表达式可以是多个值
- 条件语句每次循环都会被重新检查，因此不建议在条件语句中使用函数，尽量提前计算好条件以变量或常量代替
- 左大括号必须和条件语句在同一行
- 示例代码：
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
 }

 func test1() {
    a := 1
    for {
        a ++
        if a > 2 {
            break
        }
        fmt.Println(a)
    }
    fmt.Println("for-over1")
    /*
    // 最终输出
    2
    for-over1
    */
 }

 func test2() {
    b := 1
    for b <= 2{
        b ++
        fmt.Println(b)
    }
    fmt.Println("for-over2")
    /*
    // 最终输出
    2
    3
    for-over2
    */
 }

 func test3() {
    c := 1
    for i:=1; i<=2; i++{
        c ++
        fmt.Println(c)
    }
    fmt.Println("for-over3")
    /*
    // 最终输出
    2
    3
    for-over3
    */
 }

 func test4() {
    d := "string"
    e := 1
    l := len(d) // 不要把长度计算放置于for中,每次都会重新检查,影响性能
    for i:=1; i<=l; i++{
        e ++
        fmt.Println(e)
    }
    fmt.Println("for-over4")
    /*
    // 最终输出
    2
    3
    4
    5
    6
    7
    for-over4
    */
 }

```

---

### __switch语句__

- 可以使用任何类型或表达式作为条件语句
- 无需break, 一旦调价你符合自动终止
- 如果希望执行下一个case,需使用fallthrough语句
- 支持一个初始化表达式，可以是并行方式，右侧需跟分号
- 左大括号必须和条件语句在同一行。
- 示例代码：
```
 package main

 import (
    "fmt"
 )

 func main() {
    // test1()
    // test2()
    test3()
    // test4()
 }

 // 无需break，匹配到则停止，不会存在穿透问题
 func test1 () {
    a := 1
    switch a {
        case 0:
            fmt.Println("a=0")
        case 1:
            fmt.Println("a=1") // a=1
        default:
            fmt.Println("None")
    }
 }

 // switch 中没有表达式，表达式可放在case中
 func test2() {
    a := 1
    switch {
        case a >= 0:
            fmt.Println("a>=0") // a>=0 // 如果不添加fallthrough 只会匹配到此处，即使下面的匹配也不会往下执行
            fallthrough // 添加fallthrough， 则会继续匹配下一个case
        case a >= 1:
            fmt.Println("a>=1") // a>=1
        default:
            fmt.Println("None")

    }
 }

 // 在switch 中声明，变量范围只在switch语句块中
 func test3() {
    switch a := 1; {
        case a >= 0:
            fmt.Println("a>=0") // a>=0
        case a >= 1:
            fmt.Println("a>=1")
        default:
            fmt.Println("None")
    }
 }
```

---

### 跳转语句
- 跳转语句goto, break, continue
- 三个语法都可配合标签使用
- 标签名区分大小写，若不使用会造成编译错误
- Break 与 continue 配合标签可用于多层循环的跳出
- Goto 是调整执行位置，与其它两个语句配合标签的结果并不相同
- 实例代码：
```
  package main

  import (
    "fmt"
  )

  func main() {
    test1()
    test2()
    test3()
  }

  // 突破无限循环
  func test1 () {
    LABEL1:
        for{
            for i:=0; i< 10; i++ {
                if i>3 {

                    break LABEL1 // 单纯的break无法突破无限循环，加上LABEL1外层标签标识则可跳出
                }
                fmt.Println("i:",i)
            }
        }
        fmt.Println("over")

    /*
    // 运行结果：
    i:0
    i:1
    i:2
    i:3
    over
    */
  }

  // 测试goto语句
  func test2() {
    for {
        for i:=0; i<10; i++ {
            if i>3 {
                goto LABEL2
            }
            fmt.Println("i:",i)
        }
    }
    LABEL2:
        fmt.Println("over")

    /*
    i: 0
    i: 1
    i: 2
    i: 3
    over
    */
  }

  // continue
  func test3() {
    LABEL1:
        for i:=0; i<3; i++ {
            // 无限循环
            for {
                fmt.Println("i:",i)
                continue LABEL1 // 立马跳出无限循环
                // fmt.Println(i) // 此处continue后的代码永远不会被执行
            }

        }
        fmt.Println("over")

    /*
    i: 0
    i: 1
    i: 2
    over
    */
  }
```