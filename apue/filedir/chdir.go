//+build linux

package main

import (
	"fmt"
	"syscall"
)

func main() {
	fd, err := syscall.Open(".", syscall.O_DIRECTORY, 0777)
	if err != nil {
		fmt.Println("open error:", err)
		return
	}

	if err := syscall.Chdir("/tmp"); err != nil {
		fmt.Println("chdir error", err)
		return
	}
	fmt.Println("chdir to /tmp succeeded")

	cwd, err := syscall.Getwd()
	if err != nil {
		fmt.Println("getcwd error:", err)
		return
	}
	fmt.Println(cwd)

	if err := syscall.Fchdir(fd); err != nil {
		fmt.Println("fchdir error:", err)
		return
	}

	cwd, err = syscall.Getwd()
	if err != nil {
		fmt.Println("getcwd error:", err)
		return
	}
	fmt.Println(cwd)
}
