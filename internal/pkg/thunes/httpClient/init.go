package httpClient

import (
	"errors"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/common"
	"io"
	"net/http"
)

var (
	username = ""
	password = ""
)

// set basic auth
func Init(key string, secret string) {
	username = key
	password = secret
}

func Do(method string, url string, payload io.Reader) (*http.Response, error)  {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	if username == "" || password == "" {
		return nil, errors.New(common.MUST_INIT_HTTP_CLIENT)
	}

	// Set the auth for the request.
	req.SetBasicAuth(username, password)

	return http.DefaultClient.Do(req)
}
