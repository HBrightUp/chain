// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AirdropABI is the input ABI used to generate the binding from.
const AirdropABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"notifyAirdropAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notifyAirdropAmounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Airdrop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allAddressLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"amountInfoByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"flag\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Airdrop is an auto generated Go binding around an Ethereum contract.
type Airdrop struct {
	AirdropCaller     // Read-only binding to the contract
	AirdropTransactor // Write-only binding to the contract
	AirdropFilterer   // Log filterer for contract events
}

// AirdropCaller is an auto generated read-only Go binding around an Ethereum contract.
type AirdropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AirdropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AirdropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AirdropSession struct {
	Contract     *Airdrop          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AirdropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AirdropCallerSession struct {
	Contract *AirdropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AirdropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AirdropTransactorSession struct {
	Contract     *AirdropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AirdropRaw is an auto generated low-level Go binding around an Ethereum contract.
type AirdropRaw struct {
	Contract *Airdrop // Generic contract binding to access the raw methods on
}

// AirdropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AirdropCallerRaw struct {
	Contract *AirdropCaller // Generic read-only contract binding to access the raw methods on
}

// AirdropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AirdropTransactorRaw struct {
	Contract *AirdropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAirdrop creates a new instance of Airdrop, bound to a specific deployed contract.
func NewAirdrop(address common.Address, backend bind.ContractBackend) (*Airdrop, error) {
	contract, err := bindAirdrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Airdrop{AirdropCaller: AirdropCaller{contract: contract}, AirdropTransactor: AirdropTransactor{contract: contract}, AirdropFilterer: AirdropFilterer{contract: contract}}, nil
}

// NewAirdropCaller creates a new read-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropCaller(address common.Address, caller bind.ContractCaller) (*AirdropCaller, error) {
	contract, err := bindAirdrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropCaller{contract: contract}, nil
}

// NewAirdropTransactor creates a new write-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropTransactor(address common.Address, transactor bind.ContractTransactor) (*AirdropTransactor, error) {
	contract, err := bindAirdrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropTransactor{contract: contract}, nil
}

// NewAirdropFilterer creates a new log filterer instance of Airdrop, bound to a specific deployed contract.
func NewAirdropFilterer(address common.Address, filterer bind.ContractFilterer) (*AirdropFilterer, error) {
	contract, err := bindAirdrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AirdropFilterer{contract: contract}, nil
}

// bindAirdrop binds a generic wrapper to an already deployed contract.
func bindAirdrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AirdropABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.AirdropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transact(opts, method, params...)
}

// AllAddress is a free data retrieval call binding the contract method 0xa8cef27a.
//
// Solidity: function allAddress(uint256 ) view returns(address)
func (_Airdrop *AirdropCaller) AllAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "allAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllAddress is a free data retrieval call binding the contract method 0xa8cef27a.
//
// Solidity: function allAddress(uint256 ) view returns(address)
func (_Airdrop *AirdropSession) AllAddress(arg0 *big.Int) (common.Address, error) {
	return _Airdrop.Contract.AllAddress(&_Airdrop.CallOpts, arg0)
}

// AllAddress is a free data retrieval call binding the contract method 0xa8cef27a.
//
// Solidity: function allAddress(uint256 ) view returns(address)
func (_Airdrop *AirdropCallerSession) AllAddress(arg0 *big.Int) (common.Address, error) {
	return _Airdrop.Contract.AllAddress(&_Airdrop.CallOpts, arg0)
}

// AllAddressLength is a free data retrieval call binding the contract method 0x8b4cbe45.
//
// Solidity: function allAddressLength() view returns(uint256)
func (_Airdrop *AirdropCaller) AllAddressLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "allAddressLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllAddressLength is a free data retrieval call binding the contract method 0x8b4cbe45.
//
// Solidity: function allAddressLength() view returns(uint256)
func (_Airdrop *AirdropSession) AllAddressLength() (*big.Int, error) {
	return _Airdrop.Contract.AllAddressLength(&_Airdrop.CallOpts)
}

// AllAddressLength is a free data retrieval call binding the contract method 0x8b4cbe45.
//
// Solidity: function allAddressLength() view returns(uint256)
func (_Airdrop *AirdropCallerSession) AllAddressLength() (*big.Int, error) {
	return _Airdrop.Contract.AllAddressLength(&_Airdrop.CallOpts)
}

// AmountInfoByAddress is a free data retrieval call binding the contract method 0xc7816312.
//
// Solidity: function amountInfoByAddress(address ) view returns(uint256 amount, uint8 flag)
func (_Airdrop *AirdropCaller) AmountInfoByAddress(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	Flag   uint8
}, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "amountInfoByAddress", arg0)

	outstruct := new(struct {
		Amount *big.Int
		Flag   uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Flag = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// AmountInfoByAddress is a free data retrieval call binding the contract method 0xc7816312.
//
// Solidity: function amountInfoByAddress(address ) view returns(uint256 amount, uint8 flag)
func (_Airdrop *AirdropSession) AmountInfoByAddress(arg0 common.Address) (struct {
	Amount *big.Int
	Flag   uint8
}, error) {
	return _Airdrop.Contract.AmountInfoByAddress(&_Airdrop.CallOpts, arg0)
}

// AmountInfoByAddress is a free data retrieval call binding the contract method 0xc7816312.
//
// Solidity: function amountInfoByAddress(address ) view returns(uint256 amount, uint8 flag)
func (_Airdrop *AirdropCallerSession) AmountInfoByAddress(arg0 common.Address) (struct {
	Amount *big.Int
	Flag   uint8
}, error) {
	return _Airdrop.Contract.AmountInfoByAddress(&_Airdrop.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Airdrop *AirdropCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Airdrop *AirdropSession) Owner() (common.Address, error) {
	return _Airdrop.Contract.Owner(&_Airdrop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Airdrop *AirdropCallerSession) Owner() (common.Address, error) {
	return _Airdrop.Contract.Owner(&_Airdrop.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Airdrop *AirdropCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Airdrop *AirdropSession) Token() (common.Address, error) {
	return _Airdrop.Contract.Token(&_Airdrop.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Airdrop *AirdropCallerSession) Token() (common.Address, error) {
	return _Airdrop.Contract.Token(&_Airdrop.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(address _addr, uint256 _amount) returns()
func (_Airdrop *AirdropTransactor) Add(opts *bind.TransactOpts, _addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "add", _addr, _amount)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(address _addr, uint256 _amount) returns()
func (_Airdrop *AirdropSession) Add(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.Add(&_Airdrop.TransactOpts, _addr, _amount)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(address _addr, uint256 _amount) returns()
func (_Airdrop *AirdropTransactorSession) Add(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.Add(&_Airdrop.TransactOpts, _addr, _amount)
}

// NotifyAirdropAmount is a paid mutator transaction binding the contract method 0x8f03981f.
//
// Solidity: function notifyAirdropAmount(address _addr) returns()
func (_Airdrop *AirdropTransactor) NotifyAirdropAmount(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "notifyAirdropAmount", _addr)
}

// NotifyAirdropAmount is a paid mutator transaction binding the contract method 0x8f03981f.
//
// Solidity: function notifyAirdropAmount(address _addr) returns()
func (_Airdrop *AirdropSession) NotifyAirdropAmount(_addr common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.NotifyAirdropAmount(&_Airdrop.TransactOpts, _addr)
}

// NotifyAirdropAmount is a paid mutator transaction binding the contract method 0x8f03981f.
//
// Solidity: function notifyAirdropAmount(address _addr) returns()
func (_Airdrop *AirdropTransactorSession) NotifyAirdropAmount(_addr common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.NotifyAirdropAmount(&_Airdrop.TransactOpts, _addr)
}

// NotifyAirdropAmounts is a paid mutator transaction binding the contract method 0x13fb7600.
//
// Solidity: function notifyAirdropAmounts() returns()
func (_Airdrop *AirdropTransactor) NotifyAirdropAmounts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "notifyAirdropAmounts")
}

// NotifyAirdropAmounts is a paid mutator transaction binding the contract method 0x13fb7600.
//
// Solidity: function notifyAirdropAmounts() returns()
func (_Airdrop *AirdropSession) NotifyAirdropAmounts() (*types.Transaction, error) {
	return _Airdrop.Contract.NotifyAirdropAmounts(&_Airdrop.TransactOpts)
}

// NotifyAirdropAmounts is a paid mutator transaction binding the contract method 0x13fb7600.
//
// Solidity: function notifyAirdropAmounts() returns()
func (_Airdrop *AirdropTransactorSession) NotifyAirdropAmounts() (*types.Transaction, error) {
	return _Airdrop.Contract.NotifyAirdropAmounts(&_Airdrop.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Airdrop *AirdropTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Airdrop *AirdropSession) RenounceOwnership() (*types.Transaction, error) {
	return _Airdrop.Contract.RenounceOwnership(&_Airdrop.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Airdrop *AirdropTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Airdrop.Contract.RenounceOwnership(&_Airdrop.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Airdrop *AirdropTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Airdrop *AirdropSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferOwnership(&_Airdrop.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Airdrop *AirdropTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferOwnership(&_Airdrop.TransactOpts, newOwner)
}

// AirdropAirdropIterator is returned from FilterAirdrop and is used to iterate over the raw logs and unpacked data for Airdrop events raised by the Airdrop contract.
type AirdropAirdropIterator struct {
	Event *AirdropAirdrop // Event containing the contract specifics and raw log

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
func (it *AirdropAirdropIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropAirdrop)
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
		it.Event = new(AirdropAirdrop)
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
func (it *AirdropAirdropIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropAirdropIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropAirdrop represents a Airdrop event raised by the Airdrop contract.
type AirdropAirdrop struct {
	Sender common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAirdrop is a free log retrieval operation binding the contract event 0x8986d2aad709d6d52ca7673e78442a1ac939dc024b80334c27ca759e7658a028.
//
// Solidity: event Airdrop(address indexed sender, address indexed to, uint256 amount)
func (_Airdrop *AirdropFilterer) FilterAirdrop(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*AirdropAirdropIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Airdrop", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AirdropAirdropIterator{contract: _Airdrop.contract, event: "Airdrop", logs: logs, sub: sub}, nil
}

// WatchAirdrop is a free log subscription operation binding the contract event 0x8986d2aad709d6d52ca7673e78442a1ac939dc024b80334c27ca759e7658a028.
//
// Solidity: event Airdrop(address indexed sender, address indexed to, uint256 amount)
func (_Airdrop *AirdropFilterer) WatchAirdrop(opts *bind.WatchOpts, sink chan<- *AirdropAirdrop, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Airdrop", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropAirdrop)
				if err := _Airdrop.contract.UnpackLog(event, "Airdrop", log); err != nil {
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

// ParseAirdrop is a log parse operation binding the contract event 0x8986d2aad709d6d52ca7673e78442a1ac939dc024b80334c27ca759e7658a028.
//
// Solidity: event Airdrop(address indexed sender, address indexed to, uint256 amount)
func (_Airdrop *AirdropFilterer) ParseAirdrop(log types.Log) (*AirdropAirdrop, error) {
	event := new(AirdropAirdrop)
	if err := _Airdrop.contract.UnpackLog(event, "Airdrop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Airdrop contract.
type AirdropOwnershipTransferredIterator struct {
	Event *AirdropOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AirdropOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropOwnershipTransferred)
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
		it.Event = new(AirdropOwnershipTransferred)
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
func (it *AirdropOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropOwnershipTransferred represents a OwnershipTransferred event raised by the Airdrop contract.
type AirdropOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Airdrop *AirdropFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AirdropOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AirdropOwnershipTransferredIterator{contract: _Airdrop.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Airdrop *AirdropFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AirdropOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropOwnershipTransferred)
				if err := _Airdrop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseOwnershipTransferred(log types.Log) (*AirdropOwnershipTransferred, error) {
	event := new(AirdropOwnershipTransferred)
	if err := _Airdrop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
