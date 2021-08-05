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
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/syndtr/goleveldb/leveldb"

	golog "github.com/ipfs/go-log/v2"
	ma "github.com/multiformats/go-multiaddr"

	"google.golang.org/protobuf/proto"

	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/test-sig/sig/implement/sigapi"
	"github.com/rockiecn/test-sig/sig/implement/utils"

	"github.com/rockiecn/p2pdemo/handler"
)

//
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Change to INFO for extra info
	golog.SetAllLoggers(golog.LevelInfo)

	// Parse options from the command line
	seedF := flag.Int64("seed", 0, "set random seed for id generation")
	flag.Parse()

	// Choose random ports between 10000-10100
	rand.Seed(time.Now().UnixNano())
	port := rand.Intn(500) + 10000

	// Make a host that listens on the given multiaddress
	ha, err := hostops.MakeBasicHost(port, true, *seedF)
	if err != nil {
		log.Fatal(err)
	}

	// run listener
	lisener_done := make(chan int)
	go runListener(ctx, ha, port, lisener_done)
	<-lisener_done //wait until runlistener complete

	// run commandline
	for {
		// menu
		printMenu()

		fullAddr := hostops.GetHostAddress(ha)
		fmt.Printf("\n[ %s ]\n", fullAddr)

		var strCmd string
		var strTarget string
		fmt.Println("\n> Intput target address and cmd: ")
		fmt.Scanf("%s %s", &strTarget, &strCmd)
		if strTarget == "" || strCmd == "" {
			fmt.Printf("invalid input, need target and cmd\n")
			continue
		}

		// call
		exeCommand(ctx, ha, strTarget, strCmd)
	}
}

// set stream handler
func runListener(ctx context.Context, ha host.Host, listenPort int, listener_done chan int) {

	// Set a stream handler on host A.
	ha.SetStreamHandler("/1", func(s network.Stream) {
		fmt.Println("--> Received command 1")
		if err := handler.Cmd1Handler(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	ha.SetStreamHandler("/2", func(s network.Stream) {
		fmt.Println("--> Received command 2")
		if err := handler.Cmd2Handler(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	listener_done <- 0 // signal main to continue
}

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
	// operator send purchase to user
	case "1":
		// connect to peer, get stream
		s, err := ha.NewStream(context.Background(), peerid, "/1")
		if err != nil {
			log.Println(err)
			return
		}

		// Read from stream
		fmt.Println("--> user require purchase from operator")
		time.Sleep(100 * time.Millisecond)
		in, err := ioutil.ReadAll(s)
		if err != nil {
			log.Fatalln("Error reading :", err)
			return
		}

		// parse data
		var sigByte = in[:65]
		var purchase_marshaled = in[65:]

		// fmt.Printf("sigByte len:%d\n", len(sigByte))                       //65
		// fmt.Printf("purchase_marshaled len:%d\n", len(purchase_marshaled)) //

		// unmarshal
		purchase := &pb.Purchase{}
		if err := proto.Unmarshal(purchase_marshaled, purchase); err != nil {
			log.Fatalln("Failed to parse check:", err)
		}
		fmt.Printf("--> Received purchase:\n")
		time.Sleep(100 * time.Millisecond)

		handler.PrintPurchase(purchase)

		// verify signature of purchase

		// get operator address
		opAddrByte, err := hex.DecodeString(purchase.OperatorAddress)
		if err != nil {
			panic("decode error")
		}
		// []byte to common.Address
		opAddress := common.BytesToAddress(opAddrByte)

		// verify
		ok, _ := sigapi.Verify(purchase_marshaled, sigByte, opAddress)
		if ok {
			// create/open db
			db, err := leveldb.OpenFile("./data.db", nil)
			if err != nil {
				log.Fatal("opfen db error")
			}

			fmt.Println("<signature of purchase verify success>")
			// store have_purchased
			err = db.Put([]byte("have_purchased"), []byte("true"), nil)
			if err != nil {
				fmt.Println("db put data error")
			}
			// store purchase_marshaled
			err = db.Put([]byte("purchase_marshaled"), purchase_marshaled, nil)
			if err != nil {
				fmt.Println("db put data error")
			}
			// store purchase signature
			err = db.Put([]byte("purchase_sig"), sigByte, nil)
			if err != nil {
				fmt.Println("db put data error")
			}

			db.Close()
		} else {
			fmt.Println("<signature of purchase verify failed>")
			return
		}

	// user send cheque to storage
	case "2":
		fmt.Println("Opening stream to peerID: ", peerid)
		time.Sleep(100 * time.Millisecond)
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

		// get have_purchased
		have_purchased, err := db.Get([]byte("have_purchased"), nil)
		if err != nil {
			fmt.Println("db get data error")
		}

		// check if purchase acquired
		hp := utils.Byte2Str(have_purchased)
		if hp != "true" {
			fmt.Println("not require a purchase, run command 1 to get it.")
			return
		}

		// get purchase marshaled
		purchase_marshaled, err := db.Get([]byte("purchase_marshaled"), nil)
		if err != nil {
			fmt.Println("db get data error")
		}
		// unmarshal it
		purchase := &pb.Purchase{}
		if err := proto.Unmarshal(purchase_marshaled, purchase); err != nil {
			log.Fatalln("Failed to parse check:", err)
		}

		// get purchase siginature
		purchase_sig, err := db.Get([]byte("purchase_sig"), nil)
		if err != nil {
			fmt.Println("db get data error")
		}

		// close
		db.Close()

		// get cheque data
		cheque := &pb.Cheque{}
		cheque.Purchase = purchase
		cheque.PurchaseSig = purchase_sig
		cheque.PayAmount = 10
		cheque.StorageAddress = "b213d01542d129806d664248a380db8b12059061"

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
		var cheque_msg = []byte{}
		cheque_msg = utils.MergeSlice(cheque_sig, cheque_marshaled)

		// send cheque
		fmt.Println("--> user sending cheque to storage")
		time.Sleep(100 * time.Millisecond)
		_, err = s.Write(cheque_msg)
		if err != nil {
			log.Println(err)
			return
		}

		// close stream for reader to continue
		s.Close()

	}
}

// print command menu
func printMenu() {
	fmt.Println()
	fmt.Println("======================= Menu =======================")
	fmt.Println("cmd 1: require download cheque from operator")
	fmt.Println("cmd 2: send pay cheque to storage")
	fmt.Println("====================================================")
}
