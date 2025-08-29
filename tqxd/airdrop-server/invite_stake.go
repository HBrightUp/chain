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

// InviteStakeABI is the input ABI used to generate the binding from.
const InviteStakeABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeLowerLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakeUpLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_AITDPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_invitationRewardTotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_profitTotalAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"invitee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"}],\"name\":\"AddInvite\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"des\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AirdropDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InviterReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawProfit\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AITDPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROFIT_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aa\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"}],\"name\":\"addInviter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"batch\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"airdropOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"at\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdrawAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beginIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"}],\"name\":\"getInviteeList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"}],\"name\":\"getInviteeListLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beginIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"}],\"name\":\"getInviteeListV2\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invitee\",\"type\":\"address\"}],\"name\":\"getInviter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getInviterByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getRemainderExpire\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStakedInvitation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invitationRewardAccumulate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invitationRewardTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inviteeLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inviterToInviteeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isDeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isExpire\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onlyOwnerBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"onlyOwnerWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitAccumulate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setAITDPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_airdropOwner\",\"type\":\"address\"}],\"name\":\"setAirdropOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setEmergencyWithdrawAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"}],\"name\":\"setExpire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setInvitationRewardTotalAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setProfitTotalAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lowerLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upLimit\",\"type\":\"uint256\"}],\"name\":\"setStakeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeLowerLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeUpLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userInfoSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isInvidteeReward\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invitationReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userOfPid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// InviteStake is an auto generated Go binding around an Ethereum contract.
type InviteStake struct {
	InviteStakeCaller     // Read-only binding to the contract
	InviteStakeTransactor // Write-only binding to the contract
	InviteStakeFilterer   // Log filterer for contract events
}

// InviteStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type InviteStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InviteStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InviteStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InviteStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InviteStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InviteStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InviteStakeSession struct {
	Contract     *InviteStake      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InviteStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InviteStakeCallerSession struct {
	Contract *InviteStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// InviteStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InviteStakeTransactorSession struct {
	Contract     *InviteStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InviteStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type InviteStakeRaw struct {
	Contract *InviteStake // Generic contract binding to access the raw methods on
}

// InviteStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InviteStakeCallerRaw struct {
	Contract *InviteStakeCaller // Generic read-only contract binding to access the raw methods on
}

// InviteStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InviteStakeTransactorRaw struct {
	Contract *InviteStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInviteStake creates a new instance of InviteStake, bound to a specific deployed contract.
func NewInviteStake(address common.Address, backend bind.ContractBackend) (*InviteStake, error) {
	contract, err := bindInviteStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InviteStake{InviteStakeCaller: InviteStakeCaller{contract: contract}, InviteStakeTransactor: InviteStakeTransactor{contract: contract}, InviteStakeFilterer: InviteStakeFilterer{contract: contract}}, nil
}

// NewInviteStakeCaller creates a new read-only instance of InviteStake, bound to a specific deployed contract.
func NewInviteStakeCaller(address common.Address, caller bind.ContractCaller) (*InviteStakeCaller, error) {
	contract, err := bindInviteStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InviteStakeCaller{contract: contract}, nil
}

// NewInviteStakeTransactor creates a new write-only instance of InviteStake, bound to a specific deployed contract.
func NewInviteStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*InviteStakeTransactor, error) {
	contract, err := bindInviteStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InviteStakeTransactor{contract: contract}, nil
}

// NewInviteStakeFilterer creates a new log filterer instance of InviteStake, bound to a specific deployed contract.
func NewInviteStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*InviteStakeFilterer, error) {
	contract, err := bindInviteStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InviteStakeFilterer{contract: contract}, nil
}

// bindInviteStake binds a generic wrapper to an already deployed contract.
func bindInviteStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InviteStakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InviteStake *InviteStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InviteStake.Contract.InviteStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InviteStake *InviteStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InviteStake.Contract.InviteStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InviteStake *InviteStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InviteStake.Contract.InviteStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InviteStake *InviteStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InviteStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InviteStake *InviteStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InviteStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InviteStake *InviteStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InviteStake.Contract.contract.Transact(opts, method, params...)
}

// AITDPerBlock is a free data retrieval call binding the contract method 0xc13309c7.
//
// Solidity: function AITDPerBlock() view returns(uint256)
func (_InviteStake *InviteStakeCaller) AITDPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "AITDPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AITDPerBlock is a free data retrieval call binding the contract method 0xc13309c7.
//
// Solidity: function AITDPerBlock() view returns(uint256)
func (_InviteStake *InviteStakeSession) AITDPerBlock() (*big.Int, error) {
	return _InviteStake.Contract.AITDPerBlock(&_InviteStake.CallOpts)
}

// AITDPerBlock is a free data retrieval call binding the contract method 0xc13309c7.
//
// Solidity: function AITDPerBlock() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) AITDPerBlock() (*big.Int, error) {
	return _InviteStake.Contract.AITDPerBlock(&_InviteStake.CallOpts)
}

// PROFITBASE is a free data retrieval call binding the contract method 0xe1da8f58.
//
// Solidity: function PROFIT_BASE() view returns(uint256)
func (_InviteStake *InviteStakeCaller) PROFITBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "PROFIT_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROFITBASE is a free data retrieval call binding the contract method 0xe1da8f58.
//
// Solidity: function PROFIT_BASE() view returns(uint256)
func (_InviteStake *InviteStakeSession) PROFITBASE() (*big.Int, error) {
	return _InviteStake.Contract.PROFITBASE(&_InviteStake.CallOpts)
}

// PROFITBASE is a free data retrieval call binding the contract method 0xe1da8f58.
//
// Solidity: function PROFIT_BASE() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) PROFITBASE() (*big.Int, error) {
	return _InviteStake.Contract.PROFITBASE(&_InviteStake.CallOpts)
}

// AccPerShare is a free data retrieval call binding the contract method 0xd2890a01.
//
// Solidity: function accPerShare() view returns(uint256)
func (_InviteStake *InviteStakeCaller) AccPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "accPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccPerShare is a free data retrieval call binding the contract method 0xd2890a01.
//
// Solidity: function accPerShare() view returns(uint256)
func (_InviteStake *InviteStakeSession) AccPerShare() (*big.Int, error) {
	return _InviteStake.Contract.AccPerShare(&_InviteStake.CallOpts)
}

// AccPerShare is a free data retrieval call binding the contract method 0xd2890a01.
//
// Solidity: function accPerShare() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) AccPerShare() (*big.Int, error) {
	return _InviteStake.Contract.AccPerShare(&_InviteStake.CallOpts)
}

// AirdropOwner is a free data retrieval call binding the contract method 0x8e840cc9.
//
// Solidity: function airdropOwner() view returns(address)
func (_InviteStake *InviteStakeCaller) AirdropOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "airdropOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AirdropOwner is a free data retrieval call binding the contract method 0x8e840cc9.
//
// Solidity: function airdropOwner() view returns(address)
func (_InviteStake *InviteStakeSession) AirdropOwner() (common.Address, error) {
	return _InviteStake.Contract.AirdropOwner(&_InviteStake.CallOpts)
}

// AirdropOwner is a free data retrieval call binding the contract method 0x8e840cc9.
//
// Solidity: function airdropOwner() view returns(address)
func (_InviteStake *InviteStakeCallerSession) AirdropOwner() (common.Address, error) {
	return _InviteStake.Contract.AirdropOwner(&_InviteStake.CallOpts)
}

// At is a free data retrieval call binding the contract method 0x7fd22e39.
//
// Solidity: function at(address inviter, uint256 index) view returns(address)
func (_InviteStake *InviteStakeCaller) At(opts *bind.CallOpts, inviter common.Address, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "at", inviter, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// At is a free data retrieval call binding the contract method 0x7fd22e39.
//
// Solidity: function at(address inviter, uint256 index) view returns(address)
func (_InviteStake *InviteStakeSession) At(inviter common.Address, index *big.Int) (common.Address, error) {
	return _InviteStake.Contract.At(&_InviteStake.CallOpts, inviter, index)
}

// At is a free data retrieval call binding the contract method 0x7fd22e39.
//
// Solidity: function at(address inviter, uint256 index) view returns(address)
func (_InviteStake *InviteStakeCallerSession) At(inviter common.Address, index *big.Int) (common.Address, error) {
	return _InviteStake.Contract.At(&_InviteStake.CallOpts, inviter, index)
}

// EmergencyWithdrawAddress is a free data retrieval call binding the contract method 0x2ee72e18.
//
// Solidity: function emergencyWithdrawAddress() view returns(address)
func (_InviteStake *InviteStakeCaller) EmergencyWithdrawAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "emergencyWithdrawAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyWithdrawAddress is a free data retrieval call binding the contract method 0x2ee72e18.
//
// Solidity: function emergencyWithdrawAddress() view returns(address)
func (_InviteStake *InviteStakeSession) EmergencyWithdrawAddress() (common.Address, error) {
	return _InviteStake.Contract.EmergencyWithdrawAddress(&_InviteStake.CallOpts)
}

// EmergencyWithdrawAddress is a free data retrieval call binding the contract method 0x2ee72e18.
//
// Solidity: function emergencyWithdrawAddress() view returns(address)
func (_InviteStake *InviteStakeCallerSession) EmergencyWithdrawAddress() (common.Address, error) {
	return _InviteStake.Contract.EmergencyWithdrawAddress(&_InviteStake.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_InviteStake *InviteStakeCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "expire")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_InviteStake *InviteStakeSession) Expire() (*big.Int, error) {
	return _InviteStake.Contract.Expire(&_InviteStake.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) Expire() (*big.Int, error) {
	return _InviteStake.Contract.Expire(&_InviteStake.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_InviteStake *InviteStakeCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_InviteStake *InviteStakeSession) GetBlockNumber() (*big.Int, error) {
	return _InviteStake.Contract.GetBlockNumber(&_InviteStake.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) GetBlockNumber() (*big.Int, error) {
	return _InviteStake.Contract.GetBlockNumber(&_InviteStake.CallOpts)
}

// GetInviteeList is a free data retrieval call binding the contract method 0xe18c81f1.
//
// Solidity: function getInviteeList(uint256 beginIndex, uint256 endIndex, address inviter) view returns(address[])
func (_InviteStake *InviteStakeCaller) GetInviteeList(opts *bind.CallOpts, beginIndex *big.Int, endIndex *big.Int, inviter common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "getInviteeList", beginIndex, endIndex, inviter)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetInviteeList is a free data retrieval call binding the contract method 0xe18c81f1.
//
// Solidity: function getInviteeList(uint256 beginIndex, uint256 endIndex, address inviter) view returns(address[])
func (_InviteStake *InviteStakeSession) GetInviteeList(beginIndex *big.Int, endIndex *big.Int, inviter common.Address) ([]common.Address, error) {
	return _InviteStake.Contract.GetInviteeList(&_InviteStake.CallOpts, beginIndex, endIndex, inviter)
}

// GetInviteeList is a free data retrieval call binding the contract method 0xe18c81f1.
//
// Solidity: function getInviteeList(uint256 beginIndex, uint256 endIndex, address inviter) view returns(address[])
func (_InviteStake *InviteStakeCallerSession) GetInviteeList(beginIndex *big.Int, endIndex *big.Int, inviter common.Address) ([]common.Address, error) {
	return _InviteStake.Contract.GetInviteeList(&_InviteStake.CallOpts, beginIndex, endIndex, inviter)
}

// GetInviteeListLength is a free data retrieval call binding the contract method 0x94896c5c.
//
// Solidity: function getInviteeListLength(address inviter) view returns(uint256)
func (_InviteStake *InviteStakeCaller) GetInviteeListLength(opts *bind.CallOpts, inviter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "getInviteeListLength", inviter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInviteeListLength is a free data retrieval call binding the contract method 0x94896c5c.
//
// Solidity: function getInviteeListLength(address inviter) view returns(uint256)
func (_InviteStake *InviteStakeSession) GetInviteeListLength(inviter common.Address) (*big.Int, error) {
	return _InviteStake.Contract.GetInviteeListLength(&_InviteStake.CallOpts, inviter)
}

// GetInviteeListLength is a free data retrieval call binding the contract method 0x94896c5c.
//
// Solidity: function getInviteeListLength(address inviter) view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) GetInviteeListLength(inviter common.Address) (*big.Int, error) {
	return _InviteStake.Contract.GetInviteeListLength(&_InviteStake.CallOpts, inviter)
}

// GetInviteeListV2 is a free data retrieval call binding the contract method 0xad3cf89b.
//
// Solidity: function getInviteeListV2(uint256 beginIndex, uint256 endIndex, address inviter) view returns(address[])
func (_InviteStake *InviteStakeCaller) GetInviteeListV2(opts *bind.CallOpts, beginIndex *big.Int, endIndex *big.Int, inviter common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "getInviteeListV2", beginIndex, endIndex, inviter)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetInviteeListV2 is a free data retrieval call binding the contract method 0xad3cf89b.
//
// Solidity: function getInviteeListV2(uint256 beginIndex, uint256 endIndex, address inviter) view returns(address[])
func (_InviteStake *InviteStakeSession) GetInviteeListV2(beginIndex *big.Int, endIndex *big.Int, inviter common.Address) ([]common.Address, error) {
	return _InviteStake.Contract.GetInviteeListV2(&_InviteStake.CallOpts, beginIndex, endIndex, inviter)
}

// GetInviteeListV2 is a free data retrieval call binding the contract method 0xad3cf89b.
//
// Solidity: function getInviteeListV2(uint256 beginIndex, uint256 endIndex, address inviter) view returns(address[])
func (_InviteStake *InviteStakeCallerSession) GetInviteeListV2(beginIndex *big.Int, endIndex *big.Int, inviter common.Address) ([]common.Address, error) {
	return _InviteStake.Contract.GetInviteeListV2(&_InviteStake.CallOpts, beginIndex, endIndex, inviter)
}

// GetInviter is a free data retrieval call binding the contract method 0xd216ce6f.
//
// Solidity: function getInviter(address invitee) view returns(bool, address)
func (_InviteStake *InviteStakeCaller) GetInviter(opts *bind.CallOpts, invitee common.Address) (bool, common.Address, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "getInviter", invitee)

	if err != nil {
		return *new(bool), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetInviter is a free data retrieval call binding the contract method 0xd216ce6f.
//
// Solidity: function getInviter(address invitee) view returns(bool, address)
func (_InviteStake *InviteStakeSession) GetInviter(invitee common.Address) (bool, common.Address, error) {
	return _InviteStake.Contract.GetInviter(&_InviteStake.CallOpts, invitee)
}

// GetInviter is a free data retrieval call binding the contract method 0xd216ce6f.
//
// Solidity: function getInviter(address invitee) view returns(bool, address)
func (_InviteStake *InviteStakeCallerSession) GetInviter(invitee common.Address) (bool, common.Address, error) {
	return _InviteStake.Contract.GetInviter(&_InviteStake.CallOpts, invitee)
}

// GetInviterByIndex is a free data retrieval call binding the contract method 0x2f4c3f62.
//
// Solidity: function getInviterByIndex(uint256 idx) view returns(address, address)
func (_InviteStake *InviteStakeCaller) GetInviterByIndex(opts *bind.CallOpts, idx *big.Int) (common.Address, common.Address, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "getInviterByIndex", idx)

	if err != nil {
		return *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetInviterByIndex is a free data retrieval call binding the contract method 0x2f4c3f62.
//
// Solidity: function getInviterByIndex(uint256 idx) view returns(address, address)
func (_InviteStake *InviteStakeSession) GetInviterByIndex(idx *big.Int) (common.Address, common.Address, error) {
	return _InviteStake.Contract.GetInviterByIndex(&_InviteStake.CallOpts, idx)
}

// GetInviterByIndex is a free data retrieval call binding the contract method 0x2f4c3f62.
//
// Solidity: function getInviterByIndex(uint256 idx) view returns(address, address)
func (_InviteStake *InviteStakeCallerSession) GetInviterByIndex(idx *big.Int) (common.Address, common.Address, error) {
	return _InviteStake.Contract.GetInviterByIndex(&_InviteStake.CallOpts, idx)
}

// GetProfit is a free data retrieval call binding the contract method 0xc600e1dc.
//
// Solidity: function getProfit(address addr) view returns(uint256)
func (_InviteStake *InviteStakeCaller) GetProfit(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "getProfit", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProfit is a free data retrieval call binding the contract method 0xc600e1dc.
//
// Solidity: function getProfit(address addr) view returns(uint256)
func (_InviteStake *InviteStakeSession) GetProfit(addr common.Address) (*big.Int, error) {
	return _InviteStake.Contract.GetProfit(&_InviteStake.CallOpts, addr)
}

// GetProfit is a free data retrieval call binding the contract method 0xc600e1dc.
//
// Solidity: function getProfit(address addr) view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) GetProfit(addr common.Address) (*big.Int, error) {
	return _InviteStake.Contract.GetProfit(&_InviteStake.CallOpts, addr)
}

// GetRemainderExpire is a free data retrieval call binding the contract method 0x52ca75e7.
//
// Solidity: function getRemainderExpire(address addr) view returns(bool, uint256)
func (_InviteStake *InviteStakeCaller) GetRemainderExpire(opts *bind.CallOpts, addr common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "getRemainderExpire", addr)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetRemainderExpire is a free data retrieval call binding the contract method 0x52ca75e7.
//
// Solidity: function getRemainderExpire(address addr) view returns(bool, uint256)
func (_InviteStake *InviteStakeSession) GetRemainderExpire(addr common.Address) (bool, *big.Int, error) {
	return _InviteStake.Contract.GetRemainderExpire(&_InviteStake.CallOpts, addr)
}

// GetRemainderExpire is a free data retrieval call binding the contract method 0x52ca75e7.
//
// Solidity: function getRemainderExpire(address addr) view returns(bool, uint256)
func (_InviteStake *InviteStakeCallerSession) GetRemainderExpire(addr common.Address) (bool, *big.Int, error) {
	return _InviteStake.Contract.GetRemainderExpire(&_InviteStake.CallOpts, addr)
}

// GetStakedInvitation is a free data retrieval call binding the contract method 0x46a0da8e.
//
// Solidity: function getStakedInvitation(address addr) view returns(uint256, uint256)
func (_InviteStake *InviteStakeCaller) GetStakedInvitation(opts *bind.CallOpts, addr common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "getStakedInvitation", addr)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStakedInvitation is a free data retrieval call binding the contract method 0x46a0da8e.
//
// Solidity: function getStakedInvitation(address addr) view returns(uint256, uint256)
func (_InviteStake *InviteStakeSession) GetStakedInvitation(addr common.Address) (*big.Int, *big.Int, error) {
	return _InviteStake.Contract.GetStakedInvitation(&_InviteStake.CallOpts, addr)
}

// GetStakedInvitation is a free data retrieval call binding the contract method 0x46a0da8e.
//
// Solidity: function getStakedInvitation(address addr) view returns(uint256, uint256)
func (_InviteStake *InviteStakeCallerSession) GetStakedInvitation(addr common.Address) (*big.Int, *big.Int, error) {
	return _InviteStake.Contract.GetStakedInvitation(&_InviteStake.CallOpts, addr)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address addr) view returns(bool, bool, uint256, uint256, uint256, uint256, uint256)
func (_InviteStake *InviteStakeCaller) GetUserInfo(opts *bind.CallOpts, addr common.Address) (bool, bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "getUserInfo", addr)

	if err != nil {
		return *new(bool), *new(bool), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address addr) view returns(bool, bool, uint256, uint256, uint256, uint256, uint256)
func (_InviteStake *InviteStakeSession) GetUserInfo(addr common.Address) (bool, bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _InviteStake.Contract.GetUserInfo(&_InviteStake.CallOpts, addr)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address addr) view returns(bool, bool, uint256, uint256, uint256, uint256, uint256)
func (_InviteStake *InviteStakeCallerSession) GetUserInfo(addr common.Address) (bool, bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _InviteStake.Contract.GetUserInfo(&_InviteStake.CallOpts, addr)
}

// InvitationRewardAccumulate is a free data retrieval call binding the contract method 0xadd99be8.
//
// Solidity: function invitationRewardAccumulate() view returns(uint256)
func (_InviteStake *InviteStakeCaller) InvitationRewardAccumulate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "invitationRewardAccumulate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvitationRewardAccumulate is a free data retrieval call binding the contract method 0xadd99be8.
//
// Solidity: function invitationRewardAccumulate() view returns(uint256)
func (_InviteStake *InviteStakeSession) InvitationRewardAccumulate() (*big.Int, error) {
	return _InviteStake.Contract.InvitationRewardAccumulate(&_InviteStake.CallOpts)
}

// InvitationRewardAccumulate is a free data retrieval call binding the contract method 0xadd99be8.
//
// Solidity: function invitationRewardAccumulate() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) InvitationRewardAccumulate() (*big.Int, error) {
	return _InviteStake.Contract.InvitationRewardAccumulate(&_InviteStake.CallOpts)
}

// InvitationRewardTotalAmount is a free data retrieval call binding the contract method 0x3755c71d.
//
// Solidity: function invitationRewardTotalAmount() view returns(uint256)
func (_InviteStake *InviteStakeCaller) InvitationRewardTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "invitationRewardTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvitationRewardTotalAmount is a free data retrieval call binding the contract method 0x3755c71d.
//
// Solidity: function invitationRewardTotalAmount() view returns(uint256)
func (_InviteStake *InviteStakeSession) InvitationRewardTotalAmount() (*big.Int, error) {
	return _InviteStake.Contract.InvitationRewardTotalAmount(&_InviteStake.CallOpts)
}

// InvitationRewardTotalAmount is a free data retrieval call binding the contract method 0x3755c71d.
//
// Solidity: function invitationRewardTotalAmount() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) InvitationRewardTotalAmount() (*big.Int, error) {
	return _InviteStake.Contract.InvitationRewardTotalAmount(&_InviteStake.CallOpts)
}

// InviteeLength is a free data retrieval call binding the contract method 0xc3289eab.
//
// Solidity: function inviteeLength() view returns(uint256)
func (_InviteStake *InviteStakeCaller) InviteeLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "inviteeLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteeLength is a free data retrieval call binding the contract method 0xc3289eab.
//
// Solidity: function inviteeLength() view returns(uint256)
func (_InviteStake *InviteStakeSession) InviteeLength() (*big.Int, error) {
	return _InviteStake.Contract.InviteeLength(&_InviteStake.CallOpts)
}

// InviteeLength is a free data retrieval call binding the contract method 0xc3289eab.
//
// Solidity: function inviteeLength() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) InviteeLength() (*big.Int, error) {
	return _InviteStake.Contract.InviteeLength(&_InviteStake.CallOpts)
}

// InviterToInviteeReward is a free data retrieval call binding the contract method 0x522a0744.
//
// Solidity: function inviterToInviteeReward(address ) view returns(uint256)
func (_InviteStake *InviteStakeCaller) InviterToInviteeReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "inviterToInviteeReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviterToInviteeReward is a free data retrieval call binding the contract method 0x522a0744.
//
// Solidity: function inviterToInviteeReward(address ) view returns(uint256)
func (_InviteStake *InviteStakeSession) InviterToInviteeReward(arg0 common.Address) (*big.Int, error) {
	return _InviteStake.Contract.InviterToInviteeReward(&_InviteStake.CallOpts, arg0)
}

// InviterToInviteeReward is a free data retrieval call binding the contract method 0x522a0744.
//
// Solidity: function inviterToInviteeReward(address ) view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) InviterToInviteeReward(arg0 common.Address) (*big.Int, error) {
	return _InviteStake.Contract.InviterToInviteeReward(&_InviteStake.CallOpts, arg0)
}

// IsDeposit is a free data retrieval call binding the contract method 0xf690a3a3.
//
// Solidity: function isDeposit(address addr) view returns(bool)
func (_InviteStake *InviteStakeCaller) IsDeposit(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "isDeposit", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeposit is a free data retrieval call binding the contract method 0xf690a3a3.
//
// Solidity: function isDeposit(address addr) view returns(bool)
func (_InviteStake *InviteStakeSession) IsDeposit(addr common.Address) (bool, error) {
	return _InviteStake.Contract.IsDeposit(&_InviteStake.CallOpts, addr)
}

// IsDeposit is a free data retrieval call binding the contract method 0xf690a3a3.
//
// Solidity: function isDeposit(address addr) view returns(bool)
func (_InviteStake *InviteStakeCallerSession) IsDeposit(addr common.Address) (bool, error) {
	return _InviteStake.Contract.IsDeposit(&_InviteStake.CallOpts, addr)
}

// IsExpire is a free data retrieval call binding the contract method 0x6bde0f1c.
//
// Solidity: function isExpire(address addr) view returns(bool)
func (_InviteStake *InviteStakeCaller) IsExpire(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "isExpire", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExpire is a free data retrieval call binding the contract method 0x6bde0f1c.
//
// Solidity: function isExpire(address addr) view returns(bool)
func (_InviteStake *InviteStakeSession) IsExpire(addr common.Address) (bool, error) {
	return _InviteStake.Contract.IsExpire(&_InviteStake.CallOpts, addr)
}

// IsExpire is a free data retrieval call binding the contract method 0x6bde0f1c.
//
// Solidity: function isExpire(address addr) view returns(bool)
func (_InviteStake *InviteStakeCallerSession) IsExpire(addr common.Address) (bool, error) {
	return _InviteStake.Contract.IsExpire(&_InviteStake.CallOpts, addr)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_InviteStake *InviteStakeCaller) LastRewardBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "lastRewardBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_InviteStake *InviteStakeSession) LastRewardBlock() (*big.Int, error) {
	return _InviteStake.Contract.LastRewardBlock(&_InviteStake.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) LastRewardBlock() (*big.Int, error) {
	return _InviteStake.Contract.LastRewardBlock(&_InviteStake.CallOpts)
}

// OnlyOwnerBalance is a free data retrieval call binding the contract method 0x62479be7.
//
// Solidity: function onlyOwnerBalance() view returns(uint256)
func (_InviteStake *InviteStakeCaller) OnlyOwnerBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "onlyOwnerBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OnlyOwnerBalance is a free data retrieval call binding the contract method 0x62479be7.
//
// Solidity: function onlyOwnerBalance() view returns(uint256)
func (_InviteStake *InviteStakeSession) OnlyOwnerBalance() (*big.Int, error) {
	return _InviteStake.Contract.OnlyOwnerBalance(&_InviteStake.CallOpts)
}

// OnlyOwnerBalance is a free data retrieval call binding the contract method 0x62479be7.
//
// Solidity: function onlyOwnerBalance() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) OnlyOwnerBalance() (*big.Int, error) {
	return _InviteStake.Contract.OnlyOwnerBalance(&_InviteStake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InviteStake *InviteStakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InviteStake *InviteStakeSession) Owner() (common.Address, error) {
	return _InviteStake.Contract.Owner(&_InviteStake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InviteStake *InviteStakeCallerSession) Owner() (common.Address, error) {
	return _InviteStake.Contract.Owner(&_InviteStake.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InviteStake *InviteStakeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InviteStake *InviteStakeSession) Paused() (bool, error) {
	return _InviteStake.Contract.Paused(&_InviteStake.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InviteStake *InviteStakeCallerSession) Paused() (bool, error) {
	return _InviteStake.Contract.Paused(&_InviteStake.CallOpts)
}

// ProfitAccumulate is a free data retrieval call binding the contract method 0x993c3dd4.
//
// Solidity: function profitAccumulate() view returns(uint256)
func (_InviteStake *InviteStakeCaller) ProfitAccumulate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "profitAccumulate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitAccumulate is a free data retrieval call binding the contract method 0x993c3dd4.
//
// Solidity: function profitAccumulate() view returns(uint256)
func (_InviteStake *InviteStakeSession) ProfitAccumulate() (*big.Int, error) {
	return _InviteStake.Contract.ProfitAccumulate(&_InviteStake.CallOpts)
}

// ProfitAccumulate is a free data retrieval call binding the contract method 0x993c3dd4.
//
// Solidity: function profitAccumulate() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) ProfitAccumulate() (*big.Int, error) {
	return _InviteStake.Contract.ProfitAccumulate(&_InviteStake.CallOpts)
}

// ProfitTotalAmount is a free data retrieval call binding the contract method 0xe270c9b2.
//
// Solidity: function profitTotalAmount() view returns(uint256)
func (_InviteStake *InviteStakeCaller) ProfitTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "profitTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitTotalAmount is a free data retrieval call binding the contract method 0xe270c9b2.
//
// Solidity: function profitTotalAmount() view returns(uint256)
func (_InviteStake *InviteStakeSession) ProfitTotalAmount() (*big.Int, error) {
	return _InviteStake.Contract.ProfitTotalAmount(&_InviteStake.CallOpts)
}

// ProfitTotalAmount is a free data retrieval call binding the contract method 0xe270c9b2.
//
// Solidity: function profitTotalAmount() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) ProfitTotalAmount() (*big.Int, error) {
	return _InviteStake.Contract.ProfitTotalAmount(&_InviteStake.CallOpts)
}

// StakeLowerLimit is a free data retrieval call binding the contract method 0xc3cee9f3.
//
// Solidity: function stakeLowerLimit() view returns(uint256)
func (_InviteStake *InviteStakeCaller) StakeLowerLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "stakeLowerLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeLowerLimit is a free data retrieval call binding the contract method 0xc3cee9f3.
//
// Solidity: function stakeLowerLimit() view returns(uint256)
func (_InviteStake *InviteStakeSession) StakeLowerLimit() (*big.Int, error) {
	return _InviteStake.Contract.StakeLowerLimit(&_InviteStake.CallOpts)
}

// StakeLowerLimit is a free data retrieval call binding the contract method 0xc3cee9f3.
//
// Solidity: function stakeLowerLimit() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) StakeLowerLimit() (*big.Int, error) {
	return _InviteStake.Contract.StakeLowerLimit(&_InviteStake.CallOpts)
}

// StakeUpLimit is a free data retrieval call binding the contract method 0x78e2dba6.
//
// Solidity: function stakeUpLimit() view returns(uint256)
func (_InviteStake *InviteStakeCaller) StakeUpLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "stakeUpLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeUpLimit is a free data retrieval call binding the contract method 0x78e2dba6.
//
// Solidity: function stakeUpLimit() view returns(uint256)
func (_InviteStake *InviteStakeSession) StakeUpLimit() (*big.Int, error) {
	return _InviteStake.Contract.StakeUpLimit(&_InviteStake.CallOpts)
}

// StakeUpLimit is a free data retrieval call binding the contract method 0x78e2dba6.
//
// Solidity: function stakeUpLimit() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) StakeUpLimit() (*big.Int, error) {
	return _InviteStake.Contract.StakeUpLimit(&_InviteStake.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_InviteStake *InviteStakeCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_InviteStake *InviteStakeSession) TotalAmount() (*big.Int, error) {
	return _InviteStake.Contract.TotalAmount(&_InviteStake.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) TotalAmount() (*big.Int, error) {
	return _InviteStake.Contract.TotalAmount(&_InviteStake.CallOpts)
}

// UserInfoSet is a free data retrieval call binding the contract method 0x7399a8a1.
//
// Solidity: function userInfoSet(uint256 ) view returns(bool isInvidteeReward, uint256 stakedAmount, uint256 invitationReward, uint256 firstStakeAmount, uint256 lastBlock, uint256 profitDebt)
func (_InviteStake *InviteStakeCaller) UserInfoSet(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsInvidteeReward bool
	StakedAmount     *big.Int
	InvitationReward *big.Int
	FirstStakeAmount *big.Int
	LastBlock        *big.Int
	ProfitDebt       *big.Int
}, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "userInfoSet", arg0)

	outstruct := new(struct {
		IsInvidteeReward bool
		StakedAmount     *big.Int
		InvitationReward *big.Int
		FirstStakeAmount *big.Int
		LastBlock        *big.Int
		ProfitDebt       *big.Int
	})

	outstruct.IsInvidteeReward = out[0].(bool)
	outstruct.StakedAmount = out[1].(*big.Int)
	outstruct.InvitationReward = out[2].(*big.Int)
	outstruct.FirstStakeAmount = out[3].(*big.Int)
	outstruct.LastBlock = out[4].(*big.Int)
	outstruct.ProfitDebt = out[5].(*big.Int)

	return *outstruct, err

}

// UserInfoSet is a free data retrieval call binding the contract method 0x7399a8a1.
//
// Solidity: function userInfoSet(uint256 ) view returns(bool isInvidteeReward, uint256 stakedAmount, uint256 invitationReward, uint256 firstStakeAmount, uint256 lastBlock, uint256 profitDebt)
func (_InviteStake *InviteStakeSession) UserInfoSet(arg0 *big.Int) (struct {
	IsInvidteeReward bool
	StakedAmount     *big.Int
	InvitationReward *big.Int
	FirstStakeAmount *big.Int
	LastBlock        *big.Int
	ProfitDebt       *big.Int
}, error) {
	return _InviteStake.Contract.UserInfoSet(&_InviteStake.CallOpts, arg0)
}

// UserInfoSet is a free data retrieval call binding the contract method 0x7399a8a1.
//
// Solidity: function userInfoSet(uint256 ) view returns(bool isInvidteeReward, uint256 stakedAmount, uint256 invitationReward, uint256 firstStakeAmount, uint256 lastBlock, uint256 profitDebt)
func (_InviteStake *InviteStakeCallerSession) UserInfoSet(arg0 *big.Int) (struct {
	IsInvidteeReward bool
	StakedAmount     *big.Int
	InvitationReward *big.Int
	FirstStakeAmount *big.Int
	LastBlock        *big.Int
	ProfitDebt       *big.Int
}, error) {
	return _InviteStake.Contract.UserInfoSet(&_InviteStake.CallOpts, arg0)
}

// UserLength is a free data retrieval call binding the contract method 0x256da24f.
//
// Solidity: function userLength() view returns(uint256)
func (_InviteStake *InviteStakeCaller) UserLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "userLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLength is a free data retrieval call binding the contract method 0x256da24f.
//
// Solidity: function userLength() view returns(uint256)
func (_InviteStake *InviteStakeSession) UserLength() (*big.Int, error) {
	return _InviteStake.Contract.UserLength(&_InviteStake.CallOpts)
}

// UserLength is a free data retrieval call binding the contract method 0x256da24f.
//
// Solidity: function userLength() view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) UserLength() (*big.Int, error) {
	return _InviteStake.Contract.UserLength(&_InviteStake.CallOpts)
}

// UserOfPid is a free data retrieval call binding the contract method 0x48c3a2d8.
//
// Solidity: function userOfPid(address ) view returns(uint256)
func (_InviteStake *InviteStakeCaller) UserOfPid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InviteStake.contract.Call(opts, &out, "userOfPid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserOfPid is a free data retrieval call binding the contract method 0x48c3a2d8.
//
// Solidity: function userOfPid(address ) view returns(uint256)
func (_InviteStake *InviteStakeSession) UserOfPid(arg0 common.Address) (*big.Int, error) {
	return _InviteStake.Contract.UserOfPid(&_InviteStake.CallOpts, arg0)
}

// UserOfPid is a free data retrieval call binding the contract method 0x48c3a2d8.
//
// Solidity: function userOfPid(address ) view returns(uint256)
func (_InviteStake *InviteStakeCallerSession) UserOfPid(arg0 common.Address) (*big.Int, error) {
	return _InviteStake.Contract.UserOfPid(&_InviteStake.CallOpts, arg0)
}

// Aa is a paid mutator transaction binding the contract method 0x8466c3e6.
//
// Solidity: function aa() payable returns()
func (_InviteStake *InviteStakeTransactor) Aa(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "aa")
}

// Aa is a paid mutator transaction binding the contract method 0x8466c3e6.
//
// Solidity: function aa() payable returns()
func (_InviteStake *InviteStakeSession) Aa() (*types.Transaction, error) {
	return _InviteStake.Contract.Aa(&_InviteStake.TransactOpts)
}

// Aa is a paid mutator transaction binding the contract method 0x8466c3e6.
//
// Solidity: function aa() payable returns()
func (_InviteStake *InviteStakeTransactorSession) Aa() (*types.Transaction, error) {
	return _InviteStake.Contract.Aa(&_InviteStake.TransactOpts)
}

// AddInviter is a paid mutator transaction binding the contract method 0x61a9ab35.
//
// Solidity: function addInviter(address inviter) returns()
func (_InviteStake *InviteStakeTransactor) AddInviter(opts *bind.TransactOpts, inviter common.Address) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "addInviter", inviter)
}

// AddInviter is a paid mutator transaction binding the contract method 0x61a9ab35.
//
// Solidity: function addInviter(address inviter) returns()
func (_InviteStake *InviteStakeSession) AddInviter(inviter common.Address) (*types.Transaction, error) {
	return _InviteStake.Contract.AddInviter(&_InviteStake.TransactOpts, inviter)
}

// AddInviter is a paid mutator transaction binding the contract method 0x61a9ab35.
//
// Solidity: function addInviter(address inviter) returns()
func (_InviteStake *InviteStakeTransactorSession) AddInviter(inviter common.Address) (*types.Transaction, error) {
	return _InviteStake.Contract.AddInviter(&_InviteStake.TransactOpts, inviter)
}

// Airdrop is a paid mutator transaction binding the contract method 0xc204642c.
//
// Solidity: function airdrop(address[] batch, uint256 amount) returns()
func (_InviteStake *InviteStakeTransactor) Airdrop(opts *bind.TransactOpts, batch []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "airdrop", batch, amount)
}

// Airdrop is a paid mutator transaction binding the contract method 0xc204642c.
//
// Solidity: function airdrop(address[] batch, uint256 amount) returns()
func (_InviteStake *InviteStakeSession) Airdrop(batch []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.Airdrop(&_InviteStake.TransactOpts, batch, amount)
}

// Airdrop is a paid mutator transaction binding the contract method 0xc204642c.
//
// Solidity: function airdrop(address[] batch, uint256 amount) returns()
func (_InviteStake *InviteStakeTransactorSession) Airdrop(batch []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.Airdrop(&_InviteStake.TransactOpts, batch, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_InviteStake *InviteStakeTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_InviteStake *InviteStakeSession) Deposit() (*types.Transaction, error) {
	return _InviteStake.Contract.Deposit(&_InviteStake.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_InviteStake *InviteStakeTransactorSession) Deposit() (*types.Transaction, error) {
	return _InviteStake.Contract.Deposit(&_InviteStake.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address addr) returns()
func (_InviteStake *InviteStakeTransactor) EmergencyWithdraw(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "emergencyWithdraw", addr)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address addr) returns()
func (_InviteStake *InviteStakeSession) EmergencyWithdraw(addr common.Address) (*types.Transaction, error) {
	return _InviteStake.Contract.EmergencyWithdraw(&_InviteStake.TransactOpts, addr)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address addr) returns()
func (_InviteStake *InviteStakeTransactorSession) EmergencyWithdraw(addr common.Address) (*types.Transaction, error) {
	return _InviteStake.Contract.EmergencyWithdraw(&_InviteStake.TransactOpts, addr)
}

// OnlyOwnerWithdraw is a paid mutator transaction binding the contract method 0x96a37aee.
//
// Solidity: function onlyOwnerWithdraw(uint256 _amount) returns()
func (_InviteStake *InviteStakeTransactor) OnlyOwnerWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "onlyOwnerWithdraw", _amount)
}

// OnlyOwnerWithdraw is a paid mutator transaction binding the contract method 0x96a37aee.
//
// Solidity: function onlyOwnerWithdraw(uint256 _amount) returns()
func (_InviteStake *InviteStakeSession) OnlyOwnerWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.OnlyOwnerWithdraw(&_InviteStake.TransactOpts, _amount)
}

// OnlyOwnerWithdraw is a paid mutator transaction binding the contract method 0x96a37aee.
//
// Solidity: function onlyOwnerWithdraw(uint256 _amount) returns()
func (_InviteStake *InviteStakeTransactorSession) OnlyOwnerWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.OnlyOwnerWithdraw(&_InviteStake.TransactOpts, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InviteStake *InviteStakeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InviteStake *InviteStakeSession) RenounceOwnership() (*types.Transaction, error) {
	return _InviteStake.Contract.RenounceOwnership(&_InviteStake.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InviteStake *InviteStakeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InviteStake.Contract.RenounceOwnership(&_InviteStake.TransactOpts)
}

// SetAITDPerBlock is a paid mutator transaction binding the contract method 0xb230d698.
//
// Solidity: function setAITDPerBlock(uint256 amount) returns()
func (_InviteStake *InviteStakeTransactor) SetAITDPerBlock(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "setAITDPerBlock", amount)
}

// SetAITDPerBlock is a paid mutator transaction binding the contract method 0xb230d698.
//
// Solidity: function setAITDPerBlock(uint256 amount) returns()
func (_InviteStake *InviteStakeSession) SetAITDPerBlock(amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.SetAITDPerBlock(&_InviteStake.TransactOpts, amount)
}

// SetAITDPerBlock is a paid mutator transaction binding the contract method 0xb230d698.
//
// Solidity: function setAITDPerBlock(uint256 amount) returns()
func (_InviteStake *InviteStakeTransactorSession) SetAITDPerBlock(amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.SetAITDPerBlock(&_InviteStake.TransactOpts, amount)
}

// SetAirdropOwner is a paid mutator transaction binding the contract method 0x0f6fe11d.
//
// Solidity: function setAirdropOwner(address _airdropOwner) returns()
func (_InviteStake *InviteStakeTransactor) SetAirdropOwner(opts *bind.TransactOpts, _airdropOwner common.Address) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "setAirdropOwner", _airdropOwner)
}

// SetAirdropOwner is a paid mutator transaction binding the contract method 0x0f6fe11d.
//
// Solidity: function setAirdropOwner(address _airdropOwner) returns()
func (_InviteStake *InviteStakeSession) SetAirdropOwner(_airdropOwner common.Address) (*types.Transaction, error) {
	return _InviteStake.Contract.SetAirdropOwner(&_InviteStake.TransactOpts, _airdropOwner)
}

// SetAirdropOwner is a paid mutator transaction binding the contract method 0x0f6fe11d.
//
// Solidity: function setAirdropOwner(address _airdropOwner) returns()
func (_InviteStake *InviteStakeTransactorSession) SetAirdropOwner(_airdropOwner common.Address) (*types.Transaction, error) {
	return _InviteStake.Contract.SetAirdropOwner(&_InviteStake.TransactOpts, _airdropOwner)
}

// SetEmergencyWithdrawAddress is a paid mutator transaction binding the contract method 0x47c85634.
//
// Solidity: function setEmergencyWithdrawAddress(address addr) returns()
func (_InviteStake *InviteStakeTransactor) SetEmergencyWithdrawAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "setEmergencyWithdrawAddress", addr)
}

// SetEmergencyWithdrawAddress is a paid mutator transaction binding the contract method 0x47c85634.
//
// Solidity: function setEmergencyWithdrawAddress(address addr) returns()
func (_InviteStake *InviteStakeSession) SetEmergencyWithdrawAddress(addr common.Address) (*types.Transaction, error) {
	return _InviteStake.Contract.SetEmergencyWithdrawAddress(&_InviteStake.TransactOpts, addr)
}

// SetEmergencyWithdrawAddress is a paid mutator transaction binding the contract method 0x47c85634.
//
// Solidity: function setEmergencyWithdrawAddress(address addr) returns()
func (_InviteStake *InviteStakeTransactorSession) SetEmergencyWithdrawAddress(addr common.Address) (*types.Transaction, error) {
	return _InviteStake.Contract.SetEmergencyWithdrawAddress(&_InviteStake.TransactOpts, addr)
}

// SetExpire is a paid mutator transaction binding the contract method 0x32c27052.
//
// Solidity: function setExpire(uint256 exp) returns()
func (_InviteStake *InviteStakeTransactor) SetExpire(opts *bind.TransactOpts, exp *big.Int) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "setExpire", exp)
}

// SetExpire is a paid mutator transaction binding the contract method 0x32c27052.
//
// Solidity: function setExpire(uint256 exp) returns()
func (_InviteStake *InviteStakeSession) SetExpire(exp *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.SetExpire(&_InviteStake.TransactOpts, exp)
}

// SetExpire is a paid mutator transaction binding the contract method 0x32c27052.
//
// Solidity: function setExpire(uint256 exp) returns()
func (_InviteStake *InviteStakeTransactorSession) SetExpire(exp *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.SetExpire(&_InviteStake.TransactOpts, exp)
}

// SetInvitationRewardTotalAmount is a paid mutator transaction binding the contract method 0xcb7b2ecb.
//
// Solidity: function setInvitationRewardTotalAmount(uint256 amount) returns()
func (_InviteStake *InviteStakeTransactor) SetInvitationRewardTotalAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "setInvitationRewardTotalAmount", amount)
}

// SetInvitationRewardTotalAmount is a paid mutator transaction binding the contract method 0xcb7b2ecb.
//
// Solidity: function setInvitationRewardTotalAmount(uint256 amount) returns()
func (_InviteStake *InviteStakeSession) SetInvitationRewardTotalAmount(amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.SetInvitationRewardTotalAmount(&_InviteStake.TransactOpts, amount)
}

// SetInvitationRewardTotalAmount is a paid mutator transaction binding the contract method 0xcb7b2ecb.
//
// Solidity: function setInvitationRewardTotalAmount(uint256 amount) returns()
func (_InviteStake *InviteStakeTransactorSession) SetInvitationRewardTotalAmount(amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.SetInvitationRewardTotalAmount(&_InviteStake.TransactOpts, amount)
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_InviteStake *InviteStakeTransactor) SetPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "setPause")
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_InviteStake *InviteStakeSession) SetPause() (*types.Transaction, error) {
	return _InviteStake.Contract.SetPause(&_InviteStake.TransactOpts)
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_InviteStake *InviteStakeTransactorSession) SetPause() (*types.Transaction, error) {
	return _InviteStake.Contract.SetPause(&_InviteStake.TransactOpts)
}

// SetProfitTotalAmount is a paid mutator transaction binding the contract method 0xa1b19eda.
//
// Solidity: function setProfitTotalAmount(uint256 amount) returns()
func (_InviteStake *InviteStakeTransactor) SetProfitTotalAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "setProfitTotalAmount", amount)
}

// SetProfitTotalAmount is a paid mutator transaction binding the contract method 0xa1b19eda.
//
// Solidity: function setProfitTotalAmount(uint256 amount) returns()
func (_InviteStake *InviteStakeSession) SetProfitTotalAmount(amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.SetProfitTotalAmount(&_InviteStake.TransactOpts, amount)
}

// SetProfitTotalAmount is a paid mutator transaction binding the contract method 0xa1b19eda.
//
// Solidity: function setProfitTotalAmount(uint256 amount) returns()
func (_InviteStake *InviteStakeTransactorSession) SetProfitTotalAmount(amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.SetProfitTotalAmount(&_InviteStake.TransactOpts, amount)
}

// SetStakeLimit is a paid mutator transaction binding the contract method 0x60f62506.
//
// Solidity: function setStakeLimit(uint256 lowerLimit, uint256 upLimit) returns()
func (_InviteStake *InviteStakeTransactor) SetStakeLimit(opts *bind.TransactOpts, lowerLimit *big.Int, upLimit *big.Int) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "setStakeLimit", lowerLimit, upLimit)
}

// SetStakeLimit is a paid mutator transaction binding the contract method 0x60f62506.
//
// Solidity: function setStakeLimit(uint256 lowerLimit, uint256 upLimit) returns()
func (_InviteStake *InviteStakeSession) SetStakeLimit(lowerLimit *big.Int, upLimit *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.SetStakeLimit(&_InviteStake.TransactOpts, lowerLimit, upLimit)
}

// SetStakeLimit is a paid mutator transaction binding the contract method 0x60f62506.
//
// Solidity: function setStakeLimit(uint256 lowerLimit, uint256 upLimit) returns()
func (_InviteStake *InviteStakeTransactorSession) SetStakeLimit(lowerLimit *big.Int, upLimit *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.SetStakeLimit(&_InviteStake.TransactOpts, lowerLimit, upLimit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InviteStake *InviteStakeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InviteStake *InviteStakeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InviteStake.Contract.TransferOwnership(&_InviteStake.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InviteStake *InviteStakeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InviteStake.Contract.TransferOwnership(&_InviteStake.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_InviteStake *InviteStakeTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_InviteStake *InviteStakeSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.Withdraw(&_InviteStake.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_InviteStake *InviteStakeTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _InviteStake.Contract.Withdraw(&_InviteStake.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_InviteStake *InviteStakeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InviteStake.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_InviteStake *InviteStakeSession) Receive() (*types.Transaction, error) {
	return _InviteStake.Contract.Receive(&_InviteStake.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_InviteStake *InviteStakeTransactorSession) Receive() (*types.Transaction, error) {
	return _InviteStake.Contract.Receive(&_InviteStake.TransactOpts)
}

// InviteStakeAddInviteIterator is returned from FilterAddInvite and is used to iterate over the raw logs and unpacked data for AddInvite events raised by the InviteStake contract.
type InviteStakeAddInviteIterator struct {
	Event *InviteStakeAddInvite // Event containing the contract specifics and raw log

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
func (it *InviteStakeAddInviteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InviteStakeAddInvite)
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
		it.Event = new(InviteStakeAddInvite)
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
func (it *InviteStakeAddInviteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InviteStakeAddInviteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InviteStakeAddInvite represents a AddInvite event raised by the InviteStake contract.
type InviteStakeAddInvite struct {
	Invitee common.Address
	Inviter common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddInvite is a free log retrieval operation binding the contract event 0x9e67a0ec54544bb9be560014648e3c4909fc6365a76f11e8549f36e3638dbeae.
//
// Solidity: event AddInvite(address indexed invitee, address indexed inviter)
func (_InviteStake *InviteStakeFilterer) FilterAddInvite(opts *bind.FilterOpts, invitee []common.Address, inviter []common.Address) (*InviteStakeAddInviteIterator, error) {

	var inviteeRule []interface{}
	for _, inviteeItem := range invitee {
		inviteeRule = append(inviteeRule, inviteeItem)
	}
	var inviterRule []interface{}
	for _, inviterItem := range inviter {
		inviterRule = append(inviterRule, inviterItem)
	}

	logs, sub, err := _InviteStake.contract.FilterLogs(opts, "AddInvite", inviteeRule, inviterRule)
	if err != nil {
		return nil, err
	}
	return &InviteStakeAddInviteIterator{contract: _InviteStake.contract, event: "AddInvite", logs: logs, sub: sub}, nil
}

// WatchAddInvite is a free log subscription operation binding the contract event 0x9e67a0ec54544bb9be560014648e3c4909fc6365a76f11e8549f36e3638dbeae.
//
// Solidity: event AddInvite(address indexed invitee, address indexed inviter)
func (_InviteStake *InviteStakeFilterer) WatchAddInvite(opts *bind.WatchOpts, sink chan<- *InviteStakeAddInvite, invitee []common.Address, inviter []common.Address) (event.Subscription, error) {

	var inviteeRule []interface{}
	for _, inviteeItem := range invitee {
		inviteeRule = append(inviteeRule, inviteeItem)
	}
	var inviterRule []interface{}
	for _, inviterItem := range inviter {
		inviterRule = append(inviterRule, inviterItem)
	}

	logs, sub, err := _InviteStake.contract.WatchLogs(opts, "AddInvite", inviteeRule, inviterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InviteStakeAddInvite)
				if err := _InviteStake.contract.UnpackLog(event, "AddInvite", log); err != nil {
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

// ParseAddInvite is a log parse operation binding the contract event 0x9e67a0ec54544bb9be560014648e3c4909fc6365a76f11e8549f36e3638dbeae.
//
// Solidity: event AddInvite(address indexed invitee, address indexed inviter)
func (_InviteStake *InviteStakeFilterer) ParseAddInvite(log types.Log) (*InviteStakeAddInvite, error) {
	event := new(InviteStakeAddInvite)
	if err := _InviteStake.contract.UnpackLog(event, "AddInvite", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InviteStakeAirdropDepositIterator is returned from FilterAirdropDeposit and is used to iterate over the raw logs and unpacked data for AirdropDeposit events raised by the InviteStake contract.
type InviteStakeAirdropDepositIterator struct {
	Event *InviteStakeAirdropDeposit // Event containing the contract specifics and raw log

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
func (it *InviteStakeAirdropDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InviteStakeAirdropDeposit)
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
		it.Event = new(InviteStakeAirdropDeposit)
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
func (it *InviteStakeAirdropDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InviteStakeAirdropDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InviteStakeAirdropDeposit represents a AirdropDeposit event raised by the InviteStake contract.
type InviteStakeAirdropDeposit struct {
	Src    common.Address
	Des    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAirdropDeposit is a free log retrieval operation binding the contract event 0x39a7d1d43def1f7fcd2dbfbbf5d2ffcfdcffdece3776dad1ac48a1c9652ae936.
//
// Solidity: event AirdropDeposit(address indexed src, address indexed des, uint256 amount)
func (_InviteStake *InviteStakeFilterer) FilterAirdropDeposit(opts *bind.FilterOpts, src []common.Address, des []common.Address) (*InviteStakeAirdropDepositIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var desRule []interface{}
	for _, desItem := range des {
		desRule = append(desRule, desItem)
	}

	logs, sub, err := _InviteStake.contract.FilterLogs(opts, "AirdropDeposit", srcRule, desRule)
	if err != nil {
		return nil, err
	}
	return &InviteStakeAirdropDepositIterator{contract: _InviteStake.contract, event: "AirdropDeposit", logs: logs, sub: sub}, nil
}

// WatchAirdropDeposit is a free log subscription operation binding the contract event 0x39a7d1d43def1f7fcd2dbfbbf5d2ffcfdcffdece3776dad1ac48a1c9652ae936.
//
// Solidity: event AirdropDeposit(address indexed src, address indexed des, uint256 amount)
func (_InviteStake *InviteStakeFilterer) WatchAirdropDeposit(opts *bind.WatchOpts, sink chan<- *InviteStakeAirdropDeposit, src []common.Address, des []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var desRule []interface{}
	for _, desItem := range des {
		desRule = append(desRule, desItem)
	}

	logs, sub, err := _InviteStake.contract.WatchLogs(opts, "AirdropDeposit", srcRule, desRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InviteStakeAirdropDeposit)
				if err := _InviteStake.contract.UnpackLog(event, "AirdropDeposit", log); err != nil {
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

// ParseAirdropDeposit is a log parse operation binding the contract event 0x39a7d1d43def1f7fcd2dbfbbf5d2ffcfdcffdece3776dad1ac48a1c9652ae936.
//
// Solidity: event AirdropDeposit(address indexed src, address indexed des, uint256 amount)
func (_InviteStake *InviteStakeFilterer) ParseAirdropDeposit(log types.Log) (*InviteStakeAirdropDeposit, error) {
	event := new(InviteStakeAirdropDeposit)
	if err := _InviteStake.contract.UnpackLog(event, "AirdropDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InviteStakeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the InviteStake contract.
type InviteStakeDepositIterator struct {
	Event *InviteStakeDeposit // Event containing the contract specifics and raw log

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
func (it *InviteStakeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InviteStakeDeposit)
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
		it.Event = new(InviteStakeDeposit)
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
func (it *InviteStakeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InviteStakeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InviteStakeDeposit represents a Deposit event raised by the InviteStake contract.
type InviteStakeDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*InviteStakeDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InviteStake.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &InviteStakeDepositIterator{contract: _InviteStake.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *InviteStakeDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InviteStake.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InviteStakeDeposit)
				if err := _InviteStake.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) ParseDeposit(log types.Log) (*InviteStakeDeposit, error) {
	event := new(InviteStakeDeposit)
	if err := _InviteStake.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InviteStakeEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the InviteStake contract.
type InviteStakeEmergencyWithdrawIterator struct {
	Event *InviteStakeEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *InviteStakeEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InviteStakeEmergencyWithdraw)
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
		it.Event = new(InviteStakeEmergencyWithdraw)
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
func (it *InviteStakeEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InviteStakeEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InviteStakeEmergencyWithdraw represents a EmergencyWithdraw event raised by the InviteStake contract.
type InviteStakeEmergencyWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address) (*InviteStakeEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InviteStake.contract.FilterLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &InviteStakeEmergencyWithdrawIterator{contract: _InviteStake.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *InviteStakeEmergencyWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InviteStake.contract.WatchLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InviteStakeEmergencyWithdraw)
				if err := _InviteStake.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) ParseEmergencyWithdraw(log types.Log) (*InviteStakeEmergencyWithdraw, error) {
	event := new(InviteStakeEmergencyWithdraw)
	if err := _InviteStake.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InviteStakeInviterRewardIterator is returned from FilterInviterReward and is used to iterate over the raw logs and unpacked data for InviterReward events raised by the InviteStake contract.
type InviteStakeInviterRewardIterator struct {
	Event *InviteStakeInviterReward // Event containing the contract specifics and raw log

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
func (it *InviteStakeInviterRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InviteStakeInviterReward)
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
		it.Event = new(InviteStakeInviterReward)
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
func (it *InviteStakeInviterRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InviteStakeInviterRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InviteStakeInviterReward represents a InviterReward event raised by the InviteStake contract.
type InviteStakeInviterReward struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInviterReward is a free log retrieval operation binding the contract event 0x389563883f983db1cb1b00fca2f5c30e56ce545ee60ed77f8e07130f04c53056.
//
// Solidity: event InviterReward(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) FilterInviterReward(opts *bind.FilterOpts, user []common.Address) (*InviteStakeInviterRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InviteStake.contract.FilterLogs(opts, "InviterReward", userRule)
	if err != nil {
		return nil, err
	}
	return &InviteStakeInviterRewardIterator{contract: _InviteStake.contract, event: "InviterReward", logs: logs, sub: sub}, nil
}

// WatchInviterReward is a free log subscription operation binding the contract event 0x389563883f983db1cb1b00fca2f5c30e56ce545ee60ed77f8e07130f04c53056.
//
// Solidity: event InviterReward(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) WatchInviterReward(opts *bind.WatchOpts, sink chan<- *InviteStakeInviterReward, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InviteStake.contract.WatchLogs(opts, "InviterReward", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InviteStakeInviterReward)
				if err := _InviteStake.contract.UnpackLog(event, "InviterReward", log); err != nil {
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

// ParseInviterReward is a log parse operation binding the contract event 0x389563883f983db1cb1b00fca2f5c30e56ce545ee60ed77f8e07130f04c53056.
//
// Solidity: event InviterReward(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) ParseInviterReward(log types.Log) (*InviteStakeInviterReward, error) {
	event := new(InviteStakeInviterReward)
	if err := _InviteStake.contract.UnpackLog(event, "InviterReward", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InviteStakeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InviteStake contract.
type InviteStakeOwnershipTransferredIterator struct {
	Event *InviteStakeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InviteStakeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InviteStakeOwnershipTransferred)
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
		it.Event = new(InviteStakeOwnershipTransferred)
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
func (it *InviteStakeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InviteStakeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InviteStakeOwnershipTransferred represents a OwnershipTransferred event raised by the InviteStake contract.
type InviteStakeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InviteStake *InviteStakeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InviteStakeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InviteStake.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InviteStakeOwnershipTransferredIterator{contract: _InviteStake.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InviteStake *InviteStakeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InviteStakeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InviteStake.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InviteStakeOwnershipTransferred)
				if err := _InviteStake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InviteStake *InviteStakeFilterer) ParseOwnershipTransferred(log types.Log) (*InviteStakeOwnershipTransferred, error) {
	event := new(InviteStakeOwnershipTransferred)
	if err := _InviteStake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InviteStakeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the InviteStake contract.
type InviteStakeWithdrawIterator struct {
	Event *InviteStakeWithdraw // Event containing the contract specifics and raw log

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
func (it *InviteStakeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InviteStakeWithdraw)
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
		it.Event = new(InviteStakeWithdraw)
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
func (it *InviteStakeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InviteStakeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InviteStakeWithdraw represents a Withdraw event raised by the InviteStake contract.
type InviteStakeWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*InviteStakeWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InviteStake.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &InviteStakeWithdrawIterator{contract: _InviteStake.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *InviteStakeWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InviteStake.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InviteStakeWithdraw)
				if err := _InviteStake.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) ParseWithdraw(log types.Log) (*InviteStakeWithdraw, error) {
	event := new(InviteStakeWithdraw)
	if err := _InviteStake.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InviteStakeWithdrawProfitIterator is returned from FilterWithdrawProfit and is used to iterate over the raw logs and unpacked data for WithdrawProfit events raised by the InviteStake contract.
type InviteStakeWithdrawProfitIterator struct {
	Event *InviteStakeWithdrawProfit // Event containing the contract specifics and raw log

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
func (it *InviteStakeWithdrawProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InviteStakeWithdrawProfit)
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
		it.Event = new(InviteStakeWithdrawProfit)
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
func (it *InviteStakeWithdrawProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InviteStakeWithdrawProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InviteStakeWithdrawProfit represents a WithdrawProfit event raised by the InviteStake contract.
type InviteStakeWithdrawProfit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawProfit is a free log retrieval operation binding the contract event 0x010d214e8adbe593eebc2e78d29e88f08ddcb363fac75a9ef8c9455ba3c72dcc.
//
// Solidity: event WithdrawProfit(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) FilterWithdrawProfit(opts *bind.FilterOpts, user []common.Address) (*InviteStakeWithdrawProfitIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InviteStake.contract.FilterLogs(opts, "WithdrawProfit", userRule)
	if err != nil {
		return nil, err
	}
	return &InviteStakeWithdrawProfitIterator{contract: _InviteStake.contract, event: "WithdrawProfit", logs: logs, sub: sub}, nil
}

// WatchWithdrawProfit is a free log subscription operation binding the contract event 0x010d214e8adbe593eebc2e78d29e88f08ddcb363fac75a9ef8c9455ba3c72dcc.
//
// Solidity: event WithdrawProfit(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) WatchWithdrawProfit(opts *bind.WatchOpts, sink chan<- *InviteStakeWithdrawProfit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InviteStake.contract.WatchLogs(opts, "WithdrawProfit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InviteStakeWithdrawProfit)
				if err := _InviteStake.contract.UnpackLog(event, "WithdrawProfit", log); err != nil {
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

// ParseWithdrawProfit is a log parse operation binding the contract event 0x010d214e8adbe593eebc2e78d29e88f08ddcb363fac75a9ef8c9455ba3c72dcc.
//
// Solidity: event WithdrawProfit(address indexed user, uint256 amount)
func (_InviteStake *InviteStakeFilterer) ParseWithdrawProfit(log types.Log) (*InviteStakeWithdrawProfit, error) {
	event := new(InviteStakeWithdrawProfit)
	if err := _InviteStake.contract.UnpackLog(event, "WithdrawProfit", log); err != nil {
		return nil, err
	}
	return event, nil
}
