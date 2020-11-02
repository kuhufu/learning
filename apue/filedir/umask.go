//+build linux

package main

import (
	"fmt"
	"syscall"
)

const RWRWRW = syscall.S_IRUSR | syscall.S_IWUSR | syscall.S_IRGRP | syscall.S_IWGRP | syscall.S_IROTH | syscall.S_IWOTH

//umask用来屏蔽文件权限
func main() {
	//不屏蔽任何权限
	syscall.Umask(0)
	_, err := syscall.Creat("foo", RWRWRW)
	if err != nil {
		fmt.Println(err)
		return
	}

	//屏蔽 group 和 other 的读写权限
	syscall.Umask(syscall.S_IRGRP | syscall.S_IWGRP | syscall.S_IROTH | syscall.S_IWOTH)
	_, err = syscall.Creat("bar", RWRWRW)
	if err != nil {
		fmt.Println(err)
		return
	}
}
