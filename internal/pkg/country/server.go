package country

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/country"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	countryLab "github.com/biter777/countries"
	log "github.com/sirupsen/logrus"
)

type Server interface {
	List() (*[]ViewModel, error)
}

type server struct {
	client *httpClient.AuthClient
}

func NewServer(client *httpClient.AuthClient) Server {
	return &server{
		client: client,
	}
}

func (s *server) List() (*[]ViewModel, error) {
	countries, err := country.List(s.client, nil)
	if err != nil {
		return nil, err
	}

	viewCountries := make([]ViewModel, 0)
	for _, country := range *countries {
		info := countryLab.ByName(country.Name)
		if info == countryLab.Unknown {
			log.Errorf("[country][List]can not find the country %s\n", country.Name)
			continue
		}

		viewCountry := ViewModel{
			Name:     country.Name,
			Currency: country.IsoCode,
			Flag:     info.Emoji(),
		}

		viewCountries = append(viewCountries, viewCountry)
	}

	return &viewCountries, nil
}
