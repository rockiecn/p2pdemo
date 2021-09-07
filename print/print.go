package print

import (
	"fmt"
	"time"

	"github.com/rockiecn/p2pdemo/global"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
)

var ContractAddress string

func PrintCheque(Cheque *pb.Cheque) {
	Println100ms("------------------ Print Cheque ------------------")
	Printf100ms("->Cheque.Value: %s\n", Cheque.Value)
	Printf100ms("->Cheque.Nonce: %s\n", Cheque.Nonce)
	Printf100ms("->Cheque.OperatorAddress: %s\n", Cheque.OperatorAddress)
	Printf100ms("->Cheque.From: %s\n", Cheque.From)
	Printf100ms("->Cheque.TokenAddress: %s\n", Cheque.TokenAddress)
	Println100ms("----------------------------------------------------")
}

func PrintPayCheque(PayCheque *pb.PayCheque) {
	Println100ms("-------------------- Print PayCheque ------------------")
	Printf100ms("->Cheque.Value: %s\n", PayCheque.Cheque.Value)
	Printf100ms("->Cheque.TokenAddress: %s\n", PayCheque.Cheque.TokenAddress)
	Printf100ms("->Cheque.NodeNonce: %s\n", PayCheque.Cheque.Nonce)
	Printf100ms("->Cheque.From: %s\n", PayCheque.Cheque.From)
	Printf100ms("->Cheque.To: %s\n", PayCheque.Cheque.To)
	Printf100ms("->Cheque.OperatorAddress: %s\n", PayCheque.Cheque.OperatorAddress)
	Printf100ms("->PayCheque.CashAddress: %s\n", PayCheque.Cheque.ContractAddress)
	Printf100ms("->PayCheque.ChequeSig: %x\n", PayCheque.ChequeSig)
	Printf100ms("->PayCheque.PayValue: %s\n", PayCheque.PayValue)
	Println100ms("----------------------------------------------------")
}

// println with 100 ms delay
func Println100ms(str string) {
	fmt.Println(str)
	time.Sleep(5 * time.Millisecond)
}

// printf with 100 ms delay
func Printf100ms(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	time.Sleep(5 * time.Millisecond)
}

// print command menu
func PrintMenu() {

	var showPeer string
	if global.RemoteExist {
		showPeer = global.Peerid.Pretty()
	} else {
		showPeer = "No emote Peer"
	}

	// db, err := leveldb.OpenFile("./operator_data.db", nil)
	// if err != nil {
	// 	fmt.Println("Open db error")
	// }
	// defer db.Close()

	// var contractAddrByte []byte
	// contractAddrByte, err = db.Get([]byte("contractAddr"), nil)
	// if err != nil {
	// 	fmt.Println("Get cash address error:", err)
	// 	contractAddress = "none"
	// } else {
	// 	contractAddress = string(contractAddrByte)
	// }

	Println100ms("")
	Println100ms("")
	Println100ms("")
	Println100ms("")
	Println100ms("+++++++++++++++++++++")
	Println100ms("| Welcom to p2pdemo |")
	Println100ms("+++++++++++++++++++++")
	Println100ms("")
	Printf100ms("Remote Peer: %s\n", showPeer)
	Printf100ms("Contract Address: %s\n", ContractAddress)
	Println100ms("")
	Println100ms("              ======================= Menu =======================")
	Println100ms("               m   : [ALL]       Show menu")
	Println100ms("               r   : [ALL]       Record remote peer")
	Println100ms("               d  +: [OPERATOR]  Deploy cash contract")
	Println100ms("               gn +: [OPERATOR]  Get node nonce in contract")
	Println100ms("               re  : [OPERATOR]  Reset nonce in db")
	Println100ms("               sn  : [OPERATOR]  Show nonce in db")
	Println100ms("               g  *: [USER]      Get Cheque from operator")
	Println100ms("               s  *: [USER]      Send a PayCheque to storage")
	Println100ms("               lu  : [USER]      List user's PayCheque table")
	Println100ms("               du  : [USER]      Delete a paycheque from user db")
	Println100ms("               is *: [USER]      Increase payvalue and send cheque")
	Println100ms("               cu  : [USER]      Clear user db")
	Println100ms("               sh  : [USER]      Show paycheque info")
	Println100ms("               ls  : [STORAGE]   List storage's PayCheque table")
	Println100ms("               ds  : [STORAGE]   Delete a paycheque from storage db")
	Println100ms("               cc +: [STORAGE]   Call apply cheque")
	Println100ms("               cs  : [STORAGE]   Clear storage db")
	//fmt.Println("               t  +: [TEST]      Call retrieve in storage")
	fmt.Println("              ====================================================")
	Println100ms("")
	Println100ms("+: Need chain running.")
	Println100ms("*: Need remote peer, run command 'r'.")
	Println100ms("")

	fullAddr := hostops.GetHostAddress(hostops.HostInfo)
	Printf100ms("\nLocal Peer addres: \n[ %s ]\n", fullAddr)
}
