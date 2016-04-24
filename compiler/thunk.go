package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Thunk struct {
	Function llvm.Value
}

func (t *Thunk) Name() string {
	return ""
}

func (t *Thunk) InputTypes() []*Type {
	return []*Type{}
}

func (t *Thunk) OutputType() *Type {
	return nil
}

func (t *Thunk) Call(bc BuildContext, args []Value) Value {
	return ConstNull(bc.Intrinsics().Types.Void)
}
