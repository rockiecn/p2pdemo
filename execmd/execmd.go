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
	print.Println100ms("call deploy cash")
	callcash.CallDeploy()
}

// user get Cheque from operator
func GetCheque() {

	if !global.RemoteExist {
		print.Println100ms("No remote peer exist, record one as operator.")
		return
	}
	// connect to operator , get stream
	s, err := hostops.HostInfo.NewStream(context.Background(), global.Peerid, "/1")
	if err != nil {
		log.Println(err)
		return
	}

	// Read from stream
	print.Println100ms("--> user receive Cheque from operator")
	in, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatalln("Error reading :", err)
		return
	}

	// parse data
	sigByte := in[:65]         //65
	cashAddrByte := in[65:107] //42
	ChequeMarshaled := in[107:]

	if global.DEBUG {
		print.Printf100ms("sigByte:%x\n", sigByte)
		print.Printf100ms("cashAddr:%s\n", cashAddrByte)
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
	hash := utils.CalcHash(Cheque.From, Cheque.NodeNonce, "", 0)
	if global.DEBUG {
		print.Printf100ms("Cheque receive, hash: %x\n", hash)
	}

	// verify Cheque signature
	ok, _ := sigapi.Verify(hash, sigByte, opAddress)
	if !ok {
		print.Println100ms("<signature of Cheque verify failed>")
		return
	}

	print.Println100ms("<signature of Cheque verify success>")

	// create/open db
	db, err := leveldb.OpenFile("./paycheque.db", nil)
	if err != nil {
		log.Fatal("opfen db error")
	}

	if global.DEBUG {
		print.Printf100ms("storage address: %s\n", Cheque.To)
		print.Printf100ms("nonce: %d\n", Cheque.NodeNonce)
	}

	// Cheque key: To + nonce
	bigNonce := big.NewInt(Cheque.NodeNonce)
	ChequeKey, err := utils.GenChequeKey(Cheque.To, bigNonce)
	if err != nil {
		log.Fatal("GenChequeKey error")
		return
	}

	if global.DEBUG {
		print.Printf100ms("ChequeKey: %x\n", ChequeKey)
	}

	//fmt.Printf("---------------- cash string:%s\n", string(cashAddrByte))
	// construct pay cheque with initial payvalue=0
	PayCheque := &pb.PayCheque{}
	PayCheque.Cheque = Cheque
	PayCheque.ChequeSig = sigByte
	PayCheque.CashAddress = string(cashAddrByte)
	PayCheque.From = Cheque.From
	PayCheque.To = Cheque.To
	PayCheque.PayValue = 0

	// serialize paycheque
	var PayChequeMarshaled []byte
	PayChequeMarshaled, err = proto.Marshal(PayCheque)
	if err != nil {
		print.Println100ms("marshal pay cheque failed when user store it.")
		return
	}
	err = db.Put(ChequeKey, PayChequeMarshaled, nil)
	if err != nil {
		print.Println100ms("db put pay cheque data error")
		return
	}

	// // use ChequeHash as Cheque id to store ChequeMarshaled.
	// ChequeMarshWithSig := utils.MergeSlice(sigByte, ChequeMarshaled)
	// err = db.Put(ChequeKey, ChequeMarshWithSig, nil)
	// if err != nil {
	// 	print.Println100ms("db put data error")
	// 	return
	// }

	//
	db.Close()

	// show table
	utils.ListPayCheque(true)

}

// user send marshaled PayCheques to storage
func SendOnePayChequeByID() {
	if !global.RemoteExist {
		print.Println100ms("No remote peer exist, record one as storage.")
		return
	}

	utils.ListPayCheque(true)

	fmt.Println("-> Choose cheque ID to send.")
	var uID uint
	fmt.Scanf("%d", &uID)
	fmt.Printf("-> You choosed %d\n", uID)

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
	fmt.Println("Input ID to delete:")
	var uID uint
	fmt.Scanf("%d", &uID)
	if !(uID < uint(len(Index))) {
		fmt.Println("Invalid input")
		return
	}
	if global.UserIndex[uID] == "" {
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
	fmt.Printf("delete ID %d success.\n", uID)

	db.Close()

	utils.ListPayCheque(user)
}

// increase payvalue, then send paycheque to storage
func IncAndSendCheque() {

	utils.ListPayCheque(true)

	fmt.Println("Choose a cheque to continue:")

	var uID uint
	fmt.Scanf("%d", &uID)
	if global.UserIndex[uID] == "" {
		fmt.Println("ID not exist")
		return
	}

	keyByte, err := utils.IDtoKey(uID, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	// increase pay value in db
	err = utils.IncPayValueByKey(keyByte)
	if err != nil {
		fmt.Println(err)
		return
	}

	//
	err = utils.SendChequeByKey(keyByte)
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.ListPayCheque(true)
}

// storage call cash contract
func StorageCallCash() {
	print.Println100ms("-> Call contract")

	utils.ListPayCheque(false)

	fmt.Println("-> Choose a cheque ID:")
	var uID uint
	fmt.Scanf("%d", &uID)
	fmt.Printf("-> You choosed %d\n", uID)

	// get key from id
	keyByte, err0 := utils.IDtoKey(uID, false)
	if err0 != nil {
		fmt.Println("id to key error")
		return
	}

	// read PayCheque data from db
	// create/open db
	db, err1 := leveldb.OpenFile("./storage_paycheque.db", nil)
	if err1 != nil {
		fmt.Println("opfen db error")
		return
	}
	defer db.Close()

	// get paycheque
	PayChequeMarshaled, err2 := db.Get(keyByte, nil)
	if err2 != nil {
		fmt.Println("db get error")
		return
	}

	fmt.Printf(("PayCheque key: %x\n"), keyByte)

	// unmarshal it to get PayCheque itself
	PayCheque := &pb.PayCheque{}
	if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
		log.Fatalln("Failed to parse check:", err)
	}

	/*
		// string to common.Address
		From := common.HexToAddress(PayCheque.Cheque.From)
		// int to bigInt, nonce
		bigNonce := big.NewInt(PayCheque.Cheque.NodeNonce)
		// get storage address
		toBytes, err := hex.DecodeString(PayCheque.To)
		if err != nil {
			panic("decode error")
		}
		// []byte to common.Address
		To := common.BytesToAddress(toBytes)
		// pay amount big
		bigPay := big.NewInt(PayCheque.PayValue)
	*/

	// eth to wei
	// z18 := new(big.Int)
	// z18.SetString("1000000000000000000", 10)
	// weiPay := new(big.Int)
	// weiPay.Mul(bigPay, z18) // eth to wei

	/*
		message Cheque {
			string operator_address = 1; // operator
			string from = 2;	// user
			string to =3;	// storage
			string token_address = 4;	// token

			int64 value = 5;
			int64 node_nonce = 6;
		}
		message PayCheque {
			Cheque cheque = 1;
			bytes cheque_sig = 2; //运营商对cheque的签名

			string cash_address = 3; //运营合约地址
			string from = 4; //user地址
			string to = 5; //storage
			int64 pay_value = 6; //支付给存储节点的数额必须小于等于cheque.max_amount
		}

		[]string:
		Cheque - string operator_address = 1;
		Cheque - string from = 2;
		Cheque - string to =3;
		Cheque - string token_address = 4;
		PayCheque - string cash_address = 3;

		[]int64:
		Cheque - int64 value = 5;
		Cheque - int64 node_nonce = 6;
		PayCheque - int64 pay_value = 6;

		[]byte:
		PayCheque - bytes cheque_sig = 2; //运营商对cheque的签名
	*/
	var stringParams = []string{}
	stringParams = append(stringParams, PayCheque.Cheque.OperatorAddress)
	stringParams = append(stringParams, PayCheque.Cheque.From)
	stringParams = append(stringParams, PayCheque.Cheque.To)
	stringParams = append(stringParams, PayCheque.Cheque.TokenAddress)
	stringParams = append(stringParams, PayCheque.CashAddress)

	var intParams = []int64{}
	intParams = append(intParams, PayCheque.Cheque.Value)
	intParams = append(intParams, PayCheque.Cheque.NodeNonce)
	intParams = append(intParams, PayCheque.PayValue)

	bytesParam := PayCheque.ChequeSig

	fmt.Printf("show params\n")
	fmt.Printf("0x%s\n", stringParams[0])
	fmt.Printf("0x%s\n", stringParams[1])
	fmt.Printf("0x%s\n", stringParams[2])
	fmt.Printf("0x%s\n", stringParams[3])
	fmt.Printf("%s\n", stringParams[4])

	fmt.Printf("%d\n", intParams[0])
	fmt.Printf("%d\n", intParams[1])
	fmt.Printf("%d\n", intParams[2])

	fmt.Printf("%x\n", bytesParam)

	var paychequeContract cash.PayCheque
	paychequeContract.Cheque.OpAddr = PayCheque.Cheque.OperatorAddress
	paychequeContract.Cheque.FromAddr = PayCheque.Cheque.From
	paychequeContract.Cheque.ToAddr = PayCheque.Cheque.To
	paychequeContract.Cheque.TokenAddr = PayCheque.Cheque.TokenAddress

	bigValue := big.NewInt(PayCheque.Cheque.Value)
	paychequeContract.Cheque.Value = bigValue
	bigNonce := big.NewInt(PayCheque.Cheque.NodeNonce)
	paychequeContract.Cheque.NodeNonce = bigNonce

	//errCallApply := callcash.CallApplyPayCheque(From, bigNonce, To, bigPay, PayChequeSig)
	errCallApply := callcash.CallApplyPayCheque(paychequeContract)
	if errCallApply != nil {
		fmt.Println("callApplyPayCheque error:", errCallApply)
		fmt.Println("storage address:", PayCheque.Cheque.To)
		fmt.Println("nonce:", PayCheque.Cheque.NodeNonce)
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
		fmt.Println(err)
		return
	}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		db.Delete(iter.Key(), nil)
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
}
