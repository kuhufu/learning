package main

import "syscall"

func main() {
	path := "SqQndFWzHVIIcaTwoJTrtwhfcBjMRjiZyYKrllFlcmtiUNBJbgLqYgXPAiIoLscVmGZYFyWQWwjneRiHkJWIlzNUSkNhWHYmekqvwZcTMBWRKNlnSACVnDAxTBFoZaDi"
	for i := 0; i < 36; i++ {
		err := syscall.Mkdir(path, 0777)
		if err != nil {
			panic(err)
		}

		err = syscall.Chdir(path)
		if err != nil {
			panic(err)
		}
	}

	_, err := syscall.Getwd()
	if err != nil {
		panic(err)
	}

}
