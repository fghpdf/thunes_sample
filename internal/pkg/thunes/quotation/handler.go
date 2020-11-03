package quotation

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server interface {
	Create(params *CreateParams) (*Model, error)
	Get(id int) (*Model, error)
}

type server struct {
	client *httpClient.AuthClient
}

func NewServer(client *httpClient.AuthClient) Server {
	return &server{client: client}
}

// Create return a new quotation for a given source or destination value.
func (s *server) Create(params *CreateParams) (*Model, error) {
	url := s.client.BasicUrl + "/v2/money-transfer/quotations"

	response, err := s.client.Post(url, params)
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
func (s *server) Get(id int) (*Model, error) {
	url := fmt.Sprintf("%s/v2/money-transfer/quotations/%d", s.client.BasicUrl, id)

	response, err := s.client.Get(url, nil)
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
