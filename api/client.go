package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// PerformRequest sends a specified HTTP request
func PerformRequest(requestURL string, request *http.Request) ([]byte, error) {
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
