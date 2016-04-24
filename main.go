package main

import (
	"fmt"
	"github.com/urubas/urubas/ast"
	"github.com/urubas/urubas/jit"
	"github.com/urubas/urubas/compiler"
	"github.com/urubas/urubas/runtime"
)

func buildProgram() *ast.Program {
	main := ast.NewFunction("main")

	main.OutputType = "u.int32"
	main.Arguments = []ast.ArgumentDefinition{
		ast.ArgumentDefinition{
			Type: "u.int32",
			Edges: ast.SingleEdge(1, 0),
		},
		ast.ArgumentDefinition{
			Type: "u.int32",
			Edges: ast.SingleEdge(1, 1),
		},
	}

	main.AddNode(&ast.Node{
		ID: 1,
		Function: "u.add",
		Outputs: ast.SingleEdgeOutput("", 2, 0),
	})

	main.AddNode(&ast.Node{
		ID: 2,
		Function: "u.ret",
		Outputs: ast.SingleEdgeOutput("", -1, 0),
	})

	p := ast.NewProgram("test", "main")
	p.AddFunction(main)

	return p
}

func main() {
	driver := compiler.NewDriver(buildProgram())
	program, err := driver.Compile()

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	engine := jit.New(program)
	fn, ok := program.Functions.Find("main", []*compiler.Type{driver.Intrinsics.Types.Int32, driver.Intrinsics.Types.Int32})

	if !ok {
		fmt.Printf("Main not found!\n")
		return
	}

	ch := make(chan runtime.ChannelInputValue)
	output := &runtime.ChannelInputHandler{ch, nil}
	p, proxy := engine.Run(fn.Fn().(*compiler.Function))

	proxy.Replace(output)

	p.Inputs[0].Finish(runtime.Value(uintptr(1)))
	p.Inputs[1].Finish(runtime.Value(uintptr(2)))

	value := <-ch

	fmt.Printf("%d %v\n", value.Slot, value.Value)
}
