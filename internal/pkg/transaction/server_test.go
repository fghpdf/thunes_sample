package transaction

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/creditParty"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/transaction"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/user"
	mockTransaction "fghpdf.me/thunes_homework/test/mocks/thunes/transaction"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	quotationId := uint64(77)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreatedTransaction := &transaction.Model{
		Id:                 uint64(1),
		Status:             "10000",
		StatusMessage:      "CREATED",
		StatusClass:        "CREATED",
		StatusClassMessage: "1",
		ExternalId:         "541411823484321405",
		TransactionType:    "C2C",
	}

	params := &transaction.CreateParams{
		CreditPartyIdentifier: creditParty.IdentifierModel{
			Msisdn: "123123100",
		},
		ExternalId: mockCreatedTransaction.ExternalId,
		Sender: user.SenderModel{
			BaseUserModel: user.BaseUserModel{
				Lastname: "Tom",
			},
		},
		Beneficiary: user.BeneficiaryModel{
			BaseUserModel: user.BaseUserModel{
				Lastname: "Jerry",
			},
		},
	}

	mockClient := mockTransaction.NewMockServer(ctrl)
	mockClient.EXPECT().Create(quotationId, params).Return(mockCreatedTransaction, nil)

	viewParams := &ViewCreateParams{
		QuotationId:           quotationId,
		CreditPartyIdentifier: params.CreditPartyIdentifier,
		Sender:                params.Sender,
		Beneficiary:           params.Beneficiary,
		ExternalId:            params.ExternalId,
	}

	svc := NewServer(mockClient)
	actual, err := svc.Create(viewParams)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, mockCreatedTransaction.ExternalId, actual.ExternalId)
}
