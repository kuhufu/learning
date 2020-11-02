//+build linux

package main

import (
	"fmt"
	"syscall"
)

func main() {
	stat := new(syscall.Stat_t)

	if err := syscall.Stat("foo", stat); err != nil {
		fmt.Println(err)
		return
	}

	if err := syscall.Chmod("foo", (stat.Mode & ^uint32(syscall.S_IXGRP))|syscall.S_ISGID); err != nil {
		fmt.Println(err)
		return
	}

	if err := syscall.Chmod("bar", syscall.S_IRUSR|syscall.S_IWUSR|syscall.S_IRGRP|syscall.S_IROTH); err != nil {
		fmt.Println(err)
		return
	}
}
