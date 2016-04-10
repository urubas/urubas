package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

func InitializeInfrastructure() {
	llvm.LinkInMCJIT()
	llvm.InitializeNativeTarget()
	llvm.InitializeNativeAsmPrinter()

	initializeTypes()
	initializeArithmetic()
}
