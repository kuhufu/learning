//+build linux

package main

import (
	"fmt"
	"os"
	"syscall"
)

const (
	R_OK = 4 /* Test for read permission.  */
	W_OK = 2 /* Test for write permission.  */
	X_OK = 1 /* Test for execute permission.  */
	F_OK = 0 /* Test for existence.  */
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: perm <pathname>")
		return
	}

	path := os.Args[1]

	if err := syscall.Access(path, syscall.F_OK); err != nil {
		fmt.Println("access error:", err)
		return
	}

	if err := syscall.Access(path, X_OK); err != nil {
		fmt.Println("access error:", err)
		return
	}
}
