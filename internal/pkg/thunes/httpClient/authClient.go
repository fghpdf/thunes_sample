package httpClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
)

type AuthClient struct {
	Username string
	Password string
	BasicUrl string
}

func (authClient *AuthClient) Get(url string, payload interface{}) (*http.Response, error) {
	if payload == nil {
		return authClient.Do(http.MethodGet, url, nil)
	}

	// will output: "q=foo&all=true&page=2"
	urlQuery, err := query.Values(payload)
	if err != nil {
		return nil, err
	}

	url = fmt.Sprintf("%s?%s", url, urlQuery.Encode())
	return authClient.Do(http.MethodGet, url, nil)
}

func (authClient *AuthClient) Post(url string, payload interface{}) (*http.Response, error) {
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return authClient.Do(http.MethodPost, url, bytes.NewBuffer(bytePayload))
}

func (authClient *AuthClient) Do(method string, url string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Set the auth for the request.
	req.SetBasicAuth(authClient.Username, authClient.Password)

	return http.DefaultClient.Do(req)
}
