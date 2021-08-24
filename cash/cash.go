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

// CashMetaData contains all meta data concerning the Cash contract.
var CashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ShowString\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"Showbytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Showuint\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"stringParams\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"intParams\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"bytesParam\",\"type\":\"bytes\"}],\"name\":\"apply_cheque\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052610757806100136000396000f3fe60806040526004361061001e5760003560e01c806395910a5f14610023575b600080fd5b61003d60048036038101906100389190610389565b610053565b60405161004a91906104b0565b60405180910390f35b60007f1514c1afb85b99e3311ee11a9f096aa6d412c6ab0a39884d0a973b894acefcc3846000815181106100b0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040516100c591906104ed565b60405180910390a17fe484bc988becf0615e0d005c616f6cd84fdbe2cebff2c84e8cc1894df37f9a7f83600081518110610128577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160405161013d919061050f565b60405180910390a17f4daa2d63a96f5e545f8ff5b850799b444e2e8a26e69f59425ba735fa0528a2688260405161017491906104cb565b60405180910390a1600190509392505050565b600061019a6101958461054f565b61052a565b9050808382526020820190508260005b858110156101da57813585016101c0888261034a565b8452602084019350602083019250506001810190506101aa565b5050509392505050565b60006101f76101f28461057b565b61052a565b9050808382526020820190508285602086028201111561021657600080fd5b60005b85811015610246578161022c8882610374565b845260208401935060208301925050600181019050610219565b5050509392505050565b600061026361025e846105a7565b61052a565b90508281526020810184848401111561027b57600080fd5b610286848285610657565b509392505050565b60006102a161029c846105d8565b61052a565b9050828152602081018484840111156102b957600080fd5b6102c4848285610657565b509392505050565b600082601f8301126102dd57600080fd5b81356102ed848260208601610187565b91505092915050565b600082601f83011261030757600080fd5b81356103178482602086016101e4565b91505092915050565b600082601f83011261033157600080fd5b8135610341848260208601610250565b91505092915050565b600082601f83011261035b57600080fd5b813561036b84826020860161028e565b91505092915050565b6000813590506103838161070a565b92915050565b60008060006060848603121561039e57600080fd5b600084013567ffffffffffffffff8111156103b857600080fd5b6103c4868287016102cc565b935050602084013567ffffffffffffffff8111156103e157600080fd5b6103ed868287016102f6565b925050604084013567ffffffffffffffff81111561040a57600080fd5b61041686828701610320565b9150509250925092565b61042981610641565b82525050565b600061043a82610609565b610444818561061f565b9350610454818560208601610666565b61045d816106f9565b840191505092915050565b600061047382610614565b61047d8185610630565b935061048d818560208601610666565b610496816106f9565b840191505092915050565b6104aa8161064d565b82525050565b60006020820190506104c56000830184610420565b92915050565b600060208201905081810360008301526104e5818461042f565b905092915050565b600060208201905081810360008301526105078184610468565b905092915050565b600060208201905061052460008301846104a1565b92915050565b6000610534610545565b90506105408282610699565b919050565b6000604051905090565b600067ffffffffffffffff82111561056a576105696106ca565b5b602082029050602081019050919050565b600067ffffffffffffffff821115610596576105956106ca565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156105c2576105c16106ca565b5b6105cb826106f9565b9050602081019050919050565b600067ffffffffffffffff8211156105f3576105f26106ca565b5b6105fc826106f9565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60008115159050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610684578082015181840152602081019050610669565b83811115610693576000848401525b50505050565b6106a2826106f9565b810181811067ffffffffffffffff821117156106c1576106c06106ca565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b6107138161064d565b811461071e57600080fd5b5056fea26469706673582212203b6ade9e84d2da75f59d96431a19998aaa07011e6aa4f65ce0777b881d982ded64736f6c63430008020033",
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

// ApplyCheque is a paid mutator transaction binding the contract method 0x95910a5f.
//
// Solidity: function apply_cheque(string[] stringParams, uint256[] intParams, bytes bytesParam) payable returns(bool)
func (_Cash *CashTransactor) ApplyCheque(opts *bind.TransactOpts, stringParams []string, intParams []*big.Int, bytesParam []byte) (*types.Transaction, error) {
	return _Cash.contract.Transact(opts, "apply_cheque", stringParams, intParams, bytesParam)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0x95910a5f.
//
// Solidity: function apply_cheque(string[] stringParams, uint256[] intParams, bytes bytesParam) payable returns(bool)
func (_Cash *CashSession) ApplyCheque(stringParams []string, intParams []*big.Int, bytesParam []byte) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheque(&_Cash.TransactOpts, stringParams, intParams, bytesParam)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0x95910a5f.
//
// Solidity: function apply_cheque(string[] stringParams, uint256[] intParams, bytes bytesParam) payable returns(bool)
func (_Cash *CashTransactorSession) ApplyCheque(stringParams []string, intParams []*big.Int, bytesParam []byte) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheque(&_Cash.TransactOpts, stringParams, intParams, bytesParam)
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
