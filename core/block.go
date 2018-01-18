package core

import (
	"crypto/sha256"
	"../common"

//	"github.com/smartm2m/blockchain/common"
)

type BlockHeader struct {
	PreviousHash common.Hash
	MerkleRootHash common.Hash
	Difficulty uint64
	Nonce uint64
	Timestamp int64
	Index uint64
}

type BlockBody struct {
	Transactions []Transaction
}

type Block struct {
	Header BlockHeader
	Body BlockBody
}

func SHA2Hash(b []byte) [32]byte {
	return sha256.Sum256(b)
}

func BlockHash(b *Block) common.Hash{
	//dummy
	token := make([]byte, 32)
	return SHA2Hash(token)
}

func GetMerkleRootHash(t *Transaction) common.Hash {
	//
	//
	return SHA2Hash(nil)
}

//cb = currnet block, The name needs to be changed
func NewBlock(cb *Block, t *Transaction) *Block {
	b := &Block{
		Header: 
			BlockHeader{
				PreviousHash	: BlockHash(cb),
				//MerkleRootHash	: GetMerkleRootHash(t),
				Difficulty 		: 0,
				Nonce 			: 0,
				Timestamp 		: common.MakeTimestamp(),					
				Index 			: cb.Header.Index + 1,
			},
		Body:
			BlockBody{
				//Transactions : append(Transactions, NewTransaction(/*Parameters*/)),
			},
	}
	return b
}

func (b *Block) AddTransaction(t *Transaction) error {
	return nil
}