package core

import (
	"crypto/sha256"
	"encoding/binary"
	"time"

	"github.com/smartm2m/blockchain/validate"
)

// A BlockHeader is the header of block, which contains the information of the block.
type BlockHeader struct {
	PreviousHash   [32]byte
	MerkleRootHash [32]byte
	Difficulty     [32]byte
	Nonce          uint64
	Timestamp      int64
	Index          uint64
}

// ToBytes returns a slice of bytes of BlockHeader.
func (bh *BlockHeader) ToBytes() []byte {
	res := make([]byte, 0)

	tb := make([]byte, binary.MaxVarintLen64)
	ib := make([]byte, binary.MaxVarintLen64)
	nb := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(tb, bh.Timestamp)
	binary.PutUvarint(ib, bh.Index)
	binary.PutUvarint(nb, bh.Nonce)

	for _, b := range [][]byte{bh.PreviousHash[:], bh.MerkleRootHash[:], bh.Difficulty[:], tb, ib, nb} {
		res = append(res, b...)
	}

	return res
}

// A BlockBody is the body of block, which contains transactions.
type BlockBody struct {
	Transactions []*Transaction
}

// A Block is an element of blockchain, consisting of block header and block body.
type Block struct {
	Header BlockHeader
	Body   BlockBody
}

// NewBlock creates a new block.
func NewBlock(pb *Block) *Block {
	var diff [32]byte
	diff[0] = 10
	b := &Block{
		Header: BlockHeader{
			PreviousHash: sha256.Sum256(pb.Header.ToBytes()),
			Difficulty:   diff,
			Timestamp:    time.Now().UnixNano(),
			Index:        pb.Header.Index + 1,
		},
		Body: BlockBody{
			Transactions: nil,
		},
	}
	return b
}

// AddTransaction adds a transaction.
func (b *Block) AddTransaction(t *Transaction) error {
	b.Body.Transactions = append(b.Body.Transactions, t)
	tb := make([][]byte, 0)
	for _, v := range b.Body.Transactions {
		tb = append(tb, v.ToBytes())
	}
	b.Header.MerkleRootHash = validate.MerkleRootHash(tb)
	return nil
}
