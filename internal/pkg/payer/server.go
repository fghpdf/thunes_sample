package payer

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/payer"
)

type Server interface {
	List() (*[]ViewModel, error)
}

type server struct {
	thunesSvc payer.Server
}

func NewServer(thunesSvc payer.Server) Server {
	return &server{
		thunesSvc: thunesSvc,
	}
}

func (s *server) List() (*[]ViewModel, error) {
	payers, err := s.thunesSvc.List(nil)
	if err != nil {
		return nil, err
	}

	viewPayers := make([]ViewModel, 0)
	for _, p := range *payers {
		viewPayer := ViewModel{
			Id:               p.Id,
			Name:             p.Name,
			Currency:         p.Currency,
			CountryIsoCode:   p.CountryIsoCode,
			TransactionTypes: p.TransactionTypes,
		}

		viewPayers = append(viewPayers, viewPayer)
	}

	return &viewPayers, nil
}
