package runtime

type NodeMap map[int]*Node
type NodeList []*Node

type Process struct {
	ContextMarker int
	Nodes NodeMap
	Inputs []*Output
}

func NewProcess(contextMarker int) *Process {
	return &Process{
		ContextMarker: contextMarker,
		Nodes: make(NodeMap),
		Inputs: make([]*Output, 0),
	}
}

func (p *Process) AddNode(n *Node) {
	n.bind(p)

	p.Nodes[n.ID] = n
}

func (p *Process) Execute() {
	for _, n := range p.Nodes {
		if len(n.Inputs) == 0 {
			n.Run()
		}
	}
}
