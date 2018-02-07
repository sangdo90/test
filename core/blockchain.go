package core

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/smartm2m/blockchain/common"
)

// GlobalBlockchains is set of all blockchains.
var GlobalBlockchains []*Blockchain

// ChainsID is length of GlobalBlockchains.
var ChainsID uint64

// Blockchain is chain of blocks, consisting of ID, Blocks, Height, Genesisblcok, and CurrentBlock.
// ID is the same as index+1
type Blockchain struct {
	ID               uint64
	Blocks           []Block
	BlockchainHeight uint64
	GenesisBlock     *Block
	CandidateBlock   *Block
	TotalAmount      uint64
}

// SetSeedUsingTime sets the seed using time.
func SetSeedUsingTime() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// NewGenesisBlock creates genesis block.
func NewGenesisBlock() *Block {
	SetSeedUsingTime()
	// token := make([]byte, 32)
	// rand.Read(token)
	token := new(common.Hash)
	rand.Read(token[:])

	diff := new(common.Hash)
	diff[0] = 127

	b := &Block{
		Header: BlockHeader{
			PreviousHash: common.SHA2Hash(token[:]),
			//MerkleRootHash: MerkleRootHash(tx),
			Difficulty: *diff,
			Nonce:      0,
			Timestamp:  common.MakeTimestamp(),
			Index:      0,
		},
		Body: BlockBody{
			Transactions: nil,
		},
	}

	tx := NewTransaction(0, new(common.Address), 50)
	b.AddTransaction(tx)

	return b
}

// SelectBlockchain returns blockchain that has the input id.
func SelectBlockchain(bcid uint64) (*Blockchain, error) {
	if bcid == 0 {
		bcid = 1
	}

	if GlobalBlockchains == nil {
		return nil, errors.New("Blockchain is not exist")
	}

	if bcid > ChainsID {
		return nil, errors.New("Invalid Select Blockchain")
	}

	return GlobalBlockchains[bcid-1], nil
}

// RegisterBlockchain registers in the global blockchain.
// It should always be called when creating a new blockchain(including when cutting).
func (bc *Blockchain) RegisterBlockchain() error {
	ChainsID = ChainsID + 1
	bc.ID = ChainsID
	GlobalBlockchains = append(GlobalBlockchains, bc)
	return nil
}

// AppendBlockchain creates blockchain.
func AppendBlockchain() *Blockchain {
	b := NewGenesisBlock()
	cb := NewBlock(b)
	bc := &Blockchain{
		ID:               0, // init value
		Blocks:           []Block{*b},
		BlockchainHeight: 1, // uint64(len([]Block{*b})), // always 1
		GenesisBlock:     b,
		CandidateBlock:   cb,
		TotalAmount:      50,
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

// Block (n uint64) returns the n-th block.
func (bc *Blockchain) Block(n uint64) *Block {
	blockLength := uint64(len(bc.Blocks))
	if n < 0 || n >= blockLength || blockLength < 1 {
		return nil
	}

	return &bc.Blocks[n]
}

// String (blockchain) function provides information about the blockchain.
// TODO: Convert to string from a blockchain
func (bc *Blockchain) String() string {
	res := bytes.NewBuffer([]byte{})
	fmt.Fprintf(res, "\nID     %v\n", bc.ID)
	fmt.Fprintf(res, "Height %v\n\n", bc.BlockchainHeight)
	fmt.Fprintf(res, "Genesis Block \n%v\n", bc.GenesisBlock.String("Genesis"))
	i := 1
	for _, b := range bc.Blocks[1:] {
		fmt.Fprintf(res, "Block %d %v\n", i, b.String(""))
		i = i + 1
	}
	fmt.Fprintf(res, "Candidate Block \n%v", bc.CandidateBlock.String("Candidate"))

	return res.String()
}
