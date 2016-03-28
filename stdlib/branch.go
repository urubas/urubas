package stdlib

import (
	"github.com/urubas/urubas/vm"
)

func init() {
	vm.RegisterBlock(&vm.Block{"u.branch", true, RoutingBlock{branch}})
	vm.RegisterBlock(&vm.Block{"u.invoke", true, StandardBlock{invoke}})
}

func branch(node vm.Node) RoutingDecision {
	right := false

	if (<-node.Input(0).Input()).(bool) {
		right = true
	}

	return RoutingDecision{ "right": right, "left": !right }
}

func invoke(node vm.Node) interface{} {
	main := vm.CallingContext{
		node.Node().Metadata.(int),
		node.ValueInputCount(),
		node.StateInputCount(),
	}

	result := make(chan interface{})
	c := node.Context().CreateSubcontext(main, result)
	n := c.GetNode(main.Entrypoint)

	ConnectInputs(node, n)

	return <-result
}
