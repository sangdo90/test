package execute

import (
	"fmt"

	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

// ConsensusCommands register commands for consensus.
func ConsensusCommands() {
	_ = command.AddCommand("blockchain", command.Command{
		Name:        "copy",
		ShortName:   "cp",
		Description: "copy a new blockchain from existing blockchain",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         CopyBlockchain,
	})
}

// Copy a new blockchain from existing blockchain.
// ''CopyBlockchain(bcid uint64)''
func CopyBlockchain(bcid uint64) error {
	bc, _ := core.SelectBlockchain(bcid)
	nbcid := uint64(len(core.GlobalBlockchains))
	nbc := new(core.Blockchain)
	nbc = &core.Blockchain{
		ID:               nbcid,
		Blocks:           bc.Blocks,
		BlockchainHeight: bc.BlockchainHeight,
		GenesisBlock:     bc.GenesisBlock,
		CandidateBlock:   bc.CandidateBlock,
		TotalAmount:      bc.TotalAmount,
	}
	nbc.RegisterBlockchain()

	return nil
}

// Execution ...
func Execution(p1 string, p2 int, p3 uint64, p4 int64, p5 []byte) error {
	res := fmt.Sprintf("Execution(%s,%v,%v,%v,%v)", p1, p2, p3, p4, p5)
	log.Debug(res)
	return nil
}
