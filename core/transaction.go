package core

import (
	"github.com/smartm2m/blockchain/common"
)

//A txdata consists of Recipient, Amount and Payload.
type txdata struct {
	Recipient *common.Address
	Amount    uint64
	Payload   []byte
}

//A Transaction consists of ID and txdata.
type Transaction struct {
	ID   uint64
	Data txdata
}

//NewTransaction creates a new transaction.
func NewTransaction(from uint64, to *common.Address, amount uint64, data []byte) *Transaction {
	if len(data) > 0 {
		data = common.CopyBytes(data)
	}
	d := txdata{
		Recipient: to,
		Payload:   data,
		Amount:    amount,
	}
	return &Transaction{ID: from, Data: d}
}
