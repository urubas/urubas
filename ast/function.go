package ast

type FunctionMap []*Function

type ArgumentDefinition struct {
	Type string
	Edges []Edge
}

// A defined function
type Function struct {
	// Function name
	Name string

	// Flags
	Pure bool

	// Function nodes
	Nodes NodeMap

	// IO Definition
	Arguments []ArgumentDefinition
	OutputType string

	// Debug information
	File string
	Line string
	Column int
}

func NewFunction(name string) *Function {
	return &Function{
		Name: name,
		Nodes: make(NodeMap),
	}
}

func (f *Function) AddNode(n *Node) {
	n.sanitize()

	f.Nodes[n.ID] = n
}
