package jit

import (
	"unsafe"
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/runtime"
	"github.com/urubas/urubas/compiler"
)

type Thunk struct {
	engine *Engine
	thunk *compiler.Thunk
}

func NewThunk(engine *Engine, thunk *compiler.Thunk) *Thunk {
	return &Thunk{engine, thunk}
}

func (f *Thunk) Call(n *runtime.Node) {
	n.WaitForArguments()

	args := make([]llvm.GenericValue, len(n.Inputs))

	for i, input := range n.Inputs {
		args[i] = llvm.NewGenericValueFromPointer(unsafe.Pointer(input.Value))
	}

	ret := f.engine.engine.RunFunction(f.thunk.Function, args)

	for _, o := range n.Outputs {
		o.Send(runtime.Value(ret.Pointer()))
		o.Close(true)
	}
}
