package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

func toLlvmTypes(types []*Type) []llvm.Type {
	result := make([]llvm.Type, len(types))

	for i := range types {
		result[i] = types[i].Type
	}

	return result
}
