package core

import (
	"math/rand"
	"time"
	"../common"
)

var GlobalBlockchains []BlockChain

type BlockChain struct {
	ID    			 	uint64
	Blocks 				[]Block
	BlockChainHeight	uint64
	GenesisBlock 		*Block
	CurrentBlock		*Block
}

//end
func NewGenesisBlock() *Block {
	
	rand.Seed(time.Now().UTC().UnixNano())
	//token := make([]byte, 32)
	//rand.Read(token)
	token := new(common.Hash)
	rand.Read(token[:])

	b := &Block{
		Header: 
			BlockHeader{
				PreviousHash	: common.SHA2Hash(token[:]),
				//MerkleRootHash	: GetMerkleRootHash(transactions),
				Difficulty 		: 0,
				Nonce 			: 0,
				Timestamp 		: common.MakeTimestamp(),					
				Index 			: 0,
			},
		Body:
			BlockBody{
				//Transactions : append(Transactions, NewTransaction(/*Parameters*/)),
			},
	}

	return b
}

func NewBlockChain(id uint64) *BlockChain {
	b := NewGenesisBlock()
	bc := &BlockChain{
		ID : id,
		Blocks : []Block{*b},
		BlockChainHeight : 1,
		GenesisBlock : b,
		CurrentBlock : b,
	}
	return bc
}

//target index.. 
//-> need to changed Blockchain Struct
//-> or, need to copy front Blockchain, then attach the block
func (bc *BlockChain) AddBlock(b *Block) error {
	bc.Blocks = append(bc.Blocks, *b)
	bc.BlockChainHeight = bc.BlockChainHeight + 1
	bc.CurrentBlock = b
	
	return nil
}
