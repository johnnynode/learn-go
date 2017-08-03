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
        l := len(d) // 不要把计算长度放置于for中
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



