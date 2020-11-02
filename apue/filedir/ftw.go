package main

import (
	"fmt"
	"os"
	"syscall"
)

var (
	nreg   int //普通文件
	ndir   int //目录文件
	nblk   int //块特殊文件
	nchr   int //字符特殊文件
	nfifo  int //管道文件
	nslink int //符号链接文件
	nsock  int //套接字
	ntot   int //总文件数
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ftw <pathname>")
		return
	}

	stat := new(syscall.Stat_t)
	err := syscall.Lstat(os.Args[1], stat)
	if err != nil {
		fmt.Println("lstat error:", os.Args[1], err)
		return
	}

	if stat.Mode&syscall.S_IFMT != syscall.S_IFDIR {
		fmt.Println("path must be directory")
		return
	}

	myFtw(os.Args[1])

	ndir-- //减一，不计算当前目录

	ntot = nreg + ndir + nblk + nchr + nslink + nsock

	fmt.Println("ntot", ntot)
	fmt.Println("nreg", nreg)
	fmt.Println("ndir", ndir)
	fmt.Println("nblk", nblk)
	fmt.Println("nchr", nchr)
	fmt.Println("nslink", nslink)
	fmt.Println("nsock", nsock)
}

func myFtw(path string) {
	stat := new(syscall.Stat_t)
	err := syscall.Lstat(path, stat)
	if err != nil {
		fmt.Println("lstat error:", path, err)
		return
	}

	myFunc(stat)

	if stat.Mode&syscall.S_IFMT != syscall.S_IFDIR { //不是目录直接返回
		return
	}

	dirfd, err := syscall.Open(path, syscall.O_RDONLY|syscall.O_DIRECTORY, 0444)
	if err != nil {
		fmt.Println("open error:", path, err)
		return
	}

	//读出目录下所有文件名
	buf := make([]byte, 8192)
	_, err = syscall.ReadDirent(dirfd, buf)
	if err != nil {
		fmt.Println("read directory error:", err)
		return
	}

	//需要对 buf 解析转换成文件名数组
	_, _, names := syscall.ParseDirent(buf, 100, nil)

	for _, name := range names {
		myFtw(path + string(os.PathSeparator) + name)
	}
}

func myFunc(stat *syscall.Stat_t) {
	switch stat.Mode & syscall.S_IFMT {
	case syscall.S_IFREG:
		nreg++
	case syscall.S_IFDIR:
		ndir++
	case syscall.S_IFBLK:
		nblk++
	case syscall.S_IFCHR:
		nchr++
	case syscall.S_IFIFO:
		nfifo++
	case syscall.S_IFLNK:
		nslink++
	case syscall.S_IFSOCK:
		nsock++
	}
}
