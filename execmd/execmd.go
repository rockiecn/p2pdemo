package execmd

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/test-sig/sig/implement/sigapi"
	"github.com/rockiecn/test-sig/sig/implement/utils"
	"google.golang.org/protobuf/proto"
)

var (
	opSkByte = []byte("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
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
	cheque.PurchaseAmount = 111111
	cheque.NodeNonce = 222222
	cheque.OperatorAddress = "aaa"
	cheque.UserAddress = "bbb"
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
	fmt.Printf("--> Construct and send purchase...\n")
	time.Sleep(100 * time.Millisecond)

	// construct
	Purchase := &pb.Purchase{}
	Purchase.PurchaseAmount = 100
	Purchase.NodeNonce = 1

	Purchase.OperatorAddress = "9e0153496067c20943724b79515472195a7aedaa" // operator
	//cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d

	Purchase.UserAddress = "1ab6a9f2b90004c1269563b5da391250ede3c114" // user
	//b91c265cabae210642d66f9d59137eac2fab2674f4c1c88df3b8e9e6c1f74f9f
	Purchase.TokenAddress = "tokenaddress"

	// serialize
	purchase_marshaled, err := proto.Marshal(Purchase)
	if err != nil {
		log.Fatalln("Failed to encode cheque:", err)
	}

	// construct purchase message: sig(65 bytes) | data
	fmt.Println("-> constructing msg")
	time.Sleep(100 * time.Millisecond)

	// sign purchase
	sig, err := sigapi.Sign(purchase_marshaled, opSkByte)
	if err != nil {
		panic("sign error")
	}
	//
	// SigPurchase = sig // store signature of purchase

	//fmt.Printf("sig len:%d\n", len(sig))                               //65
	//fmt.Printf("purchase_marshaled len:%d\n", len(purchase_marshaled)) //50

	var msg = []byte{}
	msg = utils.MergeSlice(sig, purchase_marshaled)

	fmt.Println("-> sending msg")
	time.Sleep(100 * time.Millisecond)
	// send msg
	_, err = s.Write([]byte(msg))
	if err != nil {
		panic("stream write error")
	}

	fmt.Println("\n> Intput target address and cmd: ")
	time.Sleep(100 * time.Millisecond)

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

	// parse data
	var sigByte = in[:65]
	var cheque_marshaled = in[65:]
	//fmt.Println("cheque_marshaled:", cheque_marshaled)

	//fmt.Printf("sigByte len:%d\n", len(sigByte))                   //65
	//fmt.Printf("cheque_marshaled len:%d\n", len(cheque_marshaled)) //

	// unmarshal data
	cheque := &pb.Cheque{}
	if err := proto.Unmarshal(cheque_marshaled, cheque); err != nil {
		log.Fatalln("Failed to parse check:", err)
		return err
	}

	//
	PrintCheque(cheque)

	// verify signature of cheque

	// get user address
	userAddrByte, err := hex.DecodeString(cheque.Purchase.UserAddress)
	if err != nil {
		panic("decode error")
	}
	// []byte to common.Address
	userAddress := common.BytesToAddress(userAddrByte)

	// verify
	ok, verErr := sigapi.Verify(cheque_marshaled, sigByte, userAddress)
	if verErr != nil {
		log.Fatal("verify fatal error occured")
	}

	if ok {
		fmt.Println("<signature of cheque verify success>")
		// Have_Purchased = true // already have purchased
	} else {
		fmt.Println("<signature of cheque verify failed>")
	}

	fmt.Println("\n> Intput target address and cmd: ")
	time.Sleep(100 * time.Millisecond)

	return nil
}

func PrintPurchase(purchase *pb.Purchase) {
	fmt.Println("------------------ Print Purchase ------------------")
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.PurchaseAmount: %d\n", purchase.PurchaseAmount)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.NodeNonce: %d\n", purchase.NodeNonce)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.OperatorAddress: %s\n", purchase.OperatorAddress)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.UserAddress: %s\n", purchase.UserAddress)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.TokenAddress: %s\n", purchase.TokenAddress)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("----------------------------------------------------")
	time.Sleep(100 * time.Millisecond)
}

func PrintCheque(cheque *pb.Cheque) {
	fmt.Println("-------------------- Print Cheque ------------------")
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.MaxAmount: %d\n", cheque.Purchase.PurchaseAmount)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.NodeNonce: %d\n", cheque.Purchase.NodeNonce)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.OperatorAddress: %s\n", cheque.Purchase.OperatorAddress)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.UserAddress: %s\n", cheque.Purchase.UserAddress)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->purchase.TokenAddress: %s\n", cheque.Purchase.TokenAddress)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->cheque.PurchaseSig: %x\n", cheque.PurchaseSig)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->cheque.PayAmount: %d\n", cheque.PayAmount)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("->cheque.StorageAddress: %s\n", cheque.StorageAddress)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("----------------------------------------------------")
	time.Sleep(100 * time.Millisecond)
}
