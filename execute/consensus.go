package execute

import (
	"github.com/smartm2m/blockchain/consensus"
	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

// ConsensusCommands register commands for consensus.
func ConsensusCommands() {
	_ = command.AddCommand("blockchain", command.Command{
		Name:        "copy",
		ShortName:   "cp",
		Description: ": copy a new blockchain from existing blockchain",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         CopyBlockchain,
	})
	_ = command.AddCommand("", command.Command{
		Name:        "consensus",
		ShortName:   "cs",
		Description: ": consent then choose a blockchain from different blockchains",
		Commands: []command.Command{
			command.Command{
				Name:        "execute",
				ShortName:   "exec",
				Description: ": execute consensus",
				Commands:    make([]command.Command, 0),
				Flags:       nil,
				Run:         Consensus,
			},
		},
		Flags: nil,
		Run:   nil,
	})
}

// CopyBlockchain duplicates a existing blockchain.
func CopyBlockchain(bcidx uint64) error {
	log.Debug("CopyBlockchain")
	bc, _ := getBlockchain(bcidx)
	blocks := make([]core.Block, len(bc.Blocks))
	copy(blocks, bc.Blocks)
	nbc := &core.Blockchain{
		Blocks:           blocks,
		BlockchainHeight: bc.BlockchainHeight,
		GenesisBlock:     bc.GenesisBlock,
		CandidateBlock:   nil,
	}

	core.AppendBlockchain(nbc)

	return nil
}

// Consensus chooses a blockchain among discovered blockchains.
func Consensus() error {
	log.Debug("ExecuteConsensus")
	var err error
	core.GlobalBlockchains, err = consensus.Consensus(core.GlobalBlockchains)

	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info(blockchainStringInfo(core.GlobalBlockchains[0], "The longest blockchain"))
	}

	return nil
}
