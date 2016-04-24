package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type ScalarType struct {
	llvm llvm.Type

	Size int
}

func NewScalarType(t llvm.Type, size int) *ScalarType {
	return &ScalarType{t, size}
}

func (s *ScalarType) Type() llvm.Type {
	return s.llvm
}

func (s *ScalarType) Kind() TypeKind {
	return ScalarKind
}
