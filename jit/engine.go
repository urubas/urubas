package jit

import (
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/compiler"
	"github.com/urubas/urubas/runtime"
)

type Engine struct {
	engine llvm.ExecutionEngine
	program *compiler.Program
}

func New(program *compiler.Program) *Engine {
	engine, _ := llvm.NewMCJITCompiler(program.Module, llvm.MCJITCompilerOptions{})

	return &Engine{
		engine: engine,
		program: program,
	}
}

func (e *Engine) Run(f *compiler.Function) (*runtime.Process, *InputHandlerProxy) {
	p := runtime.NewProcess(0)
	poutput := &InputHandlerProxy{&runtime.NullInputHandler{}}

	for id, node := range f.Nodes {
		fn := e.createFunction(node.Function)
		n := runtime.NewNode(id, fn)

		n.Inputs = make(runtime.InputMap, node.InputCount() + node.StateInputCount())

		for i := range n.Inputs {
			n.Inputs[i] = runtime.NewInput()
		}

		p.AddNode(n)
	}

	for id, node := range p.Nodes {
		original := f.Nodes[id]

		for name, output := range original.Outputs() {
			for _, e := range output.Edges {
				var target runtime.InputHandler

				if e.Target == -1 {
					target = poutput
				} else {
					target = p.Nodes[e.Target]
				}

				node.Connect(name, target, e.Slot)
			}
		}
	}

	for i, arg := range f.Ast.Arguments {
		for _, e := range arg.Edges {
			p.Inputs[i].Edges = append(p.Inputs[i].Edges, &runtime.Edge{
				Target: p.Nodes[e.Target],
				Slot: e.Slot,
			})
		}
	}

	return p, poutput
}

func (e *Engine) createFunction(f interface{}) runtime.Function {
	return nil
}
