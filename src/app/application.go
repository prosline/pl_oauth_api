package app

import (
	"github.com/gin-gonic/gin"
	"github.com/prosline/pl_logger/logger"
	"github.com/prosline/pl_oauth_api/src/http"
	"github.com/prosline/pl_oauth_api/src/repository/db"
	"github.com/prosline/pl_oauth_api/src/repository/rest"
	access_token2 "github.com/prosline/pl_oauth_api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token2.NewService(rest.NewRepository(), db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.PUT("/oauth/access_token", atHandler.UpdateExpiration)
	logger.Info("Starting Application on Port 8080....")
	router.Run(":8080")
}
