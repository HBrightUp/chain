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
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/go-sql-driver/mysql"
)

func transferFrom(client *ethclient.Client, hexKey string, to common.Address, amount string) (string, error) {
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

	tokenAddress := common.HexToAddress("0x65867992a117586aCd311329D9E17410bBE60CCd")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	//fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(to.Bytes(), 32)
	//fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	value := big.NewInt(0)

	amount_value := new(big.Int)
	amount_value.SetString(amount, 10)
	paddedAmount := common.LeftPadBytes(amount_value.Bytes(), 32)
	gasLimit := uint64(210000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

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
	//client, err := ethclient.Dial("http://192.168.1.165:8545")
	client, err := ethclient.Dial("http://192.168.1.11:18545")
	if err != nil {
		log.Fatal(err)
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
		re := regexp.MustCompile("0x[0-9a-fA-F]{40}")
		if is_valid := re.MatchString(tolower); !is_valid {
			log.Fatal("invalid address: ", to)
			continue
		}

		toAddress := common.HexToAddress(tolower)
		amount := string(row[1])
		hexKey := privateKeys[r.Intn(keyLens)]
		hash, err := transferFrom(client, hexKey, toAddress, amount)
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
