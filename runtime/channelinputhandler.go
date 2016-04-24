package runtime

type ChannelInputValue struct {
	Slot int
	Value Value
}

type ChannelInputClosing struct {
	Slot int
	Accepted bool
}

type ChannelInputHandler struct {
	Channel chan<- ChannelInputValue
	Closer chan<- ChannelInputClosing
}

func (o *ChannelInputHandler) SendInput(slot int, value Value) {
	if o.Channel != nil {
		o.Channel <- ChannelInputValue{slot, value}
	}
}

func (o *ChannelInputHandler) CloseInput(slot int, accept bool) {
	if o.Closer != nil {
		o.Closer <- ChannelInputClosing{slot, accept}
	}
}

func (o *ChannelInputHandler) EnsureRunning() bool {
	return true
}
