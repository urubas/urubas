package jit

import (
	"github.com/urubas/urubas/runtime"
	"github.com/urubas/urubas/compiler"
)

type Function struct {
	engine *Engine
	fn *compiler.Function
}

func NewFunction(engine *Engine, fn *compiler.Function) *Function {
	return &Function{engine, fn}
}

func (f *Function) Call(n *runtime.Node) {
}
