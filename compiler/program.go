package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/ast"
)

type Program struct {
	Ast *ast.Program

	Module llvm.Module
	Functions FunctionMap

	contextMarkerCounter int
}

func NewProgram(ast *ast.Program, fns FunctionMap) *Program {
	return &Program{ast, llvm.NewModule(ast.Name), fns, 0}
}

func (p *Program) nextContextMarker() int {
	p.contextMarkerCounter++

	return p.contextMarkerCounter
}
