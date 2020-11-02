package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ftw <pathname>")
		return
	}

	path := os.Args[1]

	fd, err := syscall.Open(path, syscall.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("open error:", path, err)
		return
	}

	buf := make([]byte, 4096)
	n, err := syscall.Read(fd, buf)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	fmt.Println(string(buf[:n]))
}
