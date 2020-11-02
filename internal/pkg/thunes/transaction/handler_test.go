package transaction

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/common"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/creditParty"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/user"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	quotationId := 1
	url := fmt.Sprintf("/v2/money-transfer/quotations/%d/transactions", quotationId)
	server := common.ServerMock(url, createSuccessMock)
	defer server.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: server.URL,
	}

	externalId, err := common.Generate()
	if err != nil {
		panic(err)
	}

	params := &CreateParams{
		CreditPartyIdentifier: creditParty.IdentifierModel{
			Msisdn: "+263775892100",
		},
		RetailFee:         "1",
		RetailFeeCurrency: "EUR",
		Sender: user.SenderModel{
			BaseUserModel: user.BaseUserModel{
				Lastname:                  "Doe",
				NationalityCountryIsoCode: "FRA",
				Gender:                    "MALE",
				IdNumber:                  "502-42-0158",
			},
		},
		Beneficiary: user.BeneficiaryModel{
			BaseUserModel: user.BaseUserModel{
				Lastname:       "Jane",
				CountryIsoCode: "ZWE",
				Occupation:     "Sales Executive",
			},
		},
		ExternalId: externalId,
	}

	res, err := Create(authClient, quotationId, params)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Errorf("expected a quotation but got nil\n")
	}

	if res.ExternalId != externalId {
		t.Errorf("expected given quotation external id but got %s\n", res.ExternalId)
	}
}

func TestConfirm(t *testing.T) {
	transactionId := uint64(1)
	url := fmt.Sprintf("/v2/money-transfer/transactions/%d/confirm", transactionId)
	server := common.ServerMock(url, confirmSuccessMock)
	defer server.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: server.URL,
	}

	res, err := Confirm(authClient, transactionId)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Errorf("expected a quotation but got nil\n")
	}

	if res.Id != transactionId {
		t.Errorf("expected given id but got %d\n", res.Id)
	}
}

func createSuccessMock(w http.ResponseWriter, r *http.Request) {
	transaction := &Model{}

	params := &CreateParams{}
	_ = json.NewDecoder(r.Body).Decode(params)

	transaction.ExternalId = params.ExternalId
	transaction.Sender = params.Sender
	transaction.Beneficiary = params.Beneficiary
	transaction.CreditPartyIdentifier = params.CreditPartyIdentifier

	transaction.Id = 1

	res, _ := json.Marshal(transaction)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func confirmSuccessMock(w http.ResponseWriter, r *http.Request) {
	transaction := &Model{}

	path := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(path[len(path)-2])

	transaction.Id = uint64(id)

	res, _ := json.Marshal(transaction)
	_, _ = w.Write(res)
}
