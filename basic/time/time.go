package main

import (
	"fmt"
	"basic/time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now) // 2017-07-27 13:44:15.661319639 +0800 CST

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	p(then) // 2009-11-17 20:34:58.651387237 +0000 UTC

	p(then.Year())       // 2009
	p(then.Month())      // November
	p(then.Day())        // 17
	p(then.Hour())       // 20
	p(then.Minute())     // 34
	p(then.Second())     // 58
	p(then.Nanosecond()) // 651387237
	p(then.Location())   // UTC

	p(then.Weekday()) // Tuesday

	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	diff := now.Sub(then)
	p(diff) // 67401h9m17.009932402s

	p(diff.Hours())       // 67401.15472498123
	p(diff.Minutes())     // 4.0440692834988735e+06
	p(diff.Seconds())     // 2.426441570099324e+08
	p(diff.Nanoseconds()) // 242644157009932402

	p(then.Add(diff))  // 2017-07-27 05:44:15.661319639 +0000 UTC
	p(then.Add(-diff)) // 2002-03-11 11:25:41.641454835 +0000 UTC
}
