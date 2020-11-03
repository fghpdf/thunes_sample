package country

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/country"
	countryLab "github.com/biter777/countries"
	log "github.com/sirupsen/logrus"
)

type Server interface {
	List() (*[]ViewModel, error)
}

type server struct {
	thunesSvc country.Server
}

func NewServer(thunesSvc country.Server) Server {
	return &server{
		thunesSvc: thunesSvc,
	}
}

func (s *server) List() (*[]ViewModel, error) {
	countries, err := s.thunesSvc.List(nil)
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
			Currency: info.Currency().Alpha(),
			Flag:     info.Emoji(),
		}

		viewCountries = append(viewCountries, viewCountry)
	}

	return &viewCountries, nil
}
