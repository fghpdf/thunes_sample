package ping

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSend(t *testing.T) {
	service := serverMock()
	defer service.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: service.URL,
	}

	res, err := Send(authClient)
	if err != nil {
		t.Error(err)
	}

	if res.Status != "up" {
		t.Errorf("expected status up but got %v\n", res.Status)
	}
}

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/ping", pingResponseMock)

	service := httptest.NewServer(handler)
	return service
}

func pingResponseMock(w http.ResponseWriter, r *http.Request) {
	ping := &Model{Status: "up"}
	res, _ := json.Marshal(ping)
	_, _ = w.Write(res)
}
