package runtime

type InputReceiverDemux struct {
	receivers []InputReceiver
}

func NewInputReceiverDemux() *InputReceiverDemux {
	return &InputReceiverDemux{
		make([]InputReceiver, 0),
	}
}

func NewInputReceiverDemuxWith(inputs []InputReceiver) *InputReceiverDemux {
	d := &InputReceiverDemux{
		make([]InputReceiver, len(inputs)),
	}

	copy(d.receivers, inputs)

	return d
}

func (d *InputReceiverDemux) Add(i InputReceiver) {
	d.receivers = append(d.receivers, i)
}

func (d *InputReceiverDemux) Send(value Value) {
	for _, i := range d.receivers {
		i.Send(value)
	}
}

func (d *InputReceiverDemux) Close(accept bool) {
	for _, i := range d.receivers {
		i.Close(accept)
	}
}
