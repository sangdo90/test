package core

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/smartm2m/blockchain/common"
	"github.com/smartm2m/chainutil/console"
)

//GlobalBlockChains is set of all block chains.
var GlobalBlockChains []BlockChain

//GlobalBlockChainsLength is Length of GlobalBlockchains
var GlobalBlockChainsLength uint64

//BlockChain is chain of blocks, consisting of ID, Blocks, Height, Genesisblcok, and CurrentBlock.
//ID is the same as index+1
type BlockChain struct {
	ID               uint64
	Blocks           []Block
	BlockChainHeight uint64
	GenesisBlock     *Block
	CurrentBlock     *Block
}

//SetSeedUsingTime sets the seed using time.
func SetSeedUsingTime() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//NewGenesisBlock creates genesis block.
func NewGenesisBlock() *Block {
	SetSeedUsingTime()
	//token := make([]byte, 32)
	//rand.Read(token)
	token := new(common.Hash)
	rand.Read(token[:])

	b := &Block{
		Header: BlockHeader{
			PreviousHash: common.SHA2Hash(token[:]),
			//MerkleRootHash	: GetMerkleRootHash(transactions),
			Difficulty: 0,
			Nonce:      0,
			Timestamp:  common.MakeTimestamp(),
			Index:      0,
		},
		Body: BlockBody{
		//Transactions : append(Transactions, NewTransaction(/*Parameters*/)),
		},
	}

	return b
}

//SelectBlockChain returns blockchain that has the input id.
func SelectBlockChain(id uint64) (*BlockChain, error) {
	if id == 0 || id > GlobalBlockChainsLength {
		return nil, errors.New("Invalid Select Blockchain")
	}
	return &GlobalBlockChains[id-1], nil
}

//RegisterBlockChain registers in the global blockchain.
//It should always be called when creating a new blockchain(including when cutting).
func (bc *BlockChain) RegisterBlockChain() error {
	GlobalBlockChains = append(GlobalBlockChains, *bc)
	GlobalBlockChainsLength = GlobalBlockChainsLength + 1
	bc.ID = GlobalBlockChainsLength

	return nil
}

//CutBlockChain cuts blockchain.
func CutBlockChain(bc BlockChain, idx uint64) *BlockChain {
	bc.ID = 0 //init value
	bc.Blocks = bc.Blocks[:idx]
	bc.BlockChainHeight = idx
	bc.CurrentBlock = &bc.Blocks[idx]

	return &bc
}

//NewBlockChain creates blockchain.
func NewBlockChain() *BlockChain {
	b := NewGenesisBlock()

	bc := &BlockChain{
		ID:               0, //init value
		Blocks:           []Block{*b},
		BlockChainHeight: 1, //uint64(len([]Block{*b})), //always 1
		GenesisBlock:     b,
		CurrentBlock:     b,
	}
	return bc
}

//AddBlock adds a block to blockchain.
func (bc *BlockChain) AddBlock(blk console.Blocker) error {
	b, ok := blk.Block().(*Block)
	if !ok {
		return errors.New("Invalid block")
	}
	bc.Blocks = append(bc.Blocks, *b)
	bc.BlockChainHeight = bc.BlockChainHeight + 1
	bc.CurrentBlock = b

	return nil
}

// Block (n int) returns the n-th block.
// TODO: Getting the last block from a blockchain
// TODO: Getting the i-th block.
func (bc *BlockChain) Block(n int) console.Blocker {
	if n < 0 || n >= len(bc.Blocks) || len(bc.Blocks) < 1 {
		return nil
	}

	return &bc.Blocks[n]
}

//String (blockchain) function provides information about the blockchain.
// TODO: Convert to string from a blockchain
func (bc *BlockChain) String() string {
	res := bytes.NewBuffer([]byte{})
	fmt.Fprintf(res, "ID     %v\n", bc.ID)
	fmt.Fprintf(res, "Height %v\n", bc.BlockChainHeight)
	fmt.Fprintf(res, "Blocks: \n")
	fmt.Fprintln(res, "-----------------------------------------------------")
	for i, b := range bc.Blocks {
		fmt.Fprintf(res, "%v-th Block: \n", i)
		fmt.Fprintf(res, "%s", b.String())
		fmt.Fprintln(res, "")
	}
	fmt.Fprintln(res, "-----------------------------------------------------")

	return res.String()
}
