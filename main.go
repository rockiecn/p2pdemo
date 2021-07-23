package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"

	golog "github.com/ipfs/go-log/v2"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/golang/protobuf/proto"

	pb "github.com/rockiecn/p2pdemo/check_go"
	//"github.com/rockiecn/p2pdemo/hostops"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// LibP2P code uses golog to log messages. They log with different
	// string IDs (i.e. "swarm"). We can control the verbosity level for
	// all loggers with:
	golog.SetAllLoggers(golog.LevelInfo) // Change to INFO for extra info

	// Parse options from the command line
	listenF := flag.Int("l", 0, "wait for incoming connections")
	//targetF := flag.String("d", "", "target peer to dial")
	insecureF := flag.Bool("insecure", false, "use an unencrypted connection")
	seedF := flag.Int64("seed", 0, "set random seed for id generation")
	//cmdF := flag.String("cmd", "", "cmd to be executed")
	flag.Parse()

	if *listenF == 0 {
		log.Fatal("Please provide a port to bind on with -l")
	}

	// Make a host that listens on the given multiaddress
	ha, err := hostops.makeBasicHost(*listenF, *insecureF, *seedF)
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
	go runListener(ctx, ha, *listenF, *insecureF, lisener_done)
	<-lisener_done //wait until runlistener complete

	// commandline
	for true {
		fullAddr := hostops.getHostAddress(ha)
		fmt.Printf("\n[ %s ]\n", fullAddr)

		fmt.Printf("\n> ")
		var strCmd string
		var strTarget string
		fmt.Printf("input target and cmd: \n")
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

func runListener(ctx context.Context, ha host.Host, listenPort int, insecure bool, listener_done chan int) {
	//fullAddr := getHostAddress(ha)
	//log.Printf("I am %s\n", fullAddr)

	// Set a stream handler on host A. /echo/1.0.0 is
	// a user-defined protocol name.
	ha.SetStreamHandler("/cmd1", func(s network.Stream) {
		fmt.Println("listener received stream /cmd1")
		if err := exeCmd1(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})
	ha.SetStreamHandler("/cmd2", func(s network.Stream) {
		log.Println("listener received cmd2")
		if err := exeCmd2(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	fmt.Println("listening for connections")

	_ = insecure
	/*
		if insecure {
			log.Printf("Now run \"./p2pdemo -l %d -d %s -insecure -cmd xx\" on a different terminal\n", listenPort+1, fullAddr)
		} else {
			log.Printf("Now run \"./p2pdemo -l %d -d %s -cmd xx\" on a different terminal\n", listenPort+1, fullAddr)
		}
	*/
	listener_done <- 0 // signal main to continue

	// Wait until canceled
	//<-ctx.Done()
}

func runSender(ctx context.Context, ha host.Host, targetPeer string, cmd string, sender_done chan int) {
	//fullAddr := getHostAddress(ha)
	//fmt.Printf("I am %s\n", fullAddr)

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

	fmt.Println("sender opening stream")

	//fmt.Printf("cmd: %s\n", cmd)
	switch cmd {
	case "cmd1":
		s, err := ha.NewStream(context.Background(), peerid, "/cmd1")
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("send cmd1 param")
		_, err = s.Write([]byte("param for cmd1\n"))
		if err != nil {
			log.Println(err)
			return
		}

		/*
			//get result
			out, err := ioutil.ReadAll(s)
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("cmd1 result: %q\n", out)
		*/

		// Read data.
		in, err := ioutil.ReadAll(s)
		if err != nil {
			log.Fatalln("Error reading :", err)
		}
		check := &pb.DownloadCheck{}
		if err := proto.Unmarshal(in, check); err != nil {
			log.Fatalln("Failed to parse checj:", err)
		}
		fmt.Printf("print received struct data:\n")
		fmt.Printf("->check.MaxAmount: %d\n", check.MaxAmount)
		fmt.Printf("->check.NodeNonce: %d\n", check.NodeNonce)
		fmt.Printf("->check.From: %s\n", check.From)
		fmt.Printf("->check.To: %s\n", check.To)
		fmt.Printf("->check.TokenAddress: %s\n", check.TokenAddress)

		sender_done <- 0 // signal main to continue

	case "cmd2":
		s, err := ha.NewStream(context.Background(), peerid, "/cmd2")
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("send cmd2 param")
		_, err = s.Write([]byte("param for cmd2\n"))
		if err != nil {
			log.Println(err)
			return
		}

		// get result
		out, err := ioutil.ReadAll(s)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("cmd2 result: %q\n", out)
	}

}

//
func exeCmd1(s network.Stream) error {
	buf := bufio.NewReader(s)
	str, err := buf.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Printf("received data: %s", str)
	fmt.Printf("constructing and sending struct data...\n")
	// construct data
	check := &pb.DownloadCheck{}
	check.MaxAmount = 111111
	check.NodeNonce = 222222
	check.From = "aaa"
	check.To = "bbb"
	check.TokenAddress = "tokenaddress"
	// serialize
	out, err := proto.Marshal(check)
	if err != nil {
		log.Fatalln("Failed to encode check:", err)
	}

	// send data
	_, err = s.Write([]byte(out))

	fmt.Printf("\n> ")
	fmt.Printf("intput target and cmd: \n")

	return err
}
func exeCmd2(s network.Stream) error {
	buf := bufio.NewReader(s)
	str, err := buf.ReadString('\n')
	if err != nil {
		return err
	}

	log.Printf("executing cmd2 with param: %s", str)

	result := "cmd2 execute complete"
	_, err = s.Write([]byte(result))

	return err
}
