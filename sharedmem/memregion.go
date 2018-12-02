package sharedmem

import (
	"os"
	"syscall"
	"unsafe"
)

type MemRegion struct {
	name       string
	Mem        []byte
	fd         uintptr
	size       int
	isProducer bool
}

func NewConsumerRegion(name string, size int) (MemRegion, error) {
	return newRegion(name, size, os.O_RDONLY, 0666, syscall.PROT_READ, false)
}

func NewProducerRegion(name string, size int) (MemRegion, error) {
	r, err := newRegion(name, size, os.O_CREATE|os.O_RDWR, 0666, syscall.PROT_READ|syscall.PROT_WRITE, true)
	r.isProducer = true

	return r, err
}

func (r MemRegion) Close() error {
	err := syscall.Munmap(r.Mem)
	if err != nil {
		return err
	}

	err = os.NewFile(r.fd, "").Close()
	if err != nil {
		return err
	}

	if r.isProducer {
		key, err := syscall.BytePtrFromString(r.name)
		if err != nil {
			return err
		}

		_, _, errno := syscall.RawSyscall(syscall.SYS_SHM_UNLINK, uintptr(unsafe.Pointer(key)), uintptr(0), uintptr(0))
		if errno != 0 {
			return os.NewSyscallError("SYS_SHM_UNLINK", errno)
		}
	}

	return nil
}

func newRegion(name string, size int, shmFlags int, shmPerms, mmapPerms int, truncate bool) (MemRegion, error) {
	key, err := syscall.BytePtrFromString(name)
	if err != nil {
		return MemRegion{}, err
	}

	fd, _, errno := syscall.RawSyscall(syscall.SYS_SHM_OPEN, uintptr(unsafe.Pointer(key)), uintptr(shmFlags), uintptr(shmPerms))
	if errno != 0 {
		return MemRegion{}, os.NewSyscallError("SYS_SHM_OPEN", errno)
	}

	if truncate {
		err = syscall.Ftruncate(int(fd), int64(size))
		if err != nil {
			return MemRegion{}, err
		}
	}

	mem, err := syscall.Mmap(int(fd), 0, size, mmapPerms, syscall.MAP_SHARED)
	if err != nil {
		return MemRegion{}, err
	}

	return MemRegion{
		name: name,
		Mem:  mem,
		fd:   fd,
		size: size,
	}, nil
}
