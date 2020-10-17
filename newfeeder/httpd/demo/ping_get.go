package demo

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func PingGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"hello":"Found Me",
	})
}