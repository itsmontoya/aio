// +build linux

package aio

import (
	"log"
	"syscall"
)

func ulimitNoFile() int {
	var rl syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rl)
	if err != nil {
		log.Panic(err)
	}

	return int(rl.Cur)
}
