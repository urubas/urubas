package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

var arithmeticTypes = []string{
	"u.uint8",
	"u.uint16",
	"u.uint32",
	"u.uint64",
	"u.int8",
	"u.int16",
	"u.int32",
	"u.int64",
	"u.int32",
	"u.float",
	"u.double",
}

var intTypes = []string{
	"u.uint8",
	"u.uint16",
	"u.uint32",
	"u.uint64",
	"u.int8",
	"u.int16",
	"u.int32",
	"u.int64",
	"u.int32",
}

var unsignedTypes = []string{
	"u.uint8",
	"u.uint16",
	"u.uint32",
	"u.uint64",
	"u.uint32",
}

var signedTypes = []string{
	"u.int8",
	"u.int16",
	"u.int32",
	"u.int64",
	"u.int32",
	"u.float",
	"u.double",
}

func initializeArithmetic() {
	// +
	registerBinary("u.add", arithmeticTypes, func(bc *BuildContext, args []llvm.Value) llvm.Value {
		return bc.Builder.CreateAdd(args[0], args[1], "")
	})

	// -
	registerBinary("u.sub", arithmeticTypes, func(bc *BuildContext, args []llvm.Value) llvm.Value {
		return bc.Builder.CreateSub(args[0], args[1], "")
	})

	// *
	registerBinary("u.mul", arithmeticTypes, func(bc *BuildContext, args []llvm.Value) llvm.Value {
		return bc.Builder.CreateMul(args[0], args[1], "")
	})

	// /
	registerBinary("u.div", signedTypes, func(bc *BuildContext, args []llvm.Value) llvm.Value {
		return bc.Builder.CreateSDiv(args[0], args[1], "")
	})

	// /
	registerBinary("u.div", unsignedTypes, func(bc *BuildContext, args []llvm.Value) llvm.Value {
		return bc.Builder.CreateUDiv(args[0], args[1], "")
	})

	// %
	registerBinary("u.rem", signedTypes, func(bc *BuildContext, args []llvm.Value) llvm.Value {
		return bc.Builder.CreateSRem(args[0], args[1], "")
	})

	// %
	registerBinary("u.rem", unsignedTypes, func(bc *BuildContext, args []llvm.Value) llvm.Value {
		return bc.Builder.CreateURem(args[0], args[1], "")
	})

	// &
	registerBinary("u.and", intTypes, func(bc *BuildContext, args []llvm.Value) llvm.Value {
		return bc.Builder.CreateAnd(args[0], args[1], "")
	})

	// |
	registerBinary("u.or", intTypes, func(bc *BuildContext, args []llvm.Value) llvm.Value {
		return bc.Builder.CreateOr(args[0], args[1], "")
	})

	// ^
	registerBinary("u.xor", intTypes, func(bc *BuildContext, args []llvm.Value) llvm.Value {
		return bc.Builder.CreateXor(args[0], args[1], "")
	})

	// -
	registerUnary("u.neg", signedTypes, func(bc *BuildContext, args []llvm.Value) llvm.Value {
		return bc.Builder.CreateNeg(args[0], "")
	})

	// !
	registerUnary("u.not", []string{"u.bool"}, func(bc *BuildContext, args []llvm.Value) llvm.Value {
		return bc.Builder.CreateNot(args[0], "")
	})
}

