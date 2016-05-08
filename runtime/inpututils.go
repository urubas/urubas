package runtime

func WaitAcception(i Input) bool {
	for i.State() < InputAccepted {
		_, ok := <-i.Signal()

		if !ok {
			break
		}
	}

	return i.State() == InputAccepted
}
