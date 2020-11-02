package transaction

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

// Create return a new transaction with transfer values specified from a given quotation.
func Create(client *httpClient.AuthClient, quotationId int, params *CreateParams) (*Model, error) {
	url := fmt.Sprintf("%s/v2/money-transfer/quotations/%d/transactions", client.BasicUrl, quotationId)

	response, err := client.Post(url, params)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		errMsg := &httpClient.HttpErrorModel{}
		_ = json.NewDecoder(response.Body).Decode(errMsg)

		errorMsg := fmt.Sprintf("create transaction request error, code: {%d}, message: {%s}",
			response.StatusCode, errMsg.Errors)
		return nil, errors.New(errorMsg)
	}

	transaction := &Model{}
	err = json.NewDecoder(response.Body).Decode(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Confirm a previously-created transaction to initiate processing.
// return a given transaction
func Confirm(client *httpClient.AuthClient, transactionId uint64) (*Model, error) {
	url := fmt.Sprintf("%s/v2/money-transfer/transactions/%d/confirm", client.BasicUrl, transactionId)

	response, err := client.Post(url, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errMsg := &httpClient.HttpErrorModel{}
		_ = json.NewDecoder(response.Body).Decode(errMsg)

		errorMsg := fmt.Sprintf("confirm transaction request error, code: {%d}, message: {%s}",
			response.StatusCode, errMsg.Errors)
		return nil, errors.New(errorMsg)
	}

	transaction := &Model{}
	err = json.NewDecoder(response.Body).Decode(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
