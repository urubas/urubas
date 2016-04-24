package ast

type NodeMap map[int]*Node
type OutputMap map[string]*Output

type Literal struct {
	// Target input slot
	Slot int

	// Type
	Type string

	// Runtime value
	Value interface{}
}

type Edge struct {
	// Target node ID
	Target int

	// Target input slot
	Slot int
}

type Output struct {
	// Output name
	Name string

	// Connected edges
	Edges []Edge
}

type Node struct {
	// Node ID
	ID int

	// Thread
	ThreadID int

	// Function
	Function string

	// Input Literals
	Literals []Literal

	// Connected Inputs
	Outputs OutputMap

	// Debug Information
	// File where this node is defined
	File string

	// Line where this node is defined
	Line string

	// Column where this node is defined
	Column int
}

func SingleOutput(name string, edges []Edge) OutputMap {
	m := make(OutputMap)

	m[name] = &Output{
		Name: name,
		Edges: edges,
	}

	return m
}

func SingleEdgeOutput(name string, target, slot int) OutputMap {
	return SingleOutput(name, SingleEdge(target, slot))
}

func SingleEdge(target, slot int) []Edge {
	return []Edge{
		Edge{target, slot},
	}
}

func (n *Node) sanitize() {
	if n.Literals == nil {
		n.Literals = make([]Literal, 0)
	}

	if n.Outputs == nil {
		n.Outputs = make(OutputMap)
	}
}
