package main

import (
	"fmt"
	"syscall"
)

func main() {
	if err := syscall.Symlink("/no/exist/path", "my_symlink"); err != nil {
		fmt.Println("create symbolic link error:", err)
		return
	}

	defer func() {
		if err := syscall.Unlink("my_symlink"); err != nil {
			fmt.Println("unlink error:", err)
			return
		}
	}()

	buf := make([]byte, syscall.NAME_MAX)
	//n是符号链接指向的文件的文件路径长度
	n, err := syscall.Readlink("my_symlink", buf)
	if err != nil {
		fmt.Println("readlink error:", err)
		return
	}
	fmt.Println(n)
	fmt.Println(string(buf[:n]))
}
