package ping

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/common"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"net/http"
	"testing"
)

func TestSend(t *testing.T) {
	service := common.ServerMock("/ping", pingResponseMock)
	defer service.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: service.URL,
	}

	svc := NewServer(authClient)

	res, err := svc.Send()
	if err != nil {
		t.Error(err)
	}

	if res.Status != "up" {
		t.Errorf("expected status up but got %v\n", res.Status)
	}
}

func pingResponseMock(w http.ResponseWriter, r *http.Request) {
	ping := &Model{Status: "up"}
	res, _ := json.Marshal(ping)
	_, _ = w.Write(res)
}
