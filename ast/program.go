package ast

type NodeMap map[int]*Node
type FunctionMap []*Function
type OutputMap map[string]*Output
type TypeMap map[string]*Type

type ExecutionModel int
const (
	_ ExecutionModel = iota
	InlineExecution = iota
	ParallelExecution = iota
)

type Program struct {
	Name string
	Main string
	Types TypeMap
	Functions FunctionMap
}

type Type struct {
	Name string

	// Debug information
	File string
	Line string
	Column int
}

// A defined function
type Function struct {
	// Function name
	Name string

	// Main Nodes
	Main []int

	// Flags
	Pure bool
	ExecutionModel ExecutionModel

	// Function nodes
	Nodes NodeMap

	// IO Definition
	InputTypes []string
	OutputTypes []string

	// Debug information
	File string
	Line string
	Column int
}

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

