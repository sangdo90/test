package execute

import (
	"errors"
	"strconv"

	"github.com/smartm2m/blockchain/common"
	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

// BlockchainCommands is ...
func BlockchainCommands() {

	_ = command.AddCommand("", command.Command{
		Name:        "bc",
		Description: "manage blockchains",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run: func(args []string) error {
			log.Debug("Blockchain commands")
			return nil
		},
	})

	_ = command.AddCommand("bc", command.Command{
		Name:        "new",
		Description: "create a blockchain",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         NewBlockchain,
	})

	_ = command.AddCommand("bc", command.Command{
		Name:        "list",
		Description: "show blockchains list",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowBlockchainsList,
	})

	_ = command.AddCommand("bc", command.Command{
		Name:        "info",
		Description: "show blockchain Information",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowBlockchainInformation,
	})

	_ = command.AddCommand("bc info", command.Command{
		Name:        "all",
		Description: "show blockchains Information All",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowBlockchainsInformationAll,
	})

	_ = command.AddCommand("bc", command.Command{
		Name:        "block",
		Description: "manage blocks",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run: func(args []string) error {
			log.Debug("Block commands")
			return nil
		},
	})

	_ = command.AddCommand("bc block", command.Command{
		Name:        "info",
		Description: "show information of block",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowBlockInformation,
	})

	_ = command.AddCommand("bc block", command.Command{
		Name:        "list",
		Description: "show list of blocks",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowBlocksList,
	})

	_ = command.AddCommand("bc block", command.Command{
		Name:        "transactions",
		Description: "show list of a block's transactions",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowBlockTransactionsList,
	})

	_ = command.AddCommand("bc", command.Command{
		Name:        "cblock",
		Description: "manage candidate block",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run: func(args []string) error {
			log.Debug("Candidate Block commands")
			return nil
		},
	})

	_ = command.AddCommand("bc cblock", command.Command{
		Name:        "new",
		Description: "create a candidate block",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         NewCandidateBlock,
	})

	_ = command.AddCommand("bc cblock", command.Command{
		Name:        "attach",
		Description: "Attach candidate block to blockchain",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         AttachCandidateBlockToBlockchain,
	})

	_ = command.AddCommand("bc cblock", command.Command{
		Name:        "transaction",
		Description: "manage transaction of candidate block",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run: func(args []string) error {
			log.Debug("Block commands")
			return nil
		},
	})

	_ = command.AddCommand("bc cblock transaction", command.Command{
		Name:        "new",
		Description: "create a transaction of candidate block",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         NewTransactionInCandidateBlock,
	})

	_ = command.AddCommand("bc cblock transaction", command.Command{
		Name:        "list",
		Description: "show transactions list of candidate block",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowCandidateBlockTransactionsList,
	})
}

// NewBlockchain creates a new blockchain containing genesisblock
// ''NewBlockchain()''
func NewBlockchain(args []string) error {
	log.Debug("Create New Blockchain")
	log.Debug(common.PerforatedLine)

	bc := core.AppendBlockchain()
	bc.RegisterBlockchain()
	log.Debug("Create completed")

	return nil
}

// ShowBlockchainsList shows list of blockchains
// ''ShowBlockchainsList()''
func ShowBlockchainsList(args []string) error {
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
func ShowBlockchainInformation(args []string) error {
	log.Debug("Show Blockchain Information")
	log.Debug(common.PerforatedLine)

	if len(args) < 1 {
		return errors.New("Incorrect parameters")
	}

	//log.Debug("Blockchain ID : " + args[0])
	bcid := common.StringToUint64(args[0])
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
func ShowBlockchainsInformationAll(args []string) error {
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
func ShowBlocksList(args []string) error {
	log.Debug("Show Blocks List")
	log.Debug(common.PerforatedLine)

	if len(args) < 1 {
		return errors.New("Incorrect parameters")
	}

	bcid := common.StringToUint64(args[0])
	bc, err := core.SelectBlockchain(bcid)

	if err != nil {
		return err
	}
	log.Debug("Blockchain ID : " + args[0])
	for idx, b := range bc.Blocks {
		i := strconv.Itoa(idx)
		log.Debug(b.String("Block Index : " + i))
		log.Debug(common.PerforatedLine)
	}

	return nil
}

// ShowBlockInformation shows information of block
// ''ShowBlockInformation(bcid uint64, bidx uint64)''
func ShowBlockInformation(args []string) error {
	log.Debug("Show Block Information")
	log.Debug(common.PerforatedLine)

	if len(args) < 2 {
		return errors.New("Incorrect parameters")
	}

	bcid := common.StringToUint64(args[0])
	bidx := common.StringToUint64(args[1])

	bc, err := core.SelectBlockchain(bcid)
	if err != nil {
		return nil
	}

	if bidx > bc.BlockchainHeight-1 {
		return errors.New("Incorrect block index")
	}
	b := bc.Block(bidx)

	if bidx == 0 {
		args[1] = "Candidate"
		b = bc.CandidateBlock
	}

	log.Debug(b.String("Blockchain ID : " + args[0] + "\tBlock Index : " + args[1]))
	log.Debug(common.PerforatedLine)

	return nil
}

// ShowBlockTransactionsList shows list of block's transactions
// ''ShowBlockTransactionsList(bcid uint64, bidx uint64)''
func ShowBlockTransactionsList(args []string) error {
	log.Debug("Show Block Transactions List")
	log.Debug(common.PerforatedLine)

	if len(args) < 2 {
		return errors.New("Incorrect parameters")
	}

	bcid := common.StringToUint64(args[0])
	bidx := common.StringToUint64(args[1])

	bc, err := core.SelectBlockchain(bcid)

	if bidx > bc.BlockchainHeight-1 {
		return errors.New("Incorrect block index")
	}

	b := bc.Block(bidx)

	if bidx == 0 {
		args[1] = "Candidate"
		b = bc.CandidateBlock
	}

	if err != nil {
		return err
	}

	log.Debug(b.TransactionsString("Blockchain ID : " + args[0] + "\tBlock Index : " + args[1]))
	log.Debug(common.PerforatedLine)
	return nil
}

// NewCandidateBlock creates a new candidate block into a blockchain identified by a ID.
// Therefore, NewCandidateBlock requires a blockchain ID
// ''NewCandidateBlock(bcid uint64)''
func NewCandidateBlock(args []string) error {
	log.Debug("Create New Candidate Block")
	log.Debug(common.PerforatedLine)

	if len(args) < 1 {
		return errors.New("Incorrect parameters")
	}

	bcid := common.StringToUint64(args[0])
	bc, err := core.SelectBlockchain(bcid)
	if err != nil {
		return err
	}

	bc.CandidateBlock = core.NewBlock(core.GetLastestBlock(bcid))

	log.Debug(bc.CandidateBlock.String("Blockchain ID : " + args[0] + "'s Candidate Block"))
	log.Debug("Create completed")
	log.Debug(common.PerforatedLine)
	return nil
}

// AttachCandidateBlockToBlockchain attach candidate block into a blockchain identified by a ID.
// Therefore, AttachCandidateBlockToBlockchain requires a blockchain ID
// ''AttachCandidateBlockToBlockchain(bcid uint64)''
func AttachCandidateBlockToBlockchain(args []string) error {
	log.Debug("Attach Candidate Block to Blockchain")
	log.Debug(common.PerforatedLine)

	if len(args) < 1 {
		return errors.New("Incorrect parameters")
	}

	log.Debug("Blockchain ID : " + args[0])
	bcid := common.StringToUint64(args[0])
	bc, err := core.SelectBlockchain(bcid)

	if err != nil {
		return err
	}

	bc.AddBlock()
	log.Debug("Attach completed")
	log.Debug(common.PerforatedLine)

	return nil
}

// NewTransactionInCandidateBlock creates a new transaction and put it in a candidate block of blockchain.
// Also, need to transaction that From A has sent To B.
// Therefore, NewTransactionInCandidateBlock requires a blockchain ID, From, To(address), Amount.
// ''NewTransactionInCandidateBlock(bcid uint64, from uint64, to *common.Address, amount uint64)''
func NewTransactionInCandidateBlock(args []string) error {
	log.Debug("Create New Transaction in the Candidate block")
	log.Debug(common.PerforatedLine)

	if len(args) < 4 {
		return errors.New("Incorrect parameters")
	}

	log.Debug("Blockchain ID : " + args[0])
	bcid := common.StringToUint64(args[0])
	amount := common.StringToUint64(args[1])
	from := common.StringToUint64(args[2])
	to := common.StringToAddress(args[3])

	bc, err := core.SelectBlockchain(bcid)

	if err != nil {
		return err
	}

	t := core.NewTransaction(from, &to, amount)
	bc.CandidateBlock.AddTransaction(t)
	log.Debug("Amount : " + args[1] + "\tFrom : " + args[2] + "\tTo : " + args[3])
	log.Debug("Create completed")
	log.Debug(common.PerforatedLine)

	return nil
}

// ShowCandidateBlockTransactionsList shows a list of transactions that exist in a blockchain identified by a ID.
// Therefore, ShowCandidateBlockTransactionsList requires a blockchain ID
// ''ShowCandidateBlockTransactionsList(bcid uint64)''
func ShowCandidateBlockTransactionsList(args []string) error {
	log.Debug("Show Candidate Block Transactions list")
	log.Debug(common.PerforatedLine)

	if len(args) < 1 {
		return errors.New("Incorrect parameters")
	}

	bcid := common.StringToUint64(args[0])
	bc, err := core.SelectBlockchain(bcid)

	if err != nil {
		return err
	}

	b := bc.CandidateBlock

	log.Debug(b.TransactionsString("Blockchain ID : " + args[0] + "'s Candidate Block"))
	log.Debug(common.PerforatedLine)

	return nil
}
