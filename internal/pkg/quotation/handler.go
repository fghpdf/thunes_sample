package quotation

import (
	"fghpdf.me/thunes_homework/internal/pkg/common"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/quotation"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Create(c *gin.Context) {
	thunesSvc := quotation.NewServer(common.NewThunesClient())

	svc := NewServer(thunesSvc)

	var params ViewCreateParams
	err := c.ShouldBind(&params)
	if err != nil {
		log.Errorf("[quotation][handler][Create] bind params error: %v\n", err)
		errRes := common.ERROR_LIST_PAYER
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	createdQuotation, err := svc.Create(&params)
	if err != nil {
		log.Errorf("[quotation][handler][Create] error: %v\n", err)
		errRes := common.ERROR_LIST_PAYER
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	c.JSON(http.StatusCreated, createdQuotation)
}
