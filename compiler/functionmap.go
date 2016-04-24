package compiler

import (
	"container/list"
)

type FunctionMap struct {
	functions *list.List
}

func NewFunctionMap() FunctionMap {
	return FunctionMap{
		functions: list.New(),
	}
}

func (m *FunctionMap) Add(fn Callable) *FunctionHolder {
	holder := &FunctionHolder{fn}

	m.functions.PushBack(holder)

	return holder
}

func (m *FunctionMap) Find(name string, types []*Type) (*FunctionHolder, bool) {
	head := m.functions.Front()

	for head != nil {
		fn := head.Value.(*FunctionHolder)
		input := fn.InputTypes()
		compatible := true

		if fn.Name() != name {
			compatible = false
		}

		if custom, ok := fn.Fn().(CustomFunction); compatible && ok {
			if !custom.InputTypesMatches(types) {
				compatible = false
			}
		} else {
			if compatible {
				if len(input) != len(types) {
					compatible = false
				}
			}

			if compatible {
				for i, v := range input {
					if v.Name() != types[i].Name() {
						compatible = false
						break
					}
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

func (m *FunctionMap) All() []*FunctionHolder {
	i := 0
	res := make([]*FunctionHolder, m.functions.Len())
	head := m.functions.Front()

	for head != nil {
		res[i] = head.Value.(*FunctionHolder)

		i++
		head = head.Next()
	}

	return res
}
