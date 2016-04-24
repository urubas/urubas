package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Frame struct {
	FunctionType llvm.Type
	Function llvm.Value

	OutputType *Type
	OutputMap map[int]string

	Builder llvm.Builder
}
