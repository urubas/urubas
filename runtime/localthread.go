package runtime

import (
	"sync/atomic"
)

type LocalThread struct {
	handle int
	process *Process
	state ThreadState
	blueprint *ThreadBlueprint

	inputs []*LocalInput
	outputs []InputReceiver
}

var threadHandleCounter int32 = 0

func NewLocalThread(p *Process, blueprint *ThreadBlueprint) *LocalThread {
	handle := atomic.AddInt32(&threadHandleCounter, 1)

	th := &LocalThread{
		handle: int(handle),
		state: ThreadWaiting,
		process: p,
		blueprint: blueprint,
		outputs: make([]InputReceiver, len(blueprint.Outputs)),
		inputs: make([]*LocalInput, blueprint.StateInputCount + blueprint.ValueInputCount),
	}

	// Build input list
	for index := range th.inputs {
		th.inputs[index] = NewLocalInput()
	}

	// Build output list
	for index := range th.outputs {
		th.outputs[index] = p.GetInputReceivers(blueprint.Outputs[index])
	}

	return th
}

func (t *LocalThread) Handle() int {
	return t.handle
}

func (t *LocalThread) Process() *Process {
	return t.process
}

func (t *LocalThread) State() ThreadState {
	return t.state
}

func (t *LocalThread) Blueprint() *ThreadBlueprint {
	return t.blueprint
}

func (t *LocalThread) InputReceiver(index int) InputReceiver {
	return t.inputs[index]
}

func (t *LocalThread) Run() {
	go func(){
		defer func(){
			t.state = ThreadFinished
		}()

		t.state = ThreadRunning
		t.blueprint.Fn(t)
	}()
}

func (t *LocalThread) Thread() Thread {
	return t
}

func (t *LocalThread) Input(index int) Input {
	return t.inputs[index]
}

func (t *LocalThread) SendOutput(index int, value Value) {
	t.outputs[index].Send(value)
}

func (t *LocalThread) CloseOutput(index int, accept bool) {
	t.outputs[index].Close(accept)
}

func (t *LocalThread) SendNamedOutput(name string, value Value) {
	index, ok := t.blueprint.OutputMap[name]

	if !ok {
		return
	}

	t.SendOutput(index, value)
}

func (t *LocalThread) CloseNamedOutput(name string, accept bool) {
	index, ok := t.blueprint.OutputMap[name]

	if !ok {
		return
	}

	t.CloseOutput(index, accept)
}
