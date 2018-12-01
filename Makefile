.PHONY: all build run

all: build run

build:
		pushd ./hasher && cargo build --release --all && cbindgen > libstream_reader.h && popd

run:
		go run ./main.go
