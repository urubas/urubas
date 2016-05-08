package runtime

type NilInputReceiver struct {
}

func NewNilInputReceiver() *NilInputReceiver {
	return &NilInputReceiver{}
}

func (d *NilInputReceiver) Send(value Value) {
}

func (d *NilInputReceiver) Close(accept bool) {
}
