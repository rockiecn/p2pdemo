package utils

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
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
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/p2pdemo/print"
	"github.com/rockiecn/p2pdemo/sigapi"
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

// send a paycheque to remote peer
func SendChequeByKey(key []byte) error {
	// create/open db
	db, err := leveldb.OpenFile("./paycheque.db", nil)
	if err != nil {
		log.Println("opfen db error")
		return err
	}
	defer db.Close()

	print.Printf100ms("Opening stream to peerID: %v\n", global.Peerid)
	s, err := hostops.HostInfo.NewStream(context.Background(), global.Peerid, "/2")
	if err != nil {
		log.Println("open stream error: ", err)
		return err
	}

	var PayChequeMarshaled []byte
	PayChequeMarshaled, err = db.Get(key, nil)
	if err != nil {
		log.Println("db.Get error:", err)
		return err
	}

	// unmarshal it to get Cheque itself
	PayCheque := &pb.PayCheque{}
	if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
		log.Fatalln("Failed to parse pay check:", err)
		return err
	}

	if global.DEBUG {
		print.Println100ms("-> Pay cheque info:")
		print.PrintPayCheque(PayCheque)
	}

	// PayCheque should be created, signed and sent by user

	// calc hash from PayCheque
	hash := CalcHash(PayCheque.Cheque.From, PayCheque.Cheque.NodeNonce, PayCheque.To, PayCheque.PayValue)
	if global.DEBUG {
		print.Printf100ms("DEBUG> hash: %x\n", hash)
	}
	// sign PayCheque by user' sk
	// user address: 1ab6a9f2b90004c1269563b5da391250ede3c114
	//var userSkByte = []byte("b91c265cabae210642d66f9d59137eac2fab2674f4c1c88df3b8e9e6c1f74f9f")
	var userSkByte = global.UserSK
	PayChequeSig, err := sigapi.Sign(hash, userSkByte)
	if err != nil {
		log.Print("sign error")
		return err
	}

	if global.DEBUG {
		// for debug
		print.Printf100ms("DEBUG> From: %s\n", PayCheque.Cheque.From)
		print.Printf100ms("DEBUG> NodeNonce: %d\n", PayCheque.Cheque.NodeNonce)
		print.Printf100ms("DEBUG> To: %s\n", PayCheque.To)
		print.Printf100ms("DEBUG> PayValue: %d\n", PayCheque.PayValue)
		print.Printf100ms("DEBUG> signature: %x\n", PayChequeSig)
	}

	// send PayCheque msg to storage
	print.Println100ms("--> Sending PayCheque to storage")
	_, err = s.Write(PayChequeMarshaled)
	if err != nil {
		log.Println(err)
		return err
	}

	// close stream
	s.Close()

	return nil
}

// increase the pay value of paycheque in db by 'global.Increase'
func IncPayValueByKey(key []byte) error {
	// create/open db
	db, err := leveldb.OpenFile("./paycheque.db", nil)
	if err != nil {
		log.Println("opfen db error")
		return err
	}
	defer db.Close()

	var PayChequeMarshaled []byte
	PayChequeMarshaled, err = db.Get(key, nil)
	if err != nil {
		log.Println("db.Get error:", err)
		return err
	}

	// unmarshal it to get Cheque itself
	PayCheque := &pb.PayCheque{}
	if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
		log.Println("Failed to parse pay check:", err)
		return err
	}

	if global.DEBUG {
		print.Printf100ms("-> Pay cheque before increased:%d\n", PayCheque.PayValue)
	}

	// not enough cash
	if PayCheque.Cheque.Value < PayCheque.PayValue+global.Increase {
		fmt.Println("Cheque not enough cash.")
		return errors.New("cheque not enough cash")
	}

	// increase pay value
	PayCheque.PayValue = PayCheque.PayValue + global.Increase
	if global.DEBUG {
		print.Printf100ms("-> Pay cheque after increased:%d\n", PayCheque.PayValue)
	}

	// serialize
	PayChequeMarshaled, err = proto.Marshal(PayCheque)
	if err != nil {
		fmt.Println("marshal paycheque error:", err)
		return err
	}

	// put into db
	err = db.Put(key, PayChequeMarshaled, nil)
	if err != nil {
		fmt.Println("put paycheque error:", err)
		return err
	}

	return nil
}

// transmit ID of a pay cheque to key, use Index[]
func IDtoKey(uID uint) ([]byte, error) {

	keyByte, err := hex.DecodeString(global.Index[uID])
	if err != nil {
		fmt.Println("decode string error: ", err)
		return nil, err
	}

	return keyByte, nil

}
