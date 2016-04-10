package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/ast"
)

type UserFunction struct {
	ast *ast.Function

	// IO
	inputTypes []*Type
	outputTypes []*Type

	// LLVM
	fnType llvm.Type
	function llvm.Value
}

func (f *UserFunction) Name() string {
	return f.ast.Name
}

func (f *UserFunction) InputTypes() []*Type {
	return f.inputTypes
}

func (f *UserFunction) OutputTypes() []*Type {
	return f.outputTypes
}

func (f *UserFunction) Emit(bc *BuildContext, args []llvm.Value) llvm.Value {
	return bc.Builder.CreateCall(f.function, args, "")
}
