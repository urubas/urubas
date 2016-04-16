package runtime

import (
	"reflect"
)

type EdgeMap []*Edge

type Edge struct {
	Target *Node
	Slot int
}

type Output struct {
	Name string
	Type Type
	Edges EdgeMap
}

func (o *Output) Send(any interface{}) {
	value := reflect.ValueOf(any)
	ptr := Value(value.Pointer())

	for _, edge := range o.Edges {
		if !edge.Target.EnsureRunning() {
			continue
		}

		ch, ok := edge.Target.OpenInput(edge.Slot)

		if !ok {
			continue
		}

		ch <- ptr
	}
}

func (o *Output) Close(accept bool) {
	for _, edge := range o.Edges {
		_, ok := edge.Target.OpenInput(edge.Slot)

		if !ok {
			continue
		}

		edge.Target.CloseInput(edge.Slot, accept)
	}
}

func (o *Output) Finish(any interface{}) {
	o.Send(any)
	o.Close(true)
}
