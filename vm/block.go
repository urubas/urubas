package vm

type BlockHandler interface {
	Execute(node Node)
}

type Block struct {
	Name string
	Pure bool
	Handler BlockHandler
}

var Blocks map[string]*Block

func RegisterBlock(block *Block) {
	Blocks[block.Name] = block
}

func GetBlock(name string) *Block {
	return Blocks[name]
}

func (b *Block) Execute(node Node) {
	b.Handler.Execute(node)
}
