package demos

import (
	"fmt"
	"os"
	"sync"

	"github.com/laser/cgo-rust-ffi/hasher/cgo"
	"github.com/laser/cgo-rust-ffi/sharedpipe"
)

////////////////////////////////////////////////////////////////////////////////
// Streaming SHA256 checksum through CGO using named pipe
/////////////////////////////////////////////////////////

func RunNamedPipeDemo() {
	fmt.Println("[golang] **************************")
	fmt.Println("[golang] running named pipe demo   ")
	fmt.Println("[golang] **************************")

	name := "/tmp/pipey"

	_ = os.Remove(name) // don't blow up if file doesn't exist

	f, err := sharedpipe.NewFifo(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	// schedule goroutine which makes call into Rust lib to create digest
	go func() {
		digest := cgo.ChecksumFile(name)
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

	fmt.Println("[golang] **************************")
	fmt.Println("[golang] end named pipe demo       ")
	fmt.Println("[golang] **************************")
}
