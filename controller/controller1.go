package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "listtest is ok",
	})
}
