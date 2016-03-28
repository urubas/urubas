package stdlib

import (
	"reflect"
	"github.com/urubas/urubas/vm"
)

func WaitForAcception(node vm.Node) bool {
	count := node.InputCount()
	cases := make([]reflect.SelectCase, count)

	for i := range cases {
		cases[i].Dir = reflect.SelectRecv
		cases[i].Chan = reflect.ValueOf(node.Input(i).StateChannel())
	}

	for count > 0 {
		chosen, value, ok := reflect.Select(cases)

		if ok {
			if (value.Interface().(vm.InputState) == vm.Rejected) {
				return false
			}
		} else {
			cases[chosen].Chan = reflect.ValueOf(nil)
			count--
		}
	}

	return true
}

func FloodOutputs(node vm.Node, value interface{}) {
	for _, output := range node.Outputs() {
		output.Send(value)
	}
}

func CloseOutputs(node vm.Node, accepted bool) {
	for _, output := range node.Outputs() {
		output.Close(accepted)
	}
}

func PipeChannel(in <-chan interface{}, out chan<- interface{}) {
	for v := range in {
		out <- v
	}
}

func ConnectInputs(in vm.Node, out vm.Node) {
	for index, input := range in.Inputs() {
		go func() {
			slot := out.InputSlot(index)
			output := slot.Output()
			value := input.Input()
			state := input.StateChannel()

			for v := range value {
				output <- v
			}

			for _ = range state {}

			slot.Close(input.State() == vm.Accepted)
		}()
	}
}
