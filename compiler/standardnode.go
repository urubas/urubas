package compiler

import (
	"github.com/urubas/urubas/ast"
)

type StandardNode struct {
	id int
	ast *ast.Node
	fn *FunctionHolder
	outputs ast.OutputMap
}

func NewStandardNode(ast *ast.Node, fn *FunctionHolder) *StandardNode {
	return &StandardNode{
		id: ast.ID,
		ast: ast,
		fn: fn,
		outputs: ast.Outputs,
	}
}

func (s *StandardNode) ID() int {
	return s.id
}

func (s *StandardNode) Ast() *ast.Node {
	return s.ast
}

func (s *StandardNode) InputCount() int {
	return len(s.fn.InputTypes())
}

func (s *StandardNode) StateInputCount() int {
	return 0
}

func (s *StandardNode) Function() *FunctionHolder {
	return s.fn
}

func (s *StandardNode) Outputs() ast.OutputMap {
	return s.outputs
}
