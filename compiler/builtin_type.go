package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

func initializeTypes() {
	RegisterType(newBuiltinType("u.void", llvm.VoidType()))
	RegisterType(newBuiltinType("u.bool", llvm.Int1Type()))

	// Unsigned
	RegisterType(newBuiltinType("u.uint8", llvm.Int8Type()))
	RegisterType(newBuiltinType("u.uint16", llvm.Int16Type()))
	RegisterType(newBuiltinType("u.uint32", llvm.Int32Type()))
	RegisterType(newBuiltinType("u.uint64", llvm.Int64Type()))

	// Signed
	RegisterType(newBuiltinType("u.int8", llvm.Int8Type()))
	RegisterType(newBuiltinType("u.int16", llvm.Int16Type()))
	RegisterType(newBuiltinType("u.int32", llvm.Int32Type()))
	RegisterType(newBuiltinType("u.int64", llvm.Int64Type()))

	// Floating Point
	RegisterType(newBuiltinType("u.float", llvm.FloatType()))
	RegisterType(newBuiltinType("u.double", llvm.DoubleType()))
}

func newBuiltinType(name string, t llvm.Type) *Type {
	return &Type{name, nil, t}
}
