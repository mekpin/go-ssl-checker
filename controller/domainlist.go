package controller

import (
	"fmt"
	"go-ssl-checker/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DomainList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"alldomainport": config.Common.Alldomainport,
	})

	//debug
	fmt.Println(config.Common.Alldomainport)

}
