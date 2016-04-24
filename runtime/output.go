package runtime

type EdgeMap []*Edge

type Edge struct {
	Target InputHandler
	Slot int
}

type Output struct {
	Name string
	Edges EdgeMap
}

func (o *Output) Send(value Value) {
	for _, edge := range o.Edges {
		if !edge.Target.EnsureRunning() {
			continue
		}

		edge.Target.SendInput(edge.Slot, value)
	}
}

func (o *Output) Close(accept bool) {
	for _, edge := range o.Edges {
		edge.Target.CloseInput(edge.Slot, accept)
	}
}

func (o *Output) Finish(value Value) {
	o.Send(value)
	o.Close(true)
}
