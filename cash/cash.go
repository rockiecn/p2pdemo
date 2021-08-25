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
	OpAddr    common.Address
	FromAddr  common.Address
	ToAddr    common.Address
	TokenAddr common.Address
	Value     *big.Int
	NodeNonce *big.Int
}

// PayCheque is an auto generated low-level Go binding around an user-defined struct.
type PayCheque struct {
	Cheque    Cheque
	ChequeSig []byte
	CashAddr  common.Address
	FromAddr  common.Address
	ToAddr    common.Address
	PayValue  *big.Int
}

// CashMetaData contains all meta data concerning the Cash contract.
var CashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ShowFrom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ShowHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ShowNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"ShowPack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"ShowSig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ShowSigner\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeNonce\",\"type\":\"uint256\"}],\"internalType\":\"structCheque\",\"name\":\"cheque\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"chequeSig\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"cashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payValue\",\"type\":\"uint256\"}],\"internalType\":\"structPayCheque\",\"name\":\"paycheque\",\"type\":\"tuple\"}],\"name\":\"apply_cheque\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052610b0b806100136000396000f3fe60806040526004361061001e5760003560e01c8063babf801114610023575b600080fd5b61003d600480360381019061003891906105e9565b610053565b60405161004a9190610762565b60405180910390f35b60007fe7d038274dbda90ba9b4a73bcf23a30cabd0a1567a2b29d19d2accf5399f030982600001516020015160405161008c9190610747565b60405180910390a17faf3a06b91ffcbd286a1f609888182ab42c9bb71d1f95491bfe3bc6d300a48184826000015160a001516040516100cb91906107ff565b60405180910390a17f769b78495bdc939f1cebbfbadf4ebc075e9f897031a598312365555d70019ff2826020015160405161010691906107dd565b60405180910390a16000826000015160200151836000015160a001516000604051602001610136939291906106ff565b60405160208183030381529060405290507fa715a6946641bea3f72cecb3acc530b4c2bae0a20e41cb7b0022ef8f8750ece38160405161017691906107dd565b60405180910390a16000836000015160200151846000015160a0015160006040516020016101a6939291906106ff565b6040516020818303038152906040528051906020012090507f745731c823271a437c86a7af6d16856092b7dcf98734a5b50c09c64562e5692f816040516101ed919061077d565b60405180910390a160006102058286602001516102fb565b90507f19e241431b2317b16d0f71f8af360a0eb20fbdbfdb535358dc6136c9f1b5cada816040516102369190610747565b60405180910390a18073ffffffffffffffffffffffffffffffffffffffff1685600001516000015173ffffffffffffffffffffffffffffffffffffffff1614156102ee576000670de0b6b3a76400008660a0015161029491906108ce565b9050856080015173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156102e0573d6000803e3d6000fd5b5060019450505050506102f6565b600093505050505b919050565b6000604182511461030f57600090506103fd565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c111561036357600093505050506103fd565b601b8160ff16101561037f57601b8161037c9190610897565b90505b601b8160ff16141580156103975750601c8160ff1614155b156103a857600093505050506103fd565b600186828585604051600081526020016040526040516103cb9493929190610798565b6020604051602081039080840390855afa1580156103ed573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b60006104166104118461083f565b61081a565b90508281526020810184848401111561042e57600080fd5b610439848285610987565b509392505050565b60008135905061045081610aa7565b92915050565b600082601f83011261046757600080fd5b8135610477848260208601610403565b91505092915050565b600060c0828403121561049257600080fd5b61049c60c061081a565b905060006104ac84828501610441565b60008301525060206104c084828501610441565b60208301525060406104d484828501610441565b60408301525060606104e884828501610441565b60608301525060806104fc848285016105d4565b60808301525060a0610510848285016105d4565b60a08301525092915050565b6000610160828403121561052f57600080fd5b61053960c061081a565b9050600061054984828501610480565b60008301525060c082013567ffffffffffffffff81111561056957600080fd5b61057584828501610456565b60208301525060e061058984828501610441565b60408301525061010061059e84828501610441565b6060830152506101206105b384828501610441565b6080830152506101406105c8848285016105d4565b60a08301525092915050565b6000813590506105e381610abe565b92915050565b6000602082840312156105fb57600080fd5b600082013567ffffffffffffffff81111561061557600080fd5b6106218482850161051c565b91505092915050565b61063381610928565b82525050565b61064a61064582610928565b6109fa565b82525050565b6106598161093a565b82525050565b61066881610946565b82525050565b600061067982610870565b610683818561087b565b9350610693818560208601610996565b61069c81610a86565b840191505092915050565b60006106b460008361088c565b91506106bf82610aa4565b600082019050919050565b6106d381610970565b82525050565b6106ea6106e582610970565b610a1e565b82525050565b6106f98161097a565b82525050565b600061070b8286610639565b60148201915061071b82856106d9565b60208201915061072a826106a7565b915061073682846106d9565b602082019150819050949350505050565b600060208201905061075c600083018461062a565b92915050565b60006020820190506107776000830184610650565b92915050565b6000602082019050610792600083018461065f565b92915050565b60006080820190506107ad600083018761065f565b6107ba60208301866106f0565b6107c7604083018561065f565b6107d4606083018461065f565b95945050505050565b600060208201905081810360008301526107f7818461066e565b905092915050565b600060208201905061081460008301846106ca565b92915050565b6000610824610835565b905061083082826109c9565b919050565b6000604051905090565b600067ffffffffffffffff82111561085a57610859610a57565b5b61086382610a86565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006108a28261097a565b91506108ad8361097a565b92508260ff038211156108c3576108c2610a28565b5b828201905092915050565b60006108d982610970565b91506108e483610970565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561091d5761091c610a28565b5b828202905092915050565b600061093382610950565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156109b4578082015181840152602081019050610999565b838111156109c3576000848401525b50505050565b6109d282610a86565b810181811067ffffffffffffffff821117156109f1576109f0610a57565b5b80604052505050565b6000610a0582610a0c565b9050919050565b6000610a1782610a97565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b50565b610ab081610928565b8114610abb57600080fd5b50565b610ac781610970565b8114610ad257600080fd5b5056fea2646970667358221220020a764c52d99dce0c77bd29149346580a21907fdaa8d104cf43372209dbefcb64736f6c63430008020033",
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

// ApplyCheque is a paid mutator transaction binding the contract method 0xbabf8011.
//
// Solidity: function apply_cheque(((address,address,address,address,uint256,uint256),bytes,address,address,address,uint256) paycheque) payable returns(bool)
func (_Cash *CashTransactor) ApplyCheque(opts *bind.TransactOpts, paycheque PayCheque) (*types.Transaction, error) {
	return _Cash.contract.Transact(opts, "apply_cheque", paycheque)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0xbabf8011.
//
// Solidity: function apply_cheque(((address,address,address,address,uint256,uint256),bytes,address,address,address,uint256) paycheque) payable returns(bool)
func (_Cash *CashSession) ApplyCheque(paycheque PayCheque) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheque(&_Cash.TransactOpts, paycheque)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0xbabf8011.
//
// Solidity: function apply_cheque(((address,address,address,address,uint256,uint256),bytes,address,address,address,uint256) paycheque) payable returns(bool)
func (_Cash *CashTransactorSession) ApplyCheque(paycheque PayCheque) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheque(&_Cash.TransactOpts, paycheque)
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

// CashShowHashIterator is returned from FilterShowHash and is used to iterate over the raw logs and unpacked data for ShowHash events raised by the Cash contract.
type CashShowHashIterator struct {
	Event *CashShowHash // Event containing the contract specifics and raw log

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
func (it *CashShowHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowHash)
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
		it.Event = new(CashShowHash)
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
func (it *CashShowHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowHash represents a ShowHash event raised by the Cash contract.
type CashShowHash struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowHash is a free log retrieval operation binding the contract event 0x745731c823271a437c86a7af6d16856092b7dcf98734a5b50c09c64562e5692f.
//
// Solidity: event ShowHash(bytes32 arg0)
func (_Cash *CashFilterer) FilterShowHash(opts *bind.FilterOpts) (*CashShowHashIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowHash")
	if err != nil {
		return nil, err
	}
	return &CashShowHashIterator{contract: _Cash.contract, event: "ShowHash", logs: logs, sub: sub}, nil
}

// WatchShowHash is a free log subscription operation binding the contract event 0x745731c823271a437c86a7af6d16856092b7dcf98734a5b50c09c64562e5692f.
//
// Solidity: event ShowHash(bytes32 arg0)
func (_Cash *CashFilterer) WatchShowHash(opts *bind.WatchOpts, sink chan<- *CashShowHash) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowHash)
				if err := _Cash.contract.UnpackLog(event, "ShowHash", log); err != nil {
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

// ParseShowHash is a log parse operation binding the contract event 0x745731c823271a437c86a7af6d16856092b7dcf98734a5b50c09c64562e5692f.
//
// Solidity: event ShowHash(bytes32 arg0)
func (_Cash *CashFilterer) ParseShowHash(log types.Log) (*CashShowHash, error) {
	event := new(CashShowHash)
	if err := _Cash.contract.UnpackLog(event, "ShowHash", log); err != nil {
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

// CashShowPackIterator is returned from FilterShowPack and is used to iterate over the raw logs and unpacked data for ShowPack events raised by the Cash contract.
type CashShowPackIterator struct {
	Event *CashShowPack // Event containing the contract specifics and raw log

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
func (it *CashShowPackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowPack)
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
		it.Event = new(CashShowPack)
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
func (it *CashShowPackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowPackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowPack represents a ShowPack event raised by the Cash contract.
type CashShowPack struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowPack is a free log retrieval operation binding the contract event 0xa715a6946641bea3f72cecb3acc530b4c2bae0a20e41cb7b0022ef8f8750ece3.
//
// Solidity: event ShowPack(bytes arg0)
func (_Cash *CashFilterer) FilterShowPack(opts *bind.FilterOpts) (*CashShowPackIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowPack")
	if err != nil {
		return nil, err
	}
	return &CashShowPackIterator{contract: _Cash.contract, event: "ShowPack", logs: logs, sub: sub}, nil
}

// WatchShowPack is a free log subscription operation binding the contract event 0xa715a6946641bea3f72cecb3acc530b4c2bae0a20e41cb7b0022ef8f8750ece3.
//
// Solidity: event ShowPack(bytes arg0)
func (_Cash *CashFilterer) WatchShowPack(opts *bind.WatchOpts, sink chan<- *CashShowPack) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowPack")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowPack)
				if err := _Cash.contract.UnpackLog(event, "ShowPack", log); err != nil {
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

// ParseShowPack is a log parse operation binding the contract event 0xa715a6946641bea3f72cecb3acc530b4c2bae0a20e41cb7b0022ef8f8750ece3.
//
// Solidity: event ShowPack(bytes arg0)
func (_Cash *CashFilterer) ParseShowPack(log types.Log) (*CashShowPack, error) {
	event := new(CashShowPack)
	if err := _Cash.contract.UnpackLog(event, "ShowPack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowSigIterator is returned from FilterShowSig and is used to iterate over the raw logs and unpacked data for ShowSig events raised by the Cash contract.
type CashShowSigIterator struct {
	Event *CashShowSig // Event containing the contract specifics and raw log

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
func (it *CashShowSigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowSig)
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
		it.Event = new(CashShowSig)
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
func (it *CashShowSigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowSigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowSig represents a ShowSig event raised by the Cash contract.
type CashShowSig struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowSig is a free log retrieval operation binding the contract event 0x769b78495bdc939f1cebbfbadf4ebc075e9f897031a598312365555d70019ff2.
//
// Solidity: event ShowSig(bytes arg0)
func (_Cash *CashFilterer) FilterShowSig(opts *bind.FilterOpts) (*CashShowSigIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowSig")
	if err != nil {
		return nil, err
	}
	return &CashShowSigIterator{contract: _Cash.contract, event: "ShowSig", logs: logs, sub: sub}, nil
}

// WatchShowSig is a free log subscription operation binding the contract event 0x769b78495bdc939f1cebbfbadf4ebc075e9f897031a598312365555d70019ff2.
//
// Solidity: event ShowSig(bytes arg0)
func (_Cash *CashFilterer) WatchShowSig(opts *bind.WatchOpts, sink chan<- *CashShowSig) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowSig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowSig)
				if err := _Cash.contract.UnpackLog(event, "ShowSig", log); err != nil {
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

// ParseShowSig is a log parse operation binding the contract event 0x769b78495bdc939f1cebbfbadf4ebc075e9f897031a598312365555d70019ff2.
//
// Solidity: event ShowSig(bytes arg0)
func (_Cash *CashFilterer) ParseShowSig(log types.Log) (*CashShowSig, error) {
	event := new(CashShowSig)
	if err := _Cash.contract.UnpackLog(event, "ShowSig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowSignerIterator is returned from FilterShowSigner and is used to iterate over the raw logs and unpacked data for ShowSigner events raised by the Cash contract.
type CashShowSignerIterator struct {
	Event *CashShowSigner // Event containing the contract specifics and raw log

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
func (it *CashShowSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowSigner)
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
		it.Event = new(CashShowSigner)
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
func (it *CashShowSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowSigner represents a ShowSigner event raised by the Cash contract.
type CashShowSigner struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowSigner is a free log retrieval operation binding the contract event 0x19e241431b2317b16d0f71f8af360a0eb20fbdbfdb535358dc6136c9f1b5cada.
//
// Solidity: event ShowSigner(address arg0)
func (_Cash *CashFilterer) FilterShowSigner(opts *bind.FilterOpts) (*CashShowSignerIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "ShowSigner")
	if err != nil {
		return nil, err
	}
	return &CashShowSignerIterator{contract: _Cash.contract, event: "ShowSigner", logs: logs, sub: sub}, nil
}

// WatchShowSigner is a free log subscription operation binding the contract event 0x19e241431b2317b16d0f71f8af360a0eb20fbdbfdb535358dc6136c9f1b5cada.
//
// Solidity: event ShowSigner(address arg0)
func (_Cash *CashFilterer) WatchShowSigner(opts *bind.WatchOpts, sink chan<- *CashShowSigner) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "ShowSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowSigner)
				if err := _Cash.contract.UnpackLog(event, "ShowSigner", log); err != nil {
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

// ParseShowSigner is a log parse operation binding the contract event 0x19e241431b2317b16d0f71f8af360a0eb20fbdbfdb535358dc6136c9f1b5cada.
//
// Solidity: event ShowSigner(address arg0)
func (_Cash *CashFilterer) ParseShowSigner(log types.Log) (*CashShowSigner, error) {
	event := new(CashShowSigner)
	if err := _Cash.contract.UnpackLog(event, "ShowSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
