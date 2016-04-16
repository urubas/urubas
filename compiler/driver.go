package compiler

import (
	"errors"
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/ast"
)

type Driver struct {
	target *ast.Program
	program *Program
	builder llvm.Builder
}

func NewDriver(program *ast.Program) *Driver {
	p := &Program{
		Ast: program,
		Module: llvm.NewModule(program.Name),
		Types: make(TypeMap),
		Functions: NewFunctionMap(),
	}

	return &Driver{
		target: program,
		program: p,
		builder: llvm.NewBuilder(),
	}
}

func (d *Driver) Compile() (*Program, error) {
	// Compile Functions
	for _, ast := range d.target.Functions {
		_, err := d.CompileFunction(ast)

		if err != nil {
			return nil, err
		}
	}

	return d.program, nil
}

func (d *Driver) FindTypes(names []string) ([]*Type, bool) {
	result := make([]*Type, len(names))

	for i := range names {
		t, ok := d.FindType(names[i])

		if !ok {
			return result, false
		}

		result[i] = t
	}

	return result, true
}

func (d *Driver) FindType(name string) (*Type, bool) {
	t, ok := d.program.Types[name]

	if !ok {
		t, ok := FindType(name)
	}

	return t, ok
}

func (d *Driver) FindFunction(name string, types []*Type) (Function, bool) {
	fn, ok := d.program.Functions.Find(name, types)

	if !ok {
		fn, ok = GlobalFunctions.Find(name, types)
	}

	return fn, ok
}

func (d *Driver) CompileFunction(f *ast.Function) (Function, error) {
	var returnType llvm.Type

	// Find input types
	inputTypes, ok := d.FindTypes(f.InputTypes)

	if !ok {
		return nil, errors.New("Missing types")
	}

	// Find output types
	outputTypes, ok := d.FindTypes(f.OutputTypes)

	if !ok {
		return nil, errors.New("Missing types")
	}

	// Build native return type
	if len(outputTypes) == 1 {
		returnType = outputTypes[0].Type
	} else {
		returnType = llvm.StructType(toLlvmTypes(outputTypes), false)
	}

	// Build blocks
	blocks := make(BlockMap)
	builder := llvm.NewBuilder()
	context := &BuildContext{
		Driver: driver,
		Builder: builder,
	}

	// Build function type
	fnType := llvm.FunctionType(returnType, toLlvmTypes(inputTypes), false)

	// Build function
	fn := llvm.AddFunction(d.program.Module, f.Name, fnType);

	return &UserFunction{
		ast: f,
		inputTypes: inputTypes,
		outputTypes: outputTypes,
		blocks: blocks,
	}, nil
}
