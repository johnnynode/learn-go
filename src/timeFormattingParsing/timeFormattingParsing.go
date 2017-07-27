package main

import (
	"fmt"
	"time"
)

func main() {
	fpln := fmt.Println
	fp := fmt.Printf

	t := time.Now()
	fpln(t.Format(time.RFC3339)) // 2017-07-27T14:35:19+08:00

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	fpln(t1) // 2012-11-01 22:08:41 +0000 +0000

	fpln(t.Format("3:04PM")) // 2:35PM
	fpln(t.Format("Mon Jan _2 15:04:05 2006")) // Thu Jul 27 14:35:19 2017
	fpln(t.Format("2006-01-02T15:04:05.999999-07:00")) // 2017-07-27T14:35:19.383209+08:00
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	fpln(t2) // 0000-01-01 20:41:00 +0000 UTC

	fp("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()) // 2017-07-27T14:35:19-00:00

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	fpln(e) // parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
}