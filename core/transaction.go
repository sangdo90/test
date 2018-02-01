package core

import (
	"github.com/smartm2m/blockchain/common"
)

//A txdata consists of Recipient, Amount, Payload and Description.
type txdata struct {
	To     *common.Address
	Amount uint64
}

//A Transaction consists of ID and txdata.
type Transaction struct {
	From uint64
	Data txdata
}

//NewTransaction creates a new transaction.
func NewTransaction(from uint64, to *common.Address, amount uint64) *Transaction {
	d := txdata{
		To:     to,
		Amount: amount,
	}
	return &Transaction{From: from, Data: d}
}
