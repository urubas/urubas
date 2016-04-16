package runtime

import (
	"llvm.org/llvm/bindings/go/llvm"
)

var RuntimeDispatchType llvm.Type = llvm.FunctionType(llvm.VoidType(), []llvm.Type{

}, false)

func RuntimeDispatch(m llvm.Module) llvm.Value {
	return llvm.AddGlobal(m, RuntimeDispatchType, "__u_runtime_dispatch")
}
