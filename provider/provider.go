package provider

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/liushuochen/gotable"
	"google.golang.org/protobuf/proto"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/rockiecn/p2pdemo/cash"
	"github.com/rockiecn/p2pdemo/clientops"
	"github.com/rockiecn/p2pdemo/db"
	"github.com/rockiecn/p2pdemo/global"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/p2pdemo/print"
	"github.com/rockiecn/p2pdemo/sigapi"
	"github.com/rockiecn/p2pdemo/utils"
)

type Provider struct {
	ProDB *leveldb.DB // provider -> nonce
	//ContractAddress string      // contract address

	DBfile  string
	DBIndex []string

	ProviderAddr string // "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"
	ProviderSK   string // "cc6d63f85de8fef05446ebdd3c537c72152d0fc437fd7aa62b3019b79bd1fdd4"

	TokenAddr string // token address
	FromAddr  string // user
	ToAddr    string // storage

	db.DB
}

// init provider, need db open first
func (pro *Provider) Init() error {

	pro.DBfile = "./provider.db"
	pro.OpenDB()

	//pro.ContractAddress = ""

	pro.DBIndex = []string{}

	pro.ProviderAddr = "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"
	pro.ProviderSK = "cc6d63f85de8fef05446ebdd3c537c72152d0fc437fd7aa62b3019b79bd1fdd4"

	pro.TokenAddr = "b213d01542d129806d664248a380db8b12059061" // token address
	pro.FromAddr = "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"  // user
	pro.ToAddr = "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"    // storage

	byteAddr, err := pro.ProDB.Get([]byte("contractAddr"), nil)
	if err != nil {
		return errors.New("operator init: read contract address failed")
	}
	print.ContractAddress = string(byteAddr)

	return nil
}

func (pro *Provider) CallContract() error {
	return nil
}

// Cmd2Handler - command 2 handler, run on storage, receive PayCheque from user and record to db
func (pro *Provider) RecievePayCheck(s network.Stream) error {

	/*
		// // Read data method 1
		// in := make([]byte, 1024)
		// reader := bufio.NewReader(s)
		// n, err := reader.Read(in)
		// if err != nil {
		// 	fmt.Println("read err: ", err)
		// }
		// // get real data
		// if n > 0 {
		// 	in = in[:n]
		// }
		// fmt.Printf("in: %v", in)

		// // Read data method 2
		// reader := bufio.NewReader(s)
		// in, err := reader.ReadBytes('\n')
		// if err != nil {
		// 	return err
		// }
		// fmt.Printf("read: %v", in)
	*/
	// Read data method 3
	// Caution: Need writer to close stream first.
	in, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatalln("Error reading :", err)
		return err
	}

	// parse data
	var sigByte = in[:65]
	var PayChequeMarshaled = in[65:]

	// unmarshal data
	PayCheque := &pb.PayCheque{}
	if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
		log.Println("Failed to parse paycheck:", err)
		return err
	}

	//
	print.PrintPayCheque(PayCheque)

	//===== verify signature of PayCheque(signed by user)

	// get user address
	fromAddrByte, err := hex.DecodeString(PayCheque.Cheque.From)
	if err != nil {
		panic("decode error")
	}
	// []byte to common.Address
	From := common.BytesToAddress(fromAddrByte)

	// calc hash from PayCheque
	//hash := utils.CalcHash(PayCheque.Cheque.From, PayCheque.Cheque.Nonce, PayCheque.Cheque.To, PayCheque.PayValue)
	hash := utils.CalcPayChequeHash(PayCheque)
	fmt.Printf("paycheque hash:%x\n", hash)

	if global.DEBUG {
		fmt.Printf("hash: %x\n", hash)
		fmt.Printf("sigByte: %x\n", sigByte)
	}

	// verify PayCheque signature: []byte []byte common.Address
	ok, verErr := sigapi.Verify(hash, sigByte, From)
	if verErr != nil {
		log.Fatal("verify fatal error occured")
		return verErr
	}

	if !ok {
		print.Println100ms("<signature of PayCheque verify failed>")
	}

	print.Println100ms("<signature of PayCheque verify success>")

	// wirte PayCheque into db

	// gen Cheque key: To + nonce
	var ChequeKey []byte
	//ChequeKey, err = utils.GenChequeKey(PayCheque.Cheque.To, bigNonce)
	ChequeKey, err = utils.GenChequeKey(PayCheque.Cheque)
	if err != nil {
		log.Fatal("GenChequeKey error:", err)
	}

	if global.DEBUG {
		fmt.Println("in Cmd2Handler.")
		fmt.Printf("ChequeKey: %x\n", ChequeKey)
	}

	// use PayChequeKey as PayCheque id to store PayChequeMarshaled | paychequeSig
	err = pro.ProDB.Put(ChequeKey, in, nil)
	if err != nil {
		print.Println100ms("db put data error")
		return err
	}

	pro.ListPayCheque()

	print.PrintMenu()
	print.Println100ms("\n> Intput target address and cmd: ")

	return nil
}

//
func (pro *Provider) OpenDB() error {
	db, err := leveldb.OpenFile(pro.DBfile, nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return err
	}
	pro.ProDB = db

	return nil
}

// close db
func (pro *Provider) CloseDB() error {
	pro.ProDB.Close()

	return nil
}

// clear db
func (pro *Provider) ClearDB() error {

	iter := pro.ProDB.NewIterator(nil, nil)
	for iter.Next() {
		pro.ProDB.Delete(iter.Key(), nil)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Println("user iter db error: ", err)
		return err
	}

	pro.ListPayCheque()

	return nil
}

// delete an entry from user db, true for user, false for storage
func (pro *Provider) DeleteChequeByID() error {

	if len(pro.DBIndex) == 0 {
		fmt.Println("provider db is empty")
		return errors.New("provider db is empty")
	}

	pro.ListPayCheque()

	print.Println100ms("Input ID to delete:")
	var uID uint
	fmt.Scanf("%d", &uID)
	if !(uID < uint(len(pro.DBIndex))) {
		fmt.Println("Invalid input")
		return errors.New("invalid input")
	}
	if pro.DBIndex[uID] == "" {
		fmt.Println("ID not exist")
		return errors.New("ID not exist")
	}

	keyByte, err := hex.DecodeString(pro.DBIndex[uID])
	if err != nil {
		fmt.Println("decode string error: ", err)
		return err
	}

	err = pro.ProDB.Delete(keyByte, nil)
	if err != nil {
		fmt.Println("delete provider db error: ", err)
	}
	print.Printf100ms("delete ID %d success.\n", uID)

	pro.ListPayCheque()

	return nil

}

// list provider's db
func (pro *Provider) ListPayCheque() error {

	// update userIndex and storageIndex
	pro.UpdatePayChequeIndex()

	// show table
	table, err := gotable.Create("ID", "FROM", "TO", "VALUE", "PAYVALUE", "NONCE")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// show table

	var id int = 0
	for id < len(pro.DBIndex) {

		// get data
		var in []byte
		keyByte, err := hex.DecodeString(pro.DBIndex[id])
		if err != nil {
			fmt.Println("decodeString error:", err.Error())
			return err
		}
		in, err = pro.ProDB.Get(keyByte, nil)
		if err != nil {
			fmt.Println("db get error: ", err)
			return err
		}

		PayChequeMarshaled := in[65:]

		// unmarshal it to get Cheque itself
		PayCheque := &pb.PayCheque{}
		if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
			log.Println("Failed to parse paycheck:", err)
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
			log.Fatal(err.Error())
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
func (pro *Provider) UpdatePayChequeIndex() {

	// clear index
	pro.DBIndex = pro.DBIndex[0:0]

	iter := pro.ProDB.NewIterator(nil, nil)
	for iter.Next() {
		keyByte := iter.Key()
		pro.DBIndex = append(pro.DBIndex, hex.EncodeToString(keyByte))
		fmt.Printf("index len: %d, cap: %d\n", len(pro.DBIndex), cap(pro.DBIndex))
		fmt.Printf("index0: %s\n", (pro.DBIndex)[0])
	}
	iter.Release()

	err := iter.Error()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// transmit ID of a pay cheque to key
func (pro *Provider) IDtoKey(uID uint) ([]byte, error) {

	if !(uID < uint(len(pro.DBIndex))) {
		fmt.Println("Invalid input")
		return nil, errors.New("invalid input")
	}
	if pro.DBIndex[uID] == "" {
		fmt.Println("ID not exist")
		return nil, errors.New("ID not exist")
	}

	keyByte, err := hex.DecodeString(pro.DBIndex[uID])
	if err != nil {
		fmt.Println("decode string error: ", err)
		return nil, err
	}

	return keyByte, nil
}

// storage call cash contract
func (pro *Provider) CallCashByID() error {
	fmt.Println("-> Call contract")

	pro.ListPayCheque()

	print.Println100ms("-> Choose a cheque ID:")
	var uID uint
	fmt.Scanf("%d", &uID)
	print.Printf100ms("-> You choosed %d\n", uID)

	// get key from id
	keyByte, err := pro.IDtoKey(uID)
	if err != nil {
		fmt.Println("id to key error")
		return err
	}

	print.Printf100ms(("PayCheque key: %x\n"), keyByte)

	// get paycheque
	in, err2 := pro.ProDB.Get(keyByte, nil)
	if err2 != nil {
		fmt.Println("db get error")
		return err
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
		return err
	}

	// cheque
	var paychequeContract cash.PayCheque
	//bigValue := big.NewInt(PayCheque.Cheque.Value)
	bigValue := big.NewInt(0)
	var ok bool
	bigValue, ok = bigValue.SetString(PayCheque.Cheque.Value, 10)
	if !ok {
		print.Println100ms("big.SetString failed")
		return errors.New("big.SetString failed")
	}

	paychequeContract.Cheque.Value = bigValue
	paychequeContract.Cheque.TokenAddr = common.HexToAddress(PayCheque.Cheque.TokenAddress)
	//bigNonce := big.NewInt(PayCheque.Cheque.Nonce)
	bigNonce := big.NewInt(0)
	bigNonce, ok = bigNonce.SetString(PayCheque.Cheque.Nonce, 10)
	if !ok {
		print.Println100ms("big.SetString failed")
		return errors.New("big.SetString failed")
	}

	paychequeContract.Cheque.Nonce = bigNonce
	paychequeContract.Cheque.FromAddr = common.HexToAddress(PayCheque.Cheque.From)
	paychequeContract.Cheque.ToAddr = common.HexToAddress(PayCheque.Cheque.To)
	paychequeContract.Cheque.OpAddr = common.HexToAddress(PayCheque.Cheque.OperatorAddress)
	paychequeContract.Cheque.ContractAddr = common.HexToAddress(PayCheque.Cheque.ContractAddress)
	// paycheque
	paychequeContract.ChequeSig = PayCheque.ChequeSig
	bigPayValue := big.NewInt(0)
	bigPayValue, ok = bigPayValue.SetString(PayCheque.PayValue, 10)
	if !ok {
		print.Println100ms("big.SetString failed")
		return errors.New("big.SetString failed")
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
	errCallApply := pro.CallApplyPayCheque(paychequeContract, PayChequeSig)
	if errCallApply != nil {
		fmt.Println("callApplyPayCheque error:", errCallApply)
		fmt.Println("storage address:", PayCheque.Cheque.To)
		fmt.Println("nonce:", PayCheque.Cheque.Nonce)
	}

	return nil
}

// CallApplyCheque - send tx to contract to call apply cheque method.
func (pro *Provider) CallApplyPayCheque(paycheque cash.PayCheque, paychequeSig []byte) error {
	cli, err := clientops.GetClient(hostops.HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return err
	}
	defer cli.Close()

	auth, err := clientops.MakeAuth(pro.ProviderSK, nil, nil, big.NewInt(1000), 9000000)
	if err != nil {
		return err
	}

	// get contract instance from address
	cashInstance, err := cash.NewCash(paycheque.Cheque.ContractAddr, cli)
	if err != nil {
		fmt.Println("NewCash err: ", err)
		return err
	}

	print.Printf100ms("cheque.value: %s\n", paycheque.Cheque.Value)
	print.Printf100ms("cheque.TokenAddr: %s\n", paycheque.Cheque.TokenAddr)
	print.Printf100ms("cheque.Nonce: %s\n", paycheque.Cheque.Nonce.String())
	print.Printf100ms("cheque.FromAddr: %s\n", paycheque.Cheque.FromAddr)
	print.Printf100ms("cheque.ToAddr: %s\n", paycheque.Cheque.ToAddr)
	print.Printf100ms("cheque.OpAddr: %s\n", paycheque.Cheque.OpAddr)
	print.Printf100ms("cheque.ContractAddress: %s\n", paycheque.Cheque.ContractAddr)
	print.Printf100ms("paycheque.ChequeSig: %x\n", paycheque.ChequeSig)
	print.Printf100ms("paycheque.PayValue: %s\n", paycheque.PayValue.String())
	print.Printf100ms("paychequeSig: %x\n", paychequeSig)

	_, err = cashInstance.ApplyCheque(auth, paycheque, paychequeSig)
	if err != nil {
		fmt.Println("tx failed :", err)
		return err
	}

	fmt.Println("-> Now mine a block to complete tx.")

	return nil
}
