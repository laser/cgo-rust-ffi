DEPS:=hasher.h libhasher.a

all: $(DEPS)
.PHONY: all

$(DEPS): .install-hasher  ;

.install-hasher: rust
	cd rust && cargo build --release --all; cd ..
	find ./rust/target/release -type f -name "hasher.h" -print0 | xargs -0 ls -t | head -n 1 | xargs -I {} cp {} ./cgo/hasher.h
	find ./rust/target/release -type f -name "libhasher.a" -print0 | xargs -0 ls -t | head -n 1 | xargs -I {} cp {} ./cgo/libhasher.a
	c-for-go --ccincl --ccdefs hasher.yml
	@touch $@

clean:
	rm -rf $(DEPS) .install-hasher
	rm -rf cgo/*.go
	rm -rf cgo/*.h
	rm -rf cgo/*.a
.PHONY: clean
