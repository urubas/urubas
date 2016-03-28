package stdlib

import (
	"github.com/urubas/urubas/vm"
)

type StandardBlock struct {
	Fn func(node vm.Node) interface{}
}

func (b StandardBlock) Execute(node vm.Node) {
	defer CloseOutputs(node, WaitForAcception(node))

	if node.Node().Lazy {
		if !WaitForAcception(node) {
			return
		}
	}

	FloodOutputs(node, b.Fn(node))
}
