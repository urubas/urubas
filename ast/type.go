package ast

type TypeMap map[string]*Type

type Type struct {
	Name string

	Kind TypeKind

	// Debug information
	File string
	Line string
	Column int
}
