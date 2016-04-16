package runtime

type BuiltinFunctionFn func(n *Node)

type BuiltinFunction struct {
	name string
	inputs []Type
	outputs []Type
	fn BuiltinFunctionFn
}

func NewBuiltinFunction(name string, inputs []Type, outputs []Type, fn BuiltinFunctionFn) *BuiltinFunction {
	return &BuiltinFunction{
		name,
		inputs,
		outputs,
		fn,
	}
}

func (b *BuiltinFunction) Name() string {
	return b.name
}

func (b *BuiltinFunction) InputTypes() []Type {
	return b.inputs
}

func (b *BuiltinFunction) OutputTypes() []Type {
	return b.outputs
}

func (b *BuiltinFunction) Call(n *Node) {
	b.fn(n)
}

func RegisterBuiltin(name string, inputs []Type, outputs []Type, fn BuiltinFunctionFn) {
	GlobalFunctions.Add(NewBuiltinFunction(name, inputs, outputs, fn))
}
