package runtime

type NullInputHandler struct {}

func (o *NullInputHandler) SendInput(slot int, value Value) {
}

func (o *NullInputHandler) CloseInput(slot int, accept bool) {
}

func (o *NullInputHandler) EnsureRunning() bool {
	return true
}
