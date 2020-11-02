package quotation

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Create return a new quotation for a given source or destination value.
func Create(client *httpClient.AuthClient, params *CreateParams) (*Model, error) {
	url := client.BasicUrl + "/v2/money-transfer/quotations"

	response, err := client.Post(url, params)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		errMsg := &httpClient.HttpErrorModel{}
		_ = json.NewDecoder(response.Body).Decode(errMsg)

		errorMsg := fmt.Sprintf("create quotation request error, code: {%d}, message: {%s}",
			response.StatusCode, errMsg.Errors)
		return nil, errors.New(errorMsg)
	}

	quotation := &Model{}
	err = json.NewDecoder(response.Body).Decode(quotation)
	if err != nil {
		return nil, err
	}

	return quotation, nil
}

// Get return information for a given quotation
func Get(client *httpClient.AuthClient, id int) (*Model, error) {
	url := fmt.Sprintf("%s/v2/money-transfer/quotations/%d", client.BasicUrl, id)

	response, err := client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errMsg := &httpClient.HttpErrorModel{}
		_ = json.NewDecoder(response.Body).Decode(errMsg)

		errorMsg := fmt.Sprintf("get quotation request error, code: {%d}, message: {%s}",
			response.StatusCode, errMsg.Errors)
		return nil, errors.New(errorMsg)
	}

	quotation := &Model{}
	err = json.NewDecoder(response.Body).Decode(quotation)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return quotation, nil
}
