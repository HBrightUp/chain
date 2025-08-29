package main

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const BASE = 1e18

var (
	conn *ethclient.Client = nil
	inviteStake *InviteStake = nil
	db *sql.DB = nil
)

var (
	airdropUpperLimit string
	airdropLowerLimit string
)

var (
	key *ecdsa.PrivateKey = nil
	fromAddress common.Address
	testAddr = common.HexToAddress("0xee31e38007D819E00a386fa4308F42c13871D55D")
	serverPort string
	gasLimit uint64
)

type AirdropLimit struct {
	Upper string `form:"upper" json:"upper" binding:"required"`
	Lower string `form:"lower" json:"lower" binding:"required"`
}

type AirdropParameter struct {
	Address string `form:"address" json:"address" binding:"required"`
	Value string `form:"value" json:"value" binding:"required"`
}

type TxHash struct {
	Hash string `form:"hash" json:"hash" binding:"required"`
}

type SelectAirdropAddress struct {
	toAddress string
	hash *string
	txStatus uint
	count uint
	airdropStatus uint
	airdropTime *int64
}

type SelectAirdropHash struct {
	hash *string
	txStatus uint
	airdropStatus uint
	airdropTime *int64
}

func init() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("FATAL - ")

	cfg, err := loadIni()
	if err != nil {
		log.Println("load ini file failed, ", err.Error())
		os.Exit(1)
	}

	logFile := cfg.Section("log").Key("logFile").String()
	if logFile == "" {
		log.Println("the log file path is empty")
		os.Exit(1)
	}

	logF, err := os.OpenFile(logFile, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Println("open log file failed, ", err.Error())
		os.Exit(1)
	}

	err1 := connectDB(cfg)
	if err1 != nil {
		log.Println("connect mysql db failed, ", err1.Error())
		os.Exit(1)
	}

	rawUrl := cfg.Section("blockchain").Key("ip").String()
	var connErr error
	conn, connErr = ethclient.Dial(rawUrl)
	if connErr != nil {
		log.Println("connect blockchain failed, ", connErr.Error())
		os.Exit(1)
	}
	var gasLimitErr error
	gasLimit, gasLimitErr = cfg.Section("blockchain").Key("gasLimit").Uint64()
	if gasLimitErr != nil {
		log.Println("connect blockchain failed, ", gasLimitErr.Error())
		os.Exit(1)
	}

	if gasLimit < 21000 {
		log.Println("connect blockchain failed, ", "gasLimit < 21000, gas limit is too small")
		os.Exit(1)
	}

	s := cfg.Section("airdrop").Key("contractAddress").String()
	_, vfErr := verificationAddress(s)
	if vfErr != nil {
		log.Println("invalid contract address")
		os.Exit(1)
	}
	var inviteStakeErr error
	inviteStake, inviteStakeErr = NewInviteStake(common.HexToAddress(s), conn)
	if inviteStakeErr != nil {
		log.Println(inviteStakeErr.Error())
		os.Exit(1)
	}

	k := cfg.Section("airdrop").Key("privateKey").String()
	if strings.Contains(k, "0x") || strings.Contains(k, "0X") {
		k = k[2:]
	}
	key , _ = crypto.HexToECDSA(k)
	if key == nil {
		log.Println("Bad airdrop private key")
		os.Exit(1)
	}
	fromAddress = crypto.PubkeyToAddress(key.PublicKey)

	if fromAddress.String() == "" {
		log.Println("the from address is empty")
		os.Exit(1)
	}

	log.SetOutput(logF)
 }

func loadIni() (*ini.File, error){
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func connectDB(cfg *ini.File) error {
	if cfg == nil {
		log.Fatal("Parameter error, cfg == nil")
		return fmt.Errorf("%s", "Parameter error, cfg == nil")
	}
	ip := cfg.Section("mysql").Key("ip").String()
	port := cfg.Section("mysql").Key("port").String()
	user := cfg.Section("mysql").Key("user").String()
	password := cfg.Section("mysql").Key("password").String()
	database := cfg.Section("mysql").Key("database").String()

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{user, ":", password, "@tcp(",ip, ":", port, ")/", database, "?charset=utf8&parseTime=true"}, "")

	var err error
	db, err = sql.Open("mysql", path)
	if err != nil {
		return err
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//验证连接
	if err := db.Ping(); err != nil{
		log.Fatal("open database fail, ", err.Error())
		return err
	}
	return nil
}

func strToFloat(str string) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func verificationAddress(addr string) (string, error) {
	toLowerAddress := addr
	if strings.Contains(toLowerAddress, "0x") || strings.Contains(toLowerAddress, "0X") {
		if strings.Contains(toLowerAddress, "0X") {
			toLowerAddress = toLowerAddress[2:]
			toLowerAddress = "0x" + toLowerAddress
		}
	} else {
		toLowerAddress = "0x" + toLowerAddress
	}

	toLowerAddress = strings.ToLower(toLowerAddress)
	re := regexp.MustCompile("0x[0-9a-fA-F]{40}$")
	if isValid := re.MatchString(toLowerAddress); !isValid {
		return "", fmt.Errorf("%s", "invalid address")
	}
	return toLowerAddress, nil
}

func verificationHash(hash string) string {
	toLowerHash := hash
	if strings.Contains(toLowerHash, "0x") || strings.Contains(toLowerHash, "0X") {
		if strings.Contains(toLowerHash, "0X") {
			toLowerHash = toLowerHash[2:]
			toLowerHash = "0x" + toLowerHash
		}
	} else {
		toLowerHash = "0x" + toLowerHash
	}
	return toLowerHash
}

func strToBigInt(s string) (*big.Int, error){
	f, err := strToFloat(s)
	if err != nil {
		return nil, err
	}

	valueWei, ok := new(big.Int).SetString(fmt.Sprintf("%.0f", f * BASE), 10)
	if !ok {
		return nil, fmt.Errorf("%s", "float to bigInt failed!")
	}
	return valueWei, nil
}

func floatToBigInt(f float64) (*big.Int, error) {
	valueWei, ok := new(big.Int).SetString(fmt.Sprintf("%.0f", f * BASE), 10)
	if !ok {
		return nil, fmt.Errorf("%s", "float to bigInt failed!")
	}
	return valueWei, nil
}

func getUnixTimestamp() int64 {
	timeObj := time.Now()
	return timeObj.Unix()
}

func getReceipt(hash string) (*types.Receipt, error){
	receipt, err := conn.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func airdropLog(prefix string, v ...interface{}) {
	log.SetPrefix(prefix + " - ")
	log.Println(v)
}

func insertIntoAirdropData(toAddress, amount string, txStatus, count, airdropStatus uint) error {
	sql := fmt.Sprintf("INSERT INTO airdrop ( to_address, airdrop_amount, tx_status, airdrop_count, airdrop_status) VALUES ('%s', '%s', %d, %d, %d);",
		toAddress, amount, txStatus, count, airdropStatus)
	result, err := db.Exec(sql)
	if err != nil {
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if num != 1 {
		return fmt.Errorf("insert failed, rows affected num != 1")
	}
	return nil
}

func updateAirdropValue(toAddress, amount string, count uint) error {
	sql := fmt.Sprintf("UPDATE airdrop SET airdrop_amount = '%s', airdrop_count = %d WHERE to_address = '%s'",
		amount, count, toAddress)
	result, err := db.Exec(sql)
	if err != nil {
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		return  err
	}
	if num != 1 {
		return fmt.Errorf("UPDATE: RowsAffected != 1")
	}
	return nil
}

func updateAirdropHashTxStatusAndTime(toAddress, hash string, txStatus, airdropStatus uint, airdropTime int64) error {
	sql := fmt.Sprintf("UPDATE airdrop SET hash = '%s', tx_status = %d, airdrop_status = %d, airdrop_time = %d WHERE to_address = '%s'",
		hash, txStatus, airdropStatus, airdropTime, toAddress)
	result, err := db.Exec(sql)
	if err != nil {
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		return  err
	}
	if num != 1 {
		return fmt.Errorf("UPDATE: RowsAffected != 1")
	}
	return nil
}

func updateAirdropTxStatusAndAirdropStatus(txStatus, airdropStatus uint, hash string) error {
	sql := fmt.Sprintf("UPDATE airdrop SET tx_status = %d, airdrop_status = %d WHERE hash = '%s'",
		txStatus, airdropStatus, hash)
	result, err := db.Exec(sql)
	if err != nil {
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		return  err
	}
	if num != 1 {
		return fmt.Errorf("UPDATE: RowsAffected != 1")
	}
	return nil
}

func selectAirdropAddress(toAddress string) (SelectAirdropAddress, bool, error) {
	selectSql := fmt.Sprintf("SELECT to_address, hash, tx_status, airdrop_count, airdrop_status, airdrop_time FROM airdrop WHERE to_address = '%s';", toAddress)
	var s SelectAirdropAddress
	rows, err := db.Query(selectSql)
	if err != nil {
		return s, false, err
	}

	var i int = 0
	var flag bool = false
	for rows.Next() {
		//定义变量接收查询数据
		var address string
		var hash *string
		var txStatus, count, airdropStatus uint
		var airdropTime *int64
		err := rows.Scan(&address, &hash, &txStatus, &count, &airdropStatus, &airdropTime)
		if err != nil {
			return  s, false, err
		}
		s.toAddress = address
		s.hash = hash
		s.txStatus = txStatus
		s.count = count
		s.airdropStatus = airdropStatus
		s.airdropTime = airdropTime
		i++
		flag = true
	}
	rows.Close()

	if flag {
		if i != 1 {
			return s, false, fmt.Errorf("rows > 1")
		}
	}
	return s, flag, nil
}

func selectAirdropHash(hash string) (SelectAirdropHash, bool, error) {
	selectSql := fmt.Sprintf("SELECT hash, tx_status, airdrop_status, airdrop_time FROM airdrop WHERE hash = '%s';", hash)
	var s SelectAirdropHash
	rows, err := db.Query(selectSql)
	if err != nil {
		return s, false, err
	}

	var i int = 0
	var flag bool = false
	for rows.Next() {
		//定义变量接收查询数据
		var hash *string
		var txStatus, airdropStatus uint
		var airdropTime *int64
		err := rows.Scan(&hash, &txStatus, &airdropStatus, &airdropTime)
		if err != nil {
			return  s, false, err
		}

		s.hash = hash
		s.txStatus = txStatus
		s.airdropStatus = airdropStatus
		s.airdropTime = airdropTime
		i++
		flag = true
	}
	rows.Close()

	if flag {
		if i != 1 {
			return s, false, fmt.Errorf("rows > 1")
		}
	}
	return s, flag, nil
}

func setAirdropLimit(c *gin.Context) {
	var airdropLimit AirdropLimit

	if err := c.ShouldBindJSON(&airdropLimit); err == nil {
		if airdropLimit.Upper == "" || airdropLimit.Lower == "" {
			airdropLog("ERROR", "the parameter is empty")
			c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "the parameter is empty"})
			return
		}
		upper, err:= strToFloat(airdropLimit.Upper)
		if err != nil {
			airdropLog("ERROR", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
			return
		}
		lowe, err := strToFloat(airdropLimit.Lower)
		if err != nil {
			airdropLog("ERROR", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
			return
		}
		if upper <= lowe {
			airdropLog("ERROR", "parameter error, upper < lower")
			c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "parameter error, upper < lower"})
			return
		}

		airdropUpperLimit = airdropLimit.Upper
		airdropLowerLimit = airdropLimit.Lower

		airdropLog("INFO", "airdropUpperLimit: ", airdropLimit.Upper, "airdropLowerLimit: ", airdropLimit.Lower, ", setAirdropLimit success")
		c.JSON(http.StatusOK, gin.H{"status": 1, "message" : "success"})

	} else {
		airdropLog("ERROR", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": err.Error()})
	}
	return
}

func getAirdropLimit(c *gin.Context) {

	airdropLog("INFO", "airdropUpperLimit: ", airdropUpperLimit, "airdropLowerLimit: ", airdropLowerLimit,  ", getAirdropLimit success")
	c.JSON(http.StatusOK, gin.H{
		"status" : 1,
		"message" : "success",
		"upper" : airdropUpperLimit,
		"lower" : airdropLowerLimit,
	})
	return
}

func airdropCall(toAddress string, amount *big.Int) (string, error){

	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	var batch []common.Address
	batch = append(batch, common.HexToAddress(toAddress))

	tx, err := inviteStake.Airdrop(auth, batch, amount)
	if err != nil {
		return "", err
	}
	hash := tx.Hash().String()
	if hash == "" {
		return "", fmt.Errorf("tx hash is empty")
	}
	return hash, nil
}

func airdrop(c *gin.Context) {
	var airdrop AirdropParameter

	if err := c.ShouldBindJSON(&airdrop); err == nil {
		if airdrop.Address == "" || airdrop.Value == "" {
			airdropLog("ERROR", "the parameter is empty")
			c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "the parameter is empty"})
			return
		}

		toAddress, err := verificationAddress(airdrop.Address)
		if err != nil {
			airdropLog("ERROR", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
			return
		}

		s, flag, err := selectAirdropAddress(toAddress)
		if err != nil {
			airdropLog("ERROR", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
			return
		}
		var count uint = 1
		if flag {
			if s.txStatus == 1 {
				airdropLog("ERROR", "multiple airdrops")
				c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "multiple airdrops"})
				return
			}

			if s.airdropStatus == 1 {

				if s.airdropTime == nil || *s.airdropTime == 0 {
					airdropLog("ERROR", "the database airdrop_time field is empty")
					c.JSON(http.StatusBadRequest, gin.H{"status": 0,
						"message": "the database airdrop_time field is empty"})
					return
				}

				if (getUnixTimestamp() - (*s.airdropTime)) < 4 {
					airdropLog("ERROR", "the transaction of this account has not been submitted to the block, please try again later")
					c.JSON(http.StatusBadRequest, gin.H{"status": 0,
						"message": "the transaction of this account has not been submitted to the block, please try again later"})
					return
				}
				if s.hash == nil || *s.hash == "" {
					airdropLog("ERROR", "the database hash field is empty")
					c.JSON(http.StatusBadRequest, gin.H{"status": 0,
						"message": "the database hash field is empty"})
					return
				}
				receipt, err := getReceipt(*s.hash)
				if err != nil {
					airdropLog("ERROR", "invalid hash")
					c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "invalid hash"})
					return
				}
				if receipt == nil {
					airdropLog("ERROR", "try again later")
					c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "try again later"})
					return
				}

				if receipt.Status != 0 {
					airdropLog("ERROR", "multiple airdrops")
					err := updateAirdropTxStatusAndAirdropStatus(1, 3, *s.hash)
					if err != nil {
						airdropLog("ERROR", err.Error())
						c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
						return
					}
					c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "multiple airdrops"})
					return
				} else {
					err := updateAirdropTxStatusAndAirdropStatus(0, 2, *s.hash)
					if err != nil {
						airdropLog("ERROR", err.Error())
						c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
						return
					}
				}
			}
			if s.airdropStatus == 3 || s.airdropStatus == 4 {
				airdropLog("ERROR", "multiple airdrops")
				c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "multiple airdrops"})
				return
			}
			count = s.count + 1
			err := updateAirdropValue(toAddress, airdrop.Value, count)
			if err != nil {
				airdropLog("ERROR", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
				return
			}
		} else {
			err := insertIntoAirdropData(toAddress, airdrop.Value, 0, count, 0)
			if err != nil {
				airdropLog("ERROR", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
				return
			}
		}

		value, err := strToFloat(airdrop.Value)
		if err != nil {
			airdropLog("ERROR", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
			return
		}
		if airdropUpperLimit == "" || airdropLowerLimit == "" {
			airdropLog("ERROR", "please set airdrop limit")
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": "please set airdrop limit"})
			return
		}

		upper, err := strToFloat(airdropUpperLimit)
		if err != nil {
			airdropLog("ERROR", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
			return
		}
		lowe, err := strToFloat(airdropLowerLimit)
		if err != nil {
			airdropLog("ERROR", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
			return
		}

		if value > upper || value < lowe {
			airdropLog("ERROR", "airdrop quantity range error")
			c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "airdrop quantity range error"})
			return
		}

		valueWei, err := floatToBigInt(value)
		if err != nil {
			airdropLog("ERROR", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
			return
		}

		hash, err := airdropCall(airdrop.Address, valueWei)
		if err != nil {
			airdropLog("ERROR", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
			return
		}

		updateErr := updateAirdropHashTxStatusAndTime(toAddress, verificationHash(hash), 0, 1, getUnixTimestamp())
		if updateErr != nil {
			airdropLog("ERROR", updateErr.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": updateErr.Error()})
			return
		}
		airdropLog("INFO", "hash: ", hash, ", airdrop success")
		c.JSON(http.StatusOK, gin.H{"status" : 1, "message" : "success", "hash" : hash})
	} else {
		airdropLog("ERROR", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
	}
	return
}

func getTxStatus(c *gin.Context) {
	var hash TxHash

	if err := c.ShouldBindJSON(&hash); err == nil {
		if hash.Hash == "" {
			airdropLog("ERROR", "the parameter is empty")
			c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "the parameter is empty"})
			return
		}

		s, flag, err := selectAirdropHash(verificationHash(hash.Hash))
		if err != nil {
			airdropLog("ERROR", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": err.Error()})
			return
		}
		if flag {
			if s.airdropTime == nil || *s.airdropTime == 0{
				airdropLog("ERROR", "the database airdrop_time field is empty")
				c.JSON(http.StatusBadRequest, gin.H{"status": 0,
					"message": "the database airdrop_time field is empty"})
				return
			}

			if s.txStatus == 1 && s.airdropStatus == 3 {
				airdropLog("INFO", "hash: ", hash.Hash, ", getTxStatus success")
				c.JSON(http.StatusOK, gin.H{"status" : 1, "message" : "success", "hash" : hash.Hash})
				return
			}
		} else {
			airdropLog("ERROR", "invalid hash, ", "hash: ", hash.Hash)
			c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "invalid hash"})
			return
		}

		if (getUnixTimestamp() - (*s.airdropTime)) < 4 {
			airdropLog("ERROR", "the transaction of this account has not been submitted to the block, please try again later")
			c.JSON(http.StatusBadRequest, gin.H{"status": 0,
				"message": "the transaction of this account has not been submitted to the block, please try again later"})
			return
		}
		receipt, err := getReceipt(hash.Hash)
		if err != nil {
			airdropLog("ERROR", "invalid hash, ", "hash: ", hash.Hash, ", ", err.Error())
			err1 := updateAirdropTxStatusAndAirdropStatus(0, 2, *s.hash)
			if err1 != nil {
				airdropLog("ERROR", err1.Error())
				c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": err1.Error()})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "invalid hash"})
			return
		}
		if receipt == nil {
			airdropLog("ERROR", "try again later")
			err1 := updateAirdropTxStatusAndAirdropStatus(0, 2, *s.hash)
			if err1 != nil {
				airdropLog("ERROR", err1.Error())
				c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": err1.Error()})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "try again later"})
			return
		}

		if receipt.Status == 0 {
			airdropLog("ERROR", "failed status on chain")
			err := updateAirdropTxStatusAndAirdropStatus(0, 2, *s.hash)
			if err != nil {
				airdropLog("ERROR", err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": err.Error()})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": "failed status on chain"})
			return
		}
		err1 := updateAirdropTxStatusAndAirdropStatus(1, 3, *s.hash)
		if err1 != nil {
			airdropLog("ERROR", err1.Error())
			c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": err1.Error()})
			return
		}
		airdropLog("INFO", "hash: ", hash.Hash, ", getTxStatus success")
		c.JSON(http.StatusOK, gin.H{"status" : 1, "message" : "success", "hash" : hash.Hash})

	} else {
		airdropLog("ERROR", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
	}
	return
}

func router() {
	router := gin.Default()

	v1 := router.Group("/invite_stake_airdrop")
	{
		// set 空投数量范围
		v1.POST("/set_limit", setAirdropLimit)

		// get 空投数量范围
		v1.GET("/get_limit", getAirdropLimit)

		// airdrop
		v1.POST("/airdrop", airdrop)

		// TransactionReceipt
		v1.POST("/get_tx_status", getTxStatus)

		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	if serverPort == "" {
		serverPort = ":8080"
	}
	router.Run(serverPort)
}

func main() {
	router()
}