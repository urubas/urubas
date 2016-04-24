package jit

import (
	"github.com/urubas/urubas/runtime"
)

type InputHandlerProxy struct {
	Handler runtime.InputHandler
}

func (o *InputHandlerProxy) Replace(handler runtime.InputHandler) {
	o.Handler = handler
}

func (o *InputHandlerProxy) SendInput(slot int, value runtime.Value) {
	o.Handler.SendInput(slot, value)
}

func (o *InputHandlerProxy) CloseInput(slot int, accept bool) {
	o.Handler.CloseInput(slot, accept)
}

func (o *InputHandlerProxy) EnsureRunning() bool {
	return o.Handler.EnsureRunning()
}
