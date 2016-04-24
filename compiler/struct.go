package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type StructField struct {
	Name string
	Index int
	Type *Type
}

type StructType struct {
	llvm llvm.Type

	Fields map[string]StructField
}

func NewStructType(fields []StructField) *StructType {
	fieldMap := make(map[string]StructField)
	types := make([]llvm.Type, len(fields))

	for _, f := range fields {
		fieldMap[f.Name] = f
		types[f.Index] = f.Type.Type()
	}

	llvm := llvm.StructType(types, false)

	return &StructType{llvm, fieldMap}
}

func (s *StructType) Type() llvm.Type {
	return s.llvm
}

func (s *StructType) Kind() TypeKind {
	return StructKind
}
