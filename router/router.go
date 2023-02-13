package router

import (
	"go-ssl-checker/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) {

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Accept, Content-Type, Content-Length, Authorization, X-CLIENT-TOKEN, access-control-allow-origin"}
	config.AllowCredentials = true

	app.Use(cors.New(config))
	app.GET("/", controller.BaseHealthcheck)

	app.GET("/health", controller.Healthcheck)
	app.GET("/domainlist", controller.DomainList)

	baseGroup := app.Group("/")

	checkGroup := baseGroup.Group("/check")
	{
		checkGroup.GET("/", controller.SSLCheck)
		checkGroup.GET("/list", controller.SSLList)
	}

}
