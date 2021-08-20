package handler

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/p2pdemo/print"
	"github.com/rockiecn/p2pdemo/sigapi"
	"github.com/rockiecn/p2pdemo/utils"
	"github.com/syndtr/goleveldb/leveldb"

	// "github.com/rockiecn/sigtest/sigapi"
	// "github.com/rockiecn/sigtest/utils"

	"google.golang.org/protobuf/proto"
)

// Cmd1Handler - command 1 handler, run on user, receive purchase from operator
func Cmd1Handler(s network.Stream) error {

	print.Println100ms("--> Construct and send purchase...")

	// construct purchase
	Purchase := &pb.Purchase{}
	Purchase.PurchaseAmount = 100 // purchase 100
	//Purchase.NodeNonce = 1

	Purchase.OperatorAddress = "9e0153496067c20943724b79515472195a7aedaa"  // operator
	Purchase.UserAddress = "1ab6a9f2b90004c1269563b5da391250ede3c114"      // user
	Purchase.StorageAddress = "0xb213d01542d129806d664248a380db8b12059061" // storage
	Purchase.TokenAddress = "tokenaddress"

	// create/open db
	db, err := leveldb.OpenFile("./operator_data.db", nil)
	if err != nil {
		log.Fatal("open db error")
	}
	defer db.Close()

	var newNonce int64 = 0
	// storage -> nonce
	ret, err := db.Get([]byte(Purchase.StorageAddress), nil)
	if err != nil {
		if err.Error() == "leveldb: not found" { // no nonce at all
			db.Put([]byte(Purchase.StorageAddress), utils.Int64ToBytes(1), nil)
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
		err = db.Put([]byte(Purchase.StorageAddress), byteN, nil)
		if err != nil {
			fmt.Println("operator db put nonce error")
			return err
		}
	}

	//
	Purchase.NodeNonce = newNonce

	// serialize
	purchaseMarshaled, err := proto.Marshal(Purchase)
	if err != nil {
		log.Fatalln("Failed to encode cheque:", err)
	}

	// construct purchase message: sig(65 bytes) | data
	print.Println100ms("-> constructing msg")

	// calc hash
	hash := utils.CalcHash(Purchase.UserAddress, Purchase.NodeNonce, "", 0)
	print.Printf100ms("purchase send, hash: %x\n", hash)

	// sign purchase by operator
	var opSkByte = []byte("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
	sig, err := sigapi.Sign(hash, opSkByte)
	if err != nil {
		panic("sign error")
	}

	var msg = []byte{}
	msg = utils.MergeSlice(sig, purchaseMarshaled)

	print.Println100ms("-> sending msg")
	// send msg
	_, err = s.Write([]byte(msg))
	if err != nil {
		panic("stream write error")
	}

	print.Println100ms("\n> Intput target address and cmd: ")

	return err
}

// Cmd2Handler - command 2 handler, run on storage, receive cheque from user
func Cmd2Handler(s network.Stream) error {

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
	var chequeMarshaled = in[65:]

	// unmarshal data
	cheque := &pb.Cheque{}
	if err := proto.Unmarshal(chequeMarshaled, cheque); err != nil {
		log.Fatalln("Failed to parse check:", err)
		return err
	}

	//
	print.PrintCheque(cheque)

	//===== verify signature of cheque(signed by user)

	// get user address
	userAddrByte, err := hex.DecodeString(cheque.Purchase.UserAddress)
	if err != nil {
		panic("decode error")
	}
	// []byte to common.Address
	userAddress := common.BytesToAddress(userAddrByte)

	// calc hash from cheque
	hash := utils.CalcHash(cheque.Purchase.UserAddress, cheque.Purchase.NodeNonce, cheque.StorageAddress, cheque.PayAmount)

	// verify cheque signature: []byte []byte common.Address
	ok, verErr := sigapi.Verify(hash, sigByte, userAddress)
	if verErr != nil {
		log.Fatal("verify fatal error occured")
		return verErr
	}

	if !ok {
		print.Println100ms("<signature of cheque verify failed>")
	} else {
		print.Println100ms("<signature of cheque verify success>")

		// wirte cheque into db
		// create/open db
		db, err := leveldb.OpenFile("./storage_data.db", nil)
		if err != nil {
			log.Fatal("opfen db error")
		}
		defer db.Close()

		// gen purchase key: storageAddress + nonce
		bigNonce := big.NewInt(cheque.Purchase.NodeNonce)
		var chequeKey []byte
		chequeKey, err = utils.GenPurchaseKey(cheque.Purchase.StorageAddress, bigNonce)
		if err != nil {
			log.Fatal("GenPurchaseKey error:", err)
		}

		if utils.DEBUG {
			fmt.Println("in Cmd2Handler.")
			fmt.Printf("chequeKey: %x\n", chequeKey)
		}

		// use chequeKey as cheque id to store chequeMarshaled.
		chequeMarshWithSig := utils.MergeSlice(sigByte, chequeMarshaled)
		err = db.Put(chequeKey, chequeMarshWithSig, nil)
		if err != nil {
			print.Println100ms("db put data error")
			return err
		}

	}

	print.PrintMenu()
	print.Println100ms("\n> Intput target address and cmd: ")

	return nil
}
