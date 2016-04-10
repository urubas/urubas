package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/ast"
)

type Type struct {
	Name string

	Ast *ast.Type

	// LLVM
	Type llvm.Type
}

var Types map[string]*Type

func RegisterType(t *Type) {
	Types[t.Name] = t
}

func FindType(name string) (*Type, bool) {
	t, ok := Types[name]

	return t, ok
}
