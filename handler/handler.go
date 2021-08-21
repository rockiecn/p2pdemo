package handler

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

// Cmd1Handler - command 1 handler, run on user, receive Cheque from operator
func BuyCheckHandler(s network.Stream) error {

	print.Println100ms("--> Construct and send Cheque...")

	// construct Cheque
	Cheque := &pb.Cheque{}
	Cheque.Value = 100 // Cheque 100
	//Cheque.NodeNonce = 1

	Cheque.OperatorAddress = "9e0153496067c20943724b79515472195a7aedaa" // operator
	Cheque.From = "1ab6a9f2b90004c1269563b5da391250ede3c114"            // user
	Cheque.To = "0xb213d01542d129806d664248a380db8b12059061"            // storage
	Cheque.TokenAddress = "tokenaddress"

	// create/open db
	db, err := leveldb.OpenFile("./operator_data.db", nil)
	if err != nil {
		log.Fatal("open db error")
	}
	defer db.Close()

	var newNonce int64 = 0
	// storage -> nonce
	ret, err := db.Get([]byte(Cheque.To), nil)
	if err != nil {
		if err.Error() == "leveldb: not found" { // no nonce at all
			db.Put([]byte(Cheque.To), utils.Int64ToBytes(1), nil)
		} else {
			fmt.Println("operator db get nonce error: ", err)
			return err
		}
	} else { // increase nonce by 1
		// byte to string
		oldNonce := utils.BytesToInt64(ret)
		fmt.Println("oldNonce: ", oldNonce)
		newNonce = oldNonce + 1
		// put new nonce into db
		byteN := utils.Int64ToBytes(newNonce)
		fmt.Println("newNonce: ", newNonce)
		fmt.Printf("byteN: %v\n", byteN)
		err = db.Put([]byte(Cheque.To), byteN, nil)
		if err != nil {
			fmt.Println("operator db put nonce error")
			return err
		}
	}

	//
	Cheque.NodeNonce = newNonce

	// serialize
	ChequeMarshaled, err := proto.Marshal(Cheque)
	if err != nil {
		log.Fatalln("Failed to encode PayCheque:", err)
	}

	// construct Cheque message: sig(65 bytes) | data
	print.Println100ms("-> constructing msg")

	// calc hash
	hash := utils.CalcHash(Cheque.From, Cheque.NodeNonce, "", 0)
	print.Printf100ms("Cheque send, hash: %x\n", hash)

	// sign Cheque by operator
	var opSkByte = []byte("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
	sigByte, err := sigapi.Sign(hash, opSkByte)
	if err != nil {
		log.Fatal("sign error:", err)
	}

	cashAddrByte, err := db.Get([]byte("cashAddr"), nil)
	if err != nil {
		log.Fatal("get cash address error:", err)
	}

	if global.DEBUG {
		print.Printf100ms("sigByte:%x\n", sigByte)
		print.Printf100ms("cashAddr:%x\n", cashAddrByte)
		print.Printf100ms("ChequeMarshaled:%x\n", ChequeMarshaled)
	}

	// sig(65) | cash address(42) | cheque
	var msg = []byte{}
	msg = utils.MergeSlice(sigByte, cashAddrByte)
	msg = utils.MergeSlice(msg, ChequeMarshaled)

	print.Println100ms("-> sending msg")
	// send msg
	_, err = s.Write([]byte(msg))
	if err != nil {
		panic("stream write error")
	}

	print.Println100ms("\n> Intput target cmd: ")

	return err
}

// Cmd2Handler - command 2 handler, run on storage, receive PayCheque from user
func SendCheckHandler(s network.Stream) error {

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
	PayChequeMarshaled, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatalln("Error reading :", err)
		return err
	}

	// // parse data
	// var sigByte = in[:65]
	// var PayChequeMarshaled = in[65:]

	// unmarshal data
	PayCheque := &pb.PayCheque{}
	if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
		log.Fatalln("Failed to parse check:", err)
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

	// get sig from pay cheque
	sigByte := PayCheque.ChequeSig
	// calc hash from PayCheque
	hash := utils.CalcHash(PayCheque.Cheque.From, PayCheque.Cheque.NodeNonce, PayCheque.To, PayCheque.PayValue)

	// verify PayCheque signature: []byte []byte common.Address
	ok, verErr := sigapi.Verify(hash, sigByte, From)
	if verErr != nil {
		log.Fatal("verify fatal error occured")
		return verErr
	}

	if !ok {
		print.Println100ms("<signature of PayCheque verify failed>")
	} else {
		print.Println100ms("<signature of PayCheque verify success>")

		// wirte PayCheque into db
		// create/open db
		db, err := leveldb.OpenFile("./storage_paycheque.db", nil)
		if err != nil {
			log.Fatal("opfen db error")
		}

		// gen Cheque key: To + nonce
		bigNonce := big.NewInt(PayCheque.Cheque.NodeNonce)
		var ChequeKey []byte
		ChequeKey, err = utils.GenChequeKey(PayCheque.Cheque.To, bigNonce)
		if err != nil {
			log.Fatal("GenChequeKey error:", err)
		}

		if global.DEBUG {
			fmt.Println("in Cmd2Handler.")
			fmt.Printf("ChequeKey: %x\n", ChequeKey)
		}

		// use PayChequeKey as PayCheque id to store PayChequeMarshaled.
		//PayChequeMarshWithSig := utils.MergeSlice(sigByte, PayChequeMarshaled)
		err = db.Put(ChequeKey, PayChequeMarshaled, nil)
		if err != nil {
			print.Println100ms("db put data error")
			return err
		}

		db.Close()

	}

	utils.UpdatePayChequeIndex()
	utils.ListPayCheque()

	print.PrintMenu()
	print.Println100ms("\n> Intput target address and cmd: ")

	return nil
}
