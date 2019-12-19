package api

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type sendTxRequest struct {
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Value     string `json:"value"`
	Data      string `json:"data"`
	Nonce     uint64 `json:"nonce"`
	GasPrice  uint64 `json:"gasPrice"`
	GasLimit  uint64 `json:"gasLimit"`
	Signature string `json:"signature"`
}

type sendTxResponse struct {
	TxHash string `json:"txHash"`
	Error  string `json:"error,omitempty"`
}

// Account contains the current data for a specific wallet or account
type Account struct {
	Address string `json:"address"`
	Nonce   uint64 `json:"nonce"`
	Balance string `json:"balance"`
	//Code     string
	//CodeHash []byte
	//RootHash []byte
}

// AccountWrapper is simple wrapper type to help with deserializing the address response
type AccountWrapper struct {
	Account Account `json:"account"`
}

// SendTransaction performs the actual HTTP request to send the transaction
func SendTransaction(
	nonce uint64,
	sender string,
	receiver string,
	value string,
	gasPrice uint64,
	gasLimit uint64,
	data string,
	signature []byte,
	apiHost string) (string, error) {

	url := fmt.Sprintf("%s/transaction/send", apiHost)
	hexSignature := hex.EncodeToString(signature)

	txReq := sendTxRequest{
		Sender:    sender,
		Receiver:  receiver,
		Value:     value,
		Data:      data,
		Nonce:     nonce,
		GasPrice:  gasPrice,
		GasLimit:  gasLimit,
		Signature: hexSignature,
	}

	jsonData, _ := json.Marshal(txReq)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	body, err := performRequest(url, req)

	if err != nil {
		return "", err
	}

	var response sendTxResponse
	json.Unmarshal([]byte(body), &response)

	if response.TxHash == "" {
		return "", errors.New(response.Error)
	}

	return response.TxHash, nil
}

// GetAccount fetches the desired account's balance as well as nonce
func GetAccount(address string, apiHost string) (Account, error) {
	url := fmt.Sprintf("%s/address/%s", apiHost, address)
	req, err := http.NewRequest("GET", url, nil)

	var response AccountWrapper
	var accountResponse Account

	body, _ := performRequest(url, req)

	if err != nil {
		return accountResponse, err
	}

	json.Unmarshal([]byte(body), &response)
	accountResponse = response.Account

	return accountResponse, nil
}

func performRequest(requestURL string, request *http.Request) ([]byte, error) {
	client := &http.Client{}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("request to url %s failed", requestURL)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response from url %s", requestURL)
	}

	defer resp.Body.Close()

	return body, err
}
