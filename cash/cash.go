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
	OpAddr    string
	FromAddr  string
	ToAddr    string
	TokenAddr string
	Value     *big.Int
	NodeNonce *big.Int
}

// PayCheque is an auto generated low-level Go binding around an user-defined struct.
type PayCheque struct {
	Cheque    Cheque
	ChequeSig []byte
	CashAddr  string
	FromAddr  string
	ToAddr    string
	PayValue  *big.Int
}

// CashMetaData contains all meta data concerning the Cash contract.
var CashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ShowString\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"Showbytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Showuint\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"opAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fromAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNonce\",\"type\":\"uint256\"}],\"internalType\":\"structCheque\",\"name\":\"cheque\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"chequeSig\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"cashAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fromAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"payValue\",\"type\":\"uint256\"}],\"internalType\":\"structPayCheque\",\"name\":\"paycheque\",\"type\":\"tuple\"}],\"name\":\"apply_cheque\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052610676806100136000396000f3fe60806040526004361061001e5760003560e01c80630d141d2514610023575b600080fd5b61003d600480360381019061003891906103cd565b610053565b60405161004a9190610465565b60405180910390f35b60007f1514c1afb85b99e3311ee11a9f096aa6d412c6ab0a39884d0a973b894acefcc382604001516040516100889190610480565b60405180910390a17fe484bc988becf0615e0d005c616f6cd84fdbe2cebff2c84e8cc1894df37f9a7f826000015160a001516040516100c791906104a2565b60405180910390a160019050919050565b60006100eb6100e6846104e2565b6104bd565b90508281526020810184848401111561010357600080fd5b61010e848285610576565b509392505050565b600061012961012484610513565b6104bd565b90508281526020810184848401111561014157600080fd5b61014c848285610576565b509392505050565b600082601f83011261016557600080fd5b81356101758482602086016100d8565b91505092915050565b600082601f83011261018f57600080fd5b813561019f848260208601610116565b91505092915050565b600060c082840312156101ba57600080fd5b6101c460c06104bd565b9050600082013567ffffffffffffffff8111156101e057600080fd5b6101ec8482850161017e565b600083015250602082013567ffffffffffffffff81111561020c57600080fd5b6102188482850161017e565b602083015250604082013567ffffffffffffffff81111561023857600080fd5b6102448482850161017e565b604083015250606082013567ffffffffffffffff81111561026457600080fd5b6102708482850161017e565b6060830152506080610284848285016103b8565b60808301525060a0610298848285016103b8565b60a08301525092915050565b600060c082840312156102b657600080fd5b6102c060c06104bd565b9050600082013567ffffffffffffffff8111156102dc57600080fd5b6102e8848285016101a8565b600083015250602082013567ffffffffffffffff81111561030857600080fd5b61031484828501610154565b602083015250604082013567ffffffffffffffff81111561033457600080fd5b6103408482850161017e565b604083015250606082013567ffffffffffffffff81111561036057600080fd5b61036c8482850161017e565b606083015250608082013567ffffffffffffffff81111561038c57600080fd5b6103988482850161017e565b60808301525060a06103ac848285016103b8565b60a08301525092915050565b6000813590506103c781610629565b92915050565b6000602082840312156103df57600080fd5b600082013567ffffffffffffffff8111156103f957600080fd5b610405848285016102a4565b91505092915050565b61041781610560565b82525050565b600061042882610544565b610432818561054f565b9350610442818560208601610585565b61044b81610618565b840191505092915050565b61045f8161056c565b82525050565b600060208201905061047a600083018461040e565b92915050565b6000602082019050818103600083015261049a818461041d565b905092915050565b60006020820190506104b76000830184610456565b92915050565b60006104c76104d8565b90506104d382826105b8565b919050565b6000604051905090565b600067ffffffffffffffff8211156104fd576104fc6105e9565b5b61050682610618565b9050602081019050919050565b600067ffffffffffffffff82111561052e5761052d6105e9565b5b61053782610618565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b60008115159050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156105a3578082015181840152602081019050610588565b838111156105b2576000848401525b50505050565b6105c182610618565b810181811067ffffffffffffffff821117156105e0576105df6105e9565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b6106328161056c565b811461063d57600080fd5b5056fea2646970667358221220a6b4458b790a5bc96a32c7894634bf5ff6ed39f8ac33ccd95585431e9a51e82164736f6c63430008020033",
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

// ApplyCheque is a paid mutator transaction binding the contract method 0x0d141d25.
//
// Solidity: function apply_cheque(((string,string,string,string,uint256,uint256),bytes,string,string,string,uint256) paycheque) payable returns(bool)
func (_Cash *CashTransactor) ApplyCheque(opts *bind.TransactOpts, paycheque PayCheque) (*types.Transaction, error) {
	return _Cash.contract.Transact(opts, "apply_cheque", paycheque)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0x0d141d25.
//
// Solidity: function apply_cheque(((string,string,string,string,uint256,uint256),bytes,string,string,string,uint256) paycheque) payable returns(bool)
func (_Cash *CashSession) ApplyCheque(paycheque PayCheque) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheque(&_Cash.TransactOpts, paycheque)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0x0d141d25.
//
// Solidity: function apply_cheque(((string,string,string,string,uint256,uint256),bytes,string,string,string,uint256) paycheque) payable returns(bool)
func (_Cash *CashTransactorSession) ApplyCheque(paycheque PayCheque) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheque(&_Cash.TransactOpts, paycheque)
}

// CashShowStringIterator is returned from FilterShowString and is used to iterate over the raw logs and unpacked data for ShowString events raised by the Cash contract.
type CashShowStringIterator struct {
	Event *CashShowString // Event containing the contract specifics and raw log

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
func (it *CashShowStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowString)
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
		it.Event = new(CashShowString)
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
func (it *CashShowStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowString represents a ShowString event raised by the Cash contract.
type CashShowString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowString is a free log retrieval operation binding the contract event 0x1514c1afb85b99e3311ee11a9f096aa6d412c6ab0a39884d0a973b894acefcc3.
//
// Solidity: event ShowString(string arg0)
func (_Cash *CashFilterer) FilterShowString(opts *bind.FilterOpts) (*CashShowStringIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowString")
	if err != nil {
		return nil, err
	}
	return &CashShowStringIterator{contract: _Cash.contract, event: "ShowString", logs: logs, sub: sub}, nil
}

// WatchShowString is a free log subscription operation binding the contract event 0x1514c1afb85b99e3311ee11a9f096aa6d412c6ab0a39884d0a973b894acefcc3.
//
// Solidity: event ShowString(string arg0)
func (_Cash *CashFilterer) WatchShowString(opts *bind.WatchOpts, sink chan<- *CashShowString) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowString")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowString)
				if err := _Cash.contract.UnpackLog(event, "ShowString", log); err != nil {
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

// ParseShowString is a log parse operation binding the contract event 0x1514c1afb85b99e3311ee11a9f096aa6d412c6ab0a39884d0a973b894acefcc3.
//
// Solidity: event ShowString(string arg0)
func (_Cash *CashFilterer) ParseShowString(log types.Log) (*CashShowString, error) {
	event := new(CashShowString)
	if err := _Cash.contract.UnpackLog(event, "ShowString", log); err != nil {
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

// CashShowuintIterator is returned from FilterShowuint and is used to iterate over the raw logs and unpacked data for Showuint events raised by the Cash contract.
type CashShowuintIterator struct {
	Event *CashShowuint // Event containing the contract specifics and raw log

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
func (it *CashShowuintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowuint)
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
		it.Event = new(CashShowuint)
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
func (it *CashShowuintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowuintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowuint represents a Showuint event raised by the Cash contract.
type CashShowuint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowuint is a free log retrieval operation binding the contract event 0xe484bc988becf0615e0d005c616f6cd84fdbe2cebff2c84e8cc1894df37f9a7f.
//
// Solidity: event Showuint(uint256 arg0)
func (_Cash *CashFilterer) FilterShowuint(opts *bind.FilterOpts) (*CashShowuintIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "Showuint")
	if err != nil {
		return nil, err
	}
	return &CashShowuintIterator{contract: _Cash.contract, event: "Showuint", logs: logs, sub: sub}, nil
}

// WatchShowuint is a free log subscription operation binding the contract event 0xe484bc988becf0615e0d005c616f6cd84fdbe2cebff2c84e8cc1894df37f9a7f.
//
// Solidity: event Showuint(uint256 arg0)
func (_Cash *CashFilterer) WatchShowuint(opts *bind.WatchOpts, sink chan<- *CashShowuint) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "Showuint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowuint)
				if err := _Cash.contract.UnpackLog(event, "Showuint", log); err != nil {
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

// ParseShowuint is a log parse operation binding the contract event 0xe484bc988becf0615e0d005c616f6cd84fdbe2cebff2c84e8cc1894df37f9a7f.
//
// Solidity: event Showuint(uint256 arg0)
func (_Cash *CashFilterer) ParseShowuint(log types.Log) (*CashShowuint, error) {
	event := new(CashShowuint)
	if err := _Cash.contract.UnpackLog(event, "Showuint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
