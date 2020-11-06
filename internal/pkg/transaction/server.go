package transaction

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/transaction"
)

type Server interface {
	Create(quotationId uint64, params *ViewCreateParams) (*ViewModel, error)
	Confirm(id uint64) (*ViewModel, error)
	Get(id uint64) (*ViewModel, error)
}

type server struct {
	thunesSvc transaction.Server
}

func NewServer(thunesSvc transaction.Server) Server {
	return &server{
		thunesSvc: thunesSvc,
	}
}

func (s *server) Create(quotationId uint64, params *ViewCreateParams) (*ViewModel, error) {
	createParams := &transaction.CreateParams{
		CreditPartyIdentifier: params.CreditPartyIdentifier,
		Sender:                params.Sender,
		Beneficiary:           params.Beneficiary,
		ExternalId:            params.ExternalId,
	}

	transferResult, err := s.thunesSvc.Create(quotationId, createParams)
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

func (s *server) Confirm(id uint64) (*ViewModel, error) {
	transferResult, err := s.thunesSvc.Confirm(id)
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

func (s *server) Get(id uint64) (*ViewModel, error) {
	transferResult, err := s.thunesSvc.Get(id)
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
