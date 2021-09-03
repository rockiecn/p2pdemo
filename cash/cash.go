// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cash

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Cheque is an auto generated low-level Go binding around an user-defined struct.
type Cheque struct {
	Value        *big.Int
	TokenAddr    common.Address
	Nonce        *big.Int
	FromAddr     common.Address
	ToAddr       common.Address
	OpAddr       common.Address
	ContractAddr common.Address
}

// PayCheque is an auto generated low-level Go binding around an user-defined struct.
type PayCheque struct {
	Cheque    Cheque
	ChequeSig []byte
	PayValue  *big.Int
}

// CashMetaData contains all meta data concerning the Cash contract.
var CashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Show\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"Showbytes\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"internalType\":\"structCheque\",\"name\":\"cheque\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"chequeSig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"payValue\",\"type\":\"uint256\"}],\"internalType\":\"structPayCheque\",\"name\":\"paycheque\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"paychequeSig\",\"type\":\"bytes\"}],\"name\":\"apply_cheque\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"get_node_nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"toBytes1\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061157a806100536000396000f3fe60806040526004361061004e5760003560e01c80630ac298dc146100935780631d728a0f146100be5780633cb754d8146100fb5780635f510b7b1461013857806393b8fec4146101755761008e565b3661008e577f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f885258743334604051610084929190610f25565b60405180910390a1005b600080fd5b34801561009f57600080fd5b506100a86101a5565b6040516100b59190610f0a565b60405180910390f35b3480156100ca57600080fd5b506100e560048036038101906100e09190610ba8565b6101ce565b6040516100f291906110d0565b60405180910390f35b34801561010757600080fd5b50610122600480360381019061011d9190610c3d565b6101e6565b60405161012f9190610fae565b60405180910390f35b34801561014457600080fd5b5061015f600480360381019061015a9190610ba8565b610269565b60405161016c91906110d0565b60405180910390f35b61018f600480360381019061018a9190610bd1565b6102b2565b60405161019c9190610f4e565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60016020528060005260406000206000915090505481565b6060602067ffffffffffffffff811115610229577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f19166020018201604052801561025b5781602001600182028036833780820191505090505b509050816020820152919050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60006001600084600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548360000151604001511015610346576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033d90610fd0565b60405180910390fd5b82600001516000015183604001511115610395576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038c90611030565b60405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff16836000015160c0015173ffffffffffffffffffffffffffffffffffffffff161461040b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040290611010565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1683600001516080015173ffffffffffffffffffffffffffffffffffffffff1614610481576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610478906110b0565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16836000015160a0015173ffffffffffffffffffffffffffffffffffffffff1614610517576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050e90611070565b60405180910390fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90506000819050600061054c826101e6565b90507f5f069990bf41db1b954531895361939c97a4102bcdec49ac76981d486c9e4f1d8260405161057d91906110d0565b60405180910390a17f4daa2d63a96f5e545f8ff5b850799b444e2e8a26e69f59425ba735fa0528a268816040516105b49190610fae565b60405180910390a18186600001516040015110610606576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105fd90610ff0565b60405180910390fd5b60008660000151600001518760000151602001518860000151604001518960000151606001518a60000151608001518b6000015160a001518c6000015160c0015160405160200161065d9796959493929190610e89565b6040516020818303038152906040529050600081805190602001209050600061068a828a602001516108ce565b90508073ffffffffffffffffffffffffffffffffffffffff16896000015160a0015173ffffffffffffffffffffffffffffffffffffffff1614610702576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f990611090565b60405180910390fd5b6000838a6040015160405160200161071b929190610e61565b60405160208183030381529060405290506000818051906020012090506000610744828c6108ce565b90508073ffffffffffffffffffffffffffffffffffffffff168c600001516060015173ffffffffffffffffffffffffffffffffffffffff16146107bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b390611050565b60405180910390fd5b8b600001516080015173ffffffffffffffffffffffffffffffffffffffff166108fc8d604001519081150290604051600060405180830381858888f1935050505015801561080e573d6000803e3d6000fd5b507f737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc4898c60000151608001518d6040015160405161084c929190610f25565b60405180910390a161087060018d60000151604001516109d690919063ffffffff16565b600160008e600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001995050505050505050505092915050565b600060418251146108e257600090506109d0565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c111561093657600093505050506109d0565b601b8160ff16101561095257601b8161094f91906111cf565b90505b601b8160ff161415801561096a5750601c8160ff1614155b1561097b57600093505050506109d0565b6001868285856040516000815260200160405260405161099e9493929190610f69565b6020604051602081039080840390855afa1580156109c0573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b600081836109e49190611179565b905092915050565b60006109ff6109fa84611110565b6110eb565b905082815260208101848484011115610a1757600080fd5b610a22848285611265565b509392505050565b600081359050610a3981611516565b92915050565b600082601f830112610a5057600080fd5b8135610a608482602086016109ec565b91505092915050565b600060e08284031215610a7b57600080fd5b610a8560e06110eb565b90506000610a9584828501610b93565b6000830152506020610aa984828501610a2a565b6020830152506040610abd84828501610b93565b6040830152506060610ad184828501610a2a565b6060830152506080610ae584828501610a2a565b60808301525060a0610af984828501610a2a565b60a08301525060c0610b0d84828501610a2a565b60c08301525092915050565b60006101208284031215610b2c57600080fd5b610b3660606110eb565b90506000610b4684828501610a69565b60008301525060e082013567ffffffffffffffff811115610b6657600080fd5b610b7284828501610a3f565b602083015250610100610b8784828501610b93565b60408301525092915050565b600081359050610ba28161152d565b92915050565b600060208284031215610bba57600080fd5b6000610bc884828501610a2a565b91505092915050565b60008060408385031215610be457600080fd5b600083013567ffffffffffffffff811115610bfe57600080fd5b610c0a85828601610b19565b925050602083013567ffffffffffffffff811115610c2757600080fd5b610c3385828601610a3f565b9150509250929050565b600060208284031215610c4f57600080fd5b6000610c5d84828501610b93565b91505092915050565b610c6f81611206565b82525050565b610c86610c8182611206565b6112d8565b82525050565b610c9581611218565b82525050565b610ca481611224565b82525050565b6000610cb582611141565b610cbf818561114c565b9350610ccf818560208601611274565b610cd881611364565b840191505092915050565b6000610cee82611141565b610cf8818561115d565b9350610d08818560208601611274565b80840191505092915050565b6000610d21601483611168565b9150610d2c82611382565b602082019050919050565b6000610d44601883611168565b9150610d4f826113ab565b602082019050919050565b6000610d67601683611168565b9150610d72826113d4565b602082019050919050565b6000610d8a602b83611168565b9150610d95826113fd565b604082019050919050565b6000610dad601583611168565b9150610db88261144c565b602082019050919050565b6000610dd0602983611168565b9150610ddb82611475565b604082019050919050565b6000610df3601283611168565b9150610dfe826114c4565b602082019050919050565b6000610e16601f83611168565b9150610e21826114ed565b602082019050919050565b610e358161124e565b82525050565b610e4c610e478261124e565b6112fc565b82525050565b610e5b81611258565b82525050565b6000610e6d8285610ce3565b9150610e798284610e3b565b6020820191508190509392505050565b6000610e95828a610e3b565b602082019150610ea58289610c75565b601482019150610eb58288610e3b565b602082019150610ec58287610c75565b601482019150610ed58286610c75565b601482019150610ee58285610c75565b601482019150610ef58284610c75565b60148201915081905098975050505050505050565b6000602082019050610f1f6000830184610c66565b92915050565b6000604082019050610f3a6000830185610c66565b610f476020830184610e2c565b9392505050565b6000602082019050610f636000830184610c8c565b92915050565b6000608082019050610f7e6000830187610c9b565b610f8b6020830186610e52565b610f986040830185610c9b565b610fa56060830184610c9b565b95945050505050565b60006020820190508181036000830152610fc88184610caa565b905092915050565b60006020820190508181036000830152610fe981610d14565b9050919050565b6000602082019050818103600083015261100981610d37565b9050919050565b6000602082019050818103600083015261102981610d5a565b9050919050565b6000602082019050818103600083015261104981610d7d565b9050919050565b6000602082019050818103600083015261106981610da0565b9050919050565b6000602082019050818103600083015261108981610dc3565b9050919050565b600060208201905081810360008301526110a981610de6565b9050919050565b600060208201905081810360008301526110c981610e09565b9050919050565b60006020820190506110e56000830184610e2c565b92915050565b60006110f5611106565b905061110182826112a7565b919050565b6000604051905090565b600067ffffffffffffffff82111561112b5761112a611335565b5b61113482611364565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006111848261124e565b915061118f8361124e565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156111c4576111c3611306565b5b828201905092915050565b60006111da82611258565b91506111e583611258565b92508260ff038211156111fb576111fa611306565b5b828201905092915050565b60006112118261122e565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015611292578082015181840152602081019050611277565b838111156112a1576000848401525b50505050565b6112b082611364565b810181811067ffffffffffffffff821117156112cf576112ce611335565b5b80604052505050565b60006112e3826112ea565b9050919050565b60006112f582611375565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f6368657175652e6e6f6e636520746f6f206f6c64000000000000000000000000600082015250565b7f6e6f6e6365206d757374206c657373207468616e206d61780000000000000000600082015250565b7f636f6e74726163742061646472657373206572726f7200000000000000000000600082015250565b7f70617976616c75652073686f756c64206e6f74206578636565642076616c756560008201527f206f66206368657175652e000000000000000000000000000000000000000000602082015250565b7f696c6c6567616c20706179636865717565207369670000000000000000000000600082015250565b7f6f70657261746f722073686f756c64206265206f776e6572206f66207468697360008201527f20636f6e74726163740000000000000000000000000000000000000000000000602082015250565b7f696c6c6567616c20636865717565207369670000000000000000000000000000600082015250565b7f63616c6c6572207368756f756c64206265206368657175652e746f4164647200600082015250565b61151f81611206565b811461152a57600080fd5b50565b6115368161124e565b811461154157600080fd5b5056fea2646970667358221220e49336c5bfa749ef13b0a5ed81693414a550f41980ffa77e672a966e2434163164736f6c63430008020033",
}

// CashABI is the input ABI used to generate the binding from.
// Deprecated: Use CashMetaData.ABI instead.
var CashABI = CashMetaData.ABI

// CashBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CashMetaData.Bin instead.
var CashBin = CashMetaData.Bin

// DeployCash deploys a new Ethereum contract, binding an instance of Cash to it.
func DeployCash(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cash, error) {
	parsed, err := CashMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CashBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cash{CashCaller: CashCaller{contract: contract}, CashTransactor: CashTransactor{contract: contract}, CashFilterer: CashFilterer{contract: contract}}, nil
}

// Cash is an auto generated Go binding around an Ethereum contract.
type Cash struct {
	CashCaller     // Read-only binding to the contract
	CashTransactor // Write-only binding to the contract
	CashFilterer   // Log filterer for contract events
}

// CashCaller is an auto generated read-only Go binding around an Ethereum contract.
type CashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CashSession struct {
	Contract     *Cash             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CashCallerSession struct {
	Contract *CashCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CashTransactorSession struct {
	Contract     *CashTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CashRaw is an auto generated low-level Go binding around an Ethereum contract.
type CashRaw struct {
	Contract *Cash // Generic contract binding to access the raw methods on
}

// CashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CashCallerRaw struct {
	Contract *CashCaller // Generic read-only contract binding to access the raw methods on
}

// CashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CashTransactorRaw struct {
	Contract *CashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCash creates a new instance of Cash, bound to a specific deployed contract.
func NewCash(address common.Address, backend bind.ContractBackend) (*Cash, error) {
	contract, err := bindCash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cash{CashCaller: CashCaller{contract: contract}, CashTransactor: CashTransactor{contract: contract}, CashFilterer: CashFilterer{contract: contract}}, nil
}

// NewCashCaller creates a new read-only instance of Cash, bound to a specific deployed contract.
func NewCashCaller(address common.Address, caller bind.ContractCaller) (*CashCaller, error) {
	contract, err := bindCash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CashCaller{contract: contract}, nil
}

// NewCashTransactor creates a new write-only instance of Cash, bound to a specific deployed contract.
func NewCashTransactor(address common.Address, transactor bind.ContractTransactor) (*CashTransactor, error) {
	contract, err := bindCash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CashTransactor{contract: contract}, nil
}

// NewCashFilterer creates a new log filterer instance of Cash, bound to a specific deployed contract.
func NewCashFilterer(address common.Address, filterer bind.ContractFilterer) (*CashFilterer, error) {
	contract, err := bindCash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CashFilterer{contract: contract}, nil
}

// bindCash binds a generic wrapper to an already deployed contract.
func bindCash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cash *CashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cash.Contract.CashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cash *CashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cash.Contract.CashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cash *CashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cash.Contract.CashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cash *CashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cash *CashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cash *CashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cash.Contract.contract.Transact(opts, method, params...)
}

// GetNodeNonce is a free data retrieval call binding the contract method 0x5f510b7b.
//
// Solidity: function get_node_nonce(address node) view returns(uint256)
func (_Cash *CashCaller) GetNodeNonce(opts *bind.CallOpts, node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "get_node_nonce", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeNonce is a free data retrieval call binding the contract method 0x5f510b7b.
//
// Solidity: function get_node_nonce(address node) view returns(uint256)
func (_Cash *CashSession) GetNodeNonce(node common.Address) (*big.Int, error) {
	return _Cash.Contract.GetNodeNonce(&_Cash.CallOpts, node)
}

// GetNodeNonce is a free data retrieval call binding the contract method 0x5f510b7b.
//
// Solidity: function get_node_nonce(address node) view returns(uint256)
func (_Cash *CashCallerSession) GetNodeNonce(node common.Address) (*big.Int, error) {
	return _Cash.Contract.GetNodeNonce(&_Cash.CallOpts, node)
}

// GetOwner is a free data retrieval call binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() view returns(address)
func (_Cash *CashCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "get_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() view returns(address)
func (_Cash *CashSession) GetOwner() (common.Address, error) {
	return _Cash.Contract.GetOwner(&_Cash.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() view returns(address)
func (_Cash *CashCallerSession) GetOwner() (common.Address, error) {
	return _Cash.Contract.GetOwner(&_Cash.CallOpts)
}

// NodeNonce is a free data retrieval call binding the contract method 0x1d728a0f.
//
// Solidity: function nodeNonce(address ) view returns(uint256)
func (_Cash *CashCaller) NodeNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "nodeNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeNonce is a free data retrieval call binding the contract method 0x1d728a0f.
//
// Solidity: function nodeNonce(address ) view returns(uint256)
func (_Cash *CashSession) NodeNonce(arg0 common.Address) (*big.Int, error) {
	return _Cash.Contract.NodeNonce(&_Cash.CallOpts, arg0)
}

// NodeNonce is a free data retrieval call binding the contract method 0x1d728a0f.
//
// Solidity: function nodeNonce(address ) view returns(uint256)
func (_Cash *CashCallerSession) NodeNonce(arg0 common.Address) (*big.Int, error) {
	return _Cash.Contract.NodeNonce(&_Cash.CallOpts, arg0)
}

// ToBytes1 is a free data retrieval call binding the contract method 0x3cb754d8.
//
// Solidity: function toBytes1(uint256 x) pure returns(bytes b)
func (_Cash *CashCaller) ToBytes1(opts *bind.CallOpts, x *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "toBytes1", x)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ToBytes1 is a free data retrieval call binding the contract method 0x3cb754d8.
//
// Solidity: function toBytes1(uint256 x) pure returns(bytes b)
func (_Cash *CashSession) ToBytes1(x *big.Int) ([]byte, error) {
	return _Cash.Contract.ToBytes1(&_Cash.CallOpts, x)
}

// ToBytes1 is a free data retrieval call binding the contract method 0x3cb754d8.
//
// Solidity: function toBytes1(uint256 x) pure returns(bytes b)
func (_Cash *CashCallerSession) ToBytes1(x *big.Int) ([]byte, error) {
	return _Cash.Contract.ToBytes1(&_Cash.CallOpts, x)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0x93b8fec4.
//
// Solidity: function apply_cheque(((uint256,address,uint256,address,address,address,address),bytes,uint256) paycheque, bytes paychequeSig) payable returns(bool)
func (_Cash *CashTransactor) ApplyCheque(opts *bind.TransactOpts, paycheque PayCheque, paychequeSig []byte) (*types.Transaction, error) {
	return _Cash.contract.Transact(opts, "apply_cheque", paycheque, paychequeSig)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0x93b8fec4.
//
// Solidity: function apply_cheque(((uint256,address,uint256,address,address,address,address),bytes,uint256) paycheque, bytes paychequeSig) payable returns(bool)
func (_Cash *CashSession) ApplyCheque(paycheque PayCheque, paychequeSig []byte) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheque(&_Cash.TransactOpts, paycheque, paychequeSig)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0x93b8fec4.
//
// Solidity: function apply_cheque(((uint256,address,uint256,address,address,address,address),bytes,uint256) paycheque, bytes paychequeSig) payable returns(bool)
func (_Cash *CashTransactorSession) ApplyCheque(paycheque PayCheque, paychequeSig []byte) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheque(&_Cash.TransactOpts, paycheque, paychequeSig)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cash *CashTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cash.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cash *CashSession) Receive() (*types.Transaction, error) {
	return _Cash.Contract.Receive(&_Cash.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cash *CashTransactorSession) Receive() (*types.Transaction, error) {
	return _Cash.Contract.Receive(&_Cash.TransactOpts)
}

// CashPaidIterator is returned from FilterPaid and is used to iterate over the raw logs and unpacked data for Paid events raised by the Cash contract.
type CashPaidIterator struct {
	Event *CashPaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CashPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashPaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CashPaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CashPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashPaid represents a Paid event raised by the Cash contract.
type CashPaid struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPaid is a free log retrieval operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: event Paid(address arg0, uint256 arg1)
func (_Cash *CashFilterer) FilterPaid(opts *bind.FilterOpts) (*CashPaidIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "Paid")
	if err != nil {
		return nil, err
	}
	return &CashPaidIterator{contract: _Cash.contract, event: "Paid", logs: logs, sub: sub}, nil
}

// WatchPaid is a free log subscription operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: event Paid(address arg0, uint256 arg1)
func (_Cash *CashFilterer) WatchPaid(opts *bind.WatchOpts, sink chan<- *CashPaid) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "Paid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashPaid)
				if err := _Cash.contract.UnpackLog(event, "Paid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaid is a log parse operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: event Paid(address arg0, uint256 arg1)
func (_Cash *CashFilterer) ParsePaid(log types.Log) (*CashPaid, error) {
	event := new(CashPaid)
	if err := _Cash.contract.UnpackLog(event, "Paid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Cash contract.
type CashReceivedIterator struct {
	Event *CashReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CashReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CashReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CashReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashReceived represents a Received event raised by the Cash contract.
type CashReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Cash *CashFilterer) FilterReceived(opts *bind.FilterOpts) (*CashReceivedIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &CashReceivedIterator{contract: _Cash.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Cash *CashFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *CashReceived) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashReceived)
				if err := _Cash.contract.UnpackLog(event, "Received", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Cash *CashFilterer) ParseReceived(log types.Log) (*CashReceived, error) {
	event := new(CashReceived)
	if err := _Cash.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowIterator is returned from FilterShow and is used to iterate over the raw logs and unpacked data for Show events raised by the Cash contract.
type CashShowIterator struct {
	Event *CashShow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CashShowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CashShow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CashShowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShow represents a Show event raised by the Cash contract.
type CashShow struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShow is a free log retrieval operation binding the contract event 0x5f069990bf41db1b954531895361939c97a4102bcdec49ac76981d486c9e4f1d.
//
// Solidity: event Show(uint256 arg0)
func (_Cash *CashFilterer) FilterShow(opts *bind.FilterOpts) (*CashShowIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "Show")
	if err != nil {
		return nil, err
	}
	return &CashShowIterator{contract: _Cash.contract, event: "Show", logs: logs, sub: sub}, nil
}

// WatchShow is a free log subscription operation binding the contract event 0x5f069990bf41db1b954531895361939c97a4102bcdec49ac76981d486c9e4f1d.
//
// Solidity: event Show(uint256 arg0)
func (_Cash *CashFilterer) WatchShow(opts *bind.WatchOpts, sink chan<- *CashShow) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "Show")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShow)
				if err := _Cash.contract.UnpackLog(event, "Show", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseShow is a log parse operation binding the contract event 0x5f069990bf41db1b954531895361939c97a4102bcdec49ac76981d486c9e4f1d.
//
// Solidity: event Show(uint256 arg0)
func (_Cash *CashFilterer) ParseShow(log types.Log) (*CashShow, error) {
	event := new(CashShow)
	if err := _Cash.contract.UnpackLog(event, "Show", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowbytesIterator is returned from FilterShowbytes and is used to iterate over the raw logs and unpacked data for Showbytes events raised by the Cash contract.
type CashShowbytesIterator struct {
	Event *CashShowbytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CashShowbytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowbytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CashShowbytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CashShowbytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowbytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowbytes represents a Showbytes event raised by the Cash contract.
type CashShowbytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowbytes is a free log retrieval operation binding the contract event 0x4daa2d63a96f5e545f8ff5b850799b444e2e8a26e69f59425ba735fa0528a268.
//
// Solidity: event Showbytes(bytes arg0)
func (_Cash *CashFilterer) FilterShowbytes(opts *bind.FilterOpts) (*CashShowbytesIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "Showbytes")
	if err != nil {
		return nil, err
	}
	return &CashShowbytesIterator{contract: _Cash.contract, event: "Showbytes", logs: logs, sub: sub}, nil
}

// WatchShowbytes is a free log subscription operation binding the contract event 0x4daa2d63a96f5e545f8ff5b850799b444e2e8a26e69f59425ba735fa0528a268.
//
// Solidity: event Showbytes(bytes arg0)
func (_Cash *CashFilterer) WatchShowbytes(opts *bind.WatchOpts, sink chan<- *CashShowbytes) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "Showbytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowbytes)
				if err := _Cash.contract.UnpackLog(event, "Showbytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseShowbytes is a log parse operation binding the contract event 0x4daa2d63a96f5e545f8ff5b850799b444e2e8a26e69f59425ba735fa0528a268.
//
// Solidity: event Showbytes(bytes arg0)
func (_Cash *CashFilterer) ParseShowbytes(log types.Log) (*CashShowbytes, error) {
	event := new(CashShowbytes)
	if err := _Cash.contract.UnpackLog(event, "Showbytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
