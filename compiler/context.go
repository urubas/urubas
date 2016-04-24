package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type BuildContext interface {
	Driver() *Driver
	Builder() llvm.Builder
	Intrinsics() *Intrinsics
}
