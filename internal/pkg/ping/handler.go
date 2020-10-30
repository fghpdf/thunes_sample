package ping

import "github.com/gin-gonic/gin"

func Handler(c *gin.Context) {
	res := &Ping{Status: "up"}
	c.JSON(200, res)
}
