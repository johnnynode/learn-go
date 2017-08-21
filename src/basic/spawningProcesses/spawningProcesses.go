package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()

	grepIn.Write([] byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

/*

$ go run spawningProcesses.go
> date
2017年 7月28日 星期五 11时20分11秒 CST

> grep hello
hello grep

> ls -a -l -h
total 8
drwxr-xr-x   3 gddev4  staff   102B  7 28 11:10 .
drwxr-xr-x  65 gddev4  staff   2.2K  7 28 11:02 ..
-rw-r--r--   1 gddev4  staff   718B  7 28 11:10 spawningProcesses.go

*/