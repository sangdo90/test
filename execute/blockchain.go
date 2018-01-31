package execute

import (
	"errors"
	"strconv"

	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

// BlockchainCommands registers commands for managing blockchains.
func BlockchainCommands() {

	_ = command.AddCommand("", command.Command{
		Name:        "blockchain",
		Description: "manage blockchains",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run: func(args []string) error {
			log.Debug("Blockchain commands")
			return nil
		},
	})

	_ = command.AddCommand("blockchain", command.Command{
		Name:        "new",
		Description: "create a blockchains",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         CreateBlockchain,
	})
}

// CreateBlockchain create a new blockchain.
func CreateBlockchain(args []string) error {
	log.Debug("CreateBlockchain")
	return nil
}

// NewBlockToBlockchain creates a block and then connect it into
// a blockchain identified by a ID. Therefore, NewBlockToBlockchain
// requires a blockchain ID as the first argument.
func NewBlockToBlockchain(args []string) error {
	if len(args) < 1 {
		return errors.New("Incorrect parameters")
	}
	bid, err := strconv.Atoi(args[0])

	if err != nil {
		return err
	}

	_ = bid

	return nil
}
