// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cash

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CashABI is the input ABI used to generate the binding from.
const CashABI = "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"show\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"apply_cheque\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// CashBin is the compiled bytecode used for deploying new contracts.
var CashBin = "0x6080604052610733806100136000396000f3fe60806040526004361061001e5760003560e01c8063163dbf9814610023575b600080fd5b61003d600480360381019061003891906102eb565b610053565b60405161004a919061042b565b60405180910390f35b6000808686868660405160200161006d94939291906103dd565b60405160208183030381529060405280519060200120905060006100918285610149565b905060008190508173ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff161415610138576000670de0b6b3a7640000876100e29190610518565b90508773ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561012a573d6000803e3d6000fd5b506001945050505050610140565b600093505050505b95945050505050565b6000604182511461015d576000905061024b565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c11156101b1576000935050505061024b565b601b8160ff1610156101cd57601b816101ca91906104e1565b90505b601b8160ff16141580156101e55750601c8160ff1614155b156101f6576000935050505061024b565b600186828585604051600081526020016040526040516102199493929190610446565b6020604051602081039080840390855afa15801561023b573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b600061026461025f846104b0565b61048b565b9050828152602081018484840111156102805761027f6106a2565b5b61028b8482856105d1565b509392505050565b6000813590506102a2816106cf565b92915050565b600082601f8301126102bd576102bc61069d565b5b81356102cd848260208601610251565b91505092915050565b6000813590506102e5816106e6565b92915050565b600080600080600060a08688031215610307576103066106ac565b5b600061031588828901610293565b9550506020610326888289016102d6565b945050604061033788828901610293565b9350506060610348888289016102d6565b925050608086013567ffffffffffffffff811115610369576103686106a7565b5b610375888289016102a8565b9150509295509295909350565b61039361038e82610572565b610611565b82525050565b6103a281610584565b82525050565b6103b181610590565b82525050565b6103c86103c3826105ba565b610635565b82525050565b6103d7816105c4565b82525050565b60006103e98287610382565b6014820191506103f982866103b7565b6020820191506104098285610382565b60148201915061041982846103b7565b60208201915081905095945050505050565b60006020820190506104406000830184610399565b92915050565b600060808201905061045b60008301876103a8565b61046860208301866103ce565b61047560408301856103a8565b61048260608301846103a8565b95945050505050565b60006104956104a6565b90506104a182826105e0565b919050565b6000604051905090565b600067ffffffffffffffff8211156104cb576104ca61066e565b5b6104d4826106b1565b9050602081019050919050565b60006104ec826105c4565b91506104f7836105c4565b92508260ff0382111561050d5761050c61063f565b5b828201905092915050565b6000610523826105ba565b915061052e836105ba565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156105675761056661063f565b5b828202905092915050565b600061057d8261059a565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b6105e9826106b1565b810181811067ffffffffffffffff821117156106085761060761066e565b5b80604052505050565b600061061c82610623565b9050919050565b600061062e826106c2565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b6106d881610572565b81146106e357600080fd5b50565b6106ef816105ba565b81146106fa57600080fd5b5056fea264697066735822122031fc95d17087495be91463ab7453ae0d760b439793c003bc508a6e3b3b92bdb564736f6c63430008070033"

// DeployCash deploys a new Ethereum contract, binding an instance of Cash to it.
func DeployCash(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cash, error) {
	parsed, err := abi.JSON(strings.NewReader(CashABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CashBin), backend)
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

// ApplyCheque is a paid mutator transaction binding the contract method 0x163dbf98.
//
// Solidity: function apply_cheque(address userAddr, uint256 nonce, address stAddr, uint256 payAmount, bytes sign) payable returns(bool)
func (_Cash *CashTransactor) ApplyCheque(opts *bind.TransactOpts, userAddr common.Address, nonce *big.Int, stAddr common.Address, payAmount *big.Int, sign []byte) (*types.Transaction, error) {
	return _Cash.contract.Transact(opts, "apply_cheque", userAddr, nonce, stAddr, payAmount, sign)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0x163dbf98.
//
// Solidity: function apply_cheque(address userAddr, uint256 nonce, address stAddr, uint256 payAmount, bytes sign) payable returns(bool)
func (_Cash *CashSession) ApplyCheque(userAddr common.Address, nonce *big.Int, stAddr common.Address, payAmount *big.Int, sign []byte) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheque(&_Cash.TransactOpts, userAddr, nonce, stAddr, payAmount, sign)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0x163dbf98.
//
// Solidity: function apply_cheque(address userAddr, uint256 nonce, address stAddr, uint256 payAmount, bytes sign) payable returns(bool)
func (_Cash *CashTransactorSession) ApplyCheque(userAddr common.Address, nonce *big.Int, stAddr common.Address, payAmount *big.Int, sign []byte) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheque(&_Cash.TransactOpts, userAddr, nonce, stAddr, payAmount, sign)
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
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShow is a free log retrieval operation binding the contract event 0xfb687b8805689fd7d97b0f174d316f660a67e3dc6bfb412c484bef58f01f18e6.
//
// Solidity: event show(bytes arg0)
func (_Cash *CashFilterer) FilterShow(opts *bind.FilterOpts) (*CashShowIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "show")
	if err != nil {
		return nil, err
	}
	return &CashShowIterator{contract: _Cash.contract, event: "show", logs: logs, sub: sub}, nil
}

// WatchShow is a free log subscription operation binding the contract event 0xfb687b8805689fd7d97b0f174d316f660a67e3dc6bfb412c484bef58f01f18e6.
//
// Solidity: event show(bytes arg0)
func (_Cash *CashFilterer) WatchShow(opts *bind.WatchOpts, sink chan<- *CashShow) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "show")
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
				if err := _Cash.contract.UnpackLog(event, "show", log); err != nil {
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

// ParseShow is a log parse operation binding the contract event 0xfb687b8805689fd7d97b0f174d316f660a67e3dc6bfb412c484bef58f01f18e6.
//
// Solidity: event show(bytes arg0)
func (_Cash *CashFilterer) ParseShow(log types.Log) (*CashShow, error) {
	event := new(CashShow)
	if err := _Cash.contract.UnpackLog(event, "show", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
