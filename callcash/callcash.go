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
func CallApplyPayCheque(paycheque cash.PayCheque) error {
	cli, err := clientops.GetClient(hostops.HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return err
	}
	defer cli.Close()

	/*
		[]string:
		Cheque - string operator_address = 1;
		Cheque - string from = 2;
		Cheque - string to =3;
		Cheque - string token_address = 4;
		PayCheque - string cash_address = 3;

		[]int64:
		Cheque - int64 value = 5;
		Cheque - int64 node_nonce = 6;
		PayCheque - int64 pay_value = 6;

		[]byte:
		PayCheque - bytes cheque_sig = 2; //运营商对cheque的签名
	*/

	/*
		// byte to string
		strOpSK := hex.EncodeToString(global.OperatorSK)
		auth, err := clientops.MakeAuth(strOpSK, nil, nil, big.NewInt(1000), 3000000)
		if err != nil {
			return err
		}

		// TODO:get cash address from params

		// get contract instance from address
		cashInstance, err := cash.NewCash(common.HexToAddress(string(byteCashAddr)), cli)
		if err != nil {
			fmt.Println("NewCash err: ", err)
			return err
		}
	*/

	// fmt.Println("NewCash success: ", cashInstance)
	// fmt.Println("=== in callcash.go")
	// fmt.Println("auth:", auth)
	// fmt.Println("userAddr:", userAddr)
	// fmt.Println("nonce:", nonce)
	// fmt.Println("stAddr:", stAddr)
	// fmt.Println("PayValue:", PayValue)
	// fmt.Printf("sig:\n0x%x\n", sig)

	/*
		// call send trasaction to contract

			_, err = cashInstance.ApplyCheque(auth, userAddr, nonce, stAddr, PayValue, sig)
			if err != nil {
				fmt.Println("tx failed :", err)
				return err
			}

			fmt.Println("-> Now mine a block to complete tx.")
	*/

	// string decode to hex
	storageHexByte, err := hex.DecodeString(global.StrStorageSK)
	if err != nil {
		fmt.Println("callcash.go. decode string error", err)
		return err
	}
	// byte to string
	storageSKHexString := hex.EncodeToString(storageHexByte)
	fmt.Printf("hex strStorageSK x: %x\n", storageSKHexString)
	fmt.Printf("hex strStorageSK s: %s\n", storageSKHexString)
	auth, err := clientops.MakeAuth(storageSKHexString, nil, nil, big.NewInt(1000), 9000000)
	if err != nil {
		return err
	}

	// get contract instance from address
	cashInstance, err := cash.NewCash(common.HexToAddress(paycheque.CashAddr), cli)
	if err != nil {
		fmt.Println("NewCash err: ", err)
		return err
	}

	fmt.Printf("CashAddr: %s\n", paycheque.CashAddr)
	fmt.Printf("ChequeSig: %x\n", paycheque.ChequeSig)
	fmt.Printf("FromAddr: %s\n", paycheque.FromAddr)
	fmt.Printf("ToAddr: %s\n", paycheque.ToAddr)
	fmt.Printf("PayValue: %s\n", paycheque.PayValue.String())
	fmt.Printf("Cheque.FromAddr: %s\n", paycheque.Cheque.FromAddr)
	fmt.Printf("Cheque.OpAddr: %s\n", paycheque.Cheque.OpAddr)
	fmt.Printf("Cheque.ToAddr: %s\n", paycheque.Cheque.ToAddr)
	fmt.Printf("Cheque.NodeNonce: %s\n", paycheque.Cheque.NodeNonce.String())

	_ = auth
	_ = cashInstance

	_, err = cashInstance.ApplyCheque(auth, paycheque)
	if err != nil {
		fmt.Println("tx failed :", err)
		return err
	}

	fmt.Println("-> Now mine a block to complete tx.")

	return nil
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
		return cashAddr, err
	}
	fmt.Println("get sk: ", sk)

	// get pubkey
	publicKey := sk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
		return cashAddr, err
	}

	// pubkey to address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return cashAddr, err
	}

	// get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return cashAddr, err
	}

	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	auth, err := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))
	if err != nil {
		log.Println("NewKeyedTransactorWithChainID err:", err)
		return cashAddr, err
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
		return cashAddr, err
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
	db, err := leveldb.OpenFile("./operator_data.db", nil)
	if err != nil {
		log.Print("opfen db error:", err)
		return cashAddr, err
	}
	// store cash address
	err = db.Put([]byte("cashAddr"), []byte(cashAddr.String()), nil)
	if err != nil {
		log.Println("db put data error:", err)
		return cashAddr, err
	}
	db.Close()

	return cashAddr, nil

}
