package compiler

import (
	"errors"
	"github.com/urubas/urubas/ast"
)

type Function struct {
	Ast *ast.Function
	Program *Program

	Nodes NodeMap

	inputTypes []*Type
	outputType *Type

	contextMarker int
	thunkContextMarker int
}

func NewFunction(ast *ast.Function, program *Program) *Function {
	return &Function{
		Ast: ast,
		Program: program,
		Nodes: make(NodeMap),
		contextMarker: program.nextContextMarker(),
		thunkContextMarker: program.nextContextMarker(),
	}
}

func (f *Function) Name() string {
	return f.Ast.Name
}

func (f *Function) InputTypes() []*Type {
	return f.inputTypes
}

func (f *Function) OutputType() *Type {
	return f.outputType
}

func (f *Function) Call(bc BuildContext, args []Value) Value {
	return ConstNull(bc.Intrinsics().Types.Void)
}

func (f *Function) Prepare(d *Driver) error {
	f.inputTypes = argumentsToTypes(d, f.Ast.Arguments)
	f.outputType = d.Types[f.Ast.OutputType]

	return nil
}

func (f *Function) Build(d *Driver) error {
	err := f.prepareNodes(d)

	if err != nil {
		return err
	}

	return nil
}

func (f *Function) prepareNodes(d *Driver) error {
	order := sortNodes(f.Ast.Nodes)

	for i, id := range order {
		node := f.Ast.Nodes[id]
		slots := make(map[int]*Type)

		// Infer input types from function arguments
		for j, a := range f.Ast.Arguments {
			for _, e := range a.Edges {
				if e.Target == id {
					slots[e.Slot] = f.inputTypes[j]
				}
			}
		}

		// Infer input types from node outputs
		// Note that any node that outputs to this one
		// Is necessarily behind it on the sorted node list
		for j := 0; j < i; j++ {
			n := f.Nodes[order[j]]

			for _, o := range n.Ast().Outputs {
				for _, e := range o.Edges {
					if e.Target == id {
						slots[e.Slot] = n.Function().OutputType()
					}
				}
			}
		}

		inputs := make([]*Type, len(slots))

		for i, t := range slots {
			inputs[i] = t
		}

		fn, ok := d.Functions.Find(node.Function, inputs)

		if !ok {
			return errors.New("Function " + node.Function + " not found")
		}

		f.Nodes[id] = NewStandardNode(node, fn)
	}

	return nil
}
