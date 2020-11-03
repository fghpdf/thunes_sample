package country

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/country"
	mockCountry "fghpdf.me/thunes_homework/test/mocks/thunes/country"
	countryLab "github.com/biter777/countries"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCountries := &[]country.Model{
		{
			Name:    countryLab.Japan.String(),
			IsoCode: countryLab.Japan.Alpha3(),
		},
		{
			Name:    countryLab.China.String(),
			IsoCode: countryLab.China.Alpha3(),
		},
		{
			Name:    countryLab.KEN.String(),
			IsoCode: countryLab.KEN.Alpha3(),
		},
	}

	expectedCountries := &[]ViewModel{
		{
			Name:     countryLab.Japan.String(),
			Currency: countryLab.Japan.Currency().Alpha(),
			Flag:     countryLab.Japan.Emoji(),
		},
		{
			Name:     countryLab.China.String(),
			Currency: countryLab.China.Currency().Alpha(),
			Flag:     countryLab.China.Emoji(),
		},
		{
			Name:     countryLab.KEN.String(),
			Currency: countryLab.KEN.Currency().Alpha(),
			Flag:     countryLab.KEN.Emoji(),
		},
	}

	mockClient := mockCountry.NewMockServer(ctrl)
	mockClient.EXPECT().List(nil).Return(mockCountries, nil)

	svc := NewServer(mockClient)
	actualCountries, err := svc.List()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expectedCountries, actualCountries)
}
