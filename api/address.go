package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Account contains the current data for a specific wallet or account
type Account struct {
	Address string `json:"address,omitempty"`
	Nonce   uint64 `json:"nonce,omitempty"`
	Balance string `json:"balance"`
	//Code     string
	//CodeHash []byte
	//RootHash []byte
}

// AccountWrapper is simple wrapper type to help with deserializing the address response
type AccountWrapper struct {
	Account Account `json:"account"`
}

// GetAccount fetches the desired account's balance as well as nonce
func GetAccount(address string, apiHost string) (Account, error) {
	url := fmt.Sprintf("%s/address/%s", apiHost, address)
	req, err := http.NewRequest("GET", url, nil)

	var response AccountWrapper
	var accountResponse Account

	body, err := PerformRequest(url, req)

	if err != nil {
		return accountResponse, err
	}

	json.Unmarshal([]byte(body), &response)
	accountResponse = response.Account

	return accountResponse, nil
}

// GetBalance fetches the balance of a specific account
func GetBalance(address string, apiHost string) (Account, error) {
	url := fmt.Sprintf("%s/address/%s/balance", apiHost, address)
	req, err := http.NewRequest("GET", url, nil)

	var accountResponse Account

	if err != nil {
		return accountResponse, err
	}

	body, err := PerformRequest(url, req)

	if err != nil {
		return accountResponse, err
	}

	json.Unmarshal([]byte(body), &accountResponse)

	return accountResponse, nil
}
