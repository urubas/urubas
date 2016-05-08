package runtime

type LocalInput struct {
	signal chan InputState
	state InputState
	value Value
}

func NewLocalInput() *LocalInput {
	return &LocalInput{
		signal: make(chan InputState),
		state: InputPending,
	}
}

func (i *LocalInput) Signal() <-chan InputState {
	return i.signal
}

func (i *LocalInput) State() InputState {
	return i.state
}

func (i *LocalInput) Value() Value {
	return i.value
}

func (i *LocalInput) Send(value Value) {
	i.value = value
	i.setState(InputAvailable)
}

func (i *LocalInput) Close(accept bool) {
	var state InputState

	if accept {
		state = InputAccepted
	} else {
		state = InputRejected
	}

	i.setState(state)
	close(i.signal)
}

func (i *LocalInput) setState(state InputState) {
	if i.state < state {
		i.state = state
		i.signal <- state
	}
}
