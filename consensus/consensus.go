package consensus

import (
	"errors"

	"github.com/smartm2m/blockchain/core"
)

// Consensus deduces the longest blockchain among discovered blockchains.
func Consensus(bcList []*core.Blockchain) ([]*core.Blockchain, error) {
	longestChain := 0
	var err error
	conflicted := errors.New("Cannot decide the longest blockchain")

	for i := 1; i < len(bcList); i++ {
		if bcList[longestChain].BlockchainHeight < bcList[i].BlockchainHeight {
			longestChain = i
			err = nil
		} else if bcList[longestChain].BlockchainHeight == bcList[i].BlockchainHeight {
			err = conflicted
		}
	}

	if err != nil {
		return bcList, err
	}

	return bcList[longestChain : longestChain+1], nil
}
