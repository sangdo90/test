package execute

import (
	"errors"
	"strconv"

	"github.com/smartm2m/blockchain/common"
	"github.com/smartm2m/blockchain/consensus/pow"
	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/blockchain/validate"
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

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
					command.Command{
						Name:        "transactions",
						ShortName:   "t",
						Description: "show list of a block's transactions",
						Commands:    make([]command.Command, 0),
						Flags:       nil,
						Run:         ShowBlockTransactionsList,
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
					command.Command{
						Name:        "transaction",
						ShortName:   "t",
						Description: "manage transaction of candidate block",
						Commands: []command.Command{
							command.Command{
								Name:        "new",
								Description: "create a transaction of candidate block",
								Commands:    make([]command.Command, 0),
								Flags:       nil,
								Run:         NewTransactionInCandidateBlock,
							},
							command.Command{
								Name:        "list",
								ShortName:   "ls",
								Description: "show transactions list of candidate block",
								Commands:    make([]command.Command, 0),
								Flags:       nil,
								Run:         ShowCandidateBlockTransactionsList,
							},
						},
						Flags: nil,
						Run:   nil,
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
	log.Debug(common.PerforatedLine)

	bc := core.AppendBlockchain()
	bc.RegisterBlockchain()
	mr, _ := validate.MerkleRootHash(bc.GenesisBlock)
	bc.GenesisBlock.Header.MerkleRootHash = common.BytesToHash(mr)
	log.Debug("Create completed")

	return nil
}

// ShowBlockchainsList shows list of blockchains
// ''ShowBlockchainsList()''
func ShowBlockchainsList() error {
	log.Debug("Show Blockchains List")
	log.Debug(common.PerforatedLine)

	var blockchainsList []uint64
	for _, bc := range core.GlobalBlockchains {
		blockchainsList = append(blockchainsList, bc.ID)
	}

	log.Debug(common.Uint64ArrayToString(blockchainsList, ", "))
	log.Debug(common.PerforatedLine)
	return nil
}

// ShowBlockchainInformation shows information of blockchain identified by a ID.
// Therefore, ShowBlockchainInformation requires a blockchain ID
// ''ShowBlockchainInformation(bcid uint64)''
func ShowBlockchainInformation(bcid uint64) error {
	log.Debug("Show Blockchain Information")
	log.Debug(common.PerforatedLine)

	bc, err := core.SelectBlockchain(bcid)
	if err != nil {
		return err
	}

	log.Debug(bc.String())
	log.Debug(common.PerforatedLine)

	return nil
}

// ShowBlockchainsInformationAll shows inforamtion of all blockchains
// ''ShowBlockchainsInformationAll()''
func ShowBlockchainsInformationAll() error {
	log.Debug("Show Blockchains Information All")
	log.Debug(common.PerforatedLine)

	for _, bc := range core.GlobalBlockchains {
		log.Debug(bc.String())
		log.Debug(common.PerforatedLine)
	}

	return nil
}

// ShowBlocksList shows list of block
// ''ShowBlocksList(bcid uint64)''
func ShowBlocksList(bcid uint64) error {
	log.Debug("Show Blocks List")
	log.Debug(common.PerforatedLine)

	bc, err := core.SelectBlockchain(bcid)
	if err != nil {
		return err
	}

	log.Debug("Blockchain ID : " + strconv.FormatUint(bcid, 10))
	for idx, b := range bc.Blocks {
		i := strconv.Itoa(idx)
		log.Debug(b.String("Block Index : " + i))
		log.Debug(common.PerforatedLine)
	}

	return nil
}

// ShowBlockInformation shows information of block
// ''ShowBlockInformation(bcid uint64, bidx uint64)''
func ShowBlockInformation(bcid uint64, bidx uint64) error {
	log.Debug("Show Block Information")
	log.Debug(common.PerforatedLine)

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
	log.Debug(common.PerforatedLine)

	return nil
}

// ShowBlockTransactionsList shows list of block's transactions
// ''ShowBlockTransactionsList(bcid uint64, bidx uint64)''
func ShowBlockTransactionsList(bcid uint64, bidx uint64) error {
	log.Debug("Show Block Transactions List")
	log.Debug(common.PerforatedLine)

	bcids := strconv.FormatUint(bcid, 10)
	bidxs := strconv.FormatUint(bidx, 10)

	bc, err := core.SelectBlockchain(bcid)
	if bidx > bc.BlockchainHeight-1 {
		return errors.New("Incorrect block index")
	}

	b := bc.Block(bidx)

	if bidx == 0 {
		bidxs = "Candidate"
		b = bc.CandidateBlock
	}

	if err != nil {
		return err
	}

	log.Debug(b.TransactionsString("Blockchain ID : " + bcids + "\tBlock Index : " + bidxs))
	log.Debug(common.PerforatedLine)
	return nil
}

// NewCandidateBlock creates a new candidate block into a blockchain identified by a ID.
// Therefore, NewCandidateBlock requires a blockchain ID
// ''NewCandidateBlock(bcid uint64)''
func NewCandidateBlock(bcid uint64) error {
	log.Debug("Create New Candidate Block")
	log.Debug(common.PerforatedLine)

	bcids := strconv.FormatUint(bcid, 10)
	bc, err := core.SelectBlockchain(bcid)
	if err != nil {
		return err
	}

	bc.CandidateBlock = core.NewBlock(core.GetLastestBlock(bcid))

	log.Debug(bc.CandidateBlock.String("Blockchain ID : " + bcids + "'s Candidate Block"))
	log.Debug("Create completed")
	log.Debug(common.PerforatedLine)
	return nil
}

// AttachCandidateBlockToBlockchain attach candidate block into a blockchain identified by a ID.
// Therefore, AttachCandidateBlockToBlockchain requires a blockchain ID
// ''AttachCandidateBlockToBlockchain(bcid uint64)''
func AttachCandidateBlockToBlockchain(bcid uint64) error {
	log.Debug("Attach Candidate Block to Blockchain")
	log.Debug(common.PerforatedLine)

	bcids := strconv.FormatUint(bcid, 10)
	log.Debug("Blockchain ID : " + bcids)
	bc, err := core.SelectBlockchain(bcid)

	if err != nil {
		return err
	}

	// pow
	check, nonce := pow.Mining(bc.CandidateBlock)
	if check {
		bc.AddBlock()
		log.Debug(nonce)
		log.Debug("Attach completed")
		log.Debug(common.PerforatedLine)
	} else {
		log.Debug("Fail to attach")
	}

	return nil
}

// NewTransactionInCandidateBlock creates a new transaction and put it in a candidate block of blockchain.
// Also, need to transaction that From A has sent To B.
// Therefore, NewTransactionInCandidateBlock requires a blockchain ID, From, To(address), Amount.
// ''NewTransactionInCandidateBlock(bcid uint64, amount uint64, from uint64, to *common.Address)''
func NewTransactionInCandidateBlock(bcid uint64, amount uint64, from uint64, to []byte) error {
	log.Debug("Create New Transaction in the Candidate block")
	log.Debug(common.PerforatedLine)
	bcids := strconv.FormatUint(bcid, 10)
	amounts := strconv.FormatUint(amount, 10)
	froms := strconv.FormatUint(from, 10)
	tos := common.BytesToString(to)
	toa := common.StringToAddress(tos)
	// tos bug
	log.Debug("Blockchain ID : " + bcids)

	bc, err := core.SelectBlockchain(bcid)

	if err != nil {
		return err
	}

	t := core.NewTransaction(from, &toa, amount)
	bc.CandidateBlock.AddTransaction(t)

	mr, _ := validate.MerkleRootHash(bc.CandidateBlock)
	bc.CandidateBlock.Header.MerkleRootHash = common.BytesToHash(mr)
	log.Info(mr)

	log.Debug("Amount : " + amounts + "\tFrom : " + froms + "\tTo : " + tos)
	log.Debug("Create completed")
	log.Debug(common.PerforatedLine)

	return nil
}

// ShowCandidateBlockTransactionsList shows a list of transactions that exist in a blockchain identified by a ID.
// Therefore, ShowCandidateBlockTransactionsList requires a blockchain ID
// ''ShowCandidateBlockTransactionsList(bcid uint64)''
func ShowCandidateBlockTransactionsList(bcid uint64) error {
	log.Debug("Show Candidate Block Transactions list")
	log.Debug(common.PerforatedLine)

	bcids := strconv.FormatUint(bcid, 10)
	bc, err := core.SelectBlockchain(bcid)

	if err != nil {
		return err
	}

	b := bc.CandidateBlock

	log.Debug(b.TransactionsString("Blockchain ID : " + bcids + "'s Candidate Block"))
	log.Debug(common.PerforatedLine)

	return nil
}
