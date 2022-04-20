// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// TokenVestingVestingSchedule is an auto generated low-level Go binding around an user-defined struct.
type TokenVestingVestingSchedule struct {
	Initialized        bool
	Beneficiary        common.Address
	Cliff              *big.Int
	Start              *big.Int
	Duration           *big.Int
	SlicePeriodSeconds *big.Int
	Revocable          bool
	AmountTotal        *big.Int
	Released           *big.Int
	Revoked            bool
	Locked             bool
}

// AbiMetaData contains all meta data concerning the Abi contract.
var AbiMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"token_\",\"internalType\":\"address\"}]},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Released\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Revoked\",\"inputs\":[],\"anonymous\":false},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"computeNextVestingScheduleIdForHolder\",\"inputs\":[{\"type\":\"address\",\"name\":\"holder\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"computeReleasableAmount\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"vestingScheduleId\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"computeVestingScheduleIdForAddressAndIndex\",\"inputs\":[{\"type\":\"address\",\"name\":\"holder\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"createVestingSchedule\",\"inputs\":[{\"type\":\"address\",\"name\":\"_beneficiary\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_start\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_cliff\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_duration\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_slicePeriodSeconds\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"_revocable\",\"internalType\":\"bool\"},{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structTokenVesting.VestingSchedule\",\"components\":[{\"type\":\"bool\",\"name\":\"initialized\",\"internalType\":\"bool\"},{\"type\":\"address\",\"name\":\"beneficiary\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"cliff\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"start\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"duration\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"slicePeriodSeconds\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"revocable\",\"internalType\":\"bool\"},{\"type\":\"uint256\",\"name\":\"amountTotal\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"released\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"revoked\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]}],\"name\":\"getLastVestingScheduleForHolder\",\"inputs\":[{\"type\":\"address\",\"name\":\"holder\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getToken\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"getVestingIdAtIndex\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structTokenVesting.VestingSchedule\",\"components\":[{\"type\":\"bool\",\"name\":\"initialized\",\"internalType\":\"bool\"},{\"type\":\"address\",\"name\":\"beneficiary\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"cliff\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"start\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"duration\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"slicePeriodSeconds\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"revocable\",\"internalType\":\"bool\"},{\"type\":\"uint256\",\"name\":\"amountTotal\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"released\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"revoked\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]}],\"name\":\"getVestingSchedule\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"vestingScheduleId\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structTokenVesting.VestingSchedule\",\"components\":[{\"type\":\"bool\",\"name\":\"initialized\",\"internalType\":\"bool\"},{\"type\":\"address\",\"name\":\"beneficiary\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"cliff\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"start\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"duration\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"slicePeriodSeconds\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"revocable\",\"internalType\":\"bool\"},{\"type\":\"uint256\",\"name\":\"amountTotal\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"released\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"revoked\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]}],\"name\":\"getVestingScheduleByAddressAndIndex\",\"inputs\":[{\"type\":\"address\",\"name\":\"holder\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getVestingSchedulesCount\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getVestingSchedulesCountByBeneficiary\",\"inputs\":[{\"type\":\"address\",\"name\":\"_beneficiary\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getVestingSchedulesTotalAmount\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getWithdrawableAmount\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"total\",\"internalType\":\"uint256\"}],\"name\":\"investorWithdraw\",\"inputs\":[{\"type\":\"bool\",\"name\":\"keepWrapped\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"release\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"vestingScheduleId\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"revoke\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"vestingScheduleId\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setLock\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"vestingScheduleId\",\"internalType\":\"bytes32\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"transferOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"withdraw\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"receive\",\"stateMutability\":\"payable\"}]",
}

// AbiABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiMetaData.ABI instead.
var AbiABI = AbiMetaData.ABI

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// ComputeNextVestingScheduleIdForHolder is a free data retrieval call binding the contract method 0xf7c469f0.
//
// Solidity: function computeNextVestingScheduleIdForHolder(address holder) view returns(bytes32)
func (_Abi *AbiCaller) ComputeNextVestingScheduleIdForHolder(opts *bind.CallOpts, holder common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "computeNextVestingScheduleIdForHolder", holder)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeNextVestingScheduleIdForHolder is a free data retrieval call binding the contract method 0xf7c469f0.
//
// Solidity: function computeNextVestingScheduleIdForHolder(address holder) view returns(bytes32)
func (_Abi *AbiSession) ComputeNextVestingScheduleIdForHolder(holder common.Address) ([32]byte, error) {
	return _Abi.Contract.ComputeNextVestingScheduleIdForHolder(&_Abi.CallOpts, holder)
}

// ComputeNextVestingScheduleIdForHolder is a free data retrieval call binding the contract method 0xf7c469f0.
//
// Solidity: function computeNextVestingScheduleIdForHolder(address holder) view returns(bytes32)
func (_Abi *AbiCallerSession) ComputeNextVestingScheduleIdForHolder(holder common.Address) ([32]byte, error) {
	return _Abi.Contract.ComputeNextVestingScheduleIdForHolder(&_Abi.CallOpts, holder)
}

// ComputeReleasableAmount is a free data retrieval call binding the contract method 0xea1bb3d5.
//
// Solidity: function computeReleasableAmount(bytes32 vestingScheduleId) view returns(uint256)
func (_Abi *AbiCaller) ComputeReleasableAmount(opts *bind.CallOpts, vestingScheduleId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "computeReleasableAmount", vestingScheduleId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeReleasableAmount is a free data retrieval call binding the contract method 0xea1bb3d5.
//
// Solidity: function computeReleasableAmount(bytes32 vestingScheduleId) view returns(uint256)
func (_Abi *AbiSession) ComputeReleasableAmount(vestingScheduleId [32]byte) (*big.Int, error) {
	return _Abi.Contract.ComputeReleasableAmount(&_Abi.CallOpts, vestingScheduleId)
}

// ComputeReleasableAmount is a free data retrieval call binding the contract method 0xea1bb3d5.
//
// Solidity: function computeReleasableAmount(bytes32 vestingScheduleId) view returns(uint256)
func (_Abi *AbiCallerSession) ComputeReleasableAmount(vestingScheduleId [32]byte) (*big.Int, error) {
	return _Abi.Contract.ComputeReleasableAmount(&_Abi.CallOpts, vestingScheduleId)
}

// ComputeVestingScheduleIdForAddressAndIndex is a free data retrieval call binding the contract method 0x8af104da.
//
// Solidity: function computeVestingScheduleIdForAddressAndIndex(address holder, uint256 index) pure returns(bytes32)
func (_Abi *AbiCaller) ComputeVestingScheduleIdForAddressAndIndex(opts *bind.CallOpts, holder common.Address, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "computeVestingScheduleIdForAddressAndIndex", holder, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeVestingScheduleIdForAddressAndIndex is a free data retrieval call binding the contract method 0x8af104da.
//
// Solidity: function computeVestingScheduleIdForAddressAndIndex(address holder, uint256 index) pure returns(bytes32)
func (_Abi *AbiSession) ComputeVestingScheduleIdForAddressAndIndex(holder common.Address, index *big.Int) ([32]byte, error) {
	return _Abi.Contract.ComputeVestingScheduleIdForAddressAndIndex(&_Abi.CallOpts, holder, index)
}

// ComputeVestingScheduleIdForAddressAndIndex is a free data retrieval call binding the contract method 0x8af104da.
//
// Solidity: function computeVestingScheduleIdForAddressAndIndex(address holder, uint256 index) pure returns(bytes32)
func (_Abi *AbiCallerSession) ComputeVestingScheduleIdForAddressAndIndex(holder common.Address, index *big.Int) ([32]byte, error) {
	return _Abi.Contract.ComputeVestingScheduleIdForAddressAndIndex(&_Abi.CallOpts, holder, index)
}

// GetLastVestingScheduleForHolder is a free data retrieval call binding the contract method 0x7e913dc6.
//
// Solidity: function getLastVestingScheduleForHolder(address holder) view returns((bool,address,uint256,uint256,uint256,uint256,bool,uint256,uint256,bool,bool))
func (_Abi *AbiCaller) GetLastVestingScheduleForHolder(opts *bind.CallOpts, holder common.Address) (TokenVestingVestingSchedule, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getLastVestingScheduleForHolder", holder)

	if err != nil {
		return *new(TokenVestingVestingSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenVestingVestingSchedule)).(*TokenVestingVestingSchedule)

	return out0, err

}

// GetLastVestingScheduleForHolder is a free data retrieval call binding the contract method 0x7e913dc6.
//
// Solidity: function getLastVestingScheduleForHolder(address holder) view returns((bool,address,uint256,uint256,uint256,uint256,bool,uint256,uint256,bool,bool))
func (_Abi *AbiSession) GetLastVestingScheduleForHolder(holder common.Address) (TokenVestingVestingSchedule, error) {
	return _Abi.Contract.GetLastVestingScheduleForHolder(&_Abi.CallOpts, holder)
}

// GetLastVestingScheduleForHolder is a free data retrieval call binding the contract method 0x7e913dc6.
//
// Solidity: function getLastVestingScheduleForHolder(address holder) view returns((bool,address,uint256,uint256,uint256,uint256,bool,uint256,uint256,bool,bool))
func (_Abi *AbiCallerSession) GetLastVestingScheduleForHolder(holder common.Address) (TokenVestingVestingSchedule, error) {
	return _Abi.Contract.GetLastVestingScheduleForHolder(&_Abi.CallOpts, holder)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Abi *AbiCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Abi *AbiSession) GetToken() (common.Address, error) {
	return _Abi.Contract.GetToken(&_Abi.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Abi *AbiCallerSession) GetToken() (common.Address, error) {
	return _Abi.Contract.GetToken(&_Abi.CallOpts)
}

// GetVestingIdAtIndex is a free data retrieval call binding the contract method 0xf9079b37.
//
// Solidity: function getVestingIdAtIndex(uint256 index) view returns(bytes32)
func (_Abi *AbiCaller) GetVestingIdAtIndex(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getVestingIdAtIndex", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVestingIdAtIndex is a free data retrieval call binding the contract method 0xf9079b37.
//
// Solidity: function getVestingIdAtIndex(uint256 index) view returns(bytes32)
func (_Abi *AbiSession) GetVestingIdAtIndex(index *big.Int) ([32]byte, error) {
	return _Abi.Contract.GetVestingIdAtIndex(&_Abi.CallOpts, index)
}

// GetVestingIdAtIndex is a free data retrieval call binding the contract method 0xf9079b37.
//
// Solidity: function getVestingIdAtIndex(uint256 index) view returns(bytes32)
func (_Abi *AbiCallerSession) GetVestingIdAtIndex(index *big.Int) ([32]byte, error) {
	return _Abi.Contract.GetVestingIdAtIndex(&_Abi.CallOpts, index)
}

// GetVestingSchedule is a free data retrieval call binding the contract method 0x9ef346b4.
//
// Solidity: function getVestingSchedule(bytes32 vestingScheduleId) view returns((bool,address,uint256,uint256,uint256,uint256,bool,uint256,uint256,bool,bool))
func (_Abi *AbiCaller) GetVestingSchedule(opts *bind.CallOpts, vestingScheduleId [32]byte) (TokenVestingVestingSchedule, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getVestingSchedule", vestingScheduleId)

	if err != nil {
		return *new(TokenVestingVestingSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenVestingVestingSchedule)).(*TokenVestingVestingSchedule)

	return out0, err

}

// GetVestingSchedule is a free data retrieval call binding the contract method 0x9ef346b4.
//
// Solidity: function getVestingSchedule(bytes32 vestingScheduleId) view returns((bool,address,uint256,uint256,uint256,uint256,bool,uint256,uint256,bool,bool))
func (_Abi *AbiSession) GetVestingSchedule(vestingScheduleId [32]byte) (TokenVestingVestingSchedule, error) {
	return _Abi.Contract.GetVestingSchedule(&_Abi.CallOpts, vestingScheduleId)
}

// GetVestingSchedule is a free data retrieval call binding the contract method 0x9ef346b4.
//
// Solidity: function getVestingSchedule(bytes32 vestingScheduleId) view returns((bool,address,uint256,uint256,uint256,uint256,bool,uint256,uint256,bool,bool))
func (_Abi *AbiCallerSession) GetVestingSchedule(vestingScheduleId [32]byte) (TokenVestingVestingSchedule, error) {
	return _Abi.Contract.GetVestingSchedule(&_Abi.CallOpts, vestingScheduleId)
}

// GetVestingScheduleByAddressAndIndex is a free data retrieval call binding the contract method 0xf51321d7.
//
// Solidity: function getVestingScheduleByAddressAndIndex(address holder, uint256 index) view returns((bool,address,uint256,uint256,uint256,uint256,bool,uint256,uint256,bool,bool))
func (_Abi *AbiCaller) GetVestingScheduleByAddressAndIndex(opts *bind.CallOpts, holder common.Address, index *big.Int) (TokenVestingVestingSchedule, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getVestingScheduleByAddressAndIndex", holder, index)

	if err != nil {
		return *new(TokenVestingVestingSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenVestingVestingSchedule)).(*TokenVestingVestingSchedule)

	return out0, err

}

// GetVestingScheduleByAddressAndIndex is a free data retrieval call binding the contract method 0xf51321d7.
//
// Solidity: function getVestingScheduleByAddressAndIndex(address holder, uint256 index) view returns((bool,address,uint256,uint256,uint256,uint256,bool,uint256,uint256,bool,bool))
func (_Abi *AbiSession) GetVestingScheduleByAddressAndIndex(holder common.Address, index *big.Int) (TokenVestingVestingSchedule, error) {
	return _Abi.Contract.GetVestingScheduleByAddressAndIndex(&_Abi.CallOpts, holder, index)
}

// GetVestingScheduleByAddressAndIndex is a free data retrieval call binding the contract method 0xf51321d7.
//
// Solidity: function getVestingScheduleByAddressAndIndex(address holder, uint256 index) view returns((bool,address,uint256,uint256,uint256,uint256,bool,uint256,uint256,bool,bool))
func (_Abi *AbiCallerSession) GetVestingScheduleByAddressAndIndex(holder common.Address, index *big.Int) (TokenVestingVestingSchedule, error) {
	return _Abi.Contract.GetVestingScheduleByAddressAndIndex(&_Abi.CallOpts, holder, index)
}

// GetVestingSchedulesCount is a free data retrieval call binding the contract method 0x13083617.
//
// Solidity: function getVestingSchedulesCount() view returns(uint256)
func (_Abi *AbiCaller) GetVestingSchedulesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getVestingSchedulesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestingSchedulesCount is a free data retrieval call binding the contract method 0x13083617.
//
// Solidity: function getVestingSchedulesCount() view returns(uint256)
func (_Abi *AbiSession) GetVestingSchedulesCount() (*big.Int, error) {
	return _Abi.Contract.GetVestingSchedulesCount(&_Abi.CallOpts)
}

// GetVestingSchedulesCount is a free data retrieval call binding the contract method 0x13083617.
//
// Solidity: function getVestingSchedulesCount() view returns(uint256)
func (_Abi *AbiCallerSession) GetVestingSchedulesCount() (*big.Int, error) {
	return _Abi.Contract.GetVestingSchedulesCount(&_Abi.CallOpts)
}

// GetVestingSchedulesCountByBeneficiary is a free data retrieval call binding the contract method 0x5a7bb69a.
//
// Solidity: function getVestingSchedulesCountByBeneficiary(address _beneficiary) view returns(uint256)
func (_Abi *AbiCaller) GetVestingSchedulesCountByBeneficiary(opts *bind.CallOpts, _beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getVestingSchedulesCountByBeneficiary", _beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestingSchedulesCountByBeneficiary is a free data retrieval call binding the contract method 0x5a7bb69a.
//
// Solidity: function getVestingSchedulesCountByBeneficiary(address _beneficiary) view returns(uint256)
func (_Abi *AbiSession) GetVestingSchedulesCountByBeneficiary(_beneficiary common.Address) (*big.Int, error) {
	return _Abi.Contract.GetVestingSchedulesCountByBeneficiary(&_Abi.CallOpts, _beneficiary)
}

// GetVestingSchedulesCountByBeneficiary is a free data retrieval call binding the contract method 0x5a7bb69a.
//
// Solidity: function getVestingSchedulesCountByBeneficiary(address _beneficiary) view returns(uint256)
func (_Abi *AbiCallerSession) GetVestingSchedulesCountByBeneficiary(_beneficiary common.Address) (*big.Int, error) {
	return _Abi.Contract.GetVestingSchedulesCountByBeneficiary(&_Abi.CallOpts, _beneficiary)
}

// GetVestingSchedulesTotalAmount is a free data retrieval call binding the contract method 0x48deb471.
//
// Solidity: function getVestingSchedulesTotalAmount() view returns(uint256)
func (_Abi *AbiCaller) GetVestingSchedulesTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getVestingSchedulesTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestingSchedulesTotalAmount is a free data retrieval call binding the contract method 0x48deb471.
//
// Solidity: function getVestingSchedulesTotalAmount() view returns(uint256)
func (_Abi *AbiSession) GetVestingSchedulesTotalAmount() (*big.Int, error) {
	return _Abi.Contract.GetVestingSchedulesTotalAmount(&_Abi.CallOpts)
}

// GetVestingSchedulesTotalAmount is a free data retrieval call binding the contract method 0x48deb471.
//
// Solidity: function getVestingSchedulesTotalAmount() view returns(uint256)
func (_Abi *AbiCallerSession) GetVestingSchedulesTotalAmount() (*big.Int, error) {
	return _Abi.Contract.GetVestingSchedulesTotalAmount(&_Abi.CallOpts)
}

// GetWithdrawableAmount is a free data retrieval call binding the contract method 0x90be10cc.
//
// Solidity: function getWithdrawableAmount() view returns(uint256)
func (_Abi *AbiCaller) GetWithdrawableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getWithdrawableAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableAmount is a free data retrieval call binding the contract method 0x90be10cc.
//
// Solidity: function getWithdrawableAmount() view returns(uint256)
func (_Abi *AbiSession) GetWithdrawableAmount() (*big.Int, error) {
	return _Abi.Contract.GetWithdrawableAmount(&_Abi.CallOpts)
}

// GetWithdrawableAmount is a free data retrieval call binding the contract method 0x90be10cc.
//
// Solidity: function getWithdrawableAmount() view returns(uint256)
func (_Abi *AbiCallerSession) GetWithdrawableAmount() (*big.Int, error) {
	return _Abi.Contract.GetWithdrawableAmount(&_Abi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiSession) Owner() (common.Address, error) {
	return _Abi.Contract.Owner(&_Abi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiCallerSession) Owner() (common.Address, error) {
	return _Abi.Contract.Owner(&_Abi.CallOpts)
}

// CreateVestingSchedule is a paid mutator transaction binding the contract method 0x17e289e9.
//
// Solidity: function createVestingSchedule(address _beneficiary, uint256 _start, uint256 _cliff, uint256 _duration, uint256 _slicePeriodSeconds, bool _revocable, uint256 _amount) returns()
func (_Abi *AbiTransactor) CreateVestingSchedule(opts *bind.TransactOpts, _beneficiary common.Address, _start *big.Int, _cliff *big.Int, _duration *big.Int, _slicePeriodSeconds *big.Int, _revocable bool, _amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "createVestingSchedule", _beneficiary, _start, _cliff, _duration, _slicePeriodSeconds, _revocable, _amount)
}

// CreateVestingSchedule is a paid mutator transaction binding the contract method 0x17e289e9.
//
// Solidity: function createVestingSchedule(address _beneficiary, uint256 _start, uint256 _cliff, uint256 _duration, uint256 _slicePeriodSeconds, bool _revocable, uint256 _amount) returns()
func (_Abi *AbiSession) CreateVestingSchedule(_beneficiary common.Address, _start *big.Int, _cliff *big.Int, _duration *big.Int, _slicePeriodSeconds *big.Int, _revocable bool, _amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.CreateVestingSchedule(&_Abi.TransactOpts, _beneficiary, _start, _cliff, _duration, _slicePeriodSeconds, _revocable, _amount)
}

// CreateVestingSchedule is a paid mutator transaction binding the contract method 0x17e289e9.
//
// Solidity: function createVestingSchedule(address _beneficiary, uint256 _start, uint256 _cliff, uint256 _duration, uint256 _slicePeriodSeconds, bool _revocable, uint256 _amount) returns()
func (_Abi *AbiTransactorSession) CreateVestingSchedule(_beneficiary common.Address, _start *big.Int, _cliff *big.Int, _duration *big.Int, _slicePeriodSeconds *big.Int, _revocable bool, _amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.CreateVestingSchedule(&_Abi.TransactOpts, _beneficiary, _start, _cliff, _duration, _slicePeriodSeconds, _revocable, _amount)
}

// InvestorWithdraw is a paid mutator transaction binding the contract method 0x623a8716.
//
// Solidity: function investorWithdraw(bool keepWrapped) returns(uint256 total)
func (_Abi *AbiTransactor) InvestorWithdraw(opts *bind.TransactOpts, keepWrapped bool) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "investorWithdraw", keepWrapped)
}

// InvestorWithdraw is a paid mutator transaction binding the contract method 0x623a8716.
//
// Solidity: function investorWithdraw(bool keepWrapped) returns(uint256 total)
func (_Abi *AbiSession) InvestorWithdraw(keepWrapped bool) (*types.Transaction, error) {
	return _Abi.Contract.InvestorWithdraw(&_Abi.TransactOpts, keepWrapped)
}

// InvestorWithdraw is a paid mutator transaction binding the contract method 0x623a8716.
//
// Solidity: function investorWithdraw(bool keepWrapped) returns(uint256 total)
func (_Abi *AbiTransactorSession) InvestorWithdraw(keepWrapped bool) (*types.Transaction, error) {
	return _Abi.Contract.InvestorWithdraw(&_Abi.TransactOpts, keepWrapped)
}

// Release is a paid mutator transaction binding the contract method 0x66afd8ef.
//
// Solidity: function release(bytes32 vestingScheduleId, uint256 amount) returns()
func (_Abi *AbiTransactor) Release(opts *bind.TransactOpts, vestingScheduleId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "release", vestingScheduleId, amount)
}

// Release is a paid mutator transaction binding the contract method 0x66afd8ef.
//
// Solidity: function release(bytes32 vestingScheduleId, uint256 amount) returns()
func (_Abi *AbiSession) Release(vestingScheduleId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Release(&_Abi.TransactOpts, vestingScheduleId, amount)
}

// Release is a paid mutator transaction binding the contract method 0x66afd8ef.
//
// Solidity: function release(bytes32 vestingScheduleId, uint256 amount) returns()
func (_Abi *AbiTransactorSession) Release(vestingScheduleId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Release(&_Abi.TransactOpts, vestingScheduleId, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abi.Contract.RenounceOwnership(&_Abi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abi.Contract.RenounceOwnership(&_Abi.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 vestingScheduleId) returns()
func (_Abi *AbiTransactor) Revoke(opts *bind.TransactOpts, vestingScheduleId [32]byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "revoke", vestingScheduleId)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 vestingScheduleId) returns()
func (_Abi *AbiSession) Revoke(vestingScheduleId [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.Revoke(&_Abi.TransactOpts, vestingScheduleId)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 vestingScheduleId) returns()
func (_Abi *AbiTransactorSession) Revoke(vestingScheduleId [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.Revoke(&_Abi.TransactOpts, vestingScheduleId)
}

// SetLock is a paid mutator transaction binding the contract method 0xe66226a6.
//
// Solidity: function setLock(bytes32 vestingScheduleId, bool locked) returns()
func (_Abi *AbiTransactor) SetLock(opts *bind.TransactOpts, vestingScheduleId [32]byte, locked bool) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setLock", vestingScheduleId, locked)
}

// SetLock is a paid mutator transaction binding the contract method 0xe66226a6.
//
// Solidity: function setLock(bytes32 vestingScheduleId, bool locked) returns()
func (_Abi *AbiSession) SetLock(vestingScheduleId [32]byte, locked bool) (*types.Transaction, error) {
	return _Abi.Contract.SetLock(&_Abi.TransactOpts, vestingScheduleId, locked)
}

// SetLock is a paid mutator transaction binding the contract method 0xe66226a6.
//
// Solidity: function setLock(bytes32 vestingScheduleId, bool locked) returns()
func (_Abi *AbiTransactorSession) SetLock(vestingScheduleId [32]byte, locked bool) (*types.Transaction, error) {
	return _Abi.Contract.SetLock(&_Abi.TransactOpts, vestingScheduleId, locked)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferOwnership(&_Abi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferOwnership(&_Abi.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Abi *AbiTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Abi *AbiSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Withdraw(&_Abi.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Abi *AbiTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Withdraw(&_Abi.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abi *AbiTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Abi.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abi *AbiSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Abi.Contract.Fallback(&_Abi.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abi *AbiTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Abi.Contract.Fallback(&_Abi.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abi *AbiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abi *AbiSession) Receive() (*types.Transaction, error) {
	return _Abi.Contract.Receive(&_Abi.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abi *AbiTransactorSession) Receive() (*types.Transaction, error) {
	return _Abi.Contract.Receive(&_Abi.TransactOpts)
}

// AbiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Abi contract.
type AbiOwnershipTransferredIterator struct {
	Event *AbiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AbiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiOwnershipTransferred)
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
		it.Event = new(AbiOwnershipTransferred)
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
func (it *AbiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiOwnershipTransferred represents a OwnershipTransferred event raised by the Abi contract.
type AbiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AbiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AbiOwnershipTransferredIterator{contract: _Abi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AbiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiOwnershipTransferred)
				if err := _Abi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Abi *AbiFilterer) ParseOwnershipTransferred(log types.Log) (*AbiOwnershipTransferred, error) {
	event := new(AbiOwnershipTransferred)
	if err := _Abi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiReleasedIterator is returned from FilterReleased and is used to iterate over the raw logs and unpacked data for Released events raised by the Abi contract.
type AbiReleasedIterator struct {
	Event *AbiReleased // Event containing the contract specifics and raw log

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
func (it *AbiReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiReleased)
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
		it.Event = new(AbiReleased)
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
func (it *AbiReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiReleased represents a Released event raised by the Abi contract.
type AbiReleased struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReleased is a free log retrieval operation binding the contract event 0xfb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c565.
//
// Solidity: event Released(uint256 amount)
func (_Abi *AbiFilterer) FilterReleased(opts *bind.FilterOpts) (*AbiReleasedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Released")
	if err != nil {
		return nil, err
	}
	return &AbiReleasedIterator{contract: _Abi.contract, event: "Released", logs: logs, sub: sub}, nil
}

// WatchReleased is a free log subscription operation binding the contract event 0xfb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c565.
//
// Solidity: event Released(uint256 amount)
func (_Abi *AbiFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *AbiReleased) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Released")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiReleased)
				if err := _Abi.contract.UnpackLog(event, "Released", log); err != nil {
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

// ParseReleased is a log parse operation binding the contract event 0xfb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c565.
//
// Solidity: event Released(uint256 amount)
func (_Abi *AbiFilterer) ParseReleased(log types.Log) (*AbiReleased, error) {
	event := new(AbiReleased)
	if err := _Abi.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the Abi contract.
type AbiRevokedIterator struct {
	Event *AbiRevoked // Event containing the contract specifics and raw log

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
func (it *AbiRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRevoked)
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
		it.Event = new(AbiRevoked)
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
func (it *AbiRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRevoked represents a Revoked event raised by the Abi contract.
type AbiRevoked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: event Revoked()
func (_Abi *AbiFilterer) FilterRevoked(opts *bind.FilterOpts) (*AbiRevokedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Revoked")
	if err != nil {
		return nil, err
	}
	return &AbiRevokedIterator{contract: _Abi.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: event Revoked()
func (_Abi *AbiFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *AbiRevoked) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Revoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRevoked)
				if err := _Abi.contract.UnpackLog(event, "Revoked", log); err != nil {
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

// ParseRevoked is a log parse operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: event Revoked()
func (_Abi *AbiFilterer) ParseRevoked(log types.Log) (*AbiRevoked, error) {
	event := new(AbiRevoked)
	if err := _Abi.contract.UnpackLog(event, "Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
