package main

import "basic/time"
import "fmt"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i ++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)

	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/*

$ go run rateLimiting.go
request 1 2017-07-26 11:11:16.697235909 +0800 CST
request 2 2017-07-26 11:11:16.900162235 +0800 CST
request 3 2017-07-26 11:11:17.099213817 +0800 CST
request 4 2017-07-26 11:11:17.29653628 +0800 CST
request 5 2017-07-26 11:11:17.495636883 +0800 CST
request 1 2017-07-26 11:11:17.495691543 +0800 CST
request 2 2017-07-26 11:11:17.49569721 +0800 CST
request 3 2017-07-26 11:11:17.495700845 +0800 CST
request 4 2017-07-26 11:11:17.700176752 +0800 CST
request 5 2017-07-26 11:11:17.89582969 +0800 CST

*/