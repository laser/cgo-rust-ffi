package sharedpipe

import (
	"os"
	"syscall"
)

func NewFifo(name string, pipeFileFlags int, pipeFileMode os.FileMode) (*os.File, error) {
	err := syscall.Mkfifo(name, uint32(pipeFileMode))
	if err != nil {
		return nil, err
	}

	return os.OpenFile(name, pipeFileFlags, pipeFileMode)
}
