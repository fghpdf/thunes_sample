package quotation

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/amount"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/payer"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/quotation"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/service"
	mockQuotation "fghpdf.me/thunes_homework/test/mocks/thunes/quotation"
	countryLab "github.com/biter777/countries"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreatedQuotation := &quotation.Model{
		Id:         uint64(1),
		ExternalId: "213412342134",
		Payer: payer.Model{
			Id:             uint64(1),
			Name:           "Payer 1",
			Precision:      0,
			Increment:      "0.01",
			Currency:       countryLab.KEN.Currency().Alpha(),
			CountryIsoCode: countryLab.KEN.Alpha3(),
			Service: service.Model{
				Id:   1,
				Name: "MobileWallet",
			},
		},
		Mode:            "SOURCE_AMOUNT",
		TransactionType: "C2C",
		Source: amount.SourceModel{
			BaseModel: amount.BaseModel{
				Amount:   "75",
				Currency: "SGD",
			},
			CountryIsoCode: "SGP",
		},
		Destination: amount.DestinationModel{
			BaseModel: amount.BaseModel{
				Amount:   "5982.83",
				Currency: "KES",
			},
		},
		SentAmount: amount.SentModel{
			BaseModel: amount.BaseModel{
				Amount:   "75",
				Currency: "SGD",
			},
		},
		WholesaleFxRate: "79.77106237030",
		Fee: amount.FeeModel{
			BaseModel: amount.BaseModel{
				Amount:   "3",
				Currency: "SGD",
			},
		},
		CreationDate:   "2020-11-05T06:34:26Z",
		ExpirationDate: "2020-11-05T07:34:26Z",
	}

	params := &quotation.CreateParams{
		ExternalId:      mockCreatedQuotation.ExternalId,
		PayerId:         mockCreatedQuotation.Payer.Id,
		Mode:            mockCreatedQuotation.Mode,
		TransactionType: mockCreatedQuotation.TransactionType,
		Source:          mockCreatedQuotation.Source,
		Destination:     mockCreatedQuotation.Destination,
	}

	mockClient := mockQuotation.NewMockServer(ctrl)
	mockClient.EXPECT().Create(params).Return(mockCreatedQuotation, nil)

	viewParams := &ViewCreateParams{
		ExternalId:      params.ExternalId,
		PayerId:         params.PayerId,
		Mode:            params.Mode,
		TransactionType: params.TransactionType,
		Source:          params.Source,
		Destination:     params.Destination,
	}
	svc := NewServer(mockClient)
	actual, err := svc.Create(viewParams)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, mockCreatedQuotation.ExternalId, actual.ExternalId)
}
