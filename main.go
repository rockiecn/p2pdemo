package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/syndtr/goleveldb/leveldb"

	golog "github.com/ipfs/go-log/v2"
	ma "github.com/multiformats/go-multiaddr"

	"google.golang.org/protobuf/proto"

	"github.com/rockiecn/interact/callstorage"
	"github.com/rockiecn/p2pdemo/callcash"
	"github.com/rockiecn/p2pdemo/handler"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/p2pdemo/print"
	"github.com/rockiecn/p2pdemo/sigapi"
	"github.com/rockiecn/p2pdemo/utils"
)

// package level variable
var (
	Ctx    context.Context
	Ha     host.Host
	Cancel context.CancelFunc
)

// run listener
func init() {
	Ctx, Cancel = context.WithCancel(context.Background())
	defer Cancel()

	// Change to INFO for extra info
	golog.SetAllLoggers(golog.LevelInfo)

	// Parse options from the command line
	seedF := flag.Int64("seed", 0, "set random seed for id generation")
	flag.Parse()

	// Choose random ports between 10000-10100
	rand.Seed(time.Now().UnixNano())
	port := rand.Intn(500) + 10000

	// Make a host that listens on the given multiaddress
	var err error
	Ha, err = hostops.MakeBasicHost(port, true, *seedF)
	if err != nil {
		log.Fatal(err)
	}

	// run listener
	listenerDone := make(chan int)
	go runListener(Ctx, Ha, port, listenerDone)
	<-listenerDone //wait until runlistener complete
}

//
func main() {
	// run commandline
	for {
		// menu
		print.PrintMenu()

		fullAddr := hostops.GetHostAddress(Ha)
		print.Printf100ms("\n[ %s ]\n", fullAddr)

		var strCmd, strTarget string
		print.Println100ms("\n> Intput target address and cmd: ")
		fmt.Scanf("%s %s", &strTarget, &strCmd)
		if strTarget == "" || strCmd == "" {
			print.Printf100ms("invalid input, need target and cmd\n")
			continue
		}

		// execute command with cmd id
		exeCommand(Ctx, Ha, strTarget, strCmd)
	}
}

// set stream handler
func runListener(ctx context.Context, ha host.Host, listenPort int, listenerDone chan int) {

	// Set a stream handler on host A.
	ha.SetStreamHandler("/1", func(s network.Stream) {
		print.Println100ms("--> Received command 1")
		if err := handler.Cmd1Handler(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	ha.SetStreamHandler("/2", func(s network.Stream) {
		print.Println100ms("--> Received command 2")
		if err := handler.Cmd2Handler(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	listenerDone <- 0 // signal main to continue
}

// execute command
func exeCommand(ctx context.Context, ha host.Host, targetPeer string, cmd string) {

	// string to ma
	// /ip4/127.0.0.1/tcp/10043/p2p/QmZGUdbbgZ4VjKV9FPjc1Em6Hp9eRKfVV6TGWaGY7Fk4MR
	ipfsaddr, err := ma.NewMultiaddr(targetPeer)
	if err != nil {
		log.Println(err)
		return
	}
	// QmZGUdbbgZ4VjKV9FPjc1Em6Hp9eRKfVV6TGWaGY7Fk4MR
	pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	if err != nil {
		log.Println(err)
		return
	}
	// string to peer.ID
	// QmZGUdbbgZ4VjKV9FPjc1Em6Hp9eRKfVV6TGWaGY7Fk4MR
	peerid, err := peer.Decode(pid)
	if err != nil {
		log.Println(err)
		return
	}
	// /p2p/QmZGUdbbgZ4VjKV9FPjc1Em6Hp9eRKfVV6TGWaGY7Fk4MR
	targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", pid))
	// /ip4/127.0.0.1/tcp/10043
	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)
	// add to peerstore: peerID -> targetAddr
	ha.Peerstore().AddAddr(peerid, targetAddr, peerstore.PermanentAddrTTL)

	// open stream to target, with given protocol id
	switch cmd {
	// receive purchase from operator, signed by operator
	case "1":
		// connect to peer, get stream
		s, err := ha.NewStream(context.Background(), peerid, "/1")
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

		// verify signature
		ok, _ := sigapi.Verify(purchaseMarshaled, sigByte, opAddress)
		if ok {
			print.Println100ms("<signature of purchase verify success>")

			// create/open db
			db, err := leveldb.OpenFile("./data.db", nil)
			if err != nil {
				log.Fatal("opfen db error")
			}

			// store have_purchased
			err = db.Put([]byte("have_purchased"), []byte("true"), nil)
			if err != nil {
				print.Println100ms("db put data error")
			}
			// store purchase_marshaled
			err = db.Put([]byte("purchase_marshaled"), purchaseMarshaled, nil)
			if err != nil {
				print.Println100ms("db put data error")
			}
			// store purchase signature
			err = db.Put([]byte("purchase_sig"), sigByte, nil)
			if err != nil {
				print.Println100ms("db put data error")
			}

			db.Close()
		} else {
			print.Println100ms("<signature of purchase verify failed>")
			return
		}

	// send cheque to storage, signed by user
	case "2":
		print.Printf100ms("Opening stream to peerID: %v\n", peerid)
		s, err := ha.NewStream(context.Background(), peerid, "/2")
		if err != nil {
			log.Println(err)
			return
		}

		// create/open db
		db, err := leveldb.OpenFile("./data.db", nil)
		if err != nil {
			log.Fatal("opfen db error")
		}

		// read have_purchased flag from db
		havePurchased, err := db.Get([]byte("have_purchased"), nil)
		if err != nil {
			print.Println100ms("db get data error")
		}

		// check if purchase acquired
		hp := utils.Byte2Str(havePurchased)
		if hp != "true" {
			print.Println100ms("not require a purchase, run command 1 to get it.")
			return
		}

		// read purchase marshaled from db
		purchaseMarshaled, err := db.Get([]byte("purchase_marshaled"), nil)
		if err != nil {
			print.Println100ms("db get data error")
		}
		// unmarshal it to get purchase itself
		purchase := &pb.Purchase{}
		if err := proto.Unmarshal(purchaseMarshaled, purchase); err != nil {
			log.Fatalln("Failed to parse check:", err)
		}

		// read purchase siginature from db
		purchaseSig, err := db.Get([]byte("purchase_sig"), nil)
		if err != nil {
			print.Println100ms("db get data error")
		}

		// close
		db.Close()

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

		// for debug
		print.Printf100ms("UserAddress: %s\n", cheque.Purchase.UserAddress)
		print.Printf100ms("NodeNonce: %d\n", cheque.Purchase.NodeNonce)
		print.Printf100ms("StorageAddress: %s\n", cheque.StorageAddress)
		print.Printf100ms("PayAmount: %d\n", cheque.PayAmount)
		print.Printf100ms("signature: %x\n", chequeSig)

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

		// close stream for reader to continue
		s.Close()
	// call retrieve method of contract storage, for test
	case "3":
		print.Println100ms("call retrieve")
		callstorage.CallRetrieve()
	// use callcash package to deploy cash contract
	case "4":
		print.Println100ms("call deploy cash")
		callcash.CallDeploy()
		//callstorage.CallDeploy()

	// read cheque data from db, then call contract with it
	case "5":
		print.Println100ms("call applycheque in cash")

		// read cheque data from db
		// create/open db
		db, err := leveldb.OpenFile("./data.db", nil)
		if err != nil {
			log.Fatal("opfen db error")
		}
		// get cheque
		var chequeMarshaled []byte
		chequeMarshaled, err = db.Get([]byte("cheque"), nil)
		if err != nil {
			print.Println100ms("db get data error")
		}
		// get cheque signature
		var chequeSig []byte
		chequeSig, err = db.Get([]byte("cheque sig"), nil)
		if err != nil {
			print.Println100ms("db get data error")
		}
		//
		db.Close()

		// unmarshal data from bytes to struct cheque
		cheque := &pb.Cheque{}
		if err := proto.Unmarshal(chequeMarshaled, cheque); err != nil {
			log.Fatalln("Failed to parse check:", err)
			return
		}

		// // string to bytes
		// userAddrByte, err := hex.DecodeString(cheque.Purchase.UserAddress)
		// if err != nil {
		// 	panic("decode error")
		// }
		// // []byte to common.Address
		// userAddress := common.BytesToAddress(userAddrByte)

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
		callcash.CallApplyCheque(userAddress, bigN, stAddress, bigPay, chequeSig)

	}
}
