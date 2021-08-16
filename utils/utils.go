package utils

import (
	"encoding/binary"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

/*
// Str2Byte - convert string to []byte
func Str2Byte(str string) []byte {
	var ret []byte = []byte(str)
	return ret
}

// Byte2Str - convert []byte to strnig
func Byte2Str(data []byte) string {
	//var str string = string(data[:len(data)])
	var str string = string(data[:])
	return str
}
*/

// MergeSlice - merge some slice together
func MergeSlice(s1 []byte, s2 []byte) []byte {
	slice := make([]byte, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}

// Uint32ToBytes - convert uint32 to bytes
func Uint32ToBytes(n uint32) []byte {
	a := make([]byte, 4)
	binary.LittleEndian.PutUint32(a, n)
	return a
}

// BytesToUint32 - convert bytes to uint32
func BytesToUint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

// CalcHash - calculate hash from cheque
func CalcHash(userAddress string, int64Nonce int64, stAddress string, int64PayAmount int64) []byte {

	//
	userBytes, _ := hex.DecodeString(userAddress)
	// pad nonce to 32 bytes
	bigNonce := big.NewInt(int64Nonce)
	noncePad32 := common.LeftPadBytes(bigNonce.Bytes(), 32)

	// 20 bytes
	stBytes, _ := hex.DecodeString(stAddress)
	// pad pay amount to 32 bytes
	bigPay := big.NewInt(int64PayAmount)
	payPad32 := common.LeftPadBytes(bigPay.Bytes(), 32)

	// calc hash 32 bytes
	hash := crypto.Keccak256(userBytes, noncePad32, stBytes, payPad32)
	return hash
}
