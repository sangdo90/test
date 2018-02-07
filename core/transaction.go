package core

import (
	"fmt"
)

// Address represents a address for account.
type Address [20]byte

//A Transaction consists of from(Address) and txdata.
type Transaction struct {
	From Address
	Data txdata
}

// ToBytes returns a slice of bytes of Transaction
func (t *Transaction) ToBytes() []byte {
	res := t.From[:]
	res = append(res, t.Data.To[:]...)
	buf := make([]byte, 8)
	mask := uint64(0xff)
	for i := 0; i < len(buf); i++ {
		buf[i] = byte((t.Data.Amount >> uint(56-(8*i))) & mask)
	}
	res = append(res, buf...)
	return res
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
func NewTransaction(from Address, to Address, amount uint64) *Transaction {
	d := txdata{
		To:     to,
		Amount: amount,
	}
	return &Transaction{From: from, Data: d}
}
