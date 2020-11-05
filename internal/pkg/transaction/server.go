package transaction

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/transaction"
)

type Server interface {
	Create(params *ViewCreateParams) (*ViewModel, error)
}

type server struct {
	thunesSvc transaction.Server
}

func NewServer(thunesSvc transaction.Server) Server {
	return &server{
		thunesSvc: thunesSvc,
	}
}

func (s *server) Create(params *ViewCreateParams) (*ViewModel, error) {
	createParams := &transaction.CreateParams{
		CreditPartyIdentifier: params.CreditPartyIdentifier,
		Sender:                params.Sender,
		Beneficiary:           params.Beneficiary,
		ExternalId:            params.ExternalId,
	}

	transferResult, err := s.thunesSvc.Create(params.QuotationId, createParams)
	if err != nil {
		return nil, err
	}

	viewResult := &ViewModel{
		Id:                 transferResult.Id,
		Status:             transferResult.Status,
		StatusMessage:      transferResult.StatusMessage,
		StatusClass:        transferResult.StatusClass,
		StatusClassMessage: transferResult.StatusClassMessage,
		ExternalId:         transferResult.ExternalId,
		ExternalCode:       transferResult.ExternalCode,
		TransactionType:    transferResult.TransactionType,
		CreationDate:       transferResult.CreationDate,
		ExpirationDate:     transferResult.ExpirationDate,
	}

	return viewResult, nil
}
