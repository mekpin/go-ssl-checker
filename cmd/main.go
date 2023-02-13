package main

import (
	"fmt"
	"go-project-template/config"
	"go-project-template/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	router.Router(app)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Common.Port),
		Handler: app,
	}

	srv.ListenAndServe()

}
