package api

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
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

	jsonData, err := json.Marshal(txReq)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	if err != nil {
		return "", err
	}

	body, err := PerformRequest(url, req)

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
