package payer

import (
	"fghpdf.me/thunes_homework/internal/pkg/common"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/payer"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func List(c *gin.Context) {
	thunesSvc := payer.NewServer(common.NewThunesClient())

	svc := NewServer(thunesSvc)
	payers, err := svc.List()
	if err != nil {
		log.Errorf("[payer][handler][List] error: %v\n", err)
		errRes := common.ERROR_LIST_PAYER
		c.JSON(http.StatusInternalServerError, errRes)
	}

	c.JSON(http.StatusOK, payers)
}
