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
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/rockiecn/p2pdemo/cash"
	"github.com/rockiecn/p2pdemo/clientops"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/print"
)

// CallApplyCheque - send tx to contract to call apply cheque method.
func CallApplyCheque(
	userAddr common.Address,
	nonce *big.Int,
	stAddr common.Address,
	payAmount *big.Int,
	sig []byte) error {
	fmt.Println("HOST: ", hostops.HOST)
	cli, err := clientops.GetClient(hostops.HOST)
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

	// ====== read cashAddr from db
	// create/open db
	db, err := leveldb.OpenFile("./data.db", nil)
	if err != nil {
		log.Fatal("opfen db error")
	}
	//
	byteCashAddr, errGet := db.Get([]byte("cashAddr"), nil)
	if errGet != nil {
		print.Println100ms("db get data error")
	}
	//
	fmt.Printf("-> cash contract address: %s\n", string(byteCashAddr))
	db.Close()

	// get contract instance from address
	cashInstance, err := cash.NewCash(common.HexToAddress(string(byteCashAddr)), cli)

	if err != nil {
		fmt.Println("NewCash err: ", err)
		return err
	}

	// fmt.Println("NewCash success: ", cashInstance)
	// fmt.Println("=== in callcash.go")
	// fmt.Println("auth:", auth)
	// fmt.Println("userAddr:", userAddr)
	// fmt.Println("nonce:", nonce)
	// fmt.Println("stAddr:", stAddr)
	// fmt.Println("payAmount:", payAmount)
	// fmt.Printf("sig:\n0x%x\n", sig)

	// call send trasaction to contract
	_, err = cashInstance.ApplyCheque(auth, userAddr, nonce, stAddr, payAmount, sig)
	if err != nil {
		fmt.Println("tx failed :", err)
		return err
	}

	fmt.Println("-> Now mine a block to complete tx.")

	return err
}

// CallDeploy - deploy cash contract
func CallDeploy() (common.Address, error) {
	var cashAddr common.Address

	fmt.Println("HOST: ", hostops.HOST)
	client, err := clientops.GetClient(hostops.HOST)
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

	// ====== store cashAddr to db
	// create/open db
	db, err := leveldb.OpenFile("./data.db", nil)
	if err != nil {
		log.Fatal("opfen db error")
	}
	// store have_purchased
	err = db.Put([]byte("cashAddr"), []byte(cashAddr.String()), nil)
	if err != nil {
		print.Println100ms("db put data error")
	}
	db.Close()

	return cashAddr, nil

}
