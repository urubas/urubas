package local

import (
	"github.com/urubas/urubas/vm"
	"github.com/urubas/urubas/graph"
)

type IOPoint struct {
	state vm.InputState
	channel chan interface{}
	stateChannel chan vm.InputState
	kind vm.InputKind
}

type Node struct {
	node *graph.Node
	context *Context
	running bool

	inputs []*IOPoint
	outputs map[string]vm.Output

	valueInputCount int
	stateInputCount int
}

func NewNode(context *Context, node *graph.Node, valueCount int, stateCount int) *Node {
	n := &Node{
		node: node,
		context: context,
		inputs: make([]*IOPoint, valueCount + stateCount),
		outputs: make(map[string]vm.Output),
		valueInputCount: valueCount,
		stateInputCount: stateCount,
	}

	for i := 0; i < len(n.inputs); i++ {
		var kind vm.InputKind

		if i < valueCount {
			kind = vm.ValueInput
		} else {
			kind = vm.StateInput
		}

		n.inputs[i] = &IOPoint{
			state: vm.Waiting,
			channel: make(chan interface{}, 1),
			stateChannel: make(chan vm.InputState, 1),
			kind: kind,
		}
	}

	for k, v := range node.Edges {
		targetNode := context.GetNode(v.Target)
		n.outputs[k] = targetNode.InputSlot(v.Slot)
	}

	return n
}

func (n *Node) Node() *graph.Node {
	return n.node
}

func (n *Node) Context() vm.Context {
	return n.context
}

func (n *Node) Input(index int) vm.Input {
	return n.inputs[index]
}

func (n *Node) Inputs() []vm.Input {
	inputs := make([]vm.Input, len(n.inputs))

	for i := range n.inputs {
		inputs[i] = n.inputs[i]
	}

	return inputs
}

func (n *Node) InputSlot(index int) vm.Output {
	return n.inputs[index]
}

func (n *Node) InputCount() int {
	return len(n.inputs)
}

func (n *Node) ValueInputCount() int {
	return n.valueInputCount
}

func (n *Node) StateInputCount() int {
	return n.stateInputCount
}

func (n *Node) Output(name string) vm.Output {
	return n.outputs[name]
}

func (n *Node) Outputs() []vm.Output {
	index := 0
	keys := make([]vm.Output, len(n.outputs))

	for _, v := range n.outputs {
		keys[index] = v
		index++
	}

	return keys
}

func (n *Node) OutputCount() int {
	return len(n.outputs)
}

func (n *Node) Execute() {
	if (n.running) {
		return
	}

	n.running = true
	vm.GetBlock(n.node.Type).Execute(n)
}

func (i *IOPoint) Kind() vm.InputKind {
	return i.kind
}

func (i *IOPoint) State() vm.InputState {
	return i.state
}

func (i *IOPoint) Input() <-chan interface{} {
	return i.channel
}

func (i *IOPoint) Output() chan<- interface{} {
	return i.channel
}

func (i *IOPoint) StateChannel() <-chan vm.InputState {
	return i.stateChannel
}

func (i *IOPoint) Send(value interface{}) {
	if (i.state != vm.Waiting && i.state != vm.Open) {
		return
	}

	i.channel <- value
	i.changeState(vm.Open)
}

func (i *IOPoint) Close(accepted bool) {
	var newState vm.InputState

	if accepted {
		newState = vm.Accepted
	} else {
		newState = vm.Rejected
	}

	i.changeState(newState)

	close(i.channel)
	close(i.stateChannel)
}

func (i *IOPoint) changeState(state vm.InputState) {
	if (i.state != state) {
		i.state = state
		i.stateChannel <- state
	}
}
