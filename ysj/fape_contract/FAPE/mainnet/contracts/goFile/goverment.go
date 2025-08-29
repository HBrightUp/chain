// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goverment

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

// GovermentstInviteRanking is an auto generated low-level Go binding around an user-defined struct.
type GovermentstInviteRanking struct {
	Addr             common.Address
	InviteContribute *big.Int
}

// GovermentstUserInfo is an auto generated low-level Go binding around an user-defined struct.
type GovermentstUserInfo struct {
	Superior                 common.Address
	SubordinateCounts1       *big.Int
	SubordinateCounts2       *big.Int
	GameToChain              *big.Int
	ChainToGame              *big.Int
	ExperiencePack           *big.Int
	AwardByInvited           *big.Int
	InviteAwardWithclaimed   *big.Int
	InviteAwardWithUnclaimed *big.Int
	InviteContribute         *big.Int
	InviteLatestTimestamp    *big.Int
	AwardWithGoldRanking     *big.Int
}

// GovermentMetaData contains all meta data concerning the Goverment contract.
var GovermentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"superior_\",\"type\":\"address\"}],\"name\":\"BindSuperior\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"superior_\",\"type\":\"address\"}],\"name\":\"BuyVip\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"FAPEChainToGameEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"FAPEGameToChainEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"FAPEChainToGame\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"FAPEGameToChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FAPEGiftAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"IsVipPlayer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_userIndexes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"addInviteRankingBlacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admissionTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"superior_\",\"type\":\"address\"}],\"name\":\"buyVip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimInviteRankingAward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentUserCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dapp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"employExpPack\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchequerUsdt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameAwardAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getExperiencePackCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getInviteRankingByAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInviteRankingList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inviteContribute\",\"type\":\"uint256\"}],\"internalType\":\"structGoverment.stInviteRanking[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getInvitedAwardByAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getSubordinate1Counts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getSubordinate2Counts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getSubordinateCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getSuperior\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDTAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"superior\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subordinateCounts1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subordinateCounts2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameToChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainToGame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"experiencePack\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"awardByInvited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inviteAwardWithclaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inviteAwardWithUnclaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inviteContribute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inviteLatestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"awardWithGoldRanking\",\"type\":\"uint256\"}],\"internalType\":\"structGoverment.stUserInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goldRankingAwardIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goldRankingCurentAward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goldRankingLatestAwardTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goldRankingTimesnap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goldRankingTotalItems\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inviteRankingAwardWightUnclaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inviteRankingBlacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inviteRankingContributeTotals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inviteRankingIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inviteRankingLatestAwardTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inviteRankingList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inviteRankingMinContribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inviteRankingSection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inviteRankingTimesnap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inviteRankingTotalItems\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpenExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpenGoldRanking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpenInviteRanking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"direction_\",\"type\":\"uint8\"}],\"name\":\"isSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients_\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rate_\",\"type\":\"uint256[]\"}],\"name\":\"issueGoldRankingAward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issueInviteRankingAward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxUserIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isOpenExchange_\",\"type\":\"bool\"}],\"name\":\"openExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isOpenGoldRanking_\",\"type\":\"bool\"}],\"name\":\"openGoldRanking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isOpenInviteRanking_\",\"type\":\"bool\"}],\"name\":\"openInviteRanking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"removeInviteRankingBlacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetRankingConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"admissionTicket_\",\"type\":\"uint256\"}],\"name\":\"setAdmissionTicket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dapp_\",\"type\":\"address\"}],\"name\":\"setDapp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"FAPEGiftAmount_\",\"type\":\"uint256\"}],\"name\":\"setFAPEGiftAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gameAwardAddr_\",\"type\":\"address\"}],\"name\":\"setGameAwardAddr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rankingTimesnap_\",\"type\":\"uint256\"}],\"name\":\"setGoldRankingTimesnap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rankingTotalItems_\",\"type\":\"uint256\"}],\"name\":\"setGoldRankingTotalItems\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rankingTimesnap_\",\"type\":\"uint256\"}],\"name\":\"setInviteRankingTimesnap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rankingTotalItems_\",\"type\":\"uint256\"}],\"name\":\"setInviteRankingTotalItems\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapTimesnap_\",\"type\":\"uint256\"}],\"name\":\"setSwapTimesnap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ticketNFT_\",\"type\":\"address\"}],\"name\":\"setTicketNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenFAPE_\",\"type\":\"address\"}],\"name\":\"setTokenFAPE\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapTimesnap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFAPE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GovermentABI is the input ABI used to generate the binding from.
// Deprecated: Use GovermentMetaData.ABI instead.
var GovermentABI = GovermentMetaData.ABI

// Goverment is an auto generated Go binding around an Ethereum contract.
type Goverment struct {
	GovermentCaller     // Read-only binding to the contract
	GovermentTransactor // Write-only binding to the contract
	GovermentFilterer   // Log filterer for contract events
}

// GovermentCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovermentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovermentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovermentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovermentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovermentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovermentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovermentSession struct {
	Contract     *Goverment        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovermentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovermentCallerSession struct {
	Contract *GovermentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GovermentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovermentTransactorSession struct {
	Contract     *GovermentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GovermentRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovermentRaw struct {
	Contract *Goverment // Generic contract binding to access the raw methods on
}

// GovermentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovermentCallerRaw struct {
	Contract *GovermentCaller // Generic read-only contract binding to access the raw methods on
}

// GovermentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovermentTransactorRaw struct {
	Contract *GovermentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGoverment creates a new instance of Goverment, bound to a specific deployed contract.
func NewGoverment(address common.Address, backend bind.ContractBackend) (*Goverment, error) {
	contract, err := bindGoverment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Goverment{GovermentCaller: GovermentCaller{contract: contract}, GovermentTransactor: GovermentTransactor{contract: contract}, GovermentFilterer: GovermentFilterer{contract: contract}}, nil
}

// NewGovermentCaller creates a new read-only instance of Goverment, bound to a specific deployed contract.
func NewGovermentCaller(address common.Address, caller bind.ContractCaller) (*GovermentCaller, error) {
	contract, err := bindGoverment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovermentCaller{contract: contract}, nil
}

// NewGovermentTransactor creates a new write-only instance of Goverment, bound to a specific deployed contract.
func NewGovermentTransactor(address common.Address, transactor bind.ContractTransactor) (*GovermentTransactor, error) {
	contract, err := bindGoverment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovermentTransactor{contract: contract}, nil
}

// NewGovermentFilterer creates a new log filterer instance of Goverment, bound to a specific deployed contract.
func NewGovermentFilterer(address common.Address, filterer bind.ContractFilterer) (*GovermentFilterer, error) {
	contract, err := bindGoverment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovermentFilterer{contract: contract}, nil
}

// bindGoverment binds a generic wrapper to an already deployed contract.
func bindGoverment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovermentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Goverment *GovermentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Goverment.Contract.GovermentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Goverment *GovermentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goverment.Contract.GovermentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Goverment *GovermentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Goverment.Contract.GovermentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Goverment *GovermentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Goverment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Goverment *GovermentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goverment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Goverment *GovermentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Goverment.Contract.contract.Transact(opts, method, params...)
}

// FAPEGiftAmount is a free data retrieval call binding the contract method 0xd0e8366b.
//
// Solidity: function FAPEGiftAmount() view returns(uint256)
func (_Goverment *GovermentCaller) FAPEGiftAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "FAPEGiftAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FAPEGiftAmount is a free data retrieval call binding the contract method 0xd0e8366b.
//
// Solidity: function FAPEGiftAmount() view returns(uint256)
func (_Goverment *GovermentSession) FAPEGiftAmount() (*big.Int, error) {
	return _Goverment.Contract.FAPEGiftAmount(&_Goverment.CallOpts)
}

// FAPEGiftAmount is a free data retrieval call binding the contract method 0xd0e8366b.
//
// Solidity: function FAPEGiftAmount() view returns(uint256)
func (_Goverment *GovermentCallerSession) FAPEGiftAmount() (*big.Int, error) {
	return _Goverment.Contract.FAPEGiftAmount(&_Goverment.CallOpts)
}

// IsVipPlayer is a free data retrieval call binding the contract method 0x08a15d3b.
//
// Solidity: function IsVipPlayer(address account_) view returns(uint256)
func (_Goverment *GovermentCaller) IsVipPlayer(opts *bind.CallOpts, account_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "IsVipPlayer", account_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsVipPlayer is a free data retrieval call binding the contract method 0x08a15d3b.
//
// Solidity: function IsVipPlayer(address account_) view returns(uint256)
func (_Goverment *GovermentSession) IsVipPlayer(account_ common.Address) (*big.Int, error) {
	return _Goverment.Contract.IsVipPlayer(&_Goverment.CallOpts, account_)
}

// IsVipPlayer is a free data retrieval call binding the contract method 0x08a15d3b.
//
// Solidity: function IsVipPlayer(address account_) view returns(uint256)
func (_Goverment *GovermentCallerSession) IsVipPlayer(account_ common.Address) (*big.Int, error) {
	return _Goverment.Contract.IsVipPlayer(&_Goverment.CallOpts, account_)
}

// UserIndexes is a free data retrieval call binding the contract method 0x8d225c7e.
//
// Solidity: function _userIndexes(uint256 ) view returns(address)
func (_Goverment *GovermentCaller) UserIndexes(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "_userIndexes", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserIndexes is a free data retrieval call binding the contract method 0x8d225c7e.
//
// Solidity: function _userIndexes(uint256 ) view returns(address)
func (_Goverment *GovermentSession) UserIndexes(arg0 *big.Int) (common.Address, error) {
	return _Goverment.Contract.UserIndexes(&_Goverment.CallOpts, arg0)
}

// UserIndexes is a free data retrieval call binding the contract method 0x8d225c7e.
//
// Solidity: function _userIndexes(uint256 ) view returns(address)
func (_Goverment *GovermentCallerSession) UserIndexes(arg0 *big.Int) (common.Address, error) {
	return _Goverment.Contract.UserIndexes(&_Goverment.CallOpts, arg0)
}

// AdmissionTicket is a free data retrieval call binding the contract method 0x0a664ee2.
//
// Solidity: function admissionTicket() view returns(uint256)
func (_Goverment *GovermentCaller) AdmissionTicket(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "admissionTicket")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdmissionTicket is a free data retrieval call binding the contract method 0x0a664ee2.
//
// Solidity: function admissionTicket() view returns(uint256)
func (_Goverment *GovermentSession) AdmissionTicket() (*big.Int, error) {
	return _Goverment.Contract.AdmissionTicket(&_Goverment.CallOpts)
}

// AdmissionTicket is a free data retrieval call binding the contract method 0x0a664ee2.
//
// Solidity: function admissionTicket() view returns(uint256)
func (_Goverment *GovermentCallerSession) AdmissionTicket() (*big.Int, error) {
	return _Goverment.Contract.AdmissionTicket(&_Goverment.CallOpts)
}

// CurrentUserCounts is a free data retrieval call binding the contract method 0x7e6a2fb3.
//
// Solidity: function currentUserCounts() view returns(uint256)
func (_Goverment *GovermentCaller) CurrentUserCounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "currentUserCounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentUserCounts is a free data retrieval call binding the contract method 0x7e6a2fb3.
//
// Solidity: function currentUserCounts() view returns(uint256)
func (_Goverment *GovermentSession) CurrentUserCounts() (*big.Int, error) {
	return _Goverment.Contract.CurrentUserCounts(&_Goverment.CallOpts)
}

// CurrentUserCounts is a free data retrieval call binding the contract method 0x7e6a2fb3.
//
// Solidity: function currentUserCounts() view returns(uint256)
func (_Goverment *GovermentCallerSession) CurrentUserCounts() (*big.Int, error) {
	return _Goverment.Contract.CurrentUserCounts(&_Goverment.CallOpts)
}

// Dapp is a free data retrieval call binding the contract method 0xb6de460a.
//
// Solidity: function dapp() view returns(address)
func (_Goverment *GovermentCaller) Dapp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "dapp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dapp is a free data retrieval call binding the contract method 0xb6de460a.
//
// Solidity: function dapp() view returns(address)
func (_Goverment *GovermentSession) Dapp() (common.Address, error) {
	return _Goverment.Contract.Dapp(&_Goverment.CallOpts)
}

// Dapp is a free data retrieval call binding the contract method 0xb6de460a.
//
// Solidity: function dapp() view returns(address)
func (_Goverment *GovermentCallerSession) Dapp() (common.Address, error) {
	return _Goverment.Contract.Dapp(&_Goverment.CallOpts)
}

// ExchequerUsdt is a free data retrieval call binding the contract method 0x87ae98e5.
//
// Solidity: function exchequerUsdt() view returns(uint256)
func (_Goverment *GovermentCaller) ExchequerUsdt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "exchequerUsdt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchequerUsdt is a free data retrieval call binding the contract method 0x87ae98e5.
//
// Solidity: function exchequerUsdt() view returns(uint256)
func (_Goverment *GovermentSession) ExchequerUsdt() (*big.Int, error) {
	return _Goverment.Contract.ExchequerUsdt(&_Goverment.CallOpts)
}

// ExchequerUsdt is a free data retrieval call binding the contract method 0x87ae98e5.
//
// Solidity: function exchequerUsdt() view returns(uint256)
func (_Goverment *GovermentCallerSession) ExchequerUsdt() (*big.Int, error) {
	return _Goverment.Contract.ExchequerUsdt(&_Goverment.CallOpts)
}

// GameAwardAddr is a free data retrieval call binding the contract method 0x40244d56.
//
// Solidity: function gameAwardAddr() view returns(address)
func (_Goverment *GovermentCaller) GameAwardAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "gameAwardAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameAwardAddr is a free data retrieval call binding the contract method 0x40244d56.
//
// Solidity: function gameAwardAddr() view returns(address)
func (_Goverment *GovermentSession) GameAwardAddr() (common.Address, error) {
	return _Goverment.Contract.GameAwardAddr(&_Goverment.CallOpts)
}

// GameAwardAddr is a free data retrieval call binding the contract method 0x40244d56.
//
// Solidity: function gameAwardAddr() view returns(address)
func (_Goverment *GovermentCallerSession) GameAwardAddr() (common.Address, error) {
	return _Goverment.Contract.GameAwardAddr(&_Goverment.CallOpts)
}

// GetExperiencePackCounts is a free data retrieval call binding the contract method 0x02447a21.
//
// Solidity: function getExperiencePackCounts(address account_) view returns(uint256)
func (_Goverment *GovermentCaller) GetExperiencePackCounts(opts *bind.CallOpts, account_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "getExperiencePackCounts", account_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExperiencePackCounts is a free data retrieval call binding the contract method 0x02447a21.
//
// Solidity: function getExperiencePackCounts(address account_) view returns(uint256)
func (_Goverment *GovermentSession) GetExperiencePackCounts(account_ common.Address) (*big.Int, error) {
	return _Goverment.Contract.GetExperiencePackCounts(&_Goverment.CallOpts, account_)
}

// GetExperiencePackCounts is a free data retrieval call binding the contract method 0x02447a21.
//
// Solidity: function getExperiencePackCounts(address account_) view returns(uint256)
func (_Goverment *GovermentCallerSession) GetExperiencePackCounts(account_ common.Address) (*big.Int, error) {
	return _Goverment.Contract.GetExperiencePackCounts(&_Goverment.CallOpts, account_)
}

// GetInviteRankingByAccount is a free data retrieval call binding the contract method 0xf2fd8c09.
//
// Solidity: function getInviteRankingByAccount(address account_) view returns(uint256, uint256, uint256, uint256, uint256)
func (_Goverment *GovermentCaller) GetInviteRankingByAccount(opts *bind.CallOpts, account_ common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "getInviteRankingByAccount", account_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetInviteRankingByAccount is a free data retrieval call binding the contract method 0xf2fd8c09.
//
// Solidity: function getInviteRankingByAccount(address account_) view returns(uint256, uint256, uint256, uint256, uint256)
func (_Goverment *GovermentSession) GetInviteRankingByAccount(account_ common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Goverment.Contract.GetInviteRankingByAccount(&_Goverment.CallOpts, account_)
}

// GetInviteRankingByAccount is a free data retrieval call binding the contract method 0xf2fd8c09.
//
// Solidity: function getInviteRankingByAccount(address account_) view returns(uint256, uint256, uint256, uint256, uint256)
func (_Goverment *GovermentCallerSession) GetInviteRankingByAccount(account_ common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Goverment.Contract.GetInviteRankingByAccount(&_Goverment.CallOpts, account_)
}

// GetInviteRankingList is a free data retrieval call binding the contract method 0xacff6978.
//
// Solidity: function getInviteRankingList() view returns((address,uint256)[])
func (_Goverment *GovermentCaller) GetInviteRankingList(opts *bind.CallOpts) ([]GovermentstInviteRanking, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "getInviteRankingList")

	if err != nil {
		return *new([]GovermentstInviteRanking), err
	}

	out0 := *abi.ConvertType(out[0], new([]GovermentstInviteRanking)).(*[]GovermentstInviteRanking)

	return out0, err

}

// GetInviteRankingList is a free data retrieval call binding the contract method 0xacff6978.
//
// Solidity: function getInviteRankingList() view returns((address,uint256)[])
func (_Goverment *GovermentSession) GetInviteRankingList() ([]GovermentstInviteRanking, error) {
	return _Goverment.Contract.GetInviteRankingList(&_Goverment.CallOpts)
}

// GetInviteRankingList is a free data retrieval call binding the contract method 0xacff6978.
//
// Solidity: function getInviteRankingList() view returns((address,uint256)[])
func (_Goverment *GovermentCallerSession) GetInviteRankingList() ([]GovermentstInviteRanking, error) {
	return _Goverment.Contract.GetInviteRankingList(&_Goverment.CallOpts)
}

// GetInvitedAwardByAccount is a free data retrieval call binding the contract method 0xc0a9541e.
//
// Solidity: function getInvitedAwardByAccount(address account_) view returns(uint256)
func (_Goverment *GovermentCaller) GetInvitedAwardByAccount(opts *bind.CallOpts, account_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "getInvitedAwardByAccount", account_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInvitedAwardByAccount is a free data retrieval call binding the contract method 0xc0a9541e.
//
// Solidity: function getInvitedAwardByAccount(address account_) view returns(uint256)
func (_Goverment *GovermentSession) GetInvitedAwardByAccount(account_ common.Address) (*big.Int, error) {
	return _Goverment.Contract.GetInvitedAwardByAccount(&_Goverment.CallOpts, account_)
}

// GetInvitedAwardByAccount is a free data retrieval call binding the contract method 0xc0a9541e.
//
// Solidity: function getInvitedAwardByAccount(address account_) view returns(uint256)
func (_Goverment *GovermentCallerSession) GetInvitedAwardByAccount(account_ common.Address) (*big.Int, error) {
	return _Goverment.Contract.GetInvitedAwardByAccount(&_Goverment.CallOpts, account_)
}

// GetSubordinate1Counts is a free data retrieval call binding the contract method 0x884e76ff.
//
// Solidity: function getSubordinate1Counts(address account_) view returns(uint256)
func (_Goverment *GovermentCaller) GetSubordinate1Counts(opts *bind.CallOpts, account_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "getSubordinate1Counts", account_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubordinate1Counts is a free data retrieval call binding the contract method 0x884e76ff.
//
// Solidity: function getSubordinate1Counts(address account_) view returns(uint256)
func (_Goverment *GovermentSession) GetSubordinate1Counts(account_ common.Address) (*big.Int, error) {
	return _Goverment.Contract.GetSubordinate1Counts(&_Goverment.CallOpts, account_)
}

// GetSubordinate1Counts is a free data retrieval call binding the contract method 0x884e76ff.
//
// Solidity: function getSubordinate1Counts(address account_) view returns(uint256)
func (_Goverment *GovermentCallerSession) GetSubordinate1Counts(account_ common.Address) (*big.Int, error) {
	return _Goverment.Contract.GetSubordinate1Counts(&_Goverment.CallOpts, account_)
}

// GetSubordinate2Counts is a free data retrieval call binding the contract method 0x5f89b41e.
//
// Solidity: function getSubordinate2Counts(address account_) view returns(uint256)
func (_Goverment *GovermentCaller) GetSubordinate2Counts(opts *bind.CallOpts, account_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "getSubordinate2Counts", account_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubordinate2Counts is a free data retrieval call binding the contract method 0x5f89b41e.
//
// Solidity: function getSubordinate2Counts(address account_) view returns(uint256)
func (_Goverment *GovermentSession) GetSubordinate2Counts(account_ common.Address) (*big.Int, error) {
	return _Goverment.Contract.GetSubordinate2Counts(&_Goverment.CallOpts, account_)
}

// GetSubordinate2Counts is a free data retrieval call binding the contract method 0x5f89b41e.
//
// Solidity: function getSubordinate2Counts(address account_) view returns(uint256)
func (_Goverment *GovermentCallerSession) GetSubordinate2Counts(account_ common.Address) (*big.Int, error) {
	return _Goverment.Contract.GetSubordinate2Counts(&_Goverment.CallOpts, account_)
}

// GetSubordinateCounts is a free data retrieval call binding the contract method 0xe2c3813c.
//
// Solidity: function getSubordinateCounts(address account_) view returns(uint256)
func (_Goverment *GovermentCaller) GetSubordinateCounts(opts *bind.CallOpts, account_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "getSubordinateCounts", account_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubordinateCounts is a free data retrieval call binding the contract method 0xe2c3813c.
//
// Solidity: function getSubordinateCounts(address account_) view returns(uint256)
func (_Goverment *GovermentSession) GetSubordinateCounts(account_ common.Address) (*big.Int, error) {
	return _Goverment.Contract.GetSubordinateCounts(&_Goverment.CallOpts, account_)
}

// GetSubordinateCounts is a free data retrieval call binding the contract method 0xe2c3813c.
//
// Solidity: function getSubordinateCounts(address account_) view returns(uint256)
func (_Goverment *GovermentCallerSession) GetSubordinateCounts(account_ common.Address) (*big.Int, error) {
	return _Goverment.Contract.GetSubordinateCounts(&_Goverment.CallOpts, account_)
}

// GetSuperior is a free data retrieval call binding the contract method 0x443355e5.
//
// Solidity: function getSuperior(address account_) view returns(address)
func (_Goverment *GovermentCaller) GetSuperior(opts *bind.CallOpts, account_ common.Address) (common.Address, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "getSuperior", account_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSuperior is a free data retrieval call binding the contract method 0x443355e5.
//
// Solidity: function getSuperior(address account_) view returns(address)
func (_Goverment *GovermentSession) GetSuperior(account_ common.Address) (common.Address, error) {
	return _Goverment.Contract.GetSuperior(&_Goverment.CallOpts, account_)
}

// GetSuperior is a free data retrieval call binding the contract method 0x443355e5.
//
// Solidity: function getSuperior(address account_) view returns(address)
func (_Goverment *GovermentCallerSession) GetSuperior(account_ common.Address) (common.Address, error) {
	return _Goverment.Contract.GetSuperior(&_Goverment.CallOpts, account_)
}

// GetUSDTAddress is a free data retrieval call binding the contract method 0x8af3c40e.
//
// Solidity: function getUSDTAddress() pure returns(address)
func (_Goverment *GovermentCaller) GetUSDTAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "getUSDTAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUSDTAddress is a free data retrieval call binding the contract method 0x8af3c40e.
//
// Solidity: function getUSDTAddress() pure returns(address)
func (_Goverment *GovermentSession) GetUSDTAddress() (common.Address, error) {
	return _Goverment.Contract.GetUSDTAddress(&_Goverment.CallOpts)
}

// GetUSDTAddress is a free data retrieval call binding the contract method 0x8af3c40e.
//
// Solidity: function getUSDTAddress() pure returns(address)
func (_Goverment *GovermentCallerSession) GetUSDTAddress() (common.Address, error) {
	return _Goverment.Contract.GetUSDTAddress(&_Goverment.CallOpts)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address account_) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Goverment *GovermentCaller) GetUserInfo(opts *bind.CallOpts, account_ common.Address) (GovermentstUserInfo, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "getUserInfo", account_)

	if err != nil {
		return *new(GovermentstUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(GovermentstUserInfo)).(*GovermentstUserInfo)

	return out0, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address account_) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Goverment *GovermentSession) GetUserInfo(account_ common.Address) (GovermentstUserInfo, error) {
	return _Goverment.Contract.GetUserInfo(&_Goverment.CallOpts, account_)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address account_) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Goverment *GovermentCallerSession) GetUserInfo(account_ common.Address) (GovermentstUserInfo, error) {
	return _Goverment.Contract.GetUserInfo(&_Goverment.CallOpts, account_)
}

// GoldRankingAwardIndex is a free data retrieval call binding the contract method 0x044a7750.
//
// Solidity: function goldRankingAwardIndex() view returns(uint256)
func (_Goverment *GovermentCaller) GoldRankingAwardIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "goldRankingAwardIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GoldRankingAwardIndex is a free data retrieval call binding the contract method 0x044a7750.
//
// Solidity: function goldRankingAwardIndex() view returns(uint256)
func (_Goverment *GovermentSession) GoldRankingAwardIndex() (*big.Int, error) {
	return _Goverment.Contract.GoldRankingAwardIndex(&_Goverment.CallOpts)
}

// GoldRankingAwardIndex is a free data retrieval call binding the contract method 0x044a7750.
//
// Solidity: function goldRankingAwardIndex() view returns(uint256)
func (_Goverment *GovermentCallerSession) GoldRankingAwardIndex() (*big.Int, error) {
	return _Goverment.Contract.GoldRankingAwardIndex(&_Goverment.CallOpts)
}

// GoldRankingCurentAward is a free data retrieval call binding the contract method 0xc64809f1.
//
// Solidity: function goldRankingCurentAward() view returns(uint256)
func (_Goverment *GovermentCaller) GoldRankingCurentAward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "goldRankingCurentAward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GoldRankingCurentAward is a free data retrieval call binding the contract method 0xc64809f1.
//
// Solidity: function goldRankingCurentAward() view returns(uint256)
func (_Goverment *GovermentSession) GoldRankingCurentAward() (*big.Int, error) {
	return _Goverment.Contract.GoldRankingCurentAward(&_Goverment.CallOpts)
}

// GoldRankingCurentAward is a free data retrieval call binding the contract method 0xc64809f1.
//
// Solidity: function goldRankingCurentAward() view returns(uint256)
func (_Goverment *GovermentCallerSession) GoldRankingCurentAward() (*big.Int, error) {
	return _Goverment.Contract.GoldRankingCurentAward(&_Goverment.CallOpts)
}

// GoldRankingLatestAwardTimestamp is a free data retrieval call binding the contract method 0x4268074d.
//
// Solidity: function goldRankingLatestAwardTimestamp() view returns(uint256)
func (_Goverment *GovermentCaller) GoldRankingLatestAwardTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "goldRankingLatestAwardTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GoldRankingLatestAwardTimestamp is a free data retrieval call binding the contract method 0x4268074d.
//
// Solidity: function goldRankingLatestAwardTimestamp() view returns(uint256)
func (_Goverment *GovermentSession) GoldRankingLatestAwardTimestamp() (*big.Int, error) {
	return _Goverment.Contract.GoldRankingLatestAwardTimestamp(&_Goverment.CallOpts)
}

// GoldRankingLatestAwardTimestamp is a free data retrieval call binding the contract method 0x4268074d.
//
// Solidity: function goldRankingLatestAwardTimestamp() view returns(uint256)
func (_Goverment *GovermentCallerSession) GoldRankingLatestAwardTimestamp() (*big.Int, error) {
	return _Goverment.Contract.GoldRankingLatestAwardTimestamp(&_Goverment.CallOpts)
}

// GoldRankingTimesnap is a free data retrieval call binding the contract method 0x3d0bdcf5.
//
// Solidity: function goldRankingTimesnap() view returns(uint256)
func (_Goverment *GovermentCaller) GoldRankingTimesnap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "goldRankingTimesnap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GoldRankingTimesnap is a free data retrieval call binding the contract method 0x3d0bdcf5.
//
// Solidity: function goldRankingTimesnap() view returns(uint256)
func (_Goverment *GovermentSession) GoldRankingTimesnap() (*big.Int, error) {
	return _Goverment.Contract.GoldRankingTimesnap(&_Goverment.CallOpts)
}

// GoldRankingTimesnap is a free data retrieval call binding the contract method 0x3d0bdcf5.
//
// Solidity: function goldRankingTimesnap() view returns(uint256)
func (_Goverment *GovermentCallerSession) GoldRankingTimesnap() (*big.Int, error) {
	return _Goverment.Contract.GoldRankingTimesnap(&_Goverment.CallOpts)
}

// GoldRankingTotalItems is a free data retrieval call binding the contract method 0xb82dff51.
//
// Solidity: function goldRankingTotalItems() view returns(uint256)
func (_Goverment *GovermentCaller) GoldRankingTotalItems(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "goldRankingTotalItems")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GoldRankingTotalItems is a free data retrieval call binding the contract method 0xb82dff51.
//
// Solidity: function goldRankingTotalItems() view returns(uint256)
func (_Goverment *GovermentSession) GoldRankingTotalItems() (*big.Int, error) {
	return _Goverment.Contract.GoldRankingTotalItems(&_Goverment.CallOpts)
}

// GoldRankingTotalItems is a free data retrieval call binding the contract method 0xb82dff51.
//
// Solidity: function goldRankingTotalItems() view returns(uint256)
func (_Goverment *GovermentCallerSession) GoldRankingTotalItems() (*big.Int, error) {
	return _Goverment.Contract.GoldRankingTotalItems(&_Goverment.CallOpts)
}

// InviteRankingAwardWightUnclaim is a free data retrieval call binding the contract method 0xabc32268.
//
// Solidity: function inviteRankingAwardWightUnclaim() view returns(uint256)
func (_Goverment *GovermentCaller) InviteRankingAwardWightUnclaim(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "inviteRankingAwardWightUnclaim")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteRankingAwardWightUnclaim is a free data retrieval call binding the contract method 0xabc32268.
//
// Solidity: function inviteRankingAwardWightUnclaim() view returns(uint256)
func (_Goverment *GovermentSession) InviteRankingAwardWightUnclaim() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingAwardWightUnclaim(&_Goverment.CallOpts)
}

// InviteRankingAwardWightUnclaim is a free data retrieval call binding the contract method 0xabc32268.
//
// Solidity: function inviteRankingAwardWightUnclaim() view returns(uint256)
func (_Goverment *GovermentCallerSession) InviteRankingAwardWightUnclaim() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingAwardWightUnclaim(&_Goverment.CallOpts)
}

// InviteRankingBlacklist is a free data retrieval call binding the contract method 0x5de983f3.
//
// Solidity: function inviteRankingBlacklist(address ) view returns(bool)
func (_Goverment *GovermentCaller) InviteRankingBlacklist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "inviteRankingBlacklist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InviteRankingBlacklist is a free data retrieval call binding the contract method 0x5de983f3.
//
// Solidity: function inviteRankingBlacklist(address ) view returns(bool)
func (_Goverment *GovermentSession) InviteRankingBlacklist(arg0 common.Address) (bool, error) {
	return _Goverment.Contract.InviteRankingBlacklist(&_Goverment.CallOpts, arg0)
}

// InviteRankingBlacklist is a free data retrieval call binding the contract method 0x5de983f3.
//
// Solidity: function inviteRankingBlacklist(address ) view returns(bool)
func (_Goverment *GovermentCallerSession) InviteRankingBlacklist(arg0 common.Address) (bool, error) {
	return _Goverment.Contract.InviteRankingBlacklist(&_Goverment.CallOpts, arg0)
}

// InviteRankingContributeTotals is a free data retrieval call binding the contract method 0xc7aa1aca.
//
// Solidity: function inviteRankingContributeTotals() view returns(uint256)
func (_Goverment *GovermentCaller) InviteRankingContributeTotals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "inviteRankingContributeTotals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteRankingContributeTotals is a free data retrieval call binding the contract method 0xc7aa1aca.
//
// Solidity: function inviteRankingContributeTotals() view returns(uint256)
func (_Goverment *GovermentSession) InviteRankingContributeTotals() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingContributeTotals(&_Goverment.CallOpts)
}

// InviteRankingContributeTotals is a free data retrieval call binding the contract method 0xc7aa1aca.
//
// Solidity: function inviteRankingContributeTotals() view returns(uint256)
func (_Goverment *GovermentCallerSession) InviteRankingContributeTotals() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingContributeTotals(&_Goverment.CallOpts)
}

// InviteRankingIndex is a free data retrieval call binding the contract method 0xa2ca4c3d.
//
// Solidity: function inviteRankingIndex() view returns(uint256)
func (_Goverment *GovermentCaller) InviteRankingIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "inviteRankingIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteRankingIndex is a free data retrieval call binding the contract method 0xa2ca4c3d.
//
// Solidity: function inviteRankingIndex() view returns(uint256)
func (_Goverment *GovermentSession) InviteRankingIndex() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingIndex(&_Goverment.CallOpts)
}

// InviteRankingIndex is a free data retrieval call binding the contract method 0xa2ca4c3d.
//
// Solidity: function inviteRankingIndex() view returns(uint256)
func (_Goverment *GovermentCallerSession) InviteRankingIndex() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingIndex(&_Goverment.CallOpts)
}

// InviteRankingLatestAwardTimestamp is a free data retrieval call binding the contract method 0x24fac067.
//
// Solidity: function inviteRankingLatestAwardTimestamp() view returns(uint256)
func (_Goverment *GovermentCaller) InviteRankingLatestAwardTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "inviteRankingLatestAwardTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteRankingLatestAwardTimestamp is a free data retrieval call binding the contract method 0x24fac067.
//
// Solidity: function inviteRankingLatestAwardTimestamp() view returns(uint256)
func (_Goverment *GovermentSession) InviteRankingLatestAwardTimestamp() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingLatestAwardTimestamp(&_Goverment.CallOpts)
}

// InviteRankingLatestAwardTimestamp is a free data retrieval call binding the contract method 0x24fac067.
//
// Solidity: function inviteRankingLatestAwardTimestamp() view returns(uint256)
func (_Goverment *GovermentCallerSession) InviteRankingLatestAwardTimestamp() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingLatestAwardTimestamp(&_Goverment.CallOpts)
}

// InviteRankingList is a free data retrieval call binding the contract method 0x6f131fde.
//
// Solidity: function inviteRankingList(uint256 ) view returns(address)
func (_Goverment *GovermentCaller) InviteRankingList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "inviteRankingList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InviteRankingList is a free data retrieval call binding the contract method 0x6f131fde.
//
// Solidity: function inviteRankingList(uint256 ) view returns(address)
func (_Goverment *GovermentSession) InviteRankingList(arg0 *big.Int) (common.Address, error) {
	return _Goverment.Contract.InviteRankingList(&_Goverment.CallOpts, arg0)
}

// InviteRankingList is a free data retrieval call binding the contract method 0x6f131fde.
//
// Solidity: function inviteRankingList(uint256 ) view returns(address)
func (_Goverment *GovermentCallerSession) InviteRankingList(arg0 *big.Int) (common.Address, error) {
	return _Goverment.Contract.InviteRankingList(&_Goverment.CallOpts, arg0)
}

// InviteRankingMinContribute is a free data retrieval call binding the contract method 0xc81d040d.
//
// Solidity: function inviteRankingMinContribute() view returns(uint256)
func (_Goverment *GovermentCaller) InviteRankingMinContribute(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "inviteRankingMinContribute")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteRankingMinContribute is a free data retrieval call binding the contract method 0xc81d040d.
//
// Solidity: function inviteRankingMinContribute() view returns(uint256)
func (_Goverment *GovermentSession) InviteRankingMinContribute() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingMinContribute(&_Goverment.CallOpts)
}

// InviteRankingMinContribute is a free data retrieval call binding the contract method 0xc81d040d.
//
// Solidity: function inviteRankingMinContribute() view returns(uint256)
func (_Goverment *GovermentCallerSession) InviteRankingMinContribute() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingMinContribute(&_Goverment.CallOpts)
}

// InviteRankingSection is a free data retrieval call binding the contract method 0x91e0f15d.
//
// Solidity: function inviteRankingSection() view returns(uint256)
func (_Goverment *GovermentCaller) InviteRankingSection(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "inviteRankingSection")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteRankingSection is a free data retrieval call binding the contract method 0x91e0f15d.
//
// Solidity: function inviteRankingSection() view returns(uint256)
func (_Goverment *GovermentSession) InviteRankingSection() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingSection(&_Goverment.CallOpts)
}

// InviteRankingSection is a free data retrieval call binding the contract method 0x91e0f15d.
//
// Solidity: function inviteRankingSection() view returns(uint256)
func (_Goverment *GovermentCallerSession) InviteRankingSection() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingSection(&_Goverment.CallOpts)
}

// InviteRankingTimesnap is a free data retrieval call binding the contract method 0xe6c9a2bc.
//
// Solidity: function inviteRankingTimesnap() view returns(uint256)
func (_Goverment *GovermentCaller) InviteRankingTimesnap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "inviteRankingTimesnap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteRankingTimesnap is a free data retrieval call binding the contract method 0xe6c9a2bc.
//
// Solidity: function inviteRankingTimesnap() view returns(uint256)
func (_Goverment *GovermentSession) InviteRankingTimesnap() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingTimesnap(&_Goverment.CallOpts)
}

// InviteRankingTimesnap is a free data retrieval call binding the contract method 0xe6c9a2bc.
//
// Solidity: function inviteRankingTimesnap() view returns(uint256)
func (_Goverment *GovermentCallerSession) InviteRankingTimesnap() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingTimesnap(&_Goverment.CallOpts)
}

// InviteRankingTotalItems is a free data retrieval call binding the contract method 0xc6241e78.
//
// Solidity: function inviteRankingTotalItems() view returns(uint256)
func (_Goverment *GovermentCaller) InviteRankingTotalItems(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "inviteRankingTotalItems")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteRankingTotalItems is a free data retrieval call binding the contract method 0xc6241e78.
//
// Solidity: function inviteRankingTotalItems() view returns(uint256)
func (_Goverment *GovermentSession) InviteRankingTotalItems() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingTotalItems(&_Goverment.CallOpts)
}

// InviteRankingTotalItems is a free data retrieval call binding the contract method 0xc6241e78.
//
// Solidity: function inviteRankingTotalItems() view returns(uint256)
func (_Goverment *GovermentCallerSession) InviteRankingTotalItems() (*big.Int, error) {
	return _Goverment.Contract.InviteRankingTotalItems(&_Goverment.CallOpts)
}

// IsOpenExchange is a free data retrieval call binding the contract method 0x52feb70b.
//
// Solidity: function isOpenExchange() view returns(bool)
func (_Goverment *GovermentCaller) IsOpenExchange(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "isOpenExchange")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpenExchange is a free data retrieval call binding the contract method 0x52feb70b.
//
// Solidity: function isOpenExchange() view returns(bool)
func (_Goverment *GovermentSession) IsOpenExchange() (bool, error) {
	return _Goverment.Contract.IsOpenExchange(&_Goverment.CallOpts)
}

// IsOpenExchange is a free data retrieval call binding the contract method 0x52feb70b.
//
// Solidity: function isOpenExchange() view returns(bool)
func (_Goverment *GovermentCallerSession) IsOpenExchange() (bool, error) {
	return _Goverment.Contract.IsOpenExchange(&_Goverment.CallOpts)
}

// IsOpenGoldRanking is a free data retrieval call binding the contract method 0xd7d69f9e.
//
// Solidity: function isOpenGoldRanking() view returns(bool)
func (_Goverment *GovermentCaller) IsOpenGoldRanking(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "isOpenGoldRanking")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpenGoldRanking is a free data retrieval call binding the contract method 0xd7d69f9e.
//
// Solidity: function isOpenGoldRanking() view returns(bool)
func (_Goverment *GovermentSession) IsOpenGoldRanking() (bool, error) {
	return _Goverment.Contract.IsOpenGoldRanking(&_Goverment.CallOpts)
}

// IsOpenGoldRanking is a free data retrieval call binding the contract method 0xd7d69f9e.
//
// Solidity: function isOpenGoldRanking() view returns(bool)
func (_Goverment *GovermentCallerSession) IsOpenGoldRanking() (bool, error) {
	return _Goverment.Contract.IsOpenGoldRanking(&_Goverment.CallOpts)
}

// IsOpenInviteRanking is a free data retrieval call binding the contract method 0x3624e306.
//
// Solidity: function isOpenInviteRanking() view returns(bool)
func (_Goverment *GovermentCaller) IsOpenInviteRanking(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "isOpenInviteRanking")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpenInviteRanking is a free data retrieval call binding the contract method 0x3624e306.
//
// Solidity: function isOpenInviteRanking() view returns(bool)
func (_Goverment *GovermentSession) IsOpenInviteRanking() (bool, error) {
	return _Goverment.Contract.IsOpenInviteRanking(&_Goverment.CallOpts)
}

// IsOpenInviteRanking is a free data retrieval call binding the contract method 0x3624e306.
//
// Solidity: function isOpenInviteRanking() view returns(bool)
func (_Goverment *GovermentCallerSession) IsOpenInviteRanking() (bool, error) {
	return _Goverment.Contract.IsOpenInviteRanking(&_Goverment.CallOpts)
}

// IsSwap is a free data retrieval call binding the contract method 0x2618c5ea.
//
// Solidity: function isSwap(address account_, uint8 direction_) view returns(bool)
func (_Goverment *GovermentCaller) IsSwap(opts *bind.CallOpts, account_ common.Address, direction_ uint8) (bool, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "isSwap", account_, direction_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSwap is a free data retrieval call binding the contract method 0x2618c5ea.
//
// Solidity: function isSwap(address account_, uint8 direction_) view returns(bool)
func (_Goverment *GovermentSession) IsSwap(account_ common.Address, direction_ uint8) (bool, error) {
	return _Goverment.Contract.IsSwap(&_Goverment.CallOpts, account_, direction_)
}

// IsSwap is a free data retrieval call binding the contract method 0x2618c5ea.
//
// Solidity: function isSwap(address account_, uint8 direction_) view returns(bool)
func (_Goverment *GovermentCallerSession) IsSwap(account_ common.Address, direction_ uint8) (bool, error) {
	return _Goverment.Contract.IsSwap(&_Goverment.CallOpts, account_, direction_)
}

// MaxUserIndex is a free data retrieval call binding the contract method 0x3cbeb546.
//
// Solidity: function maxUserIndex() view returns(uint256)
func (_Goverment *GovermentCaller) MaxUserIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "maxUserIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxUserIndex is a free data retrieval call binding the contract method 0x3cbeb546.
//
// Solidity: function maxUserIndex() view returns(uint256)
func (_Goverment *GovermentSession) MaxUserIndex() (*big.Int, error) {
	return _Goverment.Contract.MaxUserIndex(&_Goverment.CallOpts)
}

// MaxUserIndex is a free data retrieval call binding the contract method 0x3cbeb546.
//
// Solidity: function maxUserIndex() view returns(uint256)
func (_Goverment *GovermentCallerSession) MaxUserIndex() (*big.Int, error) {
	return _Goverment.Contract.MaxUserIndex(&_Goverment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Goverment *GovermentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Goverment *GovermentSession) Owner() (common.Address, error) {
	return _Goverment.Contract.Owner(&_Goverment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Goverment *GovermentCallerSession) Owner() (common.Address, error) {
	return _Goverment.Contract.Owner(&_Goverment.CallOpts)
}

// SwapTimesnap is a free data retrieval call binding the contract method 0x3168b8eb.
//
// Solidity: function swapTimesnap() view returns(uint256)
func (_Goverment *GovermentCaller) SwapTimesnap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "swapTimesnap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapTimesnap is a free data retrieval call binding the contract method 0x3168b8eb.
//
// Solidity: function swapTimesnap() view returns(uint256)
func (_Goverment *GovermentSession) SwapTimesnap() (*big.Int, error) {
	return _Goverment.Contract.SwapTimesnap(&_Goverment.CallOpts)
}

// SwapTimesnap is a free data retrieval call binding the contract method 0x3168b8eb.
//
// Solidity: function swapTimesnap() view returns(uint256)
func (_Goverment *GovermentCallerSession) SwapTimesnap() (*big.Int, error) {
	return _Goverment.Contract.SwapTimesnap(&_Goverment.CallOpts)
}

// TicketNFT is a free data retrieval call binding the contract method 0xb393391b.
//
// Solidity: function ticketNFT() view returns(address)
func (_Goverment *GovermentCaller) TicketNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "ticketNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TicketNFT is a free data retrieval call binding the contract method 0xb393391b.
//
// Solidity: function ticketNFT() view returns(address)
func (_Goverment *GovermentSession) TicketNFT() (common.Address, error) {
	return _Goverment.Contract.TicketNFT(&_Goverment.CallOpts)
}

// TicketNFT is a free data retrieval call binding the contract method 0xb393391b.
//
// Solidity: function ticketNFT() view returns(address)
func (_Goverment *GovermentCallerSession) TicketNFT() (common.Address, error) {
	return _Goverment.Contract.TicketNFT(&_Goverment.CallOpts)
}

// TokenFAPE is a free data retrieval call binding the contract method 0x89694dcf.
//
// Solidity: function tokenFAPE() view returns(address)
func (_Goverment *GovermentCaller) TokenFAPE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Goverment.contract.Call(opts, &out, "tokenFAPE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFAPE is a free data retrieval call binding the contract method 0x89694dcf.
//
// Solidity: function tokenFAPE() view returns(address)
func (_Goverment *GovermentSession) TokenFAPE() (common.Address, error) {
	return _Goverment.Contract.TokenFAPE(&_Goverment.CallOpts)
}

// TokenFAPE is a free data retrieval call binding the contract method 0x89694dcf.
//
// Solidity: function tokenFAPE() view returns(address)
func (_Goverment *GovermentCallerSession) TokenFAPE() (common.Address, error) {
	return _Goverment.Contract.TokenFAPE(&_Goverment.CallOpts)
}

// FAPEChainToGame is a paid mutator transaction binding the contract method 0x4bc7e6db.
//
// Solidity: function FAPEChainToGame(address account_, uint256 amount_) returns(bool)
func (_Goverment *GovermentTransactor) FAPEChainToGame(opts *bind.TransactOpts, account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "FAPEChainToGame", account_, amount_)
}

// FAPEChainToGame is a paid mutator transaction binding the contract method 0x4bc7e6db.
//
// Solidity: function FAPEChainToGame(address account_, uint256 amount_) returns(bool)
func (_Goverment *GovermentSession) FAPEChainToGame(account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.FAPEChainToGame(&_Goverment.TransactOpts, account_, amount_)
}

// FAPEChainToGame is a paid mutator transaction binding the contract method 0x4bc7e6db.
//
// Solidity: function FAPEChainToGame(address account_, uint256 amount_) returns(bool)
func (_Goverment *GovermentTransactorSession) FAPEChainToGame(account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.FAPEChainToGame(&_Goverment.TransactOpts, account_, amount_)
}

// FAPEGameToChain is a paid mutator transaction binding the contract method 0x27f03572.
//
// Solidity: function FAPEGameToChain(address account_, uint256 amount_) returns(bool)
func (_Goverment *GovermentTransactor) FAPEGameToChain(opts *bind.TransactOpts, account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "FAPEGameToChain", account_, amount_)
}

// FAPEGameToChain is a paid mutator transaction binding the contract method 0x27f03572.
//
// Solidity: function FAPEGameToChain(address account_, uint256 amount_) returns(bool)
func (_Goverment *GovermentSession) FAPEGameToChain(account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.FAPEGameToChain(&_Goverment.TransactOpts, account_, amount_)
}

// FAPEGameToChain is a paid mutator transaction binding the contract method 0x27f03572.
//
// Solidity: function FAPEGameToChain(address account_, uint256 amount_) returns(bool)
func (_Goverment *GovermentTransactorSession) FAPEGameToChain(account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.FAPEGameToChain(&_Goverment.TransactOpts, account_, amount_)
}

// AddInviteRankingBlacklist is a paid mutator transaction binding the contract method 0x7dedb468.
//
// Solidity: function addInviteRankingBlacklist(address account_) returns(bool)
func (_Goverment *GovermentTransactor) AddInviteRankingBlacklist(opts *bind.TransactOpts, account_ common.Address) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "addInviteRankingBlacklist", account_)
}

// AddInviteRankingBlacklist is a paid mutator transaction binding the contract method 0x7dedb468.
//
// Solidity: function addInviteRankingBlacklist(address account_) returns(bool)
func (_Goverment *GovermentSession) AddInviteRankingBlacklist(account_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.AddInviteRankingBlacklist(&_Goverment.TransactOpts, account_)
}

// AddInviteRankingBlacklist is a paid mutator transaction binding the contract method 0x7dedb468.
//
// Solidity: function addInviteRankingBlacklist(address account_) returns(bool)
func (_Goverment *GovermentTransactorSession) AddInviteRankingBlacklist(account_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.AddInviteRankingBlacklist(&_Goverment.TransactOpts, account_)
}

// BuyVip is a paid mutator transaction binding the contract method 0x30c947b1.
//
// Solidity: function buyVip(address account_, address superior_) returns(uint256)
func (_Goverment *GovermentTransactor) BuyVip(opts *bind.TransactOpts, account_ common.Address, superior_ common.Address) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "buyVip", account_, superior_)
}

// BuyVip is a paid mutator transaction binding the contract method 0x30c947b1.
//
// Solidity: function buyVip(address account_, address superior_) returns(uint256)
func (_Goverment *GovermentSession) BuyVip(account_ common.Address, superior_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.BuyVip(&_Goverment.TransactOpts, account_, superior_)
}

// BuyVip is a paid mutator transaction binding the contract method 0x30c947b1.
//
// Solidity: function buyVip(address account_, address superior_) returns(uint256)
func (_Goverment *GovermentTransactorSession) BuyVip(account_ common.Address, superior_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.BuyVip(&_Goverment.TransactOpts, account_, superior_)
}

// ClaimInviteRankingAward is a paid mutator transaction binding the contract method 0x7f8f5a4b.
//
// Solidity: function claimInviteRankingAward() returns(bool)
func (_Goverment *GovermentTransactor) ClaimInviteRankingAward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "claimInviteRankingAward")
}

// ClaimInviteRankingAward is a paid mutator transaction binding the contract method 0x7f8f5a4b.
//
// Solidity: function claimInviteRankingAward() returns(bool)
func (_Goverment *GovermentSession) ClaimInviteRankingAward() (*types.Transaction, error) {
	return _Goverment.Contract.ClaimInviteRankingAward(&_Goverment.TransactOpts)
}

// ClaimInviteRankingAward is a paid mutator transaction binding the contract method 0x7f8f5a4b.
//
// Solidity: function claimInviteRankingAward() returns(bool)
func (_Goverment *GovermentTransactorSession) ClaimInviteRankingAward() (*types.Transaction, error) {
	return _Goverment.Contract.ClaimInviteRankingAward(&_Goverment.TransactOpts)
}

// EmployExpPack is a paid mutator transaction binding the contract method 0xfd2d3244.
//
// Solidity: function employExpPack(address account_, uint256 amount_) returns(bool)
func (_Goverment *GovermentTransactor) EmployExpPack(opts *bind.TransactOpts, account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "employExpPack", account_, amount_)
}

// EmployExpPack is a paid mutator transaction binding the contract method 0xfd2d3244.
//
// Solidity: function employExpPack(address account_, uint256 amount_) returns(bool)
func (_Goverment *GovermentSession) EmployExpPack(account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.EmployExpPack(&_Goverment.TransactOpts, account_, amount_)
}

// EmployExpPack is a paid mutator transaction binding the contract method 0xfd2d3244.
//
// Solidity: function employExpPack(address account_, uint256 amount_) returns(bool)
func (_Goverment *GovermentTransactorSession) EmployExpPack(account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.EmployExpPack(&_Goverment.TransactOpts, account_, amount_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Goverment *GovermentTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Goverment *GovermentSession) Initialize() (*types.Transaction, error) {
	return _Goverment.Contract.Initialize(&_Goverment.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Goverment *GovermentTransactorSession) Initialize() (*types.Transaction, error) {
	return _Goverment.Contract.Initialize(&_Goverment.TransactOpts)
}

// IssueGoldRankingAward is a paid mutator transaction binding the contract method 0xfbb205f9.
//
// Solidity: function issueGoldRankingAward(address[] recipients_, uint256[] rate_) returns(bool)
func (_Goverment *GovermentTransactor) IssueGoldRankingAward(opts *bind.TransactOpts, recipients_ []common.Address, rate_ []*big.Int) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "issueGoldRankingAward", recipients_, rate_)
}

// IssueGoldRankingAward is a paid mutator transaction binding the contract method 0xfbb205f9.
//
// Solidity: function issueGoldRankingAward(address[] recipients_, uint256[] rate_) returns(bool)
func (_Goverment *GovermentSession) IssueGoldRankingAward(recipients_ []common.Address, rate_ []*big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.IssueGoldRankingAward(&_Goverment.TransactOpts, recipients_, rate_)
}

// IssueGoldRankingAward is a paid mutator transaction binding the contract method 0xfbb205f9.
//
// Solidity: function issueGoldRankingAward(address[] recipients_, uint256[] rate_) returns(bool)
func (_Goverment *GovermentTransactorSession) IssueGoldRankingAward(recipients_ []common.Address, rate_ []*big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.IssueGoldRankingAward(&_Goverment.TransactOpts, recipients_, rate_)
}

// IssueInviteRankingAward is a paid mutator transaction binding the contract method 0xfc36a31e.
//
// Solidity: function issueInviteRankingAward() returns(bool)
func (_Goverment *GovermentTransactor) IssueInviteRankingAward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "issueInviteRankingAward")
}

// IssueInviteRankingAward is a paid mutator transaction binding the contract method 0xfc36a31e.
//
// Solidity: function issueInviteRankingAward() returns(bool)
func (_Goverment *GovermentSession) IssueInviteRankingAward() (*types.Transaction, error) {
	return _Goverment.Contract.IssueInviteRankingAward(&_Goverment.TransactOpts)
}

// IssueInviteRankingAward is a paid mutator transaction binding the contract method 0xfc36a31e.
//
// Solidity: function issueInviteRankingAward() returns(bool)
func (_Goverment *GovermentTransactorSession) IssueInviteRankingAward() (*types.Transaction, error) {
	return _Goverment.Contract.IssueInviteRankingAward(&_Goverment.TransactOpts)
}

// OpenExchange is a paid mutator transaction binding the contract method 0x74d54aa5.
//
// Solidity: function openExchange(bool isOpenExchange_) returns(bool)
func (_Goverment *GovermentTransactor) OpenExchange(opts *bind.TransactOpts, isOpenExchange_ bool) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "openExchange", isOpenExchange_)
}

// OpenExchange is a paid mutator transaction binding the contract method 0x74d54aa5.
//
// Solidity: function openExchange(bool isOpenExchange_) returns(bool)
func (_Goverment *GovermentSession) OpenExchange(isOpenExchange_ bool) (*types.Transaction, error) {
	return _Goverment.Contract.OpenExchange(&_Goverment.TransactOpts, isOpenExchange_)
}

// OpenExchange is a paid mutator transaction binding the contract method 0x74d54aa5.
//
// Solidity: function openExchange(bool isOpenExchange_) returns(bool)
func (_Goverment *GovermentTransactorSession) OpenExchange(isOpenExchange_ bool) (*types.Transaction, error) {
	return _Goverment.Contract.OpenExchange(&_Goverment.TransactOpts, isOpenExchange_)
}

// OpenGoldRanking is a paid mutator transaction binding the contract method 0x43b97874.
//
// Solidity: function openGoldRanking(bool isOpenGoldRanking_) returns(bool)
func (_Goverment *GovermentTransactor) OpenGoldRanking(opts *bind.TransactOpts, isOpenGoldRanking_ bool) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "openGoldRanking", isOpenGoldRanking_)
}

// OpenGoldRanking is a paid mutator transaction binding the contract method 0x43b97874.
//
// Solidity: function openGoldRanking(bool isOpenGoldRanking_) returns(bool)
func (_Goverment *GovermentSession) OpenGoldRanking(isOpenGoldRanking_ bool) (*types.Transaction, error) {
	return _Goverment.Contract.OpenGoldRanking(&_Goverment.TransactOpts, isOpenGoldRanking_)
}

// OpenGoldRanking is a paid mutator transaction binding the contract method 0x43b97874.
//
// Solidity: function openGoldRanking(bool isOpenGoldRanking_) returns(bool)
func (_Goverment *GovermentTransactorSession) OpenGoldRanking(isOpenGoldRanking_ bool) (*types.Transaction, error) {
	return _Goverment.Contract.OpenGoldRanking(&_Goverment.TransactOpts, isOpenGoldRanking_)
}

// OpenInviteRanking is a paid mutator transaction binding the contract method 0xcb55be0b.
//
// Solidity: function openInviteRanking(bool isOpenInviteRanking_) returns(bool)
func (_Goverment *GovermentTransactor) OpenInviteRanking(opts *bind.TransactOpts, isOpenInviteRanking_ bool) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "openInviteRanking", isOpenInviteRanking_)
}

// OpenInviteRanking is a paid mutator transaction binding the contract method 0xcb55be0b.
//
// Solidity: function openInviteRanking(bool isOpenInviteRanking_) returns(bool)
func (_Goverment *GovermentSession) OpenInviteRanking(isOpenInviteRanking_ bool) (*types.Transaction, error) {
	return _Goverment.Contract.OpenInviteRanking(&_Goverment.TransactOpts, isOpenInviteRanking_)
}

// OpenInviteRanking is a paid mutator transaction binding the contract method 0xcb55be0b.
//
// Solidity: function openInviteRanking(bool isOpenInviteRanking_) returns(bool)
func (_Goverment *GovermentTransactorSession) OpenInviteRanking(isOpenInviteRanking_ bool) (*types.Transaction, error) {
	return _Goverment.Contract.OpenInviteRanking(&_Goverment.TransactOpts, isOpenInviteRanking_)
}

// RemoveInviteRankingBlacklist is a paid mutator transaction binding the contract method 0x02766bd2.
//
// Solidity: function removeInviteRankingBlacklist(address account_) returns(bool)
func (_Goverment *GovermentTransactor) RemoveInviteRankingBlacklist(opts *bind.TransactOpts, account_ common.Address) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "removeInviteRankingBlacklist", account_)
}

// RemoveInviteRankingBlacklist is a paid mutator transaction binding the contract method 0x02766bd2.
//
// Solidity: function removeInviteRankingBlacklist(address account_) returns(bool)
func (_Goverment *GovermentSession) RemoveInviteRankingBlacklist(account_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.RemoveInviteRankingBlacklist(&_Goverment.TransactOpts, account_)
}

// RemoveInviteRankingBlacklist is a paid mutator transaction binding the contract method 0x02766bd2.
//
// Solidity: function removeInviteRankingBlacklist(address account_) returns(bool)
func (_Goverment *GovermentTransactorSession) RemoveInviteRankingBlacklist(account_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.RemoveInviteRankingBlacklist(&_Goverment.TransactOpts, account_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Goverment *GovermentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Goverment *GovermentSession) RenounceOwnership() (*types.Transaction, error) {
	return _Goverment.Contract.RenounceOwnership(&_Goverment.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Goverment *GovermentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Goverment.Contract.RenounceOwnership(&_Goverment.TransactOpts)
}

// ResetRankingConfig is a paid mutator transaction binding the contract method 0xf92605e3.
//
// Solidity: function resetRankingConfig() returns(bool)
func (_Goverment *GovermentTransactor) ResetRankingConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "resetRankingConfig")
}

// ResetRankingConfig is a paid mutator transaction binding the contract method 0xf92605e3.
//
// Solidity: function resetRankingConfig() returns(bool)
func (_Goverment *GovermentSession) ResetRankingConfig() (*types.Transaction, error) {
	return _Goverment.Contract.ResetRankingConfig(&_Goverment.TransactOpts)
}

// ResetRankingConfig is a paid mutator transaction binding the contract method 0xf92605e3.
//
// Solidity: function resetRankingConfig() returns(bool)
func (_Goverment *GovermentTransactorSession) ResetRankingConfig() (*types.Transaction, error) {
	return _Goverment.Contract.ResetRankingConfig(&_Goverment.TransactOpts)
}

// SetAdmissionTicket is a paid mutator transaction binding the contract method 0xfa3e54df.
//
// Solidity: function setAdmissionTicket(uint256 admissionTicket_) returns(bool)
func (_Goverment *GovermentTransactor) SetAdmissionTicket(opts *bind.TransactOpts, admissionTicket_ *big.Int) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "setAdmissionTicket", admissionTicket_)
}

// SetAdmissionTicket is a paid mutator transaction binding the contract method 0xfa3e54df.
//
// Solidity: function setAdmissionTicket(uint256 admissionTicket_) returns(bool)
func (_Goverment *GovermentSession) SetAdmissionTicket(admissionTicket_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetAdmissionTicket(&_Goverment.TransactOpts, admissionTicket_)
}

// SetAdmissionTicket is a paid mutator transaction binding the contract method 0xfa3e54df.
//
// Solidity: function setAdmissionTicket(uint256 admissionTicket_) returns(bool)
func (_Goverment *GovermentTransactorSession) SetAdmissionTicket(admissionTicket_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetAdmissionTicket(&_Goverment.TransactOpts, admissionTicket_)
}

// SetDapp is a paid mutator transaction binding the contract method 0x4a02a9f2.
//
// Solidity: function setDapp(address dapp_) returns(bool)
func (_Goverment *GovermentTransactor) SetDapp(opts *bind.TransactOpts, dapp_ common.Address) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "setDapp", dapp_)
}

// SetDapp is a paid mutator transaction binding the contract method 0x4a02a9f2.
//
// Solidity: function setDapp(address dapp_) returns(bool)
func (_Goverment *GovermentSession) SetDapp(dapp_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.SetDapp(&_Goverment.TransactOpts, dapp_)
}

// SetDapp is a paid mutator transaction binding the contract method 0x4a02a9f2.
//
// Solidity: function setDapp(address dapp_) returns(bool)
func (_Goverment *GovermentTransactorSession) SetDapp(dapp_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.SetDapp(&_Goverment.TransactOpts, dapp_)
}

// SetFAPEGiftAmount is a paid mutator transaction binding the contract method 0x56ca454d.
//
// Solidity: function setFAPEGiftAmount(uint256 FAPEGiftAmount_) returns(bool)
func (_Goverment *GovermentTransactor) SetFAPEGiftAmount(opts *bind.TransactOpts, FAPEGiftAmount_ *big.Int) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "setFAPEGiftAmount", FAPEGiftAmount_)
}

// SetFAPEGiftAmount is a paid mutator transaction binding the contract method 0x56ca454d.
//
// Solidity: function setFAPEGiftAmount(uint256 FAPEGiftAmount_) returns(bool)
func (_Goverment *GovermentSession) SetFAPEGiftAmount(FAPEGiftAmount_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetFAPEGiftAmount(&_Goverment.TransactOpts, FAPEGiftAmount_)
}

// SetFAPEGiftAmount is a paid mutator transaction binding the contract method 0x56ca454d.
//
// Solidity: function setFAPEGiftAmount(uint256 FAPEGiftAmount_) returns(bool)
func (_Goverment *GovermentTransactorSession) SetFAPEGiftAmount(FAPEGiftAmount_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetFAPEGiftAmount(&_Goverment.TransactOpts, FAPEGiftAmount_)
}

// SetGameAwardAddr is a paid mutator transaction binding the contract method 0x4949890e.
//
// Solidity: function setGameAwardAddr(address gameAwardAddr_) returns(bool)
func (_Goverment *GovermentTransactor) SetGameAwardAddr(opts *bind.TransactOpts, gameAwardAddr_ common.Address) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "setGameAwardAddr", gameAwardAddr_)
}

// SetGameAwardAddr is a paid mutator transaction binding the contract method 0x4949890e.
//
// Solidity: function setGameAwardAddr(address gameAwardAddr_) returns(bool)
func (_Goverment *GovermentSession) SetGameAwardAddr(gameAwardAddr_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.SetGameAwardAddr(&_Goverment.TransactOpts, gameAwardAddr_)
}

// SetGameAwardAddr is a paid mutator transaction binding the contract method 0x4949890e.
//
// Solidity: function setGameAwardAddr(address gameAwardAddr_) returns(bool)
func (_Goverment *GovermentTransactorSession) SetGameAwardAddr(gameAwardAddr_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.SetGameAwardAddr(&_Goverment.TransactOpts, gameAwardAddr_)
}

// SetGoldRankingTimesnap is a paid mutator transaction binding the contract method 0x706f721c.
//
// Solidity: function setGoldRankingTimesnap(uint256 rankingTimesnap_) returns(bool)
func (_Goverment *GovermentTransactor) SetGoldRankingTimesnap(opts *bind.TransactOpts, rankingTimesnap_ *big.Int) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "setGoldRankingTimesnap", rankingTimesnap_)
}

// SetGoldRankingTimesnap is a paid mutator transaction binding the contract method 0x706f721c.
//
// Solidity: function setGoldRankingTimesnap(uint256 rankingTimesnap_) returns(bool)
func (_Goverment *GovermentSession) SetGoldRankingTimesnap(rankingTimesnap_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetGoldRankingTimesnap(&_Goverment.TransactOpts, rankingTimesnap_)
}

// SetGoldRankingTimesnap is a paid mutator transaction binding the contract method 0x706f721c.
//
// Solidity: function setGoldRankingTimesnap(uint256 rankingTimesnap_) returns(bool)
func (_Goverment *GovermentTransactorSession) SetGoldRankingTimesnap(rankingTimesnap_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetGoldRankingTimesnap(&_Goverment.TransactOpts, rankingTimesnap_)
}

// SetGoldRankingTotalItems is a paid mutator transaction binding the contract method 0x932a264a.
//
// Solidity: function setGoldRankingTotalItems(uint256 rankingTotalItems_) returns(bool)
func (_Goverment *GovermentTransactor) SetGoldRankingTotalItems(opts *bind.TransactOpts, rankingTotalItems_ *big.Int) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "setGoldRankingTotalItems", rankingTotalItems_)
}

// SetGoldRankingTotalItems is a paid mutator transaction binding the contract method 0x932a264a.
//
// Solidity: function setGoldRankingTotalItems(uint256 rankingTotalItems_) returns(bool)
func (_Goverment *GovermentSession) SetGoldRankingTotalItems(rankingTotalItems_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetGoldRankingTotalItems(&_Goverment.TransactOpts, rankingTotalItems_)
}

// SetGoldRankingTotalItems is a paid mutator transaction binding the contract method 0x932a264a.
//
// Solidity: function setGoldRankingTotalItems(uint256 rankingTotalItems_) returns(bool)
func (_Goverment *GovermentTransactorSession) SetGoldRankingTotalItems(rankingTotalItems_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetGoldRankingTotalItems(&_Goverment.TransactOpts, rankingTotalItems_)
}

// SetInviteRankingTimesnap is a paid mutator transaction binding the contract method 0x0e0633ab.
//
// Solidity: function setInviteRankingTimesnap(uint256 rankingTimesnap_) returns(bool)
func (_Goverment *GovermentTransactor) SetInviteRankingTimesnap(opts *bind.TransactOpts, rankingTimesnap_ *big.Int) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "setInviteRankingTimesnap", rankingTimesnap_)
}

// SetInviteRankingTimesnap is a paid mutator transaction binding the contract method 0x0e0633ab.
//
// Solidity: function setInviteRankingTimesnap(uint256 rankingTimesnap_) returns(bool)
func (_Goverment *GovermentSession) SetInviteRankingTimesnap(rankingTimesnap_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetInviteRankingTimesnap(&_Goverment.TransactOpts, rankingTimesnap_)
}

// SetInviteRankingTimesnap is a paid mutator transaction binding the contract method 0x0e0633ab.
//
// Solidity: function setInviteRankingTimesnap(uint256 rankingTimesnap_) returns(bool)
func (_Goverment *GovermentTransactorSession) SetInviteRankingTimesnap(rankingTimesnap_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetInviteRankingTimesnap(&_Goverment.TransactOpts, rankingTimesnap_)
}

// SetInviteRankingTotalItems is a paid mutator transaction binding the contract method 0xe7201ade.
//
// Solidity: function setInviteRankingTotalItems(uint256 rankingTotalItems_) returns(bool)
func (_Goverment *GovermentTransactor) SetInviteRankingTotalItems(opts *bind.TransactOpts, rankingTotalItems_ *big.Int) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "setInviteRankingTotalItems", rankingTotalItems_)
}

// SetInviteRankingTotalItems is a paid mutator transaction binding the contract method 0xe7201ade.
//
// Solidity: function setInviteRankingTotalItems(uint256 rankingTotalItems_) returns(bool)
func (_Goverment *GovermentSession) SetInviteRankingTotalItems(rankingTotalItems_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetInviteRankingTotalItems(&_Goverment.TransactOpts, rankingTotalItems_)
}

// SetInviteRankingTotalItems is a paid mutator transaction binding the contract method 0xe7201ade.
//
// Solidity: function setInviteRankingTotalItems(uint256 rankingTotalItems_) returns(bool)
func (_Goverment *GovermentTransactorSession) SetInviteRankingTotalItems(rankingTotalItems_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetInviteRankingTotalItems(&_Goverment.TransactOpts, rankingTotalItems_)
}

// SetSwapTimesnap is a paid mutator transaction binding the contract method 0x52ae7a50.
//
// Solidity: function setSwapTimesnap(uint256 swapTimesnap_) returns(bool)
func (_Goverment *GovermentTransactor) SetSwapTimesnap(opts *bind.TransactOpts, swapTimesnap_ *big.Int) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "setSwapTimesnap", swapTimesnap_)
}

// SetSwapTimesnap is a paid mutator transaction binding the contract method 0x52ae7a50.
//
// Solidity: function setSwapTimesnap(uint256 swapTimesnap_) returns(bool)
func (_Goverment *GovermentSession) SetSwapTimesnap(swapTimesnap_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetSwapTimesnap(&_Goverment.TransactOpts, swapTimesnap_)
}

// SetSwapTimesnap is a paid mutator transaction binding the contract method 0x52ae7a50.
//
// Solidity: function setSwapTimesnap(uint256 swapTimesnap_) returns(bool)
func (_Goverment *GovermentTransactorSession) SetSwapTimesnap(swapTimesnap_ *big.Int) (*types.Transaction, error) {
	return _Goverment.Contract.SetSwapTimesnap(&_Goverment.TransactOpts, swapTimesnap_)
}

// SetTicketNFT is a paid mutator transaction binding the contract method 0x6ade389b.
//
// Solidity: function setTicketNFT(address ticketNFT_) returns(bool)
func (_Goverment *GovermentTransactor) SetTicketNFT(opts *bind.TransactOpts, ticketNFT_ common.Address) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "setTicketNFT", ticketNFT_)
}

// SetTicketNFT is a paid mutator transaction binding the contract method 0x6ade389b.
//
// Solidity: function setTicketNFT(address ticketNFT_) returns(bool)
func (_Goverment *GovermentSession) SetTicketNFT(ticketNFT_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.SetTicketNFT(&_Goverment.TransactOpts, ticketNFT_)
}

// SetTicketNFT is a paid mutator transaction binding the contract method 0x6ade389b.
//
// Solidity: function setTicketNFT(address ticketNFT_) returns(bool)
func (_Goverment *GovermentTransactorSession) SetTicketNFT(ticketNFT_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.SetTicketNFT(&_Goverment.TransactOpts, ticketNFT_)
}

// SetTokenFAPE is a paid mutator transaction binding the contract method 0x42ead5a7.
//
// Solidity: function setTokenFAPE(address tokenFAPE_) returns(bool)
func (_Goverment *GovermentTransactor) SetTokenFAPE(opts *bind.TransactOpts, tokenFAPE_ common.Address) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "setTokenFAPE", tokenFAPE_)
}

// SetTokenFAPE is a paid mutator transaction binding the contract method 0x42ead5a7.
//
// Solidity: function setTokenFAPE(address tokenFAPE_) returns(bool)
func (_Goverment *GovermentSession) SetTokenFAPE(tokenFAPE_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.SetTokenFAPE(&_Goverment.TransactOpts, tokenFAPE_)
}

// SetTokenFAPE is a paid mutator transaction binding the contract method 0x42ead5a7.
//
// Solidity: function setTokenFAPE(address tokenFAPE_) returns(bool)
func (_Goverment *GovermentTransactorSession) SetTokenFAPE(tokenFAPE_ common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.SetTokenFAPE(&_Goverment.TransactOpts, tokenFAPE_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Goverment *GovermentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Goverment.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Goverment *GovermentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.TransferOwnership(&_Goverment.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Goverment *GovermentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Goverment.Contract.TransferOwnership(&_Goverment.TransactOpts, newOwner)
}

// GovermentBindSuperiorIterator is returned from FilterBindSuperior and is used to iterate over the raw logs and unpacked data for BindSuperior events raised by the Goverment contract.
type GovermentBindSuperiorIterator struct {
	Event *GovermentBindSuperior // Event containing the contract specifics and raw log

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
func (it *GovermentBindSuperiorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovermentBindSuperior)
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
		it.Event = new(GovermentBindSuperior)
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
func (it *GovermentBindSuperiorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovermentBindSuperiorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovermentBindSuperior represents a BindSuperior event raised by the Goverment contract.
type GovermentBindSuperior struct {
	Account  common.Address
	Superior common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBindSuperior is a free log retrieval operation binding the contract event 0x5ce9b5b2331d888fe42214d85e899e9a2cd1a6fcfb8e075e26d4ebc190064414.
//
// Solidity: event BindSuperior(address indexed account_, address indexed superior_)
func (_Goverment *GovermentFilterer) FilterBindSuperior(opts *bind.FilterOpts, account_ []common.Address, superior_ []common.Address) (*GovermentBindSuperiorIterator, error) {

	var account_Rule []interface{}
	for _, account_Item := range account_ {
		account_Rule = append(account_Rule, account_Item)
	}
	var superior_Rule []interface{}
	for _, superior_Item := range superior_ {
		superior_Rule = append(superior_Rule, superior_Item)
	}

	logs, sub, err := _Goverment.contract.FilterLogs(opts, "BindSuperior", account_Rule, superior_Rule)
	if err != nil {
		return nil, err
	}
	return &GovermentBindSuperiorIterator{contract: _Goverment.contract, event: "BindSuperior", logs: logs, sub: sub}, nil
}

// WatchBindSuperior is a free log subscription operation binding the contract event 0x5ce9b5b2331d888fe42214d85e899e9a2cd1a6fcfb8e075e26d4ebc190064414.
//
// Solidity: event BindSuperior(address indexed account_, address indexed superior_)
func (_Goverment *GovermentFilterer) WatchBindSuperior(opts *bind.WatchOpts, sink chan<- *GovermentBindSuperior, account_ []common.Address, superior_ []common.Address) (event.Subscription, error) {

	var account_Rule []interface{}
	for _, account_Item := range account_ {
		account_Rule = append(account_Rule, account_Item)
	}
	var superior_Rule []interface{}
	for _, superior_Item := range superior_ {
		superior_Rule = append(superior_Rule, superior_Item)
	}

	logs, sub, err := _Goverment.contract.WatchLogs(opts, "BindSuperior", account_Rule, superior_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovermentBindSuperior)
				if err := _Goverment.contract.UnpackLog(event, "BindSuperior", log); err != nil {
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

// ParseBindSuperior is a log parse operation binding the contract event 0x5ce9b5b2331d888fe42214d85e899e9a2cd1a6fcfb8e075e26d4ebc190064414.
//
// Solidity: event BindSuperior(address indexed account_, address indexed superior_)
func (_Goverment *GovermentFilterer) ParseBindSuperior(log types.Log) (*GovermentBindSuperior, error) {
	event := new(GovermentBindSuperior)
	if err := _Goverment.contract.UnpackLog(event, "BindSuperior", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovermentBuyVipIterator is returned from FilterBuyVip and is used to iterate over the raw logs and unpacked data for BuyVip events raised by the Goverment contract.
type GovermentBuyVipIterator struct {
	Event *GovermentBuyVip // Event containing the contract specifics and raw log

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
func (it *GovermentBuyVipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovermentBuyVip)
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
		it.Event = new(GovermentBuyVip)
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
func (it *GovermentBuyVipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovermentBuyVipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovermentBuyVip represents a BuyVip event raised by the Goverment contract.
type GovermentBuyVip struct {
	Account  common.Address
	Superior common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBuyVip is a free log retrieval operation binding the contract event 0x7253d156974fc04c2d24feb1b925ae89963ace56ccb9efde4d0f706a98ddd0a6.
//
// Solidity: event BuyVip(address indexed account_, address indexed superior_)
func (_Goverment *GovermentFilterer) FilterBuyVip(opts *bind.FilterOpts, account_ []common.Address, superior_ []common.Address) (*GovermentBuyVipIterator, error) {

	var account_Rule []interface{}
	for _, account_Item := range account_ {
		account_Rule = append(account_Rule, account_Item)
	}
	var superior_Rule []interface{}
	for _, superior_Item := range superior_ {
		superior_Rule = append(superior_Rule, superior_Item)
	}

	logs, sub, err := _Goverment.contract.FilterLogs(opts, "BuyVip", account_Rule, superior_Rule)
	if err != nil {
		return nil, err
	}
	return &GovermentBuyVipIterator{contract: _Goverment.contract, event: "BuyVip", logs: logs, sub: sub}, nil
}

// WatchBuyVip is a free log subscription operation binding the contract event 0x7253d156974fc04c2d24feb1b925ae89963ace56ccb9efde4d0f706a98ddd0a6.
//
// Solidity: event BuyVip(address indexed account_, address indexed superior_)
func (_Goverment *GovermentFilterer) WatchBuyVip(opts *bind.WatchOpts, sink chan<- *GovermentBuyVip, account_ []common.Address, superior_ []common.Address) (event.Subscription, error) {

	var account_Rule []interface{}
	for _, account_Item := range account_ {
		account_Rule = append(account_Rule, account_Item)
	}
	var superior_Rule []interface{}
	for _, superior_Item := range superior_ {
		superior_Rule = append(superior_Rule, superior_Item)
	}

	logs, sub, err := _Goverment.contract.WatchLogs(opts, "BuyVip", account_Rule, superior_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovermentBuyVip)
				if err := _Goverment.contract.UnpackLog(event, "BuyVip", log); err != nil {
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

// ParseBuyVip is a log parse operation binding the contract event 0x7253d156974fc04c2d24feb1b925ae89963ace56ccb9efde4d0f706a98ddd0a6.
//
// Solidity: event BuyVip(address indexed account_, address indexed superior_)
func (_Goverment *GovermentFilterer) ParseBuyVip(log types.Log) (*GovermentBuyVip, error) {
	event := new(GovermentBuyVip)
	if err := _Goverment.contract.UnpackLog(event, "BuyVip", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovermentFAPEChainToGameEventIterator is returned from FilterFAPEChainToGameEvent and is used to iterate over the raw logs and unpacked data for FAPEChainToGameEvent events raised by the Goverment contract.
type GovermentFAPEChainToGameEventIterator struct {
	Event *GovermentFAPEChainToGameEvent // Event containing the contract specifics and raw log

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
func (it *GovermentFAPEChainToGameEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovermentFAPEChainToGameEvent)
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
		it.Event = new(GovermentFAPEChainToGameEvent)
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
func (it *GovermentFAPEChainToGameEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovermentFAPEChainToGameEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovermentFAPEChainToGameEvent represents a FAPEChainToGameEvent event raised by the Goverment contract.
type GovermentFAPEChainToGameEvent struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFAPEChainToGameEvent is a free log retrieval operation binding the contract event 0xdf76a1c38ca5f6615a96601f61d8df9ee900ab99d80dce76a4d6a1f96ad75f50.
//
// Solidity: event FAPEChainToGameEvent(address indexed account_, uint256 amount_)
func (_Goverment *GovermentFilterer) FilterFAPEChainToGameEvent(opts *bind.FilterOpts, account_ []common.Address) (*GovermentFAPEChainToGameEventIterator, error) {

	var account_Rule []interface{}
	for _, account_Item := range account_ {
		account_Rule = append(account_Rule, account_Item)
	}

	logs, sub, err := _Goverment.contract.FilterLogs(opts, "FAPEChainToGameEvent", account_Rule)
	if err != nil {
		return nil, err
	}
	return &GovermentFAPEChainToGameEventIterator{contract: _Goverment.contract, event: "FAPEChainToGameEvent", logs: logs, sub: sub}, nil
}

// WatchFAPEChainToGameEvent is a free log subscription operation binding the contract event 0xdf76a1c38ca5f6615a96601f61d8df9ee900ab99d80dce76a4d6a1f96ad75f50.
//
// Solidity: event FAPEChainToGameEvent(address indexed account_, uint256 amount_)
func (_Goverment *GovermentFilterer) WatchFAPEChainToGameEvent(opts *bind.WatchOpts, sink chan<- *GovermentFAPEChainToGameEvent, account_ []common.Address) (event.Subscription, error) {

	var account_Rule []interface{}
	for _, account_Item := range account_ {
		account_Rule = append(account_Rule, account_Item)
	}

	logs, sub, err := _Goverment.contract.WatchLogs(opts, "FAPEChainToGameEvent", account_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovermentFAPEChainToGameEvent)
				if err := _Goverment.contract.UnpackLog(event, "FAPEChainToGameEvent", log); err != nil {
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

// ParseFAPEChainToGameEvent is a log parse operation binding the contract event 0xdf76a1c38ca5f6615a96601f61d8df9ee900ab99d80dce76a4d6a1f96ad75f50.
//
// Solidity: event FAPEChainToGameEvent(address indexed account_, uint256 amount_)
func (_Goverment *GovermentFilterer) ParseFAPEChainToGameEvent(log types.Log) (*GovermentFAPEChainToGameEvent, error) {
	event := new(GovermentFAPEChainToGameEvent)
	if err := _Goverment.contract.UnpackLog(event, "FAPEChainToGameEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovermentFAPEGameToChainEventIterator is returned from FilterFAPEGameToChainEvent and is used to iterate over the raw logs and unpacked data for FAPEGameToChainEvent events raised by the Goverment contract.
type GovermentFAPEGameToChainEventIterator struct {
	Event *GovermentFAPEGameToChainEvent // Event containing the contract specifics and raw log

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
func (it *GovermentFAPEGameToChainEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovermentFAPEGameToChainEvent)
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
		it.Event = new(GovermentFAPEGameToChainEvent)
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
func (it *GovermentFAPEGameToChainEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovermentFAPEGameToChainEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovermentFAPEGameToChainEvent represents a FAPEGameToChainEvent event raised by the Goverment contract.
type GovermentFAPEGameToChainEvent struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFAPEGameToChainEvent is a free log retrieval operation binding the contract event 0x5d8f8b30dfdc79f8bf9c04d973bed21d3b44c646a206ebf0477c58e99cbe6634.
//
// Solidity: event FAPEGameToChainEvent(address indexed account_, uint256 amount_)
func (_Goverment *GovermentFilterer) FilterFAPEGameToChainEvent(opts *bind.FilterOpts, account_ []common.Address) (*GovermentFAPEGameToChainEventIterator, error) {

	var account_Rule []interface{}
	for _, account_Item := range account_ {
		account_Rule = append(account_Rule, account_Item)
	}

	logs, sub, err := _Goverment.contract.FilterLogs(opts, "FAPEGameToChainEvent", account_Rule)
	if err != nil {
		return nil, err
	}
	return &GovermentFAPEGameToChainEventIterator{contract: _Goverment.contract, event: "FAPEGameToChainEvent", logs: logs, sub: sub}, nil
}

// WatchFAPEGameToChainEvent is a free log subscription operation binding the contract event 0x5d8f8b30dfdc79f8bf9c04d973bed21d3b44c646a206ebf0477c58e99cbe6634.
//
// Solidity: event FAPEGameToChainEvent(address indexed account_, uint256 amount_)
func (_Goverment *GovermentFilterer) WatchFAPEGameToChainEvent(opts *bind.WatchOpts, sink chan<- *GovermentFAPEGameToChainEvent, account_ []common.Address) (event.Subscription, error) {

	var account_Rule []interface{}
	for _, account_Item := range account_ {
		account_Rule = append(account_Rule, account_Item)
	}

	logs, sub, err := _Goverment.contract.WatchLogs(opts, "FAPEGameToChainEvent", account_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovermentFAPEGameToChainEvent)
				if err := _Goverment.contract.UnpackLog(event, "FAPEGameToChainEvent", log); err != nil {
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

// ParseFAPEGameToChainEvent is a log parse operation binding the contract event 0x5d8f8b30dfdc79f8bf9c04d973bed21d3b44c646a206ebf0477c58e99cbe6634.
//
// Solidity: event FAPEGameToChainEvent(address indexed account_, uint256 amount_)
func (_Goverment *GovermentFilterer) ParseFAPEGameToChainEvent(log types.Log) (*GovermentFAPEGameToChainEvent, error) {
	event := new(GovermentFAPEGameToChainEvent)
	if err := _Goverment.contract.UnpackLog(event, "FAPEGameToChainEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovermentInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Goverment contract.
type GovermentInitializedIterator struct {
	Event *GovermentInitialized // Event containing the contract specifics and raw log

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
func (it *GovermentInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovermentInitialized)
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
		it.Event = new(GovermentInitialized)
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
func (it *GovermentInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovermentInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovermentInitialized represents a Initialized event raised by the Goverment contract.
type GovermentInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Goverment *GovermentFilterer) FilterInitialized(opts *bind.FilterOpts) (*GovermentInitializedIterator, error) {

	logs, sub, err := _Goverment.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GovermentInitializedIterator{contract: _Goverment.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Goverment *GovermentFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GovermentInitialized) (event.Subscription, error) {

	logs, sub, err := _Goverment.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovermentInitialized)
				if err := _Goverment.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Goverment *GovermentFilterer) ParseInitialized(log types.Log) (*GovermentInitialized, error) {
	event := new(GovermentInitialized)
	if err := _Goverment.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovermentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Goverment contract.
type GovermentOwnershipTransferredIterator struct {
	Event *GovermentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GovermentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovermentOwnershipTransferred)
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
		it.Event = new(GovermentOwnershipTransferred)
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
func (it *GovermentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovermentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovermentOwnershipTransferred represents a OwnershipTransferred event raised by the Goverment contract.
type GovermentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Goverment *GovermentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GovermentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Goverment.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GovermentOwnershipTransferredIterator{contract: _Goverment.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Goverment *GovermentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovermentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Goverment.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovermentOwnershipTransferred)
				if err := _Goverment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Goverment *GovermentFilterer) ParseOwnershipTransferred(log types.Log) (*GovermentOwnershipTransferred, error) {
	event := new(GovermentOwnershipTransferred)
	if err := _Goverment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
