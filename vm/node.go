package vm

import (
	"github.com/urubas/urubas/graph"
)

type InputState int
const (
	_ InputState = iota
	Waiting  = iota
	Open     = iota
	Accepted = iota
	Rejected = iota
)

type InputKind int
const (
	_ InputKind = iota
	ValueInput = iota
	StateInput = iota
)

type Input interface {
	Kind() InputKind
	State() InputState

	Input() <-chan interface{}
	StateChannel() <-chan InputState
}

type Output interface {
	Output() chan<- interface{}
	Send(value interface{})
	Close(accepted bool)
}

type Node interface {
	Node() *graph.Node
	Context() Context

	Input(index int) Input
	Inputs() []Input
	InputSlot(index int) Output
	ValueInputCount() int
	StateInputCount() int
	InputCount() int

	Output(name string) Output
	Outputs() []Output
	OutputCount() int

	Execute()
}
