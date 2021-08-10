package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/syndtr/goleveldb/leveldb"

	golog "github.com/ipfs/go-log/v2"
	ma "github.com/multiformats/go-multiaddr"

	"google.golang.org/protobuf/proto"

	"github.com/rockiecn/interact/callcash"
	"github.com/rockiecn/interact/callstorage"
	"github.com/rockiecn/p2pdemo/handler"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/p2pdemo/print"
	"github.com/rockiecn/sigtest/sigapi"
	"github.com/rockiecn/sigtest/utils"
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
	listener_done := make(chan int)
	go runListener(Ctx, Ha, port, listener_done)
	<-listener_done //wait until runlistener complete
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

		// execute command
		exeCommand(Ctx, Ha, strTarget, strCmd)
	}
}

// set stream handler
func runListener(ctx context.Context, ha host.Host, listenPort int, listener_done chan int) {

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

	listener_done <- 0 // signal main to continue
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
	// operator send purchase to user, signed by operator
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
		var purchase_marshaled = in[65:]

		// unmarshal
		purchase := &pb.Purchase{}
		if err := proto.Unmarshal(purchase_marshaled, purchase); err != nil {
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
		ok, _ := sigapi.Verify(purchase_marshaled, sigByte, opAddress)
		if ok {
			// create/open db
			db, err := leveldb.OpenFile("./data.db", nil)
			if err != nil {
				log.Fatal("opfen db error")
			}

			print.Println100ms("<signature of purchase verify success>")
			// store have_purchased
			err = db.Put([]byte("have_purchased"), []byte("true"), nil)
			if err != nil {
				print.Println100ms("db put data error")
			}
			// store purchase_marshaled
			err = db.Put([]byte("purchase_marshaled"), purchase_marshaled, nil)
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

	// user send cheque to storage, signed by user
	case "2":
		print.Printf100ms("Opening stream to peerID: ", peerid)
		s, err := ha.NewStream(context.Background(), peerid, "/2")
		if err != nil {
			log.Println(err)
			return
		}

		// storage address:
		// b213d01542d129806d664248a380db8b12059061
		// storage sk:
		// aa03c94703e40a3f9e694a002dcb250182970917a7cd2346f2dfd337ada154f5

		// create/open db
		db, err := leveldb.OpenFile("./data.db", nil)
		if err != nil {
			log.Fatal("opfen db error")
		}

		// read have_purchased flag from db
		have_purchased, err := db.Get([]byte("have_purchased"), nil)
		if err != nil {
			print.Println100ms("db get data error")
		}

		// check if purchase acquired
		hp := utils.Byte2Str(have_purchased)
		if hp != "true" {
			print.Println100ms("not require a purchase, run command 1 to get it.")
			return
		}

		// read purchase marshaled from db
		purchase_marshaled, err := db.Get([]byte("purchase_marshaled"), nil)
		if err != nil {
			print.Println100ms("db get data error")
		}
		// unmarshal it to get purchase itself
		purchase := &pb.Purchase{}
		if err := proto.Unmarshal(purchase_marshaled, purchase); err != nil {
			log.Fatalln("Failed to parse check:", err)
		}

		// read purchase siginature from db
		purchase_sig, err := db.Get([]byte("purchase_sig"), nil)
		if err != nil {
			print.Println100ms("db get data error")
		}

		// close
		db.Close()

		// cheque should be created, signed and sent by user

		// create cheque
		cheque := &pb.Cheque{}
		cheque.Purchase = purchase
		cheque.PurchaseSig = purchase_sig
		cheque.PayAmount = 10
		cheque.StorageAddress = "b213d01542d129806d664248a380db8b12059061"

		// // sign
		// hash := crypto.Keccak256(utils.IntToBytes(purchase.NodeNonce))
		// print.Printf100ms("%v", hash)

		// serialize
		cheque_marshaled, err := proto.Marshal(cheque)
		if err != nil {
			log.Fatalln("Failed to encode cheque:", err)
		}

		// sign cheque
		var userSkByte = []byte("b91c265cabae210642d66f9d59137eac2fab2674f4c1c88df3b8e9e6c1f74f9f")
		cheque_sig, err := sigapi.Sign(cheque_marshaled, userSkByte)
		if err != nil {
			panic("sign error")
		}

		// construct cheque message: sig(65 bytes) | data
		cheque_msg := utils.MergeSlice(cheque_sig, cheque_marshaled)

		// send cheque
		print.Println100ms("--> user sending cheque to storage")
		_, err = s.Write(cheque_msg)
		if err != nil {
			log.Println(err)
			return
		}

		// close stream for reader to continue
		s.Close()

	case "3":
		print.Println100ms("call retrieve")
		callstorage.CallRetrieve()
	case "4":
		print.Println100ms("call deploy cash")
		callcash.CallDeploy()
	case "5":
		print.Println100ms("call applycheque in cash")
		callcash.CallApplyCheque()
	}
}
