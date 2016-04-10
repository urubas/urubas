package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/ast"
)

type Compiler struct {
	Program *Program
	Module *llvm.Module
}

type BuildContext struct {
	Compiler *Compiler

	Builder *llvm.Builder
	Node *ast.Node
}

