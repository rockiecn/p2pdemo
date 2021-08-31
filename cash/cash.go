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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"internalType\":\"structCheque\",\"name\":\"cheque\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"chequeSig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"payValue\",\"type\":\"uint256\"}],\"internalType\":\"structPayCheque\",\"name\":\"paycheque\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"paychequeSig\",\"type\":\"bytes\"}],\"name\":\"apply_cheque\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"get_node_nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061130e806100536000396000f3fe6080604052600436106100435760003560e01c80630ac298dc146100885780631d728a0f146100b35780635f510b7b146100f057806393b8fec41461012d57610083565b36610083577f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f885258743334604051610079929190610cdb565b60405180910390a1005b600080fd5b34801561009457600080fd5b5061009d61015d565b6040516100aa9190610cc0565b60405180910390f35b3480156100bf57600080fd5b506100da60048036038101906100d591906109e3565b610186565b6040516100e79190610e44565b60405180910390f35b3480156100fc57600080fd5b50610117600480360381019061011291906109e3565b61019e565b6040516101249190610e44565b60405180910390f35b61014760048036038101906101429190610a0c565b6101e7565b6040516101549190610d04565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60016020528060005260406000206000915090505481565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60006001600084600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054836000015160400151101561027b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027290610d64565b60405180910390fd5b826000015160000151836040015111156102ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c190610da4565b60405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff16836000015160c0015173ffffffffffffffffffffffffffffffffffffffff1614610340576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033790610d84565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1683600001516080015173ffffffffffffffffffffffffffffffffffffffff16146103b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ad90610e24565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16836000015160a0015173ffffffffffffffffffffffffffffffffffffffff161461044c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044390610de4565b60405180910390fd5b6000836000015160000151846000015160200151856000015160400151866000015160600151876000015160800151886000015160a00151896000015160c001516040516020016104a39796959493929190610c3f565b604051602081830303815290604052905060008185604001516040516020016104cd929190610c17565b6040516020818303038152906040529050600082805190602001209050600082805190602001209050600061050683896020015161071f565b90506000610514838961071f565b90508173ffffffffffffffffffffffffffffffffffffffff16896000015160a0015173ffffffffffffffffffffffffffffffffffffffff161461058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058390610e04565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1689600001516060015173ffffffffffffffffffffffffffffffffffffffff1614610602576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f990610dc4565b60405180910390fd5b6000670de0b6b3a76400008a6040015161061c9190610f69565b905089600001516080015173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561066c573d6000803e3d6000fd5b507f737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc4898a6000015160800151826040516106a6929190610cdb565b60405180910390a160018a60000151604001516106c39190610edc565b600160008c600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600197505050505050505092915050565b600060418251146107335760009050610821565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c11156107875760009350505050610821565b601b8160ff1610156107a357601b816107a09190610f32565b90505b601b8160ff16141580156107bb5750601c8160ff1614155b156107cc5760009350505050610821565b600186828585604051600081526020016040526040516107ef9493929190610d1f565b6020604051602081039080840390855afa158015610811573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b600061083a61083584610e84565b610e5f565b90508281526020810184848401111561085257600080fd5b61085d848285611022565b509392505050565b600081359050610874816112aa565b92915050565b600082601f83011261088b57600080fd5b813561089b848260208601610827565b91505092915050565b600060e082840312156108b657600080fd5b6108c060e0610e5f565b905060006108d0848285016109ce565b60008301525060206108e484828501610865565b60208301525060406108f8848285016109ce565b604083015250606061090c84828501610865565b606083015250608061092084828501610865565b60808301525060a061093484828501610865565b60a08301525060c061094884828501610865565b60c08301525092915050565b6000610120828403121561096757600080fd5b6109716060610e5f565b90506000610981848285016108a4565b60008301525060e082013567ffffffffffffffff8111156109a157600080fd5b6109ad8482850161087a565b6020830152506101006109c2848285016109ce565b60408301525092915050565b6000813590506109dd816112c1565b92915050565b6000602082840312156109f557600080fd5b6000610a0384828501610865565b91505092915050565b60008060408385031215610a1f57600080fd5b600083013567ffffffffffffffff811115610a3957600080fd5b610a4585828601610954565b925050602083013567ffffffffffffffff811115610a6257600080fd5b610a6e8582860161087a565b9150509250929050565b610a8181610fc3565b82525050565b610a98610a9382610fc3565b611095565b82525050565b610aa781610fd5565b82525050565b610ab681610fe1565b82525050565b6000610ac782610eb5565b610ad18185610ec0565b9350610ae1818560208601611031565b80840191505092915050565b6000610afa601483610ecb565b9150610b058261113f565b602082019050919050565b6000610b1d601683610ecb565b9150610b2882611168565b602082019050919050565b6000610b40602b83610ecb565b9150610b4b82611191565b604082019050919050565b6000610b63601583610ecb565b9150610b6e826111e0565b602082019050919050565b6000610b86602983610ecb565b9150610b9182611209565b604082019050919050565b6000610ba9601283610ecb565b9150610bb482611258565b602082019050919050565b6000610bcc601f83610ecb565b9150610bd782611281565b602082019050919050565b610beb8161100b565b82525050565b610c02610bfd8261100b565b6110b9565b82525050565b610c1181611015565b82525050565b6000610c238285610abc565b9150610c2f8284610bf1565b6020820191508190509392505050565b6000610c4b828a610bf1565b602082019150610c5b8289610a87565b601482019150610c6b8288610bf1565b602082019150610c7b8287610a87565b601482019150610c8b8286610a87565b601482019150610c9b8285610a87565b601482019150610cab8284610a87565b60148201915081905098975050505050505050565b6000602082019050610cd56000830184610a78565b92915050565b6000604082019050610cf06000830185610a78565b610cfd6020830184610be2565b9392505050565b6000602082019050610d196000830184610a9e565b92915050565b6000608082019050610d346000830187610aad565b610d416020830186610c08565b610d4e6040830185610aad565b610d5b6060830184610aad565b95945050505050565b60006020820190508181036000830152610d7d81610aed565b9050919050565b60006020820190508181036000830152610d9d81610b10565b9050919050565b60006020820190508181036000830152610dbd81610b33565b9050919050565b60006020820190508181036000830152610ddd81610b56565b9050919050565b60006020820190508181036000830152610dfd81610b79565b9050919050565b60006020820190508181036000830152610e1d81610b9c565b9050919050565b60006020820190508181036000830152610e3d81610bbf565b9050919050565b6000602082019050610e596000830184610be2565b92915050565b6000610e69610e7a565b9050610e758282611064565b919050565b6000604051905090565b600067ffffffffffffffff821115610e9f57610e9e6110f2565b5b610ea882611121565b9050602081019050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b6000610ee78261100b565b9150610ef28361100b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610f2757610f266110c3565b5b828201905092915050565b6000610f3d82611015565b9150610f4883611015565b92508260ff03821115610f5e57610f5d6110c3565b5b828201905092915050565b6000610f748261100b565b9150610f7f8361100b565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610fb857610fb76110c3565b5b828202905092915050565b6000610fce82610feb565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561104f578082015181840152602081019050611034565b8381111561105e576000848401525b50505050565b61106d82611121565b810181811067ffffffffffffffff8211171561108c5761108b6110f2565b5b80604052505050565b60006110a0826110a7565b9050919050565b60006110b282611132565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f6368657175652e6e6f6e636520746f6f206f6c64000000000000000000000000600082015250565b7f636f6e74726163742061646472657373206572726f7200000000000000000000600082015250565b7f70617976616c75652073686f756c64206e6f74206578636565642076616c756560008201527f206f66206368657175652e000000000000000000000000000000000000000000602082015250565b7f696c6c6567616c20706179636865717565207369670000000000000000000000600082015250565b7f6f70657261746f722073686f756c64206265206f776e6572206f66207468697360008201527f20636f6e74726163740000000000000000000000000000000000000000000000602082015250565b7f696c6c6567616c20636865717565207369670000000000000000000000000000600082015250565b7f63616c6c6572207368756f756c64206265206368657175652e746f4164647200600082015250565b6112b381610fc3565b81146112be57600080fd5b50565b6112ca8161100b565b81146112d557600080fd5b5056fea2646970667358221220adcc06c86f24edb7296cf7f8440f1c5409a3edffa60cecc02d4e05f5d72f6a0564736f6c63430008020033",
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
