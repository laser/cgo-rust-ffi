// +build linux, cgo darwin,cgo

package main

/*
#cgo LDFLAGS: -L${SRCDIR}/hasher/target/release -lhasher
#include "./hasher/libhasher.h"
*/
import "C"
import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"sync"
	"syscall"
)

func randomBytes(len int) []byte {
	slice := make([]byte, len)

	_, err := io.ReadFull(rand.Reader, slice)
	if err != nil {
		panic(err)
	}

	return slice
}

func main() {
	namedPipeDemo()
}

////////////////////////////////////////////////////////////////////////////////
// Streaming SHA256 checksum through CGO using named pipe
/////////////////////////////////////////////////////////

func namedPipeDemo() {
	fmt.Println("[golang] ***********************")
	fmt.Println("[golang] running named pipe demo")
	fmt.Println("[golang] ***********************")

	pipeFile := "/tmp/pipey"

	err := os.Remove(pipeFile)
	if err != nil {
		panic(err)
	}

	err = syscall.Mkfifo(pipeFile, 0666)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(pipeFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	// schedule goroutine which makes call into Rust lib to create digest
	go func() {
		digest := C.GoString(C.checksum_file(C.CString(pipeFile)))
		fmt.Printf("[golang] digest=%s\n", digest)
		wg.Done()
	}()

	// main goroutine writes to named pipe
	for iter := 0; iter < 2; iter++ {
		m := 80000
		n, err := f.Write(randomBytes(m))
		if err != nil {
			panic(err)
		}
		if n != m {
			panic("wrong length")
		}

		fmt.Printf("[golang] wrote n=%d bytes\n", n)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("[golang] wrote everything")
	wg.Wait()
}
