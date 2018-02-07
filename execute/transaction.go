package execute

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/smartm2m/blockchain/core"
	"github.com/smartm2m/chainutil/console/command"
	"github.com/smartm2m/chainutil/log"
)

// TransactionCommands registers console commands for transactions.
func TransactionCommands() {
	_ = command.AddCommand("", command.Command{
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
	})
}

// NewTransactionInCandidateBlock creates a new transaction and put it in a candidate block of blockchain.
// Also, need to transaction that From A has sent To B.
// Therefore, NewTransactionInCandidateBlock requires a blockchain index, From, To(address), Amount.
// ''NewTransactionInCandidateBlock(bcid uint64, amount uint64, from uint64, to *common.Address)''
func NewTransactionInCandidateBlock(bcid uint64, amount uint64, from []byte, to []byte) error {
	log.Info("Create New Transaction in the Candidate block")
	log.Info(perforatedLine)
	bcids := strconv.FormatUint(bcid, 10)
	amounts := strconv.FormatUint(amount, 10)

	tl, fl := len(to), len(from)
	if len(to) > 20 {
		tl = 20
	}
	if len(from) > 20 {
		fl = 20
	}
	var toa, froma core.Address

	copy(toa[:], to[:tl])
	copy(froma[:], from[:fl])

	log.Debug("Blockchain index : " + bcids)

	bc, err := getBlockchain(bcid)
	if err != nil {
		return err
	}

	t := core.NewTransaction(froma, toa, amount)
	_ = bc.CandidateBlock.AddTransaction(t)
	log.Info("Amount : " + amounts + "\tFrom : " + froma.ToString() + "\tTo : " + toa.ToString())
	log.Debug("Create completed")
	log.Info(perforatedLine)

	return nil
}

// ShowCandidateBlockTransactionsList shows a list of transactions that exist in a blockchain identified by a index.
// Therefore, ShowCandidateBlockTransactionsList requires a blockchain index
// ''ShowCandidateBlockTransactionsList(bcid uint64)''
func ShowCandidateBlockTransactionsList(bcid uint64) error {
	log.Info("Show Candidate Block Transactions list")
	log.Info(perforatedLine)

	bcids := strconv.FormatUint(bcid, 10)
	bc, err := getBlockchain(bcid)

	if err != nil {
		return err
	}

	b := bc.CandidateBlock

	log.Info(transactionsString(b, "Blockchain index : "+bcids+"'s Candidate Block"))
	log.Info(perforatedLine)

	return nil
}

func transactionsString(b *core.Block, name string) string {
	res := name + "\nNum\tAmount\tFrom\t\t\t\t\t\tTo\n"
	buffer := bytes.NewBuffer([]byte{})

	for idx, t := range b.Body.Transactions {
		fmt.Fprintf(buffer, strconv.Itoa(idx+1))
		fmt.Fprintf(buffer, "\t%v", t.Data.Amount)
		fmt.Fprintf(buffer, "\t%v", t.From.ToString())
		fmt.Fprintf(buffer, "\t%v\n", t.Data.To.ToString())
	}

	res = res + buffer.String()
	return res
}
