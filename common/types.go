package common

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//PerforatedLine is perforated line
const PerforatedLine string = "\n-----------------------------------------------------\n"

// Address is type of transaction address(20 bytes)
type Address [20]byte

// Hash is type of block hash value(32 bytes)
type Hash [32]byte

// SHA2Hash is byte to Hash using SHA2 algorithm
func SHA2Hash(b []byte) Hash {
	return sha256.Sum256(b)
}

// MakeTimestamp makes timestamp
func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// SetAddressBytes sets bytes to address(20 bytes)
func (a *Address) SetAddressBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-20:]
	}

	copy(a[20-len(b):], b)
}

// SetHashBytes sets bytes to hash(32 bytes)
func (h *Hash) SetHashBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-32:]
	}

	copy(h[32-len(b):], b)
}

// BytesToAddress converts bytes to address
func BytesToAddress(b []byte) Address {
	var a Address
	a.SetAddressBytes(b)
	return a
}

// BytesToHash converts bytes to hash
func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetHashBytes(b)
	return h
}

// Hex2Bytes converts hex to bytes
func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)

	return h
}

// FromHex converts string to byte
func FromHex(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return Hex2Bytes(s)
}

// StringToUint64 converts string to uint64
func StringToUint64(s string) uint64 {
	u, _ := strconv.ParseUint(s, 10, 64)
	return u
}

// BytesToString converts bytes to string
func BytesToString(b []byte) string {
	return string(b[:])
}

// StringToAddress converts string to address
func StringToAddress(s string) Address { return BytesToAddress([]byte(s)) }

// StringToHash converts string to hash
func StringToHash(s string) Hash { return BytesToHash([]byte(s)) }

// AddressToString converts hash to string
func AddressToString(a Address) string { return base64.StdEncoding.EncodeToString(a[:]) }

// HashToString converts hash to string
func HashToString(h Hash) string { return base64.StdEncoding.EncodeToString(h[:]) }

// HexToHash converts hex to hash
func HexToHash(s string) Hash { return BytesToHash(FromHex(s)) }

// CopyBytes copies bytes
func CopyBytes(b []byte) (copiedBytes []byte) {
	if b == nil {
		return nil
	}
	copiedBytes = make([]byte, len(b))
	copy(copiedBytes, b)

	return
}

// Uint64ArrayToString converts uint64 array to string
func Uint64ArrayToString(a []uint64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
