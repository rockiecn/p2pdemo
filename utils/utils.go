package utils

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/liushuochen/gotable"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/protobuf/proto"

	"github.com/rockiecn/p2pdemo/global"
	"github.com/rockiecn/p2pdemo/pb"
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

// calculate Cheque hash with marshaled Cheque, used as Cheque id.
func CalcChequeHash(ChequeMarshaled []byte) []byte {
	// unmarshal Cheque
	Cheque := &pb.Cheque{}
	if err := proto.Unmarshal(ChequeMarshaled, Cheque); err != nil {
		log.Fatalln("Failed to parse check:", err)
	}

	// calc hash 32 bytes
	bigNonce := big.NewInt(Cheque.NodeNonce)
	hash := crypto.Keccak256([]byte(Cheque.To), bigNonce.Bytes())
	return hash
}

// generate Cheque key from storage address and nonce
func GenChequeKey(addr string, nonce *big.Int) ([]byte, error) {
	addrByte := []byte(addr)

	keyByte := MergeSlice(addrByte, nonce.Bytes())
	if global.DEBUG {
		fmt.Printf("in GenChequeKey\n")
		fmt.Printf("storage addr:%s\n", []byte(addr))
		fmt.Printf("nonce:%x\n", nonce.Bytes())
		fmt.Printf("keyByte:%x\n", keyByte)
	}
	return keyByte, nil
}

// Update Index
func UpdatePayChequeIndex() {
	// clear index
	global.Index = global.Index[0:0]

	// create/open db
	db, err := leveldb.OpenFile("./paycheque.db", nil)
	if err != nil {
		log.Fatal("opfen db error")
	}
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		keyByte := iter.Key()
		global.Index = append(global.Index, hex.EncodeToString(keyByte))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// user list data in pay cheque db
func ListPayCheque() {
	UpdatePayChequeIndex()
	// create/open db
	db, err := leveldb.OpenFile("./paycheque.db", nil)
	if err != nil {
		log.Fatal("opfen db error")
	}
	defer db.Close()

	// show table
	table, err := gotable.Create("ID", "FROM", "TO", "VALUE", "PAYVALUE", "NONCE")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// show table
	var id int = 0
	for id < len(global.Index) {

		// get data
		var PayChequeMarshaled []byte
		keyByte, err := hex.DecodeString(global.Index[id])
		if err != nil {
			fmt.Println("decodeString error:", err.Error())
			return
		}
		PayChequeMarshaled, _ = db.Get(keyByte, nil)
		// unmarshal it to get Cheque itself
		PayCheque := &pb.PayCheque{}
		if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
			log.Fatalln("Failed to parse check:", err)
		}

		// transmit to string
		strID := strconv.Itoa(id)
		strValue := strconv.FormatInt(PayCheque.Cheque.Value, 10)
		strPayValue := strconv.FormatInt(PayCheque.PayValue, 10)
		strNonce := strconv.FormatInt(PayCheque.Cheque.NodeNonce, 10)

		//
		value := map[string]string{
			"ID":       strID,
			"FROM":     PayCheque.From,
			"TO":       PayCheque.To,
			"VALUE":    strValue,
			"PAYVALUE": strPayValue,
			"NONCE":    strNonce,
		}
		err = table.AddRow(value)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		id++
	}

	//r, _ := table.Json(4)
	//fmt.Println(r)
	//table.CloseBorder()
	table.PrintTable()
}
