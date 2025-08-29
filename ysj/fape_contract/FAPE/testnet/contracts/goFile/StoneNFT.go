// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StoneNFT

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

// StoneNFTMetaData contains all meta data concerning the StoneNFT contract.
var StoneNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"addBlacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"at\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getCardNo\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getTokenIdByHHolders\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"isInBlacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStartMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cardNo_\",\"type\":\"uint256\"}],\"name\":\"mintNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"removeBlacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseUri_\",\"type\":\"string\"}],\"name\":\"setBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"genfee_\",\"type\":\"uint256\"}],\"name\":\"setGGeneratefee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isStartMint_\",\"type\":\"bool\"}],\"name\":\"setIsStartMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stone\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StoneNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use StoneNFTMetaData.ABI instead.
var StoneNFTABI = StoneNFTMetaData.ABI

// StoneNFT is an auto generated Go binding around an Ethereum contract.
type StoneNFT struct {
	StoneNFTCaller     // Read-only binding to the contract
	StoneNFTTransactor // Write-only binding to the contract
	StoneNFTFilterer   // Log filterer for contract events
}

// StoneNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoneNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoneNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoneNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoneNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoneNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoneNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoneNFTSession struct {
	Contract     *StoneNFT         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoneNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoneNFTCallerSession struct {
	Contract *StoneNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StoneNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoneNFTTransactorSession struct {
	Contract     *StoneNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StoneNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoneNFTRaw struct {
	Contract *StoneNFT // Generic contract binding to access the raw methods on
}

// StoneNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoneNFTCallerRaw struct {
	Contract *StoneNFTCaller // Generic read-only contract binding to access the raw methods on
}

// StoneNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoneNFTTransactorRaw struct {
	Contract *StoneNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStoneNFT creates a new instance of StoneNFT, bound to a specific deployed contract.
func NewStoneNFT(address common.Address, backend bind.ContractBackend) (*StoneNFT, error) {
	contract, err := bindStoneNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StoneNFT{StoneNFTCaller: StoneNFTCaller{contract: contract}, StoneNFTTransactor: StoneNFTTransactor{contract: contract}, StoneNFTFilterer: StoneNFTFilterer{contract: contract}}, nil
}

// NewStoneNFTCaller creates a new read-only instance of StoneNFT, bound to a specific deployed contract.
func NewStoneNFTCaller(address common.Address, caller bind.ContractCaller) (*StoneNFTCaller, error) {
	contract, err := bindStoneNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoneNFTCaller{contract: contract}, nil
}

// NewStoneNFTTransactor creates a new write-only instance of StoneNFT, bound to a specific deployed contract.
func NewStoneNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*StoneNFTTransactor, error) {
	contract, err := bindStoneNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoneNFTTransactor{contract: contract}, nil
}

// NewStoneNFTFilterer creates a new log filterer instance of StoneNFT, bound to a specific deployed contract.
func NewStoneNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*StoneNFTFilterer, error) {
	contract, err := bindStoneNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoneNFTFilterer{contract: contract}, nil
}

// bindStoneNFT binds a generic wrapper to an already deployed contract.
func bindStoneNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoneNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StoneNFT *StoneNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StoneNFT.Contract.StoneNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StoneNFT *StoneNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoneNFT.Contract.StoneNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StoneNFT *StoneNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StoneNFT.Contract.StoneNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StoneNFT *StoneNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StoneNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StoneNFT *StoneNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoneNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StoneNFT *StoneNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StoneNFT.Contract.contract.Transact(opts, method, params...)
}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index_) view returns(address)
func (_StoneNFT *StoneNFTCaller) At(opts *bind.CallOpts, index_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "at", index_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index_) view returns(address)
func (_StoneNFT *StoneNFTSession) At(index_ *big.Int) (common.Address, error) {
	return _StoneNFT.Contract.At(&_StoneNFT.CallOpts, index_)
}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index_) view returns(address)
func (_StoneNFT *StoneNFTCallerSession) At(index_ *big.Int) (common.Address, error) {
	return _StoneNFT.Contract.At(&_StoneNFT.CallOpts, index_)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_StoneNFT *StoneNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_StoneNFT *StoneNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _StoneNFT.Contract.BalanceOf(&_StoneNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_StoneNFT *StoneNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _StoneNFT.Contract.BalanceOf(&_StoneNFT.CallOpts, owner)
}

// GenerateFee is a free data retrieval call binding the contract method 0x2fe12dfb.
//
// Solidity: function generateFee() view returns(uint256)
func (_StoneNFT *StoneNFTCaller) GenerateFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "generateFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenerateFee is a free data retrieval call binding the contract method 0x2fe12dfb.
//
// Solidity: function generateFee() view returns(uint256)
func (_StoneNFT *StoneNFTSession) GenerateFee() (*big.Int, error) {
	return _StoneNFT.Contract.GenerateFee(&_StoneNFT.CallOpts)
}

// GenerateFee is a free data retrieval call binding the contract method 0x2fe12dfb.
//
// Solidity: function generateFee() view returns(uint256)
func (_StoneNFT *StoneNFTCallerSession) GenerateFee() (*big.Int, error) {
	return _StoneNFT.Contract.GenerateFee(&_StoneNFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_StoneNFT *StoneNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_StoneNFT *StoneNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _StoneNFT.Contract.GetApproved(&_StoneNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_StoneNFT *StoneNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _StoneNFT.Contract.GetApproved(&_StoneNFT.CallOpts, tokenId)
}

// GetCardNo is a free data retrieval call binding the contract method 0x518e130a.
//
// Solidity: function getCardNo(address account_) view returns(uint256[])
func (_StoneNFT *StoneNFTCaller) GetCardNo(opts *bind.CallOpts, account_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "getCardNo", account_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCardNo is a free data retrieval call binding the contract method 0x518e130a.
//
// Solidity: function getCardNo(address account_) view returns(uint256[])
func (_StoneNFT *StoneNFTSession) GetCardNo(account_ common.Address) ([]*big.Int, error) {
	return _StoneNFT.Contract.GetCardNo(&_StoneNFT.CallOpts, account_)
}

// GetCardNo is a free data retrieval call binding the contract method 0x518e130a.
//
// Solidity: function getCardNo(address account_) view returns(uint256[])
func (_StoneNFT *StoneNFTCallerSession) GetCardNo(account_ common.Address) ([]*big.Int, error) {
	return _StoneNFT.Contract.GetCardNo(&_StoneNFT.CallOpts, account_)
}

// GetTokenIdByHHolders is a free data retrieval call binding the contract method 0x1d660aae.
//
// Solidity: function getTokenIdByHHolders(address account_) view returns(uint256[])
func (_StoneNFT *StoneNFTCaller) GetTokenIdByHHolders(opts *bind.CallOpts, account_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "getTokenIdByHHolders", account_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenIdByHHolders is a free data retrieval call binding the contract method 0x1d660aae.
//
// Solidity: function getTokenIdByHHolders(address account_) view returns(uint256[])
func (_StoneNFT *StoneNFTSession) GetTokenIdByHHolders(account_ common.Address) ([]*big.Int, error) {
	return _StoneNFT.Contract.GetTokenIdByHHolders(&_StoneNFT.CallOpts, account_)
}

// GetTokenIdByHHolders is a free data retrieval call binding the contract method 0x1d660aae.
//
// Solidity: function getTokenIdByHHolders(address account_) view returns(uint256[])
func (_StoneNFT *StoneNFTCallerSession) GetTokenIdByHHolders(account_ common.Address) ([]*big.Int, error) {
	return _StoneNFT.Contract.GetTokenIdByHHolders(&_StoneNFT.CallOpts, account_)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_StoneNFT *StoneNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_StoneNFT *StoneNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _StoneNFT.Contract.IsApprovedForAll(&_StoneNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_StoneNFT *StoneNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _StoneNFT.Contract.IsApprovedForAll(&_StoneNFT.CallOpts, owner, operator)
}

// IsInBlacklist is a free data retrieval call binding the contract method 0x9caf9b00.
//
// Solidity: function isInBlacklist(address account_) view returns(bool)
func (_StoneNFT *StoneNFTCaller) IsInBlacklist(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "isInBlacklist", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInBlacklist is a free data retrieval call binding the contract method 0x9caf9b00.
//
// Solidity: function isInBlacklist(address account_) view returns(bool)
func (_StoneNFT *StoneNFTSession) IsInBlacklist(account_ common.Address) (bool, error) {
	return _StoneNFT.Contract.IsInBlacklist(&_StoneNFT.CallOpts, account_)
}

// IsInBlacklist is a free data retrieval call binding the contract method 0x9caf9b00.
//
// Solidity: function isInBlacklist(address account_) view returns(bool)
func (_StoneNFT *StoneNFTCallerSession) IsInBlacklist(account_ common.Address) (bool, error) {
	return _StoneNFT.Contract.IsInBlacklist(&_StoneNFT.CallOpts, account_)
}

// IsStartMint is a free data retrieval call binding the contract method 0xb174cca1.
//
// Solidity: function isStartMint() view returns(bool)
func (_StoneNFT *StoneNFTCaller) IsStartMint(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "isStartMint")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStartMint is a free data retrieval call binding the contract method 0xb174cca1.
//
// Solidity: function isStartMint() view returns(bool)
func (_StoneNFT *StoneNFTSession) IsStartMint() (bool, error) {
	return _StoneNFT.Contract.IsStartMint(&_StoneNFT.CallOpts)
}

// IsStartMint is a free data retrieval call binding the contract method 0xb174cca1.
//
// Solidity: function isStartMint() view returns(bool)
func (_StoneNFT *StoneNFTCallerSession) IsStartMint() (bool, error) {
	return _StoneNFT.Contract.IsStartMint(&_StoneNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StoneNFT *StoneNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StoneNFT *StoneNFTSession) Name() (string, error) {
	return _StoneNFT.Contract.Name(&_StoneNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StoneNFT *StoneNFTCallerSession) Name() (string, error) {
	return _StoneNFT.Contract.Name(&_StoneNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StoneNFT *StoneNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StoneNFT *StoneNFTSession) Owner() (common.Address, error) {
	return _StoneNFT.Contract.Owner(&_StoneNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StoneNFT *StoneNFTCallerSession) Owner() (common.Address, error) {
	return _StoneNFT.Contract.Owner(&_StoneNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StoneNFT *StoneNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StoneNFT *StoneNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _StoneNFT.Contract.OwnerOf(&_StoneNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StoneNFT *StoneNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _StoneNFT.Contract.OwnerOf(&_StoneNFT.CallOpts, tokenId)
}

// Stone is a free data retrieval call binding the contract method 0xcbb3e145.
//
// Solidity: function stone(uint256 ) view returns(uint256)
func (_StoneNFT *StoneNFTCaller) Stone(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "stone", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stone is a free data retrieval call binding the contract method 0xcbb3e145.
//
// Solidity: function stone(uint256 ) view returns(uint256)
func (_StoneNFT *StoneNFTSession) Stone(arg0 *big.Int) (*big.Int, error) {
	return _StoneNFT.Contract.Stone(&_StoneNFT.CallOpts, arg0)
}

// Stone is a free data retrieval call binding the contract method 0xcbb3e145.
//
// Solidity: function stone(uint256 ) view returns(uint256)
func (_StoneNFT *StoneNFTCallerSession) Stone(arg0 *big.Int) (*big.Int, error) {
	return _StoneNFT.Contract.Stone(&_StoneNFT.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StoneNFT *StoneNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StoneNFT *StoneNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StoneNFT.Contract.SupportsInterface(&_StoneNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StoneNFT *StoneNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StoneNFT.Contract.SupportsInterface(&_StoneNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StoneNFT *StoneNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StoneNFT *StoneNFTSession) Symbol() (string, error) {
	return _StoneNFT.Contract.Symbol(&_StoneNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StoneNFT *StoneNFTCallerSession) Symbol() (string, error) {
	return _StoneNFT.Contract.Symbol(&_StoneNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_StoneNFT *StoneNFTCaller) TokenURI(opts *bind.CallOpts, tokenId_ *big.Int) (string, error) {
	var out []interface{}
	err := _StoneNFT.contract.Call(opts, &out, "tokenURI", tokenId_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_StoneNFT *StoneNFTSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _StoneNFT.Contract.TokenURI(&_StoneNFT.CallOpts, tokenId_)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_StoneNFT *StoneNFTCallerSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _StoneNFT.Contract.TokenURI(&_StoneNFT.CallOpts, tokenId_)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x9cfe42da.
//
// Solidity: function addBlacklist(address account_) returns(bool)
func (_StoneNFT *StoneNFTTransactor) AddBlacklist(opts *bind.TransactOpts, account_ common.Address) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "addBlacklist", account_)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x9cfe42da.
//
// Solidity: function addBlacklist(address account_) returns(bool)
func (_StoneNFT *StoneNFTSession) AddBlacklist(account_ common.Address) (*types.Transaction, error) {
	return _StoneNFT.Contract.AddBlacklist(&_StoneNFT.TransactOpts, account_)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x9cfe42da.
//
// Solidity: function addBlacklist(address account_) returns(bool)
func (_StoneNFT *StoneNFTTransactorSession) AddBlacklist(account_ common.Address) (*types.Transaction, error) {
	return _StoneNFT.Contract.AddBlacklist(&_StoneNFT.TransactOpts, account_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_StoneNFT *StoneNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_StoneNFT *StoneNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoneNFT.Contract.Approve(&_StoneNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_StoneNFT *StoneNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoneNFT.Contract.Approve(&_StoneNFT.TransactOpts, to, tokenId)
}

// MintNFT is a paid mutator transaction binding the contract method 0x3c168eab.
//
// Solidity: function mintNFT(address account_, uint256 cardNo_) returns(uint256)
func (_StoneNFT *StoneNFTTransactor) MintNFT(opts *bind.TransactOpts, account_ common.Address, cardNo_ *big.Int) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "mintNFT", account_, cardNo_)
}

// MintNFT is a paid mutator transaction binding the contract method 0x3c168eab.
//
// Solidity: function mintNFT(address account_, uint256 cardNo_) returns(uint256)
func (_StoneNFT *StoneNFTSession) MintNFT(account_ common.Address, cardNo_ *big.Int) (*types.Transaction, error) {
	return _StoneNFT.Contract.MintNFT(&_StoneNFT.TransactOpts, account_, cardNo_)
}

// MintNFT is a paid mutator transaction binding the contract method 0x3c168eab.
//
// Solidity: function mintNFT(address account_, uint256 cardNo_) returns(uint256)
func (_StoneNFT *StoneNFTTransactorSession) MintNFT(account_ common.Address, cardNo_ *big.Int) (*types.Transaction, error) {
	return _StoneNFT.Contract.MintNFT(&_StoneNFT.TransactOpts, account_, cardNo_)
}

// RemoveBlacklist is a paid mutator transaction binding the contract method 0xeb91e651.
//
// Solidity: function removeBlacklist(address account_) returns(bool)
func (_StoneNFT *StoneNFTTransactor) RemoveBlacklist(opts *bind.TransactOpts, account_ common.Address) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "removeBlacklist", account_)
}

// RemoveBlacklist is a paid mutator transaction binding the contract method 0xeb91e651.
//
// Solidity: function removeBlacklist(address account_) returns(bool)
func (_StoneNFT *StoneNFTSession) RemoveBlacklist(account_ common.Address) (*types.Transaction, error) {
	return _StoneNFT.Contract.RemoveBlacklist(&_StoneNFT.TransactOpts, account_)
}

// RemoveBlacklist is a paid mutator transaction binding the contract method 0xeb91e651.
//
// Solidity: function removeBlacklist(address account_) returns(bool)
func (_StoneNFT *StoneNFTTransactorSession) RemoveBlacklist(account_ common.Address) (*types.Transaction, error) {
	return _StoneNFT.Contract.RemoveBlacklist(&_StoneNFT.TransactOpts, account_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StoneNFT *StoneNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StoneNFT *StoneNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _StoneNFT.Contract.RenounceOwnership(&_StoneNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StoneNFT *StoneNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StoneNFT.Contract.RenounceOwnership(&_StoneNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_StoneNFT *StoneNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_StoneNFT *StoneNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoneNFT.Contract.SafeTransferFrom(&_StoneNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_StoneNFT *StoneNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoneNFT.Contract.SafeTransferFrom(&_StoneNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_StoneNFT *StoneNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_StoneNFT *StoneNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _StoneNFT.Contract.SafeTransferFrom0(&_StoneNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_StoneNFT *StoneNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _StoneNFT.Contract.SafeTransferFrom0(&_StoneNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StoneNFT *StoneNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StoneNFT *StoneNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _StoneNFT.Contract.SetApprovalForAll(&_StoneNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StoneNFT *StoneNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _StoneNFT.Contract.SetApprovalForAll(&_StoneNFT.TransactOpts, operator, approved)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_StoneNFT *StoneNFTTransactor) SetBaseUri(opts *bind.TransactOpts, baseUri_ string) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "setBaseUri", baseUri_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_StoneNFT *StoneNFTSession) SetBaseUri(baseUri_ string) (*types.Transaction, error) {
	return _StoneNFT.Contract.SetBaseUri(&_StoneNFT.TransactOpts, baseUri_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_StoneNFT *StoneNFTTransactorSession) SetBaseUri(baseUri_ string) (*types.Transaction, error) {
	return _StoneNFT.Contract.SetBaseUri(&_StoneNFT.TransactOpts, baseUri_)
}

// SetGGeneratefee is a paid mutator transaction binding the contract method 0x9213c512.
//
// Solidity: function setGGeneratefee(uint256 genfee_) returns()
func (_StoneNFT *StoneNFTTransactor) SetGGeneratefee(opts *bind.TransactOpts, genfee_ *big.Int) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "setGGeneratefee", genfee_)
}

// SetGGeneratefee is a paid mutator transaction binding the contract method 0x9213c512.
//
// Solidity: function setGGeneratefee(uint256 genfee_) returns()
func (_StoneNFT *StoneNFTSession) SetGGeneratefee(genfee_ *big.Int) (*types.Transaction, error) {
	return _StoneNFT.Contract.SetGGeneratefee(&_StoneNFT.TransactOpts, genfee_)
}

// SetGGeneratefee is a paid mutator transaction binding the contract method 0x9213c512.
//
// Solidity: function setGGeneratefee(uint256 genfee_) returns()
func (_StoneNFT *StoneNFTTransactorSession) SetGGeneratefee(genfee_ *big.Int) (*types.Transaction, error) {
	return _StoneNFT.Contract.SetGGeneratefee(&_StoneNFT.TransactOpts, genfee_)
}

// SetIsStartMint is a paid mutator transaction binding the contract method 0xd7114a37.
//
// Solidity: function setIsStartMint(bool isStartMint_) returns()
func (_StoneNFT *StoneNFTTransactor) SetIsStartMint(opts *bind.TransactOpts, isStartMint_ bool) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "setIsStartMint", isStartMint_)
}

// SetIsStartMint is a paid mutator transaction binding the contract method 0xd7114a37.
//
// Solidity: function setIsStartMint(bool isStartMint_) returns()
func (_StoneNFT *StoneNFTSession) SetIsStartMint(isStartMint_ bool) (*types.Transaction, error) {
	return _StoneNFT.Contract.SetIsStartMint(&_StoneNFT.TransactOpts, isStartMint_)
}

// SetIsStartMint is a paid mutator transaction binding the contract method 0xd7114a37.
//
// Solidity: function setIsStartMint(bool isStartMint_) returns()
func (_StoneNFT *StoneNFTTransactorSession) SetIsStartMint(isStartMint_ bool) (*types.Transaction, error) {
	return _StoneNFT.Contract.SetIsStartMint(&_StoneNFT.TransactOpts, isStartMint_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_StoneNFT *StoneNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_StoneNFT *StoneNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoneNFT.Contract.TransferFrom(&_StoneNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_StoneNFT *StoneNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoneNFT.Contract.TransferFrom(&_StoneNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StoneNFT *StoneNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StoneNFT *StoneNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StoneNFT.Contract.TransferOwnership(&_StoneNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StoneNFT *StoneNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StoneNFT.Contract.TransferOwnership(&_StoneNFT.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StoneNFT *StoneNFTTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoneNFT.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StoneNFT *StoneNFTSession) Withdraw() (*types.Transaction, error) {
	return _StoneNFT.Contract.Withdraw(&_StoneNFT.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StoneNFT *StoneNFTTransactorSession) Withdraw() (*types.Transaction, error) {
	return _StoneNFT.Contract.Withdraw(&_StoneNFT.TransactOpts)
}

// StoneNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StoneNFT contract.
type StoneNFTApprovalIterator struct {
	Event *StoneNFTApproval // Event containing the contract specifics and raw log

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
func (it *StoneNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoneNFTApproval)
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
		it.Event = new(StoneNFTApproval)
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
func (it *StoneNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoneNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoneNFTApproval represents a Approval event raised by the StoneNFT contract.
type StoneNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StoneNFT *StoneNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*StoneNFTApprovalIterator, error) {

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

	logs, sub, err := _StoneNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoneNFTApprovalIterator{contract: _StoneNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StoneNFT *StoneNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StoneNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _StoneNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoneNFTApproval)
				if err := _StoneNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StoneNFT *StoneNFTFilterer) ParseApproval(log types.Log) (*StoneNFTApproval, error) {
	event := new(StoneNFTApproval)
	if err := _StoneNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoneNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the StoneNFT contract.
type StoneNFTApprovalForAllIterator struct {
	Event *StoneNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *StoneNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoneNFTApprovalForAll)
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
		it.Event = new(StoneNFTApprovalForAll)
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
func (it *StoneNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoneNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoneNFTApprovalForAll represents a ApprovalForAll event raised by the StoneNFT contract.
type StoneNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StoneNFT *StoneNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*StoneNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _StoneNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StoneNFTApprovalForAllIterator{contract: _StoneNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StoneNFT *StoneNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StoneNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _StoneNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoneNFTApprovalForAll)
				if err := _StoneNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_StoneNFT *StoneNFTFilterer) ParseApprovalForAll(log types.Log) (*StoneNFTApprovalForAll, error) {
	event := new(StoneNFTApprovalForAll)
	if err := _StoneNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoneNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StoneNFT contract.
type StoneNFTOwnershipTransferredIterator struct {
	Event *StoneNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StoneNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoneNFTOwnershipTransferred)
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
		it.Event = new(StoneNFTOwnershipTransferred)
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
func (it *StoneNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoneNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoneNFTOwnershipTransferred represents a OwnershipTransferred event raised by the StoneNFT contract.
type StoneNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StoneNFT *StoneNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StoneNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StoneNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StoneNFTOwnershipTransferredIterator{contract: _StoneNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StoneNFT *StoneNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StoneNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StoneNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoneNFTOwnershipTransferred)
				if err := _StoneNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StoneNFT *StoneNFTFilterer) ParseOwnershipTransferred(log types.Log) (*StoneNFTOwnershipTransferred, error) {
	event := new(StoneNFTOwnershipTransferred)
	if err := _StoneNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoneNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StoneNFT contract.
type StoneNFTTransferIterator struct {
	Event *StoneNFTTransfer // Event containing the contract specifics and raw log

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
func (it *StoneNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoneNFTTransfer)
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
		it.Event = new(StoneNFTTransfer)
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
func (it *StoneNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoneNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoneNFTTransfer represents a Transfer event raised by the StoneNFT contract.
type StoneNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StoneNFT *StoneNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*StoneNFTTransferIterator, error) {

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

	logs, sub, err := _StoneNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoneNFTTransferIterator{contract: _StoneNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StoneNFT *StoneNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StoneNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _StoneNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoneNFTTransfer)
				if err := _StoneNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StoneNFT *StoneNFTFilterer) ParseTransfer(log types.Log) (*StoneNFTTransfer, error) {
	event := new(StoneNFTTransfer)
	if err := _StoneNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
