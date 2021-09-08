package operator

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/protobuf/proto"

	"github.com/rockiecn/p2pdemo/cash"
	"github.com/rockiecn/p2pdemo/clientops"
	"github.com/rockiecn/p2pdemo/db"
	"github.com/rockiecn/p2pdemo/hostops"
	"github.com/rockiecn/p2pdemo/pb"
	"github.com/rockiecn/p2pdemo/print"
	"github.com/rockiecn/p2pdemo/sigapi"
	"github.com/rockiecn/p2pdemo/utils"
)

type Operator struct {
	OpDB *leveldb.DB // provider -> nonce

	DBfile string

	OperatorAddr string // "5B38Da6a701c568545dCfcB03FcB875f56beddC4"
	OperatorSK   string // "503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb"

	TokenAddr string // token address
	FromAddr  string // user
	ToAddr    string // storage

	db.DB // interface DB
}

// init operator, need db open first
func (op *Operator) Init() error {

	op.DBfile = "./operator.db"
	op.OpenDB()

	//op.ContractAddress = ""
	op.OperatorAddr = "5B38Da6a701c568545dCfcB03FcB875f56beddC4"
	op.OperatorSK = "503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb"

	op.TokenAddr = "b213d01542d129806d664248a380db8b12059061" // token address
	op.FromAddr = "Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"  // user
	op.ToAddr = "4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"    // storage

	byteAddr, err := op.OpDB.Get([]byte("contractAddr"), nil)
	if err != nil {
		return errors.New("operator init: read contract address failed")
	}
	print.ContractAddress = string(byteAddr)

	return nil
}

//
func (op *Operator) OpenDB() error {
	db, err := leveldb.OpenFile(op.DBfile, nil)
	if err != nil {
		fmt.Println("open db error: ", err)
		return err
	}
	op.OpDB = db

	return nil
}

// close db
func (op *Operator) CloseDB() error {
	op.OpDB.Close()

	return nil
}

// operator deploy contract
func (op *Operator) DeployContract() (common.Address, error) {

	var contractAddr common.Address

	fmt.Println("HOST: ", hostops.HOST)
	client, err := clientops.GetClient(hostops.HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return contractAddr, err
	}
	defer client.Close()

	// get sk
	sk, err := crypto.HexToECDSA(op.OperatorSK)
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

	fmt.Println("auth success")

	contractAddr, _, _, err = cash.DeployCash(auth, client)
	if err != nil {
		fmt.Println("deployCashErr:", err)
		return contractAddr, err
	}
	fmt.Println("contractAddr:", contractAddr.String())
	fmt.Println("value:", auth.Value.String())

	// ====== store contractAddr to db
	// store cash address
	err = op.OpDB.Put([]byte("contractAddr"), []byte(contractAddr.String()), nil)
	if err != nil {
		fmt.Println("db put data error:", err)
		return contractAddr, err
	}

	print.ContractAddress = contractAddr.String()

	return contractAddr, nil
}

// get current provider nonce stored in db
func (op *Operator) GetNonceLocal(nodeAddr string) (*big.Int, error) {

	//byteNonce := utils.Int64ToBytes(0)
	bigNonce := big.NewInt(0)

	// storage -> nonce
	byteNonce, err := op.OpDB.Get([]byte(nodeAddr), nil)
	if err != nil {
		fmt.Println("get nonce error: ", err)
		return nil, err
	}
	bigNonce.SetBytes(byteNonce)

	return bigNonce, nil
}

// set a provider's nonce in db
func (op *Operator) SetNonceLocal(nodeAddr string, nonce *big.Int) error {
	op.OpenDB()
	defer op.CloseDB()

	// storage -> nonce
	err := op.OpDB.Put([]byte(nodeAddr), nonce.Bytes(), nil)
	if err != nil {
		fmt.Println("reset nonce error: ", err)
		return err
	}

	return nil
}

// call get contract node nonce
func (op *Operator) GetNonceContract(node common.Address) (*big.Int, error) {
	cli, err := clientops.GetClient(hostops.HOST)
	if err != nil {
		fmt.Println("failed to dial geth", err)
		return nil, err
	}
	defer cli.Close()

	// ====== get contractAddr from db

	// store cash address
	byteContractAddr, err := op.OpDB.Get([]byte("contractAddr"), nil)
	if err != nil {
		fmt.Println("db get data error:", err)
		return nil, err
	}

	AddressContract := common.HexToAddress(string(byteContractAddr))
	// get contract instance from address
	cashInstance, err2 := cash.NewCash(AddressContract, cli)
	if err2 != nil {
		fmt.Println("NewCash err: ", err2)
		return nil, err2
	}

	bigNonce, err := cashInstance.GetNodeNonce(nil, node)
	if err != nil {
		fmt.Println("tx failed :", err)
		return nil, err
	}

	fmt.Println()
	fmt.Println("Node: ", node.String())
	fmt.Println("Node nonce: ", bigNonce.String())

	return bigNonce, nil
}

// operator generate a cheque
func (op *Operator) GenerateCheque() (*pb.Cheque, error) {

	fmt.Println("--> Construct and send Cheque...")

	// construct Cheque

	cheque := pb.Cheque{}

	cheque.Value = "100000000000000000000"   // Cheque 100
	cheque.TokenAddress = op.TokenAddr       // token address
	cheque.From = op.FromAddr                // user
	cheque.To = op.ToAddr                    // storage
	cheque.OperatorAddress = op.OperatorAddr // operator

	// storage -> nonce
	nonce, err := op.OpDB.Get([]byte(cheque.To), nil)
	if err != nil {
		if err.Error() == "leveldb: not found" { // no nonce at all
			op.OpDB.Put([]byte(cheque.To), utils.Int64ToBytes(0), nil)
		} else {
			fmt.Println("operator db get nonce error: ", err)
			return nil, err
		}
	}

	// increase nonce by 1
	// byte to string
	bigOne := big.NewInt(1)
	bigNonce := big.NewInt(0)
	bigNonce = bigNonce.SetBytes(nonce)
	fmt.Println("bigNonce: ", bigNonce.String())
	bigNonce = bigNonce.Add(bigNonce, bigOne)
	fmt.Println("bigNonce: ", bigNonce.String())

	// record nonce into db
	err = op.OpDB.Put([]byte(cheque.To), bigNonce.Bytes(), nil)
	if err != nil {
		fmt.Println("operator db put nonce error")
		return nil, err
	}

	//
	cheque.Nonce = bigNonce.String()

	contractAddrByte, err := op.OpDB.Get([]byte("contractAddr"), nil)
	if err != nil {
		log.Println("!! get cash address error:", err)
		return nil, err
	}

	// contract address, delete prefix '0x'
	contractAddrByte = contractAddrByte[2:]

	cheque.ContractAddress = string(contractAddrByte) // contract address

	// if global.DEBUG {
	// 	print.Printf100ms("sigByte:%x\n", sigByte)
	// 	print.Printf100ms("contractAddr:%s\n", cheque.ContractAddress)
	// 	print.Printf100ms("cheque:%x\n", ChequeMarshaled)
	// }

	return &cheque, nil
}

// operator send cheque to user
func (op *Operator) SendCheque(s network.Stream, cheque *pb.Cheque) error {
	// serialize
	ChequeMarshaled, err := proto.Marshal(cheque)
	if err != nil {
		fmt.Println("Failed to encode PayCheque:", err)
		return err
	}

	// construct Cheque message: sig(65 bytes) | data
	fmt.Println("-> constructing msg")

	// calc cheque hash
	hash := utils.CalcChequeHash(cheque)

	fmt.Printf("Cheque send, hash: %x\n", hash)

	// sign Cheque by operator
	sigByte, err := sigapi.Sign(hash, []byte(op.OperatorSK))
	if err != nil {
		log.Fatal("sign error:", err)
	}

	// sig(65) | cheque
	var msg = []byte{}
	msg = utils.MergeSlice(sigByte, ChequeMarshaled)

	fmt.Println("-> sending msg")
	// send msg
	_, err = s.Write([]byte(msg))
	if err != nil {
		panic("stream write error")
	}

	return nil
}

// called when an user connecting, generate and send a cheque to user
func (op *Operator) GenAndSendCheque(s network.Stream) error {

	cheque, err := op.GenerateCheque()
	if err != nil {
		fmt.Println("generate cheque error:", err)
		return err
	}

	err = op.SendCheque(s, cheque)
	if err != nil {
		fmt.Println("send cheque error:", err)
		return err
	}

	fmt.Println("\n> Intput target cmd: ")

	return nil
}
