package routers

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()

	r.Use(requestid.New())
	r.Use(Logging())
	r.Use(gin.Recovery())

	return r
}
