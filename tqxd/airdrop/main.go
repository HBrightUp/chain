package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/go-sql-driver/mysql"
)

func transfer(client *ethclient.Client, hexKey string, to common.Address, amount string) (string, error) {
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

	value := new(big.Int)
	value.SetString(amount, 10)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	var data []byte
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	fmt.Printf("%s", signedTx.Hash().Hex())
	fmt.Println()

	return signedTx.Hash().Hex(), nil
}

func main() {
	//client, err := ethclient.Dial("http://http-testnet.aitd.io")
	client, err := ethclient.Dial("http://192.168.1.165:8545")

	if err != nil {
		log.Fatalf("Failed to connect to ethereum: %v", err)
		return
	}
	defer client.Close()

	if err := initDB(); err != nil {
		return
	}

	f, err := excelize.OpenFile("config/config.xlsx")
	if err != nil {
		log.Fatal(err)
		return
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
		return
	}

	var privateKeys []string
	for i, row := range rows {
		if i == 0 {
			continue
		}
		for _, colCell := range row {
			privateKeys = append(privateKeys, colCell)
			//fmt.Print(colCell, "\t")
		}
	}

	keyLens := len(privateKeys)
	if keyLens == 0 {
		log.Fatal("key length is zero")
		return
	}

	f1, err := excelize.OpenFile("config/airdrop.xlsx")
	if err != nil {
		log.Fatal(err)
		return
	}

	rows, err = f1.GetRows("Sheet1")
	if err != nil {
		log.Fatalln(err)
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i, row := range rows {
		if i == 0 {
			continue
		}
		to := row[0]
		tolower := strings.ToLower(to)
		re := regexp.MustCompile("0x[0-9a-fA-F]{40}$")
		if is_valid := re.MatchString(tolower); !is_valid {
			log.Fatal("invalid address: ", to)
			continue
		}

		toAddress := common.HexToAddress(tolower)
		amount := string(row[1])
		hexKey := privateKeys[r.Intn(keyLens)]
		hash, err := transfer(client, hexKey, toAddress, amount)
		if err != nil {
			log.Fatal(err)
			continue
		}

		index := fmt.Sprintf("C%d", i+1)
		f1.SetCellValue("Sheet1", index, hash)

		_, err = db.Exec("insert into t_airdrop(address,hash,amount) values(?,?,?)", to, hash, amount)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Duration(10 * time.Millisecond))
	}

	if err := f1.Save(); err != nil {
		fmt.Println(err)
	}
	db.Close()
}
