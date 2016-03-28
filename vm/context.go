package vm

import (
	"github.com/urubas/urubas/graph"
)

type CallingContext struct {
	Entrypoint int
	ValueCount int
	StateCount int
}

type Context interface {
	Program() *graph.Program
	ResultChannel() chan interface{}
	Execute([]<-chan interface{})
	GetNode(node int) Node
	CreateSubcontext(main CallingContext, result chan interface{}) Context
}
