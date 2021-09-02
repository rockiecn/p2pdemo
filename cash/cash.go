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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Flag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"internalType\":\"structCheque\",\"name\":\"cheque\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"chequeSig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"payValue\",\"type\":\"uint256\"}],\"internalType\":\"structPayCheque\",\"name\":\"paycheque\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"paychequeSig\",\"type\":\"bytes\"}],\"name\":\"apply_cheque\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"get_node_nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611722806100536000396000f3fe6080604052600436106100435760003560e01c80630ac298dc146100885780631d728a0f146100b35780635f510b7b146100f057806393b8fec41461012d57610083565b36610083577f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f885258743334604051610079929190610f5a565b60405180910390a1005b600080fd5b34801561009457600080fd5b5061009d61015d565b6040516100aa9190610f3f565b60405180910390f35b3480156100bf57600080fd5b506100da60048036038101906100d59190610bdb565b610186565b6040516100e791906111b6565b60405180910390f35b3480156100fc57600080fd5b5061011760048036038101906101129190610bdb565b61019e565b60405161012491906111b6565b60405180910390f35b61014760048036038101906101429190610c04565b6101e7565b6040516101549190610f83565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60016020528060005260406000206000915090505481565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60006001600084600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054836000015160400151101561027b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610272906110d6565b60405180910390fd5b7f7fd5fca7bc379fa5c0330db55f9b725ddbe03460514a6fa71918966f3848779f60016040516102ab9190610fe3565b60405180910390a182600001516000015183604001511115610302576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f990611116565b60405180910390fd5b7f7fd5fca7bc379fa5c0330db55f9b725ddbe03460514a6fa71918966f3848779f60026040516103329190610ffe565b60405180910390a13073ffffffffffffffffffffffffffffffffffffffff16836000015160c0015173ffffffffffffffffffffffffffffffffffffffff16146103b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a7906110f6565b60405180910390fd5b7f7fd5fca7bc379fa5c0330db55f9b725ddbe03460514a6fa71918966f3848779f60036040516103e09190611019565b60405180910390a13373ffffffffffffffffffffffffffffffffffffffff1683600001516080015173ffffffffffffffffffffffffffffffffffffffff161461045e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045590611196565b60405180910390fd5b7f7fd5fca7bc379fa5c0330db55f9b725ddbe03460514a6fa71918966f3848779f600460405161048e9190611034565b60405180910390a160008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16836000015160a0015173ffffffffffffffffffffffffffffffffffffffff161461052c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052390611156565b60405180910390fd5b7f7fd5fca7bc379fa5c0330db55f9b725ddbe03460514a6fa71918966f3848779f600560405161055c919061104f565b60405180910390a16000836000015160000151846000015160200151856000015160400151866000015160600151876000015160800151886000015160a00151896000015160c001516040516020016105bb9796959493929190610ebe565b604051602081830303815290604052905060008185604001516040516020016105e5929190610e96565b6040516020818303038152906040529050600082805190602001209050600082805190602001209050600061061e838960200151610917565b9050600061062c8389610917565b90508173ffffffffffffffffffffffffffffffffffffffff16896000015160a0015173ffffffffffffffffffffffffffffffffffffffff16146106a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069b90611176565b60405180910390fd5b7f7fd5fca7bc379fa5c0330db55f9b725ddbe03460514a6fa71918966f3848779f60066040516106d4919061106a565b60405180910390a18073ffffffffffffffffffffffffffffffffffffffff1689600001516060015173ffffffffffffffffffffffffffffffffffffffff1614610752576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074990611136565b60405180910390fd5b7f7fd5fca7bc379fa5c0330db55f9b725ddbe03460514a6fa71918966f3848779f60076040516107829190611085565b60405180910390a16000670de0b6b3a76400008a604001516107a491906112db565b905089600001516080015173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156107f4573d6000803e3d6000fd5b507f737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc4898a60000151608001518260405161082e929190610f5a565b60405180910390a17f7fd5fca7bc379fa5c0330db55f9b725ddbe03460514a6fa71918966f3848779f600860405161086691906110a0565b60405180910390a160018a6000015160400151610883919061124e565b600160008c600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f7fd5fca7bc379fa5c0330db55f9b725ddbe03460514a6fa71918966f3848779f60096040516108fe91906110bb565b60405180910390a1600197505050505050505092915050565b6000604182511461092b5760009050610a19565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c111561097f5760009350505050610a19565b601b8160ff16101561099b57601b8161099891906112a4565b90505b601b8160ff16141580156109b35750601c8160ff1614155b156109c45760009350505050610a19565b600186828585604051600081526020016040526040516109e79493929190610f9e565b6020604051602081039080840390855afa158015610a09573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b6000610a32610a2d846111f6565b6111d1565b905082815260208101848484011115610a4a57600080fd5b610a55848285611436565b509392505050565b600081359050610a6c816116be565b92915050565b600082601f830112610a8357600080fd5b8135610a93848260208601610a1f565b91505092915050565b600060e08284031215610aae57600080fd5b610ab860e06111d1565b90506000610ac884828501610bc6565b6000830152506020610adc84828501610a5d565b6020830152506040610af084828501610bc6565b6040830152506060610b0484828501610a5d565b6060830152506080610b1884828501610a5d565b60808301525060a0610b2c84828501610a5d565b60a08301525060c0610b4084828501610a5d565b60c08301525092915050565b60006101208284031215610b5f57600080fd5b610b6960606111d1565b90506000610b7984828501610a9c565b60008301525060e082013567ffffffffffffffff811115610b9957600080fd5b610ba584828501610a72565b602083015250610100610bba84828501610bc6565b60408301525092915050565b600081359050610bd5816116d5565b92915050565b600060208284031215610bed57600080fd5b6000610bfb84828501610a5d565b91505092915050565b60008060408385031215610c1757600080fd5b600083013567ffffffffffffffff811115610c3157600080fd5b610c3d85828601610b4c565b925050602083013567ffffffffffffffff811115610c5a57600080fd5b610c6685828601610a72565b9150509250929050565b610c7981611335565b82525050565b610c90610c8b82611335565b6114a9565b82525050565b610c9f81611347565b82525050565b610cae81611353565b82525050565b6000610cbf82611227565b610cc98185611232565b9350610cd9818560208601611445565b80840191505092915050565b610cee81611394565b82525050565b610cfd816113a6565b82525050565b610d0c816113b8565b82525050565b610d1b816113ca565b82525050565b610d2a816113dc565b82525050565b610d39816113ee565b82525050565b610d4881611400565b82525050565b610d5781611412565b82525050565b610d6681611424565b82525050565b6000610d7960148361123d565b9150610d8482611553565b602082019050919050565b6000610d9c60168361123d565b9150610da78261157c565b602082019050919050565b6000610dbf602b8361123d565b9150610dca826115a5565b604082019050919050565b6000610de260158361123d565b9150610ded826115f4565b602082019050919050565b6000610e0560298361123d565b9150610e108261161d565b604082019050919050565b6000610e2860128361123d565b9150610e338261166c565b602082019050919050565b6000610e4b601f8361123d565b9150610e5682611695565b602082019050919050565b610e6a8161137d565b82525050565b610e81610e7c8261137d565b6114cd565b82525050565b610e9081611387565b82525050565b6000610ea28285610cb4565b9150610eae8284610e70565b6020820191508190509392505050565b6000610eca828a610e70565b602082019150610eda8289610c7f565b601482019150610eea8288610e70565b602082019150610efa8287610c7f565b601482019150610f0a8286610c7f565b601482019150610f1a8285610c7f565b601482019150610f2a8284610c7f565b60148201915081905098975050505050505050565b6000602082019050610f546000830184610c70565b92915050565b6000604082019050610f6f6000830185610c70565b610f7c6020830184610e61565b9392505050565b6000602082019050610f986000830184610c96565b92915050565b6000608082019050610fb36000830187610ca5565b610fc06020830186610e87565b610fcd6040830185610ca5565b610fda6060830184610ca5565b95945050505050565b6000602082019050610ff86000830184610ce5565b92915050565b60006020820190506110136000830184610cf4565b92915050565b600060208201905061102e6000830184610d03565b92915050565b60006020820190506110496000830184610d12565b92915050565b60006020820190506110646000830184610d21565b92915050565b600060208201905061107f6000830184610d30565b92915050565b600060208201905061109a6000830184610d3f565b92915050565b60006020820190506110b56000830184610d4e565b92915050565b60006020820190506110d06000830184610d5d565b92915050565b600060208201905081810360008301526110ef81610d6c565b9050919050565b6000602082019050818103600083015261110f81610d8f565b9050919050565b6000602082019050818103600083015261112f81610db2565b9050919050565b6000602082019050818103600083015261114f81610dd5565b9050919050565b6000602082019050818103600083015261116f81610df8565b9050919050565b6000602082019050818103600083015261118f81610e1b565b9050919050565b600060208201905081810360008301526111af81610e3e565b9050919050565b60006020820190506111cb6000830184610e61565b92915050565b60006111db6111ec565b90506111e78282611478565b919050565b6000604051905090565b600067ffffffffffffffff82111561121157611210611506565b5b61121a82611535565b9050602081019050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b60006112598261137d565b91506112648361137d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611299576112986114d7565b5b828201905092915050565b60006112af82611387565b91506112ba83611387565b92508260ff038211156112d0576112cf6114d7565b5b828201905092915050565b60006112e68261137d565b91506112f18361137d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561132a576113296114d7565b5b828202905092915050565b60006113408261135d565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b600061139f8261137d565b9050919050565b60006113b18261137d565b9050919050565b60006113c38261137d565b9050919050565b60006113d58261137d565b9050919050565b60006113e78261137d565b9050919050565b60006113f98261137d565b9050919050565b600061140b8261137d565b9050919050565b600061141d8261137d565b9050919050565b600061142f8261137d565b9050919050565b82818337600083830152505050565b60005b83811015611463578082015181840152602081019050611448565b83811115611472576000848401525b50505050565b61148182611535565b810181811067ffffffffffffffff821117156114a05761149f611506565b5b80604052505050565b60006114b4826114bb565b9050919050565b60006114c682611546565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f6368657175652e6e6f6e636520746f6f206f6c64000000000000000000000000600082015250565b7f636f6e74726163742061646472657373206572726f7200000000000000000000600082015250565b7f70617976616c75652073686f756c64206e6f74206578636565642076616c756560008201527f206f66206368657175652e000000000000000000000000000000000000000000602082015250565b7f696c6c6567616c20706179636865717565207369670000000000000000000000600082015250565b7f6f70657261746f722073686f756c64206265206f776e6572206f66207468697360008201527f20636f6e74726163740000000000000000000000000000000000000000000000602082015250565b7f696c6c6567616c20636865717565207369670000000000000000000000000000600082015250565b7f63616c6c6572207368756f756c64206265206368657175652e746f4164647200600082015250565b6116c781611335565b81146116d257600080fd5b50565b6116de8161137d565b81146116e957600080fd5b5056fea26469706673582212207c3c2002916235cf405991b74877df781e63f81c7d6047364a57f20a6474471d64736f6c63430008020033",
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

// CashFlagIterator is returned from FilterFlag and is used to iterate over the raw logs and unpacked data for Flag events raised by the Cash contract.
type CashFlagIterator struct {
	Event *CashFlag // Event containing the contract specifics and raw log

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
func (it *CashFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashFlag)
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
		it.Event = new(CashFlag)
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
func (it *CashFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashFlag represents a Flag event raised by the Cash contract.
type CashFlag struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFlag is a free log retrieval operation binding the contract event 0x7fd5fca7bc379fa5c0330db55f9b725ddbe03460514a6fa71918966f3848779f.
//
// Solidity: event Flag(uint256 arg0)
func (_Cash *CashFilterer) FilterFlag(opts *bind.FilterOpts) (*CashFlagIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "Flag")
	if err != nil {
		return nil, err
	}
	return &CashFlagIterator{contract: _Cash.contract, event: "Flag", logs: logs, sub: sub}, nil
}

// WatchFlag is a free log subscription operation binding the contract event 0x7fd5fca7bc379fa5c0330db55f9b725ddbe03460514a6fa71918966f3848779f.
//
// Solidity: event Flag(uint256 arg0)
func (_Cash *CashFilterer) WatchFlag(opts *bind.WatchOpts, sink chan<- *CashFlag) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "Flag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashFlag)
				if err := _Cash.contract.UnpackLog(event, "Flag", log); err != nil {
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

// ParseFlag is a log parse operation binding the contract event 0x7fd5fca7bc379fa5c0330db55f9b725ddbe03460514a6fa71918966f3848779f.
//
// Solidity: event Flag(uint256 arg0)
func (_Cash *CashFilterer) ParseFlag(log types.Log) (*CashFlag, error) {
	event := new(CashFlag)
	if err := _Cash.contract.UnpackLog(event, "Flag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
