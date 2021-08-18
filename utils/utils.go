package utils

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rockiecn/p2pdemo/pb"
	"google.golang.org/protobuf/proto"
)

const DEBUG bool = true

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

// calculate purchase hash with marshaled purchase, used as purchase id.
func CalcPurchaseHash(purchaseMarshaled []byte) []byte {
	// unmarshal purchase
	purchase := &pb.Purchase{}
	if err := proto.Unmarshal(purchaseMarshaled, purchase); err != nil {
		log.Fatalln("Failed to parse check:", err)
	}

	// calc hash 32 bytes
	bigNonce := big.NewInt(purchase.NodeNonce)
	hash := crypto.Keccak256([]byte(purchase.StorageAddress), bigNonce.Bytes())
	return hash
}

// generate purchase key from storage address and nonce
func GenPurchaseKey(addr string, nonce *big.Int) ([]byte, error) {
	addrByte := []byte(addr)

	keyByte := MergeSlice(addrByte, nonce.Bytes())
	if DEBUG {
		fmt.Printf("in GenPurchaseKey\n")
		fmt.Printf("storage addr:%s\n", []byte(addr))
		fmt.Printf("nonce:%x\n", nonce.Bytes())
		fmt.Printf("keyByte:%x\n", keyByte)
	}
	return keyByte, nil
}
