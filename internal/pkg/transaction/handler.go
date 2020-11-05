package transaction

import (
	"fghpdf.me/thunes_homework/internal/pkg/common"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/transaction"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Create(c *gin.Context) {
	thunesSvc := transaction.NewServer(common.NewThunesClient())

	svc := NewServer(thunesSvc)

	var params ViewCreateParams
	err := c.ShouldBind(&params)
	if err != nil {
		log.Errorf("[transaction][handler][Create] bind params error: %v\n", err)
		errRes := common.ERROR_CREATE_TRANSACTION_BIND
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	var urlParams ViewCreateUrlParams
	err = c.ShouldBindUri(&urlParams)
	if err != nil {
		log.Errorf("[transaction][handler][Create] bind params error: %v\n", err)
		errRes := common.ERROR_CREATE_TRANSACTION_BIND
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	createdTransaction, err := svc.Create(urlParams.QuotationId, &params)
	if err != nil {
		log.Errorf("[transaction][handler][Create] error: %v\n", err)
		errRes := common.ERROR_CREATE_TRANSACTION
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	c.JSON(http.StatusCreated, createdTransaction)
}

func Confirm(c *gin.Context) {
	thunesSvc := transaction.NewServer(common.NewThunesClient())

	svc := NewServer(thunesSvc)

	var urlParams ViewConfirmUrlParams
	err := c.ShouldBindUri(&urlParams)
	if err != nil {
		log.Errorf("[transaction][handler][Confirm] bind params error: %v\n", err)
		errRes := common.ERROR_CREATE_TRANSACTION_BIND
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	createdTransaction, err := svc.Confirm(urlParams.TransactionId)
	if err != nil {
		log.Errorf("[transaction][handler][Confirm] error: %v\n", err)
		errRes := common.ERROR_CREATE_TRANSACTION
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	c.JSON(http.StatusOK, createdTransaction)
}
