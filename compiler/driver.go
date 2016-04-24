package compiler

import (
	"github.com/urubas/urubas/ast"
)

type Driver struct {
	target *ast.Program

	Program *Program
	Intrinsics *Intrinsics
	Functions FunctionMap
	Types TypeMap
}

func NewDriver(program *ast.Program) *Driver {
	fns := NewFunctionMap()

	d := &Driver{
		target: program,
		Functions: fns,
		Types: make(TypeMap),
		Program: NewProgram(program, fns),
	}

	d.Intrinsics = NewIntrinsics(d)

	return d
}

func (d *Driver) Compile() (*Program, error) {
	functions := make([]*FunctionHolder, len(d.target.Functions))

	// Create Functions
	for i, ast := range d.target.Functions {
		fn := NewFunction(ast, d.Program)

		fn.Prepare(d)

		holder := d.Functions.Add(fn)
		functions[i] = holder
	}

	// Build Functions
	for _, f := range functions {
		fn := f.Fn().(*Function)
		err := fn.Build(d)

		if err != nil {
			return nil, err
		}
	}

	return d.Program, nil
}
