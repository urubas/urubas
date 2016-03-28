package local

import (
	// "sync"
	"github.com/urubas/urubas/vm"
	"github.com/urubas/urubas/graph"
)

type Context struct {
	main vm.CallingContext
	program *graph.Program
	result chan interface{}

	// mutex sync.Mutex
	nodes map[int]*Node
	nodeMap map[int]*graph.Node
}

func NewContext(program *graph.Program, result chan interface{}) *Context {
	c := &Context{
		program: program,
		result: result,

		nodes: make(map[int]*Node),
		nodeMap: make(map[int]*graph.Node),

		main: vm.CallingContext{
			program.Main,
			0,
			0,
		},
	}

	for _, node := range program.Nodes {
		c.nodeMap[node.ID] = node
	}

	return c
}

func (c *Context) Program() *graph.Program {
	return c.program
}

func (c *Context) ResultChannel() chan interface{} {
	return c.result
}

func (c *Context) Execute(args []<-chan interface{}) {
	c.main.ValueCount = len(args)

	node := c.GetNode(c.main.Entrypoint)

	for k, v := range args {
		go func() {
			input := node.InputSlot(k)

			for value := range v {
				input.Send(value)
			}

			input.Close(true)
		}()
	}
}

func (c *Context) GetNode(node int) vm.Node {
	// c.mutex.Lock()

	result, ok := c.nodes[node]

	if !ok {
		value, state := 0, 0

		if node == c.main.Entrypoint {
			value, state = c.main.ValueCount, c.main.StateCount
		} else {
			value, state = getInputCount(c.program, node)
		}

		result = NewNode(c, c.nodeMap[node], value, state)

		c.nodes[node] = result
	}

	// c.mutex.Unlock()

	go result.Execute()

	return result
}

func (c *Context) CreateSubcontext(main vm.CallingContext, result chan interface{}) vm.Context {
	return &Context{
		main: main,
		program: c.program,
		result: c.result,

		nodes: make(map[int]*Node),
		nodeMap: c.nodeMap,
	}
}

func getInputCount(p *graph.Program, node int) (int, int) {
	value, state := 0, 0

	for _, v := range p.Nodes {
		for _, e := range v.Edges {
			if e.Target == node {
				if (e.Slot < 0) {
					state++
				} else {
					value++
				}
			}
		}
	}

	return value, state
}
