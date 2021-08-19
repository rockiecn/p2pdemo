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
	"github.com/rockiecn/p2pdemo/handler"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/print"
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

		// parse peerid
		peerid, err := parsePeerID(hostops.HostInfo, strTarget)
		if err != nil {
			log.Println(err)
			return
		}
		// execute command with cmd id
		//exeCommand(Ctx, hostops.HostInfo, strTarget, strCmd)
		switch strCmd {
		case "1":
			// operator send purchase to user
			execmd.ExeCmd1(Ctx, hostops.HostInfo, peerid)
		case "2":
			// user send cheque to storage
			execmd.ExeCmd2(Ctx, hostops.HostInfo, peerid)
		case "3":
			// test
			execmd.ExeCmd3(Ctx, hostops.HostInfo, peerid)
		case "4":
			// deploy cash
			execmd.ExeCmd4(Ctx, hostops.HostInfo, peerid)
		case "5":
			// call cash contract
			execmd.ExeCmd5(Ctx, hostops.HostInfo, peerid)
		case "6":
			// list user_db
			execmd.ExeCmd6(Ctx, hostops.HostInfo, peerid)
		case "7":
			// list storage_db
		case "8":
			// delete an entry of user_db
		case "9":
			// delete an entry of storage_db
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
