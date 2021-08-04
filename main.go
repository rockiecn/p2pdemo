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

	golog "github.com/ipfs/go-log/v2"
	ma "github.com/multiformats/go-multiaddr"

	"google.golang.org/protobuf/proto"

	"github.com/rockiecn/p2pdemo/execmd"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/test-sig/sig/implement/sigapi"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// LibP2P code uses golog to log messages. They log with different
	// string IDs (i.e. "swarm"). We can control the verbosity level for
	// all loggers with:
	golog.SetAllLoggers(golog.LevelInfo) // Change to INFO for extra info

	// Parse options from the command line
	//listenF := flag.Int("l", 0, "wait for incoming connections")
	//targetF := flag.String("d", "", "target peer to dial")
	//insecureF := flag.Bool("insecure", false, "use an unencrypted connection")
	seedF := flag.Int64("seed", 0, "set random seed for id generation")
	//cmdF := flag.String("cmd", "", "cmd to be executed")
	flag.Parse()

	/*
		if *listenF == 0 {
			log.Fatal("Please provide a port to bind on with -l")
		}
	*/

	// Choose random ports between 10000-10100
	rand.Seed(time.Now().UnixNano())
	port := rand.Intn(500) + 10000

	// Make a host that listens on the given multiaddress
	//ha, err := hostops.MakeBasicHost(*listenF, *insecureF, *seedF)
	ha, err := hostops.MakeBasicHost(port, true, *seedF)
	if err != nil {
		log.Fatal(err)
	}

	// run listener
	// contact with goroutine
	lisener_done := make(chan int)
	go runListener(ctx, ha, port, true, lisener_done)
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
		runSender(ctx, ha, strTarget, strCmd)
	}
}

// set stream handler
func runListener(ctx context.Context, ha host.Host, listenPort int, insecure bool, listener_done chan int) {

	// Set a stream handler on host A. /echo/1.0.0 is
	// a user-defined protocol name.
	ha.SetStreamHandler("/cmd1", func(s network.Stream) {
		//fmt.Println("Listener received stream /cmd1")
		if err := execmd.ExeCmd1(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	ha.SetStreamHandler("/1", func(s network.Stream) {
		fmt.Println("---> Received command 1")
		if err := execmd.SendPurchase(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	ha.SetStreamHandler("/2", func(s network.Stream) {
		fmt.Println("---> Received command 2")
		if err := execmd.ReceiveCheque(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	_ = insecure
	/*
		if insecure {
			log.Printf("Now run \"./p2pdemo -l %d -d %s -insecure -cmd xx\" on a different terminal\n", listenPort+1, fullAddr)
		} else {
			log.Printf("Now run \"./p2pdemo -l %d -d %s -cmd xx\" on a different terminal\n", listenPort+1, fullAddr)
		}

		// Wait until canceled
		<-ctx.Done()
	*/
	listener_done <- 0 // signal main to continue
}

// open stream to target, with given protocol id
func runSender(ctx context.Context, ha host.Host, targetPeer string, cmd string) {

	//string to ma
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

	// execute command
	switch cmd {
	case "cmd1":
		fmt.Println("Opening stream...")
		s, err := ha.NewStream(context.Background(), peerid, "/cmd1")
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Send data...")
		_, err = s.Write([]byte("param for cmd1\n"))
		if err != nil {
			log.Println(err)
			return
		}

		// Read data.
		in, err := ioutil.ReadAll(s)
		if err != nil {
			log.Fatalln("Error reading :", err)
			return
		}

		// unmarshal data
		purchase := &pb.Purchase{}
		if err := proto.Unmarshal(in, purchase); err != nil {
			log.Fatalln("Failed to parse checj:", err)
		}
		fmt.Printf("Received struct data:\n")
		fmt.Printf("->purchase.MaxAmount: %d\n", purchase.MaxAmount)
		fmt.Printf("->purchase.NodeNonce: %d\n", purchase.NodeNonce)
		fmt.Printf("->purchase.From: %s\n", purchase.From)
		fmt.Printf("->purchase.To: %s\n", purchase.To)
		fmt.Printf("->purchase.TokenAddress: %s\n", purchase.TokenAddress)

		//sender_done <- 0 // signal main to continue

	// user require download cheque from operator
	case "1":
		// connect to peer, get stream
		s, err := ha.NewStream(context.Background(), peerid, "/1")
		if err != nil {
			log.Println(err)
			return
		}

		// Read from stream
		fmt.Println("---> user require purchase from operator")
		time.Sleep(100 * time.Millisecond)
		in, err := ioutil.ReadAll(s)
		if err != nil {
			log.Fatalln("Error reading :", err)
			return
		}

		// parse data
		var sigByte = in[:65]
		var purchase_marshaled = in[65:]

		fmt.Printf("sigByte len:%d\n", len(sigByte)) //65
		//fmt.Printf("skByte len:%d\n", len(skByte))                         //64
		fmt.Printf("purchase_marshaled len:%d\n", len(purchase_marshaled)) //50

		// unmarshal
		purchase := &pb.Purchase{}
		if err := proto.Unmarshal(purchase_marshaled, purchase); err != nil {
			log.Fatalln("Failed to parse check:", err)
		}
		fmt.Printf("---> Received purchase:\n")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("--------------------- purchase ---------------------")
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("->purchase.MaxAmount: %d\n", purchase.MaxAmount)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("->purchase.NodeNonce: %d\n", purchase.NodeNonce)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("->purchase.From: %s\n", purchase.From)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("->purchase.To: %s\n", purchase.To)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("->purchase.TokenAddress: %s\n", purchase.TokenAddress)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("----------------------------------------------------")
		time.Sleep(100 * time.Millisecond)

		// verify signature of purchase
		//var fromAddr = "9e0153496067c20943724b79515472195a7aedaa"
		// get from address
		fromAddrByte, err := hex.DecodeString(purchase.From)
		if err != nil {
			panic("decode error")
		}
		// []byte to common.Address
		fromAddress := common.BytesToAddress(fromAddrByte)

		// verify
		ok, _ := sigapi.Verify(purchase_marshaled, sigByte, fromAddress)
		if ok {
			fmt.Println("signature of purchase verify success")
		} else {
			fmt.Println("signature of purchase verify failed")
			return
		}
	// user send pay cheque to storage
	case "2":
		fmt.Println("Opening stream to peerID: ", peerid)
		time.Sleep(100 * time.Millisecond)
		s, err := ha.NewStream(context.Background(), peerid, "/2")
		if err != nil {
			log.Println(err)
			return
		}

		// construct download cheque
		purchase := &pb.Purchase{}
		purchase.MaxAmount = 1000
		purchase.NodeNonce = 1
		purchase.From = "user address"
		purchase.To = "storage address"
		purchase.TokenAddress = "tokenaddress"

		// construct pay cheque
		cheque := &pb.Cheque{}
		cheque.Purchase = purchase
		byteArr := []byte("download sign")
		cheque.DownloadSign = byteArr
		cheque.PayAmount = 100
		cheque.OperatorAddress = "operator address"

		// serialize
		out, err := proto.Marshal(cheque)
		if err != nil {
			log.Fatalln("Failed to encode cheque:", err)
		}
		//fmt.Printf("out: %v\n", out)

		// send pay cheque
		fmt.Println("---> user sending cheque to storage")
		time.Sleep(100 * time.Millisecond)
		_, err = s.Write(out)
		if err != nil {
			log.Println(err)
			return
		}
		s.Close()

		// // test
		// log.Println("sender saying hello")
		// _, err = s.Write([]byte("Hello, world!\n"))
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }
		// s.Close()

		//sender_done <- 0 // signal main to continue
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
