package vm

import (
	"github.com/urubas/urubas/ast"
	"github.com/urubas/urubas/compiler"
)

type OutputMap map[string]Output

type Edge struct {
	Target *Node
	Slot int
}

type Output struct {
	Name string
	Type *compiler.Type
	Edges []Edge
}

type Input struct {
	Type *compiler.Type
}

type Node struct {
	ID int

	Function compiler.Function

	Inputs []Input
	Outputs OutputMap
}
