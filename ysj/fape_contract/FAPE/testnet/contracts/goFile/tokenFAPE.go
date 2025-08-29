// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenFAPE

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

// TokenFAPEMetaData contains all meta data concerning the TokenFAPE contract.
var TokenFAPEMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients_\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values_\",\"type\":\"uint256[]\"}],\"name\":\"airTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundsLockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"getAccountByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAccount\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenFAPEABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenFAPEMetaData.ABI instead.
var TokenFAPEABI = TokenFAPEMetaData.ABI

// TokenFAPE is an auto generated Go binding around an Ethereum contract.
type TokenFAPE struct {
	TokenFAPECaller     // Read-only binding to the contract
	TokenFAPETransactor // Write-only binding to the contract
	TokenFAPEFilterer   // Log filterer for contract events
}

// TokenFAPECaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenFAPECaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFAPETransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenFAPETransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFAPEFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFAPEFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFAPESession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenFAPESession struct {
	Contract     *TokenFAPE        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenFAPECallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenFAPECallerSession struct {
	Contract *TokenFAPECaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenFAPETransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenFAPETransactorSession struct {
	Contract     *TokenFAPETransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenFAPERaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenFAPERaw struct {
	Contract *TokenFAPE // Generic contract binding to access the raw methods on
}

// TokenFAPECallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenFAPECallerRaw struct {
	Contract *TokenFAPECaller // Generic read-only contract binding to access the raw methods on
}

// TokenFAPETransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenFAPETransactorRaw struct {
	Contract *TokenFAPETransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenFAPE creates a new instance of TokenFAPE, bound to a specific deployed contract.
func NewTokenFAPE(address common.Address, backend bind.ContractBackend) (*TokenFAPE, error) {
	contract, err := bindTokenFAPE(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenFAPE{TokenFAPECaller: TokenFAPECaller{contract: contract}, TokenFAPETransactor: TokenFAPETransactor{contract: contract}, TokenFAPEFilterer: TokenFAPEFilterer{contract: contract}}, nil
}

// NewTokenFAPECaller creates a new read-only instance of TokenFAPE, bound to a specific deployed contract.
func NewTokenFAPECaller(address common.Address, caller bind.ContractCaller) (*TokenFAPECaller, error) {
	contract, err := bindTokenFAPE(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFAPECaller{contract: contract}, nil
}

// NewTokenFAPETransactor creates a new write-only instance of TokenFAPE, bound to a specific deployed contract.
func NewTokenFAPETransactor(address common.Address, transactor bind.ContractTransactor) (*TokenFAPETransactor, error) {
	contract, err := bindTokenFAPE(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFAPETransactor{contract: contract}, nil
}

// NewTokenFAPEFilterer creates a new log filterer instance of TokenFAPE, bound to a specific deployed contract.
func NewTokenFAPEFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFAPEFilterer, error) {
	contract, err := bindTokenFAPE(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFAPEFilterer{contract: contract}, nil
}

// bindTokenFAPE binds a generic wrapper to an already deployed contract.
func bindTokenFAPE(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenFAPEABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFAPE *TokenFAPERaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFAPE.Contract.TokenFAPECaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFAPE *TokenFAPERaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFAPE.Contract.TokenFAPETransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFAPE *TokenFAPERaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFAPE.Contract.TokenFAPETransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFAPE *TokenFAPECallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFAPE.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFAPE *TokenFAPETransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFAPE.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFAPE *TokenFAPETransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFAPE.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenFAPE *TokenFAPECaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenFAPE.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenFAPE *TokenFAPESession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenFAPE.Contract.Allowance(&_TokenFAPE.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenFAPE *TokenFAPECallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenFAPE.Contract.Allowance(&_TokenFAPE.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenFAPE *TokenFAPECaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenFAPE.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenFAPE *TokenFAPESession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TokenFAPE.Contract.BalanceOf(&_TokenFAPE.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenFAPE *TokenFAPECallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TokenFAPE.Contract.BalanceOf(&_TokenFAPE.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenFAPE *TokenFAPECaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenFAPE.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenFAPE *TokenFAPESession) Decimals() (uint8, error) {
	return _TokenFAPE.Contract.Decimals(&_TokenFAPE.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenFAPE *TokenFAPECallerSession) Decimals() (uint8, error) {
	return _TokenFAPE.Contract.Decimals(&_TokenFAPE.CallOpts)
}

// FundsLockTime is a free data retrieval call binding the contract method 0x46fd3782.
//
// Solidity: function fundsLockTime() view returns(uint256)
func (_TokenFAPE *TokenFAPECaller) FundsLockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenFAPE.contract.Call(opts, &out, "fundsLockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundsLockTime is a free data retrieval call binding the contract method 0x46fd3782.
//
// Solidity: function fundsLockTime() view returns(uint256)
func (_TokenFAPE *TokenFAPESession) FundsLockTime() (*big.Int, error) {
	return _TokenFAPE.Contract.FundsLockTime(&_TokenFAPE.CallOpts)
}

// FundsLockTime is a free data retrieval call binding the contract method 0x46fd3782.
//
// Solidity: function fundsLockTime() view returns(uint256)
func (_TokenFAPE *TokenFAPECallerSession) FundsLockTime() (*big.Int, error) {
	return _TokenFAPE.Contract.FundsLockTime(&_TokenFAPE.CallOpts)
}

// GetAccountByIndex is a free data retrieval call binding the contract method 0xedc01752.
//
// Solidity: function getAccountByIndex(uint256 index_) view returns(address)
func (_TokenFAPE *TokenFAPECaller) GetAccountByIndex(opts *bind.CallOpts, index_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenFAPE.contract.Call(opts, &out, "getAccountByIndex", index_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountByIndex is a free data retrieval call binding the contract method 0xedc01752.
//
// Solidity: function getAccountByIndex(uint256 index_) view returns(address)
func (_TokenFAPE *TokenFAPESession) GetAccountByIndex(index_ *big.Int) (common.Address, error) {
	return _TokenFAPE.Contract.GetAccountByIndex(&_TokenFAPE.CallOpts, index_)
}

// GetAccountByIndex is a free data retrieval call binding the contract method 0xedc01752.
//
// Solidity: function getAccountByIndex(uint256 index_) view returns(address)
func (_TokenFAPE *TokenFAPECallerSession) GetAccountByIndex(index_ *big.Int) (common.Address, error) {
	return _TokenFAPE.Contract.GetAccountByIndex(&_TokenFAPE.CallOpts, index_)
}

// GetAccountLength is a free data retrieval call binding the contract method 0x22de3c91.
//
// Solidity: function getAccountLength() view returns(uint256)
func (_TokenFAPE *TokenFAPECaller) GetAccountLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenFAPE.contract.Call(opts, &out, "getAccountLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountLength is a free data retrieval call binding the contract method 0x22de3c91.
//
// Solidity: function getAccountLength() view returns(uint256)
func (_TokenFAPE *TokenFAPESession) GetAccountLength() (*big.Int, error) {
	return _TokenFAPE.Contract.GetAccountLength(&_TokenFAPE.CallOpts)
}

// GetAccountLength is a free data retrieval call binding the contract method 0x22de3c91.
//
// Solidity: function getAccountLength() view returns(uint256)
func (_TokenFAPE *TokenFAPECallerSession) GetAccountLength() (*big.Int, error) {
	return _TokenFAPE.Contract.GetAccountLength(&_TokenFAPE.CallOpts)
}

// GetAllAccount is a free data retrieval call binding the contract method 0x9dc5368d.
//
// Solidity: function getAllAccount() view returns(address[])
func (_TokenFAPE *TokenFAPECaller) GetAllAccount(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenFAPE.contract.Call(opts, &out, "getAllAccount")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllAccount is a free data retrieval call binding the contract method 0x9dc5368d.
//
// Solidity: function getAllAccount() view returns(address[])
func (_TokenFAPE *TokenFAPESession) GetAllAccount() ([]common.Address, error) {
	return _TokenFAPE.Contract.GetAllAccount(&_TokenFAPE.CallOpts)
}

// GetAllAccount is a free data retrieval call binding the contract method 0x9dc5368d.
//
// Solidity: function getAllAccount() view returns(address[])
func (_TokenFAPE *TokenFAPECallerSession) GetAllAccount() ([]common.Address, error) {
	return _TokenFAPE.Contract.GetAllAccount(&_TokenFAPE.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenFAPE *TokenFAPECaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenFAPE.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenFAPE *TokenFAPESession) Name() (string, error) {
	return _TokenFAPE.Contract.Name(&_TokenFAPE.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenFAPE *TokenFAPECallerSession) Name() (string, error) {
	return _TokenFAPE.Contract.Name(&_TokenFAPE.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFAPE *TokenFAPECaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFAPE.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFAPE *TokenFAPESession) Owner() (common.Address, error) {
	return _TokenFAPE.Contract.Owner(&_TokenFAPE.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFAPE *TokenFAPECallerSession) Owner() (common.Address, error) {
	return _TokenFAPE.Contract.Owner(&_TokenFAPE.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenFAPE *TokenFAPECaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenFAPE.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenFAPE *TokenFAPESession) Symbol() (string, error) {
	return _TokenFAPE.Contract.Symbol(&_TokenFAPE.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenFAPE *TokenFAPECallerSession) Symbol() (string, error) {
	return _TokenFAPE.Contract.Symbol(&_TokenFAPE.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenFAPE *TokenFAPECaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenFAPE.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenFAPE *TokenFAPESession) TotalSupply() (*big.Int, error) {
	return _TokenFAPE.Contract.TotalSupply(&_TokenFAPE.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenFAPE *TokenFAPECallerSession) TotalSupply() (*big.Int, error) {
	return _TokenFAPE.Contract.TotalSupply(&_TokenFAPE.CallOpts)
}

// AirTransfer is a paid mutator transaction binding the contract method 0x74ef1d3c.
//
// Solidity: function airTransfer(address[] recipients_, uint256[] values_) returns(bool)
func (_TokenFAPE *TokenFAPETransactor) AirTransfer(opts *bind.TransactOpts, recipients_ []common.Address, values_ []*big.Int) (*types.Transaction, error) {
	return _TokenFAPE.contract.Transact(opts, "airTransfer", recipients_, values_)
}

// AirTransfer is a paid mutator transaction binding the contract method 0x74ef1d3c.
//
// Solidity: function airTransfer(address[] recipients_, uint256[] values_) returns(bool)
func (_TokenFAPE *TokenFAPESession) AirTransfer(recipients_ []common.Address, values_ []*big.Int) (*types.Transaction, error) {
	return _TokenFAPE.Contract.AirTransfer(&_TokenFAPE.TransactOpts, recipients_, values_)
}

// AirTransfer is a paid mutator transaction binding the contract method 0x74ef1d3c.
//
// Solidity: function airTransfer(address[] recipients_, uint256[] values_) returns(bool)
func (_TokenFAPE *TokenFAPETransactorSession) AirTransfer(recipients_ []common.Address, values_ []*big.Int) (*types.Transaction, error) {
	return _TokenFAPE.Contract.AirTransfer(&_TokenFAPE.TransactOpts, recipients_, values_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenFAPE *TokenFAPETransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenFAPE *TokenFAPESession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.Contract.Approve(&_TokenFAPE.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenFAPE *TokenFAPETransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.Contract.Approve(&_TokenFAPE.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenFAPE *TokenFAPETransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenFAPE *TokenFAPESession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.Contract.DecreaseAllowance(&_TokenFAPE.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenFAPE *TokenFAPETransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.Contract.DecreaseAllowance(&_TokenFAPE.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenFAPE *TokenFAPETransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenFAPE *TokenFAPESession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.Contract.IncreaseAllowance(&_TokenFAPE.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenFAPE *TokenFAPETransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.Contract.IncreaseAllowance(&_TokenFAPE.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFAPE *TokenFAPETransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFAPE.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFAPE *TokenFAPESession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenFAPE.Contract.RenounceOwnership(&_TokenFAPE.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFAPE *TokenFAPETransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenFAPE.Contract.RenounceOwnership(&_TokenFAPE.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TokenFAPE *TokenFAPETransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TokenFAPE *TokenFAPESession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.Contract.Transfer(&_TokenFAPE.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TokenFAPE *TokenFAPETransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.Contract.Transfer(&_TokenFAPE.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TokenFAPE *TokenFAPETransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TokenFAPE *TokenFAPESession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.Contract.TransferFrom(&_TokenFAPE.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TokenFAPE *TokenFAPETransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFAPE.Contract.TransferFrom(&_TokenFAPE.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFAPE *TokenFAPETransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenFAPE.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFAPE *TokenFAPESession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenFAPE.Contract.TransferOwnership(&_TokenFAPE.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFAPE *TokenFAPETransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenFAPE.Contract.TransferOwnership(&_TokenFAPE.TransactOpts, newOwner)
}

// TokenFAPEApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TokenFAPE contract.
type TokenFAPEApprovalIterator struct {
	Event *TokenFAPEApproval // Event containing the contract specifics and raw log

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
func (it *TokenFAPEApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFAPEApproval)
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
		it.Event = new(TokenFAPEApproval)
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
func (it *TokenFAPEApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFAPEApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFAPEApproval represents a Approval event raised by the TokenFAPE contract.
type TokenFAPEApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenFAPE *TokenFAPEFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenFAPEApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenFAPE.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenFAPEApprovalIterator{contract: _TokenFAPE.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenFAPE *TokenFAPEFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenFAPEApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenFAPE.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFAPEApproval)
				if err := _TokenFAPE.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_TokenFAPE *TokenFAPEFilterer) ParseApproval(log types.Log) (*TokenFAPEApproval, error) {
	event := new(TokenFAPEApproval)
	if err := _TokenFAPE.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFAPEOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenFAPE contract.
type TokenFAPEOwnershipTransferredIterator struct {
	Event *TokenFAPEOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenFAPEOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFAPEOwnershipTransferred)
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
		it.Event = new(TokenFAPEOwnershipTransferred)
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
func (it *TokenFAPEOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFAPEOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFAPEOwnershipTransferred represents a OwnershipTransferred event raised by the TokenFAPE contract.
type TokenFAPEOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenFAPE *TokenFAPEFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenFAPEOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenFAPE.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenFAPEOwnershipTransferredIterator{contract: _TokenFAPE.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenFAPE *TokenFAPEFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenFAPEOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenFAPE.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFAPEOwnershipTransferred)
				if err := _TokenFAPE.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenFAPE *TokenFAPEFilterer) ParseOwnershipTransferred(log types.Log) (*TokenFAPEOwnershipTransferred, error) {
	event := new(TokenFAPEOwnershipTransferred)
	if err := _TokenFAPE.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFAPETransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenFAPE contract.
type TokenFAPETransferIterator struct {
	Event *TokenFAPETransfer // Event containing the contract specifics and raw log

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
func (it *TokenFAPETransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFAPETransfer)
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
		it.Event = new(TokenFAPETransfer)
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
func (it *TokenFAPETransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFAPETransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFAPETransfer represents a Transfer event raised by the TokenFAPE contract.
type TokenFAPETransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenFAPE *TokenFAPEFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenFAPETransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFAPE.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenFAPETransferIterator{contract: _TokenFAPE.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenFAPE *TokenFAPEFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenFAPETransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFAPE.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFAPETransfer)
				if err := _TokenFAPE.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_TokenFAPE *TokenFAPEFilterer) ParseTransfer(log types.Log) (*TokenFAPETransfer, error) {
	event := new(TokenFAPETransfer)
	if err := _TokenFAPE.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
