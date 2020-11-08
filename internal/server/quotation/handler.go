package quotation

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/quotation"
	"fghpdf.me/thunes_homework/internal/server/common"
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
		errRes := common.ERROR_CREATE_QUOTATION_BIND
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	createdQuotation, err := svc.Create(&params)
	if err != nil {
		log.Errorf("[quotation][handler][Create] error: %v\n", err)
		errRes := common.ERROR_CREATE_QUOTATION
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	c.JSON(http.StatusCreated, createdQuotation)
}
