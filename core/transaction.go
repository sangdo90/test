package core

import (
	"github.com/smartm2m/blockchain/common"
)

type txdata struct {
	Recipient *common.Address
	Amount    uint64
	Payload   []byte
}

type Transaction struct {
	ID   uint64
	Data txdata
}

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
