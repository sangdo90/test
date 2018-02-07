package execute

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

const perforatedLine string = "\n-----------------------------------------------------\n"

// BlockchainCommands is ...
func BlockchainCommands() {
	_ = command.AddCommand("", command.Command{
		Name:        "blockchain",
		ShortName:   "bc",
		Description: "manage blockchains",
		Commands: []command.Command{
			command.Command{
				Name:        "new",
				Description: "create a blockchain",
				Commands:    make([]command.Command, 0),
				Flags:       nil,
				Run:         NewBlockchain,
			},
			command.Command{
				Name:        "list",
				ShortName:   "ls",
				Description: "show blockchains list",
				Commands:    make([]command.Command, 0),
				Flags:       nil,
				Run:         ShowBlockchainsList,
			},
			command.Command{
				Name:          "info",
				Description:   "show blockchain Information",
				Commands:      nil,
				Flags:         nil,
				DefaultParams: []interface{}{uint64(1)},
				Run:           ShowBlockchainInformation,
			},
		},
		Flags: nil,
		Run:   nil,
	})
}

// NewBlockchain creates a new blockchain containing genesisblock
// ''NewBlockchain()''
func NewBlockchain() error {
	log.Debug("Create New Blockchain")
	log.Info(perforatedLine)

	bc := core.NewBlockchain()
	core.AppendBlockchain(bc)

	log.Debug("Create completed")

	return nil
}

// ShowBlockchainsList shows list of blockchains
// ''ShowBlockchainsList()''
func ShowBlockchainsList() error {
	log.Debug("Show Blockchains List")
	log.Info(perforatedLine)

	result := ""
	for _, bc := range core.GlobalBlockchains {
		result += fmt.Sprintf("%v ", bc.ID)
	}

	log.Info(result)
	log.Info(perforatedLine)

	return nil
}

// ShowBlockchainInformation shows information of blockchain identified by a ID.
// Therefore, ShowBlockchainInformation requires a blockchain ID
// ''ShowBlockchainInformation(bcid uint64)''
func ShowBlockchainInformation(bcid uint64) error {
	log.Debug("Show Blockchain Information")
	log.Info(perforatedLine)
	bc, err := getBlockchain(bcid)

	if err != nil {
		return err
	}

	log.Info(blockchainStringInfo(bc, ""))
	log.Info(perforatedLine)

	return nil
}

// getBlockchain gets blockchain
func getBlockchain(bcid uint64) (*core.Blockchain, error) {
	if bcid > uint64(len(core.GlobalBlockchains)) {
		return nil, errors.New("Invalid Select Blockchain")
	}

	return core.GlobalBlockchains[bcid-1], nil
}

// blockchainStringInfo provides information(string) about the blockchain.
func blockchainStringInfo(bc *core.Blockchain, title string) string {
	buffer := bytes.NewBuffer([]byte{})
	fmt.Fprintf(buffer, "\nID     %v\n", bc.ID)
	fmt.Fprintf(buffer, "Height %v\n\n", bc.BlockchainHeight)
	fmt.Fprintf(buffer, "Genesis Block \n%v\n", blockStringInfo(bc.GenesisBlock, "Genesis"))
	fmt.Fprintf(buffer, "Candidate Block \n%v", blockStringInfo(bc.CandidateBlock, "Candidate"))

	res := title + "\n" + buffer.String()
	return res
}
