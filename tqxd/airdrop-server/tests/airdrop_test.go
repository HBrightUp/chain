package tests

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)
type AirdropParameter struct {
	Address string
	Value string
}

type TxHash struct {
	Hash string
}

func toBytes(v interface{}) (*bytes.Buffer, int, error){
	b, err := json.Marshal(v)
	if err != nil {
		return nil, 0, err
	}
	buf := new(bytes.Buffer)
	n, err := buf.Write(b)
	if err != nil {
		return nil, n, err
	}
	return buf, n, err
}

type success struct {
	Hash string
	Message string
	Status int
}

func airdrop(address, value string) string {
	var requestUrl string = "http://127.0.0.1:8080/invite_stake_airdrop/airdrop"

	a := AirdropParameter{
		address,
		value,
	}
	buf, _, err := toBytes(a)
	if err != nil {
		return ""
	}

	request, err := http.Post(requestUrl, "text/json", buf)
	if err != nil {
		return ""
	}

	defer request.Body.Close()

	if request.StatusCode == 200 {
		rb, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return ""
		}
		var s success
		err1 := json.Unmarshal(rb, &s)
		if err1 != nil {
			return ""
		}
		return s.Hash
	}
	return ""
}

func getTxStatus(hash string) {
	var requestUrl string = "http://127.0.0.1:8080/invite_stake_airdrop/get_tx_status"
	h := TxHash{
		Hash: hash,
	}
	buf, _, err := toBytes(h)
	if err != nil {
		return
	}
	client := &http.Client{}
	requestGet, _:= http.NewRequest("POST", requestUrl, buf)

	request, err := client.Do(requestGet)
	if err != nil {
		return
	}
	defer request.Body.Close()
}

func createAccount() (string, error){

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", err
	}

	//privateKeyBytes := crypto.FromECDSA(privateKey)
	//fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("cannot convert")
	}
	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address, nil
}

func TestAirdrop(t *testing.T) {

	airdropPool := make(map[string]string)

	var count int = 20000

	for i := 0; true; i++ {
		addr, err := createAccount()
		if err != nil {
			continue
		}
		airdropPool[addr] = "0.0011"
		if len(airdropPool) == count {
			break
		}
	}
	t.Log("len(airdropPool): ", len(airdropPool))

	var hash []string
	for k, v :=range airdropPool {
		h := airdrop(k, v)
		if h == "" {
			continue
		}
		hash = append(hash, h)
		time.Sleep(time.Duration(100) * time.Millisecond)
	}

	t.Log("len(hash): ", len(hash))


	time.Sleep(time.Duration(5) * time.Second)

	for _, h := range hash {
		getTxStatus(h)
	}

}

func TestGetTxStatus(t *testing.T) {
	getTxStatus("0x1782e00c63abb72b285fda7766e33daa8f0f8bfacb22d2d193ab2a73d3a428b2")
}

func TestCreateAccount(t *testing.T) {
	createAccount()
}

func t1() (string, error) {
	return "t1", fmt.Errorf("t1")
}

func t2() (string, error) {
	return "t2", fmt.Errorf("t2")
}

func TestErr(t *testing.T)  {

	vt1, err := t1()
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(vt1)
	var vt2 string
	vt2, err = t2()
	if err != nil {
		t.Log(err)
	}
	t.Log(vt2)

	vt3, err := t2()
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(vt3)

}