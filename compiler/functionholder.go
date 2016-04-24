package compiler

type FunctionHolder struct {
	fn Callable
}

func (f *FunctionHolder) Fn() Callable {
	return f.fn
}

func (f *FunctionHolder) Replace(fn Callable) {
	// TODO: Check if new function is compatible with the old one

	f.fn = fn
}

func (f *FunctionHolder) Name() string {
	return f.fn.Name()
}

func (f *FunctionHolder) InputTypes() []*Type {
	return f.fn.InputTypes()
}

func (f *FunctionHolder) OutputType() *Type {
	return f.fn.OutputType()
}

func (f *FunctionHolder) Call(bc BuildContext, args []Value) Value {
	return f.fn.Call(bc, args)
}
