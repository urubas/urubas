all: urubas

clean:
	rm urubas

urubas:
	go build -ldflags "-extldflags '-rpath ${GOPATH}/src/llvm.org/llvm/bindings/go/llvm/workdir/llvm_build/lib'"

.PHONY: all clean urubas

