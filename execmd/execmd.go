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

// operator send purchase to user
func ExeCmd1() {
	// connect to peer, get stream
	s, err := hostops.HostInfo.NewStream(context.Background(), utils.Peerid, "/1")
	if err != nil {
		log.Println(err)
		return
	}

	// Read from stream
	print.Println100ms("--> user receive purchase from operator")
	in, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatalln("Error reading :", err)
		return
	}

	// parse data
	var sigByte = in[:65]
	var purchaseMarshaled = in[65:]

	// unmarshal
	purchase := &pb.Purchase{}
	if err := proto.Unmarshal(purchaseMarshaled, purchase); err != nil {
		log.Fatalln("Failed to parse check:", err)
	}
	print.Printf100ms("--> Received purchase:\n")

	print.PrintPurchase(purchase)

	// verify signature of purchase, signed by operator

	// string to byte
	opAddrByte, err := hex.DecodeString(purchase.OperatorAddress)
	if err != nil {
		panic("decode error")
	}
	// []byte to common.Address
	opAddress := common.BytesToAddress(opAddrByte)

	// calc hash
	hash := utils.CalcHash(purchase.UserAddress, purchase.NodeNonce, "", 0)
	print.Printf100ms("purchase receive, hash: %x\n", hash)

	// verify purchase signature
	ok, _ := sigapi.Verify(hash, sigByte, opAddress)
	if !ok {
		print.Println100ms("<signature of purchase verify failed>")
		return
	} else {
		print.Println100ms("<signature of purchase verify success>")

		// create/open db
		db, err := leveldb.OpenFile("./user_data.db", nil)
		if err != nil {
			log.Fatal("opfen db error")
		}

		// // calc purchase hash
		// purchaseHash := utils.CalcPurchaseHash(purchaseMarshaled)
		// fmt.Printf("purchaseHash(purchase id): %x\n", purchaseHash)

		// gen purchase key: storageAddress + nonce
		if utils.DEBUG {
			fmt.Printf("in main\n")
			fmt.Printf("storage address: %s\n", purchase.StorageAddress)
			fmt.Printf("nonce: %d\n", purchase.NodeNonce)
		}

		bigNonce := big.NewInt(purchase.NodeNonce)
		purchaseKey, err := utils.GenPurchaseKey(purchase.StorageAddress, bigNonce)
		if err != nil {
			log.Fatal("GenPurchaseKey error")
			return
		}

		if utils.DEBUG {
			fmt.Println("in main")
			fmt.Printf("purchaseKey: %x\n", purchaseKey)
		}

		// use purchaseHash as purchase id to store purchaseMarshaled.
		// store purchase_marshaled
		purchaseMarshWithSig := utils.MergeSlice(sigByte, purchaseMarshaled)
		err = db.Put(purchaseKey, purchaseMarshWithSig, nil)
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

// user send cheque to storage
func ExeCmd2() {
	// create/open db
	db, err := leveldb.OpenFile("./user_data.db", nil)
	if err != nil {
		log.Fatal("opfen db error")
	}
	defer db.Close()

	// navigate purchases
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
		fmt.Printf(("purchase key: %x\n"), key)

		purchaseSig := purMarshalWithSig[:65]
		purchaseMarshaled := purMarshalWithSig[65:]

		// unmarshal it to get purchase itself
		purchase := &pb.Purchase{}
		if err := proto.Unmarshal(purchaseMarshaled, purchase); err != nil {
			log.Fatalln("Failed to parse check:", err)
		}

		// cheque should be created, signed and sent by user

		// create cheque
		cheque := &pb.Cheque{}
		cheque.Purchase = purchase
		cheque.PurchaseSig = purchaseSig
		cheque.PayAmount = 10 //wei
		cheque.StorageAddress = "b213d01542d129806d664248a380db8b12059061"

		// calc hash from cheque
		hash := utils.CalcHash(cheque.Purchase.UserAddress, cheque.Purchase.NodeNonce, cheque.StorageAddress, cheque.PayAmount)
		print.Printf100ms("hash: %x\n", hash)
		// sign cheque by user' sk
		// user address: 1ab6a9f2b90004c1269563b5da391250ede3c114
		var userSkByte = []byte("b91c265cabae210642d66f9d59137eac2fab2674f4c1c88df3b8e9e6c1f74f9f")
		chequeSig, err := sigapi.Sign(hash, userSkByte)
		if err != nil {
			panic("sign error")
		}

		if utils.DEBUG {
			// for debug
			print.Printf100ms("DEBUG> UserAddress: %s\n", cheque.Purchase.UserAddress)
			print.Printf100ms("DEBUG> NodeNonce: %d\n", cheque.Purchase.NodeNonce)
			print.Printf100ms("DEBUG> StorageAddress: %s\n", cheque.StorageAddress)
			print.Printf100ms("DEBUG> PayAmount: %d\n", cheque.PayAmount)
			print.Printf100ms("DEBUG> signature: %x\n", chequeSig)
		}

		// serialize
		chequeMarshaled, err := proto.Marshal(cheque)
		if err != nil {
			log.Fatalln("Failed to encode cheque:", err)
		}

		// construct cheque message: signature(65 bytes) | marshaled cheqe
		chequeMsg := utils.MergeSlice(chequeSig, chequeMarshaled)

		// send cheque msg to storage
		print.Println100ms("--> user sending cheque to storage")
		_, err = s.Write(chequeMsg)
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

func ExeCmd3() {
	print.Println100ms("call retrieve")
	callstorage.CallRetrieve()
}

// deploy cash
func ExeCmd4() {
	print.Println100ms("call deploy cash")
	callcash.CallDeploy()
}

// call cash contract
func ExeCmd5() {
	print.Println100ms("call applycheque in cash")

	// read cheque data from db
	// create/open db
	db, err := leveldb.OpenFile("./storage_data.db", nil)
	if err != nil {
		log.Fatal("opfen db error")
	}
	defer db.Close()

	// navigate purchases
	iter := db.NewIterator(nil, nil)
	for iter.Next() {

		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		chequeMarshWithSig := iter.Value()
		fmt.Printf(("cheque key: %x\n"), key)

		chequeSig := chequeMarshWithSig[:65]
		chequeMarshaled := chequeMarshWithSig[65:]

		// unmarshal it to get cheque itself
		cheque := &pb.Cheque{}
		if err := proto.Unmarshal(chequeMarshaled, cheque); err != nil {
			log.Fatalln("Failed to parse check:", err)
		}

		// string to common.Address
		userAddress := common.HexToAddress(cheque.Purchase.UserAddress)

		// int to bigInt, nonce
		bigN := big.NewInt(cheque.Purchase.NodeNonce)

		// get storage address
		stAddrBytes, err := hex.DecodeString(cheque.StorageAddress)
		if err != nil {
			panic("decode error")
		}
		// []byte to common.Address
		stAddress := common.BytesToAddress(stAddrBytes)

		// pay amount big
		bigPay := big.NewInt(cheque.PayAmount)

		// // call contract
		// z18 := new(big.Int)
		// z18.SetString("1000000000000000000", 10)
		// weiPay := new(big.Int)
		// weiPay.Mul(bigPay, z18) // eth to wei

		// fmt.Println("bigPay: ", bigPay.String())
		// fmt.Println("z18: ", z18.String())
		// fmt.Println("weiPay: ", weiPay.String())

		//
		errCallApply := callcash.CallApplyCheque(userAddress, bigN, stAddress, bigPay, chequeSig)
		if errCallApply != nil {
			log.Fatalln("callApplyCheque error:", err)
			log.Fatalln("storage address:", cheque.Purchase.StorageAddress)
			log.Fatalln("nonce:", cheque.Purchase.NodeNonce)
		}

		fmt.Println("continue?(y/n)")
		var ctn string
		fmt.Scanf("%s", &ctn)
		if ctn != "y" {
			break
		}
	}
}

// list user_db
func ExeCmd6() {
	utils.UpdateIndex()
	utils.ListUserDB()
}

// delete an entry of user db
func ExeCmd7() {

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
