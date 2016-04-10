package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type builtinFunctionEmitter func(bc *BuildContext, args []llvm.Value) llvm.Value

type builtinFunction struct {
	name string

	// IO
	inputTypes []*Type
	outputTypes []*Type

	// Emitter
	emit builtinFunctionEmitter
}

func newbuiltinFunction(name string, i, o []*Type, emit builtinFunctionEmitter) *builtinFunction {
	return &builtinFunction{name, i, o, emit}
}

func (f *builtinFunction) Name() string {
	return f.name
}

func (f *builtinFunction) InputTypes() []*Type {
	return f.inputTypes
}

func (f *builtinFunction) OutputTypes() []*Type {
	return f.outputTypes
}

func (f *builtinFunction) Emit(bc *BuildContext, args []llvm.Value) llvm.Value {
	return f.emit(bc, args)
}

func registerUnary(name string, types []string, fn builtinFunctionEmitter) {
	for _, t := range types {
		t, _ := FindType(t)

		GlobalFunctions.Add(newbuiltinFunction(name, []*Type{t}, []*Type{t}, fn))
	}
}

func registerBinary(name string, types []string, fn builtinFunctionEmitter) {
	for _, t := range types {
		t, _ := FindType(t)

		GlobalFunctions.Add(newbuiltinFunction(name, []*Type{t, t}, []*Type{t}, fn))
	}
}
