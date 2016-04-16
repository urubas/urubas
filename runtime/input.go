package runtime

import (
	"unsafe"
)

type Value unsafe.Pointer

type InputState int
const (
	_ InputState = iota
	InputWaiting = iota
	InputOpen = iota
	InputAccepted = iota
	InputRejected = iota
)

type Input struct {
	Value chan Value
	State chan InputState
	CurrentState InputState
}

func (i *Input) ensureState(state InputState) {
	if i.CurrentState == state {
		return
	}

	i.CurrentState = state;
	i.State <- state

	if state == InputAccepted || state == InputRejected {
		close(i.Value)
		close(i.State)
	}
}
