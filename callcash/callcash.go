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
	"github.com/rockiecn/p2pdemo/global"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/print"
)

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
	sk, err := crypto.HexToECDSA(global.StrOperatorSK)
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
	if !ok1 {
		fmt.Println("SetString: error")
		fmt.Println("big number SetString error")
		return contractAddr, err
	}
	auth.Value = bn                 // deploy 100 eth
	auth.GasLimit = uint64(7000000) // in units
	auth.GasPrice = gasPrice

	print.Printf100ms("auth success: %v\n", auth)

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

	global.ContractAddress = contractAddr.String()

	return contractAddr, nil
}

// CallApplyCheque - send tx to contract to call apply cheque method.
func CallApplyPayCheque(paycheque cash.PayCheque, paychequeSig []byte) error {
	cli, err := clientops.GetClient(hostops.HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return err
	}
	defer cli.Close()

	auth, err := clientops.MakeAuth(global.StrStorageSK, nil, nil, big.NewInt(1000), 9000000)
	if err != nil {
		return err
	}

	// get contract instance from address
	cashInstance, err := cash.NewCash(paycheque.Cheque.ContractAddr, cli)
	if err != nil {
		fmt.Println("NewCash err: ", err)
		return err
	}

	print.Printf100ms("cheque.value: %s\n", paycheque.Cheque.Value)
	print.Printf100ms("cheque.TokenAddr: %s\n", paycheque.Cheque.TokenAddr)
	print.Printf100ms("cheque.Nonce: %s\n", paycheque.Cheque.Nonce.String())
	print.Printf100ms("cheque.FromAddr: %s\n", paycheque.Cheque.FromAddr)
	print.Printf100ms("cheque.ToAddr: %s\n", paycheque.Cheque.ToAddr)
	print.Printf100ms("cheque.OpAddr: %s\n", paycheque.Cheque.OpAddr)
	print.Printf100ms("cheque.ContractAddress: %s\n", paycheque.Cheque.ContractAddr)
	print.Printf100ms("paycheque.ChequeSig: %x\n", paycheque.ChequeSig)
	print.Printf100ms("paycheque.PayValue: %s\n", paycheque.PayValue.String())
	print.Printf100ms("paychequeSig: %x\n", paychequeSig)

	_, err = cashInstance.ApplyCheque(auth, paycheque, paychequeSig)
	if err != nil {
		fmt.Println("tx failed :", err)
		return err
	}

	fmt.Println("-> Now mine a block to complete tx.")

	return nil
}

// call get contract node nonce
func CallGetNodeNonce(node common.Address) error {
	cli, err := clientops.GetClient(hostops.HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return err
	}
	defer cli.Close()

	// ====== get contractAddr from db
	// create/open db
	db, err := leveldb.OpenFile("./operator_data.db", nil)
	if err != nil {
		log.Print("opfen db error:", err)
		return err
	}
	// store cash address
	byteContractAddr, err := db.Get([]byte("contractAddr"), nil)
	if err != nil {
		log.Println("db put data error:", err)
		return err
	}
	db.Close()

	AddressContract := common.HexToAddress(string(byteContractAddr))
	// get contract instance from address
	cashInstance, err2 := cash.NewCash(AddressContract, cli)
	if err2 != nil {
		fmt.Println("NewCash err: ", err2)
		return err2
	}

	bigNonce, err := cashInstance.GetNodeNonce(nil, node)
	if err != nil {
		fmt.Println("tx failed :", err)
		return err
	}

	fmt.Println()
	fmt.Println("Node: ", node.String())
	fmt.Println("Node nonce: ", bigNonce.String())

	return nil
}
