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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ShowChequeHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"ShowChequeSig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ShowChequeSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ShowFrom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ShowNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"ShowPayCheckPack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ShowPayChequeHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ShowPayChequeSigner\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"internalType\":\"structCheque\",\"name\":\"cheque\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"chequeSig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"payValue\",\"type\":\"uint256\"}],\"internalType\":\"structPayCheque\",\"name\":\"paycheque\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"paychequeSig\",\"type\":\"bytes\"}],\"name\":\"apply_cheque\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"get_node_nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052610dfa806100136000396000f3fe6080604052600436106100345760003560e01c80631d728a0f146100395780635f510b7b1461007657806393b8fec4146100b3575b600080fd5b34801561004557600080fd5b50610060600480360381019061005b91906107ea565b6100e3565b60405161006d9190610afc565b60405180910390f35b34801561008257600080fd5b5061009d600480360381019061009891906107ea565b6100fb565b6040516100aa9190610afc565b60405180910390f35b6100cd60048036038101906100c89190610813565b610143565b6040516100da9190610a5f565b60405180910390f35b60006020528060005260406000206000915090505481565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60007fe7d038274dbda90ba9b4a73bcf23a30cabd0a1567a2b29d19d2accf5399f030983600001516060015160405161017c9190610a44565b60405180910390a17faf3a06b91ffcbd286a1f609888182ab42c9bb71d1f95491bfe3bc6d300a481848360000151604001516040516101bb9190610afc565b60405180910390a17f749c515763629ce6c817505ed6443ae1cc62a6cdc59359c25057f575b418d0e183602001516040516101f69190610ada565b60405180910390a16000836000015160000151846000015160200151856000015160400151866000015160600151876000015160800151886000015160a00151896000015160c001516040516020016102559796959493929190610931565b60405160208183030381529060405290506000846000015160000151856000015160200151866000015160400151876000015160600151886000015160800151896000015160a001518a6000015160c001518b604001516040516020016102c39897969594939291906109b2565b60405160208183030381529060405290507fe254bf5fc12a7e7953bcdfd673d53edf48f7b614e7189fbfd2ddc3514d3abd48816040516103039190610ada565b60405180910390a16000828051906020012090506000828051906020012090507f58e8d67fffbff939e478cac4f59c309af09bfa219751076057be6947f5703a97826040516103529190610a7a565b60405180910390a17f267ef9c051dadf48f9d78139f8afcaab9b9ea923e2f62607fb90558d169399f0816040516103899190610a7a565b60405180910390a160006103a1838960200151610526565b90507fd992130aab86d2580aa88e21392e36739b1621554343d409620f0fb1e698b33d816040516103d29190610a44565b60405180910390a160006103e68389610526565b90507f129e8dff958b080ea26f69bc288d31037ee2127de5b7e3784e72ee7fdb8901e2816040516104179190610a44565b60405180910390a18173ffffffffffffffffffffffffffffffffffffffff16896000015160a0015173ffffffffffffffffffffffffffffffffffffffff1614801561049557508073ffffffffffffffffffffffffffffffffffffffff1689600001516060015173ffffffffffffffffffffffffffffffffffffffff16145b15610515576000670de0b6b3a76400008a604001516104b49190610bc0565b905089600001516080015173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610504573d6000803e3d6000fd5b506001975050505050505050610520565b600096505050505050505b92915050565b6000604182511461053a5760009050610628565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c111561058e5760009350505050610628565b601b8160ff1610156105aa57601b816105a79190610b89565b90505b601b8160ff16141580156105c25750601c8160ff1614155b156105d35760009350505050610628565b600186828585604051600081526020016040526040516105f69493929190610a95565b6020604051602081039080840390855afa158015610618573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b600061064161063c84610b3c565b610b17565b90508281526020810184848401111561065957600080fd5b610664848285610c79565b509392505050565b60008135905061067b81610d96565b92915050565b600082601f83011261069257600080fd5b81356106a284826020860161062e565b91505092915050565b600060e082840312156106bd57600080fd5b6106c760e0610b17565b905060006106d7848285016107d5565b60008301525060206106eb8482850161066c565b60208301525060406106ff848285016107d5565b60408301525060606107138482850161066c565b60608301525060806107278482850161066c565b60808301525060a061073b8482850161066c565b60a08301525060c061074f8482850161066c565b60c08301525092915050565b6000610120828403121561076e57600080fd5b6107786060610b17565b90506000610788848285016106ab565b60008301525060e082013567ffffffffffffffff8111156107a857600080fd5b6107b484828501610681565b6020830152506101006107c9848285016107d5565b60408301525092915050565b6000813590506107e481610dad565b92915050565b6000602082840312156107fc57600080fd5b600061080a8482850161066c565b91505092915050565b6000806040838503121561082657600080fd5b600083013567ffffffffffffffff81111561084057600080fd5b61084c8582860161075b565b925050602083013567ffffffffffffffff81111561086957600080fd5b61087585828601610681565b9150509250929050565b61088881610c1a565b82525050565b61089f61089a82610c1a565b610cec565b82525050565b6108ae81610c2c565b82525050565b6108bd81610c38565b82525050565b60006108ce82610b6d565b6108d88185610b78565b93506108e8818560208601610c88565b6108f181610d78565b840191505092915050565b61090581610c62565b82525050565b61091c61091782610c62565b610d10565b82525050565b61092b81610c6c565b82525050565b600061093d828a61090b565b60208201915061094d828961088e565b60148201915061095d828861090b565b60208201915061096d828761088e565b60148201915061097d828661088e565b60148201915061098d828561088e565b60148201915061099d828461088e565b60148201915081905098975050505050505050565b60006109be828b61090b565b6020820191506109ce828a61088e565b6014820191506109de828961090b565b6020820191506109ee828861088e565b6014820191506109fe828761088e565b601482019150610a0e828661088e565b601482019150610a1e828561088e565b601482019150610a2e828461090b565b6020820191508190509998505050505050505050565b6000602082019050610a59600083018461087f565b92915050565b6000602082019050610a7460008301846108a5565b92915050565b6000602082019050610a8f60008301846108b4565b92915050565b6000608082019050610aaa60008301876108b4565b610ab76020830186610922565b610ac460408301856108b4565b610ad160608301846108b4565b95945050505050565b60006020820190508181036000830152610af481846108c3565b905092915050565b6000602082019050610b1160008301846108fc565b92915050565b6000610b21610b32565b9050610b2d8282610cbb565b919050565b6000604051905090565b600067ffffffffffffffff821115610b5757610b56610d49565b5b610b6082610d78565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b6000610b9482610c6c565b9150610b9f83610c6c565b92508260ff03821115610bb557610bb4610d1a565b5b828201905092915050565b6000610bcb82610c62565b9150610bd683610c62565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610c0f57610c0e610d1a565b5b828202905092915050565b6000610c2582610c42565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015610ca6578082015181840152602081019050610c8b565b83811115610cb5576000848401525b50505050565b610cc482610d78565b810181811067ffffffffffffffff82111715610ce357610ce2610d49565b5b80604052505050565b6000610cf782610cfe565b9050919050565b6000610d0982610d89565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b610d9f81610c1a565b8114610daa57600080fd5b50565b610db681610c62565b8114610dc157600080fd5b5056fea26469706673582212204d15163b537f232c2f9f2368bacba92b5be3555b385a11a4059f05fd7b31fc2e64736f6c63430008020033",
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

// CashShowChequeHashIterator is returned from FilterShowChequeHash and is used to iterate over the raw logs and unpacked data for ShowChequeHash events raised by the Cash contract.
type CashShowChequeHashIterator struct {
	Event *CashShowChequeHash // Event containing the contract specifics and raw log

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
func (it *CashShowChequeHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowChequeHash)
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
		it.Event = new(CashShowChequeHash)
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
func (it *CashShowChequeHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowChequeHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowChequeHash represents a ShowChequeHash event raised by the Cash contract.
type CashShowChequeHash struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowChequeHash is a free log retrieval operation binding the contract event 0x58e8d67fffbff939e478cac4f59c309af09bfa219751076057be6947f5703a97.
//
// Solidity: event ShowChequeHash(bytes32 arg0)
func (_Cash *CashFilterer) FilterShowChequeHash(opts *bind.FilterOpts) (*CashShowChequeHashIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowChequeHash")
	if err != nil {
		return nil, err
	}
	return &CashShowChequeHashIterator{contract: _Cash.contract, event: "ShowChequeHash", logs: logs, sub: sub}, nil
}

// WatchShowChequeHash is a free log subscription operation binding the contract event 0x58e8d67fffbff939e478cac4f59c309af09bfa219751076057be6947f5703a97.
//
// Solidity: event ShowChequeHash(bytes32 arg0)
func (_Cash *CashFilterer) WatchShowChequeHash(opts *bind.WatchOpts, sink chan<- *CashShowChequeHash) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowChequeHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowChequeHash)
				if err := _Cash.contract.UnpackLog(event, "ShowChequeHash", log); err != nil {
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

// ParseShowChequeHash is a log parse operation binding the contract event 0x58e8d67fffbff939e478cac4f59c309af09bfa219751076057be6947f5703a97.
//
// Solidity: event ShowChequeHash(bytes32 arg0)
func (_Cash *CashFilterer) ParseShowChequeHash(log types.Log) (*CashShowChequeHash, error) {
	event := new(CashShowChequeHash)
	if err := _Cash.contract.UnpackLog(event, "ShowChequeHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowChequeSigIterator is returned from FilterShowChequeSig and is used to iterate over the raw logs and unpacked data for ShowChequeSig events raised by the Cash contract.
type CashShowChequeSigIterator struct {
	Event *CashShowChequeSig // Event containing the contract specifics and raw log

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
func (it *CashShowChequeSigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowChequeSig)
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
		it.Event = new(CashShowChequeSig)
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
func (it *CashShowChequeSigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowChequeSigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowChequeSig represents a ShowChequeSig event raised by the Cash contract.
type CashShowChequeSig struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowChequeSig is a free log retrieval operation binding the contract event 0x749c515763629ce6c817505ed6443ae1cc62a6cdc59359c25057f575b418d0e1.
//
// Solidity: event ShowChequeSig(bytes arg0)
func (_Cash *CashFilterer) FilterShowChequeSig(opts *bind.FilterOpts) (*CashShowChequeSigIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowChequeSig")
	if err != nil {
		return nil, err
	}
	return &CashShowChequeSigIterator{contract: _Cash.contract, event: "ShowChequeSig", logs: logs, sub: sub}, nil
}

// WatchShowChequeSig is a free log subscription operation binding the contract event 0x749c515763629ce6c817505ed6443ae1cc62a6cdc59359c25057f575b418d0e1.
//
// Solidity: event ShowChequeSig(bytes arg0)
func (_Cash *CashFilterer) WatchShowChequeSig(opts *bind.WatchOpts, sink chan<- *CashShowChequeSig) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowChequeSig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowChequeSig)
				if err := _Cash.contract.UnpackLog(event, "ShowChequeSig", log); err != nil {
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

// ParseShowChequeSig is a log parse operation binding the contract event 0x749c515763629ce6c817505ed6443ae1cc62a6cdc59359c25057f575b418d0e1.
//
// Solidity: event ShowChequeSig(bytes arg0)
func (_Cash *CashFilterer) ParseShowChequeSig(log types.Log) (*CashShowChequeSig, error) {
	event := new(CashShowChequeSig)
	if err := _Cash.contract.UnpackLog(event, "ShowChequeSig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowChequeSignerIterator is returned from FilterShowChequeSigner and is used to iterate over the raw logs and unpacked data for ShowChequeSigner events raised by the Cash contract.
type CashShowChequeSignerIterator struct {
	Event *CashShowChequeSigner // Event containing the contract specifics and raw log

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
func (it *CashShowChequeSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowChequeSigner)
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
		it.Event = new(CashShowChequeSigner)
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
func (it *CashShowChequeSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowChequeSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowChequeSigner represents a ShowChequeSigner event raised by the Cash contract.
type CashShowChequeSigner struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowChequeSigner is a free log retrieval operation binding the contract event 0xd992130aab86d2580aa88e21392e36739b1621554343d409620f0fb1e698b33d.
//
// Solidity: event ShowChequeSigner(address arg0)
func (_Cash *CashFilterer) FilterShowChequeSigner(opts *bind.FilterOpts) (*CashShowChequeSignerIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowChequeSigner")
	if err != nil {
		return nil, err
	}
	return &CashShowChequeSignerIterator{contract: _Cash.contract, event: "ShowChequeSigner", logs: logs, sub: sub}, nil
}

// WatchShowChequeSigner is a free log subscription operation binding the contract event 0xd992130aab86d2580aa88e21392e36739b1621554343d409620f0fb1e698b33d.
//
// Solidity: event ShowChequeSigner(address arg0)
func (_Cash *CashFilterer) WatchShowChequeSigner(opts *bind.WatchOpts, sink chan<- *CashShowChequeSigner) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowChequeSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowChequeSigner)
				if err := _Cash.contract.UnpackLog(event, "ShowChequeSigner", log); err != nil {
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

// ParseShowChequeSigner is a log parse operation binding the contract event 0xd992130aab86d2580aa88e21392e36739b1621554343d409620f0fb1e698b33d.
//
// Solidity: event ShowChequeSigner(address arg0)
func (_Cash *CashFilterer) ParseShowChequeSigner(log types.Log) (*CashShowChequeSigner, error) {
	event := new(CashShowChequeSigner)
	if err := _Cash.contract.UnpackLog(event, "ShowChequeSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowFromIterator is returned from FilterShowFrom and is used to iterate over the raw logs and unpacked data for ShowFrom events raised by the Cash contract.
type CashShowFromIterator struct {
	Event *CashShowFrom // Event containing the contract specifics and raw log

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
func (it *CashShowFromIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowFrom)
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
		it.Event = new(CashShowFrom)
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
func (it *CashShowFromIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowFromIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowFrom represents a ShowFrom event raised by the Cash contract.
type CashShowFrom struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowFrom is a free log retrieval operation binding the contract event 0xe7d038274dbda90ba9b4a73bcf23a30cabd0a1567a2b29d19d2accf5399f0309.
//
// Solidity: event ShowFrom(address arg0)
func (_Cash *CashFilterer) FilterShowFrom(opts *bind.FilterOpts) (*CashShowFromIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowFrom")
	if err != nil {
		return nil, err
	}
	return &CashShowFromIterator{contract: _Cash.contract, event: "ShowFrom", logs: logs, sub: sub}, nil
}

// WatchShowFrom is a free log subscription operation binding the contract event 0xe7d038274dbda90ba9b4a73bcf23a30cabd0a1567a2b29d19d2accf5399f0309.
//
// Solidity: event ShowFrom(address arg0)
func (_Cash *CashFilterer) WatchShowFrom(opts *bind.WatchOpts, sink chan<- *CashShowFrom) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowFrom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowFrom)
				if err := _Cash.contract.UnpackLog(event, "ShowFrom", log); err != nil {
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

// ParseShowFrom is a log parse operation binding the contract event 0xe7d038274dbda90ba9b4a73bcf23a30cabd0a1567a2b29d19d2accf5399f0309.
//
// Solidity: event ShowFrom(address arg0)
func (_Cash *CashFilterer) ParseShowFrom(log types.Log) (*CashShowFrom, error) {
	event := new(CashShowFrom)
	if err := _Cash.contract.UnpackLog(event, "ShowFrom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowNonceIterator is returned from FilterShowNonce and is used to iterate over the raw logs and unpacked data for ShowNonce events raised by the Cash contract.
type CashShowNonceIterator struct {
	Event *CashShowNonce // Event containing the contract specifics and raw log

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
func (it *CashShowNonceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowNonce)
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
		it.Event = new(CashShowNonce)
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
func (it *CashShowNonceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowNonce represents a ShowNonce event raised by the Cash contract.
type CashShowNonce struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowNonce is a free log retrieval operation binding the contract event 0xaf3a06b91ffcbd286a1f609888182ab42c9bb71d1f95491bfe3bc6d300a48184.
//
// Solidity: event ShowNonce(uint256 arg0)
func (_Cash *CashFilterer) FilterShowNonce(opts *bind.FilterOpts) (*CashShowNonceIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowNonce")
	if err != nil {
		return nil, err
	}
	return &CashShowNonceIterator{contract: _Cash.contract, event: "ShowNonce", logs: logs, sub: sub}, nil
}

// WatchShowNonce is a free log subscription operation binding the contract event 0xaf3a06b91ffcbd286a1f609888182ab42c9bb71d1f95491bfe3bc6d300a48184.
//
// Solidity: event ShowNonce(uint256 arg0)
func (_Cash *CashFilterer) WatchShowNonce(opts *bind.WatchOpts, sink chan<- *CashShowNonce) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowNonce")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowNonce)
				if err := _Cash.contract.UnpackLog(event, "ShowNonce", log); err != nil {
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

// ParseShowNonce is a log parse operation binding the contract event 0xaf3a06b91ffcbd286a1f609888182ab42c9bb71d1f95491bfe3bc6d300a48184.
//
// Solidity: event ShowNonce(uint256 arg0)
func (_Cash *CashFilterer) ParseShowNonce(log types.Log) (*CashShowNonce, error) {
	event := new(CashShowNonce)
	if err := _Cash.contract.UnpackLog(event, "ShowNonce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowPayCheckPackIterator is returned from FilterShowPayCheckPack and is used to iterate over the raw logs and unpacked data for ShowPayCheckPack events raised by the Cash contract.
type CashShowPayCheckPackIterator struct {
	Event *CashShowPayCheckPack // Event containing the contract specifics and raw log

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
func (it *CashShowPayCheckPackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowPayCheckPack)
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
		it.Event = new(CashShowPayCheckPack)
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
func (it *CashShowPayCheckPackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowPayCheckPackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowPayCheckPack represents a ShowPayCheckPack event raised by the Cash contract.
type CashShowPayCheckPack struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowPayCheckPack is a free log retrieval operation binding the contract event 0xe254bf5fc12a7e7953bcdfd673d53edf48f7b614e7189fbfd2ddc3514d3abd48.
//
// Solidity: event ShowPayCheckPack(bytes arg0)
func (_Cash *CashFilterer) FilterShowPayCheckPack(opts *bind.FilterOpts) (*CashShowPayCheckPackIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowPayCheckPack")
	if err != nil {
		return nil, err
	}
	return &CashShowPayCheckPackIterator{contract: _Cash.contract, event: "ShowPayCheckPack", logs: logs, sub: sub}, nil
}

// WatchShowPayCheckPack is a free log subscription operation binding the contract event 0xe254bf5fc12a7e7953bcdfd673d53edf48f7b614e7189fbfd2ddc3514d3abd48.
//
// Solidity: event ShowPayCheckPack(bytes arg0)
func (_Cash *CashFilterer) WatchShowPayCheckPack(opts *bind.WatchOpts, sink chan<- *CashShowPayCheckPack) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowPayCheckPack")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowPayCheckPack)
				if err := _Cash.contract.UnpackLog(event, "ShowPayCheckPack", log); err != nil {
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

// ParseShowPayCheckPack is a log parse operation binding the contract event 0xe254bf5fc12a7e7953bcdfd673d53edf48f7b614e7189fbfd2ddc3514d3abd48.
//
// Solidity: event ShowPayCheckPack(bytes arg0)
func (_Cash *CashFilterer) ParseShowPayCheckPack(log types.Log) (*CashShowPayCheckPack, error) {
	event := new(CashShowPayCheckPack)
	if err := _Cash.contract.UnpackLog(event, "ShowPayCheckPack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowPayChequeHashIterator is returned from FilterShowPayChequeHash and is used to iterate over the raw logs and unpacked data for ShowPayChequeHash events raised by the Cash contract.
type CashShowPayChequeHashIterator struct {
	Event *CashShowPayChequeHash // Event containing the contract specifics and raw log

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
func (it *CashShowPayChequeHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowPayChequeHash)
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
		it.Event = new(CashShowPayChequeHash)
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
func (it *CashShowPayChequeHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowPayChequeHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowPayChequeHash represents a ShowPayChequeHash event raised by the Cash contract.
type CashShowPayChequeHash struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowPayChequeHash is a free log retrieval operation binding the contract event 0x267ef9c051dadf48f9d78139f8afcaab9b9ea923e2f62607fb90558d169399f0.
//
// Solidity: event ShowPayChequeHash(bytes32 arg0)
func (_Cash *CashFilterer) FilterShowPayChequeHash(opts *bind.FilterOpts) (*CashShowPayChequeHashIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowPayChequeHash")
	if err != nil {
		return nil, err
	}
	return &CashShowPayChequeHashIterator{contract: _Cash.contract, event: "ShowPayChequeHash", logs: logs, sub: sub}, nil
}

// WatchShowPayChequeHash is a free log subscription operation binding the contract event 0x267ef9c051dadf48f9d78139f8afcaab9b9ea923e2f62607fb90558d169399f0.
//
// Solidity: event ShowPayChequeHash(bytes32 arg0)
func (_Cash *CashFilterer) WatchShowPayChequeHash(opts *bind.WatchOpts, sink chan<- *CashShowPayChequeHash) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowPayChequeHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowPayChequeHash)
				if err := _Cash.contract.UnpackLog(event, "ShowPayChequeHash", log); err != nil {
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

// ParseShowPayChequeHash is a log parse operation binding the contract event 0x267ef9c051dadf48f9d78139f8afcaab9b9ea923e2f62607fb90558d169399f0.
//
// Solidity: event ShowPayChequeHash(bytes32 arg0)
func (_Cash *CashFilterer) ParseShowPayChequeHash(log types.Log) (*CashShowPayChequeHash, error) {
	event := new(CashShowPayChequeHash)
	if err := _Cash.contract.UnpackLog(event, "ShowPayChequeHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowPayChequeSignerIterator is returned from FilterShowPayChequeSigner and is used to iterate over the raw logs and unpacked data for ShowPayChequeSigner events raised by the Cash contract.
type CashShowPayChequeSignerIterator struct {
	Event *CashShowPayChequeSigner // Event containing the contract specifics and raw log

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
func (it *CashShowPayChequeSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowPayChequeSigner)
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
		it.Event = new(CashShowPayChequeSigner)
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
func (it *CashShowPayChequeSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowPayChequeSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowPayChequeSigner represents a ShowPayChequeSigner event raised by the Cash contract.
type CashShowPayChequeSigner struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowPayChequeSigner is a free log retrieval operation binding the contract event 0x129e8dff958b080ea26f69bc288d31037ee2127de5b7e3784e72ee7fdb8901e2.
//
// Solidity: event ShowPayChequeSigner(address arg0)
func (_Cash *CashFilterer) FilterShowPayChequeSigner(opts *bind.FilterOpts) (*CashShowPayChequeSignerIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowPayChequeSigner")
	if err != nil {
		return nil, err
	}
	return &CashShowPayChequeSignerIterator{contract: _Cash.contract, event: "ShowPayChequeSigner", logs: logs, sub: sub}, nil
}

// WatchShowPayChequeSigner is a free log subscription operation binding the contract event 0x129e8dff958b080ea26f69bc288d31037ee2127de5b7e3784e72ee7fdb8901e2.
//
// Solidity: event ShowPayChequeSigner(address arg0)
func (_Cash *CashFilterer) WatchShowPayChequeSigner(opts *bind.WatchOpts, sink chan<- *CashShowPayChequeSigner) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowPayChequeSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowPayChequeSigner)
				if err := _Cash.contract.UnpackLog(event, "ShowPayChequeSigner", log); err != nil {
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

// ParseShowPayChequeSigner is a log parse operation binding the contract event 0x129e8dff958b080ea26f69bc288d31037ee2127de5b7e3784e72ee7fdb8901e2.
//
// Solidity: event ShowPayChequeSigner(address arg0)
func (_Cash *CashFilterer) ParseShowPayChequeSigner(log types.Log) (*CashShowPayChequeSigner, error) {
	event := new(CashShowPayChequeSigner)
	if err := _Cash.contract.UnpackLog(event, "ShowPayChequeSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
