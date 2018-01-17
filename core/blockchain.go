package core

var GlobalBlockchains []BlockChain

type BlockChain struct {
	ID     uint64
	Blocks []Block
}

func NewBlockChain( /*Parameters*/ ) *BlockChain {
	return nil
}

func (c *BlockChain) AddBlock(b *Block) error {
	return nil
}
