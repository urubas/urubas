package runtime

type InputState int32
const (
	_ InputState = iota

	InputPending = iota
	InputAvailable = iota
	InputAccepted = iota
	InputRejected = iota
)

type Input interface {
	Signal() <-chan InputState
	State() InputState
	Value() Value
}

type InputReceiver interface {
	Send(value Value)
	Close(accept bool)
}
