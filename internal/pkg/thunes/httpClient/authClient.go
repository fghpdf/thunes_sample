package httpClient

import (
	"io"
	"net/http"
)

type AuthClient struct {
	Username string
	Password string
	BasicUrl string
}

func (authClient *AuthClient) Do(method string, url string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	// Set the auth for the request.
	req.SetBasicAuth(authClient.Username, authClient.Password)

	return http.DefaultClient.Do(req)
}
