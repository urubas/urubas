package graph

type Edge struct {
	Target int
	Slot int
}

type EdgeMap map[string]Edge

type Node struct {
	ID int
	Type string
	Edges EdgeMap
	Metadata interface{}
	Lazy bool
}
