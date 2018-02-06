package execute

import (
	"errors"
	"fmt"
	"strconv"

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
				Name:        "info",
				Description: "show blockchain Information",
				Commands: []command.Command{
					command.Command{
						Name:        "all",
						ShortName:   "a",
						Description: "show blockchains Information All",
						Commands:    make([]command.Command, 0),
						Flags:       nil,
						Run:         ShowBlockchainsInformationAll,
					},
				},
				Flags:         nil,
				DefaultParams: []interface{}{uint64(1)},
				Run:           ShowBlockchainInformation,
			},
			command.Command{
				Name:        "block",
				ShortName:   "b",
				Description: "manage blocks",
				Commands: []command.Command{
					command.Command{
						Name:        "info",
						Description: "show information of block",
						Commands:    make([]command.Command, 0),
						Flags:       nil,
						Run:         ShowBlockInformation,
					},
					command.Command{
						Name:        "list",
						ShortName:   "ls",
						Description: "show list of blocks",
						Commands:    make([]command.Command, 0),
						Flags:       nil,
						Run:         ShowBlocksList,
					},
				},
				Flags: nil,
				Run:   nil,
			},
			command.Command{
				Name:        "cblock",
				ShortName:   "cb",
				Description: "manage candidate block",
				Commands: []command.Command{
					command.Command{
						Name:        "new",
						Description: "create a candidate block",
						Commands:    make([]command.Command, 0),
						Flags:       nil,
						Run:         NewCandidateBlock,
					},
					command.Command{
						Name:        "attach",
						Description: "Attach candidate block to blockchain",
						Commands:    make([]command.Command, 0),
						Flags:       nil,
						Run:         AttachCandidateBlockToBlockchain,
					},
				},
				Flags: nil,
				Run:   nil,
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

	bc := core.NewBlockchain()
	core.AppendBlockchain(bc)

	log.Debug("Create completed")

	return nil
}

// ShowBlockchainsList shows list of blockchains
// ''ShowBlockchainsList()''
func ShowBlockchainsList() error {
	log.Debug("Show Blockchains List")

	result := ""
	for _, bc := range core.GlobalBlockchains {
		result += fmt.Sprintf("%v ", bc.ID)
	}

	log.Info(result)
	log.Debug(perforatedLine)
	return nil
}

// ShowBlockchainInformation shows information of blockchain identified by a ID.
// Therefore, ShowBlockchainInformation requires a blockchain ID
// ''ShowBlockchainInformation(bcid uint64)''
func ShowBlockchainInformation(bcid uint64) error {
	log.Debug("Show Blockchain Information")
	log.Debug(perforatedLine)

	bc, err := core.SelectBlockchain(bcid)
	if err != nil {
		return err
	}

	log.Debug(bc.String())
	log.Debug(perforatedLine)

	return nil
}

// ShowBlockchainsInformationAll shows inforamtion of all blockchains
// ''ShowBlockchainsInformationAll()''
func ShowBlockchainsInformationAll() error {
	log.Debug("Show Blockchains Information All")
	log.Debug(perforatedLine)

	for _, bc := range core.GlobalBlockchains {
		log.Debug(bc.String())
		log.Debug(perforatedLine)
	}

	return nil
}

// ShowBlocksList shows list of block
// ''ShowBlocksList(bcid uint64)''
func ShowBlocksList(bcid uint64) error {
	log.Debug("Show Blocks List")
	log.Debug(perforatedLine)

	bc, err := core.SelectBlockchain(bcid)
	if err != nil {
		return err
	}

	log.Debug("Blockchain ID : " + strconv.FormatUint(bcid, 10))
	for idx, b := range bc.Blocks {
		i := strconv.Itoa(idx)
		log.Debug(b.String("Block Index : " + i))
		log.Debug(perforatedLine)
	}

	return nil
}

// ShowBlockInformation shows information of block
// ''ShowBlockInformation(bcid uint64, bidx uint64)''
func ShowBlockInformation(bcid uint64, bidx uint64) error {
	log.Debug("Show Block Information")
	log.Debug(perforatedLine)

	bcids := strconv.FormatUint(bcid, 10)
	bidxs := strconv.FormatUint(bidx, 10)

	bc, err := core.SelectBlockchain(bcid)
	if err != nil {
		return nil
	}

	if bidx > bc.BlockchainHeight-1 {
		return errors.New("Incorrect block index")
	}
	b := bc.Block(bidx)

	if bidx == 0 {
		bidxs = "Candidate"
		b = bc.CandidateBlock
	}

	log.Debug(b.String("Blockchain ID : " + bcids + "\tBlock Index : " + bidxs))
	log.Debug(perforatedLine)

	return nil
}

// NewCandidateBlock creates a new candidate block into a blockchain identified by a ID.
// Therefore, NewCandidateBlock requires a blockchain ID
// ''NewCandidateBlock(bcid uint64)''
func NewCandidateBlock(bcid uint64) error {
	log.Debug("Create New Candidate Block")
	log.Debug(perforatedLine)

	bcids := strconv.FormatUint(bcid, 10)
	bc, err := core.SelectBlockchain(bcid)
	if err != nil {
		return err
	}

	bc.CandidateBlock = core.NewBlock(core.GetLastestBlock(bcid))

	log.Debug(bc.CandidateBlock.String("Blockchain ID : " + bcids + "'s Candidate Block"))
	log.Debug("Create completed")
	log.Debug(perforatedLine)
	return nil
}

// AttachCandidateBlockToBlockchain attach candidate block into a blockchain identified by a ID.
// Therefore, AttachCandidateBlockToBlockchain requires a blockchain ID
// ''AttachCandidateBlockToBlockchain(bcid uint64)''
func AttachCandidateBlockToBlockchain(bcid uint64) error {
	log.Debug("Attach Candidate Block to Blockchain")
	log.Debug(perforatedLine)

	bcids := strconv.FormatUint(bcid, 10)
	log.Debug("Blockchain ID : " + bcids)
	bc, err := core.SelectBlockchain(bcid)

	if err != nil {
		return err
	}

	bc.AddBlock()
	log.Debug("Attach completed")
	log.Debug(perforatedLine)

	return nil
}
