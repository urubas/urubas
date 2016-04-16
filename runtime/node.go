package runtime

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
	Context *Context

	Function Function
	State NodeState

	Inputs InputMap
	StateInputs InputMap

	Outputs OutputMap
}

func NewNode(id int, context *Context, fn Function) *Node {
	return &Node{
		ID: id,
		Function: fn,
		State: NodeIdle,
		Inputs: make(InputMap, 0),
		StateInputs: make(InputMap, 0),
		Outputs: make(OutputMap),
	}
}

func (n *Node) Run() {
	if n.State != NodeIdle {
		return
	}

	n.State = NodeRunning

	n.Function.Call(n)

	// TODO: Probably needs synchronization
	if n.State == NodeRunning {
		n.State = NodeFinished
	}
}

func (n *Node) EnsureRunning() bool {
	switch (n.State) {
		case NodeIdle:
			go n.Run()
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

func (n *Node) OpenInput(slot int) (chan<- Value, bool) {
	input := n.getInput(slot)

	if input.CurrentState == InputAccepted || input.CurrentState == InputRejected {
		return nil, false
	}

	input.ensureState(InputOpen)

	return input.Value, true
}

func (n *Node) CloseInput(slot int, accept bool) {
	var state InputState

	input := n.getInput(slot)

	if input.CurrentState == InputAccepted || input.CurrentState == InputRejected {
		return
	}

	if accept {
		state = InputAccepted
	} else {
		state = InputRejected
	}

	input.ensureState(state)

	if accept {
		n.checkAcception()
	} else {
		n.State = NodeRejected
	}
}

func (n *Node) getInput(slot int) *Input {
	var input *Input

	if slot < 0 {
		input = n.Inputs[slot]
	} else {
		input = n.Inputs[slot * -1 - 1]
	}

	return input
}

func (n *Node) checkAcception() {
	for _, i := range n.Inputs {
		if i.CurrentState != InputAccepted {
			return
		}
	}

	for _, i := range n.StateInputs {
		if i.CurrentState != InputAccepted {
			return
		}
	}

	n.State = NodeAccepted
}

func (n *Node) bind(c *Context) {
	n.Context = c
}
