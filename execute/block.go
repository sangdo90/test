package execute

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

// BlockchainCommands is ...
func BlockCommands() {
	_ = command.AddCommand("", command.Command{
		Name:        "block",
		ShortName:   "b",
		Description: "manage blocks",
		Commands: []command.Command{
			command.Command{
				Name:          "new",
				Description:   "create a candidate block",
				Commands:      make([]command.Command, 0),
				Flags:         nil,
				DefaultParams: []interface{}{uint64(1)},
				Run:           NewCandidateBlock,
			},
			command.Command{
				Name:          "attach",
				Description:   "Attach candidate block to blockchain",
				Commands:      make([]command.Command, 0),
				Flags:         nil,
				DefaultParams: []interface{}{uint64(1)},
				Run:           AttachCandidateBlockToBlockchain,
			},
			command.Command{
				Name:        "list",
				ShortName:   "ls",
				Description: "show list of blocks",
				Commands:    make([]command.Command, 0),
				Flags:       nil,
				Run:         ShowBlocksList,
			},
			command.Command{
				Name:          "info",
				Description:   "show information of block",
				Commands:      make([]command.Command, 0),
				Flags:         nil,
				DefaultParams: []interface{}{uint64(1), uint64(0)},
				Run:           ShowBlockInformation,
			},
		},
		Flags: nil,
		Run:   nil,
	})
}

// ShowBlocksList shows list of block
// ''ShowBlocksList(bcid uint64)''
func ShowBlocksList(bcid uint64) error {
	log.Debug("Show Blocks List")
	log.Info(perforatedLine)

	bc, err := getBlockchain(bcid)

	if err != nil {
		return err
	}

	log.Info("Blockchain ID : " + strconv.FormatUint(bcid, 10))

	for idx, b := range bc.Blocks {
		i := strconv.Itoa(idx)
		log.Info(blockStringInfo(&b, "Block Index : "+i))
		log.Info(perforatedLine)
	}

	return nil
}

// ShowBlockInformation shows information of block
// ''ShowBlockInformation(bcid uint64, bidx uint64)''
func ShowBlockInformation(bcid uint64, bidx uint64) error {
	log.Debug("Show Block Information")
	log.Info(perforatedLine)

	bcids := strconv.FormatUint(bcid, 10)
	bidxs := strconv.FormatUint(bidx, 10)
	bc, err := getBlockchain(bcid)

	if err != nil {
		return err
	}

	if bidx > bc.BlockchainHeight-1 {
		return errors.New("Incorrect block index")
	}

	b := &bc.Blocks[bidx]

	if bidx == 0 {
		bidxs = "Candidate"
		b = bc.CandidateBlock
	}

	log.Info(blockStringInfo(b, "Blockchain ID : "+bcids+"\tBlock Index : "+bidxs))
	log.Info(perforatedLine)

	return nil
}

// NewCandidateBlock creates a new candidate block into a blockchain identified by a ID.
// Therefore, NewCandidateBlock requires a blockchain ID
// ''NewCandidateBlock(bcid uint64)''
func NewCandidateBlock(bcid uint64) error {
	log.Debug("Create New Candidate Block")
	log.Info(perforatedLine)

	bcids := strconv.FormatUint(bcid, 10)
	bc, err := getBlockchain(bcid)

	if err != nil {
		return err
	}

	bc.CandidateBlock = core.NewBlock(&bc.Blocks[bc.BlockchainHeight-1])

	log.Info(blockStringInfo(bc.CandidateBlock, "Blockchain ID : "+bcids+"'s Candidate Block"))
	log.Info(perforatedLine)
	log.Debug("Create completed")
	return nil
}

// AttachCandidateBlockToBlockchain attach candidate block into a blockchain identified by a ID.
// Therefore, AttachCandidateBlockToBlockchain requires a blockchain ID
// ''AttachCandidateBlockToBlockchain(bcid uint64)''
func AttachCandidateBlockToBlockchain(bcid uint64) error {
	log.Debug("Attach Candidate Block to Blockchain")
	log.Info(perforatedLine)

	bcids := strconv.FormatUint(bcid, 10)
	log.Debug("Blockchain ID : " + bcids)
	bc, err := getBlockchain(bcid)

	if err != nil {
		return err
	}

	bc.AddBlock()
	log.Info(perforatedLine)
	log.Debug("Attach completed")
	return nil
}

// BlockStringInfo provides information(string) about the block.
func blockStringInfo(b *core.Block, title string) string {
	buffer := bytes.NewBuffer([]byte{})
	fmt.Fprintf(buffer, "PreviousHash     %v\n", b.Header.PreviousHash)
	fmt.Fprintf(buffer, "Timestamp        %v\n", b.Header.Timestamp)
	fmt.Fprintf(buffer, "Index            %v\n", b.Header.Index)
	fmt.Fprintf(buffer, "Transactions     %v\n", len(b.Body.Transactions))

	res := title + "\n" + buffer.String()
	return res
}
