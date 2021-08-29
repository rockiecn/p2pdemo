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
	Printf100ms("->Cheque.Nonce: %d\n", Cheque.Nonce)
	Printf100ms("->Cheque.OperatorAddress: %s\n", Cheque.OperatorAddress)
	Printf100ms("->Cheque.From: %s\n", Cheque.From)
	Printf100ms("->Cheque.TokenAddress: %s\n", Cheque.TokenAddress)
	Println100ms("----------------------------------------------------")
}

func PrintPayCheque(PayCheque *pb.PayCheque) {
	Println100ms("-------------------- Print PayCheque ------------------")
	Printf100ms("->Cheque.Value: %d\n", PayCheque.Cheque.Value)
	Printf100ms("->Cheque.TokenAddress: %s\n", PayCheque.Cheque.TokenAddress)
	Printf100ms("->Cheque.NodeNonce: %d\n", PayCheque.Cheque.Nonce)
	Printf100ms("->Cheque.From: %s\n", PayCheque.Cheque.From)
	Printf100ms("->Cheque.To: %s\n", PayCheque.Cheque.To)
	Printf100ms("->Cheque.OperatorAddress: %s\n", PayCheque.Cheque.OperatorAddress)
	Printf100ms("->PayCheque.CashAddress: %s\n", PayCheque.Cheque.ContractAddress)
	Printf100ms("->PayCheque.ChequeSig: %x\n", PayCheque.ChequeSig)
	Printf100ms("->PayCheque.PayValue: %d\n", PayCheque.PayValue)
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
	fmt.Println("               r  -: [ALL]       Record remote peer")
	fmt.Println("               d  +: [OPERATOR]  Deploy cash contract")
	fmt.Println("               g  *: [USER]      Get Cheque from operator")
	fmt.Println("               s  *: [USER]      Send a PayCheque to storage")
	fmt.Println("               lu -: [USER]      List user's PayCheque table")
	fmt.Println("               du -: [USER]      Delete a paycheque from user db")
	fmt.Println("               is *: [USER]      Increase payvalue and send cheque")
	fmt.Println("               cu -: [USER]      Clear user db")
	fmt.Println("               sh -: [USER]      Show paycheque info")
	fmt.Println("               ls -: [STORAGE]   List storage's PayCheque table")
	fmt.Println("               ds -: [STORAGE]   Delete a paycheque from storage db")
	fmt.Println("               cc +: [STORAGE]   Call apply cheque")
	fmt.Println("               cs -: [STORAGE]   Clear storage db")
	//fmt.Println("               t  +: [TEST]      Call retrieve in storage")
	fmt.Println("              ====================================================")

	fullAddr := hostops.GetHostAddress(hostops.HostInfo)
	Printf100ms("\nLocal Peer addres: \n[ %s ]\n", fullAddr)
}
