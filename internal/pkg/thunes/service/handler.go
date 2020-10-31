package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fmt"
	"net/http"
)

// List return a list of all services available to the caller
func List(client *httpClient.AuthClient, params *ListParams) (*[]Model, error) {
	url := client.BasicUrl + "/v2/money-transfer/services"
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
		errorMsg := fmt.Sprintf("list services request error, code: {%d}, message: {%s}",
			response.StatusCode, response.Status)
		return nil, errors.New(errorMsg)
	}

	services := &[]Model{}
	err = json.NewDecoder(response.Body).Decode(services)
	if err != nil {
		return nil, err
	}

	return services, nil
}
