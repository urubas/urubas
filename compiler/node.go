package compiler

import (
	"github.com/urubas/urubas/ast"
)

type NodeMap map[int]Node

type Node interface {
	ID() int
	Ast() *ast.Node
	Function() *FunctionHolder

	InputCount() int
	StateInputCount() int
	Outputs() ast.OutputMap
}
