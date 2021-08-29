package callcash

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/rockiecn/p2pdemo/cash"
	"github.com/rockiecn/p2pdemo/clientops"
	"github.com/rockiecn/p2pdemo/global"
	"github.com/rockiecn/p2pdemo/hostops"
)

// CallApplyCheque - send tx to contract to call apply cheque method.
// func CallApplyPayCheque(
// 	userAddr common.Address,
// 	nonce *big.Int,
// 	stAddr common.Address,
// 	PayValue *big.Int,
// 	sig []byte) error {
// 	fmt.Println("HOST: ", hostops.HOST)
func CallApplyPayCheque(paycheque cash.PayCheque, paychequeSig []byte) error {
	cli, err := clientops.GetClient(hostops.HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return err
	}
	defer cli.Close()

	// string decode to hex
	storageHexByte, err := hex.DecodeString(global.StrStorageSK)
	if err != nil {
		fmt.Println("callcash.go. decode string error", err)
		return err
	}
	// byte to string
	storageSKHexString := hex.EncodeToString(storageHexByte)
	// fmt.Printf("hex strStorageSK x: %x\n", storageSKHexString)
	// fmt.Printf("hex strStorageSK s: %s\n", storageSKHexString)
	auth, err := clientops.MakeAuth(storageSKHexString, nil, nil, big.NewInt(1000), 9000000)
	if err != nil {
		return err
	}

	// get contract instance from address
	cashInstance, err := cash.NewCash(paycheque.Cheque.ContractAddr, cli)
	if err != nil {
		fmt.Println("NewCash err: ", err)
		return err
	}

	fmt.Printf("cheque.value: %s\n", paycheque.Cheque.Value)
	fmt.Printf("cheque.TokenAddr: %s\n", paycheque.Cheque.TokenAddr)
	fmt.Printf("cheque.Nonce: %s\n", paycheque.Cheque.Nonce.String())
	fmt.Printf("cheque.FromAddr: %s\n", paycheque.Cheque.FromAddr)
	fmt.Printf("cheque.ToAddr: %s\n", paycheque.Cheque.ToAddr)
	fmt.Printf("cheque.OpAddr: %s\n", paycheque.Cheque.OpAddr)
	fmt.Printf("cheque.ContractAddress: %s\n", paycheque.Cheque.ContractAddr)
	fmt.Printf("paycheque.ChequeSig: %x\n", paycheque.ChequeSig)
	fmt.Printf("paycheque.PayValue: %s\n", paycheque.PayValue.String())
	fmt.Printf("paychequeSig: %x\n", paychequeSig)

	_, err = cashInstance.ApplyCheque(auth, paycheque, paychequeSig)
	if err != nil {
		fmt.Println("tx failed :", err)
		return err
	}

	fmt.Println("-> Now mine a block to complete tx.")

	return nil
}

// CallDeploy - deploy cash contract
func CallDeploy() (common.Address, error) {
	var contractAddr common.Address

	fmt.Println("HOST: ", hostops.HOST)
	client, err := clientops.GetClient(hostops.HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return contractAddr, err
	}
	defer client.Close()

	// get sk
	//9e0153496067c20943724b79515472195a7aedaa
	sk, err := crypto.HexToECDSA("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
	if err != nil {
		fmt.Println("HexToECDSA err: ", err)
		return contractAddr, err
	}
	fmt.Println("get sk: ", sk)

	// get pubkey
	publicKey := sk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
		return contractAddr, err
	}

	// pubkey to address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return contractAddr, err
	}

	// get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return contractAddr, err
	}

	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	auth, err := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))
	if err != nil {
		log.Println("NewKeyedTransactorWithChainID err:", err)
		return contractAddr, err
	}

	// set nonce
	auth.Nonce = big.NewInt(int64(nonce))
	// string to bigint
	bn := new(big.Int)
	bn, ok1 := bn.SetString("100000000000000000000", 10) // deploy 100 eth
	//bn, ok1 := bn.SetString("1000000000000000000", 10) // deploy 1 eth
	//bn, ok1 := bn.SetString("1000", 10) // deploy 100 eth
	if !ok1 {
		fmt.Println("SetString: error")
		fmt.Println("big number SetString error")
		return contractAddr, err
	}
	auth.Value = bn                 // deploy 100 eth
	auth.GasLimit = uint64(7000000) // in units
	auth.GasPrice = gasPrice

	fmt.Printf("auth success: %v\n", auth)

	contractAddr, _, _, err = cash.DeployCash(auth, client)
	if err != nil {
		log.Println("deployCashErr:", err)
		return contractAddr, err
	}
	log.Println("contractAddr:", contractAddr.String())
	log.Println("value:", auth.Value.String())

	// ====== store contractAddr to db
	// create/open db
	db, err := leveldb.OpenFile("./operator_data.db", nil)
	if err != nil {
		log.Print("opfen db error:", err)
		return contractAddr, err
	}
	// store cash address
	err = db.Put([]byte("contractAddr"), []byte(contractAddr.String()), nil)
	if err != nil {
		log.Println("db put data error:", err)
		return contractAddr, err
	}
	db.Close()

	return contractAddr, nil

}
