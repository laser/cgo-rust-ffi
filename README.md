# cgo-rust-ffi

Various experiments and demos relating to Golang talking to Rust via CGO

## Tools

1. Install Rust: `curl https://sh.rustup.rs -sSf | sh`
1. Install Golang: https://golang.org/dl/

## Build and Run Tests

To build and run tests, run `make`. You should see output similar to:

```shell
12:23 $ make
pushd ./hasher && cargo build --release --all && cbindgen > libstream_reader.h && popd
~/go/src/github.com/laser/cgo-rust-ffi/hasher ~/go/src/github.com/laser/cgo-rust-ffi
   Compiling hasher v0.1.0 (/Users/erinswenson-healey/go/src/github.com/laser/cgo-rust-ffi/hasher)
    Finished release [optimized] target(s) in 0.61s
~/go/src/github.com/laser/cgo-rust-ffi
go run ./main.go
[golang] ***********************
[golang] running named pipe demo
[golang] ***********************
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[golang] wrote n=80000 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[rust]   read n=8192 bytes
[golang] wrote n=80000 bytes
[golang] wrote everything
[rust]   read n=4352 bytes
[rust]   EOF
[golang] digest=e0902e35221a2d8d003059ffb460ffa0
```
