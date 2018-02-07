package core

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/smartm2m/blockchain/common"
)

// A BlockHeader is the header of block, which contains the information of the block.
type BlockHeader struct {
	PreviousHash   common.Hash
	MerkleRootHash common.Hash
	Difficulty     common.Hash
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
	//buffer.WriteString(strconv.FormatUint(bh.Difficulty, 10))
	buffer.WriteString(common.HashToString(bh.Difficulty))
	buffer.WriteString(strconv.FormatUint(bh.Nonce, 10))
	buffer.WriteString(strconv.FormatInt(bh.Timestamp, 10))
	buffer.WriteString(strconv.FormatUint(bh.Index, 10))

	var s = common.StringToHash(buffer.String())
	return common.SHA2Hash(s[:])
}

// MerkleRootHash is ...
func MerkleRootHash(t *Transaction) common.Hash {
	return common.SHA2Hash(nil)
}

// GetLastestBlock gest lastest block in blockchain identified by a ID.
// Therefore, GetLastestBlock requires a blockchain ID.
func GetLastestBlock(bcid uint64) *Block {
	bc, _ := SelectBlockchain(bcid)
	return &bc.Blocks[bc.BlockchainHeight-1]
}

// NewBlock creates a new block.
func NewBlock(pb *Block) *Block {
	diff := new(common.Hash)
	diff[0] = 10
	b := &Block{
		Header: BlockHeader{
			PreviousHash: BlockHeaderHash(pb.Header),
			//MerkleRootHash: MerkleRootHash(t),
			Difficulty: *diff, // need to static variable diff
			Nonce:      0,     // need to static variable nonce
			Timestamp:  common.MakeTimestamp(),
			Index:      pb.Header.Index + 1,
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
	return nil
}

// String (block) function provides information about the block.
func (b *Block) String(name string) string {

	buffer := bytes.NewBuffer([]byte{})
	fmt.Fprintf(buffer, "PreviousHash     %v\n", b.Header.PreviousHash)
	fmt.Fprintf(buffer, "MerkleRootHash   %v\n", b.Header.MerkleRootHash)
	fmt.Fprintf(buffer, "Difficulty       %v\n", b.Header.Difficulty)
	fmt.Fprintf(buffer, "Nonce            %v\n", b.Header.Nonce)
	fmt.Fprintf(buffer, "Timestamp        %v\n", b.Header.Timestamp)
	fmt.Fprintf(buffer, "Index            %v\n", b.Header.Index)
	fmt.Fprintf(buffer, "Transactions     %v\n", len(b.Body.Transactions))

	res := name + "\n" + buffer.String()
	return res
}

// TransactionsString returns string of transactions infomartion
func (b *Block) TransactionsString(name string) string {
	res := name + "\nNum\tAmount\tFrom\tTo\n"
	buffer := bytes.NewBuffer([]byte{})

	for idx, t := range b.Body.Transactions {
		fmt.Fprintf(buffer, strconv.Itoa(idx+1))
		fmt.Fprintf(buffer, "\t%v", t.Data.Amount)
		fmt.Fprintf(buffer, "\t%v", t.From)
		fmt.Fprintf(buffer, "\t%v\n", t.Data.To)
	}

	res = res + buffer.String()
	return res
}
