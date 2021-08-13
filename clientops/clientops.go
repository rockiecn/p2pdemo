package clientops

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const HOST = "http://localhost:8545"

//
func GetClient(endPoint string) (*ethclient.Client, error) {
	rpcClient, err := rpc.Dial(endPoint)
	if err != nil {
		log.Println("rpc.Dial err:", err)
		return nil, err
	}

	conn := ethclient.NewClient(rpcClient)
	return conn, nil
}

//MakeAuth make the transactOpts to call contract
func MakeAuth(hexSk string, moneyToContract, nonce, gasPrice *big.Int, gasLimit uint64) (*bind.TransactOpts, error) {
	auth := &bind.TransactOpts{}
	sk, err := crypto.HexToECDSA(hexSk)
	if err != nil {
		log.Println("HexToECDSA err: ", err)
		return auth, err
	}

	auth = bind.NewKeyedTransactor(sk)
	auth.GasPrice = gasPrice
	auth.Value = moneyToContract //放进合约里的钱
	auth.Nonce = nonce
	auth.GasLimit = gasLimit
	return auth, nil
}

//
func QueryBalance(account string) (*big.Int, error) {
	var result string

	client, err := rpc.Dial(HOST)
	if err != nil {
		log.Println("rpc.dial err:", err)
		return big.NewInt(0), err
	}

	retryCount := 0
	for {
		retryCount++

		err = client.Call(&result, "eth_getBalance", account, "latest")

		if err != nil {
			if retryCount > 3 {
				return big.NewInt(0), err
			}
			time.Sleep(1000)
			continue
		} else {
			log.Println("call getbalance success: result = ", result)
		}
		//balance := utils.HexToBigInt(result)

		// hex to bitInt
		trimResult := strings.TrimPrefix(result, "0x")
		balance := new(big.Int)
		balance.SetString(trimResult, 16)

		log.Println("in queryBalance")
		log.Println("result:", result)
		log.Println("balance:", balance)
		return balance, nil

		//log.Println("balance:", result)
		//return nil, nil
	}
}

//
func TransferTo(value *big.Int, toAddr, eth string) error {

	client, err := ethclient.Dial(eth)
	if err != nil {
		fmt.Println("rpc.Dial err", err)
		log.Fatal(err)
	}

	//account0 cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d
	//account2 920ffe90f05741f3b27aeec8a843f870d51b2a2d30d65afcb7390c0851af39f3
	privateKey, err := crypto.HexToECDSA("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("from addr:", fromAddress.String())

	gasLimit := uint64(23000)           // in units
	gasPrice := big.NewInt(30000000000) // in wei (30 gwei)

	toAddress := common.HexToAddress(toAddr[2:])
	log.Println("toAddress: ", toAddress)

	var chainID *big.Int
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("client.NetworkID error,use the default chainID")
		chainID = big.NewInt(666)
	}
	log.Println("networkID: ", networkID)

	log.Println("constructing and sending tx..")

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
		//continue
	}

	gasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
		//continue
	}

	// construct tx
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// sign tx
	//log.Println("chainID: ", chainID)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Println("sign transaction failed")
		return err
		//continue
	}

	// send tx
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Println("send transcation failed:", err)
		return err
		//continue
	} else {
		log.Println("send transaction success")
	}

	/*
		log.Println("quering balance ...")

		//balance, _ := QueryBalance(addr, eth)
		balance, err := QueryBalance(addr)
		if err != nil {
			log.Println("query balance error")
		} else {
			log.Println("query balance success")
		}

		log.Println("balance：", balance)
		log.Println("value", value)

		log.Println(addr, "'s Balance now:", balance.String())
	*/

	// wait 200 seconds?
	/*
		fmt.Println(addr, "'s Balance now:", balance.String(), ", waiting for transfer success")
		t := 20 * (qCount + 1)
		time.Sleep(time.Duration(t) * time.Second)
	*/

	log.Println("transfer ", value.String(), "wei to", toAddr, "complete")
	return nil
}
