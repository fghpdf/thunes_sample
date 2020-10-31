package common

import (
	"net/http"
	"net/http/httptest"
)

func ServerMock(url string, mockHandler func(http.ResponseWriter, *http.Request)) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc(url, mockHandler)

	server := httptest.NewServer(handler)
	return server
}
