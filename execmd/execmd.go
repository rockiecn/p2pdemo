package execmd

import (
	"bufio"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/rockiecn/p2pdemo/pb"
	"google.golang.org/protobuf/proto"
)

//
func ExeCmd1(s network.Stream) error {
	buf := bufio.NewReader(s)
	str, err := buf.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Printf("Received data: %s", str)
	fmt.Printf("Constructing and sending struct data...\n")
	// construct data
	cheque := &pb.DownloadCheque{}
	cheque.MaxAmount = 111111
	cheque.NodeNonce = 222222
	cheque.From = "aaa"
	cheque.To = "bbb"
	cheque.TokenAddress = "tokenaddress"
	// serialize
	out, err := proto.Marshal(cheque)
	if err != nil {
		log.Fatalln("Failed to encode cheque:", err)
	}

	// send data
	_, err = s.Write([]byte(out))

	fmt.Printf("\n> ")
	fmt.Printf("Intput target address and cmd: \n")

	return err
}
func ExeDownloadCheque(s network.Stream) error {

	//fmt.Printf("Received data: %s", str)
	fmt.Printf("Constructing and sending download cheque...\n")
	// construct
	cheque := &pb.DownloadCheque{}
	cheque.MaxAmount = 1000
	cheque.NodeNonce = 1
	cheque.From = "user address"
	cheque.To = "storage address"
	cheque.TokenAddress = "tokenaddress"
	// serialize
	out, err := proto.Marshal(cheque)
	if err != nil {
		log.Fatalln("Failed to encode cheque:", err)
	}

	fmt.Printf("out: %v\n", out)

	// send data
	_, err = s.Write([]byte(out))

	fmt.Printf("\n> ")
	fmt.Printf("Intput target address and cmd: \n")

	return err
}

func ExePayCheque(s network.Stream) error {

	// Read data.
	fmt.Println("In handler ..")

	// test
	in := make([]byte, 1024)
	reader := bufio.NewReader(s)
	n, err := reader.Read(in)
	if err != nil {
		fmt.Println("read err: ", err)
	}
	//fmt.Println("read n: ", n)

	// get real data
	if n > 0 {
		in = in[:n]
	}
	fmt.Printf("in: %v", in)

	// // unmarshal data
	// fmt.Println("unmarshaling data")
	// pay_cheque := &pb.PayCheque{}
	// //download_cheque := &pb.DownloadCheque{}
	// if err := proto.Unmarshal(in, pay_cheque); err != nil {
	// 	log.Fatalln("Failed to parse check:", err)
	// 	return err
	// }

	// fmt.Printf("Print pay cheque:\n")
	// fmt.Printf("%v", pay_cheque)

	// // ok
	// reader := bufio.NewReader(s)
	// in, err := reader.ReadBytes('\n')
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("read: %v", in)

	return nil
}
