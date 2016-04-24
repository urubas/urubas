package compiler

type Callable interface {
	Name() string
	InputTypes() []*Type
	OutputType() *Type
	Call(bc BuildContext, args []Value) Value
}

type CustomFunction interface {
	InputTypesMatches(types []*Type) bool
}
