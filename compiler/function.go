package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Emitable interface {
	Emit(bc *BuildContext, args []llvm.Value) llvm.Value
}
