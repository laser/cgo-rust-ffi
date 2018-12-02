package main

import (
	"fmt"

	"github.com/laser/cgo-rust-ffi/demos"
)

func main() {
	demos.RunNamedPipeDemo()
	fmt.Printf("\n\n\n")
	demos.RunSharedMemDemo()
}
