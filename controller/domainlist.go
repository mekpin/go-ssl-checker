package controller

import (
	"fmt"
	"go-ssl-checker/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DomainList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"domain1":       config.Common.Domain1,
		"domain2":       config.Common.Domain2,
		"domain3":       config.Common.Domain3,
		"domain4":       config.Common.Domain4,
		"domain5":       config.Common.Domain5,
		"alldomainport": config.Common.Alldomainport,
	})

	//debug
	fmt.Println(config.Common.Domain1)
	fmt.Println(config.Common.Domain2)
	fmt.Println(config.Common.Domain3)
	fmt.Println(config.Common.Domain4)
	fmt.Println(config.Common.Domain5)
	fmt.Println(config.Common.Alldomainport)

}
