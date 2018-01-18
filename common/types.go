package common

import (
	"time"
	"encoding/hex"
	"encoding/base64"
)

type Hash [32]byte

func MakeTimestamp() int64{
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-32:]
	}

	copy(h[32-len(b):], b)
}

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)

	return h
}

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

func StringToHash(s string) Hash { return BytesToHash([]byte(s)) }
func HexToHash(s string) Hash    { return BytesToHash(FromHex(s)) }
func HashToString(h Hash) string { return base64.StdEncoding.EncodeToString(h[:])}