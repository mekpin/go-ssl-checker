package controller

import (
	"fmt"
	"go-ssl-checker/config"
	"go-ssl-checker/service/manifest"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListManifest(c *gin.Context) {
	fmt.Println(config.Manifest.InventoryPath)

	list, err := manifest.ParseInventory(false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, list)
}
