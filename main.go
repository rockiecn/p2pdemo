package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"

	golog "github.com/ipfs/go-log/v2"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/rockiecn/p2pdemo/execmd"
	"github.com/rockiecn/p2pdemo/global"
	"github.com/rockiecn/p2pdemo/handler"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/print"
)

// run listener
func init() {
	var Cancel context.CancelFunc
	global.Ctx, Cancel = context.WithCancel(context.Background())
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
	go runListener(global.Ctx, hostops.HostInfo, port, listenerDone)
	<-listenerDone //wait until runlistener complete
}

//
func main() {

	// local peer
	//fullAddr := hostops.GetHostAddress(hostops.HostInfo)
	//fmt.Printf("\nLocal peer address: \n[ %s ]\n", fullAddr)

	// menu
	print.PrintMenu()

	// run commandline
	for {

		var strCmd string
		print.Println100ms("\n> Intput command: ")
		fmt.Scanf("%s", &strCmd)
		if strCmd == "" {
			continue
		}

		// execute command with cmd id
		//exeCommand(Ctx, hostops.HostInfo, strTarget, strCmd)
		switch strCmd {
		// show menu
		case "m":
			print.PrintMenu()
		// recoredRemote
		case "r":
			recoredRemote()
		// DeployCash
		case "d":
			execmd.DeployCash()
		// Get contract nonce
		case "gn":
			execmd.GetContractNonce()
		case "re":
			execmd.ResetNonceInOperatorDB()
		case "sn":
			execmd.ShowNonceInOperatorDB()
		// user get cheque from operator
		case "g":
			execmd.GetCheque()
		// Send One PayCheque to storage
		case "s":
			execmd.SendOnePayChequeByID()
		// list user's cheque db
		case "lu":
			execmd.ListDB(true)
		// delete a cheque of user
		case "du":
			execmd.DeleteChequeByID(true)
		// Inc And Send a Cheque to storage
		case "is":
			execmd.IncAndSendCheque()
		// list storage's cheque db
		case "ls":
			execmd.ListDB(false)
		// delete a cheque of storage
		case "ds":
			execmd.DeleteChequeByID(false)
		// call cash
		case "cc":
			execmd.StorageCallCash()
		// TestCall
		case "t":
			execmd.TestCall()
		case "cu":
			execmd.ClearDB(true)
		case "cs":
			execmd.ClearDB(false)
		case "sh":
			execmd.ShowPayChequeByID()
		default:
			print.Printf100ms("invalid input.\n")
		}
	}
}

// set stream handler
func runListener(ctx context.Context, ha host.Host, listenPort int, listenerDone chan int) {

	// executed handler when a stream opened.
	ha.SetStreamHandler("/1", func(s network.Stream) {
		print.Println100ms("--> Received command 1")
		if err := handler.BuyCheckHandler(s); err != nil {
			log.Println(err)
			s.Reset()
		}
		// send data over, close stream for receiver to get data.
		s.Close()
	})

	// handler for cmd 2
	ha.SetStreamHandler("/2", func(s network.Stream) {
		print.Println100ms("--> Received command 2")
		if err := handler.SendCheckHandler(s); err != nil {
			log.Println(err)
			s.Reset()
		}

		s.Close()
	})

	listenerDone <- 0 // signal main to continue
}

// connect to a peer
func recoredRemote() {
	var strTarget string

	print.Println100ms("-> Intput peer address: ")
	fmt.Scanf("%s", &strTarget)
	if strTarget == "" {
		print.Printf100ms("invalid input, need target address.\n")
		return
	}

	// parse peerid
	var err error
	global.Peerid, err = parsePeerID(hostops.HostInfo, strTarget)
	if err != nil {
		log.Println(err)
		return
	}

	global.RemoteExist = true
	print.Println100ms("-> Recorded.")
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
