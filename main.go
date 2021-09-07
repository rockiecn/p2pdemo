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

	"github.com/rockiecn/p2pdemo/app"
	"github.com/rockiecn/p2pdemo/global"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/print"
)

//
func main() {
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

	//
	app := &app.App{}
	app.Init()

	// run listener
	listenerDone := make(chan int)
	go runListener(app, global.Ctx, hostops.HostInfo, port, listenerDone)
	<-listenerDone //wait until runlistener complete

	defer app.Op.CloseDB()
	defer app.User.CloseDB()
	defer app.Pro.CloseDB()

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

		var err error

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
			_, err = app.Op.DeployContract()
			if err != nil {
				fmt.Println("operator deploy contract error:", err)
			}
		// Get contract nonce
		case "gn":
			_, err = app.Op.GetContractNonce()
			if err != nil {
				fmt.Println("operator get contract error:", err)
			}
		case "re":
			//execmd.ResetNonceInOperatorDB()
			err = app.Op.ResetNonceInDB()
			if err != nil {
				fmt.Println("operator reset nonce error:", err)
			}
		case "sn":
			err = app.Op.ShowNonceInDB()
			if err != nil {
				fmt.Println("operator show nonce error:", err)
			}
		// user get cheque from operator
		case "g":
			_, err = app.User.BuyCheque()
			if err != nil {
				fmt.Println("user buy cheque error:", err)
			}
		// Send One PayCheque to storage
		case "s":
			err = app.User.SendOnePayChequeByID()
			if err != nil {
				fmt.Println("user send paycheque error:", err)
			}
		// list user's cheque db
		case "lu":
			err = app.User.ListPayCheque()
			if err != nil {
				fmt.Println("user list paycheque error:", err)
			}
		// delete a cheque of user
		case "du":
			app.User.DeleteChequeByID()
		// Inc And Send a Cheque to storage
		case "is":
			err = app.User.IncAndSendPayChequeByID()
			if err != nil {
				fmt.Println("user inc and send paycheque error:", err)
			}
		// list storage's cheque db
		case "ls":
			err = app.Pro.ListPayCheque()
			if err != nil {
				fmt.Println("provider list paycheque error:", err)
			}
		// delete a cheque of storage
		case "ds":
			err = app.Pro.DeleteChequeByID()
			if err != nil {
				fmt.Println("provider delete paycheque error:", err)
			}
		// call cash
		case "cc":
			//execmd.StorageCallCash()
			err = app.Pro.CallCashByID()
			if err != nil {
				fmt.Println("provider call cash error:", err)
			}

		// TestCall
		//case "t":
		//execmd.TestCall()
		case "cu":
			err = app.User.ClearDB()
			if err != nil {
				fmt.Println("user clear db error:", err)
			}
		case "cs":
			err = app.Pro.ClearDB()
			if err != nil {
				fmt.Println("provider clear db error:", err)
			}
		case "sh":
			err = app.User.ShowPayChequeByID()
			if err != nil {
				fmt.Println("user show paycheque error:", err)
			}
		default:
			print.Printf100ms("invalid input.\n")
		}
	}
}

// set stream handler
func runListener(app *app.App, ctx context.Context, ha host.Host, listenPort int, listenerDone chan int) {

	// executed handler when a stream opened.
	ha.SetStreamHandler("/1", func(s network.Stream) {
		print.Println100ms("--> Received command 1")
		if err := app.Op.BuyChequeHandler(s); err != nil {
			log.Println(err)
			s.Reset()
		}
		// send data over, close stream for receiver to get data.
		s.Close()
	})

	// handler for cmd 2
	ha.SetStreamHandler("/2", func(s network.Stream) {
		print.Println100ms("--> Received command 2")
		if err := app.Pro.SendCheckHandler(s); err != nil {
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
