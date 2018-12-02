# cgo-rust-ffi

Various experiments and demos relating to Golang talking to Rust via CGO. This
gave me a good excuse to mess around with SHM_GET, SHM_OPEN, MMAP, and mkfifo.

## Examples

1. (DONE) Go calls Rust through CGO and streams bytes using FIFO file (named pipe)
1. (DONE) Go calls Rust through CGO, sharing memory (uses SHM_OPEN + MMAP)
1. (WIP) Go allocates slice in C heap, copies from Go heap and passes pointers to Rust
1. (WIP) Go calls Rust through CGO and streams bytes over TCP socket

## Tools

1. Install Rust: `curl https://sh.rustup.rs -sSf | sh`
1. Install Golang: https://golang.org/dl/

## Caveats

1. This software has been tested with OSX only

## Build and Run Tests

To build and run demos, run `make`.
