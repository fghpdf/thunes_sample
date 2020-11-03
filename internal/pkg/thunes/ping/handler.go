package ping

import (
	"encoding/json"
	"errors"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fmt"
	"net/http"
)

type Server interface {
	Send() (*Model, error)
}

type server struct {
	client *httpClient.AuthClient
}

func NewServer(client *httpClient.AuthClient) Server {
	return &server{client: client}
}

// Send ping check connectivity
// 200 + status: ok is connectivity
func (s *server) Send() (*Model, error) {
	url := s.client.BasicUrl + "/ping"
	response, err := s.client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorMsg := fmt.Sprintf("ping request error, code: {%d}, message: {%s}",
			response.StatusCode, response.Status)
		return nil, errors.New(errorMsg)
	}

	pingModel := &Model{}
	err = json.NewDecoder(response.Body).Decode(pingModel)
	if err != nil {
		return nil, err
	}
	return pingModel, nil
}
