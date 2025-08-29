package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
	"time"
)

func TestFloatToStr(t *testing.T) {
	str := "0.001"
	f, err := strToFloat(str)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("f: ", f)
}

func TestStrToBigInt(t *testing.T) {
	str := "1"

	bigInt, ok := new(big.Int).SetString(str, 10)
	if !ok {
		t.Errorf("big.int conv failed")
		return
	}
	fmt.Println("bigInt: ", bigInt)
}


func TestContract(t *testing.T) {
	// Create an IPC based RPC connection to a remote node
	conn1, err := ethclient.Dial("http://http-testnet.aitd.io")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	inviteStake, err := NewInviteStake(common.HexToAddress("0xdFa3390B6a66Ea94eB2bf29eb2a0f06Cf7e74333"), conn1)
	if err != nil {
		t.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	blockNumber, err := inviteStake.GetBlockNumber(nil)
	if err != nil {
		t.Fatalf("Failed to retrieve tblock number: %v", blockNumber)
	}
	fmt.Println("block number: ", blockNumber)

	nonce, err := conn1.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		t.Fatalf(err.Error())
	}

	gasPrice, err := conn1.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatalf(err.Error())
	}

	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(100000)
	auth.GasPrice = gasPrice

	amount, _ := new(big.Int).SetString("1001000000000000000", 10)

	var batch []common.Address
	batch = append(batch, testAddr)

	tx, err := inviteStake.Airdrop(auth, batch, amount)
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println("hash: ", tx.Hash())
	time.Sleep(time.Duration(5) * time.Second)
	receipt, err := conn1.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf(err.Error())
	}
	if receipt == nil {
		t.Fatalf("receipt == nil")
	}
	fmt.Println("status: ", receipt.Status)
}


