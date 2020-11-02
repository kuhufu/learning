//+build linux

package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	names := os.Args[1:]
	for _, name := range names {
		lstat(name)
	}
}

//lstat 函数类似于 stat，但是当命名的文件是一个符号链接时，
//lstat 返回该符号链接的有关信息，而不是由该符号链接引用的文件的信息。
func lstat(name string) {
	stat := new(syscall.Stat_t)
	err := syscall.Lstat(name, stat)
	if err != nil {
		panic(err)
	}

	mode := stat.Mode
	var str string
	switch {
	case S_ISREG(mode):
		str = "regular"
	case S_ISDIR(mode):
		str = "directory"
	case S_ISBLK(mode):
		str = "block special"
	case S_ISCHR(mode):
		str = "character special"
	case S_ISLNK(mode):
		str = "symbol link"
	case S_ISFIFO(mode):
		str = "fifo"
	case S_ISSOCK(mode):
		str = "socket"
	default:
		str = "unknown file mode"
	}

	fmt.Printf("%v: %v\n", name, str)
	fmt.Printf("size:%v, blockSize:%v, blocks:%v\n", stat.Size, stat.Blksize, stat.Blocks)

	info, _ := os.Stat(name)
	fmt.Println(info.Size())
}

func S_ISREG(mode uint32) bool {
	return mode&syscall.S_IFMT == syscall.S_IFREG
}

func S_ISDIR(mode uint32) bool {
	return mode&syscall.S_IFMT == syscall.S_IFDIR
}

func S_ISCHR(mode uint32) bool {
	return mode&syscall.S_IFMT == syscall.S_IFCHR
}

func S_ISBLK(mode uint32) bool {
	return mode&syscall.S_IFMT == syscall.S_IFBLK
}

func S_ISLNK(mode uint32) bool {
	return mode&syscall.S_IFMT == syscall.S_IFLNK
}

func S_ISFIFO(mode uint32) bool {
	return mode&syscall.S_IFMT == syscall.S_IFIFO
}

func S_ISSOCK(mode uint32) bool {
	return mode&syscall.S_IFMT == syscall.S_IFSOCK
}
