package runtime

import (
	"errors"
)

type InputMap []*Input
type OutputMap map[string]*Output

type NodeState int
const (
	_ NodeState = iota
	NodeIdle = iota
	NodeRunning = iota
	NodeFinished = iota
	NodeAccepted = iota
	NodeRejected = iota
)

type Node struct {
	ID int
	Process *Process

	Function Function
	State NodeState

	Inputs InputMap
	Outputs OutputMap
}

func NewNode(id int, fn Function) *Node {
	return &Node{
		ID: id,
		Function: fn,
		State: NodeIdle,
		Inputs: make(InputMap, 0),
		Outputs: make(OutputMap),
	}
}

func (n *Node) Run() {
	if n.State != NodeIdle {
		return
	}

	n.State = NodeRunning

	go func() {
		n.Function.Call(n)

		// TODO: Probably needs synchronization
		if n.State == NodeRunning {
			n.State = NodeFinished
		}
	}()
}

func (n *Node) EnsureRunning() bool {
	switch (n.State) {
		case NodeIdle:
			n.Run()
			return true
		case NodeRunning:
			return true
		case NodeFinished:
			return true
		case NodeAccepted:
			return false
		case NodeRejected:
			return false
	}

	return false
}

func (n *Node) Connect(name string, target InputHandler, slot int) {
	var output *Output

	output, ok := n.Outputs[name]

	if !ok {
		output = &Output{
			Name: name,
			Edges: make(EdgeMap, 1),
		}
	}

	output.Edges = append(output.Edges, &Edge{
		Target: target,
		Slot: slot,
	})
}

func (n *Node) SendInput(slot int, value Value) {
	input := n.Inputs[slot]

	if input.State != InputWaiting {
		panic(errors.New("Input already resolved"))
	}

	input.setValue(value)
}


func (n *Node) CloseInput(slot int, accept bool) {
	var state InputState

	input := n.Inputs[slot]

	if input.State == InputAccepted || input.State == InputRejected {
		panic(errors.New("Input already resolved"))
	}

	if accept {
		state = InputAccepted
	} else {
		state = InputRejected
	}

	input.setState(state)

	if accept {
		n.checkAcception()
	} else {
		n.State = NodeRejected
	}
}

func (n *Node) checkAcception() {
	for _, i := range n.Inputs {
		if i.State != InputAccepted {
			return
		}
	}

	n.State = NodeAccepted
}

func (n *Node) bind(p *Process) {
	n.Process = p
}

func (n *Node) WaitForArguments() {
	for _, i := range n.Inputs {
		i.WaitResolve()
	}
}
