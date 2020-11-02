//+build linux

package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	_, err := syscall.Open("tempfile", syscall.O_RDWR|syscall.O_CREAT, 0)
	if err != nil {
		fmt.Println("open error:", err)
		return
	}

	if err := syscall.Unlink("tempfile"); err != nil {
		fmt.Println("unlink error:", err)
		return
	}

	fmt.Println("unlink success")
	time.Sleep(time.Second * 30)
	fmt.Println("done")
}
