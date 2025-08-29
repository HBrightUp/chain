package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func addAddress(client *ethclient.Client, hexKey string, airdrop *Airdrop) (string, error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(21000)
	auth.GasPrice = gasPrice

	amount, _ := new(big.Int).SetString("10000000000", 10)
	to := common.HexToAddress("")
	tx, err := airdrop.Add(auth, to, amount)
	if err != nil {
		return "", err
	}
	fmt.Printf("add tx sent: %s", tx.Hash().Hex())
	fmt.Println()

	return tx.Hash().Hex(), nil
}

func notify(client *ethclient.Client, hexKey string, airdrop *Airdrop) (string, error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(21000)
	auth.GasPrice = gasPrice

	tx, err := airdrop.NotifyAirdropAmounts(auth)
	if err != nil {
		return "", err
	}

	//	fmt.Printf("airdrop tx sent: %s", tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func main() {
	//client, err := ethclient.Dial("http://http-testnet.aitd.io")
	client, err := ethclient.Dial("http://18.141.247.72:8545")
	if err != nil {
		log.Fatalf("Failed to connect to ethereum: %v", err)
		return
	}
	defer client.Close()

	f, err := excelize.OpenFile("config/config.xlsx")
	if err != nil {
		log.Fatal(err)
		return
	}

	privateKey, err := f.GetCellValue("Sheet2", "A2")
	if err != nil {
		log.Fatal(err)
		return
	}

	airdropIns, err := NewAirdrop(common.HexToAddress("0xAbd3Fd6BE0C595AEA1B072faCd39Fb841DbC7891"), client)
	if err != nil {
		log.Fatal(err)
		return
	}

	addAddress(client, privateKey, airdropIns)

	tx, err := notify(client, privateKey, airdropIns)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("airdrop tx sent: %s", tx)
}
