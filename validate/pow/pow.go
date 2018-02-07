package pow

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/chainutil/log"
)

// Mining verifies a block for attaching into a blockchain.
func Mining(b *core.Block) uint64 {
	diff := b.Header.Difficulty
	fmt.Println("PoW Start")
	hash := sha256.Sum256(b.Header.ToBytes())
	i := 1
	for ; greaterThan(hash, diff); hash, i = sha256.Sum256(b.Header.ToBytes()), i+1 {
		log.Debug("Try: " + strconv.Itoa(i))
		log.Debug("Difficulty: " + hashToString(diff))
		log.Debug("Block Hash: " + hashToString(hash))
		b.Header.Nonce = b.Header.Nonce + 1
		time.Sleep(300 * time.Millisecond)

	}
	log.Debug("Tried: " + strconv.Itoa(i))
	log.Debug("Difficulty: " + hashToString(diff))
	log.Debug("Block Hash: " + hashToString(hash))

	return b.Header.Nonce
}

func hashToString(h [32]byte) string {
	res := ""
	for _, v := range h {
		res += fmt.Sprintf("%02x", v)
	}

	return res
}

func greaterThan(l, r [32]byte) bool {
	for i := 0; i < 32; i++ {
		if l[i] > r[i] {
			return true
		} else if l[i] < r[i] {
			return false
		}
	}

	return false
}
