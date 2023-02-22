package controller

import (
	"fmt"
	"go-ssl-checker/service/core"
	"go-ssl-checker/service/manifest"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SSLCheck(c *gin.Context) {
	manifests, err := manifest.ParseInventory(false)
	if err != nil {
		fmt.Println("error while parsing inventory on ssl check in file controller/sslcheck.go")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error on controller/sslcheck.go function SSLCheck": err,
		})
		return
	}
	output := core.SSLExpireCheck(manifests)
	c.JSON(http.StatusOK, output)
}

func SSLList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ssl list is reachable",
	})
}
