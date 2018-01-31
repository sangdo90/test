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
var GlobalBlockchains []Blockchain

// LongestBlockchainID is the longest blockchain's id among the blockchains.
// init value is 0
var LongestBlockchainID uint64

// GlobalBlockchainsLength is length of GlobalBlockchains.
var GlobalBlockchainsLength uint64

// Blockchain is chain of blocks, consisting of ID, Blocks, Height, Genesisblcok, and CurrentBlock.
// ID is the same as index+1
type Blockchain struct {
	ID               uint64
	Blocks           []Block
	BlockchainHeight uint64
	GenesisBlock     *Block
	CandidateBlock   *Block
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

	b := &Block{
		Header: BlockHeader{
			PreviousHash: common.SHA2Hash(token[:]),
			// MerkleRootHash	: GetMerkleRootHash(transactions),
			Difficulty: 0,
			Nonce:      0,
			Timestamp:  common.MakeTimestamp(),
			Index:      0,
		},
		Body: BlockBody{
		// Transactions : append(Transactions, NewTransaction(/*Parameters*/)),
		},
	}

	return b
}

// GetLongestBlockchain gets longest Blockchain.
func GetLongestBlockchain() *Blockchain {
	glb, _ := SelectBlockchain(LongestBlockchainID)
	return glb
}

// SelectBlockchain returns blockchain that has the input id.
func SelectBlockchain(bcid uint64) (*Blockchain, error) {
	if bcid == 0 {
		bcid = 1
	}

	if GlobalBlockchainsLength == 0 {
		return nil, errors.New("Blockchain is not exist")
	}

	if bcid > GlobalBlockchainsLength {
		return nil, errors.New("Invalid Select Blockchain")
	}

	return &GlobalBlockchains[bcid-1], nil
}

// LongestBlockchainUpdate updates the longest block chain's id(global variable, LongestBlockchainID).
func LongestBlockchainUpdate(bc *Blockchain) error {
	lbc, _ := SelectBlockchain(LongestBlockchainID)

	if lbc.BlockchainHeight < bc.BlockchainHeight {
		LongestBlockchainID = bc.ID
	}

	return nil
}

// RegisterBlockchain registers in the global blockchain.
// It should always be called when creating a new blockchain(including when cutting).
func (bc *Blockchain) RegisterBlockchain() error {
	GlobalBlockchainsLength = GlobalBlockchainsLength + 1
	bc.ID = GlobalBlockchainsLength
	GlobalBlockchains = append(GlobalBlockchains, *bc)
	return nil
}

// CopyBlockchain copies blockchain.
func CopyBlockchain(bc Blockchain) error {
	nbc := bc
	nbc.RegisterBlockchain()
	return nil
}

// NewBlockchain creates blockchain.
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
	LongestBlockchainUpdate(bc)
	return nil
}

// Block (n int) returns the n-th block.
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
	fmt.Fprintf(res, "Genesis Block \n%v\n", bc.GenesisBlock.String())
	fmt.Fprintf(res, "Candidate Block \n%v", bc.CandidateBlock.String())
	return res.String()
}
