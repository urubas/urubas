package ast

type ExecutionModel int
const (
	_ ExecutionModel = iota
	InlineExecution = iota
	ParallelExecution = iota
)

type TypeKind int
const (
	_ TypeKind = iota
	ScalarKind = iota
	ArrayKind = iota
	StructKind = iota
	ChannelKind = iota
)

type Program struct {
	Name string
	Main string
	Types TypeMap
	Functions FunctionMap
}

func NewProgram(name string, main string) *Program {
	return &Program{
		Name: name,
		Main: main,
		Types: make(TypeMap),
		Functions: make(FunctionMap, 0),
	}
}

func (p *Program) AddType(t *Type) {
	p.Types[t.Name] = t
}

func (p *Program) AddFunction(f *Function) {
	p.Functions = append(p.Functions, f)
}
