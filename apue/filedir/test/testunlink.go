package main

import "syscall"

func main() {
	path := "/dev/fd/1"

	err := syscall.Unlink(path)
	if err != nil {
		panic(err)
	}

	_, err = syscall.Creat(path, 0777)
	if err != nil {
		panic(err)
	}
}
