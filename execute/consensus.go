package execute

import (
	"errors"
	"fmt"

	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

// ConsensusCommands register commands for consensus.
func ConsensusCommands() {
	_ = command.AddCommand("", command.Command{
		Name:        "consensus",
		ShortName:   "c",
		Description: "manage consensus",
		Commands: []command.Command{
			command.Command{
				Name:        "perform",
				ShortName:   "p",
				Description: "perform a consensus algorithm.",
				Commands:    nil,
				Flags:       nil,
				Run:         PerformConsensus,
			},
			command.Command{
				Name:        "execute",
				ShortName:   "e",
				Description: "execute a test.",
				Commands:    nil,
				Flags:       nil,
				Run:         Execution,
			},
		},
		Flags: nil,
		Run:   nil,
	})
}

// PerformConsensus performs consensus for blockchains.
func PerformConsensus(p1, p2, p3 string) error {
	res := fmt.Sprintf("PerformConsensus(%s,%s,%s)", p1, p2, p3)
	log.Debug(res)

	if p1 == "2" {
		return nil
	}

	return errors.New("asdf")
}

// Execution ...
func Execution(p1 string, p2 int, p3 uint64, p4 int64, p5 []byte) error {
	res := fmt.Sprintf("Execution(%s,%v,%v,%v,%v)", p1, p2, p3, p4, p5)
	log.Debug(res)
	return nil
}
