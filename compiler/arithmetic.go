package compiler

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

func initializeArithmetic(d *Driver) {
	// +
	registerBinary(d, "u.add", arithmeticTypes, func (bc BuildContext, args []Value) Value {
		return NewValue(bc.Builder().CreateAdd(args[0].Value, args[1].Value, ""), args[0].Type)
	})

	// -
	registerBinary(d, "u.sub", arithmeticTypes, func (bc BuildContext, args []Value) Value {
		return NewValue(bc.Builder().CreateSub(args[0].Value, args[1].Value, ""), args[0].Type)
	})

	// *
	registerBinary(d, "u.mul", arithmeticTypes, func (bc BuildContext, args []Value) Value {
		return NewValue(bc.Builder().CreateMul(args[0].Value, args[1].Value, ""), args[0].Type)
	})

	// /
	registerBinary(d, "u.div", signedTypes, func (bc BuildContext, args []Value) Value {
		return NewValue(bc.Builder().CreateSDiv(args[0].Value, args[1].Value, ""), args[0].Type)
	})

	// /
	registerBinary(d, "u.div", unsignedTypes, func (bc BuildContext, args []Value) Value {
		return NewValue(bc.Builder().CreateUDiv(args[0].Value, args[1].Value, ""), args[0].Type)
	})

	// %
	registerBinary(d, "u.rem", signedTypes, func (bc BuildContext, args []Value) Value {
		return NewValue(bc.Builder().CreateSRem(args[0].Value, args[1].Value, ""), args[0].Type)
	})

	// %
	registerBinary(d, "u.rem", unsignedTypes, func (bc BuildContext, args []Value) Value {
		return NewValue(bc.Builder().CreateURem(args[0].Value, args[1].Value, ""), args[0].Type)
	})

	// &
	registerBinary(d, "u.and", intTypes, func (bc BuildContext, args []Value) Value {
		return NewValue(bc.Builder().CreateAnd(args[0].Value, args[1].Value, ""), args[0].Type)
	})

	// |
	registerBinary(d, "u.or", intTypes, func (bc BuildContext, args []Value) Value {
		return NewValue(bc.Builder().CreateOr(args[0].Value, args[1].Value, ""), args[0].Type)
	})

	// ^
	registerBinary(d, "u.xor", intTypes, func (bc BuildContext, args []Value) Value {
		return NewValue(bc.Builder().CreateXor(args[0].Value, args[1].Value, ""), args[0].Type)
	})

	// -
	registerUnary(d, "u.neg", signedTypes, func (bc BuildContext, args []Value) Value {
		return NewValue(bc.Builder().CreateNeg(args[0].Value, ""), args[0].Type)
	})

	// !
	registerUnary(d, "u.not", []string{"u.bool"}, func (bc BuildContext, args []Value) Value {
		return NewValue(bc.Builder().CreateNot(args[0].Value, ""), args[0].Type)
	})
}

