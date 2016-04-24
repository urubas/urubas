package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type ArrayType struct {
	llvm llvm.Type

	ValueType *Type
	Size int
}

func NewArrayType(t *Type, size int) *ArrayType {
	llvm := llvm.ArrayType(t.Type(), size)

	return &ArrayType{llvm, t, size}
}

func (a *ArrayType) Type() llvm.Type {
	return a.llvm
}

func (a *ArrayType) Kind() TypeKind {
	return ArrayKind
}
