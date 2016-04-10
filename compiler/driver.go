package compiler

import (
	"errors"
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/ast"
)

type Driver struct {
	target *ast.Program
	program *Program
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

func (d *Driver) CompileFunction(ast *ast.Function) (Function, error) {
	var returnType llvm.Type

	// Find input types
	inputTypes, ok := d.FindTypes(ast.InputTypes)

	if !ok {
		return nil, errors.New("Missing types")
	}

	// Find output types
	outputTypes, ok := d.FindTypes(ast.OutputTypes)

	if !ok {
		return nil, errors.New("Missing types")
	}

	// Build native return type
	if ast.ExecutionModel == ast.InlineExecution {
		if len(outputTypes) == 1 {
			returnType = outputTypes[0].Type
		} else {
			returnType = llvm.StructType(toLlvmTypes(outputTypes), false)
		}
	} else {
		returnType = llvm.VoidType()
	}

	// Build function type
	fnType := llvm.FunctionType(returnType, toLlvmTypes(inputTypes), false)

	// Build function
	fn := llvm.AddFunction(d.program.Module, ast.Name, fnType)

	return &UserFunction{
		ast: ast,
		inputTypes: inputTypes,
		outputTypes: outputTypes,
		fnType: fnType,
		function: fn,
	}, nil
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
