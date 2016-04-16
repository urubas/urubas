package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/ast"
)

type Type struct {
	name string

	Ast *ast.Type

	// LLVM
	Type llvm.Type
}
