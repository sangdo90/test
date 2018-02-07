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
	Blocks           []Block
	BlockchainHeight uint64
	GenesisBlock     *Block
	CandidateBlock   *Block
}

// NewGenesisBlock creates genesis block.
func newGenesisBlock() *Block {
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
func NewBlockchain() *Blockchain {
	b := newGenesisBlock()
	bc := &Blockchain{
		Blocks:           []Block{*b},
		BlockchainHeight: 1, // uint64(len([]Block{*b})), // always 1
		GenesisBlock:     b,
		CandidateBlock:   nil,
	}
	return bc
}

// AddBlock adds a block to blockchain.
func (bc *Blockchain) AddBlock() error {
	bc.Blocks = append(bc.Blocks, *bc.CandidateBlock)
	bc.BlockchainHeight = bc.BlockchainHeight + 1
	bc.CandidateBlock = nil
	return nil
}
