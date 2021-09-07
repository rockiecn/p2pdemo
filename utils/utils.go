package utils

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/rockiecn/p2pdemo/global"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/p2pdemo/print"
)

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

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.LittleEndian.Uint64(buf))
}

func Byte2Str(data []byte) string {
	//var str string = string(data[:len(data)])
	var str string = string(data[:])
	return str
}

// CalcHash - calculate hash from PayCheque
func CalcHash(From string, int64Nonce int64, stAddress string, int64PayAmount int64) []byte {

	//
	userBytes, _ := hex.DecodeString(From)
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

// calculate Cheque hash
func CalcChequeHash(Cheque *pb.Cheque) []byte {
	// unmarshal Cheque
	// Cheque := &pb.Cheque{}
	// if err := proto.Unmarshal(ChequeMarshaled, Cheque); err != nil {
	// 	log.Fatalln("Failed to parse check:", err)
	// }

	fmt.Println("cheque.value:", Cheque.Value)
	// calc hash 32 bytes
	//bigValue := big.NewInt(Cheque.Value)
	bigValue := big.NewInt(0)
	var ok bool
	bigValue, ok = bigValue.SetString(Cheque.Value, 10)
	if !ok {
		print.Println100ms("big.SetString failed")
		return nil
	}

	if global.DEBUG {
		fmt.Println("in calcchequehash")
		fmt.Println("bigvalue:", bigValue.String())
	}

	valuePad32 := common.LeftPadBytes(bigValue.Bytes(), 32)

	//bigNonce := big.NewInt(Cheque.Nonce)
	bigNonce := big.NewInt(0)
	bigNonce.SetString(Cheque.Nonce, 10)

	noncePad32 := common.LeftPadBytes(bigNonce.Bytes(), 32)

	tokenBytes, _ := hex.DecodeString(Cheque.TokenAddress)
	fromBytes, _ := hex.DecodeString(Cheque.From)
	toBytes, _ := hex.DecodeString(Cheque.To)
	operatorBytes, _ := hex.DecodeString(Cheque.OperatorAddress)
	contractBytes, _ := hex.DecodeString(Cheque.ContractAddress)

	hash := crypto.Keccak256(
		valuePad32,
		tokenBytes,
		noncePad32,
		fromBytes,
		toBytes,
		operatorBytes,
		contractBytes,
	)
	return hash
}

// calculate Cheque hash
func CalcPayChequeHash(PayCheque *pb.PayCheque) []byte {

	// calc hash 32 bytes
	//bigValue := big.NewInt(PayCheque.Cheque.Value)
	bigValue := big.NewInt(0)
	var ok bool
	bigValue, ok = bigValue.SetString(PayCheque.Cheque.Value, 10)
	if !ok {
		print.Println100ms("big.SetString failed")
		return nil
	}

	valuePad32 := common.LeftPadBytes(bigValue.Bytes(), 32)

	//bigNonce := big.NewInt(PayCheque.Cheque.Nonce)
	bigNonce := big.NewInt(0)
	bigNonce.SetString(PayCheque.Cheque.Nonce, 10)

	noncePad32 := common.LeftPadBytes(bigNonce.Bytes(), 32)

	//bigPayValue := big.NewInt(PayCheque.PayValue)
	bigPayValue := big.NewInt(0)
	bigPayValue, ok = bigPayValue.SetString(PayCheque.PayValue, 10)
	if !ok {
		print.Println100ms("big.SetString failed")
		return nil
	}

	payvaluePad32 := common.LeftPadBytes(bigPayValue.Bytes(), 32)

	tokenBytes, _ := hex.DecodeString(PayCheque.Cheque.TokenAddress)
	fromBytes, _ := hex.DecodeString(PayCheque.Cheque.From)
	toBytes, _ := hex.DecodeString(PayCheque.Cheque.To)
	operatorBytes, _ := hex.DecodeString(PayCheque.Cheque.OperatorAddress)
	contractBytes, _ := hex.DecodeString(PayCheque.Cheque.ContractAddress)

	hash := crypto.Keccak256(
		valuePad32,
		tokenBytes,
		noncePad32,
		fromBytes,
		toBytes,
		operatorBytes,
		contractBytes,
		payvaluePad32,
	)
	return hash
}

// generate Cheque key from operator address, storage address and nonce
//func GenChequeKey(addr string, nonce *big.Int) ([]byte, error) {
func GenChequeKey(Cheque *pb.Cheque) ([]byte, error) {

	//bigNonce := big.NewInt(Cheque.Nonce)
	bigNonce := big.NewInt(0)
	bigNonce.SetString(Cheque.Nonce, 10)

	opByte := []byte(Cheque.OperatorAddress)
	toByte := []byte(Cheque.To)

	keyByte := MergeSlice(opByte, toByte)
	keyByte = MergeSlice(keyByte, bigNonce.Bytes())
	if global.DEBUG {
		fmt.Printf("<DEBUG> in GenChequeKey\n")
		fmt.Printf("<DEBUG> storage addr:%s\n", []byte(Cheque.OperatorAddress))
		fmt.Printf("<DEBUG> nonce:%x\n", bigNonce.Bytes())
		fmt.Printf("<DEBUG> keyByte:%x\n", keyByte)
	}
	return keyByte, nil
}
