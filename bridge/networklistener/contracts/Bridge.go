// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userIdentifier\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userContractAddress\",\"type\":\"address\"}],\"name\":\"addUserContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userIdentifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506106d5806100616000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631cb159bc1461003b578063467fba0f14610057575b600080fd5b610055600480360381019061005091906103ff565b610073565b005b610071600480360381019061006c919061045f565b610131565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100cd57600080fd5b80600084846040516100e092919061051f565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b600073ffffffffffffffffffffffffffffffffffffffff166000858560405161015b92919061051f565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036101e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d790610595565b60405180910390fd5b600082826040516024016101f59291906105f3565b6040516020818303038152906040527f469c8110000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050600080868660405161028792919061051f565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826040516102da9190610688565b6000604051808303816000865af19150503d8060008114610317576040519150601f19603f3d011682016040523d82523d6000602084013e61031c565b606091505b505090508061032a57600080fd5b505050505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126103615761036061033c565b5b8235905067ffffffffffffffff81111561037e5761037d610341565b5b60208301915083600182028301111561039a57610399610346565b5b9250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103cc826103a1565b9050919050565b6103dc816103c1565b81146103e757600080fd5b50565b6000813590506103f9816103d3565b92915050565b60008060006040848603121561041857610417610332565b5b600084013567ffffffffffffffff81111561043657610435610337565b5b6104428682870161034b565b93509350506020610455868287016103ea565b9150509250925092565b6000806000806040858703121561047957610478610332565b5b600085013567ffffffffffffffff81111561049757610496610337565b5b6104a38782880161034b565b9450945050602085013567ffffffffffffffff8111156104c6576104c5610337565b5b6104d28782880161034b565b925092505092959194509250565b600081905092915050565b82818337600083830152505050565b600061050683856104e0565b93506105138385846104eb565b82840190509392505050565b600061052c8284866104fa565b91508190509392505050565b600082825260208201905092915050565b7f55736572206973206e6f74207265676973746572656400000000000000000000600082015250565b600061057f601683610538565b915061058a82610549565b602082019050919050565b600060208201905081810360008301526105ae81610572565b9050919050565b6000601f19601f8301169050919050565b60006105d28385610538565b93506105df8385846104eb565b6105e8836105b5565b840190509392505050565b6000602082019050818103600083015261060e8184866105c6565b90509392505050565b600081519050919050565b600081905092915050565b60005b8381101561064b578082015181840152602081019050610630565b60008484015250505050565b600061066282610617565b61066c8185610622565b935061067c81856020860161062d565b80840191505092915050565b60006106948284610657565b91508190509291505056fea264697066735822122051a3dee87a91cab7a5456e02153edf98775c483de89e94a4e384e881b09c3a4e64736f6c63430008130033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// AddUserContract is a paid mutator transaction binding the contract method 0x1cb159bc.
//
// Solidity: function addUserContract(string userIdentifier, address userContractAddress) returns()
func (_Bridge *BridgeTransactor) AddUserContract(opts *bind.TransactOpts, userIdentifier string, userContractAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addUserContract", userIdentifier, userContractAddress)
}

// AddUserContract is a paid mutator transaction binding the contract method 0x1cb159bc.
//
// Solidity: function addUserContract(string userIdentifier, address userContractAddress) returns()
func (_Bridge *BridgeSession) AddUserContract(userIdentifier string, userContractAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddUserContract(&_Bridge.TransactOpts, userIdentifier, userContractAddress)
}

// AddUserContract is a paid mutator transaction binding the contract method 0x1cb159bc.
//
// Solidity: function addUserContract(string userIdentifier, address userContractAddress) returns()
func (_Bridge *BridgeTransactorSession) AddUserContract(userIdentifier string, userContractAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddUserContract(&_Bridge.TransactOpts, userIdentifier, userContractAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x467fba0f.
//
// Solidity: function sendMessage(string userIdentifier, string message) returns()
func (_Bridge *BridgeTransactor) SendMessage(opts *bind.TransactOpts, userIdentifier string, message string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "sendMessage", userIdentifier, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x467fba0f.
//
// Solidity: function sendMessage(string userIdentifier, string message) returns()
func (_Bridge *BridgeSession) SendMessage(userIdentifier string, message string) (*types.Transaction, error) {
	return _Bridge.Contract.SendMessage(&_Bridge.TransactOpts, userIdentifier, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x467fba0f.
//
// Solidity: function sendMessage(string userIdentifier, string message) returns()
func (_Bridge *BridgeTransactorSession) SendMessage(userIdentifier string, message string) (*types.Transaction, error) {
	return _Bridge.Contract.SendMessage(&_Bridge.TransactOpts, userIdentifier, message)
}
