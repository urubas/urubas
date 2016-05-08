package runtime

type ThreadState int
const (
	_ ThreadState = iota
	ThreadWaiting = iota
	ThreadRunning = iota
	ThreadFinished = iota
)

type ThreadMap map[int]Thread

type ThreadInterface interface {
	Thread() Thread
	Blueprint() *ThreadBlueprint

	Input(index int) Input

	SendOutput(index int, value Value)
	CloseOutput(index int, accept bool)

	SendNamedOutput(name string, value Value)
	CloseNamedOutput(name string, accept bool)
}

type Thread interface {
	Handle() int
	Process() *Process
	Blueprint() *ThreadBlueprint
	State() ThreadState

	InputReceiver(index int) InputReceiver

	Run()
}
