package callcash

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/rockiecn/p2pdemo/cash"
	"github.com/rockiecn/p2pdemo/clientops"
)

// host
const HOST = "http://localhost:8545"

// ApplyCheque(opts *bind.TransactOpts,
// 	userAddr common.Address,
// 	nonce *big.Int,
// 	stAddr common.Address,
// 	payAmount *big.Int,
// 	sign []byte)
//
func CallApplyCheque(
	userAddr common.Address,
	nonce *big.Int,
	stAddr common.Address,
	payAmount *big.Int,
	sig []byte) error {
	fmt.Println("HOST: ", HOST)
	cli, err := clientops.GetClient(HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return err
	}
	defer cli.Close()

	// operator's sk
	hexSk := "cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d"
	auth, err := clientops.MakeAuth(hexSk, nil, nil, big.NewInt(1000), 3000000)
	if err != nil {
		return err
	}

	// need cash address
	//cashInstance, err := cash.NewCash(common.HexToAddress("0x77AA1d64C1E85Cc4AF38046FfF5bc35e394f8eAD"), cli)
	cashInstance, err := cash.NewCash(common.HexToAddress("0x5854e2529F2b6d853afb2a9Da7094A0e6C9eE802"), cli)

	if err != nil {
		fmt.Println("NewCash err: ", err)
		return err
	} else {
		fmt.Println("NewCash success: ", cashInstance)
	}

	//
	//fmt.Printf("n = %s\n", n.String())

	// address to receive money
	// toAddress := common.HexToAddress("0xb213d01542d129806d664248a380db8b12059061")

	fmt.Println("=== in callcash.go")
	fmt.Println("auth:", auth)
	fmt.Println("userAddr:", userAddr)
	fmt.Println("nonce:", nonce)
	fmt.Println("stAddr:", stAddr)
	fmt.Println("payAmount:", payAmount)
	fmt.Printf("sig:\n0x%x\n", sig)

	tx, err := cashInstance.ApplyCheque(auth, userAddr, nonce, stAddr, payAmount, sig)
	if err != nil {
		fmt.Println("tx failed :", err)
		return err
	}

	fmt.Println("tx:", tx)

	return err
}

//
func CallDeploy() (common.Address, error) {
	var cashAddr common.Address

	fmt.Println("HOST: ", HOST)
	client, err := clientops.GetClient(HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return cashAddr, err
	}
	defer client.Close()

	// get sk
	//9e0153496067c20943724b79515472195a7aedaa
	sk, err := crypto.HexToECDSA("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
	if err != nil {
		fmt.Println("HexToECDSA err: ", err)
	} else {
		fmt.Println("get sk: ", sk)
	}

	//
	publicKey := sk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	//
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	auth, err := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))
	if err != nil {
		log.Panic("NewKeyedTransactorWithChainID err:", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	// string to bigint
	bn := new(big.Int)
	bn, ok1 := bn.SetString("100000000000000000000", 10) // deploy 100 eth
	//bn, ok1 := bn.SetString("1000", 10) // deploy 100 eth
	if !ok1 {
		fmt.Println("SetString: error")
		panic("SetString error")
	}
	auth.Value = bn                 // deploy 100 eth
	auth.GasLimit = uint64(7000000) // in units
	auth.GasPrice = gasPrice

	fmt.Printf("auth success: %v\n", auth)

	cashAddr, _, _, err = cash.DeployCash(auth, client)
	if err != nil {
		log.Println("deployCashErr:", err)
		return cashAddr, err
	}
	log.Println("cashAddr:", cashAddr.String())
	log.Println("value:", auth.Value.String())
	return cashAddr, nil

}
