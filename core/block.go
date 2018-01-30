package core

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/smartm2m/blockchain/common"
	"github.com/smartm2m/chainutil/console"
)

// A BlockHeader is the header of block, which contains the information of the block.
type BlockHeader struct {
	PreviousHash   common.Hash
	MerkleRootHash common.Hash
	Difficulty     uint64
	Nonce          uint64
	Timestamp      int64
	Index          uint64
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

// BlockHeaderHash computes the hash value of the block header.
func BlockHeaderHash(bh BlockHeader) common.Hash {
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

// GetMerkleRootHash is ...
func GetMerkleRootHash(t *Transaction) common.Hash {
	return common.SHA2Hash(nil)
}

// GetLastestBlock gest lastest block in blockchain identified by a ID.
// Therefore, GetLastestBlock requires a blockchain ID.
func GetLastestBlock(bcid uint64) *Block {
	bc, _ := SelectBlockchain(bcid)
	return &bc.Blocks[bc.BlockchainHeight-1]
}

// NewBlock creates a new block.
func NewBlock(blk console.Blocker) *Block {
	bb, ok := blk.Block().(*Block)

	if !ok {
		return nil
	}

	b := &Block{
		Header: BlockHeader{
			PreviousHash: BlockHeaderHash(bb.Header),
			//MerkleRootHash	: GetMerkleRootHash(t),
			Difficulty: 0, // need to static variable diff
			Nonce:      0, // need to static variable nonce
			Timestamp:  common.MakeTimestamp(),
			Index:      bb.Header.Index + 1,
		},
		Body: BlockBody{
		//Transactions : append(Transactions, NewTransaction(/*Parameters*/)),
		},
	}
	return b
}

// AddTransaction adds a transaction.
func (b *Block) AddTransaction(t *Transaction) error {
	b.Body.Transactions = append(b.Body.Transactions, t)
	return nil
}

// Block function is an interface function that returns a block.
func (b *Block) Block() interface{} {
	return b
}

// BlockTransactions returns all transactions of block
func (b *Block) BlockTransactions() []*Transaction {
	return b.Body.Transactions
}

//String (block) function provides information about the block.
func (b *Block) String() string {
	res := bytes.NewBuffer([]byte{})
	fmt.Fprintf(res, "PreviousHash     %v\n", b.Header.PreviousHash)
	fmt.Fprintf(res, "MerkleRootHash   %v\n", b.Header.MerkleRootHash)
	fmt.Fprintf(res, "Difficulty       %v\n", b.Header.Difficulty)
	fmt.Fprintf(res, "Nonce            %v\n", b.Header.Nonce)
	fmt.Fprintf(res, "Timestamp        %v\n", b.Header.Timestamp)
	fmt.Fprintf(res, "Index            %v\n", b.Header.Index)
	fmt.Fprintf(res, "Transactions     %v\n", len(b.Body.Transactions))

	return res.String()
}
