package compiler

import (
	"strconv"
	"llvm.org/llvm/bindings/go/llvm"
	"github.com/BTBurke/toposort"
	"github.com/urubas/urubas/ast"
)

func toLlvmTypes(types []*Type) []llvm.Type {
	result := make([]llvm.Type, len(types))

	for i := range types {
		result[i] = types[i].Type()
	}

	return result
}

func argumentsToTypes(d *Driver, args []ast.ArgumentDefinition) []*Type {
	result := make([]*Type, len(args))

	for i, arg := range args {
		result[i] = d.Types[arg.Type]
	}

	return result
}

type nodeSorting struct {
	nodelist []int
	nodemap ast.NodeMap
}

func (n *nodeSorting) Len() int {
	return len(n.nodelist)
}

func (n *nodeSorting) Label(index int) string {
	return strconv.Itoa(n.nodelist[index])
}

func (n *nodeSorting) Dependencies(index int) []string {
	node := n.nodemap[n.nodelist[index]]
	result := make([]string, 0)

	for _, o := range node.Outputs {
		for _, e := range o.Edges {
			result = append(result, strconv.Itoa(e.Target))
		}
	}

	return result
}

func sortNodes(nodes ast.NodeMap) []int {
	list := make([]int, len(nodes))

	i := 0
	for id := range nodes {
		list[i] = id
		i++
	}

	sorting := &nodeSorting{
		nodelist: list,
		nodemap: nodes,
	}

	indexes, _ := toposort.SortIndex(sorting)
	result := make([]int, len(indexes))

	for i, index := range indexes {
		result[len(result) - i - 1] = sorting.nodelist[index - 1]
	}

	return result
}
