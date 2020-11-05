package payer

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/payer"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/service"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/transactionType"
	mockPayer "fghpdf.me/thunes_homework/test/mocks/thunes/payer"
	"fmt"
	countryLab "github.com/biter777/countries"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPayers := make([]payer.Model, 0)
	for index := 1; index <= 2; index++ {
		p := payer.Model{
			Id:             uint64(index),
			Name:           fmt.Sprintf("Payer %d", index),
			Precision:      0,
			Increment:      "0.01",
			Currency:       countryLab.KEN.Currency().Alpha(),
			CountryIsoCode: countryLab.KEN.Alpha3(),
			Service: service.Model{
				Id:   1,
				Name: "MobileWallet",
			},
			TransactionTypes: transactionType.TypeModel{},
		}
		mockPayers = append(mockPayers, p)
	}

	expectedPayers := make([]ViewModel, 0)
	for _, value := range mockPayers {
		p := ViewModel{
			Id:               value.Id,
			Name:             value.Name,
			Currency:         value.Currency,
			CountryIsoCode:   value.CountryIsoCode,
			TransactionTypes: value.TransactionTypes,
		}
		expectedPayers = append(expectedPayers, p)
	}

	mockClient := mockPayer.NewMockServer(ctrl)
	mockClient.EXPECT().List(nil).Return(&mockPayers, nil)

	svc := NewServer(mockClient)
	actualPayers, err := svc.List()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expectedPayers, *actualPayers)
}
