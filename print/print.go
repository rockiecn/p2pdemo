package print

import (
	"fmt"
	"time"

	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
)

func PrintPurchase(purchase *pb.Purchase) {
	Println100ms("------------------ Print Purchase ------------------")
	Printf100ms("->purchase.PurchaseAmount: %d\n", purchase.PurchaseAmount)
	Printf100ms("->purchase.NodeNonce: %d\n", purchase.NodeNonce)
	Printf100ms("->purchase.OperatorAddress: %s\n", purchase.OperatorAddress)
	Printf100ms("->purchase.UserAddress: %s\n", purchase.UserAddress)
	Printf100ms("->purchase.TokenAddress: %s\n", purchase.TokenAddress)
	Println100ms("----------------------------------------------------")
}

func PrintCheque(cheque *pb.Cheque) {
	Println100ms("-------------------- Print Cheque ------------------")
	Printf100ms("->purchase.MaxAmount: %d\n", cheque.Purchase.PurchaseAmount)
	Printf100ms("->purchase.NodeNonce: %d\n", cheque.Purchase.NodeNonce)
	Printf100ms("->purchase.OperatorAddress: %s\n", cheque.Purchase.OperatorAddress)
	Printf100ms("->purchase.UserAddress: %s\n", cheque.Purchase.UserAddress)
	Printf100ms("->purchase.TokenAddress: %s\n", cheque.Purchase.TokenAddress)
	Printf100ms("->cheque.PurchaseSig: %x\n", cheque.PurchaseSig)
	Printf100ms("->cheque.PayAmount: %d\n", cheque.PayAmount)
	Printf100ms("->cheque.StorageAddress: %s\n", cheque.StorageAddress)
	Println100ms("----------------------------------------------------")
}

// println with 100 ms delay
func Println100ms(str string) {
	fmt.Println(str)
	time.Sleep(100 * time.Millisecond)
}

// printf with 100 ms delay
func Printf100ms(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	time.Sleep(100 * time.Millisecond)
}

// print command menu
func PrintMenu() {
	fmt.Println()
	fmt.Println("======================= Menu =======================")
	fmt.Println("cmd 1: require purchase from operator")
	fmt.Println("cmd 2: send cheque to storage")
	fmt.Println("cmd 3: call retrieve in storage")
	fmt.Println("cmd 4: call deploy cash")
	fmt.Println("cmd 5: call applycheque in cash")
	fmt.Println("cmd 6: list key")
	fmt.Println("====================================================")

	fullAddr := hostops.GetHostAddress(hostops.HostInfo)
	Printf100ms("\n[ %s ]\n", fullAddr)
}
