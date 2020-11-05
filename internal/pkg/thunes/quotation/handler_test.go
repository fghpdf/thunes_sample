package quotation

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/amount"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/common"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/payer"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	server := common.ServerMock("/v2/money-transfer/quotations", createSuccessMock)
	defer server.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: server.URL,
	}

	svc := NewServer(authClient)

	externalId, err := common.Generate()
	if err != nil {
		panic(err)
	}

	params := &CreateParams{
		ExternalId:      externalId,
		Mode:            common.SOURCE_AMOUNT,
		TransactionType: "C2C",
		PayerId:         1,
		Source: amount.SourceModel{
			CountryIsoCode: "FRA",
			BaseModel: amount.BaseModel{
				Amount:   "10",
				Currency: "EUR",
			},
		},
		Destination: amount.DestinationModel{
			BaseModel: amount.BaseModel{
				Amount:   "27",
				Currency: "CNY",
			},
		},
	}

	res, err := svc.Create(params)
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

func TestGet(t *testing.T) {
	quotationId := uint64(1)
	url := fmt.Sprintf("/v2/money-transfer/quotations/%d", quotationId)
	server := common.ServerMock(url, getSuccessMock)
	defer server.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: server.URL,
	}

	svc := NewServer(authClient)

	res, err := svc.Get(quotationId)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Errorf("expected a quotation but got nil\n")
	}

	if res.Id != quotationId {
		t.Errorf("expected the Id 1 but got %d\n", res.Id)
	}
}

func createSuccessMock(w http.ResponseWriter, r *http.Request) {
	quotation := mockQuotation()

	params := &CreateParams{}
	_ = json.NewDecoder(r.Body).Decode(params)

	quotation.Id = 1
	quotation.ExternalId = params.ExternalId
	quotation.Mode = params.Mode
	quotation.TransactionType = params.TransactionType
	quotation.Source = params.Source
	quotation.Destination = params.Destination

	res, _ := json.Marshal(quotation)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func getSuccessMock(w http.ResponseWriter, r *http.Request) {
	quotation := mockQuotation()

	path := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(path[len(path)-1])

	quotation.Id = uint64(id)

	res, _ := json.Marshal(quotation)
	_, _ = w.Write(res)
}

func mockQuotation() *Model {
	externalId, err := common.Generate()
	if err != nil {
		panic(err)
	}

	quotation := &Model{
		ExternalId: externalId,
		Payer: payer.Model{
			Id:             1,
			Name:           "Payer One",
			Precision:      2,
			Increment:      "0.01",
			Currency:       "CNY",
			CountryIsoCode: "CHN",
			Service: service.Model{
				Id:   1,
				Name: "MobileWallet",
			},
		},
		Mode:            common.SOURCE_AMOUNT,
		TransactionType: "C2C",
		Source: amount.SourceModel{
			CountryIsoCode: "FRA",
			BaseModel: amount.BaseModel{
				Amount:   "10",
				Currency: "EUR",
			},
		},
		Destination: amount.DestinationModel{
			BaseModel: amount.BaseModel{
				Amount:   "27",
				Currency: "CNY",
			},
		},
		SentAmount: amount.SentModel{
			BaseModel: amount.BaseModel{
				Amount:   "10",
				Currency: "EUR",
			},
		},
		WholesaleFxRate: "2.7",
		Fee: amount.FeeModel{
			BaseModel: amount.BaseModel{
				Amount:   "1.88",
				Currency: "EUR",
			},
		},
		CreationDate:   "2016-11-02T09:07:44",
		ExpirationDate: "2016-11-03T09:07:44",
	}

	return quotation
}
