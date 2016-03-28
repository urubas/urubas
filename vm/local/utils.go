package local

func MakeLiteralChannel(literal interface{}) <-chan interface{} {
	ch := make(chan interface{}, 1)
	ch <- literal
	return ch
}
