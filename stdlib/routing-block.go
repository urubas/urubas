package stdlib

import (
	"github.com/urubas/urubas/vm"
)

type RoutingDecision map[string]bool

type RoutingBlock struct {
	Fn func(node vm.Node) RoutingDecision
}

func (b RoutingBlock) Execute(node vm.Node) {
	if node.Node().Lazy {
		if !WaitForAcception(node) {
			CloseOutputs(node, false)
			return
		}
	}

	decision := b.Fn(node)
	success := WaitForAcception(node)

	for key := range node.Node().Edges {
		accepted, ok := decision[key]

		node.Output(key).Close(success && ok && accepted)
	}
}
