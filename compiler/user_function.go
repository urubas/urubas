package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/ast"
)

type BlockMap map[int]*Block

type Block struct {

	// LLVM
	fnType llvm.Type
	function llvm.Value
}

type UserFunction struct {
	ast *ast.Function

	// IO
	inputTypes []*Type
	outputTypes []*Type

	// JIT
	blocks BlockMap
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
	if bc.Node.ExecutionModel == ast.InlineExecution {
		return bc.Builder.CreateCall(f.blocks[0].function, args, "")
	} else {
		return bc.EmitDispatch(f, args)
	}
}
