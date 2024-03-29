package payer

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/common"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/creditParty"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/service"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/transactionType"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	server := common.ServerMock("/v2/money-transfer/payers", listSuccessMock)
	defer server.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: server.URL,
	}

	svc := NewServer(authClient)

	params := &ListParams{
		PageParams: common.PageParams{
			Page:    0,
			PerPage: 50,
		},
		CountryIsoCode: "CNH",
		Currency:       "CNY",
	}

	res, err := svc.List(params)
	if err != nil {
		t.Error(err)
	}

	if res == nil || len(*res) != 5 {
		t.Errorf("expected 5 payers but got %d\n", len(*res))
	}

	if (*res)[0].Id != 1 {
		t.Errorf("expected first payer id is 1 but got %d\n", (*res)[0].Id)
	}
}

func TestGetDetail(t *testing.T) {
	payerId := uint64(1)
	url := fmt.Sprintf("/v2/money-transfer/payers/%d", payerId)
	server := common.ServerMock(url, getDetailSuccessMock)
	defer server.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: server.URL,
	}

	svc := NewServer(authClient)

	res, err := svc.GetDetail(payerId)
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

func TestGetRate(t *testing.T) {
	payerId := uint64(1)
	url := fmt.Sprintf("/v2/money-transfer/payers/%d/rates", payerId)
	server := common.ServerMock(url, getRateSuccessMock)
	defer server.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: server.URL,
	}

	svc := NewServer(authClient)

	res, err := svc.GetRate(payerId)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Errorf("expected a payer but got nil\n")
	}

	if res.DestinationCurrency != "KEN" {
		t.Errorf("expected KEN but got %s\n", res.DestinationCurrency)
	}
}

func listSuccessMock(w http.ResponseWriter, r *http.Request) {
	payers := &[5]Model{}
	for index := 0; index < 5; index++ {
		payers[index] = Model{
			Id:             uint64(index + 1),
			Name:           fmt.Sprintf("Payer %d", index+1),
			Precision:      0,
			Increment:      "0.01",
			Currency:       "CNY",
			CountryIsoCode: "CNH",
			Service: service.Model{
				Id:   1,
				Name: "MobileWallet",
			},
			TransactionTypes: transactionType.TypeModel{
				C2C: &transactionType.InfoModel{
					MinimumTransactionAmount:       "0",
					MaximumTransactionAmount:       "100",
					CreditPartyIdentifiersAccepted: [][]string{{"msisdn"}},
					RequiredSendingEntityFields:    nil,
					RequiredReceivingEntityFields:  nil,
					RequiredDocuments:              nil,
					CreditPartyInformation:         &creditParty.InfoModel{},
					CreditPartyVerification:        &creditParty.VerificationModel{},
				},
			},
		}
	}

	res, _ := json.Marshal(payers)
	_, _ = w.Write(res)
}

func getDetailSuccessMock(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(path[len(path)-1])

	payer := &Model{
		Id:             uint64(id),
		Name:           "Payer One",
		Precision:      0,
		Increment:      "0.01",
		Currency:       "CNY",
		CountryIsoCode: "CNH",
		Service: service.Model{
			Id:   1,
			Name: "MobileWallet",
		},
		TransactionTypes: transactionType.TypeModel{
			C2C: &transactionType.InfoModel{
				MinimumTransactionAmount:       "0",
				MaximumTransactionAmount:       "100",
				CreditPartyIdentifiersAccepted: [][]string{{"msisdn"}},
				RequiredSendingEntityFields:    nil,
				RequiredReceivingEntityFields:  nil,
				RequiredDocuments:              nil,
				CreditPartyInformation:         &creditParty.InfoModel{},
				CreditPartyVerification:        &creditParty.VerificationModel{},
			},
		},
	}

	res, _ := json.Marshal(payer)
	_, _ = w.Write(res)
}

func getRateSuccessMock(w http.ResponseWriter, r *http.Request) {
	rate := &RateModel{
		DestinationCurrency: "KEN",
		Rates: RateInfoModel{
			SGD: &[]RateDetailModel{
				{
					SourceAmountMax: "100",
					SourceAmountMin: "0",
					WholesaleFxRate: "1.7",
				},
			},
		},
	}

	res, _ := json.Marshal(rate)
	_, _ = w.Write(res)
}
