package stdlib

import (
	"github.com/urubas/urubas/vm"
)

func init() {
	vm.RegisterBlock(&vm.Block{"u.finish", true, StandardBlock{finish}})
}

func finish(node vm.Node) interface{} {
	if (!WaitForAcception(node)) {
		return nil
	}

	PipeChannel(node.Input(0).Input(), node.Context().ResultChannel())

	return nil
}
