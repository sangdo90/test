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

// BlockCommands contains block commands.
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
// ''ShowBlocksList(bcidx uint64)''
func ShowBlocksList(bcidx uint64) error {
	log.Debug("Show Blocks List")
	log.Info(perforatedLine)

	bc, err := getBlockchain(bcidx)

	if err != nil {
		return err
	}

	log.Info("Blockchain index : " + strconv.FormatUint(bcidx, 10))

	for idx, b := range bc.Blocks {
		i := strconv.Itoa(idx)
		log.Info(blockStringInfo(&b, "Block Index : "+i))
		log.Info(perforatedLine)
	}

	return nil
}

// ShowBlockInformation shows information of block
// ''ShowBlockInformation(bcidx uint64, bidx uint64)''
func ShowBlockInformation(bcidx uint64, bidx uint64) error {
	log.Debug("Show Block Information")
	log.Info(perforatedLine)

	bcidxs := strconv.FormatUint(bcidx, 10)
	bidxs := strconv.FormatUint(bidx, 10)
	bc, err := getBlockchain(bcidx)

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

	log.Info(blockStringInfo(b, "Blockchain index : "+bcidxs+"\tBlock Index : "+bidxs))
	log.Info(perforatedLine)

	return nil
}

// NewCandidateBlock creates a new candidate block into a blockchain identified by a index.
// Therefore, NewCandidateBlock requires a blockchain index
// ''NewCandidateBlock(bcidx uint64)''
func NewCandidateBlock(bcidx uint64) error {
	log.Debug("Create New Candidate Block")
	log.Info(perforatedLine)

	bcidxs := strconv.FormatUint(bcidx, 10)
	bc, err := getBlockchain(bcidx)

	if err != nil {
		return err
	}

	bc.CandidateBlock = core.NewBlock(&bc.Blocks[bc.BlockchainHeight-1])

	log.Info(blockStringInfo(bc.CandidateBlock, "Blockchain index : "+bcidxs+"'s Candidate Block"))
	log.Info(perforatedLine)
	log.Debug("Create completed")
	return nil
}

// AttachCandidateBlockToBlockchain attach candidate block into a blockchain identified by a index.
// Therefore, AttachCandidateBlockToBlockchain requires a blockchain index
// ''AttachCandidateBlockToBlockchain(bcidx uint64)''
func AttachCandidateBlockToBlockchain(bcidx uint64) error {
	log.Debug("Attach Candidate Block to Blockchain")
	log.Info(perforatedLine)

	bcidxs := strconv.FormatUint(bcidx, 10)
	log.Debug("Blockchain index : " + bcidxs)
	bc, err := getBlockchain(bcidx)

	if err != nil {
		return err
	}

	_ = bc.AddBlock()
	log.Info(perforatedLine)
	log.Debug("Attach completed")
	return nil
}

// blockStringInfo provides information(string) about the block.
func blockStringInfo(b *core.Block, title string) string {
	buffer := bytes.NewBuffer([]byte{})
	if b != nil {
		ph := ""
		for _, v := range b.Header.PreviousHash {
			ph += fmt.Sprintf("%02x", v)
		}
		fmt.Fprintf(buffer, "PreviousHash     %v\n", ph)
		fmt.Fprintf(buffer, "Timestamp        %v\n", b.Header.Timestamp)
		fmt.Fprintf(buffer, "Index            %v\n", b.Header.Index)
		fmt.Fprintf(buffer, "Transactions     %v\n", len(b.Body.Transactions))
		fmt.Fprintf(buffer, "%v", transactionsString(b, ""))
	}

	res := title + "\n" + buffer.String()
	return res
}
