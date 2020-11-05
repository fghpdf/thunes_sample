package payer

import (
	"encoding/json"
	"errors"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fmt"
	"net/http"
)

type Server interface {
	List(params *ListParams) (*[]Model, error)
	GetDetail(id uint64) (*Model, error)
	GetRate(id uint64) (*RateModel, error)
}

type server struct {
	client *httpClient.AuthClient
}

func NewServer(client *httpClient.AuthClient) Server {
	return &server{client: client}
}

// List return the list of available payers
func (s *server) List(params *ListParams) (*[]Model, error) {
	url := s.client.BasicUrl + "/v2/money-transfer/payers"

	response, err := s.client.Get(url, params)
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
func (s *server) GetDetail(id uint64) (*Model, error) {
	url := fmt.Sprintf("%s/v2/money-transfer/payers/%d", s.client.BasicUrl, id)

	response, err := s.client.Get(url, nil)
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

func (s *server) GetRate(id uint64) (*RateModel, error) {
	url := fmt.Sprintf("%s/v2/money-transfer/payers/%d/rates", s.client.BasicUrl, id)

	response, err := s.client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorMsg := fmt.Sprintf("get a payer rate request error, code: {%d}, message: {%s}",
			response.StatusCode, response.Status)
		return nil, errors.New(errorMsg)
	}

	rate := &RateModel{}
	err = json.NewDecoder(response.Body).Decode(rate)
	if err != nil {
		return nil, err
	}

	return rate, nil
}
