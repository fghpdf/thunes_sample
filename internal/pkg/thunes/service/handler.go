package service

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

// List return a list of all services available to the caller
func (s *server) List(params *ListParams) (*[]Model, error) {
	url := s.client.BasicUrl + "/v2/money-transfer/services"

	response, err := s.client.Get(url, params)
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
