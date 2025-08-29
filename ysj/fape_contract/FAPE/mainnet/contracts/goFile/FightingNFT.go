// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FightingNFT

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

// FightingNFTMetaData contains all meta data concerning the FightingNFT contract.
var FightingNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_isStartMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"addWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"at\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstMint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"isInWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"mintNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"removeWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseUri_\",\"type\":\"string\"}],\"name\":\"setBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isStartMint_\",\"type\":\"bool\"}],\"name\":\"setIsStartMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FightingNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use FightingNFTMetaData.ABI instead.
var FightingNFTABI = FightingNFTMetaData.ABI

// FightingNFT is an auto generated Go binding around an Ethereum contract.
type FightingNFT struct {
	FightingNFTCaller     // Read-only binding to the contract
	FightingNFTTransactor // Write-only binding to the contract
	FightingNFTFilterer   // Log filterer for contract events
}

// FightingNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type FightingNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FightingNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FightingNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FightingNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FightingNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FightingNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FightingNFTSession struct {
	Contract     *FightingNFT      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FightingNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FightingNFTCallerSession struct {
	Contract *FightingNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FightingNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FightingNFTTransactorSession struct {
	Contract     *FightingNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FightingNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type FightingNFTRaw struct {
	Contract *FightingNFT // Generic contract binding to access the raw methods on
}

// FightingNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FightingNFTCallerRaw struct {
	Contract *FightingNFTCaller // Generic read-only contract binding to access the raw methods on
}

// FightingNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FightingNFTTransactorRaw struct {
	Contract *FightingNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFightingNFT creates a new instance of FightingNFT, bound to a specific deployed contract.
func NewFightingNFT(address common.Address, backend bind.ContractBackend) (*FightingNFT, error) {
	contract, err := bindFightingNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FightingNFT{FightingNFTCaller: FightingNFTCaller{contract: contract}, FightingNFTTransactor: FightingNFTTransactor{contract: contract}, FightingNFTFilterer: FightingNFTFilterer{contract: contract}}, nil
}

// NewFightingNFTCaller creates a new read-only instance of FightingNFT, bound to a specific deployed contract.
func NewFightingNFTCaller(address common.Address, caller bind.ContractCaller) (*FightingNFTCaller, error) {
	contract, err := bindFightingNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FightingNFTCaller{contract: contract}, nil
}

// NewFightingNFTTransactor creates a new write-only instance of FightingNFT, bound to a specific deployed contract.
func NewFightingNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*FightingNFTTransactor, error) {
	contract, err := bindFightingNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FightingNFTTransactor{contract: contract}, nil
}

// NewFightingNFTFilterer creates a new log filterer instance of FightingNFT, bound to a specific deployed contract.
func NewFightingNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*FightingNFTFilterer, error) {
	contract, err := bindFightingNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FightingNFTFilterer{contract: contract}, nil
}

// bindFightingNFT binds a generic wrapper to an already deployed contract.
func bindFightingNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FightingNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FightingNFT *FightingNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FightingNFT.Contract.FightingNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FightingNFT *FightingNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FightingNFT.Contract.FightingNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FightingNFT *FightingNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FightingNFT.Contract.FightingNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FightingNFT *FightingNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FightingNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FightingNFT *FightingNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FightingNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FightingNFT *FightingNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FightingNFT.Contract.contract.Transact(opts, method, params...)
}

// IsStartMint is a free data retrieval call binding the contract method 0xbde5719e.
//
// Solidity: function _isStartMint() view returns(bool)
func (_FightingNFT *FightingNFTCaller) IsStartMint(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "_isStartMint")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStartMint is a free data retrieval call binding the contract method 0xbde5719e.
//
// Solidity: function _isStartMint() view returns(bool)
func (_FightingNFT *FightingNFTSession) IsStartMint() (bool, error) {
	return _FightingNFT.Contract.IsStartMint(&_FightingNFT.CallOpts)
}

// IsStartMint is a free data retrieval call binding the contract method 0xbde5719e.
//
// Solidity: function _isStartMint() view returns(bool)
func (_FightingNFT *FightingNFTCallerSession) IsStartMint() (bool, error) {
	return _FightingNFT.Contract.IsStartMint(&_FightingNFT.CallOpts)
}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index_) view returns(address)
func (_FightingNFT *FightingNFTCaller) At(opts *bind.CallOpts, index_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "at", index_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index_) view returns(address)
func (_FightingNFT *FightingNFTSession) At(index_ *big.Int) (common.Address, error) {
	return _FightingNFT.Contract.At(&_FightingNFT.CallOpts, index_)
}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index_) view returns(address)
func (_FightingNFT *FightingNFTCallerSession) At(index_ *big.Int) (common.Address, error) {
	return _FightingNFT.Contract.At(&_FightingNFT.CallOpts, index_)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FightingNFT *FightingNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FightingNFT *FightingNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FightingNFT.Contract.BalanceOf(&_FightingNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FightingNFT *FightingNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FightingNFT.Contract.BalanceOf(&_FightingNFT.CallOpts, owner)
}

// FirstMint is a free data retrieval call binding the contract method 0x60724708.
//
// Solidity: function firstMint() view returns(address)
func (_FightingNFT *FightingNFTCaller) FirstMint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "firstMint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FirstMint is a free data retrieval call binding the contract method 0x60724708.
//
// Solidity: function firstMint() view returns(address)
func (_FightingNFT *FightingNFTSession) FirstMint() (common.Address, error) {
	return _FightingNFT.Contract.FirstMint(&_FightingNFT.CallOpts)
}

// FirstMint is a free data retrieval call binding the contract method 0x60724708.
//
// Solidity: function firstMint() view returns(address)
func (_FightingNFT *FightingNFTCallerSession) FirstMint() (common.Address, error) {
	return _FightingNFT.Contract.FirstMint(&_FightingNFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FightingNFT *FightingNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FightingNFT *FightingNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FightingNFT.Contract.GetApproved(&_FightingNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FightingNFT *FightingNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FightingNFT.Contract.GetApproved(&_FightingNFT.CallOpts, tokenId)
}

// GetUserCounts is a free data retrieval call binding the contract method 0x79a6b5e8.
//
// Solidity: function getUserCounts() view returns(uint256)
func (_FightingNFT *FightingNFTCaller) GetUserCounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "getUserCounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserCounts is a free data retrieval call binding the contract method 0x79a6b5e8.
//
// Solidity: function getUserCounts() view returns(uint256)
func (_FightingNFT *FightingNFTSession) GetUserCounts() (*big.Int, error) {
	return _FightingNFT.Contract.GetUserCounts(&_FightingNFT.CallOpts)
}

// GetUserCounts is a free data retrieval call binding the contract method 0x79a6b5e8.
//
// Solidity: function getUserCounts() view returns(uint256)
func (_FightingNFT *FightingNFTCallerSession) GetUserCounts() (*big.Int, error) {
	return _FightingNFT.Contract.GetUserCounts(&_FightingNFT.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FightingNFT *FightingNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FightingNFT *FightingNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FightingNFT.Contract.IsApprovedForAll(&_FightingNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FightingNFT *FightingNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FightingNFT.Contract.IsApprovedForAll(&_FightingNFT.CallOpts, owner, operator)
}

// IsInWhitelist is a free data retrieval call binding the contract method 0x09fd8212.
//
// Solidity: function isInWhitelist(address account_) view returns(bool)
func (_FightingNFT *FightingNFTCaller) IsInWhitelist(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "isInWhitelist", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInWhitelist is a free data retrieval call binding the contract method 0x09fd8212.
//
// Solidity: function isInWhitelist(address account_) view returns(bool)
func (_FightingNFT *FightingNFTSession) IsInWhitelist(account_ common.Address) (bool, error) {
	return _FightingNFT.Contract.IsInWhitelist(&_FightingNFT.CallOpts, account_)
}

// IsInWhitelist is a free data retrieval call binding the contract method 0x09fd8212.
//
// Solidity: function isInWhitelist(address account_) view returns(bool)
func (_FightingNFT *FightingNFTCallerSession) IsInWhitelist(account_ common.Address) (bool, error) {
	return _FightingNFT.Contract.IsInWhitelist(&_FightingNFT.CallOpts, account_)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FightingNFT *FightingNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FightingNFT *FightingNFTSession) Name() (string, error) {
	return _FightingNFT.Contract.Name(&_FightingNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FightingNFT *FightingNFTCallerSession) Name() (string, error) {
	return _FightingNFT.Contract.Name(&_FightingNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FightingNFT *FightingNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FightingNFT *FightingNFTSession) Owner() (common.Address, error) {
	return _FightingNFT.Contract.Owner(&_FightingNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FightingNFT *FightingNFTCallerSession) Owner() (common.Address, error) {
	return _FightingNFT.Contract.Owner(&_FightingNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FightingNFT *FightingNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FightingNFT *FightingNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FightingNFT.Contract.OwnerOf(&_FightingNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FightingNFT *FightingNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FightingNFT.Contract.OwnerOf(&_FightingNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FightingNFT *FightingNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FightingNFT *FightingNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FightingNFT.Contract.SupportsInterface(&_FightingNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FightingNFT *FightingNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FightingNFT.Contract.SupportsInterface(&_FightingNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FightingNFT *FightingNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FightingNFT *FightingNFTSession) Symbol() (string, error) {
	return _FightingNFT.Contract.Symbol(&_FightingNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FightingNFT *FightingNFTCallerSession) Symbol() (string, error) {
	return _FightingNFT.Contract.Symbol(&_FightingNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_FightingNFT *FightingNFTCaller) TokenURI(opts *bind.CallOpts, tokenId_ *big.Int) (string, error) {
	var out []interface{}
	err := _FightingNFT.contract.Call(opts, &out, "tokenURI", tokenId_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_FightingNFT *FightingNFTSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _FightingNFT.Contract.TokenURI(&_FightingNFT.CallOpts, tokenId_)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_FightingNFT *FightingNFTCallerSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _FightingNFT.Contract.TokenURI(&_FightingNFT.CallOpts, tokenId_)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address account_) returns(bool)
func (_FightingNFT *FightingNFTTransactor) AddWhitelist(opts *bind.TransactOpts, account_ common.Address) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "addWhitelist", account_)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address account_) returns(bool)
func (_FightingNFT *FightingNFTSession) AddWhitelist(account_ common.Address) (*types.Transaction, error) {
	return _FightingNFT.Contract.AddWhitelist(&_FightingNFT.TransactOpts, account_)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address account_) returns(bool)
func (_FightingNFT *FightingNFTTransactorSession) AddWhitelist(account_ common.Address) (*types.Transaction, error) {
	return _FightingNFT.Contract.AddWhitelist(&_FightingNFT.TransactOpts, account_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_FightingNFT *FightingNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_FightingNFT *FightingNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FightingNFT.Contract.Approve(&_FightingNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_FightingNFT *FightingNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FightingNFT.Contract.Approve(&_FightingNFT.TransactOpts, to, tokenId)
}

// MintNFT is a paid mutator transaction binding the contract method 0x3c168eab.
//
// Solidity: function mintNFT(address player_, uint256 amount_) returns(uint256)
func (_FightingNFT *FightingNFTTransactor) MintNFT(opts *bind.TransactOpts, player_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "mintNFT", player_, amount_)
}

// MintNFT is a paid mutator transaction binding the contract method 0x3c168eab.
//
// Solidity: function mintNFT(address player_, uint256 amount_) returns(uint256)
func (_FightingNFT *FightingNFTSession) MintNFT(player_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _FightingNFT.Contract.MintNFT(&_FightingNFT.TransactOpts, player_, amount_)
}

// MintNFT is a paid mutator transaction binding the contract method 0x3c168eab.
//
// Solidity: function mintNFT(address player_, uint256 amount_) returns(uint256)
func (_FightingNFT *FightingNFTTransactorSession) MintNFT(player_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _FightingNFT.Contract.MintNFT(&_FightingNFT.TransactOpts, player_, amount_)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x78c8cda7.
//
// Solidity: function removeWhitelist(address account_) returns(bool)
func (_FightingNFT *FightingNFTTransactor) RemoveWhitelist(opts *bind.TransactOpts, account_ common.Address) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "removeWhitelist", account_)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x78c8cda7.
//
// Solidity: function removeWhitelist(address account_) returns(bool)
func (_FightingNFT *FightingNFTSession) RemoveWhitelist(account_ common.Address) (*types.Transaction, error) {
	return _FightingNFT.Contract.RemoveWhitelist(&_FightingNFT.TransactOpts, account_)
}

// RemoveWhitelist is a paid mutator transaction binding the contract method 0x78c8cda7.
//
// Solidity: function removeWhitelist(address account_) returns(bool)
func (_FightingNFT *FightingNFTTransactorSession) RemoveWhitelist(account_ common.Address) (*types.Transaction, error) {
	return _FightingNFT.Contract.RemoveWhitelist(&_FightingNFT.TransactOpts, account_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FightingNFT *FightingNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FightingNFT *FightingNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _FightingNFT.Contract.RenounceOwnership(&_FightingNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FightingNFT *FightingNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FightingNFT.Contract.RenounceOwnership(&_FightingNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FightingNFT *FightingNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FightingNFT *FightingNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FightingNFT.Contract.SafeTransferFrom(&_FightingNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FightingNFT *FightingNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FightingNFT.Contract.SafeTransferFrom(&_FightingNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_FightingNFT *FightingNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_FightingNFT *FightingNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _FightingNFT.Contract.SafeTransferFrom0(&_FightingNFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_FightingNFT *FightingNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _FightingNFT.Contract.SafeTransferFrom0(&_FightingNFT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FightingNFT *FightingNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FightingNFT *FightingNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FightingNFT.Contract.SetApprovalForAll(&_FightingNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FightingNFT *FightingNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FightingNFT.Contract.SetApprovalForAll(&_FightingNFT.TransactOpts, operator, approved)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_FightingNFT *FightingNFTTransactor) SetBaseUri(opts *bind.TransactOpts, baseUri_ string) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "setBaseUri", baseUri_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_FightingNFT *FightingNFTSession) SetBaseUri(baseUri_ string) (*types.Transaction, error) {
	return _FightingNFT.Contract.SetBaseUri(&_FightingNFT.TransactOpts, baseUri_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_FightingNFT *FightingNFTTransactorSession) SetBaseUri(baseUri_ string) (*types.Transaction, error) {
	return _FightingNFT.Contract.SetBaseUri(&_FightingNFT.TransactOpts, baseUri_)
}

// SetIsStartMint is a paid mutator transaction binding the contract method 0xd7114a37.
//
// Solidity: function setIsStartMint(bool isStartMint_) returns()
func (_FightingNFT *FightingNFTTransactor) SetIsStartMint(opts *bind.TransactOpts, isStartMint_ bool) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "setIsStartMint", isStartMint_)
}

// SetIsStartMint is a paid mutator transaction binding the contract method 0xd7114a37.
//
// Solidity: function setIsStartMint(bool isStartMint_) returns()
func (_FightingNFT *FightingNFTSession) SetIsStartMint(isStartMint_ bool) (*types.Transaction, error) {
	return _FightingNFT.Contract.SetIsStartMint(&_FightingNFT.TransactOpts, isStartMint_)
}

// SetIsStartMint is a paid mutator transaction binding the contract method 0xd7114a37.
//
// Solidity: function setIsStartMint(bool isStartMint_) returns()
func (_FightingNFT *FightingNFTTransactorSession) SetIsStartMint(isStartMint_ bool) (*types.Transaction, error) {
	return _FightingNFT.Contract.SetIsStartMint(&_FightingNFT.TransactOpts, isStartMint_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FightingNFT *FightingNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FightingNFT *FightingNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FightingNFT.Contract.TransferFrom(&_FightingNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FightingNFT *FightingNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FightingNFT.Contract.TransferFrom(&_FightingNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FightingNFT *FightingNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FightingNFT *FightingNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FightingNFT.Contract.TransferOwnership(&_FightingNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FightingNFT *FightingNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FightingNFT.Contract.TransferOwnership(&_FightingNFT.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FightingNFT *FightingNFTTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FightingNFT.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FightingNFT *FightingNFTSession) Withdraw() (*types.Transaction, error) {
	return _FightingNFT.Contract.Withdraw(&_FightingNFT.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FightingNFT *FightingNFTTransactorSession) Withdraw() (*types.Transaction, error) {
	return _FightingNFT.Contract.Withdraw(&_FightingNFT.TransactOpts)
}

// FightingNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FightingNFT contract.
type FightingNFTApprovalIterator struct {
	Event *FightingNFTApproval // Event containing the contract specifics and raw log

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
func (it *FightingNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FightingNFTApproval)
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
		it.Event = new(FightingNFTApproval)
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
func (it *FightingNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FightingNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FightingNFTApproval represents a Approval event raised by the FightingNFT contract.
type FightingNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FightingNFT *FightingNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*FightingNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FightingNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FightingNFTApprovalIterator{contract: _FightingNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FightingNFT *FightingNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FightingNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FightingNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FightingNFTApproval)
				if err := _FightingNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FightingNFT *FightingNFTFilterer) ParseApproval(log types.Log) (*FightingNFTApproval, error) {
	event := new(FightingNFTApproval)
	if err := _FightingNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FightingNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the FightingNFT contract.
type FightingNFTApprovalForAllIterator struct {
	Event *FightingNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *FightingNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FightingNFTApprovalForAll)
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
		it.Event = new(FightingNFTApprovalForAll)
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
func (it *FightingNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FightingNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FightingNFTApprovalForAll represents a ApprovalForAll event raised by the FightingNFT contract.
type FightingNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FightingNFT *FightingNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*FightingNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FightingNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &FightingNFTApprovalForAllIterator{contract: _FightingNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FightingNFT *FightingNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FightingNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FightingNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FightingNFTApprovalForAll)
				if err := _FightingNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FightingNFT *FightingNFTFilterer) ParseApprovalForAll(log types.Log) (*FightingNFTApprovalForAll, error) {
	event := new(FightingNFTApprovalForAll)
	if err := _FightingNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FightingNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FightingNFT contract.
type FightingNFTOwnershipTransferredIterator struct {
	Event *FightingNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FightingNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FightingNFTOwnershipTransferred)
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
		it.Event = new(FightingNFTOwnershipTransferred)
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
func (it *FightingNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FightingNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FightingNFTOwnershipTransferred represents a OwnershipTransferred event raised by the FightingNFT contract.
type FightingNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FightingNFT *FightingNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FightingNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FightingNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FightingNFTOwnershipTransferredIterator{contract: _FightingNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FightingNFT *FightingNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FightingNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FightingNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FightingNFTOwnershipTransferred)
				if err := _FightingNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FightingNFT *FightingNFTFilterer) ParseOwnershipTransferred(log types.Log) (*FightingNFTOwnershipTransferred, error) {
	event := new(FightingNFTOwnershipTransferred)
	if err := _FightingNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FightingNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FightingNFT contract.
type FightingNFTTransferIterator struct {
	Event *FightingNFTTransfer // Event containing the contract specifics and raw log

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
func (it *FightingNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FightingNFTTransfer)
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
		it.Event = new(FightingNFTTransfer)
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
func (it *FightingNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FightingNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FightingNFTTransfer represents a Transfer event raised by the FightingNFT contract.
type FightingNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FightingNFT *FightingNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*FightingNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FightingNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FightingNFTTransferIterator{contract: _FightingNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FightingNFT *FightingNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FightingNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FightingNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FightingNFTTransfer)
				if err := _FightingNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FightingNFT *FightingNFTFilterer) ParseTransfer(log types.Log) (*FightingNFTTransfer, error) {
	event := new(FightingNFTTransfer)
	if err := _FightingNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
