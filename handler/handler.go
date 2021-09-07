package handler

import (
	"fmt"
	"log"
	"math/big"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/rockiecn/p2pdemo/global"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/p2pdemo/print"
	"github.com/rockiecn/p2pdemo/sigapi"
	"github.com/rockiecn/p2pdemo/utils"

	// "github.com/rockiecn/sigtest/sigapi"
	// "github.com/rockiecn/sigtest/utils"

	"google.golang.org/protobuf/proto"
)

// Cmd1Handler - command 1 handler, run on operator, send Cheque to user
func BuyChequeHandler(s network.Stream) error {

	print.Println100ms("--> Construct and send Cheque...")

	// construct Cheque
	Cheque := &pb.Cheque{}
	Cheque.Value = "100000000000000000000"          // Cheque 100
	Cheque.TokenAddress = global.StrTokenAddr       // token address
	Cheque.From = global.StrFromAddr                // user
	Cheque.To = global.StrToAddr                    // storage
	Cheque.OperatorAddress = global.StrOperatorAddr // operator

	// create/open db
	db, err := leveldb.OpenFile("./operator_data.db", nil)
	if err != nil {
		log.Fatal("open db error")
	}
	defer db.Close()

	//var newNonce int64 = 0
	// storage -> nonce
	nonce, err := db.Get([]byte(Cheque.To), nil)
	if err != nil {
		if err.Error() == "leveldb: not found" { // no nonce at all
			db.Put([]byte(Cheque.To), utils.Int64ToBytes(0), nil)
		} else {
			fmt.Println("operator db get nonce error: ", err)
			return err
		}
	}

	// increase nonce by 1
	// byte to string
	bigNewNonce := big.NewInt(0)
	bigOne := big.NewInt(1)
	//oldNonce := utils.BytesToInt64(nonce)
	bigOldNonce := big.NewInt(0)
	bigOldNonce = bigOldNonce.SetBytes(nonce)
	fmt.Println("bigOldNonce: ", bigOldNonce.String())
	bigNewNonce = bigOldNonce.Add(bigOldNonce, bigOne)
	fmt.Println("bigNewNonce: ", bigNewNonce.String())

	// put new nonce into db
	//byteN := utils.Int64ToBytes(newNonce)
	//fmt.Println("newNonce: ", newNonce)
	//fmt.Printf("byteN: %v\n", byteN)
	// err = db.Put([]byte(Cheque.To), byteN, nil)
	// if err != nil {
	// 	fmt.Println("operator db put nonce error")
	// 	return err
	// }

	err = db.Put([]byte(Cheque.To), bigNewNonce.Bytes(), nil)
	if err != nil {
		fmt.Println("operator db put nonce error")
		return err
	}

	//
	Cheque.Nonce = bigNewNonce.String()

	contractAddrByte, err := db.Get([]byte("contractAddr"), nil)
	if err != nil {
		log.Println("!! get cash address error:", err)
		return err
	}

	// contract address, delete prefix '0x'
	contractAddrByte = contractAddrByte[2:]
	Cheque.ContractAddress = string(contractAddrByte) // contract address

	// if global.DEBUG {
	// 	print.Printf100ms("sigByte:%x\n", sigByte)
	// 	print.Printf100ms("contractAddr:%s\n", Cheque.ContractAddress)
	// 	print.Printf100ms("ChequeMarshaled:%x\n", ChequeMarshaled)
	// }

	// serialize
	ChequeMarshaled, err := proto.Marshal(Cheque)
	if err != nil {
		log.Fatalln("Failed to encode PayCheque:", err)
	}

	// construct Cheque message: sig(65 bytes) | data
	print.Println100ms("-> constructing msg")

	//hash := utils.CalcHash(Cheque.From, Cheque.Nonce, "", 0)
	// calc cheque hash
	hash := utils.CalcChequeHash(Cheque)

	print.Printf100ms("Cheque send, hash: %x\n", hash)

	// sign Cheque by operator
	var opSkByte = []byte(global.StrOperatorSK)

	sigByte, err := sigapi.Sign(hash, opSkByte)
	if err != nil {
		log.Fatal("sign error:", err)
	}

	// sig(65) | cheque
	var msg = []byte{}
	msg = utils.MergeSlice(sigByte, ChequeMarshaled)

	print.Println100ms("-> sending msg")
	// send msg
	_, err = s.Write([]byte(msg))
	if err != nil {
		panic("stream write error")
	}

	print.Println100ms("\n> Intput target cmd: ")

	return err
}

/*
// Cmd2Handler - command 2 handler, run on storage, receive PayCheque from user and record to db
func SendCheckHandler(s network.Stream) error {

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
	// create/open db
	db, err := leveldb.OpenFile("./storage_paycheque.db", nil)
	if err != nil {
		log.Fatal("opfen db error")
	}

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
	err = db.Put(ChequeKey, in, nil)
	if err != nil {
		print.Println100ms("db put data error")
		return err
	}

	defer utils.ListPayCheque(false)
	defer db.Close()

	print.PrintMenu()
	print.Println100ms("\n> Intput target address and cmd: ")

	return nil
}
*/
