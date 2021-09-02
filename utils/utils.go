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

// Update Index
func UpdatePayChequeIndex(user bool) {
	var dbfile string
	var Index *[]string
	if user {
		dbfile = "./paycheque.db"
		// if want destination slice to change the source slice, address of source slice needed.
		Index = &global.UserIndex
	} else {
		dbfile = "./storage_paycheque.db"
		Index = &global.StorageIndex
	}
	// clear index
	*Index = (*Index)[0:0]

	// create/open db
	db, err := leveldb.OpenFile(dbfile, nil)
	if err != nil {
		log.Fatal("opfen db error")
	}
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		keyByte := iter.Key()
		*Index = append(*Index, hex.EncodeToString(keyByte))
		fmt.Printf("index len: %d, cap: %d\n", len(*Index), cap(*Index))
		fmt.Printf("index0: %s\n", (*Index)[0])
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// user list data in pay cheque db
func ListPayCheque(user bool) {

	// update userIndex and storageIndex
	UpdatePayChequeIndex(user)

	var dbfile string
	var Index []string
	if user {
		dbfile = "./paycheque.db"
		Index = global.UserIndex
	} else {
		dbfile = "./storage_paycheque.db"
		Index = global.StorageIndex
	}

	// create/open db
	db, err := leveldb.OpenFile(dbfile, nil)
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
	for id < len(Index) {

		// get data
		var in []byte
		keyByte, err := hex.DecodeString(Index[id])
		if err != nil {
			fmt.Println("decodeString error:", err.Error())
			return
		}
		in, err = db.Get(keyByte, nil)
		if err != nil {
			fmt.Println("db get error: ", err)
			return
		}

		PayChequeMarshaled := in[65:]

		// unmarshal it to get Cheque itself
		PayCheque := &pb.PayCheque{}
		if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
			log.Println("Failed to parse paycheck:", err)
			return
		}

		// transmit to string
		strID := strconv.Itoa(id)
		//strValue := strconv.FormatInt(PayCheque.Cheque.Value, 10)
		//strPayValue := strconv.FormatInt(PayCheque.PayValue, 10)

		//strNonce := strconv.FormatInt(PayCheque.Cheque.Nonce, 10)

		//
		value := map[string]string{
			"ID":       strID,
			"FROM":     PayCheque.Cheque.From,
			"TO":       PayCheque.Cheque.To,
			"VALUE":    PayCheque.Cheque.Value,
			"PAYVALUE": PayCheque.PayValue,
			"NONCE":    PayCheque.Cheque.Nonce,
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

// user send a paycheque to remote peer
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

	msg, err2 := db.Get(key, nil)
	if err2 != nil {
		log.Println("db.Get error:", err)
		return err
	}

	// send PayCheque msg to storage
	print.Println100ms("--> Sending PayCheque and sig to storage")
	_, err = s.Write(msg)
	if err != nil {
		log.Println(err)
		return err
	}

	// close stream
	s.Close()

	return nil
}

// user increase the pay value of paycheque in db by 'global.Increase'
func IncPayValueByKey(key []byte) error {
	// create/open db
	db, err := leveldb.OpenFile("./paycheque.db", nil)
	if err != nil {
		log.Println("opfen db error")
		return err
	}
	defer db.Close()

	var msg []byte
	msg, err = db.Get(key, nil)
	if err != nil {
		log.Println("db.Get error:", err)
		return err
	}

	PayChequeMarshaled := msg[65:]

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
	bigValue := big.NewInt(0)
	bigPayvalue := big.NewInt(0)
	bigInc := big.NewInt(0)
	var ok bool
	bigValue, ok = bigValue.SetString(PayCheque.Cheque.Value, 10)
	if !ok {
		print.Println100ms("big.SetString failed")
		return errors.New("big.SetString failed")
	}

	bigPayvalue, ok = bigPayvalue.SetString(PayCheque.PayValue, 10)
	if !ok {
		print.Println100ms("big.SetString failed")
		return errors.New("big.SetString failed")
	}

	bigInc, ok = bigInc.SetString(global.Increase, 10)
	if !ok {
		print.Println100ms("big.SetString failed")
		return errors.New("big.SetString failed")
	}

	bigAdd := bigPayvalue.Add(bigPayvalue, bigInc)

	// if PayCheque.Cheque.Value < PayCheque.PayValue+global.Increase {
	// 	fmt.Println("Cheque not enough cash.")
	// 	return errors.New("cheque not enough cash")
	// }
	if bigValue.Cmp(bigAdd) == -1 {
		fmt.Println("Cheque not enough cash.")
		return errors.New("cheque not enough cash")
	}

	// increase pay value
	//PayCheque.PayValue = PayCheque.PayValue + global.Increase
	PayCheque.PayValue = bigAdd.String()

	if global.DEBUG {
		print.Printf100ms("-> Pay cheque after increased:%d\n", PayCheque.PayValue)
	}

	// serialize
	PayChequeMarshaled, err = proto.Marshal(PayCheque)
	if err != nil {
		fmt.Println("marshal paycheque error:", err)
		return err
	}

	// recompute paycheque sig
	hash := CalcPayChequeHash(PayCheque)
	if global.DEBUG {
		print.Printf100ms("DEBUG> paycheque hash: %x\n", hash)
	}
	// sign PayCheque by user
	var userSkByte = []byte(global.StrUserSK)
	PayChequeSig, err := sigapi.Sign(hash, userSkByte)
	if err != nil {
		log.Print("sign error")
		return err
	}

	// rewrite updated paycheque
	msg = MergeSlice(PayChequeSig, PayChequeMarshaled)

	// put into db
	err = db.Put(key, msg, nil)
	if err != nil {
		fmt.Println("put paycheque error:", err)
		return err
	}

	return nil
}

// transmit ID of a pay cheque to key, use Index[]
// user: true for user, false for storage
func IDtoKey(uID uint, user bool) ([]byte, error) {

	var Index []string
	if user {
		Index = global.UserIndex
	} else {
		Index = global.StorageIndex
	}

	keyByte, err := hex.DecodeString(Index[uID])
	if err != nil {
		fmt.Println("decode string error: ", err)
		return nil, err
	}

	return keyByte, nil
}

// show pay cheque info by key
func ShowPayChequeInfoByKey(key []byte) error {
	// create/open db
	db, err := leveldb.OpenFile("./paycheque.db", nil)
	if err != nil {
		log.Println("opfen db error")
		return err
	}
	defer db.Close()

	msg, err2 := db.Get(key, nil)
	if err2 != nil {
		log.Println("db.Get error:", err)
		return err
	}

	PayChequeSig := msg[:65]
	PayChequeMarshaled := msg[65:]
	// unmarshal it to get Cheque itself
	PayCheque := &pb.PayCheque{}
	if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
		log.Println("Failed to parse pay check:", err)
		return err
	}

	print.PrintPayCheque(PayCheque)
	print.Printf100ms("PayChequeSig: %x\n", PayChequeSig)

	return nil
}
