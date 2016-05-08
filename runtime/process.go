package runtime

import (
	"sync/atomic"
)

type Process struct {
	handle int
	blueprint *ProcessBlueprint
	threads ThreadMap

	inputs []InputReceiver
	outputs []*InputReceiverWrapper
}

var processHandleCounter int32 = 0

func NewProcess(blueprint *ProcessBlueprint) *Process {
	handle := atomic.AddInt32(&processHandleCounter, 1)

	p := &Process{
		handle: int(handle),
		blueprint: blueprint,
		threads: make(ThreadMap),
		inputs: make([]InputReceiver, len(blueprint.InputMap)),
		outputs: make([]*InputReceiverWrapper, blueprint.OutputCount),
	}

	// Create inputs
	for index := range p.inputs {
		v := blueprint.InputMap[index]

		p.inputs[index] = p.GetInputReceiver(v.Target, v.Slot, false)
	}

	// Create outputs
	for index := range p.outputs {
		p.outputs[index] = NewInputReceiverWrapper(nil)
	}

	return p
}

func (p *Process) Handle() int {
	return p.handle
}

func (p *Process) Blueprint() *ProcessBlueprint {
	return p.blueprint
}

func (p *Process) Thread(id int, create bool) (Thread, bool) {
	t, ok := p.threads[id]

	if !ok && create {
		blueprint, ok := p.blueprint.ThreadById(id)

		if !ok {
			return nil, false
		}

		t = NewLocalThread(p, blueprint)

		p.threads[id] = t
	}

	return t, ok
}

func (p *Process) GetInputReceiver(thread, slot int, create bool) InputReceiver {
	if thread == -1 {
		return p.outputs[slot]
	}

	t, ok := p.Thread(thread, create)

	if ok {
		return t.InputReceiver(slot)
	}

	return NewLazyInputReceiver(p, thread, slot)
}

func (p *Process) BindOutput(slot int, target InputReceiver) {
	p.outputs[slot].Bind(target)
}
