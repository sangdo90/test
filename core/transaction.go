package core

import "fmt"

// Address represents a address for account.
type Address [20]byte

//A Transaction consists of from(Address) and txdata.
type Transaction struct {
	From Address
	Data txdata
}

// ToString returns a hex string for account address
func (a Address) ToString() string {
	res := "0x"
	for _, v := range a {
		res += fmt.Sprintf("%02x", v)
	}

	return res
}

type txdata struct {
	To     Address
	Amount uint64
}

//NewTransaction creates a new transaction.
func NewTransaction(amount uint64, from Address, to Address) *Transaction {
	d := txdata{
		To:     to,
		Amount: amount,
	}
	return &Transaction{From: from, Data: d}
}
