package pow

import (
	"fmt"

	"github.com/smartm2m/blockchain/core"
)

type PoW struct {
}

func New() *PoW {
	return nil
}

func Mining(b *core.Block) (bool, uint64) {
	diff := b.Header.Difficulty
	fmt.Println("PoW Start")
	i := 1
	for diff[0] < core.BlockHeaderHash(b.Header)[0] {
		fmt.Println("Try: ", i)
		fmt.Println("Difficulty: ", diff[0])
		fmt.Println("Block Hash: ", core.BlockHeaderHash(b.Header)[0])
		b.Header.Nonce = b.Header.Nonce + 1
		i = i + 1
	}
	fmt.Println("Try: ", i)
	fmt.Println("Difficulty: ", diff[0])
	fmt.Println("Block Hash: ", core.BlockHeaderHash(b.Header)[0])
	return true, b.Header.Nonce
}
