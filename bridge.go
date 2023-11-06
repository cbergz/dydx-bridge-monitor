// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"accAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Bridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"DelegatedPowerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATE_BY_TYPE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DYDX_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_VERSION\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_nextAvailableBridgeId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_propositionPowerDelegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_propositionPowerSnapshots\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_propositionPowerSnapshotsCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_votingDelegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_votingSnapshots\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_votingSnapshotsCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"accAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"memo\",\"type\":\"bytes\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"delegateByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateByTypeBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getDelegateeByType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getPowerAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getPowerCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

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
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeCaller) DELEGATEBYTYPETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "DELEGATE_BY_TYPE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeSession) DELEGATEBYTYPETYPEHASH() ([32]byte, error) {
	return _Bridge.Contract.DELEGATEBYTYPETYPEHASH(&_Bridge.CallOpts)
}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DELEGATEBYTYPETYPEHASH() ([32]byte, error) {
	return _Bridge.Contract.DELEGATEBYTYPETYPEHASH(&_Bridge.CallOpts)
}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeCaller) DELEGATETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "DELEGATE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeSession) DELEGATETYPEHASH() ([32]byte, error) {
	return _Bridge.Contract.DELEGATETYPEHASH(&_Bridge.CallOpts)
}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DELEGATETYPEHASH() ([32]byte, error) {
	return _Bridge.Contract.DELEGATETYPEHASH(&_Bridge.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bridge *BridgeCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bridge *BridgeSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Bridge.Contract.DOMAINSEPARATOR(&_Bridge.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Bridge.Contract.DOMAINSEPARATOR(&_Bridge.CallOpts)
}

// DYDXTOKEN is a free data retrieval call binding the contract method 0x3257a4a1.
//
// Solidity: function DYDX_TOKEN() view returns(address)
func (_Bridge *BridgeCaller) DYDXTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "DYDX_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DYDXTOKEN is a free data retrieval call binding the contract method 0x3257a4a1.
//
// Solidity: function DYDX_TOKEN() view returns(address)
func (_Bridge *BridgeSession) DYDXTOKEN() (common.Address, error) {
	return _Bridge.Contract.DYDXTOKEN(&_Bridge.CallOpts)
}

// DYDXTOKEN is a free data retrieval call binding the contract method 0x3257a4a1.
//
// Solidity: function DYDX_TOKEN() view returns(address)
func (_Bridge *BridgeCallerSession) DYDXTOKEN() (common.Address, error) {
	return _Bridge.Contract.DYDXTOKEN(&_Bridge.CallOpts)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(bytes32)
func (_Bridge *BridgeCaller) EIP712DOMAIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "EIP712_DOMAIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(bytes32)
func (_Bridge *BridgeSession) EIP712DOMAIN() ([32]byte, error) {
	return _Bridge.Contract.EIP712DOMAIN(&_Bridge.CallOpts)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(bytes32)
func (_Bridge *BridgeCallerSession) EIP712DOMAIN() ([32]byte, error) {
	return _Bridge.Contract.EIP712DOMAIN(&_Bridge.CallOpts)
}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(bytes)
func (_Bridge *BridgeCaller) EIP712VERSION(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "EIP712_VERSION")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(bytes)
func (_Bridge *BridgeSession) EIP712VERSION() ([]byte, error) {
	return _Bridge.Contract.EIP712VERSION(&_Bridge.CallOpts)
}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(bytes)
func (_Bridge *BridgeCallerSession) EIP712VERSION() ([]byte, error) {
	return _Bridge.Contract.EIP712VERSION(&_Bridge.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Bridge.Contract.PERMITTYPEHASH(&_Bridge.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Bridge.Contract.PERMITTYPEHASH(&_Bridge.CallOpts)
}

// NextAvailableBridgeId is a free data retrieval call binding the contract method 0x7bc1365b.
//
// Solidity: function _nextAvailableBridgeId() view returns(uint256)
func (_Bridge *BridgeCaller) NextAvailableBridgeId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_nextAvailableBridgeId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextAvailableBridgeId is a free data retrieval call binding the contract method 0x7bc1365b.
//
// Solidity: function _nextAvailableBridgeId() view returns(uint256)
func (_Bridge *BridgeSession) NextAvailableBridgeId() (*big.Int, error) {
	return _Bridge.Contract.NextAvailableBridgeId(&_Bridge.CallOpts)
}

// NextAvailableBridgeId is a free data retrieval call binding the contract method 0x7bc1365b.
//
// Solidity: function _nextAvailableBridgeId() view returns(uint256)
func (_Bridge *BridgeCallerSession) NextAvailableBridgeId() (*big.Int, error) {
	return _Bridge.Contract.NextAvailableBridgeId(&_Bridge.CallOpts)
}

// PropositionPowerDelegates is a free data retrieval call binding the contract method 0xf0eb7f8e.
//
// Solidity: function _propositionPowerDelegates(address ) view returns(address)
func (_Bridge *BridgeCaller) PropositionPowerDelegates(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_propositionPowerDelegates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PropositionPowerDelegates is a free data retrieval call binding the contract method 0xf0eb7f8e.
//
// Solidity: function _propositionPowerDelegates(address ) view returns(address)
func (_Bridge *BridgeSession) PropositionPowerDelegates(arg0 common.Address) (common.Address, error) {
	return _Bridge.Contract.PropositionPowerDelegates(&_Bridge.CallOpts, arg0)
}

// PropositionPowerDelegates is a free data retrieval call binding the contract method 0xf0eb7f8e.
//
// Solidity: function _propositionPowerDelegates(address ) view returns(address)
func (_Bridge *BridgeCallerSession) PropositionPowerDelegates(arg0 common.Address) (common.Address, error) {
	return _Bridge.Contract.PropositionPowerDelegates(&_Bridge.CallOpts, arg0)
}

// PropositionPowerSnapshots is a free data retrieval call binding the contract method 0x5f3b6876.
//
// Solidity: function _propositionPowerSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Bridge *BridgeCaller) PropositionPowerSnapshots(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_propositionPowerSnapshots", arg0, arg1)

	outstruct := new(struct {
		BlockNumber *big.Int
		Value       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PropositionPowerSnapshots is a free data retrieval call binding the contract method 0x5f3b6876.
//
// Solidity: function _propositionPowerSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Bridge *BridgeSession) PropositionPowerSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Bridge.Contract.PropositionPowerSnapshots(&_Bridge.CallOpts, arg0, arg1)
}

// PropositionPowerSnapshots is a free data retrieval call binding the contract method 0x5f3b6876.
//
// Solidity: function _propositionPowerSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Bridge *BridgeCallerSession) PropositionPowerSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Bridge.Contract.PropositionPowerSnapshots(&_Bridge.CallOpts, arg0, arg1)
}

// PropositionPowerSnapshotsCounts is a free data retrieval call binding the contract method 0x549aa8a3.
//
// Solidity: function _propositionPowerSnapshotsCounts(address ) view returns(uint256)
func (_Bridge *BridgeCaller) PropositionPowerSnapshotsCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_propositionPowerSnapshotsCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PropositionPowerSnapshotsCounts is a free data retrieval call binding the contract method 0x549aa8a3.
//
// Solidity: function _propositionPowerSnapshotsCounts(address ) view returns(uint256)
func (_Bridge *BridgeSession) PropositionPowerSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.PropositionPowerSnapshotsCounts(&_Bridge.CallOpts, arg0)
}

// PropositionPowerSnapshotsCounts is a free data retrieval call binding the contract method 0x549aa8a3.
//
// Solidity: function _propositionPowerSnapshotsCounts(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) PropositionPowerSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.PropositionPowerSnapshotsCounts(&_Bridge.CallOpts, arg0)
}

// VotingDelegates is a free data retrieval call binding the contract method 0x537f215c.
//
// Solidity: function _votingDelegates(address ) view returns(address)
func (_Bridge *BridgeCaller) VotingDelegates(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_votingDelegates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingDelegates is a free data retrieval call binding the contract method 0x537f215c.
//
// Solidity: function _votingDelegates(address ) view returns(address)
func (_Bridge *BridgeSession) VotingDelegates(arg0 common.Address) (common.Address, error) {
	return _Bridge.Contract.VotingDelegates(&_Bridge.CallOpts, arg0)
}

// VotingDelegates is a free data retrieval call binding the contract method 0x537f215c.
//
// Solidity: function _votingDelegates(address ) view returns(address)
func (_Bridge *BridgeCallerSession) VotingDelegates(arg0 common.Address) (common.Address, error) {
	return _Bridge.Contract.VotingDelegates(&_Bridge.CallOpts, arg0)
}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Bridge *BridgeCaller) VotingSnapshots(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_votingSnapshots", arg0, arg1)

	outstruct := new(struct {
		BlockNumber *big.Int
		Value       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Bridge *BridgeSession) VotingSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Bridge.Contract.VotingSnapshots(&_Bridge.CallOpts, arg0, arg1)
}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Bridge *BridgeCallerSession) VotingSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Bridge.Contract.VotingSnapshots(&_Bridge.CallOpts, arg0, arg1)
}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_Bridge *BridgeCaller) VotingSnapshotsCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_votingSnapshotsCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_Bridge *BridgeSession) VotingSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.VotingSnapshotsCounts(&_Bridge.CallOpts, arg0)
}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) VotingSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.VotingSnapshotsCounts(&_Bridge.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bridge *BridgeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bridge *BridgeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bridge.Contract.Allowance(&_Bridge.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bridge *BridgeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bridge.Contract.Allowance(&_Bridge.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bridge *BridgeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bridge *BridgeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bridge.Contract.BalanceOf(&_Bridge.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bridge *BridgeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bridge.Contract.BalanceOf(&_Bridge.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bridge *BridgeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bridge *BridgeSession) Decimals() (uint8, error) {
	return _Bridge.Contract.Decimals(&_Bridge.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bridge *BridgeCallerSession) Decimals() (uint8, error) {
	return _Bridge.Contract.Decimals(&_Bridge.CallOpts)
}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_Bridge *BridgeCaller) GetDelegateeByType(opts *bind.CallOpts, delegator common.Address, delegationType uint8) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getDelegateeByType", delegator, delegationType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_Bridge *BridgeSession) GetDelegateeByType(delegator common.Address, delegationType uint8) (common.Address, error) {
	return _Bridge.Contract.GetDelegateeByType(&_Bridge.CallOpts, delegator, delegationType)
}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_Bridge *BridgeCallerSession) GetDelegateeByType(delegator common.Address, delegationType uint8) (common.Address, error) {
	return _Bridge.Contract.GetDelegateeByType(&_Bridge.CallOpts, delegator, delegationType)
}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_Bridge *BridgeCaller) GetPowerAtBlock(opts *bind.CallOpts, user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getPowerAtBlock", user, blockNumber, delegationType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_Bridge *BridgeSession) GetPowerAtBlock(user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	return _Bridge.Contract.GetPowerAtBlock(&_Bridge.CallOpts, user, blockNumber, delegationType)
}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_Bridge *BridgeCallerSession) GetPowerAtBlock(user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	return _Bridge.Contract.GetPowerAtBlock(&_Bridge.CallOpts, user, blockNumber, delegationType)
}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_Bridge *BridgeCaller) GetPowerCurrent(opts *bind.CallOpts, user common.Address, delegationType uint8) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getPowerCurrent", user, delegationType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_Bridge *BridgeSession) GetPowerCurrent(user common.Address, delegationType uint8) (*big.Int, error) {
	return _Bridge.Contract.GetPowerCurrent(&_Bridge.CallOpts, user, delegationType)
}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_Bridge *BridgeCallerSession) GetPowerCurrent(user common.Address, delegationType uint8) (*big.Int, error) {
	return _Bridge.Contract.GetPowerCurrent(&_Bridge.CallOpts, user, delegationType)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bridge *BridgeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bridge *BridgeSession) Name() (string, error) {
	return _Bridge.Contract.Name(&_Bridge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bridge *BridgeCallerSession) Name() (string, error) {
	return _Bridge.Contract.Name(&_Bridge.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Bridge *BridgeCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Bridge *BridgeSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Bridge.Contract.Nonces(&_Bridge.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Bridge *BridgeCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Bridge.Contract.Nonces(&_Bridge.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bridge *BridgeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bridge *BridgeSession) Symbol() (string, error) {
	return _Bridge.Contract.Symbol(&_Bridge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bridge *BridgeCallerSession) Symbol() (string, error) {
	return _Bridge.Contract.Symbol(&_Bridge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bridge *BridgeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bridge *BridgeSession) TotalSupply() (*big.Int, error) {
	return _Bridge.Contract.TotalSupply(&_Bridge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bridge *BridgeCallerSession) TotalSupply() (*big.Int, error) {
	return _Bridge.Contract.TotalSupply(&_Bridge.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bridge *BridgeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bridge *BridgeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Approve(&_Bridge.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bridge *BridgeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Approve(&_Bridge.TransactOpts, spender, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0x1d45e29c.
//
// Solidity: function bridge(uint256 amount, bytes accAddress, bytes memo) returns()
func (_Bridge *BridgeTransactor) Bridge(opts *bind.TransactOpts, amount *big.Int, accAddress []byte, memo []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "bridge", amount, accAddress, memo)
}

// Bridge is a paid mutator transaction binding the contract method 0x1d45e29c.
//
// Solidity: function bridge(uint256 amount, bytes accAddress, bytes memo) returns()
func (_Bridge *BridgeSession) Bridge(amount *big.Int, accAddress []byte, memo []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Bridge(&_Bridge.TransactOpts, amount, accAddress, memo)
}

// Bridge is a paid mutator transaction binding the contract method 0x1d45e29c.
//
// Solidity: function bridge(uint256 amount, bytes accAddress, bytes memo) returns()
func (_Bridge *BridgeTransactorSession) Bridge(amount *big.Int, accAddress []byte, memo []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Bridge(&_Bridge.TransactOpts, amount, accAddress, memo)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bridge *BridgeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bridge *BridgeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DecreaseAllowance(&_Bridge.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bridge *BridgeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DecreaseAllowance(&_Bridge.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Bridge *BridgeTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Bridge *BridgeSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Delegate(&_Bridge.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Bridge *BridgeTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Delegate(&_Bridge.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bridge *BridgeTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bridge *BridgeSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.DelegateBySig(&_Bridge.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bridge *BridgeTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.DelegateBySig(&_Bridge.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_Bridge *BridgeTransactor) DelegateByType(opts *bind.TransactOpts, delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "delegateByType", delegatee, delegationType)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_Bridge *BridgeSession) DelegateByType(delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _Bridge.Contract.DelegateByType(&_Bridge.TransactOpts, delegatee, delegationType)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_Bridge *BridgeTransactorSession) DelegateByType(delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _Bridge.Contract.DelegateByType(&_Bridge.TransactOpts, delegatee, delegationType)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bridge *BridgeTransactor) DelegateByTypeBySig(opts *bind.TransactOpts, delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "delegateByTypeBySig", delegatee, delegationType, nonce, expiry, v, r, s)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bridge *BridgeSession) DelegateByTypeBySig(delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.DelegateByTypeBySig(&_Bridge.TransactOpts, delegatee, delegationType, nonce, expiry, v, r, s)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bridge *BridgeTransactorSession) DelegateByTypeBySig(delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.DelegateByTypeBySig(&_Bridge.TransactOpts, delegatee, delegationType, nonce, expiry, v, r, s)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bridge *BridgeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bridge *BridgeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.IncreaseAllowance(&_Bridge.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bridge *BridgeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.IncreaseAllowance(&_Bridge.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bridge *BridgeTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bridge *BridgeSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Permit(&_Bridge.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bridge *BridgeTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Permit(&_Bridge.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Bridge *BridgeTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Bridge *BridgeSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Transfer(&_Bridge.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Bridge *BridgeTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Transfer(&_Bridge.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Bridge *BridgeTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Bridge *BridgeSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TransferFrom(&_Bridge.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Bridge *BridgeTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TransferFrom(&_Bridge.TransactOpts, sender, recipient, amount)
}

// BridgeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bridge contract.
type BridgeApprovalIterator struct {
	Event *BridgeApproval // Event containing the contract specifics and raw log

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
func (it *BridgeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeApproval)
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
		it.Event = new(BridgeApproval)
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
func (it *BridgeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeApproval represents a Approval event raised by the Bridge contract.
type BridgeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bridge *BridgeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BridgeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeApprovalIterator{contract: _Bridge.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bridge *BridgeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BridgeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeApproval)
				if err := _Bridge.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bridge *BridgeFilterer) ParseApproval(log types.Log) (*BridgeApproval, error) {
	event := new(BridgeApproval)
	if err := _Bridge.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBridgeIterator is returned from FilterBridge and is used to iterate over the raw logs and unpacked data for Bridge events raised by the Bridge contract.
type BridgeBridgeIterator struct {
	Event *BridgeBridge // Event containing the contract specifics and raw log

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
func (it *BridgeBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBridge)
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
		it.Event = new(BridgeBridge)
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
func (it *BridgeBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBridge represents a Bridge event raised by the Bridge contract.
type BridgeBridge struct {
	Id         *big.Int
	Amount     *big.Int
	From       common.Address
	AccAddress []byte
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBridge is a free log retrieval operation binding the contract event 0x498a04382650bc110983392ed12ab27595af8ece270a344fc70d773d2481043a.
//
// Solidity: event Bridge(uint256 indexed id, uint256 amount, address from, bytes accAddress, bytes data)
func (_Bridge *BridgeFilterer) FilterBridge(opts *bind.FilterOpts, id []*big.Int) (*BridgeBridgeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Bridge", idRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBridgeIterator{contract: _Bridge.contract, event: "Bridge", logs: logs, sub: sub}, nil
}

// WatchBridge is a free log subscription operation binding the contract event 0x498a04382650bc110983392ed12ab27595af8ece270a344fc70d773d2481043a.
//
// Solidity: event Bridge(uint256 indexed id, uint256 amount, address from, bytes accAddress, bytes data)
func (_Bridge *BridgeFilterer) WatchBridge(opts *bind.WatchOpts, sink chan<- *BridgeBridge, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Bridge", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBridge)
				if err := _Bridge.contract.UnpackLog(event, "Bridge", log); err != nil {
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

// ParseBridge is a log parse operation binding the contract event 0x498a04382650bc110983392ed12ab27595af8ece270a344fc70d773d2481043a.
//
// Solidity: event Bridge(uint256 indexed id, uint256 amount, address from, bytes accAddress, bytes data)
func (_Bridge *BridgeFilterer) ParseBridge(log types.Log) (*BridgeBridge, error) {
	event := new(BridgeBridge)
	if err := _Bridge.contract.UnpackLog(event, "Bridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the Bridge contract.
type BridgeDelegateChangedIterator struct {
	Event *BridgeDelegateChanged // Event containing the contract specifics and raw log

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
func (it *BridgeDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDelegateChanged)
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
		it.Event = new(BridgeDelegateChanged)
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
func (it *BridgeDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDelegateChanged represents a DelegateChanged event raised by the Bridge contract.
type BridgeDelegateChanged struct {
	Delegator      common.Address
	Delegatee      common.Address
	DelegationType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_Bridge *BridgeFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, delegatee []common.Address) (*BridgeDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDelegateChangedIterator{contract: _Bridge.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_Bridge *BridgeFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *BridgeDelegateChanged, delegator []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDelegateChanged)
				if err := _Bridge.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_Bridge *BridgeFilterer) ParseDelegateChanged(log types.Log) (*BridgeDelegateChanged, error) {
	event := new(BridgeDelegateChanged)
	if err := _Bridge.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDelegatedPowerChangedIterator is returned from FilterDelegatedPowerChanged and is used to iterate over the raw logs and unpacked data for DelegatedPowerChanged events raised by the Bridge contract.
type BridgeDelegatedPowerChangedIterator struct {
	Event *BridgeDelegatedPowerChanged // Event containing the contract specifics and raw log

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
func (it *BridgeDelegatedPowerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDelegatedPowerChanged)
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
		it.Event = new(BridgeDelegatedPowerChanged)
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
func (it *BridgeDelegatedPowerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDelegatedPowerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDelegatedPowerChanged represents a DelegatedPowerChanged event raised by the Bridge contract.
type BridgeDelegatedPowerChanged struct {
	User           common.Address
	Amount         *big.Int
	DelegationType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegatedPowerChanged is a free log retrieval operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_Bridge *BridgeFilterer) FilterDelegatedPowerChanged(opts *bind.FilterOpts, user []common.Address) (*BridgeDelegatedPowerChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DelegatedPowerChanged", userRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDelegatedPowerChangedIterator{contract: _Bridge.contract, event: "DelegatedPowerChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatedPowerChanged is a free log subscription operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_Bridge *BridgeFilterer) WatchDelegatedPowerChanged(opts *bind.WatchOpts, sink chan<- *BridgeDelegatedPowerChanged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DelegatedPowerChanged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDelegatedPowerChanged)
				if err := _Bridge.contract.UnpackLog(event, "DelegatedPowerChanged", log); err != nil {
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

// ParseDelegatedPowerChanged is a log parse operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_Bridge *BridgeFilterer) ParseDelegatedPowerChanged(log types.Log) (*BridgeDelegatedPowerChanged, error) {
	event := new(BridgeDelegatedPowerChanged)
	if err := _Bridge.contract.UnpackLog(event, "DelegatedPowerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bridge contract.
type BridgeTransferIterator struct {
	Event *BridgeTransfer // Event containing the contract specifics and raw log

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
func (it *BridgeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTransfer)
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
		it.Event = new(BridgeTransfer)
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
func (it *BridgeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTransfer represents a Transfer event raised by the Bridge contract.
type BridgeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BridgeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTransferIterator{contract: _Bridge.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BridgeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTransfer)
				if err := _Bridge.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) ParseTransfer(log types.Log) (*BridgeTransfer, error) {
	event := new(BridgeTransfer)
	if err := _Bridge.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
