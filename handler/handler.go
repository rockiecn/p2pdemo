package handler

import (
	"encoding/hex"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/test-sig/sig/implement/sigapi"
	"github.com/rockiecn/test-sig/sig/implement/utils"
	"google.golang.org/protobuf/proto"

	"github.com/rockiecn/p2pdemo/print"
)

// command 1 handler, operator send purchase to user
func Cmd1Handler(s network.Stream) error {

	print.Println100ms("--> Construct and send purchase...")

	// construct purchase
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
	print.Println100ms("-> constructing msg")

	// sign purchase by operator
	var opSkByte = []byte("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
	sig, err := sigapi.Sign(purchase_marshaled, opSkByte)
	if err != nil {
		panic("sign error")
	}

	var msg = []byte{}
	msg = utils.MergeSlice(sig, purchase_marshaled)

	print.Println100ms("-> sending msg")
	// send msg
	_, err = s.Write([]byte(msg))
	if err != nil {
		panic("stream write error")
	}

	print.Println100ms("\n> Intput target address and cmd: ")

	return err
}

// command 2 handler, user send cheque to storage
func Cmd2Handler(s network.Stream) error {

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
	// Caution: Need writer to close stream first.
	in, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatalln("Error reading :", err)
		return err
	}

	// parse data
	var sigByte = in[:65]
	var cheque_marshaled = in[65:]

	// unmarshal data
	cheque := &pb.Cheque{}
	if err := proto.Unmarshal(cheque_marshaled, cheque); err != nil {
		log.Fatalln("Failed to parse check:", err)
		return err
	}

	//
	print.PrintCheque(cheque)

	//===== verify signature of cheque(signed by user)

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
		print.Println100ms("<signature of cheque verify success>")
	} else {
		print.Println100ms("<signature of cheque verify failed>")
	}

	print.Println100ms("\n> Intput target address and cmd: ")

	return nil
}
