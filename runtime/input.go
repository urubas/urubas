package runtime

import (
	"unsafe"
)

type Value unsafe.Pointer

type InputState int
const (
	_ InputState = iota
	InputWaiting = iota
	InputArrived = iota
	InputAccepted = iota
	InputRejected = iota
)

type InputHandler interface {
	SendInput(slot int, value Value)
	CloseInput(slot int, accept bool)
	EnsureRunning() bool
}

type Input struct {
	Value Value
	State InputState
	Signal chan InputState
}

func NewInput() *Input {
	return &Input{
		State: InputWaiting,
		Signal: make(chan InputState, 2),
	}
}

func (i *Input) WaitArrival() {
	for i.State < InputArrived {
		<-i.Signal
	}
}

func (i *Input) WaitResolve() bool {
	for i.State < InputAccepted {
		<-i.Signal
	}

	return i.State == InputAccepted
}

func (i *Input) setValue(value Value) {
	i.Value = value
	i.setState(InputWaiting)
}

func (i *Input) setState(state InputState) {
	if i.State <= state {
		return
	}

	i.State = state;
	i.Signal <- state

	if state == InputAccepted || state == InputRejected {
		close(i.Signal)
	}
}

