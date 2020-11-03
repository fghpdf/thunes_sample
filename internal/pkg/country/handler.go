package country

import (
	"fghpdf.me/thunes_homework/internal/pkg/common"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/country"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func List(c *gin.Context) {
	thunesSvc := country.NewServer(common.NewThunesClient())

	svc := NewServer(thunesSvc)
	countries, err := svc.List()
	if err != nil {
		log.Errorf("[country][handler][List] error: %v\n", err)
		errRes := common.ERROR_LIST_COUNTRY
		c.JSON(http.StatusInternalServerError, errRes)
	}

	c.JSON(http.StatusOK, countries)
}
