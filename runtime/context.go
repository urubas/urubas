package runtime

type NodeMap map[int]*Node
type NodeList []*Node

type Context struct {
	Nodes NodeMap
}

func NewContext() *Context {
	return &Context{
		Nodes: make(NodeMap),
	}
}

func (c *Context) AddNode(n *Node) {
	n.bind(c)

	c.Nodes[n.ID] = n
}
