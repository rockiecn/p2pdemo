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
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/p2pdemo/print"
	"github.com/rockiecn/p2pdemo/sigapi"
	"github.com/rockiecn/p2pdemo/utils"
)

// deploy cash
func ExeCmd1() {
	print.Println100ms("call deploy cash")
	callcash.CallDeploy()
}

// user buy Cheque from operator
func ExeCmd2() {
	// connect to peer, get stream
	s, err := hostops.HostInfo.NewStream(context.Background(), utils.Peerid, "/1")
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
	sigByte := in[:65]
	cashAddrByte := in[65:107]
	ChequeMarshaled := in[107:]
	fmt.Printf("cashAddr:%x\n", cashAddrByte)

	// unmarshal
	Cheque := &pb.Cheque{}
	if err := proto.Unmarshal(ChequeMarshaled, Cheque); err != nil {
		log.Fatalln("Failed to parse check:", err)
	}
	print.Printf100ms("--> Received Cheque:\n")

	print.PrintCheque(Cheque)

	// verify signature of Cheque, signed by operator

	// string to byte
	opAddrByte, err := hex.DecodeString(Cheque.OperatorAddress)
	if err != nil {
		panic("decode error")
	}
	// []byte to common.Address
	opAddress := common.BytesToAddress(opAddrByte)

	// calc hash
	hash := utils.CalcHash(Cheque.From, Cheque.NodeNonce, "", 0)
	print.Printf100ms("Cheque receive, hash: %x\n", hash)

	// verify Cheque signature
	ok, _ := sigapi.Verify(hash, sigByte, opAddress)
	if !ok {
		print.Println100ms("<signature of Cheque verify failed>")
		return
	} else {
		print.Println100ms("<signature of Cheque verify success>")

		// create/open db
		db, err := leveldb.OpenFile("./user_data.db", nil)
		if err != nil {
			log.Fatal("opfen db error")
		}

		// gen Cheque key: To + nonce
		if utils.DEBUG {
			fmt.Printf("storage address: %s\n", Cheque.To)
			fmt.Printf("nonce: %d\n", Cheque.NodeNonce)
		}

		bigNonce := big.NewInt(Cheque.NodeNonce)
		ChequeKey, err := utils.GenChequeKey(Cheque.To, bigNonce)
		if err != nil {
			log.Fatal("GenChequeKey error")
			return
		}

		if utils.DEBUG {
			fmt.Printf("ChequeKey: %x\n", ChequeKey)
		}

		// use ChequeHash as Cheque id to store ChequeMarshaled.
		// store Cheque_marshaled
		ChequeMarshWithSig := utils.MergeSlice(sigByte, ChequeMarshaled)
		err = db.Put(ChequeKey, ChequeMarshWithSig, nil)
		if err != nil {
			print.Println100ms("db put data error")
			return
		}
		db.Close()

		// show table
		utils.UpdateIndex()
		utils.ListUserDB()
	}

}

// user send PayCheque to storage
func ExeCmd3() {
	// create/open db
	db, err := leveldb.OpenFile("./user_data.db", nil)
	if err != nil {
		log.Fatal("opfen db error")
	}
	defer db.Close()

	// navigate Cheques
	iter := db.NewIterator(nil, nil)
loop:
	for iter.Next() {

		print.Printf100ms("Opening stream to peerID: %v\n", utils.Peerid)
		s, err := hostops.HostInfo.NewStream(context.Background(), utils.Peerid, "/2")
		if err != nil {
			log.Println(err)
			return
		}

		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		purMarshalWithSig := iter.Value()
		fmt.Printf(("Cheque key: %x\n"), key)

		ChequeSig := purMarshalWithSig[:65]
		ChequeMarshaled := purMarshalWithSig[65:]

		// unmarshal it to get Cheque itself
		Cheque := &pb.Cheque{}
		if err := proto.Unmarshal(ChequeMarshaled, Cheque); err != nil {
			log.Fatalln("Failed to parse check:", err)
		}

		// PayCheque should be created, signed and sent by user

		// create PayCheque
		PayCheque := &pb.PayCheque{}
		PayCheque.Cheque = Cheque
		PayCheque.ChequeSig = ChequeSig
		PayCheque.PayValue = 10 //wei
		PayCheque.To = "b213d01542d129806d664248a380db8b12059061"

		// calc hash from PayCheque
		hash := utils.CalcHash(PayCheque.Cheque.From, PayCheque.Cheque.NodeNonce, PayCheque.To, PayCheque.PayValue)
		print.Printf100ms("hash: %x\n", hash)
		// sign PayCheque by user' sk
		// user address: 1ab6a9f2b90004c1269563b5da391250ede3c114
		var userSkByte = []byte("b91c265cabae210642d66f9d59137eac2fab2674f4c1c88df3b8e9e6c1f74f9f")
		PayChequeSig, err := sigapi.Sign(hash, userSkByte)
		if err != nil {
			panic("sign error")
		}

		if utils.DEBUG {
			// for debug
			print.Printf100ms("DEBUG> From: %s\n", PayCheque.Cheque.From)
			print.Printf100ms("DEBUG> NodeNonce: %d\n", PayCheque.Cheque.NodeNonce)
			print.Printf100ms("DEBUG> To: %s\n", PayCheque.To)
			print.Printf100ms("DEBUG> PayValue: %d\n", PayCheque.PayValue)
			print.Printf100ms("DEBUG> signature: %x\n", PayChequeSig)
		}

		// serialize
		PayChequeMarshaled, err := proto.Marshal(PayCheque)
		if err != nil {
			log.Fatalln("Failed to encode PayCheque:", err)
		}

		// construct PayCheque message: signature(65 bytes) | marshaled cheqe
		PayChequeMsg := utils.MergeSlice(PayChequeSig, PayChequeMarshaled)

		// send PayCheque msg to storage
		print.Println100ms("--> user sending PayCheque to storage")
		_, err = s.Write(PayChequeMsg)
		if err != nil {
			log.Println(err)
			return
		}

		s.Close()

		for {
			fmt.Println("continue?(y/n)")
			var ctn string
			fmt.Scanf("%s", &ctn)
			switch ctn {
			case "y":
				continue loop
			case "n":
				break loop
			default:
				fmt.Println("error input, input y/n")
			}
		}

	}
	fmt.Println("end of user db iterate.")

	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// list user_db
func ExeCmd4() {
	utils.UpdateIndex()
	utils.ListUserDB()
}

// delete an entry of user db
func ExeCmd5() {

	utils.UpdateIndex()
	utils.ListUserDB()

	db, err := leveldb.OpenFile("./user_data.db", nil)
	if err != nil {
		log.Fatal("opfen db error")
	}
	fmt.Println("Input ID to delete:")
	var uID uint
	fmt.Scanf("%d", &uID)
	if utils.Index[uID] == "" {
		fmt.Println("ID not exist")
		return
	}

	var keyByte []byte
	keyByte, err = hex.DecodeString(utils.Index[uID])
	if err != nil {
		fmt.Println("decode string error: ", err)
	}
	fmt.Printf("delete ID %d success.\n", uID)

	err = db.Delete(keyByte, nil)
	if err != nil {
		fmt.Println("delete user db error: ", err)
	}
	fmt.Printf("delete ID %d success.\n", uID)

	db.Close()

	utils.UpdateIndex()
	utils.ListUserDB()
}

// call cash contract
func ExeCmd6() {
	print.Println100ms("call applyPayCheque in cash")

	// read PayCheque data from db
	// create/open db
	db, err := leveldb.OpenFile("./storage_data.db", nil)
	if err != nil {
		log.Fatal("opfen db error")
	}
	defer db.Close()

	// navigate Cheques
	iter := db.NewIterator(nil, nil)
	for iter.Next() {

		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		PayChequeMarshWithSig := iter.Value()
		fmt.Printf(("PayCheque key: %x\n"), key)

		PayChequeSig := PayChequeMarshWithSig[:65]
		PayChequeMarshaled := PayChequeMarshWithSig[65:]

		// unmarshal it to get PayCheque itself
		PayCheque := &pb.PayCheque{}
		if err := proto.Unmarshal(PayChequeMarshaled, PayCheque); err != nil {
			log.Fatalln("Failed to parse check:", err)
		}

		// string to common.Address
		From := common.HexToAddress(PayCheque.Cheque.From)

		// int to bigInt, nonce
		bigN := big.NewInt(PayCheque.Cheque.NodeNonce)

		// get storage address
		stAddrBytes, err := hex.DecodeString(PayCheque.To)
		if err != nil {
			panic("decode error")
		}
		// []byte to common.Address
		stAddress := common.BytesToAddress(stAddrBytes)

		// pay amount big
		bigPay := big.NewInt(PayCheque.PayValue)

		// // call contract
		// z18 := new(big.Int)
		// z18.SetString("1000000000000000000", 10)
		// weiPay := new(big.Int)
		// weiPay.Mul(bigPay, z18) // eth to wei

		// fmt.Println("bigPay: ", bigPay.String())
		// fmt.Println("z18: ", z18.String())
		// fmt.Println("weiPay: ", weiPay.String())

		//
		errCallApply := callcash.CallApplyPayCheque(From, bigN, stAddress, bigPay, PayChequeSig)
		if errCallApply != nil {
			log.Fatalln("callApplyPayCheque error:", err)
			log.Fatalln("storage address:", PayCheque.Cheque.To)
			log.Fatalln("nonce:", PayCheque.Cheque.NodeNonce)
		}

		fmt.Println("continue?(y/n)")
		var ctn string
		fmt.Scanf("%s", &ctn)
		if ctn != "y" {
			break
		}
	}
}

func ExeCmd7() {
	print.Println100ms("call retrieve")
	callstorage.CallRetrieve()
}
