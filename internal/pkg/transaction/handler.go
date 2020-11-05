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
	createdQuotation, err := svc.Create(&params)
	if err != nil {
		log.Errorf("[transaction][handler][Create] error: %v\n", err)
		errRes := common.ERROR_CREATE_TRANSACTION
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	c.JSON(http.StatusCreated, createdQuotation)
}
