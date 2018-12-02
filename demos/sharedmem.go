// +build linux, cgo darwin,cgo

package demos

/*
#cgo LDFLAGS: -L${SRCDIR}/../hasher/target/release -lhasher
#include "../hasher/libhasher.h"
*/
import "C"
import (
	"fmt"

	"github.com/laser/cgo-rust-ffi/sharedmem"
)

////////////////////////////////////////////////////////////////////////////////
// Sharing memory with Rust through CGO using shmopen and mmap
//////////////////////////////////////////////////////////////

func RunSharedMemDemo() {
	fmt.Println("[golang] **************************")
	fmt.Println("[golang] running shared memory demo")
	fmt.Println("[golang] **************************")

	name := "/foobarbaz"
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
	digest := C.GoString(C.checksum_sharedmem(C.CString(name), C.size_t(size)))
	fmt.Printf("[golang] digest=%s\n", digest)

	fmt.Println("[golang] **************************")
	fmt.Println("[golang] end shared memory demo    ")
	fmt.Println("[golang] **************************")
}
