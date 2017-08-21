package main

import (
	"fmt"
	"basic/time"
)

var fpln = fmt.Println

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fpln(now) // 2017-07-27 14:17:25.748590827 +0800 CST

	millis := nanos / 1000000
	fpln(secs)   // 1501136245
	fpln(millis) // 1501136245748
	fpln(nanos)  // 1501136245748590827

	fpln(time.Unix(secs, 0))  // 2017-07-27 14:17:25 +0800 CST
	fpln(time.Unix(0, nanos)) // 2017-07-27 14:17:25.748590827 +0800 CST
}
