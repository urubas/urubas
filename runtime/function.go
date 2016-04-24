package runtime

type Function interface {
	Call(n *Node)
}
