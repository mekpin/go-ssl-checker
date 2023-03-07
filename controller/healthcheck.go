package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "go-ssl-checker service is alive",
	})
}

func BaseHealthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "base path is OK",
	})
}
