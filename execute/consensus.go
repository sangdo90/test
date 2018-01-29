package execute

import (
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

// ConsensusCommands register commands for consensus.
func ConsensusCommands() {
	_ = command.AddCommand("", command.Command{
		Name:        "consensus",
		Description: "manage consensus",
		Commands: []command.Command{
			command.Command{
				Name:        "perform",
				Description: "perform a consensus algorithm.",
				Commands:    nil,
				Flags:       nil,
				Run:         PerformConsensus,
			},
		},
		Flags: nil,
		Run: func(args []string) error {
			log.Debug("Consensus commands")

			return nil
		},
	})
}

// PerformConsensus performs consensus for blockchains.
func PerformConsensus(args []string) error {
	log.Debug("consensus perform")
	return nil
}
