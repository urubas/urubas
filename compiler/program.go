package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/ast"
)

type TypeMap map[string]*Type

type Program struct {
	Ast *ast.Program

	Module llvm.Module
	Types TypeMap
	Functions FunctionMap
}
