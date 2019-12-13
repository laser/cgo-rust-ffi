package demos

import (
	"fmt"
	"math/rand"

	"github.com/laser/cgo-rust-ffi/hasher/cgo"
	"github.com/laser/cgo-rust-ffi/sharedmem"
)

////////////////////////////////////////////////////////////////////////////////
// Sharing memory with Rust through CGO using shmopen and mmap
//////////////////////////////////////////////////////////////

func RunSharedMemDemo() {
	fmt.Println("[golang] **************************")
	fmt.Println("[golang] running shared memory demo")
	fmt.Println("[golang] **************************")

	name := fmt.Sprintf("/tmp/foo%d", rand.Int())
	size := 160000

	r1, err := sharedmem.NewProducerRegion(name, size)
	if err != nil {
		panic(err)
	}
	defer r1.Close()

	// write into now-shared memory
	bs := randomBytes(size)
	copy(r1.Mem[0:len(bs)], bs)
	fmt.Println("[golang] wrote everything")

	// call in to Rust, passing metadata used to reconstruct memory map
	digest := cgo.ChecksumSharedmem(name, uint(size))
	fmt.Printf("[golang] digest=%s\n", digest)

	fmt.Println("[golang] **************************")
	fmt.Println("[golang] end shared memory demo    ")
	fmt.Println("[golang] **************************")
}
