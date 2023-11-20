// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dydxtoken

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

// DydxtokenMetaData contains all meta data concerning the Dydxtoken contract.
var DydxtokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transfersRestrictedBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transferRestrictionLiftedNoLaterThan\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintingRestrictedBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintMaxPercent\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"DelegatedPowerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"TransferAllowlistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transfersRestrictedBefore\",\"type\":\"uint256\"}],\"name\":\"TransfersRestrictedBeforeUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATE_BY_TYPE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_VERSION\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_MAX_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_MIN_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_RESTRICTION_LIFTED_NO_LATER_THAN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_mintingRestrictedBefore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_propositionPowerDelegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_propositionPowerSnapshots\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_propositionPowerSnapshotsCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenTransferAllowlist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_totalSupplySnapshots\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalSupplySnapshotsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_transfersRestrictedBefore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_votingDelegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_votingSnapshots\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_votingSnapshotsCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addressesToAdd\",\"type\":\"address[]\"}],\"name\":\"addToTokenTransferAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"delegateByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateByTypeBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getDelegateeByType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getPowerAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getPowerCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addressesToRemove\",\"type\":\"address[]\"}],\"name\":\"removeFromTokenTransferAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transfersRestrictedBefore\",\"type\":\"uint256\"}],\"name\":\"updateTransfersRestrictedBefore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DydxtokenABI is the input ABI used to generate the binding from.
// Deprecated: Use DydxtokenMetaData.ABI instead.
var DydxtokenABI = DydxtokenMetaData.ABI

// Dydxtoken is an auto generated Go binding around an Ethereum contract.
type Dydxtoken struct {
	DydxtokenCaller     // Read-only binding to the contract
	DydxtokenTransactor // Write-only binding to the contract
	DydxtokenFilterer   // Log filterer for contract events
}

// DydxtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type DydxtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DydxtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DydxtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DydxtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DydxtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DydxtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DydxtokenSession struct {
	Contract     *Dydxtoken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DydxtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DydxtokenCallerSession struct {
	Contract *DydxtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DydxtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DydxtokenTransactorSession struct {
	Contract     *DydxtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DydxtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type DydxtokenRaw struct {
	Contract *Dydxtoken // Generic contract binding to access the raw methods on
}

// DydxtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DydxtokenCallerRaw struct {
	Contract *DydxtokenCaller // Generic read-only contract binding to access the raw methods on
}

// DydxtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DydxtokenTransactorRaw struct {
	Contract *DydxtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDydxtoken creates a new instance of Dydxtoken, bound to a specific deployed contract.
func NewDydxtoken(address common.Address, backend bind.ContractBackend) (*Dydxtoken, error) {
	contract, err := bindDydxtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dydxtoken{DydxtokenCaller: DydxtokenCaller{contract: contract}, DydxtokenTransactor: DydxtokenTransactor{contract: contract}, DydxtokenFilterer: DydxtokenFilterer{contract: contract}}, nil
}

// NewDydxtokenCaller creates a new read-only instance of Dydxtoken, bound to a specific deployed contract.
func NewDydxtokenCaller(address common.Address, caller bind.ContractCaller) (*DydxtokenCaller, error) {
	contract, err := bindDydxtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DydxtokenCaller{contract: contract}, nil
}

// NewDydxtokenTransactor creates a new write-only instance of Dydxtoken, bound to a specific deployed contract.
func NewDydxtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*DydxtokenTransactor, error) {
	contract, err := bindDydxtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DydxtokenTransactor{contract: contract}, nil
}

// NewDydxtokenFilterer creates a new log filterer instance of Dydxtoken, bound to a specific deployed contract.
func NewDydxtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*DydxtokenFilterer, error) {
	contract, err := bindDydxtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DydxtokenFilterer{contract: contract}, nil
}

// bindDydxtoken binds a generic wrapper to an already deployed contract.
func bindDydxtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DydxtokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dydxtoken *DydxtokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dydxtoken.Contract.DydxtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dydxtoken *DydxtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dydxtoken.Contract.DydxtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dydxtoken *DydxtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dydxtoken.Contract.DydxtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dydxtoken *DydxtokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dydxtoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dydxtoken *DydxtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dydxtoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dydxtoken *DydxtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dydxtoken.Contract.contract.Transact(opts, method, params...)
}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_Dydxtoken *DydxtokenCaller) DELEGATEBYTYPETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "DELEGATE_BY_TYPE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_Dydxtoken *DydxtokenSession) DELEGATEBYTYPETYPEHASH() ([32]byte, error) {
	return _Dydxtoken.Contract.DELEGATEBYTYPETYPEHASH(&_Dydxtoken.CallOpts)
}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_Dydxtoken *DydxtokenCallerSession) DELEGATEBYTYPETYPEHASH() ([32]byte, error) {
	return _Dydxtoken.Contract.DELEGATEBYTYPETYPEHASH(&_Dydxtoken.CallOpts)
}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_Dydxtoken *DydxtokenCaller) DELEGATETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "DELEGATE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_Dydxtoken *DydxtokenSession) DELEGATETYPEHASH() ([32]byte, error) {
	return _Dydxtoken.Contract.DELEGATETYPEHASH(&_Dydxtoken.CallOpts)
}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_Dydxtoken *DydxtokenCallerSession) DELEGATETYPEHASH() ([32]byte, error) {
	return _Dydxtoken.Contract.DELEGATETYPEHASH(&_Dydxtoken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Dydxtoken *DydxtokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Dydxtoken *DydxtokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Dydxtoken.Contract.DOMAINSEPARATOR(&_Dydxtoken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Dydxtoken *DydxtokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Dydxtoken.Contract.DOMAINSEPARATOR(&_Dydxtoken.CallOpts)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(bytes32)
func (_Dydxtoken *DydxtokenCaller) EIP712DOMAIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "EIP712_DOMAIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(bytes32)
func (_Dydxtoken *DydxtokenSession) EIP712DOMAIN() ([32]byte, error) {
	return _Dydxtoken.Contract.EIP712DOMAIN(&_Dydxtoken.CallOpts)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(bytes32)
func (_Dydxtoken *DydxtokenCallerSession) EIP712DOMAIN() ([32]byte, error) {
	return _Dydxtoken.Contract.EIP712DOMAIN(&_Dydxtoken.CallOpts)
}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(bytes)
func (_Dydxtoken *DydxtokenCaller) EIP712VERSION(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "EIP712_VERSION")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(bytes)
func (_Dydxtoken *DydxtokenSession) EIP712VERSION() ([]byte, error) {
	return _Dydxtoken.Contract.EIP712VERSION(&_Dydxtoken.CallOpts)
}

// EIP712VERSION is a free data retrieval call binding the contract method 0xeccec5a8.
//
// Solidity: function EIP712_VERSION() view returns(bytes)
func (_Dydxtoken *DydxtokenCallerSession) EIP712VERSION() ([]byte, error) {
	return _Dydxtoken.Contract.EIP712VERSION(&_Dydxtoken.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "INITIAL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_Dydxtoken *DydxtokenSession) INITIALSUPPLY() (*big.Int, error) {
	return _Dydxtoken.Contract.INITIALSUPPLY(&_Dydxtoken.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _Dydxtoken.Contract.INITIALSUPPLY(&_Dydxtoken.CallOpts)
}

// MINTMAXPERCENT is a free data retrieval call binding the contract method 0x657c7a85.
//
// Solidity: function MINT_MAX_PERCENT() view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) MINTMAXPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "MINT_MAX_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTMAXPERCENT is a free data retrieval call binding the contract method 0x657c7a85.
//
// Solidity: function MINT_MAX_PERCENT() view returns(uint256)
func (_Dydxtoken *DydxtokenSession) MINTMAXPERCENT() (*big.Int, error) {
	return _Dydxtoken.Contract.MINTMAXPERCENT(&_Dydxtoken.CallOpts)
}

// MINTMAXPERCENT is a free data retrieval call binding the contract method 0x657c7a85.
//
// Solidity: function MINT_MAX_PERCENT() view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) MINTMAXPERCENT() (*big.Int, error) {
	return _Dydxtoken.Contract.MINTMAXPERCENT(&_Dydxtoken.CallOpts)
}

// MINTMININTERVAL is a free data retrieval call binding the contract method 0x69e3b0d0.
//
// Solidity: function MINT_MIN_INTERVAL() view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) MINTMININTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "MINT_MIN_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTMININTERVAL is a free data retrieval call binding the contract method 0x69e3b0d0.
//
// Solidity: function MINT_MIN_INTERVAL() view returns(uint256)
func (_Dydxtoken *DydxtokenSession) MINTMININTERVAL() (*big.Int, error) {
	return _Dydxtoken.Contract.MINTMININTERVAL(&_Dydxtoken.CallOpts)
}

// MINTMININTERVAL is a free data retrieval call binding the contract method 0x69e3b0d0.
//
// Solidity: function MINT_MIN_INTERVAL() view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) MINTMININTERVAL() (*big.Int, error) {
	return _Dydxtoken.Contract.MINTMININTERVAL(&_Dydxtoken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Dydxtoken *DydxtokenCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Dydxtoken *DydxtokenSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Dydxtoken.Contract.PERMITTYPEHASH(&_Dydxtoken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Dydxtoken *DydxtokenCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Dydxtoken.Contract.PERMITTYPEHASH(&_Dydxtoken.CallOpts)
}

// TRANSFERRESTRICTIONLIFTEDNOLATERTHAN is a free data retrieval call binding the contract method 0xd17deb9f.
//
// Solidity: function TRANSFER_RESTRICTION_LIFTED_NO_LATER_THAN() view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) TRANSFERRESTRICTIONLIFTEDNOLATERTHAN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "TRANSFER_RESTRICTION_LIFTED_NO_LATER_THAN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TRANSFERRESTRICTIONLIFTEDNOLATERTHAN is a free data retrieval call binding the contract method 0xd17deb9f.
//
// Solidity: function TRANSFER_RESTRICTION_LIFTED_NO_LATER_THAN() view returns(uint256)
func (_Dydxtoken *DydxtokenSession) TRANSFERRESTRICTIONLIFTEDNOLATERTHAN() (*big.Int, error) {
	return _Dydxtoken.Contract.TRANSFERRESTRICTIONLIFTEDNOLATERTHAN(&_Dydxtoken.CallOpts)
}

// TRANSFERRESTRICTIONLIFTEDNOLATERTHAN is a free data retrieval call binding the contract method 0xd17deb9f.
//
// Solidity: function TRANSFER_RESTRICTION_LIFTED_NO_LATER_THAN() view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) TRANSFERRESTRICTIONLIFTEDNOLATERTHAN() (*big.Int, error) {
	return _Dydxtoken.Contract.TRANSFERRESTRICTIONLIFTEDNOLATERTHAN(&_Dydxtoken.CallOpts)
}

// MintingRestrictedBefore is a free data retrieval call binding the contract method 0xcdfeb0c1.
//
// Solidity: function _mintingRestrictedBefore() view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) MintingRestrictedBefore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "_mintingRestrictedBefore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintingRestrictedBefore is a free data retrieval call binding the contract method 0xcdfeb0c1.
//
// Solidity: function _mintingRestrictedBefore() view returns(uint256)
func (_Dydxtoken *DydxtokenSession) MintingRestrictedBefore() (*big.Int, error) {
	return _Dydxtoken.Contract.MintingRestrictedBefore(&_Dydxtoken.CallOpts)
}

// MintingRestrictedBefore is a free data retrieval call binding the contract method 0xcdfeb0c1.
//
// Solidity: function _mintingRestrictedBefore() view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) MintingRestrictedBefore() (*big.Int, error) {
	return _Dydxtoken.Contract.MintingRestrictedBefore(&_Dydxtoken.CallOpts)
}

// PropositionPowerDelegates is a free data retrieval call binding the contract method 0xf0eb7f8e.
//
// Solidity: function _propositionPowerDelegates(address ) view returns(address)
func (_Dydxtoken *DydxtokenCaller) PropositionPowerDelegates(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "_propositionPowerDelegates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PropositionPowerDelegates is a free data retrieval call binding the contract method 0xf0eb7f8e.
//
// Solidity: function _propositionPowerDelegates(address ) view returns(address)
func (_Dydxtoken *DydxtokenSession) PropositionPowerDelegates(arg0 common.Address) (common.Address, error) {
	return _Dydxtoken.Contract.PropositionPowerDelegates(&_Dydxtoken.CallOpts, arg0)
}

// PropositionPowerDelegates is a free data retrieval call binding the contract method 0xf0eb7f8e.
//
// Solidity: function _propositionPowerDelegates(address ) view returns(address)
func (_Dydxtoken *DydxtokenCallerSession) PropositionPowerDelegates(arg0 common.Address) (common.Address, error) {
	return _Dydxtoken.Contract.PropositionPowerDelegates(&_Dydxtoken.CallOpts, arg0)
}

// PropositionPowerSnapshots is a free data retrieval call binding the contract method 0x5f3b6876.
//
// Solidity: function _propositionPowerSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Dydxtoken *DydxtokenCaller) PropositionPowerSnapshots(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "_propositionPowerSnapshots", arg0, arg1)

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
func (_Dydxtoken *DydxtokenSession) PropositionPowerSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Dydxtoken.Contract.PropositionPowerSnapshots(&_Dydxtoken.CallOpts, arg0, arg1)
}

// PropositionPowerSnapshots is a free data retrieval call binding the contract method 0x5f3b6876.
//
// Solidity: function _propositionPowerSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Dydxtoken *DydxtokenCallerSession) PropositionPowerSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Dydxtoken.Contract.PropositionPowerSnapshots(&_Dydxtoken.CallOpts, arg0, arg1)
}

// PropositionPowerSnapshotsCounts is a free data retrieval call binding the contract method 0x549aa8a3.
//
// Solidity: function _propositionPowerSnapshotsCounts(address ) view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) PropositionPowerSnapshotsCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "_propositionPowerSnapshotsCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PropositionPowerSnapshotsCounts is a free data retrieval call binding the contract method 0x549aa8a3.
//
// Solidity: function _propositionPowerSnapshotsCounts(address ) view returns(uint256)
func (_Dydxtoken *DydxtokenSession) PropositionPowerSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _Dydxtoken.Contract.PropositionPowerSnapshotsCounts(&_Dydxtoken.CallOpts, arg0)
}

// PropositionPowerSnapshotsCounts is a free data retrieval call binding the contract method 0x549aa8a3.
//
// Solidity: function _propositionPowerSnapshotsCounts(address ) view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) PropositionPowerSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _Dydxtoken.Contract.PropositionPowerSnapshotsCounts(&_Dydxtoken.CallOpts, arg0)
}

// TokenTransferAllowlist is a free data retrieval call binding the contract method 0x48032ec1.
//
// Solidity: function _tokenTransferAllowlist(address ) view returns(bool)
func (_Dydxtoken *DydxtokenCaller) TokenTransferAllowlist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "_tokenTransferAllowlist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenTransferAllowlist is a free data retrieval call binding the contract method 0x48032ec1.
//
// Solidity: function _tokenTransferAllowlist(address ) view returns(bool)
func (_Dydxtoken *DydxtokenSession) TokenTransferAllowlist(arg0 common.Address) (bool, error) {
	return _Dydxtoken.Contract.TokenTransferAllowlist(&_Dydxtoken.CallOpts, arg0)
}

// TokenTransferAllowlist is a free data retrieval call binding the contract method 0x48032ec1.
//
// Solidity: function _tokenTransferAllowlist(address ) view returns(bool)
func (_Dydxtoken *DydxtokenCallerSession) TokenTransferAllowlist(arg0 common.Address) (bool, error) {
	return _Dydxtoken.Contract.TokenTransferAllowlist(&_Dydxtoken.CallOpts, arg0)
}

// TotalSupplySnapshots is a free data retrieval call binding the contract method 0x13929bbe.
//
// Solidity: function _totalSupplySnapshots(uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Dydxtoken *DydxtokenCaller) TotalSupplySnapshots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "_totalSupplySnapshots", arg0)

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

// TotalSupplySnapshots is a free data retrieval call binding the contract method 0x13929bbe.
//
// Solidity: function _totalSupplySnapshots(uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Dydxtoken *DydxtokenSession) TotalSupplySnapshots(arg0 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Dydxtoken.Contract.TotalSupplySnapshots(&_Dydxtoken.CallOpts, arg0)
}

// TotalSupplySnapshots is a free data retrieval call binding the contract method 0x13929bbe.
//
// Solidity: function _totalSupplySnapshots(uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Dydxtoken *DydxtokenCallerSession) TotalSupplySnapshots(arg0 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Dydxtoken.Contract.TotalSupplySnapshots(&_Dydxtoken.CallOpts, arg0)
}

// TotalSupplySnapshotsCount is a free data retrieval call binding the contract method 0x2b4e1a5b.
//
// Solidity: function _totalSupplySnapshotsCount() view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) TotalSupplySnapshotsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "_totalSupplySnapshotsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplySnapshotsCount is a free data retrieval call binding the contract method 0x2b4e1a5b.
//
// Solidity: function _totalSupplySnapshotsCount() view returns(uint256)
func (_Dydxtoken *DydxtokenSession) TotalSupplySnapshotsCount() (*big.Int, error) {
	return _Dydxtoken.Contract.TotalSupplySnapshotsCount(&_Dydxtoken.CallOpts)
}

// TotalSupplySnapshotsCount is a free data retrieval call binding the contract method 0x2b4e1a5b.
//
// Solidity: function _totalSupplySnapshotsCount() view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) TotalSupplySnapshotsCount() (*big.Int, error) {
	return _Dydxtoken.Contract.TotalSupplySnapshotsCount(&_Dydxtoken.CallOpts)
}

// TransfersRestrictedBefore is a free data retrieval call binding the contract method 0x1ff06fdf.
//
// Solidity: function _transfersRestrictedBefore() view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) TransfersRestrictedBefore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "_transfersRestrictedBefore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransfersRestrictedBefore is a free data retrieval call binding the contract method 0x1ff06fdf.
//
// Solidity: function _transfersRestrictedBefore() view returns(uint256)
func (_Dydxtoken *DydxtokenSession) TransfersRestrictedBefore() (*big.Int, error) {
	return _Dydxtoken.Contract.TransfersRestrictedBefore(&_Dydxtoken.CallOpts)
}

// TransfersRestrictedBefore is a free data retrieval call binding the contract method 0x1ff06fdf.
//
// Solidity: function _transfersRestrictedBefore() view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) TransfersRestrictedBefore() (*big.Int, error) {
	return _Dydxtoken.Contract.TransfersRestrictedBefore(&_Dydxtoken.CallOpts)
}

// VotingDelegates is a free data retrieval call binding the contract method 0x537f215c.
//
// Solidity: function _votingDelegates(address ) view returns(address)
func (_Dydxtoken *DydxtokenCaller) VotingDelegates(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "_votingDelegates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingDelegates is a free data retrieval call binding the contract method 0x537f215c.
//
// Solidity: function _votingDelegates(address ) view returns(address)
func (_Dydxtoken *DydxtokenSession) VotingDelegates(arg0 common.Address) (common.Address, error) {
	return _Dydxtoken.Contract.VotingDelegates(&_Dydxtoken.CallOpts, arg0)
}

// VotingDelegates is a free data retrieval call binding the contract method 0x537f215c.
//
// Solidity: function _votingDelegates(address ) view returns(address)
func (_Dydxtoken *DydxtokenCallerSession) VotingDelegates(arg0 common.Address) (common.Address, error) {
	return _Dydxtoken.Contract.VotingDelegates(&_Dydxtoken.CallOpts, arg0)
}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Dydxtoken *DydxtokenCaller) VotingSnapshots(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "_votingSnapshots", arg0, arg1)

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
func (_Dydxtoken *DydxtokenSession) VotingSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Dydxtoken.Contract.VotingSnapshots(&_Dydxtoken.CallOpts, arg0, arg1)
}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_Dydxtoken *DydxtokenCallerSession) VotingSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _Dydxtoken.Contract.VotingSnapshots(&_Dydxtoken.CallOpts, arg0, arg1)
}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) VotingSnapshotsCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "_votingSnapshotsCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_Dydxtoken *DydxtokenSession) VotingSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _Dydxtoken.Contract.VotingSnapshotsCounts(&_Dydxtoken.CallOpts, arg0)
}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) VotingSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _Dydxtoken.Contract.VotingSnapshotsCounts(&_Dydxtoken.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Dydxtoken *DydxtokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Dydxtoken.Contract.Allowance(&_Dydxtoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Dydxtoken.Contract.Allowance(&_Dydxtoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Dydxtoken *DydxtokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Dydxtoken.Contract.BalanceOf(&_Dydxtoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Dydxtoken.Contract.BalanceOf(&_Dydxtoken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Dydxtoken *DydxtokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Dydxtoken *DydxtokenSession) Decimals() (uint8, error) {
	return _Dydxtoken.Contract.Decimals(&_Dydxtoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Dydxtoken *DydxtokenCallerSession) Decimals() (uint8, error) {
	return _Dydxtoken.Contract.Decimals(&_Dydxtoken.CallOpts)
}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_Dydxtoken *DydxtokenCaller) GetDelegateeByType(opts *bind.CallOpts, delegator common.Address, delegationType uint8) (common.Address, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "getDelegateeByType", delegator, delegationType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_Dydxtoken *DydxtokenSession) GetDelegateeByType(delegator common.Address, delegationType uint8) (common.Address, error) {
	return _Dydxtoken.Contract.GetDelegateeByType(&_Dydxtoken.CallOpts, delegator, delegationType)
}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_Dydxtoken *DydxtokenCallerSession) GetDelegateeByType(delegator common.Address, delegationType uint8) (common.Address, error) {
	return _Dydxtoken.Contract.GetDelegateeByType(&_Dydxtoken.CallOpts, delegator, delegationType)
}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) GetPowerAtBlock(opts *bind.CallOpts, user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "getPowerAtBlock", user, blockNumber, delegationType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_Dydxtoken *DydxtokenSession) GetPowerAtBlock(user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	return _Dydxtoken.Contract.GetPowerAtBlock(&_Dydxtoken.CallOpts, user, blockNumber, delegationType)
}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) GetPowerAtBlock(user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	return _Dydxtoken.Contract.GetPowerAtBlock(&_Dydxtoken.CallOpts, user, blockNumber, delegationType)
}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) GetPowerCurrent(opts *bind.CallOpts, user common.Address, delegationType uint8) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "getPowerCurrent", user, delegationType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_Dydxtoken *DydxtokenSession) GetPowerCurrent(user common.Address, delegationType uint8) (*big.Int, error) {
	return _Dydxtoken.Contract.GetPowerCurrent(&_Dydxtoken.CallOpts, user, delegationType)
}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) GetPowerCurrent(user common.Address, delegationType uint8) (*big.Int, error) {
	return _Dydxtoken.Contract.GetPowerCurrent(&_Dydxtoken.CallOpts, user, delegationType)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dydxtoken *DydxtokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dydxtoken *DydxtokenSession) Name() (string, error) {
	return _Dydxtoken.Contract.Name(&_Dydxtoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dydxtoken *DydxtokenCallerSession) Name() (string, error) {
	return _Dydxtoken.Contract.Name(&_Dydxtoken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Dydxtoken *DydxtokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Dydxtoken.Contract.Nonces(&_Dydxtoken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Dydxtoken.Contract.Nonces(&_Dydxtoken.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dydxtoken *DydxtokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dydxtoken *DydxtokenSession) Owner() (common.Address, error) {
	return _Dydxtoken.Contract.Owner(&_Dydxtoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dydxtoken *DydxtokenCallerSession) Owner() (common.Address, error) {
	return _Dydxtoken.Contract.Owner(&_Dydxtoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dydxtoken *DydxtokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dydxtoken *DydxtokenSession) Symbol() (string, error) {
	return _Dydxtoken.Contract.Symbol(&_Dydxtoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dydxtoken *DydxtokenCallerSession) Symbol() (string, error) {
	return _Dydxtoken.Contract.Symbol(&_Dydxtoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dydxtoken *DydxtokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dydxtoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dydxtoken *DydxtokenSession) TotalSupply() (*big.Int, error) {
	return _Dydxtoken.Contract.TotalSupply(&_Dydxtoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dydxtoken *DydxtokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Dydxtoken.Contract.TotalSupply(&_Dydxtoken.CallOpts)
}

// AddToTokenTransferAllowlist is a paid mutator transaction binding the contract method 0x8aaee234.
//
// Solidity: function addToTokenTransferAllowlist(address[] addressesToAdd) returns()
func (_Dydxtoken *DydxtokenTransactor) AddToTokenTransferAllowlist(opts *bind.TransactOpts, addressesToAdd []common.Address) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "addToTokenTransferAllowlist", addressesToAdd)
}

// AddToTokenTransferAllowlist is a paid mutator transaction binding the contract method 0x8aaee234.
//
// Solidity: function addToTokenTransferAllowlist(address[] addressesToAdd) returns()
func (_Dydxtoken *DydxtokenSession) AddToTokenTransferAllowlist(addressesToAdd []common.Address) (*types.Transaction, error) {
	return _Dydxtoken.Contract.AddToTokenTransferAllowlist(&_Dydxtoken.TransactOpts, addressesToAdd)
}

// AddToTokenTransferAllowlist is a paid mutator transaction binding the contract method 0x8aaee234.
//
// Solidity: function addToTokenTransferAllowlist(address[] addressesToAdd) returns()
func (_Dydxtoken *DydxtokenTransactorSession) AddToTokenTransferAllowlist(addressesToAdd []common.Address) (*types.Transaction, error) {
	return _Dydxtoken.Contract.AddToTokenTransferAllowlist(&_Dydxtoken.TransactOpts, addressesToAdd)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Dydxtoken *DydxtokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Dydxtoken *DydxtokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.Approve(&_Dydxtoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Dydxtoken *DydxtokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.Approve(&_Dydxtoken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Dydxtoken *DydxtokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Dydxtoken *DydxtokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.DecreaseAllowance(&_Dydxtoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Dydxtoken *DydxtokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.DecreaseAllowance(&_Dydxtoken.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Dydxtoken *DydxtokenTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Dydxtoken *DydxtokenSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Dydxtoken.Contract.Delegate(&_Dydxtoken.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Dydxtoken *DydxtokenTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Dydxtoken.Contract.Delegate(&_Dydxtoken.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dydxtoken *DydxtokenTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dydxtoken *DydxtokenSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dydxtoken.Contract.DelegateBySig(&_Dydxtoken.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dydxtoken *DydxtokenTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dydxtoken.Contract.DelegateBySig(&_Dydxtoken.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_Dydxtoken *DydxtokenTransactor) DelegateByType(opts *bind.TransactOpts, delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "delegateByType", delegatee, delegationType)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_Dydxtoken *DydxtokenSession) DelegateByType(delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _Dydxtoken.Contract.DelegateByType(&_Dydxtoken.TransactOpts, delegatee, delegationType)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_Dydxtoken *DydxtokenTransactorSession) DelegateByType(delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _Dydxtoken.Contract.DelegateByType(&_Dydxtoken.TransactOpts, delegatee, delegationType)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dydxtoken *DydxtokenTransactor) DelegateByTypeBySig(opts *bind.TransactOpts, delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "delegateByTypeBySig", delegatee, delegationType, nonce, expiry, v, r, s)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dydxtoken *DydxtokenSession) DelegateByTypeBySig(delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dydxtoken.Contract.DelegateByTypeBySig(&_Dydxtoken.TransactOpts, delegatee, delegationType, nonce, expiry, v, r, s)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dydxtoken *DydxtokenTransactorSession) DelegateByTypeBySig(delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dydxtoken.Contract.DelegateByTypeBySig(&_Dydxtoken.TransactOpts, delegatee, delegationType, nonce, expiry, v, r, s)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Dydxtoken *DydxtokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Dydxtoken *DydxtokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.IncreaseAllowance(&_Dydxtoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Dydxtoken *DydxtokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.IncreaseAllowance(&_Dydxtoken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_Dydxtoken *DydxtokenTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "mint", recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_Dydxtoken *DydxtokenSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.Mint(&_Dydxtoken.TransactOpts, recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_Dydxtoken *DydxtokenTransactorSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.Mint(&_Dydxtoken.TransactOpts, recipient, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dydxtoken *DydxtokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dydxtoken *DydxtokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dydxtoken.Contract.Permit(&_Dydxtoken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dydxtoken *DydxtokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dydxtoken.Contract.Permit(&_Dydxtoken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RemoveFromTokenTransferAllowlist is a paid mutator transaction binding the contract method 0x8d48f4f1.
//
// Solidity: function removeFromTokenTransferAllowlist(address[] addressesToRemove) returns()
func (_Dydxtoken *DydxtokenTransactor) RemoveFromTokenTransferAllowlist(opts *bind.TransactOpts, addressesToRemove []common.Address) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "removeFromTokenTransferAllowlist", addressesToRemove)
}

// RemoveFromTokenTransferAllowlist is a paid mutator transaction binding the contract method 0x8d48f4f1.
//
// Solidity: function removeFromTokenTransferAllowlist(address[] addressesToRemove) returns()
func (_Dydxtoken *DydxtokenSession) RemoveFromTokenTransferAllowlist(addressesToRemove []common.Address) (*types.Transaction, error) {
	return _Dydxtoken.Contract.RemoveFromTokenTransferAllowlist(&_Dydxtoken.TransactOpts, addressesToRemove)
}

// RemoveFromTokenTransferAllowlist is a paid mutator transaction binding the contract method 0x8d48f4f1.
//
// Solidity: function removeFromTokenTransferAllowlist(address[] addressesToRemove) returns()
func (_Dydxtoken *DydxtokenTransactorSession) RemoveFromTokenTransferAllowlist(addressesToRemove []common.Address) (*types.Transaction, error) {
	return _Dydxtoken.Contract.RemoveFromTokenTransferAllowlist(&_Dydxtoken.TransactOpts, addressesToRemove)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dydxtoken *DydxtokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dydxtoken *DydxtokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dydxtoken.Contract.RenounceOwnership(&_Dydxtoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dydxtoken *DydxtokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dydxtoken.Contract.RenounceOwnership(&_Dydxtoken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Dydxtoken *DydxtokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Dydxtoken *DydxtokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.Transfer(&_Dydxtoken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Dydxtoken *DydxtokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.Transfer(&_Dydxtoken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Dydxtoken *DydxtokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Dydxtoken *DydxtokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.TransferFrom(&_Dydxtoken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Dydxtoken *DydxtokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.TransferFrom(&_Dydxtoken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dydxtoken *DydxtokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dydxtoken *DydxtokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dydxtoken.Contract.TransferOwnership(&_Dydxtoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dydxtoken *DydxtokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dydxtoken.Contract.TransferOwnership(&_Dydxtoken.TransactOpts, newOwner)
}

// UpdateTransfersRestrictedBefore is a paid mutator transaction binding the contract method 0xad36dafd.
//
// Solidity: function updateTransfersRestrictedBefore(uint256 transfersRestrictedBefore) returns()
func (_Dydxtoken *DydxtokenTransactor) UpdateTransfersRestrictedBefore(opts *bind.TransactOpts, transfersRestrictedBefore *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.contract.Transact(opts, "updateTransfersRestrictedBefore", transfersRestrictedBefore)
}

// UpdateTransfersRestrictedBefore is a paid mutator transaction binding the contract method 0xad36dafd.
//
// Solidity: function updateTransfersRestrictedBefore(uint256 transfersRestrictedBefore) returns()
func (_Dydxtoken *DydxtokenSession) UpdateTransfersRestrictedBefore(transfersRestrictedBefore *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.UpdateTransfersRestrictedBefore(&_Dydxtoken.TransactOpts, transfersRestrictedBefore)
}

// UpdateTransfersRestrictedBefore is a paid mutator transaction binding the contract method 0xad36dafd.
//
// Solidity: function updateTransfersRestrictedBefore(uint256 transfersRestrictedBefore) returns()
func (_Dydxtoken *DydxtokenTransactorSession) UpdateTransfersRestrictedBefore(transfersRestrictedBefore *big.Int) (*types.Transaction, error) {
	return _Dydxtoken.Contract.UpdateTransfersRestrictedBefore(&_Dydxtoken.TransactOpts, transfersRestrictedBefore)
}

// DydxtokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Dydxtoken contract.
type DydxtokenApprovalIterator struct {
	Event *DydxtokenApproval // Event containing the contract specifics and raw log

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
func (it *DydxtokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DydxtokenApproval)
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
		it.Event = new(DydxtokenApproval)
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
func (it *DydxtokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DydxtokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DydxtokenApproval represents a Approval event raised by the Dydxtoken contract.
type DydxtokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Dydxtoken *DydxtokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DydxtokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Dydxtoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DydxtokenApprovalIterator{contract: _Dydxtoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Dydxtoken *DydxtokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DydxtokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Dydxtoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DydxtokenApproval)
				if err := _Dydxtoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Dydxtoken *DydxtokenFilterer) ParseApproval(log types.Log) (*DydxtokenApproval, error) {
	event := new(DydxtokenApproval)
	if err := _Dydxtoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DydxtokenDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the Dydxtoken contract.
type DydxtokenDelegateChangedIterator struct {
	Event *DydxtokenDelegateChanged // Event containing the contract specifics and raw log

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
func (it *DydxtokenDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DydxtokenDelegateChanged)
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
		it.Event = new(DydxtokenDelegateChanged)
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
func (it *DydxtokenDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DydxtokenDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DydxtokenDelegateChanged represents a DelegateChanged event raised by the Dydxtoken contract.
type DydxtokenDelegateChanged struct {
	Delegator      common.Address
	Delegatee      common.Address
	DelegationType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_Dydxtoken *DydxtokenFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, delegatee []common.Address) (*DydxtokenDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Dydxtoken.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &DydxtokenDelegateChangedIterator{contract: _Dydxtoken.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_Dydxtoken *DydxtokenFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *DydxtokenDelegateChanged, delegator []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _Dydxtoken.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DydxtokenDelegateChanged)
				if err := _Dydxtoken.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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
func (_Dydxtoken *DydxtokenFilterer) ParseDelegateChanged(log types.Log) (*DydxtokenDelegateChanged, error) {
	event := new(DydxtokenDelegateChanged)
	if err := _Dydxtoken.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DydxtokenDelegatedPowerChangedIterator is returned from FilterDelegatedPowerChanged and is used to iterate over the raw logs and unpacked data for DelegatedPowerChanged events raised by the Dydxtoken contract.
type DydxtokenDelegatedPowerChangedIterator struct {
	Event *DydxtokenDelegatedPowerChanged // Event containing the contract specifics and raw log

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
func (it *DydxtokenDelegatedPowerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DydxtokenDelegatedPowerChanged)
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
		it.Event = new(DydxtokenDelegatedPowerChanged)
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
func (it *DydxtokenDelegatedPowerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DydxtokenDelegatedPowerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DydxtokenDelegatedPowerChanged represents a DelegatedPowerChanged event raised by the Dydxtoken contract.
type DydxtokenDelegatedPowerChanged struct {
	User           common.Address
	Amount         *big.Int
	DelegationType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegatedPowerChanged is a free log retrieval operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_Dydxtoken *DydxtokenFilterer) FilterDelegatedPowerChanged(opts *bind.FilterOpts, user []common.Address) (*DydxtokenDelegatedPowerChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Dydxtoken.contract.FilterLogs(opts, "DelegatedPowerChanged", userRule)
	if err != nil {
		return nil, err
	}
	return &DydxtokenDelegatedPowerChangedIterator{contract: _Dydxtoken.contract, event: "DelegatedPowerChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatedPowerChanged is a free log subscription operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_Dydxtoken *DydxtokenFilterer) WatchDelegatedPowerChanged(opts *bind.WatchOpts, sink chan<- *DydxtokenDelegatedPowerChanged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Dydxtoken.contract.WatchLogs(opts, "DelegatedPowerChanged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DydxtokenDelegatedPowerChanged)
				if err := _Dydxtoken.contract.UnpackLog(event, "DelegatedPowerChanged", log); err != nil {
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
func (_Dydxtoken *DydxtokenFilterer) ParseDelegatedPowerChanged(log types.Log) (*DydxtokenDelegatedPowerChanged, error) {
	event := new(DydxtokenDelegatedPowerChanged)
	if err := _Dydxtoken.contract.UnpackLog(event, "DelegatedPowerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DydxtokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dydxtoken contract.
type DydxtokenOwnershipTransferredIterator struct {
	Event *DydxtokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DydxtokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DydxtokenOwnershipTransferred)
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
		it.Event = new(DydxtokenOwnershipTransferred)
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
func (it *DydxtokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DydxtokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DydxtokenOwnershipTransferred represents a OwnershipTransferred event raised by the Dydxtoken contract.
type DydxtokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dydxtoken *DydxtokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DydxtokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dydxtoken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DydxtokenOwnershipTransferredIterator{contract: _Dydxtoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dydxtoken *DydxtokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DydxtokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dydxtoken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DydxtokenOwnershipTransferred)
				if err := _Dydxtoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dydxtoken *DydxtokenFilterer) ParseOwnershipTransferred(log types.Log) (*DydxtokenOwnershipTransferred, error) {
	event := new(DydxtokenOwnershipTransferred)
	if err := _Dydxtoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DydxtokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Dydxtoken contract.
type DydxtokenTransferIterator struct {
	Event *DydxtokenTransfer // Event containing the contract specifics and raw log

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
func (it *DydxtokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DydxtokenTransfer)
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
		it.Event = new(DydxtokenTransfer)
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
func (it *DydxtokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DydxtokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DydxtokenTransfer represents a Transfer event raised by the Dydxtoken contract.
type DydxtokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Dydxtoken *DydxtokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DydxtokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dydxtoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DydxtokenTransferIterator{contract: _Dydxtoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Dydxtoken *DydxtokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DydxtokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dydxtoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DydxtokenTransfer)
				if err := _Dydxtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Dydxtoken *DydxtokenFilterer) ParseTransfer(log types.Log) (*DydxtokenTransfer, error) {
	event := new(DydxtokenTransfer)
	if err := _Dydxtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DydxtokenTransferAllowlistUpdatedIterator is returned from FilterTransferAllowlistUpdated and is used to iterate over the raw logs and unpacked data for TransferAllowlistUpdated events raised by the Dydxtoken contract.
type DydxtokenTransferAllowlistUpdatedIterator struct {
	Event *DydxtokenTransferAllowlistUpdated // Event containing the contract specifics and raw log

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
func (it *DydxtokenTransferAllowlistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DydxtokenTransferAllowlistUpdated)
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
		it.Event = new(DydxtokenTransferAllowlistUpdated)
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
func (it *DydxtokenTransferAllowlistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DydxtokenTransferAllowlistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DydxtokenTransferAllowlistUpdated represents a TransferAllowlistUpdated event raised by the Dydxtoken contract.
type DydxtokenTransferAllowlistUpdated struct {
	Account   common.Address
	IsAllowed bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferAllowlistUpdated is a free log retrieval operation binding the contract event 0x80e7b8bab7ab72e47957c0c472ede1b89bb3e7d2ba30cd37c2d6b000b49a960a.
//
// Solidity: event TransferAllowlistUpdated(address account, bool isAllowed)
func (_Dydxtoken *DydxtokenFilterer) FilterTransferAllowlistUpdated(opts *bind.FilterOpts) (*DydxtokenTransferAllowlistUpdatedIterator, error) {

	logs, sub, err := _Dydxtoken.contract.FilterLogs(opts, "TransferAllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return &DydxtokenTransferAllowlistUpdatedIterator{contract: _Dydxtoken.contract, event: "TransferAllowlistUpdated", logs: logs, sub: sub}, nil
}

// WatchTransferAllowlistUpdated is a free log subscription operation binding the contract event 0x80e7b8bab7ab72e47957c0c472ede1b89bb3e7d2ba30cd37c2d6b000b49a960a.
//
// Solidity: event TransferAllowlistUpdated(address account, bool isAllowed)
func (_Dydxtoken *DydxtokenFilterer) WatchTransferAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *DydxtokenTransferAllowlistUpdated) (event.Subscription, error) {

	logs, sub, err := _Dydxtoken.contract.WatchLogs(opts, "TransferAllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DydxtokenTransferAllowlistUpdated)
				if err := _Dydxtoken.contract.UnpackLog(event, "TransferAllowlistUpdated", log); err != nil {
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

// ParseTransferAllowlistUpdated is a log parse operation binding the contract event 0x80e7b8bab7ab72e47957c0c472ede1b89bb3e7d2ba30cd37c2d6b000b49a960a.
//
// Solidity: event TransferAllowlistUpdated(address account, bool isAllowed)
func (_Dydxtoken *DydxtokenFilterer) ParseTransferAllowlistUpdated(log types.Log) (*DydxtokenTransferAllowlistUpdated, error) {
	event := new(DydxtokenTransferAllowlistUpdated)
	if err := _Dydxtoken.contract.UnpackLog(event, "TransferAllowlistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DydxtokenTransfersRestrictedBeforeUpdatedIterator is returned from FilterTransfersRestrictedBeforeUpdated and is used to iterate over the raw logs and unpacked data for TransfersRestrictedBeforeUpdated events raised by the Dydxtoken contract.
type DydxtokenTransfersRestrictedBeforeUpdatedIterator struct {
	Event *DydxtokenTransfersRestrictedBeforeUpdated // Event containing the contract specifics and raw log

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
func (it *DydxtokenTransfersRestrictedBeforeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DydxtokenTransfersRestrictedBeforeUpdated)
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
		it.Event = new(DydxtokenTransfersRestrictedBeforeUpdated)
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
func (it *DydxtokenTransfersRestrictedBeforeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DydxtokenTransfersRestrictedBeforeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DydxtokenTransfersRestrictedBeforeUpdated represents a TransfersRestrictedBeforeUpdated event raised by the Dydxtoken contract.
type DydxtokenTransfersRestrictedBeforeUpdated struct {
	TransfersRestrictedBefore *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterTransfersRestrictedBeforeUpdated is a free log retrieval operation binding the contract event 0xd7b9c9321b627ff004969698b4616502d6b7305a588e685489e91c46fca509a9.
//
// Solidity: event TransfersRestrictedBeforeUpdated(uint256 transfersRestrictedBefore)
func (_Dydxtoken *DydxtokenFilterer) FilterTransfersRestrictedBeforeUpdated(opts *bind.FilterOpts) (*DydxtokenTransfersRestrictedBeforeUpdatedIterator, error) {

	logs, sub, err := _Dydxtoken.contract.FilterLogs(opts, "TransfersRestrictedBeforeUpdated")
	if err != nil {
		return nil, err
	}
	return &DydxtokenTransfersRestrictedBeforeUpdatedIterator{contract: _Dydxtoken.contract, event: "TransfersRestrictedBeforeUpdated", logs: logs, sub: sub}, nil
}

// WatchTransfersRestrictedBeforeUpdated is a free log subscription operation binding the contract event 0xd7b9c9321b627ff004969698b4616502d6b7305a588e685489e91c46fca509a9.
//
// Solidity: event TransfersRestrictedBeforeUpdated(uint256 transfersRestrictedBefore)
func (_Dydxtoken *DydxtokenFilterer) WatchTransfersRestrictedBeforeUpdated(opts *bind.WatchOpts, sink chan<- *DydxtokenTransfersRestrictedBeforeUpdated) (event.Subscription, error) {

	logs, sub, err := _Dydxtoken.contract.WatchLogs(opts, "TransfersRestrictedBeforeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DydxtokenTransfersRestrictedBeforeUpdated)
				if err := _Dydxtoken.contract.UnpackLog(event, "TransfersRestrictedBeforeUpdated", log); err != nil {
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

// ParseTransfersRestrictedBeforeUpdated is a log parse operation binding the contract event 0xd7b9c9321b627ff004969698b4616502d6b7305a588e685489e91c46fca509a9.
//
// Solidity: event TransfersRestrictedBeforeUpdated(uint256 transfersRestrictedBefore)
func (_Dydxtoken *DydxtokenFilterer) ParseTransfersRestrictedBeforeUpdated(log types.Log) (*DydxtokenTransfersRestrictedBeforeUpdated, error) {
	event := new(DydxtokenTransfersRestrictedBeforeUpdated)
	if err := _Dydxtoken.contract.UnpackLog(event, "TransfersRestrictedBeforeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
