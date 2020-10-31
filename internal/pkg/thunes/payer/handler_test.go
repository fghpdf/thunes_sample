package payer

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/creditParty"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/service"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/transactionType"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {
	server := serverMock("/v2/money-transfer/payers", listSuccessMock)
	defer server.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: server.URL,
	}

	params := &ListParams{
		Page:           0,
		PerPage:        50,
		CountryIsoCode: "CNH",
		Currency:       "CNY",
	}

	res, err := List(authClient, params)
	if err != nil {
		t.Error(err)
	}

	if res == nil && len(*res) != 5 {
		t.Errorf("expected 5 payers but got %d\n", len(*res))
	}

	if (*res)[0].Id != 1 {
		t.Errorf("expected first payer id is 1 but got %d\n", (*res)[0].Id)
	}
}

func TestGetDetail(t *testing.T) {
	payerId := 1
	url := fmt.Sprintf("/v2/money-transfer/payers/%d", payerId)
	server := serverMock(url, getDetailMock)
	defer server.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: server.URL,
	}

	res, err := GetDetail(authClient, payerId)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Errorf("expected a payer but got nil\n")
	}

	if res.Id != payerId {
		t.Errorf("expected the payer id is 1 but got %d\n", res.Id)
	}
}

func serverMock(url string, mockHandler func(http.ResponseWriter, *http.Request)) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc(url, mockHandler)

	server := httptest.NewServer(handler)
	return server
}

func listSuccessMock(w http.ResponseWriter, r *http.Request) {
	payers := &[5]Model{}
	for index := 0; index < 5; index++ {
		payers[index] = Model{
			Id:             index + 1,
			Name:           fmt.Sprintf("Payer %d", index+1),
			Precision:      0,
			Increment:      "0.01",
			Currency:       "CNY",
			CountryIsoCode: "CNH",
			Service: service.Model{
				Id:   1,
				Name: "MobileWallet",
			},
			TransactionTypes: transactionType.Model{
				C2C: transactionType.InfoModel{
					MinimumTransactionAmount:       0,
					MaximumTransactionAmount:       100,
					CreditPartyIdentifiersAccepted: [][]string{{"msisdn"}},
					RequiredSendingEntityFields:    nil,
					RequiredReceivingEntityFields:  nil,
					RequiredDocuments:              nil,
					CreditPartyInformation:         creditParty.InfoModel{},
					CreditPartyVerification:        creditParty.VerificationModel{},
				},
			},
		}
	}

	res, _ := json.Marshal(payers)
	_, _ = w.Write(res)
}

func getDetailMock(w http.ResponseWriter, r *http.Request) {
	payer := &Model{
		Id:             1,
		Name:           "Payer One",
		Precision:      0,
		Increment:      "0.01",
		Currency:       "CNY",
		CountryIsoCode: "CNH",
		Service: service.Model{
			Id:   1,
			Name: "MobileWallet",
		},
		TransactionTypes: transactionType.Model{
			C2C: transactionType.InfoModel{
				MinimumTransactionAmount:       0,
				MaximumTransactionAmount:       100,
				CreditPartyIdentifiersAccepted: [][]string{{"msisdn"}},
				RequiredSendingEntityFields:    nil,
				RequiredReceivingEntityFields:  nil,
				RequiredDocuments:              nil,
				CreditPartyInformation:         creditParty.InfoModel{},
				CreditPartyVerification:        creditParty.VerificationModel{},
			},
		},
	}

	res, _ := json.Marshal(payer)
	_, _ = w.Write(res)
}