//+build linux

package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	var stat syscall.Stat_t
	var times [2]syscall.Timeval

	for _, path := range os.Args[1:] {
		if err := syscall.Stat(path, &stat); err != nil {
			fmt.Println("stat error:", err)
			return
		}

		fd, err := syscall.Open(path, syscall.O_RDWR|syscall.O_TRUNC, 0666)
		if err != nil {
			fmt.Println("open error:", err)
			return
		}

		//保持 access time 访问时间不变
		times[0].Sec = stat.Atim.Sec
		times[0].Usec = stat.Atim.Nsec / 1000

		//保持 modify time 修改时间不变
		times[1].Sec = stat.Mtim.Sec
		times[1].Usec = stat.Mtim.Nsec / 1000

		if err = syscall.Futimes(fd, times[:]); err != nil {
			fmt.Println("futimes error:", err)
			return
		}
	}
}
