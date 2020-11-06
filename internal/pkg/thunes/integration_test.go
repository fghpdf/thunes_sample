//+build integration

package thunes

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/amount"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/common"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/creditParty"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/payer"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/quotation"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/transaction"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/user"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"testing"
	"time"
)

const (
	SUCCESS_MSISDN               = "+263775892100"
	CURRENTLY_UNAVAILABLE_MSISDN = "+263775892117"
	BARRED_BENEFICIARY_MSISDN    = "+263775892104"
	LIMITATIONS_MSISDN           = "+263775892111"

	SUCCESS_STATUS               = "70000"
	CURRENTLY_UNAVAILABLE_STATUS = "90400"
	BARRED_BENEFICIARY_STATUS    = "90201"
	LIMITATIONS_STATUS           = "90305"
)

func TestSuccessFlow(t *testing.T) {
	testFlow(t, SUCCESS_MSISDN, SUCCESS_STATUS)
}

func TestDeclinedPayerCurrentlyUnavailable(t *testing.T) {
	testFlow(t, CURRENTLY_UNAVAILABLE_MSISDN, CURRENTLY_UNAVAILABLE_STATUS)
}

func TestBarredBeneficiary(t *testing.T) {
	testFlow(t, BARRED_BENEFICIARY_MSISDN, BARRED_BENEFICIARY_STATUS)
}

func TestLimitations(t *testing.T) {
	testFlow(t, LIMITATIONS_MSISDN, LIMITATIONS_STATUS)
}

func testFlow(t *testing.T, msisdn string, expectedStatus string) {
	client := newAuthClient()

	// Step One: list all payers
	payers, err := listPayers(client)
	if err != nil {
		t.Error(err)
	}

	// Step two: choose payer 0 and create quotation
	quotation, err := createQuotation(client, &(*payers)[0])
	if err != nil {
		t.Error(err)
	}

	// Step three: create transaction and use .xxx117 msisdn
	transaction, err := createTransaction(client, quotation, msisdn)
	if err != nil {
		t.Error(err)
	}

	// Step four: confirm transaction
	confirmedTransaction, err := confirmTransaction(client, transaction.Id)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, confirmedTransaction.Id, transaction.Id)
	assert.Equal(t, confirmedTransaction.Status, "20000")

	// wait for completed
	time.Sleep(6 * time.Second)

	// Step five: get transaction status
	finalTransaction, err := getTransaction(client, transaction.Id)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, finalTransaction.Id, transaction.Id)
	assert.Equal(t, finalTransaction.Status, expectedStatus)
}

func newAuthClient() *httpClient.AuthClient {
	viper.BindEnv("thunes.APIKey", "THUNES_APIKEY")
	viper.BindEnv("thunes.APISecret", "THUNES_APISECRET")

	return &httpClient.AuthClient{
		Username: viper.GetString("thunes.APIKey"),
		Password: viper.GetString("thunes.APISecret"),
		BasicUrl: "https://api-mt.pre.thunes.com",
	}
}

func listPayers(client *httpClient.AuthClient) (*[]payer.Model, error) {
	params := &payer.ListParams{
		PageParams: common.PageParams{
			Page:    0,
			PerPage: 50,
		},
		CountryIsoCode: "KEN", // sure exist
	}
	svc := payer.NewServer(client)
	return svc.List(params)
}

func createQuotation(client *httpClient.AuthClient, payer *payer.Model) (*quotation.Model, error) {
	externalId, err := common.Generate()
	if err != nil {
		return nil, err
	}
	params := &quotation.CreateParams{
		ExternalId:      externalId,
		PayerId:         payer.Id,
		Mode:            common.SOURCE_AMOUNT,
		TransactionType: "C2C",
		Source: amount.SourceModel{
			CountryIsoCode: "SGP",
			BaseModel: amount.BaseModel{
				Amount:   "90",
				Currency: "SGD",
			},
		},
		Destination: amount.DestinationModel{
			BaseModel: amount.BaseModel{
				Currency: payer.Currency,
			},
		},
	}

	svc := quotation.NewServer(client)
	return svc.Create(params)
}

func createTransaction(client *httpClient.AuthClient, quotation *quotation.Model, msisdn string) (*transaction.Model, error) {
	externalId, err := common.Generate()
	if err != nil {
		return nil, err
	}
	params := &transaction.CreateParams{
		CreditPartyIdentifier: creditParty.IdentifierModel{
			Msisdn: msisdn,
		},
		Sender: user.SenderModel{
			BaseUserModel: user.BaseUserModel{
				Lastname: "Qu",
			},
		},
		Beneficiary: user.BeneficiaryModel{
			BaseUserModel: user.BaseUserModel{
				Lastname: "Luo",
			},
		},
		ExternalId: externalId,
	}

	svc := transaction.NewServer(client)

	return svc.Create(quotation.Id, params)
}

func confirmTransaction(client *httpClient.AuthClient, id uint64) (*transaction.Model, error) {
	svc := transaction.NewServer(client)
	return svc.Confirm(id)
}

func getTransaction(client *httpClient.AuthClient, id uint64) (*transaction.Model, error) {
	svc := transaction.NewServer(client)
	return svc.Get(id)
}
