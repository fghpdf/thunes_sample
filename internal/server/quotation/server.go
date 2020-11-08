package quotation

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/quotation"
)

type Server interface {
	Create(params *ViewCreateParams) (*ViewModel, error)
}

type server struct {
	thunesSvc quotation.Server
}

func NewServer(thunesSvc quotation.Server) Server {
	return &server{
		thunesSvc: thunesSvc,
	}
}

func (s *server) Create(params *ViewCreateParams) (*ViewModel, error) {
	createParams := &quotation.CreateParams{
		ExternalId:      params.ExternalId,
		PayerId:         params.PayerId,
		Mode:            params.Mode,
		TransactionType: params.TransactionType,
		Source:          params.Source,
		Destination:     params.Destination,
	}
	quotation, err := s.thunesSvc.Create(createParams)
	if err != nil {
		return nil, err
	}

	viewQuotation := &ViewModel{
		Id:              quotation.Id,
		ExternalId:      quotation.ExternalId,
		Payer:           quotation.Payer,
		Mode:            quotation.Mode,
		TransactionType: quotation.TransactionType,
		Source:          quotation.Source,
		Destination:     quotation.Destination,
		SentAmount:      quotation.SentAmount,
		WholesaleFxRate: quotation.WholesaleFxRate,
		Fee:             quotation.Fee,
		CreationDate:    quotation.CreationDate,
		ExpirationDate:  quotation.ExpirationDate,
	}

	return viewQuotation, nil
}
