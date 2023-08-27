package golang

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainGin() {
	r := gin.Default()
	r.GET("/health/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.Run(":3001")
}
