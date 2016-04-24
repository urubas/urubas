package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type IntrinsicTypes struct {
	Void *Type
	Bool *Type

	UInt8 *Type
	UInt16 *Type
	UInt32 *Type
	UInt64 *Type

	Int8 *Type
	Int16 *Type
	Int32 *Type
	Int64 *Type

	Float *Type
	Double *Type

	String *Type
}

type IntrinsicFunctions struct {
}

type Intrinsics struct {
	Types IntrinsicTypes
	Functions IntrinsicFunctions
}

func NewIntrinsics(d *Driver) *Intrinsics {
	createScalarIntrinsicType := func(name string, t llvm.Type, size int) *Type {
		tt := NewType(nil, name, NewScalarType(t, size))

		d.Types[name] = tt

		return tt
	}

	i := &Intrinsics{
		Types: IntrinsicTypes{
			Void: createScalarIntrinsicType("u.void", llvm.VoidType(), 0),
			Bool: createScalarIntrinsicType("u.bool", llvm.Int1Type(), 1),

			UInt8: createScalarIntrinsicType("u.uint8", llvm.Int8Type(), 1),
			UInt16: createScalarIntrinsicType("u.uint16", llvm.Int16Type(), 2),
			UInt32: createScalarIntrinsicType("u.uint32", llvm.Int32Type(), 4),
			UInt64: createScalarIntrinsicType("u.uint64", llvm.Int64Type(), 8),


			Int8: createScalarIntrinsicType("u.int8", llvm.Int8Type(), 1),
			Int16: createScalarIntrinsicType("u.int16", llvm.Int16Type(), 2),
			Int32: createScalarIntrinsicType("u.int32", llvm.Int32Type(), 4),
			Int64: createScalarIntrinsicType("u.int64", llvm.Int64Type(), 8),

			Float: createScalarIntrinsicType("u.float", llvm.FloatType(), 4),
			Double: createScalarIntrinsicType("u.double", llvm.DoubleType(), 8),
		},
		Functions: IntrinsicFunctions{
		},
	}

	initializeArithmetic(d)
	initializeFlow(d)

	return i
}
