package ping

import (
	"encoding/json"
	"errors"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fmt"
	"net/http"
)

// Send ping check connectivity
// 200 + status: ok is connectivity
func Send(authClient *httpClient.AuthClient) (*Model, error) {
	url := authClient.BasicUrl + "/ping"
	response, err := authClient.Get(url, nil)
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
