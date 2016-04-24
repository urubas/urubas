package compiler

type intrinsicFunctionEmitter func(bc BuildContext, args []Value) Value

type intrinsicFunction struct {
	name string

	// IO
	inputTypes []*Type
	outputType *Type

	// Emitter
	emit intrinsicFunctionEmitter
}

func newIntrinsicFunction(name string, i []*Type, o *Type, emit intrinsicFunctionEmitter) *intrinsicFunction {
	return &intrinsicFunction{name, i, o, emit}
}

func (f *intrinsicFunction) Name() string {
	return f.name
}

func (f *intrinsicFunction) InputTypes() []*Type {
	return f.inputTypes
}

func (f *intrinsicFunction) OutputType() *Type {
	return f.outputType
}

func (f *intrinsicFunction) Call(bc BuildContext, args []Value) Value {
	return f.emit(bc, args)
}

func registerUnary(d *Driver, name string, types []string, fn intrinsicFunctionEmitter) {
	for _, t := range types {
		t := d.Types[t]

		d.Functions.Add(newIntrinsicFunction(name, []*Type{t}, t, fn))
	}
}

func registerBinary(d *Driver, name string, types []string, fn intrinsicFunctionEmitter) {
	for _, t := range types {
		t := d.Types[t]

		d.Functions.Add(newIntrinsicFunction(name, []*Type{t, t}, t, fn))
	}
}
