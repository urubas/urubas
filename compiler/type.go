package compiler

import (
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/urubas/urubas/ast"
)

type TypeKind int
const (
	_ TypeKind = iota
	ScalarKind = iota
	ArrayKind = iota
	StructKind = iota
	ChannelKind = iota
)

type UnderlyingType interface {
	Type() llvm.Type
	Kind() TypeKind
}

type TypeMap map[string]*Type

type Type struct {
	ast *ast.Type
	name string
	underlying UnderlyingType
}

func NewType(ast *ast.Type, name string, underlying UnderlyingType) *Type {
	return &Type{
		ast: ast,
		name: name,
		underlying: underlying,
	}
}

func NewArray(ast *ast.Type, name string, t *Type, size int) *Type {
	return NewType(ast, name, NewArrayType(t, size))
}

func NewStruct(ast *ast.Type, name string, fields []StructField) *Type {
	return NewType(ast, name, NewStructType(fields))
}

func (t *Type) Ast() *ast.Type {
	return t.ast
}

func (t *Type) Name() string {
	return t.name
}

func (t *Type) Type() llvm.Type {
	return t.underlying.Type()
}

func (t *Type) Kind() TypeKind {
	return t.underlying.Kind()
}

func (t *Type) UnderlyingType() UnderlyingType {
	return t.underlying
}
