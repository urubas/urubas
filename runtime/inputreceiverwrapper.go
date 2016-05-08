package runtime

type InputReceiverWrapper struct {
	receiver InputReceiver
}

func NewInputReceiverWrapper(r InputReceiver) *InputReceiverWrapper {
	return &InputReceiverWrapper{r}
}

func (d *InputReceiverWrapper) Send(value Value) {
	d.receiver.Send(value)
}

func (d *InputReceiverWrapper) Close(accept bool) {
	d.receiver.Close(accept)
}

func (d *InputReceiverWrapper) Bind(receiver InputReceiver) {
	d.receiver = receiver
}
