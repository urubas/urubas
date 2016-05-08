package runtime

import (
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestThreadLifetime(t *testing.T) {
	blueprint := &ProcessBlueprint{
		InputMap: make([][]ThreadVertex, 0),
		OutputCount: 1,
		Threads: []*ThreadBlueprint{
			&ThreadBlueprint{
				ID: 1,
				ValueInputCount: 0,
				StateInputCount: 0,
				OutputMap: make(map[string]int),
				Outputs: [][]ThreadVertex{
					[]ThreadVertex{
						{-1, 0},
					},
				},
				Fn: func(ti ThreadInterface) {
					ti.SendOutput(0, Value(reflect.ValueOf(1337)))
					ti.CloseOutput(0, true)
				},
			},
		},
	}

	output := NewLocalInput()

	process := NewProcess(blueprint)
	process.BindOutput(0, output)

	thread, _ := process.Thread(1, true)
	thread.Run()

	WaitAcceptance(output)

	assert.EqualValues(t, reflect.Value(output.Value()).Int(), 1337, "Output is not correct")
}
