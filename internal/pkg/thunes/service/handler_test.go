package service

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/common"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"net/http"
	"testing"
)

func TestList(t *testing.T) {
	server := common.ServerMock("/v2/money-transfer/services", listSuccessMock)
	defer server.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: server.URL,
	}

	params := &ListParams{
		PageParams: common.PageParams{
			Page:    0,
			PerPage: 50,
		},
		CountryIsoCode: "CNH",
	}

	res, err := List(authClient, params)
	if err != nil {
		t.Error(err)
	}

	if res == nil || len(*res) != 3 {
		t.Errorf("expected 5 payers but got %d\n", len(*res))
	}

	if (*res)[0].Id != 1 {
		t.Errorf("expected first payer id is 1 but got %d\n", (*res)[0].Id)
	}
}

func listSuccessMock(w http.ResponseWriter, r *http.Request) {
	services := &[3]Model{
		{
			Name: "MobileWallet",
			Id:   1,
		},
		{
			Name: "BankAccount",
			Id:   2,
		},
		{
			Name: "CashPickup",
			Id:   3,
		},
	}

	res, _ := json.Marshal(services)
	_, _ = w.Write(res)
}
