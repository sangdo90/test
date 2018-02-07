package validate

import (
	"crypto/sha256"

	"github.com/smartm2m/blockchain/core"
)

type TxContent struct {
	x string
}

//CalculateHash hashes the values of a TestContent
func (t TxContent) CalculateHash() []byte {
	h := sha256.New()
	h.Write([]byte(t.x))
	return h.Sum(nil)
}

//Equals tests for equality of two Contents
func (t TxContent) Equals(other Content) bool {
	return t.x == other.(TxContent).x
}

// Making MerkleTree when transaction be attached to block
func MerkleRootHash(b *core.Block) ([]byte, error) {
	var list []Content
	for _, txs := range b.Body.Transactions {
		list = append(list, TxContent{x: core.TransactionHash(txs)})
	}
	tree, _ := NewTree(list)
	mr := tree.MerkleRoot()

	//Verify a specific content in in the tree
	//vc := t.VerifyContent(tree.MerkleRoot(), list[0])
	//fmt.Println("Verify Content: ", vc)

	return mr, nil
}
