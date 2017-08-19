### __Concurrency__

- 很多人都是冲着 Go 大肆宣扬的高并发而忍不住跃跃欲试，但其实从源码的解析来看，goroutine 只是由官方实现的超级“线程池”而已。
- 不过话说回来，每个实例 4-5KB 的栈内存占用和由于实现机制而大幅减少的创建和销毁开销，是制造 Go 号称的高并发的根本原因。
- 另外，goroutine 的简单易用，也在语言层面上给予了开发者巨大的便利。
- 并发不是并行：Concurrency Is Not Parallelism
- 并发主要由切换时间片来实现“同时”运行，在并行则是直接利用
- 多核实现多线程的运行，但 Go 可以设置使用核数，以发挥多核计算机的能力。
- Goroutine 奉行通过通信来共享内存，而不是共享内存来通信。

### __Channel__

- Channel 是 goroutine 沟通的桥梁，大都是阻塞同步的
- 通过 make 创建，close 关闭
- Channel 是引用类型
- 可以使用 for range 来迭代不断操作 channel
- 可以设置单向或双向通道
- 可以设置缓存大小，在未被填满前不会发生阻塞

### __Select__

- 可处理一个或多个 channel 的发送与接收
- 同时有多个可用的 channel时按随机顺序处理
- 可用空的 select 来阻塞 main 函数
- 可设置超时

### __示例代码__
```
package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)


func main() {
	// 说明，由于存在异步和阻塞等情况，测试时候需要一个函数一个函数的执行
	// 执行一个函数时，注释掉其他测试函数
	test1T()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7T()
	// test8T()
	// test9T()
	// test10T()
	// test11T()
	// test12()
	// test13()
	// test14()
}

func test1() {
	fmt.Println("Go Go Go !!!")
}

func test1T() {
	fmt.Println("=============test1T=============")
	go test1() // 先启动goroutine
	time.Sleep(2 * time.Second) // 程序等待2s后退出
}

// 使用channel 来进行消息的发送
func test2() {
	fmt.Println("=============test2=============")
	c := make(chan bool) // 创建一个channel
	// 下面是用匿名函数
	go func() {
		fmt.Println("Go2") // Go2
		c <- true // c 存进一个bool型
	}()
	<-c // 在外面取出来 函数在执行到这一步的时候是阻塞状态，等c 存完之后才会执行这里，结束运行
}

// 可以使用 for range 来迭代不断操作 channel，等待channel完成，利用close函数，最后退出程序
func test3() {
	fmt.Println("=============test3=============")
	c := make(chan bool) // 创建一个channel ，make 即能存，又能取 (属于双向通道)
	// 下面是用匿名函数
	go func() {
		fmt.Println("Go2")
		c <- true // c 存进一个bool型
		close(c) // 在迭代时一定要注意close，否则会死锁，最终崩溃退出。
	}()
	for v := range c {
		fmt.Println(v) // true
	}
}

// 测试 使用无缓存的channel
func test4() {
	fmt.Println("=============test4=============")
	c := make(chan bool)
	go func() {
		fmt.Println("Go3")
		c <- true
	}()
	<- c // main 函数先执行，再去执行 goroutine， 此处是一直等待状态, 直到取出东西来，main函数退出
}

// 测试 使用无缓存的channel， 尝试位置调换, 无缓存是同步阻塞状态，直到读取到，才会结束
func test5() {
	fmt.Println("=============test5=============")
	c := make(chan bool)
	go func() {
		fmt.Println("Go3") // Go3
		<- c
	}()
	c <- true
}

// 测试使用有缓存的channel, 并且对调位置 ，有缓存是异步的
func test6() {
	fmt.Println("=============test6=============")
	c := make(chan bool, 1)
	go func() {
		fmt.Println("Go3") // 未输出
		<- c

	}()
	 c <- true // 直接退出
}

func test7() {
	a := 1
	for i := 0; i<10000000; i++ {
		a += i
	}
	fmt.Println(a)
}

// 循环10个 goroutine ，什么都没输出
func test7T() {
	fmt.Println("=============test7T=============")
	for i := 0; i < 10; i++ {
		go test7()
	}
}

func test8(c chan bool, index int) {
	a := 1
	for i := 0; i<100; i++ {
		a += i
	}
	fmt.Println(index, a)
	if index == 9 {
		c <- true
	}
}

// 循环输出10个 goroutine , 但结果并未按照想象中的，10个分别输出执行，可能是go的api版本升级导致的, 默认使用多核处理
// 意想不到的结果是 每个goroutine 都是独立的，
// 如果第9个先输出，那么其他的可能并不会输出而是直接结束
func test8T() {
	fmt.Println("=============test8T=============")
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go test8(c, i)
	}
	<-c
}

// 循环10个 goroutine , 但是每个goroutine 都是独立的，如果第9个先输出，那么其他的可能并不会输出而是直接结束
// 同样使用test8 , 现在的情况是，输出结果和test8T是一样的
func test9T() {
	fmt.Println("=============test9T=============")
	runtime.GOMAXPROCS(runtime.NumCPU()) // 使用多核CPU，分配不定的, 不能够用最后一个goroutine的结果，作为10个都完成的判定
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go test8(c, i)
	}
	<-c
}

func test10(c chan bool, index int) {
	a := 1
	for i := 0; i<100; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}

// 设置一个缓存长度为10的channel, 使用此种方法，10个channel会分别输出执行
func test10T() {
	fmt.Println("=============test10T=============")
	runtime.GOMAXPROCS(runtime.NumCPU()) // 使用多核CPU，分配不定的, 不能够用最后一个goroutine的结果，作为10个都完成的判定
	c := make(chan bool, 10) // 此处channel 长度为10
	for i := 0; i < 10; i++ {
		go test10(c, i)
	}
	// 循环取10次
	for i := 0; i < 10; i ++ {
		<-c
	}
}

// 使用第二种方法来完成执行10个channel
// 通过sync 同步包中的waitgroup 来创建一个任务组，在任务组中添加要完成的任务数
// 在没完成一次任务，就标记一次档，待完成任务数就会往下减，直到为0时，全部任务就会完成。

func test11(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 100; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}

func test11T() {
	fmt.Println("=============test11T=============")
	runtime.GOMAXPROCS(runtime.NumCPU()) // 使用多核CPU，分配不定的, 不能够用最后一个goroutine的结果，作为10个都完成的判定
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go test11(&wg, i)
	}
	wg.Wait()
}

// 使用select 测试多个channel
func test12() {
	fmt.Println("=============test12=============")
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2) // 设置一个缓存为2的channel
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1",v)
			case v, ok := <- c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1)
	close(c2)

	for i := 0; i < 2; i ++ {
		<- o
	}

	/*
	// 程序运行结果：
	c1 1
	c2 hi
	c1 3
	c2 hello

	 */
}

func test13() {
	fmt.Println("=============test13=============")
	c := make(chan int)
	go func() {
		// 不断读出c的值
		for v := range c {
			fmt.Println(v)
		}
	}()

	// 随机向c中写入0或者1
	for i := 0; i < 6; i++ {
		select {
		case c <- 0:
		case c <- 1:
		}
	}

	/*
	// 随机输出结果：
	1
	1
	0
	0
	0

	 */
}

// select 可设置超时
func test14() {
	fmt.Println("=============test14=============")
	c := make(chan bool)
	select {
	case v:= <-c: // 此处未向c中放入数据，所以程序一直在等待
		fmt.Println(v)
	case <- time.After(3 * time.Second): // time.After 返回一个channel，case被执行，输出 Timeout After 3 second
		fmt.Println("Timeout After 3 second")
	}

	/*
	// 程序在3s后的运行结果：
	Timeout After 3 second
	 */

}


```

