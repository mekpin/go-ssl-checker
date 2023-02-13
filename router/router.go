package router

import (
	"go-project-template/controller"

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

	baseGroup := app.Group("/")

	checkGroup := baseGroup.Group("/check")
	{
		checkGroup.GET("/list", controller.ListTest)
	}

}
