package core

type Hash []byte

type Block struct {
	PreviousHash Hash
	Transactions []Transaction
}

func NewBlock( /*Parameters*/ ) *Block {
	return nil
}

func (b *Block) AddTransaction(t *Transaction) error {
}
