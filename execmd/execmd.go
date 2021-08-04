package execmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/test-sig/sig/implement/sigapi"
	"github.com/rockiecn/test-sig/sig/implement/utils"
	"google.golang.org/protobuf/proto"
)

var (
	skByte = []byte("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
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
	cheque := &pb.Purchase{}
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

// 1 handler
func SendPurchase(s network.Stream) error {

	//fmt.Printf("Received data: %s", str)
	fmt.Printf("---> Construct and send purchase...\n")
	time.Sleep(100 * time.Millisecond)
	// construct
	purchase := &pb.Purchase{}
	purchase.MaxAmount = 100
	purchase.NodeNonce = 1
	purchase.From = "9e0153496067c20943724b79515472195a7aedaa"
	purchase.To = "1ab6a9f2b90004c1269563b5da391250ede3c114"
	purchase.TokenAddress = "tokenaddress"
	// serialize
	purchase_marshaled, err := proto.Marshal(purchase)
	if err != nil {
		log.Fatalln("Failed to encode cheque:", err)
	}

	// construct msg
	// sig(65 bytes) | data
	// sign
	sig, err := sigapi.Sign(purchase_marshaled, skByte)
	if err != nil {
		panic("sign error")
	}

	//fmt.Printf("sig:%s\n", sig)
	//fmt.Printf("purchase_marshaled:%s\n", purchase_marshaled)

	fmt.Printf("sig len:%d\n", len(sig))                               //65
	fmt.Printf("purchase_marshaled len:%d\n", len(purchase_marshaled)) //50

	var msg = []byte{}
	msg = utils.MergeSlice(sig, purchase_marshaled)
	fmt.Printf("msg len:%d\n", len(msg))

	// send msg
	_, err = s.Write([]byte(msg))
	if err != nil {
		panic("stream write error")
	}

	fmt.Println("\n> Intput target address and cmd: ")

	return err
}

// 2 handler
func ReceiveCheque(s network.Stream) error {

	// // Read data method 1
	// in := make([]byte, 1024)
	// reader := bufio.NewReader(s)
	// n, err := reader.Read(in)
	// if err != nil {
	// 	fmt.Println("read err: ", err)
	// }
	// // get real data
	// if n > 0 {
	// 	in = in[:n]
	// }
	// fmt.Printf("in: %v", in)

	// // Read data method 2
	// reader := bufio.NewReader(s)
	// in, err := reader.ReadBytes('\n')
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("read: %v", in)

	// Read data method 3
	// Caution: Need writer close stream to continue.
	in, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatalln("Error reading :", err)
		return err
	}

	// unmarshal data
	cheque := &pb.Cheque{}
	if err := proto.Unmarshal(in, cheque); err != nil {
		log.Fatalln("Failed to parse check:", err)
		return err
	}

	fmt.Printf("Print cheque:\n")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("--------------------- cheque -------------------")
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.MaxAmount: %d\n", cheque.Purchase.MaxAmount)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.NodeNonce: %d\n", cheque.Purchase.NodeNonce)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.From: %s\n", cheque.Purchase.From)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.To: %s\n", cheque.Purchase.To)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.TokenAddress: %s\n", cheque.Purchase.TokenAddress)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->cheque.DownloadSign: %x\n", cheque.DownloadSign)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->cheque.PayAmount: %d\n", cheque.PayAmount)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->cheque.OperatorAddress: %s\n", cheque.OperatorAddress)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("----------------------------------------------------")
	time.Sleep(100 * time.Millisecond)

	return nil
}
