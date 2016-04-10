package vm

import (
	"github.com/urubas/urubas/ast"
	"github.com/urubas/urubas/compiler"
)

type NodeMap map[int]*Node

type Context struct {
	Nodes NodeMap
}


