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
var Ctx context.Context

// run listener
func init() {
	var Cancel context.CancelFunc
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
	hostops.HostInfo, err = hostops.MakeBasicHost(port, true, *seedF)
	if err != nil {
		log.Fatal(err)
	}

	// run listener
	listenerDone := make(chan int)
	go runListener(Ctx, hostops.HostInfo, port, listenerDone)
	<-listenerDone //wait until runlistener complete
}

//
func main() {
	// run commandline
	for {
		// menu
		print.PrintMenu()

		var strCmd, strTarget string
		print.Println100ms("\n> Intput target address and cmd: ")
		fmt.Scanf("%s %s", &strTarget, &strCmd)
		if strTarget == "" || strCmd == "" {
			print.Printf100ms("invalid input, need target and cmd\n")
			continue
		}

		// execute command with cmd id
		exeCommand(Ctx, hostops.HostInfo, strTarget, strCmd)
	}
}

// exeCommand - execute specified command according to the cmd param.
// 1.receive purchase from operator, signed by operator
// 2.send cheque to storage, signed by user
// 3.call retrieve method of contract storage, for test
// 4.use callcash package to deploy cash contract
// 5.read cheque data from db, then call contract with it
func exeCommand(ctx context.Context, ha host.Host, targetPeer string, cmd string) {
	// get peerid
	peerid, err := parsePeerID(ha, targetPeer)
	if err != nil {
		log.Println(err)
		return
	}
	switch cmd {
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
			defer db.Close()

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
		}
	case "2":
		// create/open db
		db, err := leveldb.OpenFile("./user_data.db", nil)
		if err != nil {
			log.Fatal("opfen db error")
		}
		defer db.Close()

		// navigate purchases
		iter := db.NewIterator(nil, nil)
		for iter.Next() {

			print.Printf100ms("Opening stream to peerID: %v\n", peerid)
			s, err := ha.NewStream(context.Background(), peerid, "/2")
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

			fmt.Println("continue?(y/n)")
			var ctn string
			fmt.Scanf("%s", &ctn)
			if ctn != "y" {
				break
			}
		}
		fmt.Println("end of user db iterate.")

		iter.Release()
		err = iter.Error()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "3":
		print.Println100ms("call retrieve")
		callstorage.CallRetrieve()
	case "4":
		print.Println100ms("call deploy cash")
		callcash.CallDeploy()
	case "5":
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
}

// set stream handler
func runListener(ctx context.Context, ha host.Host, listenPort int, listenerDone chan int) {

	// executed handler when a stream opened.
	ha.SetStreamHandler("/1", func(s network.Stream) {
		print.Println100ms("--> Received command 1")
		if err := handler.Cmd1Handler(s); err != nil {
			log.Println(err)
			s.Reset()
		}
		// send data over, close stream for receiver to get data.
		s.Close()

	})
	// handler for cmd 2
	ha.SetStreamHandler("/2", func(s network.Stream) {
		print.Println100ms("--> Received command 2")
		if err := handler.Cmd2Handler(s); err != nil {
			log.Println(err)
			s.Reset()
		}

		s.Close()
	})

	listenerDone <- 0 // signal main to continue
}

// parse peerid from targetPeer, and add it to peerstore
func parsePeerID(ha host.Host, targetPeer string) (peer.ID, error) {
	// string to ma
	// /ip4/127.0.0.1/tcp/10043/p2p/QmZGUdbbgZ4VjKV9FPjc1Em6Hp9eRKfVV6TGWaGY7Fk4MR
	ipfsaddr, err := ma.NewMultiaddr(targetPeer)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// QmZGUdbbgZ4VjKV9FPjc1Em6Hp9eRKfVV6TGWaGY7Fk4MR
	pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// string to peer.ID
	// QmZGUdbbgZ4VjKV9FPjc1Em6Hp9eRKfVV6TGWaGY7Fk4MR
	peerid, err := peer.Decode(pid)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// /p2p/QmZGUdbbgZ4VjKV9FPjc1Em6Hp9eRKfVV6TGWaGY7Fk4MR
	targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", pid))
	// /ip4/127.0.0.1/tcp/10043
	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)
	// add to peerstore: peerID -> targetAddr
	ha.Peerstore().AddAddr(peerid, targetAddr, peerstore.PermanentAddrTTL)

	return peerid, nil
}
