package main

import (
	"context"
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//client, err := ethclient.Dial("http://http-testnet.aitd.io")
	//client, err := ethclient.Dial("http://192.168.1.165:8545")
	client, err := ethclient.Dial("http://192.168.1.11:18545")
	if err != nil {
		log.Fatalf("Failed to connect to ethereum: %v", err)
		return
	}
	defer client.Close()

	if err := initDB(); err != nil {
		return
	}

	f, err := excelize.OpenFile("config/airdrop.xlsx")
	if err != nil {
		log.Fatal(err)
		return
	}

	cols, err := f.GetCols("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, col := range cols[2] {
		if i == 0 {
			continue
		}
		fmt.Println(col)

		index := fmt.Sprintf("D%d", i+1)
		status := 0
		sqlStr := "update t_airdrop set status=? where hash=?"
		txHash := common.HexToHash(col)
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			f.SetCellValue("Sheet1", index, "failed")

			_, update_err := db.Exec(sqlStr, 2, col)
			if update_err != nil {
				fmt.Println("新增数据错误", update_err, col)
			}

			continue
		}

		if receipt.Status == 1 {
			f.SetCellValue("Sheet1", index, "success")
			status = 1
		} else {
			f.SetCellValue("Sheet1", index, "failed")
			status = 2
		}

		_, update_err := db.Exec(sqlStr, status, col)
		if update_err != nil {
			fmt.Println("新增数据错误", update_err, col)
		}
	}

	if err := f.Save(); err != nil {
		fmt.Println(err)
	}

	db.Close()
}
