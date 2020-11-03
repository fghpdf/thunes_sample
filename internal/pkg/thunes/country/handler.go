package country

import (
	"encoding/json"
	"errors"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fmt"
	"net/http"
)

type Server interface {
	List(params *ListParams) (*[]Model, error)
}

type server struct {
	client *httpClient.AuthClient
}

func NewServer(client *httpClient.AuthClient) Server {
	return &server{client: client}
}

// List return a list of countries for all money transfer services available for the caller
func (s *server) List(params *ListParams) (*[]Model, error) {
	url := s.client.BasicUrl + "/v2/money-transfer/countries"

	response, err := s.client.Get(url, params)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorMsg := fmt.Sprintf("list countries request error, code: {%d}, message: {%s}",
			response.StatusCode, response.Status)
		return nil, errors.New(errorMsg)
	}

	countries := &[]Model{}
	err = json.NewDecoder(response.Body).Decode(countries)
	if err != nil {
		return nil, err
	}

	return countries, nil
}
