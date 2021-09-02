package execmd

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/protobuf/proto"

	"github.com/ethereum/go-ethereum/common"

	"github.com/rockiecn/interact/callstorage"
	"github.com/rockiecn/p2pdemo/callcash"
	"github.com/rockiecn/p2pdemo/cash"
	"github.com/rockiecn/p2pdemo/global"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/p2pdemo/print"
	"github.com/rockiecn/p2pdemo/sigapi"
	"github.com/rockiecn/p2pdemo/utils"
)

// deploy cash
func DeployCash() {
	fmt.Println("call deploy cash")
	callcash.CallDeploy()
}

// user get Cheque from operator
func GetCheque() {

	if !global.RemoteExist {
		fmt.Println("No remote peer exist, record one as operator.")
		return
	}
	// connect to operator , get stream
	s, err := hostops.HostInfo.NewStream(context.Background(), global.Peerid, "/1")
	if err != nil {
		log.Println(err)
		return
	}

	// Read from stream
	fmt.Println("--> user receive Cheque from operator")
	in, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatalln("Error reading :", err)
		return
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
		return
	}

	fmt.Println("<signature of Cheque verify success>")

	// create/open db
	db, err := leveldb.OpenFile("./paycheque.db", nil)
	if err != nil {
		log.Fatal("opfen db error")
	}

	// Cheque key: To + nonce
	ChequeKey, err := utils.GenChequeKey(Cheque)
	if err != nil {
		log.Fatal("GenChequeKey error")
		return
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
		print.Printf100ms("Value: %d\n", Cheque.Value)
		print.Printf100ms("TokenAddress: %s\n", Cheque.TokenAddress)
		print.Printf100ms("Nonce: %d\n", Cheque.Nonce)
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
		return
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
		return
	}

	if global.DEBUG {
		print.Printf100ms("payvalue: %d\n", PayCheque.PayValue)
		print.Printf100ms("DEBUG> paycheque sig: %x\n", PayChequeSig)
	}

	msg := utils.MergeSlice(PayChequeSig, PayChequeMarshaled)

	//////////////////

	// db: paycheque_sig | paycheque_marshaled
	err = db.Put(ChequeKey, msg, nil)
	if err != nil {
		fmt.Println("db put pay cheque data error")
		return
	}

	//
	db.Close()

	// show table
	utils.ListPayCheque(true)
}

// user send marshaled PayCheques to storage
func SendOnePayChequeByID() {
	if !global.RemoteExist {
		fmt.Println("No remote peer exist, record one as storage.")
		return
	}

	utils.ListPayCheque(true)

	print.Println100ms("-> Choose cheque ID to send.")
	var uID uint
	fmt.Scanf("%d", &uID)
	print.Printf100ms("-> You choosed %d\n", uID)

	// get key from id
	keyByte, err := utils.IDtoKey(uID, true)
	if err != nil {
		fmt.Println("ID to Key error", err)
		return
	}

	utils.SendChequeByKey(keyByte)
}

// list db, true for user, false for storage
func ListDB(user bool) {
	utils.ListPayCheque(user)
}

// delete an entry from user db, true for user, false for storage
func DeleteChequeByID(user bool) {

	var dbfile string
	var Index []string

	if user {
		dbfile = "./paycheque.db"
		Index = global.UserIndex
	} else {
		dbfile = "./storage_paycheque.db"
		Index = global.StorageIndex
	}

	utils.ListPayCheque(user)

	db, err := leveldb.OpenFile(dbfile, nil)
	if err != nil {
		log.Fatal("opfen db error")
	}
	print.Println100ms("Input ID to delete:")
	var uID uint
	fmt.Scanf("%d", &uID)
	if !(uID < uint(len(Index))) {
		fmt.Println("Invalid input")
		return
	}
	if Index[uID] == "" {
		fmt.Println("ID not exist")
		return
	}

	var keyByte []byte
	keyByte, err = hex.DecodeString(Index[uID])
	if err != nil {
		fmt.Println("decode string error: ", err)
	}

	err = db.Delete(keyByte, nil)
	if err != nil {
		fmt.Println("delete user db error: ", err)
	}
	print.Printf100ms("delete ID %d success.\n", uID)

	db.Close()

	utils.ListPayCheque(user)
}

//
func ShowPayChequeByID() {

	var Index []string = global.UserIndex

	// show user's paycheque table
	utils.ListPayCheque(true)

	fmt.Println("Input ID to show:")
	var uID uint
	fmt.Scanf("%d", &uID)
	if !(uID < uint(len(Index))) {
		fmt.Println("Invalid input")
		return
	}
	if Index[uID] == "" {
		fmt.Println("ID not exist")
		return
	}

	keyByte, err := hex.DecodeString(Index[uID])
	if err != nil {
		fmt.Println("decode string error: ", err)
	}

	utils.ShowPayChequeInfoByKey(keyByte)
}

// increase payvalue, then send paycheque to storage
func IncAndSendCheque() {

	utils.ListPayCheque(true)

	print.Println100ms("Choose a cheque to continue:")

	var uID uint
	fmt.Scanf("%d", &uID)
	if global.UserIndex[uID] == "" {
		fmt.Println("ID not exist")
		return
	}

	keyByte, err := utils.IDtoKey(uID, true)
	if err != nil {
		fmt.Println("", err)
		return
	}

	// increase pay value in db
	err = utils.IncPayValueByKey(keyByte)
	if err != nil {
		fmt.Println("", err)
		return
	}

	//
	err = utils.SendChequeByKey(keyByte)
	if err != nil {
		fmt.Println("", err)
		return
	}

	utils.ListPayCheque(true)
}

// storage call cash contract
func StorageCallCash() {
	fmt.Println("-> Call contract")

	utils.ListPayCheque(false)

	print.Println100ms("-> Choose a cheque ID:")
	var uID uint
	fmt.Scanf("%d", &uID)
	print.Printf100ms("-> You choosed %d\n", uID)

	// get key from id
	keyByte, err0 := utils.IDtoKey(uID, false)
	if err0 != nil {
		fmt.Println("id to key error")
		return
	}

	print.Printf100ms(("PayCheque key: %x\n"), keyByte)

	// read PayCheque data from db
	// create/open db
	db, err1 := leveldb.OpenFile("./storage_paycheque.db", nil)
	if err1 != nil {
		fmt.Println("opfen db error")
		return
	}
	defer db.Close()

	// get paycheque
	in, err2 := db.Get(keyByte, nil)
	if err2 != nil {
		fmt.Println("db get error")
		return
	}

	PayChequeSig := in[:65]
	PayChequeMarshaled := in[65:]

	print.Println100ms("---- in StorageCallCash")
	//fmt.Println("PayChequeSig:", PayChequeSig)
	//fmt.Println("PayChequeMarshaled:", PayChequeMarshaled)

	// unmarshal it to get PayCheque itself
	PayCheque := &pb.PayCheque{}
	if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
		log.Println("Failed to parse paycheck:", err)
		return
	}

	// eth to wei
	// z18 := new(big.Int)
	// z18.SetString("1000000000000000000", 10)
	// weiPay := new(big.Int)
	// weiPay.Mul(bigPay, z18) // eth to wei

	// cheque
	var paychequeContract cash.PayCheque
	//bigValue := big.NewInt(PayCheque.Cheque.Value)
	bigValue := big.NewInt(0)
	var ok bool
	bigValue, ok = bigValue.SetString(PayCheque.Cheque.Value, 10)
	if !ok {
		print.Println100ms("big.SetString failed")
		return
	}

	paychequeContract.Cheque.Value = bigValue
	paychequeContract.Cheque.TokenAddr = common.HexToAddress(PayCheque.Cheque.TokenAddress)
	//bigNonce := big.NewInt(PayCheque.Cheque.Nonce)
	bigNonce := big.NewInt(0)
	bigNonce, ok = bigNonce.SetString(PayCheque.Cheque.Nonce, 10)
	if !ok {
		print.Println100ms("big.SetString failed")
		return
	}

	paychequeContract.Cheque.Nonce = bigNonce
	paychequeContract.Cheque.FromAddr = common.HexToAddress(PayCheque.Cheque.From)
	paychequeContract.Cheque.ToAddr = common.HexToAddress(PayCheque.Cheque.To)
	paychequeContract.Cheque.OpAddr = common.HexToAddress(PayCheque.Cheque.OperatorAddress)
	paychequeContract.Cheque.ContractAddr = common.HexToAddress(PayCheque.Cheque.ContractAddress)
	// paycheque
	paychequeContract.ChequeSig = PayCheque.ChequeSig
	//bigPayValue := big.NewInt(PayCheque.PayValue)
	bigPayValue := big.NewInt(0)
	bigPayValue, ok = bigPayValue.SetString(PayCheque.PayValue, 10)
	if !ok {
		print.Println100ms("big.SetString failed")
		return
	}

	paychequeContract.PayValue = bigPayValue

	print.Println100ms("------------- show paycheque contract ---------------")
	print.Printf100ms("paychequeContract.Cheque.Value: %s\n", paychequeContract.Cheque.Value.String())
	print.Printf100ms("paychequeContract.Cheque.TokenAddr: %s\n", paychequeContract.Cheque.TokenAddr)

	print.Printf100ms("paychequeContract.Cheque.Nonce: %s\n", paychequeContract.Cheque.Nonce.String())
	print.Printf100ms("paychequeContract.Cheque.FromAddr: %s\n", paychequeContract.Cheque.FromAddr)
	print.Printf100ms("paychequeContract.Cheque.ToAddr: %s\n", paychequeContract.Cheque.ToAddr)
	print.Printf100ms("paychequeContract.Cheque.OpAddr: %s\n", paychequeContract.Cheque.OpAddr)
	print.Printf100ms("paychequeContract.ChequeSig: %x\n", paychequeContract.ChequeSig)
	print.Printf100ms("paychequeContract.PayValue: %s\n", paychequeContract.PayValue.String())
	print.Println100ms("")

	//errCallApply := callcash.CallApplyPayCheque(From, bigNonce, To, bigPay, PayChequeSig)
	errCallApply := callcash.CallApplyPayCheque(paychequeContract, PayChequeSig)
	if errCallApply != nil {
		fmt.Println("callApplyPayCheque error:", errCallApply)
		fmt.Println("storage address:", PayCheque.Cheque.To)
		fmt.Println("nonce:", PayCheque.Cheque.Nonce)
	}
}

func TestCall() {
	print.Println100ms("call retrieve")
	callstorage.CallRetrieve()
}

// delete all data of DB
func ClearDB(user bool) {
	var dbfile string
	if user {
		dbfile = "./paycheque.db"
	} else {
		dbfile = "./storage_paycheque.db"
	}

	db, err := leveldb.OpenFile(dbfile, nil)
	if err != nil {
		fmt.Println("", err)
		return
	}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		db.Delete(iter.Key(), nil)
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println("", err)
		return
	}
	defer db.Close()
}

// get contract nonce
func GetContractNonce() {

	AddressTo := common.HexToAddress(global.StrToAddr)
	//print.Printf100ms("address to :%s\n", AddressTo.String())

	err := callcash.CallGetNodeNonce(AddressTo)
	if err != nil {
		fmt.Println("call get nonce error: ", err)
		return
	}
}

// set nonce of operator db to 0.
func ResetNonceInOperatorDB() {

	// create/open db
	db, err := leveldb.OpenFile("./operator_data.db", nil)
	if err != nil {
		log.Fatal("open db error")
	}
	defer db.Close()

	//byteNonce := utils.Int64ToBytes(0)
	bigNonce := big.NewInt(0)

	// storage -> nonce
	err = db.Put([]byte(global.StrToAddr), bigNonce.Bytes(), nil)
	if err != nil {
		fmt.Println("reset nonce error: ", err)
		return
	}
}

func ShowNonceInOperatorDB() {

	// create/open db
	db, err := leveldb.OpenFile("./operator_data.db", nil)
	if err != nil {
		log.Fatal("open db error")
	}
	defer db.Close()

	//byteNonce := utils.Int64ToBytes(0)
	bigNonce := big.NewInt(0)

	// storage -> nonce
	byteNonce, err := db.Get([]byte(global.StrToAddr), nil)
	if err != nil {
		fmt.Println("get nonce error: ", err)
		return
	}
	bigNonce.SetBytes(byteNonce)
	fmt.Println("nonce:", bigNonce.String())
}
