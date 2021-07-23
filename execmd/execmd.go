package execmd

import (
	"bufio"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p-core/network"
	pb "github.com/rockiecn/p2pdemo/check_go"
	"google.golang.org/protobuf/proto"
)

//
func ExeCmd1(s network.Stream) error {
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
func ExeCmd2(s network.Stream) error {
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
