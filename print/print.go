package print

import (
	"fmt"
	"log"
	"time"

	"github.com/rockiecn/p2pdemo/global"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
)

func PrintCheque(Cheque *pb.Cheque) {
	Println100ms("------------------ Print Cheque ------------------")
	Printf100ms("->Cheque.Value: %d\n", Cheque.Value)
	Printf100ms("->Cheque.NodeNonce: %d\n", Cheque.NodeNonce)
	Printf100ms("->Cheque.OperatorAddress: %s\n", Cheque.OperatorAddress)
	Printf100ms("->Cheque.From: %s\n", Cheque.From)
	Printf100ms("->Cheque.TokenAddress: %s\n", Cheque.TokenAddress)
	Println100ms("----------------------------------------------------")
}

func PrintPayCheque(PayCheque *pb.PayCheque) {
	Println100ms("-------------------- Print PayCheque ------------------")
	Printf100ms("->Cheque.MaxAmount: %d\n", PayCheque.Cheque.Value)
	Printf100ms("->Cheque.NodeNonce: %d\n", PayCheque.Cheque.NodeNonce)
	Printf100ms("->Cheque.OperatorAddress: %s\n", PayCheque.Cheque.OperatorAddress)
	Printf100ms("->Cheque.From: %s\n", PayCheque.Cheque.From)
	Printf100ms("->Cheque.TokenAddress: %s\n", PayCheque.Cheque.TokenAddress)
	Printf100ms("->PayCheque.ChequeSig: %x\n", PayCheque.ChequeSig)
	Printf100ms("->PayCheque.CashAddress: %s\n", PayCheque.CashAddress)
	Printf100ms("->PayCheque.PayValue: %d\n", PayCheque.PayValue)
	Printf100ms("->PayCheque.From: %s\n", PayCheque.From)
	Printf100ms("->PayCheque.To: %s\n", PayCheque.To)
	Println100ms("----------------------------------------------------")
}

// println with 100 ms delay
func Println100ms(str string) {
	fmt.Println(str)
	time.Sleep(10 * time.Millisecond)
}

// printf with 100 ms delay
func Printf100ms(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	time.Sleep(10 * time.Millisecond)
}

// print command menu
func PrintMenu() {

	var showPeer string
	var err error
	if global.RemoteExist {
		showPeer = global.Peerid.Pretty()
	} else {
		//showPeer = "No emote Peer"
		// string to peer.ID
		showPeer = "No emote Peer"
		if err != nil {
			log.Println(err)
			return
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("+++++++++++++++++++++")
	fmt.Printf("| Welcom to p2pdemo |  Remote Peer: %s\n", showPeer)
	fmt.Println("+++++++++++++++++++++")
	fmt.Println()

	fmt.Println("              ======================= Menu =======================")
	fmt.Println("               0-: [ALL]       Record remote peer")
	fmt.Println("               1+: [OPERATOR]  Deploy cash contract")
	fmt.Println("               2*: [USER]      Get Cheque from operator")
	fmt.Println("               3*: [USER]      Send PayCheque to storage")
	fmt.Println("               4-: [USER]      List PayCheque table")
	fmt.Println("               5-: [USER]      Delete a paycheque from db")
	fmt.Println("               6*: [USER]      Increase payvalue and send cheque")
	fmt.Println("               7+: [STORAGE]   Call apply cheque")
	fmt.Println("               8+: [TEST]      Call retrieve in storage")
	fmt.Println("              ====================================================")

	fullAddr := hostops.GetHostAddress(hostops.HostInfo)
	Printf100ms("\nLocal Peer addres: \n[ %s ]\n", fullAddr)
}
