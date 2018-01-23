package core

import (
	"strconv"
	"../common"
	"bytes"
//	"github.com/smartm2m/blockchain/common"
)

type BlockHeader struct {
	PreviousHash 	common.Hash
	MerkleRootHash	common.Hash
	Difficulty		uint64
	Nonce			uint64
	Timestamp		int64
	Index			uint64
}

type BlockBody struct {
	Transactions []*Transaction
}

type Block struct {
	Header	BlockHeader
	Body	BlockBody
}

func BlockHeaderHash(bh BlockHeader) common.Hash{
	var buffer bytes.Buffer	
	buffer.WriteString(common.HashToString(bh.PreviousHash))
	buffer.WriteString(common.HashToString(bh.MerkleRootHash))
	buffer.WriteString(strconv.Itoa(int(bh.Difficulty)))
	buffer.WriteString(strconv.Itoa(int(bh.Nonce)))
	buffer.WriteString(strconv.Itoa(int(bh.Timestamp)))
	buffer.WriteString(strconv.Itoa(int(bh.Index)))

	var s = common.StringToHash(buffer.String())
	return common.SHA2Hash(s[:])
}

func GetMerkleRootHash(t *Transaction) common.Hash {
	//
	//
	return common.SHA2Hash(nil)
}

//cb = currnet block, The name needs to be changed
func NewBlock(cb *Block, t *Transaction) *Block {
	b := &Block{
		Header: 
			BlockHeader{
				PreviousHash	: BlockHeaderHash(cb.Header),
				//MerkleRootHash	: GetMerkleRootHash(t),
				Difficulty 		: 0, // need to static variable diff
				Nonce 			: 0, // need to static variable nonce
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
	b.Body.Transactions = append(b.Body.Transactions, t)
	return nil
}