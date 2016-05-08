package runtime

import (
	"errors"
)

type LazyInputReceiver struct {
	process *Process
	target int
	slot int
	receiver InputReceiver
}

func NewLazyInputReceiver(p *Process, target, slot int) *LazyInputReceiver {
	return &LazyInputReceiver{
		process: p,
		target: target,
		slot: slot,
	}
}

func (d *LazyInputReceiver) Send(value Value) {
	d.realize().Send(value)
}

func (d *LazyInputReceiver) Close(accept bool) {
	d.realize().Close(accept)
}

func (d *LazyInputReceiver) realize() InputReceiver {
	if d.receiver == nil {
		thread, ok := d.process.Thread(d.target, true)

		if !ok {
			panic(errors.New("bound to inexistent thread"))
		}

		d.receiver = thread.InputReceiver(d.slot)
	}

	return d.receiver
}
