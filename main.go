package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"

	golog "github.com/ipfs/go-log/v2"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/golang/protobuf/proto"

	"github.com/rockiecn/p2pdemo/execmd"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
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

	/*
		if *targetF == "" {
			log.Println("running listener.")
			go runListener(ctx, ha, *listenF, *insecureF)
		} else {
			runSender(ctx, ha, *targetF, *cmdF)
		}
	*/

	// run listener
	// contact with goroutine
	lisener_done := make(chan int)
	sender_done := make(chan int)
	//go runListener(ctx, ha, *listenF, *insecureF, lisener_done)
	go runListener(ctx, ha, port, true, lisener_done)
	<-lisener_done //wait until runlistener complete

	// run commandline
	for {
		// menu
		printMenu()

		fullAddr := hostops.GetHostAddress(ha)
		fmt.Printf("\n[ %s ]\n", fullAddr)

		fmt.Printf("\n> ")
		var strCmd string
		var strTarget string
		fmt.Printf("Input target address and cmd: \n")
		fmt.Scanf("%s %s", &strTarget, &strCmd)
		if strTarget == "" || strCmd == "" {
			fmt.Printf("invalid input, need target and cmd\n")
			continue
		}
		go runSender(ctx, ha, strTarget, strCmd, sender_done)
		<-sender_done // wait util sender complete
		//runSender(ctx, ha, *targetF, *cmdF)
	}
}

// set stream handler
func runListener(ctx context.Context, ha host.Host, listenPort int, insecure bool, listener_done chan int) {

	// Set a stream handler on host A. /echo/1.0.0 is
	// a user-defined protocol name.
	ha.SetStreamHandler("/cmd1", func(s network.Stream) {
		fmt.Println("Listener received stream /cmd1")
		if err := execmd.ExeCmd1(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	ha.SetStreamHandler("/1", func(s network.Stream) {
		//log.Println("Listener received cmd2")
		if err := execmd.ExeDownloadCheque(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	ha.SetStreamHandler("/2", func(s network.Stream) {
		log.Println("Listener received 2")
		if err := execmd.ExePayCheque(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	//fmt.Println("listening for connections")

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
func runSender(ctx context.Context, ha host.Host, targetPeer string, cmd string, sender_done chan int) {

	// The following code extracts target's the peer ID from the
	// given multiaddress
	ipfsaddr, err := ma.NewMultiaddr(targetPeer)
	if err != nil {
		log.Println(err)
		return
	}

	pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	if err != nil {
		log.Println(err)
		return
	}

	peerid, err := peer.Decode(pid)
	if err != nil {
		log.Println(err)
		return
	}

	// Decapsulate the /ipfs/<peerID> part from the target
	// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
	targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", pid))
	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)
	// We have a peer ID and a targetAddr so we add it to the peerstore
	// so LibP2P knows how to contact it
	ha.Peerstore().AddAddr(peerid, targetAddr, peerstore.PermanentAddrTTL)

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
		download_cheque := &pb.DownloadCheque{}
		if err := proto.Unmarshal(in, download_cheque); err != nil {
			log.Fatalln("Failed to parse checj:", err)
		}
		fmt.Printf("Received struct data:\n")
		fmt.Printf("->download_cheque.MaxAmount: %d\n", download_cheque.MaxAmount)
		fmt.Printf("->download_cheque.NodeNonce: %d\n", download_cheque.NodeNonce)
		fmt.Printf("->download_cheque.From: %s\n", download_cheque.From)
		fmt.Printf("->download_cheque.To: %s\n", download_cheque.To)
		fmt.Printf("->download_cheque.TokenAddress: %s\n", download_cheque.TokenAddress)

		sender_done <- 0 // signal main to continue

	// user require download cheque from operator
	case "1":
		s, err := ha.NewStream(context.Background(), peerid, "/1")
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
		download_cheque := &pb.DownloadCheque{}
		if err := proto.Unmarshal(in, download_cheque); err != nil {
			log.Fatalln("Failed to parse checj:", err)
		}
		fmt.Printf("Received download cheque:\n")
		fmt.Printf("->download_cheque.MaxAmount: %d\n", download_cheque.MaxAmount)
		fmt.Printf("->download_cheque.NodeNonce: %d\n", download_cheque.NodeNonce)
		fmt.Printf("->download_cheque.From: %s\n", download_cheque.From)
		fmt.Printf("->download_cheque.To: %s\n", download_cheque.To)
		fmt.Printf("->download_cheque.TokenAddress: %s\n", download_cheque.TokenAddress)

		sender_done <- 0 // signal main to continue

	// send pay cheque to storage
	case "2":
		fmt.Println("cmd 2")
		fmt.Println("Opening stream to peerID: ", peerid)
		s, err := ha.NewStream(context.Background(), peerid, "/2")
		if err != nil {
			log.Println(err)
			return
		}

		/*
			// construct download cheque
			dc := &pb.DownloadCheque{}
			dc.MaxAmount = 1000
			dc.NodeNonce = 1
			dc.From = "user address"
			dc.To = "storage address"
			dc.TokenAddress = "tokenaddress"

			// construct pay cheque
			pc := &pb.PayCheque{}
			pc.Dc = dc
			byteArr := []byte("download sign")
			pc.DownloadSign = byteArr
			pc.PayAmount = 100
			pc.OperatorAddress = "operator address"

			// serialize
			out, err := proto.Marshal(pc)
			if err != nil {
				log.Fatalln("Failed to encode cheque:", err)
			}
			out = append(out, '\n')
			fmt.Printf("out: %v\n", out)

			// send pay cheque
			fmt.Println("sending out to stream")
			_, err = s.Write(out)
			if err != nil {
				log.Println(err)
				return
			}
		*/

		log.Println("sender saying hello")
		out := []byte{'a', 'b', 'c'}
		_, err = s.Write([]byte(out))
		//_, err = s.Write([]byte("Hello, world!\n"))
		if err != nil {
			log.Println(err)
			return
		}

		sender_done <- 0 // signal main to continue
	}
}

func printMenu() {
	fmt.Println("\n----------------Menu-----------------\n")
	fmt.Println("1. require download cheque from operator\n")
	fmt.Println("2. send pay cheque to storage\n")
	fmt.Println("-------------------------------------\n")
}
