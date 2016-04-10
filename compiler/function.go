package compiler

import (
	"container/list"
	"llvm.org/llvm/bindings/go/llvm"
)

type Function interface {
	Name() string
	InputTypes() []*Type
	OutputTypes() []*Type

	Emit(bc *BuildContext, args []llvm.Value) llvm.Value
}

type FunctionMap struct {
	functions *list.List
}

var GlobalFunctions FunctionMap = NewFunctionMap()

func NewFunctionMap() FunctionMap {
	return FunctionMap{
		functions: list.New(),
	}
}

func (m *FunctionMap) Add(fn Function) {
	m.functions.PushBack(fn)
}

func (m *FunctionMap) Find(name string, types []*Type) (Function, bool) {
	head := m.functions.Front()

	for head != nil {
		fn := head.Value.(Function)
		input := fn.InputTypes()
		compatible := true

		if fn.Name() != name {
			compatible = false
		}

		if compatible {
			if len(input) != len(types) {
				compatible = false
			}
		}

		if compatible {
			for i, v := range input {
				if v.Name != types[i].Name {
					compatible = false
					break
				}
			}
		}

		if compatible {
			return fn, true
		}

		head = head.Next()
	}

	return nil, false
}
