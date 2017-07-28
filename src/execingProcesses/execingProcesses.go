package main

import (
	"syscall"
	"os"
	"os/exec"
)

func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

/*

$ go run execingProcesses.go
total 8
drwxr-xr-x   3 gddev4  staff   102B  7 28 11:46 .
drwxr-xr-x  66 gddev4  staff   2.2K  7 28 11:42 ..
-rw-r--r--   1 gddev4  staff   300B  7 28 11:46 execingProcesses.go

*/