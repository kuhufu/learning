package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: mkdir <path>")
		return
	}

	for _, path := range os.Args[1:] {
		err := syscall.Mkdir(path, 0777)
		if err != nil {
			fmt.Println("mkdir error:", err)
			continue
		}
	}



	time.Sleep(time.Second * 10)
}
