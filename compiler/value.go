package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Value struct {
	Value llvm.Value
	Type *Type
}

func NewValue(value llvm.Value, t *Type) Value {
	return Value{value, t}
}

func ConstNull(t *Type) Value {
	return NewValue(llvm.ConstNull(t.Type()), t)
}

func ConstInt(t *Type, i uint64) Value {
	return NewValue(llvm.ConstInt(t.Type(), i, false), t)
}

func ConstFloat(t *Type, i float64) Value {
	return NewValue(llvm.ConstFloat(t.Type(), i), t)
}
