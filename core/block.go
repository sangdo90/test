package core

import (
	"crypto/sha256"
	"encoding/binary"
	"time"
)

// A BlockHeader is the header of block, which contains the information of the block.
type BlockHeader struct {
	PreviousHash [32]byte
	Timestamp    int64
	Index        uint64
}

// ToBytes returns a slice of bytes of BlockHeader.
func (bh *BlockHeader) ToBytes() []byte {
	res := make([]byte, 0)

	tb := make([]byte, binary.MaxVarintLen64)
	ib := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(tb, bh.Timestamp)
	binary.PutUvarint(ib, bh.Index)

	for _, b := range [][]byte{bh.PreviousHash[:], tb, ib} {
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
	b := &Block{
		Header: BlockHeader{
			PreviousHash: sha256.Sum256(pb.Header.toBytes()),
			Timestamp:    time.Now().UnixNano(),
			Index:        pb.Header.Index + 1,
		},
		Body: BlockBody{
			Transactions: nil,
		},
	}
	return b
}
