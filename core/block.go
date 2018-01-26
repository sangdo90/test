package core

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/smartm2m/blockchain/common"
	"github.com/smartm2m/chainutil/console"
)

type BlockHeader struct {
	PreviousHash   common.Hash
	MerkleRootHash common.Hash
	Difficulty     uint64
	Nonce          uint64
	Timestamp      int64
	Index          uint64
}

type BlockBody struct {
	Transactions []*Transaction
}

type Block struct {
	Header BlockHeader
	Body   BlockBody
}

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

func GetMerkleRootHash(t *Transaction) common.Hash {
	return common.SHA2Hash(nil)
}

//cb = currnet block, The name needs to be changed
func NewBlock(blk console.Blocker) *Block {
	cb, ok := blk.Block().(*Block)

	if !ok {
		return nil
	}

	b := &Block{
		Header: BlockHeader{
			PreviousHash: BlockHeaderHash(cb.Header),
			//MerkleRootHash	: GetMerkleRootHash(t),
			Difficulty: 0, // need to static variable diff
			Nonce:      0, // need to static variable nonce
			Timestamp:  common.MakeTimestamp(),
			Index:      cb.Header.Index + 1,
		},
		Body: BlockBody{
		//Transactions : append(Transactions, NewTransaction(/*Parameters*/)),
		},
	}
	return b
}

func (b *Block) AddTransaction(t *Transaction) error {
	b.Body.Transactions = append(b.Body.Transactions, t)
	return nil
}

func (b *Block) Block() interface{} {
	return b
}

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
