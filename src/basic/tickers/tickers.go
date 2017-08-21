package main

import "basic/time"
import "fmt"

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

/*

$ go run tickers.go
Tick at 2017-07-26 10:49:54.477527436 +0800 CST
Tick at 2017-07-26 10:49:54.977305838 +0800 CST
Tick at 2017-07-26 10:49:55.47756701 +0800 CST
Ticker stopped

*/