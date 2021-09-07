package user

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/liushuochen/gotable"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/protobuf/proto"

	"github.com/rockiecn/p2pdemo/db"
	"github.com/rockiecn/p2pdemo/global"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/p2pdemo/print"
	"github.com/rockiecn/p2pdemo/sigapi"
	"github.com/rockiecn/p2pdemo/utils"
)

type User struct {
	UserDB          *leveldb.DB // provider -> nonce
	ContractAddress string      // contract address

	DBfile  string
	DBIndex []string

	UserAddr string // "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"
	UserSK   string // "7e5bfb82febc4c2c8529167104271ceec190eafdca277314912eaabdb67c6e5f"

	db.DB // interface DB
}

// init user, need db open first
func (user *User) Init() error {
	user.DBfile = "./user.db"

	user.OpenDB()
	defer user.CloseDB()

	user.ContractAddress = ""

	user.DBIndex = []string{}

	user.UserAddr = "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"
	user.UserSK = "7e5bfb82febc4c2c8529167104271ceec190eafdca277314912eaabdb67c6e5f"

	byteAddr, err := user.UserDB.Get([]byte("contractAddr"), nil)
	if err != nil {
		return errors.New("operator init: read contract address failed")
	}
	user.ContractAddress = string(byteAddr)

	return nil
}

// user get Cheque from operator
func (user *User) BuyCheque() (*pb.PayCheque, error) {
	user.OpenDB()
	defer user.CloseDB()

	if !global.RemoteExist {
		fmt.Println("No remote peer exist, record one as operator.")
		return nil, errors.New("no remote peer")
	}
	// connect to operator , get stream
	s, err := hostops.HostInfo.NewStream(context.Background(), global.Peerid, "/1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Read from stream
	fmt.Println("--> user receive Cheque from operator")
	in, err := ioutil.ReadAll(s)
	if err != nil {
		fmt.Println("Error reading :", err)
		return nil, err
	}

	// parse data
	sigByte := in[:65] //65
	ChequeMarshaled := in[65:]

	if global.DEBUG {
		print.Printf100ms("sigByte:%x\n", sigByte)
		print.Printf100ms("ChequeMarshaled:%x\n", ChequeMarshaled)
	}

	// unmarshal
	Cheque := &pb.Cheque{}
	if err := proto.Unmarshal(ChequeMarshaled, Cheque); err != nil {
		log.Fatalln("Failed to parse check:", err)
	}

	if global.DEBUG {
		print.Printf100ms("--> Received Cheque:\n")
		print.PrintCheque(Cheque)
	}

	// verify signature of Cheque, signed by operator

	// ======  string to common.Address , for verify ======
	// string to byte
	opAddrByte, err := hex.DecodeString(Cheque.OperatorAddress)
	if err != nil {
		panic("decode error")
	}
	// []byte to common.Address, add 0x prefix automaticlly.
	opAddress := common.BytesToAddress(opAddrByte)

	// calc hash for verify cheque sig
	hash := utils.CalcChequeHash(Cheque)
	if global.DEBUG {
		print.Printf100ms("Cheque receive, hash: %x\n", hash)
	}

	// verify Cheque signature
	ok, _ := sigapi.Verify(hash, sigByte, opAddress)
	if !ok {
		fmt.Println("<signature of Cheque verify failed>")
		return nil, errors.New("cheque signature verify failed")
	}

	fmt.Println("<signature of Cheque verify success>")

	// Cheque key: To + nonce
	ChequeKey, err := utils.GenChequeKey(Cheque)
	if err != nil {
		log.Fatal("GenChequeKey error")
		return nil, errors.New("generate cheque key failed")
	}

	if global.DEBUG {
		print.Printf100ms("ChequeKey: %x\n", ChequeKey)
	}

	//print.Printf100ms("---------------- cash string:%s\n", string(cashAddrByte))
	// construct pay cheque with initial payvalue=0
	PayCheque := &pb.PayCheque{}
	PayCheque.Cheque = Cheque
	PayCheque.ChequeSig = sigByte
	PayCheque.PayValue = (big.NewInt(0)).String()

	// show generated paycheque
	if global.DEBUG {
		print.Println100ms("---------- show generated paycheque -----------")
		print.Printf100ms("Value: %s\n", Cheque.Value)
		print.Printf100ms("TokenAddress: %s\n", Cheque.TokenAddress)
		print.Printf100ms("Nonce: %s\n", Cheque.Nonce)
		print.Printf100ms("From: %s\n", Cheque.From)
		print.Printf100ms("To: %s\n", Cheque.To)
		print.Printf100ms("OperatorAddress: %s\n", Cheque.OperatorAddress)
		print.Printf100ms("ContractAddress: %s\n", Cheque.ContractAddress)
		print.Printf100ms("cheque sig: %x\n", PayCheque.ChequeSig)
		print.Println100ms("")
	}
	// serialize paycheque
	var PayChequeMarshaled []byte
	PayChequeMarshaled, err = proto.Marshal(PayCheque)
	if err != nil {
		fmt.Println("marshal pay cheque failed when user store it.")
		return nil, errors.New("marshal paycheqe failed")
	}

	/////////////

	// calc hash from PayCheque
	hash = utils.CalcPayChequeHash(PayCheque)
	if global.DEBUG {
		print.Printf100ms("DEBUG> paycheque hash: %x\n", hash)
	}
	// sign PayCheque by user' sk
	var userSkByte = []byte(global.StrUserSK)
	PayChequeSig, err := sigapi.Sign(hash, userSkByte)
	if err != nil {
		log.Print("sign error")
		return nil, errors.New("sign error")
	}

	if global.DEBUG {
		print.Printf100ms("payvalue: %d\n", PayCheque.PayValue)
		print.Printf100ms("DEBUG> paycheque sig: %x\n", PayChequeSig)
	}

	msg := utils.MergeSlice(PayChequeSig, PayChequeMarshaled)

	//////////////////
	// db: paycheque_sig | paycheque_marshaled
	err = user.UserDB.Put(ChequeKey, msg, nil)
	if err != nil {
		fmt.Println("db put pay cheque data error")
		return nil, errors.New("db put pay cheque failed")
	}

	// show table
	user.ListPayCheque()
	return PayCheque, nil
}

// show paycheque info by key
func (user *User) GetPayChequeByKey(key []byte) (*pb.PayCheque, error) {

	user.OpenDB()
	defer user.CloseDB()

	msg, err := user.UserDB.Get(key, nil)
	if err != nil {
		log.Println("db.Get error:", err)
		return nil, err
	}

	//PayChequeSig := msg[:65]
	PayChequeMarshaled := msg[65:]
	// unmarshal it to get Cheque itself
	PayCheque := &pb.PayCheque{}
	if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
		log.Println("Failed to parse pay check:", err)
		return nil, err
	}

	//print.PrintPayCheque(PayCheque)
	//print.Printf100ms("PayChequeSig: %x\n", PayChequeSig)

	return PayCheque, nil
}

//
func (user *User) ShowPayChequeByID() error {

	// show user's paycheque table
	user.ListPayCheque()

	fmt.Println("Input ID to show:")
	var uID uint
	fmt.Scanf("%d", &uID)
	print.Printf100ms("-> You choosed %d\n", uID)

	keyByte, err := user.IDtoKey(uID)
	if err != nil {
		fmt.Println("ID to Key error: ", err)
		return err
	}

	//utils.ShowPayChequeInfoByKey(keyByte)
	PayCheque, err := user.GetPayChequeByKey(keyByte)
	if err != nil {
		fmt.Println("get paycheque error:", err)
		return err
	}

	print.PrintPayCheque(PayCheque)

	return nil

}

// user list data in pay cheque db
func (user *User) ListPayCheque() error {

	// update userIndex and storageIndex
	user.UpdatePayChequeIndex()

	// show table
	table, err := gotable.Create("ID", "FROM", "TO", "VALUE", "PAYVALUE", "NONCE")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// show table
	user.OpenDB()
	defer user.CloseDB()

	var id int = 0
	for id < len(user.DBIndex) {

		// get data
		var in []byte
		keyByte, err := hex.DecodeString(user.DBIndex[id])
		if err != nil {
			fmt.Println("decodeString error:", err.Error())
			return err
		}
		in, err = user.UserDB.Get(keyByte, nil)
		if err != nil {
			fmt.Println("db get error: ", err)
			return err
		}

		PayChequeMarshaled := in[65:]

		// unmarshal it to get Cheque itself
		PayCheque := &pb.PayCheque{}
		if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
			fmt.Println("Failed to parse paycheck:", err)
			return err
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
			fmt.Println(err)
			return err
		}
		id++
	}

	//r, _ := table.Json(4)
	//fmt.Println(r)
	//table.CloseBorder()
	table.PrintTable()

	return nil
}

// Update Index
func (user *User) UpdatePayChequeIndex() {

	user.OpenDB()
	defer user.CloseDB()

	// clear index
	user.DBIndex = user.DBIndex[0:0]

	iter := user.UserDB.NewIterator(nil, nil)
	for iter.Next() {
		keyByte := iter.Key()
		user.DBIndex = append(user.DBIndex, hex.EncodeToString(keyByte))
		fmt.Printf("index len: %d, cap: %d\n", len(user.DBIndex), cap(user.DBIndex))
		fmt.Printf("index0: %s\n", (user.DBIndex)[0])
	}
	iter.Release()

	err := iter.Error()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// user increase the pay value of paycheque in db by 'global.Increase'
func (user *User) IncPayValueByKey(key []byte) error {

	user.OpenDB()
	defer user.CloseDB()

	msg, err := user.UserDB.Get(key, nil)
	if err != nil {
		fmt.Println("db.Get error:", err)
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
		print.Printf100ms("-> Pay cheque before increased:%s\n", PayCheque.PayValue)
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
		print.Printf100ms("-> Pay cheque after increased:%s\n", PayCheque.PayValue)
	}

	// serialize
	PayChequeMarshaled, err = proto.Marshal(PayCheque)
	if err != nil {
		fmt.Println("marshal paycheque error:", err)
		return err
	}

	// recompute paycheque sig
	hash := utils.CalcPayChequeHash(PayCheque)
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
	msg = utils.MergeSlice(PayChequeSig, PayChequeMarshaled)

	// put into db
	err = user.UserDB.Put(key, msg, nil)
	if err != nil {
		fmt.Println("put paycheque error:", err)
		return err
	}

	return nil
}

//
func (user *User) SendOnePayChequeByID() error {
	if !global.RemoteExist {
		fmt.Println("No remote peer exist, record one as storage.")
		return errors.New("no remote peer")
	}

	user.ListPayCheque()

	print.Println100ms("-> Choose cheque ID to send.")
	var uID uint
	fmt.Scanf("%d", &uID)
	print.Printf100ms("-> You choosed %d\n", uID)

	// get key from id
	keyByte, err := user.IDtoKey(uID)
	if err != nil {
		fmt.Println("ID to Key error:", err)
		return err
	}

	user.SendPayChequeByKey(keyByte)

	return nil
}

//
func (user *User) IncAndSendPayChequeByID() error {
	if !global.RemoteExist {
		fmt.Println("No remote peer exist, record one as storage.")
		return errors.New("no remote peer")
	}

	user.ListPayCheque()

	print.Println100ms("-> Choose cheque ID to send.")
	var uID uint
	fmt.Scanf("%d", &uID)
	print.Printf100ms("-> You choosed %d\n", uID)

	// get key from id
	keyByte, err := user.IDtoKey(uID)
	if err != nil {
		fmt.Println("ID to Key error:", err)
		return err
	}

	err = user.IncPayValueByKey(keyByte)
	if err != nil {
		fmt.Println("inc payvalue error:", err)
		return err
	}

	err = user.SendPayChequeByKey(keyByte)
	if err != nil {
		fmt.Println("send paycheque error:", err)
		return err
	}

	return nil
}

//
func (user *User) SendPayChequeByKey(key []byte) error {

	user.OpenDB()
	defer user.CloseDB()

	print.Printf100ms("Opening stream to peerID: %v\n", global.Peerid)
	s, err := hostops.HostInfo.NewStream(context.Background(), global.Peerid, "/2")
	if err != nil {
		log.Println("open stream error: ", err)
		return err
	}

	msg, err2 := user.UserDB.Get(key, nil)
	if err2 != nil {
		log.Println("db.Get error:", err2)
		return err2
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

//
func (user *User) OpenDB() error {
	db, err := leveldb.OpenFile(user.DBfile, nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return err
	}
	user.UserDB = db

	return nil
}

// close db
func (user *User) CloseDB() error {
	user.UserDB.Close()

	return nil
}

// clear db
func (user *User) ClearDB() error {

	user.OpenDB()

	iter := user.UserDB.NewIterator(nil, nil)
	for iter.Next() {
		user.UserDB.Delete(iter.Key(), nil)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Println("user iter db error: ", err)
		return err
	}

	defer user.ListPayCheque()
	defer user.CloseDB()

	return nil
}

// delete an entry from user db, true for user, false for storage
func (user *User) DeleteChequeByID() {

	if len(user.DBIndex) == 0 {
		fmt.Println("user db is empty")
		return
	}

	user.ListPayCheque()

	user.OpenDB()

	print.Println100ms("Input ID to delete:")
	var uID uint
	fmt.Scanf("%d", &uID)
	if !(uID < uint(len(user.DBIndex))) {
		fmt.Println("Invalid input")
		return
	}
	if user.DBIndex[uID] == "" {
		fmt.Println("ID not exist")
		return
	}

	keyByte, err := user.IDtoKey(uID)
	if err != nil {
		fmt.Println("decode string error: ", err)
	}

	err = user.UserDB.Delete(keyByte, nil)
	if err != nil {
		fmt.Println("delete user db error: ", err)
	}
	print.Printf100ms("delete ID %d success.\n", uID)

	defer user.ListPayCheque()
	defer user.CloseDB()

}

// transmit ID of a pay cheque to key
func (user *User) IDtoKey(uID uint) ([]byte, error) {

	if !(uID < uint(len(user.DBIndex))) {
		fmt.Println("Invalid input")
		return nil, errors.New("invalid input")
	}
	if user.DBIndex[uID] == "" {
		fmt.Println("ID not exist")
		return nil, errors.New("ID not exist")
	}

	keyByte, err := hex.DecodeString(user.DBIndex[uID])
	if err != nil {
		fmt.Println("decode string error: ", err)
		return nil, err
	}

	return keyByte, nil
}
