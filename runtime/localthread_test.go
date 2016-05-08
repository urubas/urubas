package runtime

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestThreadLifetime(t *testing.T) {
	var th *LocalThread

	ch := make(chan interface{})

	th = NewLocalThread(nil, simpleBlueprint(func(ti ThreadInterface) {
		assert.EqualValues(t, th.ID(), ti.Thread().ID(), "Thread interface thread isn't the same as the original thread")
		assert.EqualValues(t, ThreadRunning, th.State(), "Expected thread state to be ThreadRunning")

		close(ch)
	}))

	assert.NotEqual(t, 0, th.ID(), "Thread ID can't be 0")
	assert.EqualValues(t, ThreadWaiting, th.State(), "Expected thread state to be ThreadWaiting")

	th.Run()

	// Wait thread to exit
	for _ = range(ch) {}

	assert.EqualValues(t, ThreadFinished, th.State(), "Expected thread state to be ThreadFinished")
}

func TestThreadInput(t *testing.T) {
	ch := make(chan interface{})

	th := NewLocalThread(nil, simpleBlueprint(func(ti ThreadInterface) {

		close(ch)
	}))

	th.Run()

	// Wait thread to exit
	for _ = range(ch) {}
}

func simpleBlueprint(fn ThreadFunc) *ThreadBlueprint {
	return &ThreadBlueprint{
		0,
		0,
		0,
		nil,
		nil,
		fn,
	}
}
