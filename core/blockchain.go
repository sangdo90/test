package core

import (
	"crypto/sha256"
	"time"
)

// GlobalBlockchains is set of all blockchains.
var GlobalBlockchains []*Blockchain

// Blockchain is chain of blocks, consisting of ID, Blocks, Height, Genesisblcok, and CurrentBlock.
// ID is the same as index+1
type Blockchain struct {
	ID               uint64
	Blocks           []Block
	BlockchainHeight uint64
	GenesisBlock     *Block
	CandidateBlock   *Block
}

// NewGenesisBlock creates genesis block.
func NewGenesisBlock() *Block {
	b := &Block{
		Header: BlockHeader{
			PreviousHash: sha256.Sum256([]byte{}),
			Timestamp:    time.Now().UnixNano(),
			Index:        0,
		},
		Body: BlockBody{},
	}

	return b
}

// AppendBlockchain appends a blockchain to GlobalBlockchains
func AppendBlockchain(bc *Blockchain) error {
	GlobalBlockchains = append(GlobalBlockchains, bc)
	return nil
}

// NewBlockchain creates blockchain.
// TODO: ID is not updated, needs to modification.
// TODO: ID should be considered to have a unique value,
// TODO: even if the blockchain is deleted.
func NewBlockchain() *Blockchain {
	b := NewGenesisBlock()
	cb := NewBlock(b)
	bc := &Blockchain{
		ID:               0, // init value
		Blocks:           []Block{*b},
		BlockchainHeight: 1, // uint64(len([]Block{*b})), // always 1
		GenesisBlock:     b,
		CandidateBlock:   cb,
	}
	return bc
}

// AddBlock adds a block to blockchain.
func (bc *Blockchain) AddBlock() error {
	bc.Blocks = append(bc.Blocks, *bc.CandidateBlock)
	bc.BlockchainHeight = bc.BlockchainHeight + 1
	bc.CandidateBlock = NewBlock(bc.CandidateBlock)
	return nil
}
