package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/ast"
)

type BuildContext struct {
	Driver *Driver
	Builder llvm.Builder
	Node *ast.Node
}


