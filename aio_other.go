// +build !linux

package aio

func ulimitNoFile() int { return 1024 }
