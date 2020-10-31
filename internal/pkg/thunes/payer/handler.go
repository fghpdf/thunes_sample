package payer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fmt"
	"net/http"
)

// List return the list of available payers
func List(client *httpClient.AuthClient, params *ListParams) (*[]Model, error) {
	url := client.BasicUrl + "/v2/money-transfer/payers"
	byteParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(http.MethodGet, url, bytes.NewReader(byteParams))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorMsg := fmt.Sprintf("list payers request error, code: {%d}, message: {%s}",
			response.StatusCode, response.Status)
		return nil, errors.New(errorMsg)
	}

	payers := &[]Model{}
	err = json.NewDecoder(response.Body).Decode(payers)
	if err != nil {
		return nil, err
	}

	return payers, nil
}

// GetDetail return a payer detail for given id
func GetDetail(client *httpClient.AuthClient, id int) (*Model, error) {
	url := fmt.Sprintf("%s/v2/money-transfer/payers/%d", client.BasicUrl, id)

	response, err := client.Do(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorMsg := fmt.Sprintf("get a payer detail request error, code: {%d}, message: {%s}",
			response.StatusCode, response.Status)
		return nil, errors.New(errorMsg)
	}

	payer := &Model{}
	err = json.NewDecoder(response.Body).Decode(payer)
	if err != nil {
		return nil, err
	}

	return payer, nil
}
