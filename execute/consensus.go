package execute

import (
	"strconv"

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
				Run:         ExecuteConsensus,
			},
		},
		Flags: nil,
		Run:   nil,
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

// execute consensus
// ''ExecuteConsensus()''
func ExecuteConsensus() error {
	var maximum, eqaulnum, bcid uint64
	maximum = 0
	eqaulnum = 0
	bcid = 0

	for _, bc := range core.GlobalBlockchains {
		if bc.BlockchainHeight > maximum {
			maximum = bc.BlockchainHeight
			bcid = bc.ID
		}
	}
	for _, bc := range core.GlobalBlockchains {
		if bc.BlockchainHeight == maximum {
			eqaulnum = eqaulnum + 1
		}
	}

	if eqaulnum > 1 {
		log.Info("Chains are collide.")
	} else {
		log.Info("Choose blockchain ID" + strconv.FormatUint(bcid, 10))
		ChooseBlockchain(bcid)
	}

	return nil
}

func ChooseBlockchain(bcid uint64) error {
	bc, _ := core.SelectBlockchain(bcid)
	//nbcid := uint64(len(core.GlobalBlockchains))
	nbc := new(core.Blockchain)
	nbc = &core.Blockchain{
		ID:               0,
		Blocks:           bc.Blocks,
		BlockchainHeight: bc.BlockchainHeight,
		GenesisBlock:     bc.GenesisBlock,
		CandidateBlock:   bc.CandidateBlock,
		TotalAmount:      bc.TotalAmount,
	}
	core.GlobalBlockchains = core.GlobalBlockchains[:0]
	core.ChainsID = 0
	core.GlobalBlockchains = append(core.GlobalBlockchains, nbc)

	return nil
}
