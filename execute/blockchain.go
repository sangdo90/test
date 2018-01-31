package execute

import (
	"bytes"
	"errors"

	"github.com/smartm2m/blockchain/common"
	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

// BlockchainCommands is ...
func BlockchainCommands() {

	_ = command.AddCommand("", command.Command{
		Name:        "transaction",
		Description: "manage transactions",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run: func(args []string) error {
			log.Debug("Transaction commands")
			return nil
		},
	})

	_ = command.AddCommand("transaction", command.Command{
		Name:        "new",
		Description: "create a candidate transaction",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         AddTransactionToCandidateBlock,
	})

	_ = command.AddCommand("transaction", command.Command{
		Name:        "list",
		Description: "show candidate transactions list",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowCandidateBlockTransactionsList,
	})

	_ = command.AddCommand("", command.Command{
		Name:        "block",
		Description: "manage blocks",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run: func(args []string) error {
			log.Debug("Block commands")
			return nil
		},
	})

	_ = command.AddCommand("block", command.Command{
		Name:        "new",
		Description: "create a block",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         NewBlock,
	})

	_ = command.AddCommand("block", command.Command{
		Name:        "attach",
		Description: "Attach block to blockchain",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         AttachBlockToBlockchain,
	})

	_ = command.AddCommand("block", command.Command{
		Name:        "info",
		Description: "show block information",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowBlockInformation,
	})

	_ = command.AddCommand("block", command.Command{
		Name:        "transactions",
		Description: "show transactions of block",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowBlockTransactionsList,
	})

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
		Description: "create a blockchain",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         NewBlockchain,
	})

	_ = command.AddCommand("blockchain", command.Command{
		Name:        "blocks",
		Description: "show list of blocks",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowBlocksList,
	})

	_ = command.AddCommand("blockchain", command.Command{
		Name:        "list",
		Description: "show blockchains list",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowBlockchainsList,
	})

	_ = command.AddCommand("blockchain", command.Command{
		Name:        "info",
		Description: "show blockchain Information",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowBlockchainInformation,
	})

	_ = command.AddCommand("blockchain", command.Command{
		Name:        "info all",
		Description: "show blockchains Information All",
		Commands:    make([]command.Command, 0),
		Flags:       nil,
		Run:         ShowBlockchainsInformationAll,
	})
}

// AddTransactionToCandidateBlock creates a new transaction and put it in a candidate block of blockchain.
// Also, need to transaction that From A has sent To B.
// Therefore, AddTransactionToCandidateBlock requires a blockchain ID, From, To(address), Amount.
// ''AddTransactionToCandidateBlcok(bcid uint64, from uint64, to *common.Address, amount uint64)''
func AddTransactionToCandidateBlock(args []string) error {
	log.Debug("Create New Transaction in the Candidate block")
	if len(args) < 4 {
		return errors.New("Incorrect parameters")
	}

	log.Debug("Blockchain ID : " + args[0])
	bcid := common.StringToUint64(args[0])
	from := common.StringToUint64(args[1])
	to := common.StringToAddress(args[2])
	amount := common.StringToUint64(args[3])
	bc, err := core.SelectBlockchain(bcid)

	if err != nil {
		return err
	}

	t := core.NewTransaction(from, &to, amount)
	bc.CandidateBlock.AddTransaction(t)
	return nil
}

// NewBlock creates a new candidate block into a blockchain identified by a ID.
// Therefore, NewBlock requires a blockchain ID
// ''NewBlock(bcid uint64)''
func NewBlock(args []string) error {
	log.Debug("Create New Candidate Block")
	if len(args) < 1 {
		return errors.New("Incorrect parameters")
	}

	log.Debug("Blockchain ID : " + args[0])
	bcid := common.StringToUint64(args[0])
	bc, err := core.SelectBlockchain(bcid)
	if err != nil {
		return err
	}

	bc.CandidateBlock = core.NewBlock(core.GetLastestBlock(bcid))
	return nil
}

// NewBlockchain creates a new blockchain containing genesisblock
// ''NewBlockchain()''
func NewBlockchain(args []string) error {
	log.Debug("Create New Blockchain")
	bc := core.NewBlockchain()
	bc.RegisterBlockchain()
	return nil
}

// AttachBlockToBlockchain attach candidate block into a blockchain identified by a ID.
// Therefore, AttachBlockToBlockchain requires a blockchain ID
// ''AttachBlockToBlockchain(bcid uint64)''
func AttachBlockToBlockchain(args []string) error {
	log.Debug("Attach Block to Blockchain")
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
	return nil
}

// ShowCandidateBlockTransactionsList shows a list of transactions that exist in a blockchain identified by a ID.
// Therefore, ShowCandidateBlockTransactionsList requires a blockchain ID
// ''ShowCandidateBlockTransactionsList(bcid uint64)''
func ShowCandidateBlockTransactionsList(args []string) error {
	log.Debug("Show Candidate Block Transactions list")
	if len(args) < 1 {
		return errors.New("Incorrect parameters")
	}

	log.Debug("Blockchain ID : " + args[0])
	bcid := common.StringToUint64(args[0])
	bc, err := core.SelectBlockchain(bcid)

	if err != nil {
		return err
	}
	b := bc.CandidateBlock
	ts := b.BlockTransactions()
	var buffer bytes.Buffer
	buffer.WriteString(b.String())
	buffer.WriteString("-----------------------------------------------------\n")
	buffer.WriteString("0\tFrom\tTo\tAmount\n")
	for idx, t := range ts {
		buffer.WriteString(string(idx) + t.String())
	}
	buffer.WriteString("-----------------------------------------------------\n")

	log.Debug(buffer.String())
	return nil
}

// ShowBlockInformation shows information of block
// ''ShowBlockInformation(bcid uint64, bidx uint64)''
func ShowBlockInformation(args []string) error {
	log.Debug("Show Block Information")
	if len(args) < 2 {
		return errors.New("Incorrect parameters")
	}

	log.Debug("Blockchain ID : " + args[0] + "\n")
	log.Debug("Block index : " + args[1] + "\n")
	bcid := common.StringToUint64(args[0])
	bidx := common.StringToUint64(args[1])
	bc, err := core.SelectBlockchain(bcid)
	if err != nil {
		return nil
	}

	log.Debug(bc.Block(bidx).String())
	return nil
}

// ShowBlockTransactionsList shows list of block's transactions
// ''ShowBlockTransactionsList(bcid uint64, bidx uint64)''
func ShowBlockTransactionsList(args []string) error {
	log.Debug("Show Block Transactions List")
	if len(args) < 2 {
		return errors.New("Incorrect parameters")
	}

	bcid := common.StringToUint64(args[0])
	bidx := common.StringToUint64(args[1])

	bc, err := core.SelectBlockchain(bcid)
	b := bc.Block(bidx)
	ts := b.BlockTransactions()
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	buffer.WriteString(b.String())
	buffer.WriteString("-----------------------------------------------------\n")
	buffer.WriteString("0\tFrom\tTo\tAmount\n")
	for idx, t := range ts {
		buffer.WriteString(string(idx) + t.String())
	}
	buffer.WriteString("-----------------------------------------------------\n")

	log.Debug(buffer.String())
	return nil
}

// ShowBlocksList shows list of block
// ''ShowBlocksList(bcid uint64)''
func ShowBlocksList(args []string) error {
	log.Debug("Show Blocks List")
	if len(args) < 1 {
		return errors.New("Incorrect parameters")
	}

	bcid := common.StringToUint64(args[0])
	bc, err := core.SelectBlockchain(bcid)

	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	buffer.WriteString(bc.String())
	for idx, b := range bc.Blocks {
		buffer.WriteString("-----------------------------------------------------\n")
		buffer.WriteString(string(idx) + "-th Block: \n")
		buffer.WriteString(b.String() + "\n")
		buffer.WriteString("-----------------------------------------------------\n")
	}

	log.Debug(buffer.String())
	return nil
}

// ShowBlockchainsList shows list of blockchains
// ''ShowBlockchainsList()''
func ShowBlockchainsList(args []string) error {
	log.Debug("Show Blockchains List")
	var idx uint64
	var blockchainsList []uint64
	for idx = 0; idx < core.GlobalBlockchainsLength; idx++ {
		blockchainsList = append(blockchainsList, core.GlobalBlockchains[idx].ID)
	}

	log.Debug(common.Uint64ArrayToString(blockchainsList, ", "))
	return nil
}

// ShowBlockchainInformation shows information of blockchain identified by a ID.
// Therefore, ShowBlockchainInformation requires a blockchain ID
// ''ShowBlockchainInformation(bcid uint64)''
func ShowBlockchainInformation(args []string) error {
	log.Debug("Show Blockchain Information")
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
	return nil
}

// ShowBlockchainsInformationAll shows inforamtion of all blockchains
// ''ShowBlockchainsInformationAll()''
func ShowBlockchainsInformationAll(args []string) error {
	log.Debug("Show Blockchains Information All")
	var idx uint64
	var buffer bytes.Buffer
	for idx = 0; idx < core.GlobalBlockchainsLength; idx++ {
		bc := core.GlobalBlockchains[idx]
		buffer.WriteString(bc.String())
	}
	log.Debug(buffer)
	return nil
}
