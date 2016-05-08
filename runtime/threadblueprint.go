package runtime

type ThreadFunc func(thread ThreadInterface)

type ThreadVertex struct {
	Target int
	Slot int
}

type ThreadBlueprint struct {
	ID int

	ValueInputCount int
	StateInputCount int

	OutputMap map[string]int
	Outputs [][]ThreadVertex

	Fn ThreadFunc
}
