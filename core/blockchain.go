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

var GlobalBlockchains []BlockChain

type BlockChain struct {
	ID               uint64
	Blocks           []Block
	BlockChainHeight uint64
	GenesisBlock     *Block
	CurrentBlock     *Block
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//end
func NewGenesisBlock() *Block {
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

func NewBlockChain() *BlockChain {
	b := NewGenesisBlock()

	bc := &BlockChain{
		ID:               rand.Uint64(),
		Blocks:           []Block{*b},
		BlockChainHeight: 1,
		GenesisBlock:     b,
		CurrentBlock:     b,
	}
	return bc
}

//target index..
//-> need to changed Blockchain Struct
//-> or, need to copy front Blockchain, then attach the block
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

// TODO: Getting the last block from a blockchain
// TODO: Getting the i-th block.
func (bc *BlockChain) Block(n int) console.Blocker {
	if n < 0 || n >= len(bc.Blocks) || len(bc.Blocks) < 1 {
		return nil
	}

	return &bc.Blocks[n]
}

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
